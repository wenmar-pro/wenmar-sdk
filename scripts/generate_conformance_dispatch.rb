#!/usr/bin/env ruby
# frozen_string_literal: true

# Builds conformance/go/dispatch.gen.go and conformance/ruby/dispatch.gen.rb
# from spec/operations.json. These map every manifest operation (and any test
# aliases) to a callable that executes the operation against a client.
#
# The dispatch lambdas receive `args` which contains "pathParams", "query",
# and "requestBody" maps (string keys). They reconstruct the SDK call.
#
# Usage: ruby scripts/generate_conformance_dispatch.rb [manifest]
#   manifest defaults to spec/operations.json

require "json"

MANIFEST_PATH = ARGV[0] || "spec/operations.json"
MANIFEST = JSON.parse(File.read(MANIFEST_PATH))
OPS = MANIFEST["operations"]

# Test operations that aren't plain manifest IDs but reference a manifest
# operation. Value is the manifest operation id they exercise.
TEST_ALIASES = {
  "list_customer_vehicles" => "list_customers_vehicles",
  "list_customer_work_orders" => "list_customers_work_orders",
  "list_vehicle_work_orders" => "list_vehicles_work_orders",
  "list_customers_with_params" => "list_customers",
  "list_customers_with_params_paginated" => "list_customers",
  "list_customers_paginated" => "list_customers",
  "list_work_orders_paginated" => "list_work_orders",
  "check_duplicate" => "check_vehicle_duplicate"
}.freeze

def pascal(identifier)
  identifier.split("_").reject(&:empty?).map { |p| p[0].upcase + p[1..] }.join
end

def go_param_name(name)
  parts = name.split("_")
  (parts.shift || "id") + parts.map { |p| p[0].upcase + p[1..] }.join
end

def op_by_id(id)
  OPS.find { |o| o["id"] == id }
end

# Go: expression to pull a path param from args.
def go_path_param(pp)
  "intArg(args[\"pathParams\"], \"#{pp}\")"
end

# Go: expression to pull a query param.
def go_query_expr(op)
  return "nil" if (op["queryParams"] || []).empty?

  q = op["queryParams"].map do |p|
    "#{pascal(p["name"])}: ptrVal(args[\"query\"], \"#{p["name"]}\")"
  end.join(", ")
  "&#{pascal(op["id"])}Params{#{q}}"
end

# Go: expression to build a request body value from args["requestBody"].
def go_body_expr(op)
  return "nil" unless op["requestSchema"]

  shape = op["requestShape"]
  schema = op["requestSchema"]
  if shape && shape["wrapper"]
    wrapper = shape["wrapper"]
    # Build the nested wrapper struct.
    "buildWrapper[#{schema}](#{wrapper.inspect}, args[\"requestBody\"])"
  elsif shape && shape["flat"]
    fields = shape["flat"].map do |f|
      "#{pascal(f)}: flatVal(args[\"requestBody\"], #{f.inspect})"
    end.join(", ")
    "#{schema}{#{fields}}"
  else
    # no shape info: pass the request body through as the schema type
    "#{schema}{}" # empty; conformance rarely exercises these
  end
end

# Build a single Go dispatch entry.
def build_go_entry(key, manifest_id, variant = :default)
  op = op_by_id(manifest_id)
  return nil unless op

  m = pascal(manifest_id)
  path_params = op["pathParams"]

  args = []
  args += path_params.map { |p| go_path_param(p) }
  args << go_query_expr(op) unless (op["queryParams"] || []).empty?
  body_expr = go_body_expr(op)
  args << body_expr if op["requestSchema"]

  call = "c.#{m}(ctx#{args.empty? ? "" : ", " + args.join(", ")})"

  case variant
  when :paginated
    body = <<~GO
      "KEY": func(ctx context.Context, t *testing.T, c *wenmar.Client, args map[string]interface{}) (interface{}, error) {
      	resp, err := CALL
      	if err != nil {
      		return nil, err
      	}
      	return paginateBody(c, resp, 5)
      },
    GO
  when :with_params
    body = <<~GO
      "KEY": func(ctx context.Context, t *testing.T, c *wenmar.Client, args map[string]interface{}) (interface{}, error) {
      	resp, err := CALL
      	if err != nil {
      		return nil, err
      	}
      	return decodeBody(resp.Body)
      },
    GO
  else
    body = <<~GO
      "KEY": func(ctx context.Context, t *testing.T, c *wenmar.Client, args map[string]interface{}) (interface{}, error) {
      	resp, err := CALL
      	if err != nil {
      		return nil, err
      	}
      	return decodeBody(resp.Body)
      },
    GO
  end

  body.gsub("KEY", key).gsub("CALL", call)
end

go_dispatch = +""
ruby_dispatch = +""
go_all = []
ruby_all = []

OPS.each do |op|
  key = op["id"]
  entry = build_go_entry(key, key)
  next unless entry

  go_dispatch << entry << "\n"
  go_all << key
  ruby_dispatch << %Q{    "#{key}" => ->(client, args) { client.#{key}(buildArgs(client, args)) },\n}
  ruby_all << key
end

# Alias entries used only by the conformance tests.
{
  "list_customers_paginated" => [:default, :paginated],
  "list_work_orders_paginated" => [:default, :paginated],
  "list_customers_with_params_paginated" => [:default, :paginated],
  "list_customers_with_params" => [:default, :with_params]
}.each do |key, (manifest_id, variant)|
  entry = build_go_entry(key, manifest_id, variant)
  next unless entry

  go_dispatch << entry << "\n"
  go_all << key
end

# Simple aliases that map to a different manifest id (no variant).
TEST_ALIASES.each do |key, manifest_id|
  next if go_all.include?(key)

  entry = build_go_entry(key, manifest_id)
  next unless entry

  go_dispatch << entry << "\n"
  go_all << key
  ruby_dispatch << %Q{    "#{key}" => ->(client, args) { client.#{manifest_id}(buildArgs(client, args)) },\n}
  ruby_all << key
end

# Ensure Ruby dispatch includes aliases even if already added.
ruby_all.uniq!
go_all.uniq!

go_ops_array = go_all.map { |k| "\t#{k.inspect}" }.join(",\n")
ruby_ops_array = ruby_all.map { |k| %Q{    #{k.inspect}} }.join(",\n")

go_file = <<~GO
	// Code generated by scripts/generate_conformance_dispatch.rb from spec/operations.json. DO NOT EDIT.
	package conformance

	import (
		"context"
		"testing"

		wenmar "github.com/wenmar-pro/wenmar-sdk/go/wenmar"
	)

	type operationFunc func(ctx context.Context, t *testing.T, client *wenmar.Client, args map[string]interface{}) (interface{}, error)

	var dispatch = map[string]operationFunc{
	#{go_dispatch}
	}

	// allOperations lists every operation covered by the dispatch map.
	var allOperations = []string{
	#{go_ops_array}
	}
GO

ruby_file = <<~RUBY
	# frozen_string_literal: true
	# AUTO-GENERATED by scripts/generate_conformance_dispatch.rb from spec/operations.json. DO NOT EDIT.

	module Conformance
	  DISPATCH = {
	#{ruby_dispatch}
	  }.freeze

	  ALL_OPERATIONS = [
	#{ruby_ops_array}
	  ].freeze
	end
RUBY

File.write("conformance/go/dispatch.gen.go", go_file)
File.write("conformance/ruby/dispatch.gen.rb", ruby_file)
puts "Wrote conformance/go/dispatch.gen.go and conformance/ruby/dispatch.gen.rb"
