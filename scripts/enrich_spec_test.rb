require "minitest/autorun"
require "yaml"

class EnrichSpecTest < Minitest::Test
  def setup
    @input = {
      "openapi" => "3.0.0",
      "info" => { "title" => "Wenmar Pro API", "version" => "1.0.0" },
      "paths" => {
        "/api/customers" => {
          "get" => { "summary" => "index", "responses" => { "200" => customer_list_response } },
          "post" => { "summary" => "create", "responses" => { "200" => customer_show_response } }
        },
        "/api/customers/{id}" => {
          "get" => { "summary" => "show", "responses" => { "200" => customer_show_response, "404" => error_response } }
        },
        "/api/vehicles/{id}" => {
          "get" => { "summary" => "show", "responses" => { "200" => vehicle_show_response, "404" => error_response } }
        },
        "/api/work_orders" => {
          "get" => { "summary" => "index", "responses" => { "200" => work_order_list_response } }
        },
        "/api/work_orders/{id}" => {
          "get" => { "summary" => "show", "responses" => { "200" => work_order_show_response } }
        }
      }
    }
  end

  def test_adds_operation_id_to_each_operation
    result = enrich(@input)
    assert_equal "list_customers", result["paths"]["/api/customers"]["get"]["operationId"]
    assert_equal "create_customer", result["paths"]["/api/customers"]["post"]["operationId"]
    assert_equal "show_customer", result["paths"]["/api/customers/{id}"]["get"]["operationId"]
    assert_equal "show_vehicle", result["paths"]["/api/vehicles/{id}"]["get"]["operationId"]
    assert_equal "list_work_orders", result["paths"]["/api/work_orders"]["get"]["operationId"]
    assert_equal "show_work_order", result["paths"]["/api/work_orders/{id}"]["get"]["operationId"]
  end

  def test_adds_tags_grouped_by_resource
    result = enrich(@input)
    assert_equal ["customers"], result["paths"]["/api/customers"]["get"]["tags"]
    assert_equal ["customers"], result["paths"]["/api/customers/{id}"]["get"]["tags"]
    assert_equal ["vehicles"], result["paths"]["/api/vehicles/{id}"]["get"]["tags"]
    assert_equal ["work_orders"], result["paths"]["/api/work_orders"]["get"]["tags"]
  end

  def test_adds_servers_with_base_url
    result = enrich(@input)
    assert_equal [{ "url" => "https://app.wenmarpro.com" }], result["servers"]
  end

  def test_adds_human_readable_descriptions
    result = enrich(@input)
    assert_match(/List all customers/i, result["paths"]["/api/customers"]["get"]["description"])
    assert_match(/Create a customer/i, result["paths"]["/api/customers"]["post"]["description"])
    assert_match(/Show a customer/i, result["paths"]["/api/customers/{id}"]["get"]["description"])
  end

  def test_does_not_modify_the_input_spec
    original = Marshal.load(Marshal.dump(@input))
    enrich(@input)
    assert_equal original, @input
  end

  def test_extracts_reusable_schema_components
    result = enrich(@input)
    refute_nil result["components"]
    refute_nil result["components"]["schemas"]
    refute_nil result["components"]["schemas"]["Customer"]
    refute_nil result["components"]["schemas"]["Vehicle"]
    refute_nil result["components"]["schemas"]["WorkOrder"]
    refute_nil result["components"]["schemas"]["Error"]
  end

  def test_replaces_inline_schemas_with_ref
    result = enrich(@input)
    customer_schema = result["paths"]["/api/customers"]["get"]["responses"]["200"]["content"]["application/json"]["schema"]
    refute_nil customer_schema["properties"]["data"]["items"]["$ref"]
    assert_equal "#/components/schemas/Customer", customer_schema["properties"]["data"]["items"]["$ref"]
  end

  def test_error_responses_use_error_component_ref
    result = enrich(@input)
    error_schema = result["paths"]["/api/customers/{id}"]["get"]["responses"]["404"]["content"]["application/json"]["schema"]
    assert_equal "#/components/schemas/Error", error_schema["properties"]["error"]["$ref"]
  end

  def test_convention_based_operation_id_for_list
    input = make_spec_with("/api/vehicles" => { "get" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "list_vehicles", result["paths"]["/api/vehicles"]["get"]["operationId"]
  end

  def test_convention_based_operation_id_for_show
    input = make_spec_with("/api/vehicles/{id}" => { "get" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "show_vehicle", result["paths"]["/api/vehicles/{id}"]["get"]["operationId"]
  end

  def test_convention_based_operation_id_for_create
    input = make_spec_with("/api/vehicles" => { "post" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "create_vehicle", result["paths"]["/api/vehicles"]["post"]["operationId"]
  end

  def test_convention_based_operation_id_for_update
    input = make_spec_with("/api/customers/{id}" => { "patch" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "update_customer", result["paths"]["/api/customers/{id}"]["patch"]["operationId"]
  end

  def test_convention_based_operation_id_for_delete
    input = make_spec_with("/api/customers/{id}" => { "delete" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "delete_customer", result["paths"]["/api/customers/{id}"]["delete"]["operationId"]
  end

  def test_convention_handles_work_orders_singularization
    input = make_spec_with("/api/work_orders" => { "post" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "create_work_order", result["paths"]["/api/work_orders"]["post"]["operationId"]
  end

  private

  def make_spec_with(paths)
    {
      "openapi" => "3.0.0",
      "info" => { "title" => "Wenmar Pro API", "version" => "1.0.0" },
      "paths" => paths
    }
  end

  def enrich(spec)
    require_relative "enrich_spec"
    EnrichSpec.call(spec)
  end

  def json_response(schema)
    { "description" => "response", "content" => { "application/json" => { "schema" => schema } } }
  end

  def customer_list_response
    json_response(
      "type" => "object",
      "properties" => {
        "data" => {
          "type" => "array",
          "items" => {
            "type" => "object",
            "properties" => { "id" => { "type" => "integer" }, "full_name" => { "type" => "string" } }
          }
        }
      }
    )
  end

  def customer_show_response
    json_response(
      "type" => "object",
      "properties" => {
        "data" => {
          "type" => "object",
          "properties" => { "id" => { "type" => "integer" }, "full_name" => { "type" => "string" } }
        }
      }
    )
  end

  def vehicle_show_response
    json_response(
      "type" => "object",
      "properties" => {
        "data" => {
          "type" => "object",
          "properties" => { "id" => { "type" => "integer" }, "make" => { "type" => "string" } }
        }
      }
    )
  end

  def work_order_list_response
    json_response(
      "type" => "object",
      "properties" => {
        "data" => {
          "type" => "array",
          "items" => {
            "type" => "object",
            "properties" => { "id" => { "type" => "integer" }, "status" => { "type" => "string" } }
          }
        }
      }
    )
  end

  def work_order_show_response
    json_response(
      "type" => "object",
      "properties" => {
        "data" => {
          "type" => "object",
          "properties" => { "id" => { "type" => "integer" }, "status" => { "type" => "string" } }
        }
      }
    )
  end

  def error_response
    json_response(
      "type" => "object",
      "properties" => {
        "error" => {
          "type" => "object",
          "properties" => { "code" => { "type" => "string" }, "message" => { "type" => "string" }, "details" => { "type" => "object" } }
        }
      }
    )
  end
end
