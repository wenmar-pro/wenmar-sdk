#!/usr/bin/env ruby
# frozen_string_literal: true

# CI gate that enforces parity between the operations manifest, the conformance
# test cases, and the generated Go/Ruby dispatch files.
#
#  1. spec/operations.json exists and has >= 1 operation.
#  2. Every operation referenced by conformance/tests/*.json resolves to a
#     manifest operation (directly or via a known alias).
#  3. Every manifest operation is present in both dispatch files.
#  4. No extra operations in the dispatch files that are not in the manifest
#     (except test-only aliases).
#
# Exit 0 with counts, or exit 1 with a diagnostic.
#
# Usage: ruby scripts/check_conformance_parity.rb

require "json"

ROOT = File.expand_path("..", __dir__)
MANIFEST_PATH = File.join(ROOT, "spec", "operations.json")
GO_DISPATCH = File.join(ROOT, "conformance", "go", "dispatch.gen.go")
RUBY_DISPATCH = File.join(ROOT, "conformance", "ruby", "dispatch.gen.rb")

# Test operations that map to a manifest operation under a different name.
# These are intentional aliases the conformance runner uses; keep in sync with
# generate_conformance_dispatch.rb.
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

def fail!(msg)
  warn "ERROR: #{msg}"
  exit 1
end

def manifest_operations
  fail!("manifest not found at #{MANIFEST_PATH}; run make generate first") unless File.exist?(MANIFEST_PATH)
  data = JSON.parse(File.read(MANIFEST_PATH))
  ops = data["operations"]
  fail!("manifest has no operations") if ops.nil? || ops.empty?
  ops.map { |o| o["id"] }
end

def test_operations
  Dir[File.join(ROOT, "conformance", "tests", "*.json")].flat_map do |f|
    JSON.parse(File.read(f)).map { |tc| tc["operation"] }
  end
end

def dispatch_operations(path, key_regex)
  return [] unless File.exist?(path)
  content = File.read(path)
  content.scan(key_regex).flatten
end

manifest = manifest_operations
manifest_set = manifest.to_h { |id| [id, true] }

test_ops = test_operations
test_missing = test_ops.uniq.reject { |op| manifest_set.key?(op) || TEST_ALIASES.key?(op) }
unless test_missing.empty?
  fail!("conformance tests reference operations missing from the manifest: #{test_missing.join(', ')}")
end

# Resolve each test operation to a manifest operation (direct or alias).
resolved = test_ops.uniq.map { |op| TEST_ALIASES[op] || op }
uncovered = resolved.reject { |op| manifest_set.key?(op) }
unless uncovered.empty?
  fail!("conformance tests resolve to operations not in the manifest: #{uncovered.join(', ')}")
end

go_ops = dispatch_operations(GO_DISPATCH, /"([a-z0-9_]+)"\s*:\s*func/)
ruby_ops = dispatch_operations(RUBY_DISPATCH, /"([a-z0-9_]+)"\s*=>/)

go_missing = manifest.reject { |op| go_ops.include?(op) }
ruby_missing = manifest.reject { |op| ruby_ops.include?(op) }
unless go_missing.empty? && ruby_missing.empty?
  fail!("manifest operations missing from dispatch: go=[#{go_missing.join(', ')}] ruby=[#{ruby_missing.join(', ')}]")
end

allowed = manifest + TEST_ALIASES.keys
go_extra = go_ops.reject { |op| allowed.include?(op) }
ruby_extra = ruby_ops.reject { |op| allowed.include?(op) }
unless go_extra.empty? && ruby_extra.empty?
  fail!("dispatch files contain operations not in the manifest: go=[#{go_extra.join(', ')}] ruby=[#{ruby_extra.join(', ')}]")
end

puts "OK: manifest=#{manifest.size} test-cases=#{test_ops.size} go-dispatch=#{go_ops.size} ruby-dispatch=#{ruby_ops.size} parity holds"
