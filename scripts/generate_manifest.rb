#!/usr/bin/env ruby
# frozen_string_literal: true

# Builds spec/operations.json from the enriched spec. This language-agnostic
# manifest drives the Go wrapper, Ruby resource, and conformance dispatch
# generators plus the parity CI gate.
#
# Usage: ruby scripts/generate_manifest.rb [enriched_spec] [output]
#   enriched_spec defaults to spec/openapi.enriched.yaml
#   output        defaults to spec/operations.json

require "yaml"
require "json"

ENRICHED_PATH = ARGV[0] || "spec/openapi.enriched.yaml"
OUT_PATH = ARGV[1] || "spec/operations.json"

# Stable manifest version. This is a schema/format version, not a date, so
# the generated file is byte-stable across runs and the drift check in
# `make check` does not fail on a new day.
MANIFEST_VERSION = "2026-08-30"

SPEC = YAML.load_file(ENRICHED_PATH, aliases: true)

def normalize_tag(tag)
  tag.downcase.gsub(/[^a-z0-9]+/, "_").gsub(/_+/, "_").sub(/_+$/, "")
end

def type_of(param)
  t = param.dig("schema", "type") || param["type"]
  return "string" if t.nil?

  t
end

# A list is paginated when it returns a bare array whose next page is
# advertised via the `Link` header (`rel="next"`). The manifest records these
# so SDK generators can expose pagination + GetAll helpers.
def paginated?(method, operation)
  return false unless method == "get"

  responses = operation["responses"] || {}
  ok = responses["200"] || responses["201"]
  return false unless ok

  schema = ok.dig("content", "application/json", "schema")
  schema && schema["type"] == "array"
end

def response_schema(operation)
  ok = operation.dig("responses", "200") || operation.dig("responses", "201")
  schema = ok&.dig("content", "application/json", "schema")
  return nil unless schema

  ref = schema["$ref"] || schema.dig("items", "$ref")
  ref&.split("/")&.last
end

# Whether the success (2xx) response carries an application/json body at all.
# Distinct from response_schema: some ops return an inline object (anonymous
# JSON200) that has no component name.
def response_body?(operation)
  ok = operation.dig("responses", "200") || operation.dig("responses", "201")
  !ok&.dig("content", "application/json").nil?
end

# Describes how the request body should be passed from SDK callers.
#   {"wrapper" => "customer"} -> pass { "customer" => attrs }
#   {"flat" => ["f1", "f2"]}  -> pass keyword args
#   {"wrapper" => nil}        -> pass the generated struct through unchanged
def request_shape(schema_name)
  return nil unless schema_name

  schema = SPEC.dig("components", "schemas", schema_name)
  return nil unless schema

  required = schema["required"] || []
  props = schema["properties"] || {}
  if required.size == 1 && props.size == 1
    return { "wrapper" => required.first }
  end
  return { "flat" => required } unless required.empty?

  { "wrapper" => nil }
end

schemas = (SPEC.dig("components", "schemas") || {}).keys.sort

operations = SPEC.fetch("paths", {}).flat_map do |path, methods|
  methods.filter_map do |method, operation|
    next unless %w[get post patch delete].include?(method)
    id = operation["operationId"]
    warn("missing operationId for #{method.upcase} #{path}") unless id
    next unless id

    {
      "id" => id,
      "method" => method,
      "path" => path,
      "tag" => normalize_tag(Array(operation["tags"]).first || "general"),
      "pathParams" => path.scan(/\{([^}]+)\}/).flatten,
      "pathParamTypes" => (operation["parameters"] || []).each_with_object({}) do |p, h|
        next unless p["in"] == "path"
        h[p["name"]] = type_of(p)
      end,
      "queryParams" => (operation["parameters"] || []).filter_map do |p|
        next unless p["in"] == "query"

        {
          "name" => p["name"],
          "type" => type_of(p),
          "required" => p["required"] == true,
          "description" => p["description"]
        }
      end,
      "requestSchema" => operation["x-wenmar-request-schema"],
      "requestShape" => request_shape(operation["x-wenmar-request-schema"]),
      "paginated" => paginated?(method, operation),
      "responseSchema" => response_schema(operation),
      "responseBody" => response_body?(operation),
      "summary" => operation["summary"]
    }
  end
end

File.write(OUT_PATH, JSON.pretty_generate(
                        "version" => MANIFEST_VERSION,
                        "source" => ENRICHED_PATH,
                        "schemas" => schemas,
                        "operations" => operations
                      ))
puts "Wrote #{OUT_PATH} with #{operations.size} operations"
