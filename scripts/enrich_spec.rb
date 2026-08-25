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

  def self.call(input)
    spec = Marshal.load(Marshal.dump(input)) # deep copy

    spec["servers"] = [{ "url" => "https://app.wenmarpro.com" }]
    spec["paths"] ||= {}

    extract_components!(spec)

    spec["paths"].each do |path, methods|
      resource = path.split("/")[2] # "/api/customers" => "customers"
      methods.each do |method, operation|
        operation["operationId"] = derive_operation_id(path, method)
        operation["tags"] = [resource] if resource
        operation["description"] = derive_description(operation["operationId"])
      end
    end

    spec
  end

  def self.derive_operation_id(path, method)
    resource = path.split("/")[2] # "customers", "vehicles", "work_orders"
    return nil unless resource

    singularized = SINGULAR_OVERRIDES[resource] || resource.chomp("s")
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
