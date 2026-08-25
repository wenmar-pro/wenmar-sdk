require_relative "spec_helper"

module WenmarPro
  class ClientTest < TestCase
    def test_initialize_requires_token
      assert_raises(ArgumentError) { Client.new(token: "") }
    end

    def test_initialize_sets_base_url
      client = Client.new(token: "test", base_url: "https://api.example.com")
      assert_equal "https://api.example.com", client.base_url
    end

    def test_list_customers
      stub_api(:get, "/api/customers", { data: [{ id: 1, full_name: "Jane" }] })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.list_customers
      assert_kind_of Array, result["data"]
      assert_equal 1, result["data"].first["id"]
    end

    def test_show_customer
      stub_api(:get, "/api/customers/1", { data: { id: 1, full_name: "Jane" } })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.show_customer(1)
      assert_equal 1, result["data"]["id"]
    end

    def test_create_customer
      stub_api(:post, "/api/customers", { data: { id: 1, full_name: "Jane" } }, status: 201)
      client = Client.new(token: "test", base_url: @base_url)
      result = client.create_customer(full_name: "Jane")
      assert_equal "Jane", result["data"]["full_name"]
    end

    def test_show_customer_not_found
      stub_api(:get, "/api/customers/999", { error: { code: "not_found", message: "Not found", details: {} } }, status: 404)
      client = Client.new(token: "test", base_url: @base_url)
      err = assert_raises(WenmarPro::Error) { client.show_customer(999) }
      assert_equal "not_found", err.code
      assert_equal 404, err.status
    end

    def test_sets_bearer_auth_header
      stub = stub_request(:get, "#{@base_url}/api/customers")
             .with(headers: { "Authorization" => "Bearer my-token" })
             .to_return(status: 200, body: { data: [] }.to_json, headers: { "Content-Type" => "application/json" })
      client = Client.new(token: "my-token", base_url: @base_url)
      client.list_customers
      assert_requested stub
    end

    private

    def stub_api(method, path, body, status: 200)
      stub_request(method, "#{@base_url}#{path}")
        .to_return(status: status, body: body.to_json, headers: { "Content-Type" => "application/json" })
    end
  end
end
