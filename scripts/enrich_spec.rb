require "yaml"

module EnrichSpec
  SINGULAR_OVERRIDES = {
    "work_orders" => "work_order",
    "service_categories" => "service_category"
  }.freeze

  RESOURCE_SCHEMAS = {
    "customers" => "Customer",
    "vehicles" => "Vehicle",
    "work_orders" => "WorkOrder",
    "drivers" => "Driver",
    "statements" => "Statement",
    "vendors" => "Vendor",
    "service_categories" => "ServiceCategory",
    "purchase_orders" => "PurchaseOrder",
    "return_orders" => "ReturnOrder",
    "sublet_orders" => "SubletOrder"
  }.freeze

  # Collection sub-actions that are NOT standard CRUD on a resource. The key is
  # "METHOD path"; the value is the semantic operationId. The physical path may
  # differ from the operation name (e.g. /vehicles/vin_decode -> decode_vin).
  SUB_ACTION_IDS = {
    "get /vehicles/vin_decode"       => "decode_vin",
    "get /vehicles/check_duplicate"  => "check_vehicle_duplicate",
    "get /customers/check_duplicate" => "check_customer_duplicate",
    # seed_defaults returns a {created, message} summary, not a ServiceCategory.
    "post /service_categories/seed_defaults" => "seed_defaults_service_categories",
    # Shipped operationIds (0.4.1) kept stable across the 2026-08-31 path
    # renames (merge/transfer pluralized; lifecycle flattened) + semantic
    # names for new surface (close_zero, payment AR send/reverse).
    "post /customers/{id}/merges"                             => "merge_customer",
    "post /vehicles/{id}/merges"                              => "merge_vehicle",
    "post /vehicles/{id}/transfers"                           => "transfer_vehicle",
    "patch /work_orders/{id}/close"                           => "close_work_order",
    "patch /work_orders/{id}/close_as_paid"                   => "close_work_order_as_paid",
    "patch /work_orders/{id}/close_zero"                      => "close_work_order_zero",
    "patch /work_orders/{id}/decline_all"                     => "decline_all_work_order_services",
    "patch /work_orders/{id}/reopen"                           => "reopen_work_order",
    "patch /work_orders/{id}/return_to_board"                 => "return_work_order_to_board",
    "patch /work_orders/{id}/save_for_later"                  => "save_work_order_for_later",
    "get /work_orders/{id}/service_history"                   => "show_work_order_service_history",
    "get /work_orders/{id}/declined_services"                 => "show_work_order_declined_services",
    "post /work_orders/{work_order_id}/payments/send_to_ar"   => "send_work_order_payment_to_ar",
    "delete /work_orders/{work_order_id}/payments/reverse_ar" => "reverse_work_order_payment_ar"
  }.freeze

  # Explicit operationIds for nested/sub-resource endpoints whose auto-derived
  # id would collide (e.g. drivers list vs show both derive to get_customers_driver)
  # or whose semantic name differs from the generic derivation. Keyed by
  # "METHOD path". These are stable contract names — do not rename once shipped.
  NESTED_OPERATION_IDS = {
    "get /customers/{customer_id}/drivers"              => "list_customers_drivers",
    "post /customers/{customer_id}/drivers"             => "create_driver",
    "get /customers/{customer_id}/drivers/{id}"         => "show_driver",
    "patch /customers/{customer_id}/drivers/{id}"       => "update_driver",
    "delete /customers/{customer_id}/drivers/{id}"       => "delete_driver",
    "get /customers/{customer_id}/statements"           => "list_customers_statements",
    "get /statements/{id}"                              => "show_statement",
    "get /vendors"                                      => "list_vendors",
    "get /vendors/{id}"                                 => "show_vendor",
    "get /work_orders/{work_order_id}/estimate"         => "show_work_order_estimate",
    "get /work_orders/{work_order_id}/wip"              => "show_work_order_wip",
    "get /work_orders/{work_order_id}/inspection"        => "show_work_order_inspection",
    "get /work_orders/{work_order_id}/parts"            => "show_work_order_parts",
    "get /work_orders/{work_order_id}/payments"          => "show_work_order_payments",
    "post /work_orders/{work_order_id}/payments"         => "create_work_order_payment",
    # Full CRUD expansion (2026-08-28): tags, lookup/prefill
    "get /settings/tags"                                  => "list_tags",
    "patch /settings/tags"                                => "update_tags",
    "get /customers/lookup"                               => "lookup_customer",
    "get /vehicles/lookup"                                => "lookup_vehicle",
    "get /vehicles/prefill"                               => "prefill_vehicle",
    "get /customers/{customer_id}/vehicles"               => "list_customers_vehicles",
    "get /customers/{customer_id}/work_orders"            => "list_customers_work_orders",
    "get /vehicles/{vehicle_id}/work_orders"              => "list_vehicles_work_orders",
    # Authorization endpoints (synced 2026-08-29; paths unchanged 2026-08-31)
    "post /work_orders/{work_order_id}/authorization"                => "create_work_order_authorization",
    "post /work_orders/{work_order_id}/authorization/update_decisions" => "update_work_order_authorization_decisions",
    # Work order service line-item sub-actions (synced 2026-08-31): the
    # auto-derived id for these deeply-nested paths would all collapse to
    # "patch_work_orders_service_line_items", so pin stable, distinct ids.
    "patch /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/add_to_inventory" => "add_work_order_service_line_item_to_inventory",
    "post /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/duplicate"          => "duplicate_work_order_service_line_item",
    "patch /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/pull"              => "pull_work_order_service_line_item",
    "patch /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/refresh_price"     => "refresh_work_order_service_line_item_price",
    "patch /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_pull"         => "undo_pull_work_order_service_line_item",
    "patch /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_return"       => "undo_return_work_order_service_line_item",
    "patch /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/update_part_status" => "update_work_order_service_line_item_part_status"
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

    assert_unique_operation_ids!(spec)

    name_request_schemas!(spec)

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
    # Explicit semantic/action overrides are the stable contract — keep them.
    sub_action = SUB_ACTION_IDS["#{method} #{path}"]
    return sub_action if sub_action

    nested = NESTED_OPERATION_IDS["#{method} #{path}"]
    return nested if nested

    segments = path.split("/").reject(&:empty?)
    return nil if segments.empty?

    # The resource chain is the meaningful path context. Drop {param}
    # placeholders and literal numeric IDs (e.g. /customers/1043910089/merge).
    # Hyphenated segments (e.g. lead-sources) are normalized to underscores so
    # the resulting operationId is a valid Go/Ruby identifier.
    nouns = segments.reject { |s| s.start_with?("{") || s =~ /\A\d+\z/ }.map { |s| s.tr("-", "_") }
    return nil if nouns.empty?

    last_is_param = segments.last.start_with?("{") || segments.last =~ /\A\d+\z/

    verb = case method
           when "get"    then last_is_param ? "show" : "list"
           when "post"   then "create"
           when "patch", "put" then "update"
           when "delete" then "delete"
           else method
           end

    # Singularize the final noun for single-resource verbs. Keep it plural for
    # list, and keep terminal action segments intact (e.g. close, hard_delete).
    singularize_last = case verb
                       when "list"   then false
                       when "create" then true
                       else last_is_param
                       end

    parts = [verb]
    nouns.each_with_index do |noun, i|
      is_last = i == nouns.length - 1
      parts << (is_last && singularize_last ? singularize(noun) : noun)
    end
    parts.join("_")
  end

  def self.singularize(word)
    SINGULAR_OVERRIDES[word] || word.chomp("s")
  end

  def self.assert_unique_operation_ids!(spec)
    seen = {}
    duplicates = []
    spec["paths"].each do |path, methods|
      methods.each do |method, op|
        next unless op.is_a?(Hash)
        id = op["operationId"]
        next if id.nil?
        if seen.key?(id)
          duplicates << "#{id} seen at #{seen[id]} and #{method.upcase} #{path}"
        else
          seen[id] = "#{method.upcase} #{path}"
        end
      end
    end
    return if duplicates.empty?
    warn "Duplicate operationIds detected:\n#{duplicates.join("\n")}"
    exit 1
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

  # Hoists inline JSON request bodies into named component schemas so
  # oapi-codegen emits a named request struct. Each request schema is named
  # "<PascalCase operationId>Request" and the operation is annotated with
  # x-wenmar-request-schema so downstream generators can find it.
  def self.name_request_schemas!(spec)
    spec["components"] ||= {}
    spec["components"]["schemas"] ||= {}

    spec["paths"].each do |path, methods|
      methods.each do |method, operation|
        body = operation.dig("requestBody", "content", "application/json", "schema")
        next unless body && !body["$ref"]

        operation_id = operation["operationId"]
        next unless operation_id

        schema_name = pascal_case("#{operation_id}_request")
        spec["components"]["schemas"][schema_name] = body
        operation["requestBody"]["content"]["application/json"]["schema"] =
          { "$ref" => "#/components/schemas/#{schema_name}" }
        operation["x-wenmar-request-schema"] = schema_name
      end
    end
  end

  def self.pascal_case(identifier)
    identifier.split("_").reject(&:empty?).map { |part| part[0].upcase + part[1..] }.join
  end

  def self.extract_components!(spec)
    spec["components"] ||= {}
    spec["components"]["schemas"] ||= {}

    spec["paths"].each do |path, methods|
      segments = path.split("/").reject(&:empty?)

      # Resolve the resource + schema name. Top-level resource paths
      # (e.g. /customers, /customers/{id}) map directly. Nested sub-resource
      # paths (e.g. /customers/{customer_id}/drivers) resolve to the LAST
      # segment when that segment is a registered resource (drivers,
      # statements, vendors). Sub-collection/action paths (vehicles,
      # work_orders, history, estimate, wip, inspection, parts, payments)
      # are handled by their top-level paths and must not pollute the parent
      # schema.
      # Resolve the resource + schema name. Top-level resource paths
      # (e.g. /customers, /customers/{id}) map directly. The `orders`
      # namespace (e.g. /orders/purchase_orders) prefixes the resource, so
      # segments[1] is the resource. Nested sub-resource paths
      # (e.g. /customers/{customer_id}/vehicles) resolve to the nested
      # segment when it's a registered resource. Sub-collection/action paths
      # (vehicles, work_orders, history, estimate, wip, inspection, parts,
      # payments) are handled by their top-level paths and must not pollute
      # the parent schema.
      namespace = segments[0] == "orders" && !RESOURCE_SCHEMAS.key?(segments[0])
      if namespace
        # orders namespace: resource is the second segment
        # (e.g. /orders/purchase_orders). Deeper sub-resources
        # (e.g. /orders/purchase_orders/{id}/cancellations) return the
        # base resource schema.
        resource = segments[1]
        schema_name = RESOURCE_SCHEMAS[resource]
        next unless schema_name
      else
        resource = segments[0]
        schema_name = RESOURCE_SCHEMAS[resource]
        if segments.length > 2 && RESOURCE_SCHEMAS.key?(segments[2])
          # Nested sub-resource paths (e.g. /customers/{customer_id}/vehicles)
          # resolve to the third segment when it's a registered resource.
          # Sub-collection/action paths (vehicles, work_orders, history,
          # estimate, wip, inspection, parts, payments) are handled by their
          # top-level paths and must not pollute the parent schema.
          resource = segments[2]
          schema_name = RESOURCE_SCHEMAS[segments[2]]
        end
        next unless schema_name
      end

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
