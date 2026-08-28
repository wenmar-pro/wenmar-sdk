require "yaml"

module EnrichSpec
  SINGULAR_OVERRIDES = {
    "work_orders" => "work_order"
  }.freeze

  RESOURCE_SCHEMAS = {
    "customers" => "Customer",
    "vehicles" => "Vehicle",
    "work_orders" => "WorkOrder"
  }.freeze

  # Collection sub-actions that are NOT standard CRUD on a resource. The key is
  # "METHOD path"; the value is the semantic operationId. The physical path may
  # differ from the operation name (e.g. /vehicles/vin_decode -> decode_vin).
  SUB_ACTION_IDS = {
    "get /vehicles/vin_decode"       => "decode_vin",
    "get /vehicles/check_duplicate"  => "check_duplicate"
  }.freeze

  def self.call(input)
    spec = Marshal.load(Marshal.dump(input)) # deep copy

    spec["servers"] = [{ "url" => "https://app.wenmarpro.com" }]
    spec["paths"] ||= {}

    extract_components!(spec)

    spec["paths"].each do |path, methods|
      resource = path.split("/")[1] # "/customers" => "customers"
      methods.each do |method, operation|
        operation["operationId"] = derive_operation_id(path, method)
        operation["tags"] = [resource] if resource
        operation["description"] = derive_description(operation["operationId"])
        operation["x-curl-example"] = derive_curl_example(path, method)
      end
    end

    apply_example_overrides!(spec)
    apply_overlays!(spec)

    spec
  end

  # Applies the simplified OpenAPI Overlay files in overlays/*.yaml to the
  # enriched spec. This is a simple YAML merge, not a full OpenAPI Overlay
  # implementation. Supported targets:
  #   "$.paths.*.*"                 — every operation
  #   "$.paths['/path'].method"     — a single operation
  def self.apply_overlays!(spec)
    overlay_dir = File.expand_path("../overlays", __dir__)
    Dir.glob(File.join(overlay_dir, "*.yaml")).sort.each do |file|
      overlay = YAML.load_file(file)
      (overlay["actions"] || []).each do |action|
        target = action["target"]
        update = action["update"] || {}
        next unless target

        if target == "$.paths.*.*"
          (spec["paths"] || {}).each_value do |methods|
            methods.each_value { |operation| operation.merge!(Marshal.load(Marshal.dump(update))) }
          end
        elsif target =~ %r{\A\$\.paths\['([^']+)'\]\.(\w+)\z}
          path, method = Regexp.last_match(1), Regexp.last_match(2)
          operation = spec.dig("paths", path, method)
          operation.merge!(Marshal.load(Marshal.dump(update))) if operation
        end
      end
    end
  end

  def self.derive_operation_id(path, method)
    # Sub-actions (e.g. /vehicles/vin_decode) map to a named operationId.
    sub_action = SUB_ACTION_IDS["#{method} #{path}"]
    return sub_action if sub_action

    segments = path.split("/").reject(&:empty?)
    resource = segments[0] # "customers", "vehicles", "work_orders"
    return nil unless resource

    singularized = SINGULAR_OVERRIDES[resource] || resource.chomp("s")

    # Nested resource paths (e.g. /work_orders/{work_order_id}/payments) get a
    # distinct operationId so they don't collide with the top-level resource.
    if segments.length >= 3
      nested = segments[2]
      nested_singular = nested.chomp("s")
      # Deeper nesting (e.g. .../vehicles/{vehicle_id}/history) appends the
      # final segment to keep operationIds unique.
      if segments.length >= 5
        return "#{method}_#{resource}_#{nested_singular}_#{segments[4]}"
      end
      return "#{method}_#{resource}_#{nested_singular}"
    end

    has_id = path.include?("{id}")

    case [method, has_id]
    when ["get", false] then "list_#{resource}"
    when ["get", true]  then "show_#{singularized}"
    when ["post", false] then "create_#{singularized}"
    when ["patch", true], ["put", true] then "update_#{singularized}"
    when ["delete", true] then "delete_#{singularized}"
    when ["delete", false] then "delete_#{resource}"
    else "#{method}_#{resource}"
    end
  end

  def self.derive_description(operation_id)
    return nil unless operation_id
    verb, resource = operation_id.split("_", 2)
    readable = resource.tr("_", " ")
    case verb
    when "list"   then "List all #{readable}, paginated via the Link header."
    when "show"   then "Show a #{readable} by ID."
    when "create" then "Create a #{readable}."
    when "update" then "Update a #{readable} by ID."
    when "delete" then "Delete a #{readable} by ID."
    end
  end

  # Generates a copy-paste cURL command for a given path + method.
  def self.derive_curl_example(path, method)
    base = "https://app.wenmarpro.com"
    url = if path.include?("{id}")
      "#{base}#{path.split('{')[0]}<id>.json"
    else
      "#{base}#{path}.json"
    end
    headers = '-H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN"'
    case method
    when "get"
      "curl #{headers} #{url}"
    when "post", "patch"
      "curl -X #{method.upcase} #{headers} -H \"Content-Type: application/json\" \\\n     -d '{\"...\":\"...\"}' #{url}"
    when "delete"
      "curl -X DELETE #{headers} #{url}"
    end
  end

  # Realistic example values matching bc3's documentation quality bar. Applied
  # to component schema `example` fields.
  EXAMPLE_OVERRIDES = {
    "Customer" => {
      "id" => 7,
      "full_name" => "Jane Doe",
      "company_name" => "Acme Auto",
      "emails" => [{ "id" => 42, "address" => "jane@acmeauto.com", "label" => "work", "primary" => true }]
    },
    "Vehicle" => {
      "id" => 13,
      "make" => "Honda",
      "model" => "Civic",
      "year" => 2020,
      "vin" => "1HGCN2345ABC",
      "odometer" => { "reading" => 182_340, "unit" => "km" }
    },
    "WorkOrder" => {
      "id" => 42,
      "status" => "in_progress",
      "work_order_number" => "WO-0042",
      "totals" => { "subtotal_cents" => 12_000, "currency" => "CAD" }
    }
  }.freeze

  def self.apply_example_overrides!(spec)
    schemas = spec.dig("components", "schemas")
    return unless schemas

    EXAMPLE_OVERRIDES.each do |name, example|
      schema = schemas[name]
      next unless schema

      props = schema["properties"]
      next unless props

      props.each do |prop_name, prop_schema|
        next unless example.key?(prop_name)

        # Prefer the nested example for object-valued props; fall back to scalar.
        value = example[prop_name]
        if prop_schema["type"] == "object" && prop_schema["properties"] && value.is_a?(Hash)
          prop_schema["properties"].each do |sub_name, sub_schema|
            sub_schema["example"] = value[sub_name] if value.key?(sub_name)
          end
        else
          prop_schema["example"] = value
        end
      end
    end
  end

  def self.extract_components!(spec)
    spec["components"] ||= {}
    spec["components"]["schemas"] ||= {}

    spec["paths"].each do |path, methods|
      segments = path.split("/").reject(&:empty?)
      # Only top-level resource paths (e.g. /customers, /customers/{id}) map to
      # a component schema. Nested paths (e.g. /customers/{customer_id}/vehicles)
      # are sub-resources and must not pollute the parent schema.
      next unless segments.length <= 2

      resource = segments[0]
      schema_name = RESOURCE_SCHEMAS[resource]
      next unless schema_name

      methods.each do |method, operation|
        # decode_vin / check_duplicate return free-form data, not a resource.
        next if SUB_ACTION_IDS["#{method} #{path}"]

        responses = operation["responses"] || {}
        responses.each do |status, response|
          # Only 2xx responses carry the resource representation; error
          # envelopes (4xx/5xx) are extracted separately below.
          next unless status.to_s.start_with?("2")

          content = response["content"]
          next unless content && content["application/json"]

          schema = content["application/json"]["schema"]
          next unless schema

          # Candidate component: the item schema of a bare array, or the bare
          # object schema itself. Prefer the RICHEST representation (show
          # responses carry metrics that list partials omit) so generated SDK
          # types expose the full attribute surface.
          candidate = if schema["type"] == "array" && schema["items"] && schema["items"]["properties"]
            schema["items"].dup
          elsif schema["properties"]
            schema.dup
          else
            nil
          end
          next unless candidate

          existing = spec["components"]["schemas"][schema_name]
          if existing.nil? || candidate["properties"].keys.length > existing["properties"].keys.length
            spec["components"]["schemas"][schema_name] = candidate
          end

          if schema["type"] == "array"
            schema["items"] = { "$ref" => "#/components/schemas/#{schema_name}" }
          else
            content["application/json"]["schema"] = { "$ref" => "#/components/schemas/#{schema_name}" }
          end
        end
      end
    end

    # Extract error schema
    spec["paths"].each do |_path, methods|
      methods.each do |_method, operation|
        responses = operation["responses"] || {}
        responses.each do |status, response|
          next unless status.to_s.start_with?("4") || status.to_s.start_with?("5")
          content = response["content"]
          next unless content && content["application/json"]

          schema = content["application/json"]["schema"]
          next unless schema && schema["properties"] && schema["properties"]["error"]

          unless spec["components"]["schemas"]["Error"]
            spec["components"]["schemas"]["Error"] = schema["properties"]["error"].dup
          end
          schema["properties"]["error"] = { "$ref" => "#/components/schemas/Error" }
        end
      end
    end

    spec
  end

  def self.run(input_path, output_path)
    spec = YAML.load_file(input_path)
    enriched = call(spec)
    File.write(output_path, enriched.to_yaml)
    puts "Enriched spec written to #{output_path}"
  end
end

if __FILE__ == $0
  input = ARGV[0] || "spec/openapi.yaml"
  output = ARGV[1] || "spec/openapi.enriched.yaml"
  EnrichSpec.run(input, output)
end
