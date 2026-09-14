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
require_relative "generator_utils"

ROOT = File.expand_path("..", __dir__)
MANIFEST_PATH = File.join(ROOT, "spec", "operations.json")
GO_DISPATCH = File.join(ROOT, "conformance", "go", "dispatch.gen.go")
RUBY_DISPATCH = File.join(ROOT, "conformance", "ruby", "dispatch.gen.rb")
COVERAGE_ALLOWLIST = File.join(ROOT, "conformance", "coverage_allowlist.json")

TEST_ALIASES = GeneratorUtils::TEST_ALIASES

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

# Coverage check: every manifest operation must have a conformance scenario,
# unless it is explicitly listed in conformance/coverage_allowlist.json.
#
# The allowlist may only shrink. If a NEW operation ships without a scenario
# (and isn't listed) the gate fails hard — the allowlist exists to record the
# known backlog, not to permit new gaps. Setting WENMAR_STRICT_COVERAGE=1
# additionally fails when the allowlist is non-empty, so the backlog can be
# driven to zero.
resolved_test_ops = test_ops.uniq.map { |op| TEST_ALIASES[op] || op }
untested = manifest.reject { |op| resolved_test_ops.include?(op) }

allowlist = if File.exist?(COVERAGE_ALLOWLIST)
              JSON.parse(File.read(COVERAGE_ALLOWLIST))["operations"] || []
            else
              []
            end

# Allowlist entries that are now tested must be removed (keep it shrinking).
stale_allowlist = allowlist.reject { |op| untested.include?(op) }
unless stale_allowlist.empty?
  fail!("coverage_allowlist.json lists operations that now have scenarios (remove them): #{stale_allowlist.sort.join(', ')}")
end

# Allowlist entries that no longer exist at all are also stale.
unknown_allowlist = allowlist - manifest
unless unknown_allowlist.empty?
  fail!("coverage_allowlist.json lists unknown operations (remove them): #{unknown_allowlist.sort.join(', ')}")
end

newly_untested = untested - allowlist
unless newly_untested.empty?
  fail!("operations have no conformance test scenario and are not allowlisted:\n  #{newly_untested.sort.join("\n  ")}\n" \
        "Add scenarios in conformance/tests/*.json, or record them in conformance/coverage_allowlist.json.")
end

if untested.any?
  puts "WARNING: #{untested.size} operations have no conformance test scenario (all allowlisted):"
  untested.sort.each { |op| puts "  - #{op}" }
  puts "coverage_allowlist.json may only shrink; add scenarios and remove ids from it."
  if ENV["WENMAR_STRICT_COVERAGE"] == "1"
    fail!("WENMAR_STRICT_COVERAGE=1 but #{untested.size} operations remain untested (allowlist non-empty)")
  end
else
  puts "All #{manifest.size} operations have at least one test scenario."
end

puts "OK: manifest=#{manifest.size} test-cases=#{test_ops.size} go-dispatch=#{go_ops.size} ruby-dispatch=#{ruby_ops.size} parity holds"
