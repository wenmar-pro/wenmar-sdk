require "minitest/autorun"
require "yaml"

class EnrichSpecTest < Minitest::Test
  def setup
    @input = {
      "openapi" => "3.0.0",
      "info" => { "title" => "Wenmar Pro API", "version" => "1.0.0" },
      "paths" => {
        "/customers" => {
          "get" => { "summary" => "index", "responses" => { "200" => customer_list_response } },
          "post" => { "summary" => "create", "responses" => { "200" => customer_show_response } }
        },
        "/customers/{id}" => {
          "get" => { "summary" => "show", "responses" => { "200" => customer_show_response, "404" => error_response } },
          "patch" => { "summary" => "update", "responses" => { "200" => customer_show_response } },
          "delete" => { "summary" => "destroy", "responses" => { "204" => { "description" => "no content" } } }
        },
        "/vehicles" => {
          "get" => { "summary" => "index", "responses" => { "200" => vehicle_list_response } },
          "post" => { "summary" => "create", "responses" => { "201" => vehicle_show_response } }
        },
        "/vehicles/{id}" => {
          "get" => { "summary" => "show", "responses" => { "200" => vehicle_show_response } },
          "patch" => { "summary" => "update", "responses" => { "200" => vehicle_show_response } },
          "delete" => { "summary" => "destroy", "responses" => { "204" => { "description" => "no content" } } }
        },
        "/vehicles/vin_decode" => {
          "get" => { "summary" => "vin_decode", "responses" => { "200" => vin_decode_response } }
        },
        "/vehicles/check_duplicate" => {
          "get" => { "summary" => "check_duplicate", "responses" => { "200" => check_duplicate_response } }
        },
        "/work_orders" => {
          "get" => { "summary" => "index", "responses" => { "200" => work_order_list_response } }
        },
        "/work_orders/{id}" => {
          "get" => { "summary" => "show", "responses" => { "200" => work_order_show_response } }
        }
      }
    }
  end

  def test_adds_operation_id_to_each_operation
    result = enrich(@input)
    assert_equal "list_customers", result["paths"]["/customers"]["get"]["operationId"]
    assert_equal "create_customer", result["paths"]["/customers"]["post"]["operationId"]
    assert_equal "show_customer", result["paths"]["/customers/{id}"]["get"]["operationId"]
    assert_equal "update_customer", result["paths"]["/customers/{id}"]["patch"]["operationId"]
    assert_equal "delete_customer", result["paths"]["/customers/{id}"]["delete"]["operationId"]
    assert_equal "show_vehicle", result["paths"]["/vehicles/{id}"]["get"]["operationId"]
    assert_equal "list_work_orders", result["paths"]["/work_orders"]["get"]["operationId"]
    assert_equal "show_work_order", result["paths"]["/work_orders/{id}"]["get"]["operationId"]
  end

  def test_adds_sub_action_operation_ids
    result = enrich(@input)
    assert_equal "decode_vin", result["paths"]["/vehicles/vin_decode"]["get"]["operationId"]
    assert_equal "check_duplicate", result["paths"]["/vehicles/check_duplicate"]["get"]["operationId"]
  end

  def test_adds_tags_grouped_by_resource
    result = enrich(@input)
    assert_equal ["customers"], result["paths"]["/customers"]["get"]["tags"]
    assert_equal ["customers"], result["paths"]["/customers/{id}"]["get"]["tags"]
    assert_equal ["vehicles"], result["paths"]["/vehicles/{id}"]["get"]["tags"]
    assert_equal ["work_orders"], result["paths"]["/work_orders"]["get"]["tags"]
  end

  def test_adds_servers_with_base_url
    result = enrich(@input)
    assert_equal [{ "url" => "https://app.wenmarpro.com" }], result["servers"]
  end

  def test_adds_human_readable_descriptions
    result = enrich(@input)
    assert_match(/List all customers/i, result["paths"]["/customers"]["get"]["description"])
    assert_match(/Create a customer/i, result["paths"]["/customers"]["post"]["description"])
    assert_match(/Show a customer/i, result["paths"]["/customers/{id}"]["get"]["description"])
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

  def test_replaces_array_responses_with_ref
    result = enrich(@input)
    customer_list_schema = result["paths"]["/customers"]["get"]["responses"]["200"]["content"]["application/json"]["schema"]
    assert_equal "array", customer_list_schema["type"]
    refute_nil customer_list_schema["items"]["$ref"]
    assert_equal "#/components/schemas/Customer", customer_list_schema["items"]["$ref"]
  end

  def test_replaces_bare_object_responses_with_ref
    result = enrich(@input)
    customer_show_schema = result["paths"]["/customers/{id}"]["get"]["responses"]["200"]["content"]["application/json"]["schema"]
    assert_equal "#/components/schemas/Customer", customer_show_schema["$ref"]
  end

  def test_error_responses_use_error_component_ref
    result = enrich(@input)
    error_schema = result["paths"]["/customers/{id}"]["get"]["responses"]["404"]["content"]["application/json"]["schema"]
    assert_equal "#/components/schemas/Error", error_schema["properties"]["error"]["$ref"]
  end

  def test_sub_action_responses_are_not_extracted_as_resources
    result = enrich(@input)
    decode_schema = result["paths"]["/vehicles/vin_decode"]["get"]["responses"]["200"]["content"]["application/json"]["schema"]
    refute_nil decode_schema["properties"]
    refute_equal "#/components/schemas/Vehicle", decode_schema["$ref"]
  end

  def test_convention_based_operation_id_for_list
    input = make_spec_with("/vehicles" => { "get" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "list_vehicles", result["paths"]["/vehicles"]["get"]["operationId"]
  end

  def test_convention_based_operation_id_for_show
    input = make_spec_with("/vehicles/{id}" => { "get" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "show_vehicle", result["paths"]["/vehicles/{id}"]["get"]["operationId"]
  end

  def test_convention_based_operation_id_for_create
    input = make_spec_with("/vehicles" => { "post" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "create_vehicle", result["paths"]["/vehicles"]["post"]["operationId"]
  end

  def test_convention_based_operation_id_for_update
    input = make_spec_with("/customers/{id}" => { "patch" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "update_customer", result["paths"]["/customers/{id}"]["patch"]["operationId"]
  end

  def test_convention_based_operation_id_for_delete
    input = make_spec_with("/customers/{id}" => { "delete" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "delete_customer", result["paths"]["/customers/{id}"]["delete"]["operationId"]
  end

  def test_convention_handles_work_orders_singularization
    input = make_spec_with("/work_orders" => { "post" => { "responses" => {} } })
    result = enrich(input)
    assert_equal "create_work_order", result["paths"]["/work_orders"]["post"]["operationId"]
  end

  def test_adds_curl_example_to_each_operation
    result = enrich(@input)
    get_example = result["paths"]["/customers"]["get"]["x-curl-example"]
    assert_includes get_example, "curl"
    assert_includes get_example, "https://app.wenmarpro.com/customers.json"
    assert_includes get_example, "Bearer $WENMAR_TOKEN"

    create_example = result["paths"]["/customers"]["post"]["x-curl-example"]
    assert_includes create_example, "-X POST"

    show_example = result["paths"]["/customers/{id}"]["get"]["x-curl-example"]
    assert_includes show_example, "https://app.wenmarpro.com/customers/<id>.json"
  end

  def test_applies_example_overrides_to_schemas
    result = enrich(@input)
    customer = result["components"]["schemas"]["Customer"]
    props = customer["properties"]
    assert_equal 7, props["id"]["example"]
    assert_equal "Jane Doe", props["full_name"]["example"]
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
      "type" => "array",
      "items" => {
        "type" => "object",
        "properties" => { "id" => { "type" => "integer" }, "full_name" => { "type" => "string" } }
      }
    )
  end

  def customer_show_response
    json_response(
      "type" => "object",
      "properties" => { "id" => { "type" => "integer" }, "full_name" => { "type" => "string" } }
    )
  end

  def vehicle_list_response
    json_response(
      "type" => "array",
      "items" => {
        "type" => "object",
        "properties" => { "id" => { "type" => "integer" }, "make" => { "type" => "string" } }
      }
    )
  end

  def vehicle_show_response
    json_response(
      "type" => "object",
      "properties" => { "id" => { "type" => "integer" }, "make" => { "type" => "string" } }
    )
  end

  def vin_decode_response
    json_response(
      "type" => "object",
      "properties" => { "make" => { "type" => "string" }, "model" => { "type" => "string" }, "vin" => { "type" => "string" } }
    )
  end

  def check_duplicate_response
    json_response(
      "type" => "object",
      "properties" => { "matches" => { "type" => "array", "items" => { "type" => "object", "properties" => { "id" => { "type" => "integer" } } } } }
    )
  end

  def work_order_list_response
    json_response(
      "type" => "array",
      "items" => {
        "type" => "object",
        "properties" => { "id" => { "type" => "integer" }, "status" => { "type" => "string" } }
      }
    )
  end

  def work_order_show_response
    json_response(
      "type" => "object",
      "properties" => { "id" => { "type" => "integer" }, "status" => { "type" => "string" } }
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
