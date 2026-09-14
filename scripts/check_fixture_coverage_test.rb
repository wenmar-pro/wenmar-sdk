# frozen_string_literal: true

require "minitest/autorun"
require "yaml"
require "json"
require "tmpdir"
require "open3"

# Failure-behavior tests for the fixture-coverage gate. The gate is a
# subprocess that exits 1 on failure. We run it against a temporary spec +
# fixture set to prove it fails loudly instead of silently passing.
class CheckFixtureCoverageTest < Minitest::Test
  ROOT = File.expand_path("..", __dir__)
  GATE = File.join(ROOT, "scripts", "check_fixture_coverage.rb")

  def run_gate(spec:, manifest:, fixtures:)
    Dir.mktmpdir do |dir|
      spec_path = File.join(dir, "spec.yaml")
      manifest_path = File.join(dir, "manifest.yaml")
      fixtures_dir = File.join(dir, "fixtures")
      FileUtils.mkdir_p(fixtures_dir)
      File.write(spec_path, YAML.dump(spec))
      File.write(manifest_path, YAML.dump(manifest))
      fixtures.each { |name, body| File.write(File.join(fixtures_dir, name), JSON.generate(body)) }
      Open3.capture2e(
        { "WENMAR_SPEC_PATH" => spec_path, "WENMAR_FIXTURE_MANIFEST" => manifest_path,
          "WENMAR_FIXTURES_DIR" => fixtures_dir },
        "ruby", GATE
      )
    end
  end

  def spec_with(schema)
    {
      "paths" => {
        "/things" => {
          "get" => {
            "operationId" => "list_things",
            "responses" => { "200" => { "content" => { "application/json" => { "schema" => schema } } } }
          }
        }
      },
      "components" => { "schemas" => {} }
    }
  end

  def manifest
    { "targets" => [{ "id" => "thing-list", "operation" => "list_things", "fixture" => "things.json" }] }
  end

  def test_committed_fixtures_pass
    out, status = Open3.capture2e("ruby", "scripts/check_fixture_coverage.rb", chdir: ROOT)
    assert status.success?, "expected fixture coverage to pass, got: #{out}"
    assert_includes out, "fixture targets validated"
  end

  def test_unhandled_schema_type_fails_loudly
    schema = { "type" => "object", "properties" => { "blob" => { "type" => "file" } } }
    out, status = run_gate(spec: spec_with(schema), manifest: manifest, fixtures: { "things.json" => { "blob" => "x" } })
    refute status.success?, "expected an unhandled schema type to fail the gate"
    assert_includes out, "unhandled schema type"
  end

  def test_unresolved_ref_fails_loudly
    schema = { "type" => "object", "properties" => { "thing" => { "$ref" => "#/components/schemas/Missing" } } }
    out, status = run_gate(spec: spec_with(schema), manifest: manifest, fixtures: { "things.json" => { "thing" => {} } })
    refute status.success?, "expected an unresolved $ref to fail the gate"
    assert_includes out, "unresolved $ref"
  end

  def test_number_type_is_validated
    schema = { "type" => "object", "properties" => { "amount" => { "type" => "number" } } }
    out, status = run_gate(spec: spec_with(schema), manifest: manifest, fixtures: { "things.json" => { "amount" => "nope" } })
    refute status.success?, "expected a string where number is required to fail"
    assert_includes out, "expected number"
  end
end
