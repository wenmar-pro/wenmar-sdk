require "yaml"

module EnrichSpec
  OPERATION_IDS = {
    ["/api/customers", "get"] => "list_customers",
    ["/api/customers", "post"] => "create_customer",
    ["/api/customers/{id}", "get"] => "show_customer",
    ["/api/vehicles/{id}", "get"] => "show_vehicle",
    ["/api/work_orders", "get"] => "list_work_orders",
    ["/api/work_orders/{id}", "get"] => "show_work_order"
  }.freeze

  DESCRIPTIONS = {
    ["/api/customers", "get"] => "List all customers, paginated via the Link header.",
    ["/api/customers", "post"] => "Create a customer.",
    ["/api/customers/{id}", "get"] => "Show a customer by ID.",
    ["/api/vehicles/{id}", "get"] => "Show a single vehicle by ID.",
    ["/api/work_orders", "get"] => "List all work orders, paginated via the Link header.",
    ["/api/work_orders/{id}", "get"] => "Show a single work order by ID."
  }.freeze

  RESOURCE_SCHEMAS = {
    "customers" => "Customer",
    "vehicles" => "Vehicle",
    "work_orders" => "WorkOrder"
  }.freeze

  def self.call(input)
    spec = Marshal.load(Marshal.dump(input)) # deep copy

    spec["servers"] = [{ "url" => "https://app.wenmarpro.com" }]
    spec["paths"] ||= {}

    extract_components!(spec)

    spec["paths"].each do |path, methods|
      resource = path.split("/")[2] # "/api/customers" => "customers"
      methods.each do |method, operation|
        key = [path, method]
        operation["operationId"] = OPERATION_IDS[key] if OPERATION_IDS[key]
        operation["tags"] = [resource] if resource
        operation["description"] = DESCRIPTIONS[key] if DESCRIPTIONS[key]
      end
    end

    spec
  end

  def self.extract_components!(spec)
    spec["components"] ||= {}
    spec["components"]["schemas"] ||= {}

    spec["paths"].each do |path, methods|
      resource = path.split("/")[2]
      schema_name = RESOURCE_SCHEMAS[resource]
      next unless schema_name

      methods.each do |_method, operation|
        responses = operation["responses"] || {}
        responses.each do |_status, response|
          content = response["content"]
          next unless content && content["application/json"]

          schema = content["application/json"]["schema"]
          next unless schema && schema["properties"] && schema["properties"]["data"]

          data = schema["properties"]["data"]
          if data["type"] == "array" && data["items"] && data["items"]["properties"]
            unless spec["components"]["schemas"][schema_name]
              spec["components"]["schemas"][schema_name] = data["items"].dup
            end
            data["items"] = { "$ref" => "#/components/schemas/#{schema_name}" }
          elsif data["properties"]
            unless spec["components"]["schemas"][schema_name]
              spec["components"]["schemas"][schema_name] = data.dup
            end
            schema["properties"]["data"] = { "$ref" => "#/components/schemas/#{schema_name}" }
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
