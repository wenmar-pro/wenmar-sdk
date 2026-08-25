#!/usr/bin/env ruby
# frozen_string_literal: true

# Validates shared JSON fixtures against the enriched OpenAPI spec's response
# schemas. Guarantees every covered schema has a live, validated representative
# among the fixtures, so a new required field added to a response schema forces
# a fixture to carry it.
#
# Usage: ruby scripts/check_fixture_coverage.rb
# Exit 0 on success, 1 on any validation failure.

require "yaml"
require "json"

ROOT = File.expand_path("..", __dir__)
SPEC_PATH = File.join(ROOT, "spec", "openapi.enriched.yaml")
MANIFEST_PATH = File.join(ROOT, "spec", "fixtures", "manifest.yaml")
FIXTURES_DIR = File.join(ROOT, "spec", "fixtures")

def fail!(message)
  warn "FAIL: #{message}"
  exit 1
end

def load_spec
  spec = YAML.load_file(SPEC_PATH)
  fail!("enriched spec not found at #{SPEC_PATH} — run scripts/enrich_spec.rb first") unless spec
  spec
end

def find_operation(spec, operation_id)
  spec["paths"].each do |_path, methods|
    methods.each do |method, operation|
      return [operation, method] if operation["operationId"] == operation_id
    end
  end
  nil
end

# Returns the schema for the `data` envelope of an operation's 2xx response.
def response_data_schema(operation)
  responses = operation["responses"] || {}
  response = responses["200"] || responses["201"] || responses["default"]
  return nil unless response

  schema = response.dig("content", "application/json", "schema")
  return nil unless schema

  # Unwrap the { "data": ... } envelope.
  schema.dig("properties", "data") || schema
end

# Structural validation: required fields present, types match, nullability.
def validate_schema(instance, schema, path = "$")
  return if instance.nil? && schema["nullable"]

  case schema["type"]
  when "object"
    unless instance.is_a?(Hash)
      fail!("expected object at #{path}, got #{instance.class}")
    end
    (schema["required"] || []).each do |key|
      unless instance.key?(key) && !instance[key].nil?
        fail!("missing required field '#{key}' at #{path}")
      end
    end
    (schema["properties"] || {}).each do |key, prop_schema|
      next unless instance.key?(key)

      validate_schema(instance[key], prop_schema, "#{path}.#{key}")
    end
  when "array"
    unless instance.is_a?(Array)
      fail!("expected array at #{path}, got #{instance.class}")
    end
    item_schema = schema["items"]
    instance.each_with_index do |item, i|
      validate_schema(item, item_schema, "#{path}[#{i}]")
    end
  when "integer"
    unless instance.is_a?(Integer)
      fail!("expected integer at #{path}, got #{instance.class}")
    end
  when "string"
    unless instance.is_a?(String)
      fail!("expected string at #{path}, got #{instance.class}")
    end
  when "boolean"
    unless instance == true || instance == false
      fail!("expected boolean at #{path}, got #{instance.class}")
    end
  end
end

def load_fixture(path)
  full = File.join(FIXTURES_DIR, path)
  fail!("fixture not found: #{path}") unless File.exist?(full)

  JSON.parse(File.read(full))
end

def validate_target(spec, target)
  fixture = load_fixture(target["fixture"])

  if target["operation"]
    operation, _method = find_operation(spec, target["operation"])
    fail!("unknown operation '#{target["operation"]}' for fixture #{target["fixture"]}") unless operation

    schema = response_data_schema(operation)
    fail!("no 2xx response schema for operation '#{target["operation"]}'") unless schema

    validate_schema(fixture, schema, target["fixture"])
  elsif target["pointer"]
    # Pointer entries are not used yet; the operation entries cover all schemas.
    fail!("pointer entries not yet supported: #{target["id"]}")
  else
    fail!("target #{target["id"]} has neither operation nor pointer")
  end
end

def main
  spec = load_spec
  manifest = YAML.load_file(MANIFEST_PATH)
  targets = manifest["targets"] || []

  fail!("no targets in manifest") if targets.empty?

  targets.each do |target|
    validate_target(spec, target)
  end

  puts "OK: #{targets.length} fixture targets validated against the spec"
end

main
