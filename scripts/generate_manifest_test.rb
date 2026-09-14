# frozen_string_literal: true

require "minitest/autorun"
require "json"
require "tmpdir"

# Golden test for the manifest generator. It regenerates spec/operations.json
# from the committed enriched spec and compares byte-for-byte against the
# committed manifest. This guards the four downstream generators that consume
# it (Go wrapper, Ruby resources, conformance dispatch, parity gates): if any
# of them would drift, this fails first with a clear diff surface.
#
# It also asserts the emitted operations are sorted by id, which is what makes
# regeneration independent of the upstream YAML path order.
class GenerateManifestTest < Minitest::Test
  ROOT = File.expand_path("..", __dir__)
  ENRICHED = File.join(ROOT, "spec", "openapi.enriched.yaml")
  COMMITTED = File.join(ROOT, "spec", "operations.json")
  GENERATOR = File.join(ROOT, "scripts", "generate_manifest.rb")

  def generate_manifest
    Dir.mktmpdir do |dir|
      out = File.join(dir, "operations.json")
      # Run from the repo root with the same relative paths the Makefile uses so
      # the manifest's `source` field matches the committed file.
      ok = Dir.chdir(ROOT) do
        system("ruby", "scripts/generate_manifest.rb", "spec/openapi.enriched.yaml", out,
               out: File::NULL, err: File::NULL)
      end
      raise "generate_manifest.rb failed" unless ok

      return JSON.parse(File.read(out))
    end
  end

  def test_regenerated_manifest_matches_committed
    generated = generate_manifest
    committed = JSON.parse(File.read(COMMITTED))
    assert_equal committed, generated,
                 "spec/operations.json has drifted — run `make generate` and commit"
  end

  def test_operations_are_sorted_by_id
    ids = generate_manifest.fetch("operations").map { |o| o.fetch("id") }
    assert_equal ids.sort, ids, "manifest operations must be emitted sorted by id for determinism"
  end

  def test_manifest_carries_expected_top_level_keys
    manifest = generate_manifest
    assert_equal "2026-08-30", manifest.fetch("version")
    refute_empty manifest.fetch("schemas")
    refute_empty manifest.fetch("operations")
  end
end
