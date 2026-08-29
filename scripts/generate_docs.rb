# frozen_string_literal: true

require "yaml"
require "json"
require "fileutils"

# Generates LLM-friendly markdown API docs from the enriched OpenAPI spec.
#
# Usage:
#   ruby scripts/generate_docs.rb [spec_path] [out_dir]
#
# Defaults: spec/openapi.enriched.yaml -> docs/api
#
# Outputs:
#   sections/{tag}.md — per-resource endpoint reference, grouped by spec tag
#   api-reference.md  — endpoint table for every operation
#
# Narrative docs (README, conventions, authentication, errors, pagination) are
# hand-written and never touched. Input is the enriched spec, which already
# carries operationIds, tags, descriptions, and x-curl-example.
module GenerateDocs
  HTTP_METHODS = %w[get post patch put delete head options trace].freeze
  ACRONYMS = %w[vip vin wip url id api vin].freeze

  module_function

  # Returns { "tag" => [operation hashes] }, sorted by tag.
  # Each operation hash carries everything render_operation needs, including
  # the names of component schemas it references (for schema sections).
  def document(spec)
    ops = extract_operations(spec)
    ops.each { |op| op[:schemas] = referenced_schemas(op, spec) }
    ops.group_by { |op| op[:tag] }.sort.to_h
  end

  def extract_operations(spec)
    spec.fetch("paths", {}).flat_map do |path, methods|
      methods.filter_map do |method, operation|
        next unless HTTP_METHODS.include?(method)
        next unless operation.is_a?(Hash)

        tag = (operation["tags"] || []).first ||
              path.split("/").reject(&:empty?)[0] ||
              "misc"
        description = [operation["description"], titleize_operation_id(operation["summary"])]
                      .compact.map(&:strip).reject(&:empty?).first
        {
          tag: tag,
          path: path,
          method: method.upcase,
          operation_id: operation["operationId"],
          description: description,
          parameters: (operation["parameters"] || []).select { |p| p["in"] != "header" },
          request_body: operation["requestBody"],
          responses: operation["responses"] || {},
          curl: operation["x-curl-example"]
        }
      end
    end.sort_by { |op| [op[:path], op[:method]] }
  end

  def titleize_operation_id(operation_id)
    return "" unless operation_id

    words = operation_id.split("_")
    first = ACRONYMS.include?(words[0]) ? words[0].upcase : words[0].capitalize
    rest = words[1..].map { |w| ACRONYMS.include?(w) ? w.upcase : w.downcase }
    ([first] + rest).join(" ")
  end

  # Returns the names of component schemas referenced anywhere under an
  # operation (responses, request bodies, nested arrays / objects).
  def referenced_schemas(op, spec)
    schemas = spec.dig("components", "schemas") || {}
    found = []
    walk = lambda do |node|
      case node
      when Hash
        if (ref = node["$ref"]) && ref.start_with?("#/components/schemas/")
          found << ref.split("/").last
        end
        node.each_value { |v| walk.call(v) }
      when Array
        node.each { |v| walk.call(v) }
      end
    end
    walk.call(
      "parameters" => op[:parameters],
      "requestBody" => op[:request_body],
      "responses" => op[:responses]
    )
    found.uniq & schemas.keys
  end

  # --- Operation rendering ---------------------------------------------------

  def render_operation(op, fixtures: {})
    squeeze_newlines(markdown(op, fixtures: fixtures))
  end

  # Collapses every run of 3+ newlines to exactly 2 (one blank line).
  def squeeze_newlines(string)
    string.gsub(/\n{3,}/, "\n\n")
  end

  def markdown(op, fixtures: {})
    String.new.tap do |out|
      out << "## #{titleize_operation_id(op[:operation_id])}\n\n"
      out << "```\n#{op[:method]} #{op[:path]}\n```\n\n"
      desc = op[:description].to_s.strip
      out << "#{desc}\n\n" unless desc.empty?
      out << params(op)
      out << body(op)
      out << responses(op)
      out << example(op, fixtures)
      out << curl(op)
    end
  end

  # Renders a pretty-printed JSON example from a fixture when one is available
  # for this operationId. Placed after the response shape, before the cURL.
  def example(op, fixtures)
    data = fixtures[op[:operation_id]]
    return "" unless data

    json = JSON.pretty_generate(data)
    "**Example**\n\n```json\n#{json}\n```\n\n"
  end

  def params(op)
    return "" if op[:parameters].empty?

    rows = op[:parameters].map do |p|
      "| `#{p["name"]}` | #{field_type(p["schema"] || {})} | #{p["required"] ? "Yes" : "No"} |"
    end

    table("Param", "Type", "Required", rows)
  end

  def body(op)
    schema = op.dig(:request_body, "content", "application/json", "schema")
    return "" unless schema

    props = schema["properties"] || {}
    if props.length == 1
      key, inner = props.first
      if inner["type"] == "object" && inner["properties"]
        fields = render_schema_table(inner)
        return "**Request body** — wrapper key `#{key}`:\n\n#{fields}\n"
      end
    end

    fields = render_schema_table(schema)
    fields.empty? ? "" : "**Request body**:\n\n#{fields}\n"
  end

  def responses(op)
    blocks = op[:responses].keys.sort.filter_map do |status|
      response = op[:responses][status]
      schema = response.dig("content", "application/json", "schema")

      if schema.nil?
        next "**Response #{status}** — no content.\n"
      end

      shape = describe_response_shape(schema)
      next if shape.nil?

      out = +"**Response #{status}**"
      out << " — #{shape[:detail]}" if shape[:detail]
      out << "\n\n#{shape[:table]}" if shape[:table]
      out << "\n"
      out
    end

    blocks.empty? ? "" : "#{blocks.join("\n")}\n"
  end

  # Returns { detail:, table: } or nil when nothing beyond the status is shown.
  def describe_response_shape(schema)
    if (ref = schema["$ref"])
      name = ref.split("/").last
      return { detail: "[#{name}](##{name.downcase}-schema)" }
    end

    if schema["type"] == "array"
      items = schema["items"] || {}
      if (ref = items["$ref"])
        name = ref.split("/").last
        return { detail: "array of [#{name}](##{name.downcase}-schema)" }
      end
      table = render_schema_table(items)
      return { detail: "array", table: table.empty? ? nil : table }
    end

    # Error envelope: { "error": <Error ref> } — show only a link, not a
    # one-row table restating the envelope.
    props = schema["properties"] || {}
    if props.keys == ["error"] && props["error"]["$ref"]
      ref_name = props["error"]["$ref"].split("/").last
      return { detail: "[#{ref_name}](##{ref_name.downcase}-schema) error envelope" }
    end

    table = render_schema_table(schema)
    return nil if table.empty? && schema["description"].nil?

    { detail: schema["description"], table: table.empty? ? nil : table }
  end

  def curl(op)
    example = op[:curl]
    return "" if example.nil? || example.strip.empty?

    "```bash\n#{example}\n```\n\n"
  end

  # --- Schema tables -----------------------------------------------------------

  def render_schema_table(schema)
    props = schema && schema["properties"]
    return "" unless props && !props.empty?

    required = schema["required"] || []
    rows = props.map do |name, prop|
      "| `#{name}` | #{field_type(prop)} | #{required.include?(name) ? "Yes" : "No"} |"
    end

    table = table("Field", "Type", "Required", rows)

    # Nested objects get a labelled sub-table (markdown tables cannot nest).
    props.each do |name, prop|
      next unless prop["type"] == "object" && prop["properties"]

      table << "\n`#{name}` — object:\n"
      table << render_schema_table(prop)
    end

    table
  end

  def table(headers1, headers2, headers3, rows)
    <<~MD
      | #{headers1} | #{headers2} | #{headers3} |
      |---|---|---|
      #{rows.join("\n")}

    MD
  end

  def field_type(prop)
    if (ref = prop["$ref"])
      name = ref.split("/").last
      return "[#{name}](##{name.downcase}-schema)"
    end

    case prop["type"]
    when "array"
      items = prop["items"] || {}
      if (ref = items["$ref"])
        name = ref.split("/").last
        "array of [#{name}](##{name.downcase}-schema)"
      elsif items["properties"]
        "array of object"
      else
        "array of #{items["type"] || "any"}"
      end
    when "object" then "object"
    when nil then "any"
    else
      prop["nullable"] ? "#{prop["type"]} \\| null" : prop["type"]
    end
  end

  # --- api-reference.md endpoint rows --------------------------------------------

  def endpoint_rows(spec)
    extract_operations(spec).map do |op|
      {
        tag: op[:tag],
        method: op[:method],
        path: op[:path],
        operation_id: op[:operation_id],
        description: op[:description]
      }
    end
  end

  # --- File assembly -------------------------------------------------------------

  def write_sections(outdir, spec, fixtures: {})
    sections_dir = File.join(outdir, "sections")
    FileUtils.mkdir_p(sections_dir)

    document(spec).each do |tag, ops|
      File.write(File.join(sections_dir, "#{tag}.md"), render_section(tag, ops, spec, fixtures: fixtures))
    end
  end

  def render_section(tag, ops, spec, fixtures: {})
    title = section_title(tag)
    ops_markdown = ops.map { |op| render_operation(op, fixtures: fixtures) }.join
    schema_blocks = schema_blocks(tag, ops, spec)

    squeeze_newlines(<<~MD)
      # #{title}

      <!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
           Run: make docs -->

      #{ops_markdown}
      #{schema_blocks}
    MD
  end

  # Loads fixtures from a manifest.yaml + JSON files in the given directory.
  # Returns { operationId => parsed_json }.
  def load_fixtures(fixtures_dir)
    manifest_path = File.join(fixtures_dir, "manifest.yaml")
    return {} unless File.exist?(manifest_path)

    manifest = YAML.load_file(manifest_path)
    targets = manifest["targets"] || []

    targets.filter_map do |target|
      op = target["operation"]
      fixture = target["fixture"]
      next unless op && fixture

      path = File.join(fixtures_dir, fixture)
      next unless File.exist?(path)

      [op, JSON.parse(File.read(path))]
    end.to_h
  end

  # Renders "### {Name} schema" blocks for every component schema referenced
  # by operations under this tag. Anchors match the [Name](#name-schema) links
  # emitted by field_type / describe_response_shape.
  def schema_blocks(_tag, ops, spec)
    schemas = spec.dig("components", "schemas") || {}
    ops.flat_map { |op| op[:schemas] }.uniq.filter_map do |name|
      schema = schemas[name]
      next if schema.nil?

      table = render_schema_table(schema)
      next if table.empty?

      <<~MD

        ---

        ### #{name} schema {##{name.downcase}-schema}

        #{table}
      MD
    end.join
  end

  def section_title(tag)
    {
      "account" => "Account",
      "customer_tags" => "Customer Tags",
      "customers" => "Customers",
      "drivers" => "Drivers",
      "locations" => "Locations",
      "service_categories" => "Service Categories",
      "statements" => "Statements",
      "team" => "Team",
      "vendors" => "Vendors",
      "vehicles" => "Vehicles",
      "work_orders" => "Work Orders"
    }.fetch(tag) { tag.tr("_", " ").split.map(&:capitalize).join(" ") }
  end

  def write_api_reference(outdir, spec)
    rows = endpoint_rows(spec).sort_by { |r| [r[:tag], r[:path], r[:method]] }
    table = +"| Method | Path | Operation | Description |\n|---|---|---|---|\n"
    rows.each do |r|
      desc = (r[:description] || "").strip
      table << "| #{r[:method]} | `#{r[:path]}` | `#{r[:operation_id]}` | #{desc} |\n"
    end

    content = <<~MD
      # Wenmar Pro API Reference

      <!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
           Run: make docs -->

      Base URL: `https://app.wenmarpro.com`

      For detailed per-resource docs, see [sections/](sections/).

      ## Endpoints

      #{table}
      For common error and pagination shapes, see [errors](errors.md) and [pagination](pagination.md).
    MD

    FileUtils.mkdir_p(outdir)
    File.write(File.join(outdir, "api-reference.md"), content)
  end

  # --- llm-compact.md -----------------------------------------------------------

  # A single-file compact view of the entire API surface — no cURL, no examples,
  # no schema tables. Just method, path, params, and response shape per
  # operation. Optimized for stuffing into a single LLM context read.
  def write_llm_compact(outdir, spec)
    groups = document(spec)
    body = String.new

    groups.each do |tag, ops|
      body << "## #{section_title(tag)}\n\n"
      ops.each do |op|
        body << compact_operation_line(op)
      end
      body << "\n"
    end

    content = <<~MD
      # Wenmar Pro API — Compact Reference

      <!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
           Run: make docs -->

      Base URL: `https://app.wenmarpro.com`
      Auth: `Authorization: Bearer <token>`
      Responses: bare objects/arrays, no envelope. Errors: `{ "error": { code, message, details } }`.

      #{body}
      ## Schemas

      #{compact_schemas(spec)}
    MD

    FileUtils.mkdir_p(outdir)
    File.write(File.join(outdir, "llm-compact.md"), content)
  end

  def compact_operation_line(op)
    line = "- **#{op[:method]} #{op[:path]}**"
    params = op[:parameters].map { |p| p["name"] }
    line << " ?#{params.join(",")}" unless params.empty?
    line << " -> #{compact_response(op[:responses])}"
    line << "\n"
  end

  def compact_response(responses)
    successes = responses.keys.sort.filter_map do |status|
      schema = responses.dig(status, "content", "application/json", "schema")
      next "no content" if schema.nil?

      shape = compact_shape(schema)
      "#{status}: #{shape}"
    end
    successes.join(" | ")
  end

  def compact_shape(schema)
    if (ref = schema["$ref"])
      return ref.split("/").last
    end

    case schema["type"]
    when "array"
      items = schema["items"] || {}
      if (ref = items["$ref"])
        "array of #{ref.split("/").last}"
      elsif items["properties"]
        "array of object"
      else
        "array of #{items["type"] || "any"}"
      end
    when "object"
      props = schema["properties"] || {}
      if props.keys == ["error"] && props["error"]["$ref"]
        "error envelope"
      elsif props.empty?
        "object"
      else
        "object{#{props.keys.join(",")}}"
      end
    else
      schema["type"] || "any"
    end
  end

  def compact_schemas(spec)
    schemas = spec.dig("components", "schemas") || {}
    return "" if schemas.empty?

    lines = schemas.map do |name, schema|
      props = schema["properties"] || {}
      required = schema["required"] || []
      fields = props.map do |field, prop|
        req = required.include?(field) ? "*" : ""
        "#{field}#{req}:#{compact_shape(prop)}"
      end
      "- **#{name}**: {#{fields.join(", ")}}"
    end
    lines.join("\n") + "\n"
  end
end

if __FILE__ == $0
  input = ARGV.fetch(0, "spec/openapi.enriched.yaml")
  outdir = ARGV.fetch(1, "docs/api")
  fixtures_dir = ARGV.fetch(2, "spec/fixtures")
  spec = YAML.load_file(input)
  fixtures = GenerateDocs.load_fixtures(fixtures_dir)
  GenerateDocs.write_sections(outdir, spec, fixtures: fixtures)
  GenerateDocs.write_api_reference(outdir, spec)
  GenerateDocs.write_llm_compact(outdir, spec)
  puts "API docs written to #{outdir} (#{fixtures.size} fixtures loaded)"
end