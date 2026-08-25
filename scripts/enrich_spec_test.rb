require "minitest/autorun"
require "yaml"

class EnrichSpecTest < Minitest::Test
  def setup
    @input = {
      "openapi" => "3.0.0",
      "info" => { "title" => "Wenmar Pro API", "version" => "1.0.0" },
      "paths" => {
        "/api/customers" => {
          "get" => { "summary" => "index", "responses" => {} },
          "post" => { "summary" => "create", "responses" => {} }
        },
        "/api/customers/{id}" => {
          "get" => { "summary" => "show", "responses" => {} }
        },
        "/api/vehicles/{id}" => {
          "get" => { "summary" => "show", "responses" => {} }
        },
        "/api/work_orders" => {
          "get" => { "summary" => "index", "responses" => {} }
        },
        "/api/work_orders/{id}" => {
          "get" => { "summary" => "show", "responses" => {} }
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

  private

  def enrich(spec)
    require_relative "enrich_spec"
    EnrichSpec.call(spec)
  end
end
