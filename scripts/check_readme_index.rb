#!/usr/bin/env ruby
# frozen_string_literal: true

# CI gate that verifies the generated "API endpoints" index block in
# docs/api/README.md matches the current enriched spec. The rest of README.md
# is hand-written and freely editable; only the marker-delimited block is
# generated, so we compare just that block rather than diffing the whole file.
#
# Usage: ruby scripts/check_readme_index.rb
# Exit 0 on success, 1 with a diagnostic on drift.

require "yaml"
require_relative "generate_docs"

ROOT = File.expand_path("..", __dir__)
SPEC_PATH = File.join(ROOT, "spec", "openapi.enriched.yaml")
README_PATH = File.join(ROOT, "docs", "api", "README.md")
START_MARKER = "<!-- START API ENDPOINTS -->"
END_MARKER = "<!-- END API ENDPOINTS -->"

def fail!(msg)
  warn "ERROR: #{msg}"
  exit 1
end

fail!("README not found at #{README_PATH}") unless File.exist?(README_PATH)
fail!("enriched spec not found at #{SPEC_PATH}") unless File.exist?(SPEC_PATH)

spec = YAML.load_file(SPEC_PATH)
content = File.read(README_PATH)
start_idx = content.index(START_MARKER)
end_idx = content.index(END_MARKER)
fail!("README is missing the API ENDPOINTS marker block") unless start_idx && end_idx && end_idx > start_idx

actual_block = content[start_idx...(end_idx + END_MARKER.length)]

tags = GenerateDocs.document(spec).keys
expected_lines = tags.map do |tag|
  title = GenerateDocs.section_title(tag)
  anchor = title.downcase.tr(" ", "-")
  "- [#{title}](sections/#{tag}.md##{anchor})"
end
expected_block = "#{START_MARKER}\n#{expected_lines.join("\n")}\n#{END_MARKER}"

if actual_block == expected_block
  puts "OK: README endpoint index covers all #{tags.size} tags"
else
  fail!(
    "docs/api/README.md API endpoints index is stale — run `make docs` and commit.\n" \
    "Expected #{tags.size} entries between the markers."
  )
end
