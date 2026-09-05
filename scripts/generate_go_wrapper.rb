#!/usr/bin/env ruby
# frozen_string_literal: true

# Builds go/wenmar/operations.gen.go and go/wenmar/models.gen.go from
# spec/operations.json. These are the public wrapper methods and type aliases.
#
# Usage: ruby scripts/generate_go_wrapper.rb [manifest]
#   manifest defaults to spec/operations.json

require "json"

MANIFEST_PATH = ARGV[0] || "spec/operations.json"
MANIFEST = JSON.parse(File.read(MANIFEST_PATH))
OPS = MANIFEST["operations"]

# Map an operation id (snake_case) to a Go exported method/type name.
def pascal(identifier)
  identifier.split("_").reject(&:empty?).map { |p| p[0].upcase + p[1..] }.join
end

# Convert a snake_case path parameter name to a camelCase Go identifier.
def go_param_name(name)
  parts = name.split("_")
  (parts.shift || "id") + parts.map { |p| p[0].upcase + p[1..] }.join
end

def path_param_types(op)
  types = {}
  op["pathParams"].each do |p|
    t = (op["pathParamTypes"] || {})[p] || "integer"
    types[go_param_name(p)] = t == "integer" ? "int" : "string"
  end
  types
end

def path_param_signature(op)
  types = path_param_types(op)
  op["pathParams"].map { |p| "#{go_param_name(p)} #{types[go_param_name(p)]}" }.join(", ")
end

def path_param_args(op)
  op["pathParams"].map { |p| go_param_name(p) }.join(", ")
end

def has_query_params?(op)
  !(op["queryParams"] || []).empty?
end

def param_struct(op)
  "#{pascal(op["id"])}Params"
end

# Whether an operation gets a typed ListResult wrapper (and thus its raw
# method is renamed to ListXxxRaw). Only paginated list_* operations with a
# known item schema qualify.
def typed_list?(op)
  op["paginated"] && op["responseSchema"] && op["id"].start_with?("list_")
end

# Build a single operation method. Each wrapper returns the generated raw
# response struct (aliased into the wenmar package by models.gen.go) and
# centralizes hook firing + error parsing. For paginated list operations the
# method is named ListXxxRaw; the typed ListXxx is emitted by build_typed_list.
def build_operation(op)
  m = pascal(op["id"])
  method_name = typed_list?(op) ? "#{m}Raw" : m
  body_type = op["requestSchema"] # wenmar-alias name
  params = []
  params << path_param_signature(op) unless op["pathParams"].empty?
  params << "params *#{param_struct(op)}" if has_query_params?(op)
  params << "body #{body_type}" if body_type

  gen_fn = "#{m}WithResponse"
  gen_args = op["pathParams"].map { |p| go_param_name(p) }
  gen_args << "params" if has_query_params?(op)
  gen_args << "body" if body_type

  # Multipart uploads (e.g. import validate) have no named request schema and
  # oapi-codegen only emits a `WithBodyWithResponse(ctx, contentType, body
  # io.Reader, ...)` variant — there is no no-arg `WithResponse`. Route these
  # through the WithBody variant so the wrapper compiles.
  if op["requestContentType"] && op["requestContentType"] != "application/json"
    gen_fn = "#{m}WithBodyWithResponse"
    gen_args << '"multipart/form-data"'
    gen_args << "body"
    params << "body io.Reader"
  end

  call = "c.gen.#{gen_fn}(ctx#{gen_args.empty? ? "" : ", " + gen_args.join(", ")})"

  <<~GO
    // #{method_name} runs the #{op["id"]} operation (#{op["method"].upcase} #{op["path"]}).
    func (c *Client) #{method_name}(ctx context.Context#{params.empty? ? "" : ", " + params.join(", ")}) (*#{m}Response, error) {
    	ctx = c.hooks.OnOperationStart(ctx, OperationInfo{Operation: "#{m}"})
    	resp, err := #{call}
    	if err != nil {
    		c.hooks.OnOperationEnd(ctx, OperationInfo{Operation: "#{m}"}, OperationResult{Operation: "#{m}", Err: err})
    		return nil, err
    	}
    	if resp.StatusCode() >= 400 {
    		perr := parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
    		c.hooks.OnOperationEnd(ctx, OperationInfo{Operation: "#{m}"}, OperationResult{Operation: "#{m}", Err: perr})
    		return nil, perr
    	}
    	c.hooks.OnOperationEnd(ctx, OperationInfo{Operation: "#{m}"}, OperationResult{Operation: "#{m}"})
    	return resp, nil
    }
  GO
end

def build_operation_bodies
  OPS.map { |op| build_operation(op) }.join("\n")
end

# Build a typed ListXxx method that returns *ListResult[ItemType] for
# paginated operations. The raw method is ListXxxRaw (from build_operation).
def build_typed_list(op)
  return "" unless typed_list?(op)

  m = pascal(op["id"])
  item = op["responseSchema"]
  raw_method = "#{m}Raw"

  params = ["ctx context.Context"]
  params << path_param_signature(op) unless op["pathParams"].empty?
  params << "params *#{param_struct(op)}" if has_query_params?(op)

  call_args = ["ctx"]
  call_args += op["pathParams"].map { |p| go_param_name(p) }
  call_args << "params" if has_query_params?(op)

  <<~GO
    // #{m} fetches a page of #{op["id"]} results as a typed ListResult.
    // Use .Items for the current page, .HasNext() / .Next(ctx) to paginate.
    func (c *Client) #{m}(#{params.join(", ")}) (*ListResult[#{item}], error) {
    	resp, err := c.#{raw_method}(#{call_args.join(", ")})
    	if err != nil {
    		return nil, err
    	}
    	return newListResultFromResponse[#{item}](resp.Body, resp.HTTPResponse.Header, c), nil
    }
  GO
end

def build_typed_list_bodies
  OPS.map { |op| build_typed_list(op) }.join("\n")
end

# Operations that return binary content (no JSON200 field in the response).
# These get a Download* method that returns *DownloadResult with a
# streaming body instead of buffering into []byte.
BINARY_DOWNLOAD_OPS = [
  "list_customers_export_download",
  "list_expenses_export_download",
  "list_tire_storage_slots_export_download",
  "list_tires_export_download",
]

def build_download(op)
  return "" unless BINARY_DOWNLOAD_OPS.include?(op["id"])

  m = pascal(op["id"])
  download_name = m.sub(/\AList/, "Download").sub(/Download\z/, "")
  params = ["ctx context.Context"]
  params << path_param_signature(op) unless op["pathParams"].empty?
  params << "params *#{param_struct(op)}" if has_query_params?(op)

  path_params = op["pathParams"].map { |p| go_param_name(p) }

  if path_params.empty?
    <<~GO
      // #{download_name} downloads #{op["id"]} as a streaming body.
      // The caller must close result.Body when done.
      func (c *Client) #{download_name}(#{params.join(", ")}) (*DownloadResult, error) {
      	return c.downloadPath(ctx, "#{op["path"]}")
      }
    GO
  else
    <<~GO
      // #{download_name} downloads #{op["id"]} as a streaming body.
      // The caller must close result.Body when done.
      func (c *Client) #{download_name}(#{params.join(", ")}) (*DownloadResult, error) {
      	return c.downloadPathWithParams(ctx, "#{op["path"]}", #{path_params.join(", ")})
      }
    GO
  end
end

def build_download_bodies
  OPS.map { |op| build_download(op) }.join("\n")
end

# Build GetAll helpers for paginated list operations. They collect items from
# the first page plus every Link-header page, capped via GetAllOptions
# (default 1000 items).
def build_get_all(op)
  return "" unless typed_list?(op)

  item = op["responseSchema"] # wenmar-alias of the item type
  # Guard: if the operation path is a sub-resource (deeper than the
  # tag's root resource), the responseSchema may not match the parent
  # resource type. Warn but don't block — the manifest should be fixed.
  path_segments = op["path"].split("/").reject(&:empty?)
  tag = op["tag"]
  if path_segments.size > 2 && item == pascal(tag.sub(/s\z/, ""))
    $stderr.puts "WARNING: #{op["id"]} has responseSchema=#{item} but path has #{path_segments.size} segments — verify this is correct"
  end

  list_method = pascal(op["id"])
  base = list_method.sub(/\AList/, "")
  params = ["ctx context.Context"]
  params << path_param_signature(op) unless op["pathParams"].empty?
  params << "params *#{param_struct(op)}" if has_query_params?(op)
  params << "opts *GetAllOptions"

  call_args = ["ctx"]
  call_args += op["pathParams"].map { |p| go_param_name(p) }
  call_args << "params" if has_query_params?(op)

  <<~GO
    // GetAll#{base} auto-paginates #{op["id"]}, following the Link header.
    // Pass opts.MaxItems or opts.MaxPages to cap collection (default 1000).
    func (c *Client) GetAll#{base}(#{params.join(", ")}) ([]#{item}, error) {
    	first, err := c.#{list_method}(#{call_args.join(", ")})
    	if err != nil {
    		return nil, err
    	}
    	items, _, err := getAll[#{item}](ctx, first, opts)
    	if err != nil {
    		return nil, err
    	}
    	return items, nil
    }
  GO
end

def build_get_all_bodies
  OPS.map { |op| build_get_all(op) }.join("\n")
end

def build_models
  item_types = OPS.map { |op| op["responseSchema"] }.compact.uniq
  response_types = OPS.map { |op| pascal(op["id"]) + "Response" }.uniq
  request_types = OPS.map { |op| op["requestSchema"] }.compact.uniq
  param_types = OPS.filter_map { |op| param_struct(op) if has_query_params?(op) }.uniq

  groups = []
  groups << (item_types + response_types)
  groups << request_types
  groups << param_types

  groups.filter_map do |types|
    next if types.empty?

    "type (\n" + types.map { |t| "\t#{t} = gen.#{t}" }.join("\n") + "\n)"
  end.join("\n\n")
end

uses_multipart = OPS.any? { |op| op["requestContentType"] && op["requestContentType"] != "application/json" }

operations_header = <<~'GO'
	// Code generated by scripts/generate_go_wrapper.rb from spec/operations.json. DO NOT EDIT.
	//go:build !skip_generate

	package wenmar

	import (
		"context"
		"io"
	)
GO

operations_header = operations_header.sub("\t\"io\"\n", "") unless uses_multipart

models_header = <<~'GO'
	// Code generated by scripts/generate_go_wrapper.rb from spec/operations.json. DO NOT EDIT.

	package wenmar

	import gen "github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
GO

File.write("go/wenmar/operations.gen.go", operations_header + "\n" + build_operation_bodies + "\n" + build_typed_list_bodies + "\n" + build_get_all_bodies + "\n" + build_download_bodies)
File.write("go/wenmar/models.gen.go", models_header + "\n" + build_models + "\n")
puts "Wrote go/wenmar/operations.gen.go and go/wenmar/models.gen.go (#{OPS.size} operations)"
