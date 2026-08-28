require_relative "spec_helper"

module Wenmar
  class ClientTest < TestCase
    def test_initialize_requires_token
      assert_raises(ArgumentError) { Client.new(token: "") }
    end

    def test_initialize_sets_base_url
      client = Client.new(token: "test", base_url: "https://api.example.com")
      assert_equal "https://api.example.com", client.base_url
    end

    def test_list_customers
      stub_api(:get, "/customers", [{ id: 1, full_name: "Jane" }])
      client = Client.new(token: "test", base_url: @base_url)
      result = client.list_customers
      assert_kind_of Array, result
      assert_equal 1, result.first["id"]
    end

    def test_show_customer
      stub_api(:get, "/customers/1", { id: 1, full_name: "Jane" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.show_customer(1)
      assert_equal 1, result["id"]
    end

    def test_create_customer
      stub_api(:post, "/customers", { id: 1, full_name: "Jane" }, status: 201)
      client = Client.new(token: "test", base_url: @base_url)
      result = client.create_customer(first_name: "Jane", last_name: "Doe")
      assert_equal "Jane", result["full_name"]
    end

    def test_update_customer
      stub_api(:patch, "/customers/1", { id: 1, full_name: "Jane Doe" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.update_customer(1, first_name: "Jane")
      assert_equal 1, result["id"]
    end

    def test_show_vehicle
      stub_api(:get, "/vehicles/1", { id: 1, make: "Toyota", model: "Camry", year: 2020 })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.show_vehicle(1)
      assert_equal "Toyota", result["make"]
    end

    def test_list_vehicles
      stub_api(:get, "/vehicles", [{ id: 1, make: "Honda" }])
      client = Client.new(token: "test", base_url: @base_url)
      result = client.list_vehicles
      assert_kind_of Array, result
      assert_equal 1, result.first["id"]
    end

    def test_create_vehicle
      stub_api(:post, "/vehicles", { id: 1, make: "Honda", model: "Civic", year: 2020 }, status: 201)
      client = Client.new(token: "test", base_url: @base_url)
      result = client.create_vehicle(make: "Honda", model: "Civic", year: 2020, customer_id: 1)
      assert_equal "Honda", result["make"]
    end

    def test_update_vehicle
      stub_api(:patch, "/vehicles/1", { id: 1, make: "Toyota" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.update_vehicle(1, make: "Toyota")
      assert_equal "Toyota", result["make"]
    end

    def test_delete_vehicle
      stub_api(:delete, "/vehicles/1", nil, status: 204)
      client = Client.new(token: "test", base_url: @base_url)
      assert_nil client.delete_vehicle(1)
    end

    def test_decode_vin
      stub_request(:get, "#{@base_url}/vehicles/vin_decode")
        .with(query: { vin: "1HGCM82633A004352" })
        .to_return(status: 200, body: { make: "Honda", model: "Civic", vin: "1HGCM82633A004352" }.to_json, headers: { "Content-Type" => "application/json" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.decode_vin("1HGCM82633A004352")
      assert_equal "Honda", result["make"]
    end

    def test_check_duplicate
      stub_request(:get, "#{@base_url}/vehicles/check_duplicate")
        .with(query: { vin: "ABC123" })
        .to_return(status: 200, body: { matches: [{ id: 1, display_name: "Toyota Camry", url: "/vehicles/1", reasons: ["vin"] }] }.to_json, headers: { "Content-Type" => "application/json" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.check_duplicate("ABC123")
      assert_equal 1, result["matches"].first["id"]
    end

    def test_show_customer_not_found
      stub_api(:get, "/customers/999", { error: { code: "not_found", message: "Not found", details: {} } }, status: 404)
      client = Client.new(token: "test", base_url: @base_url)
      err = assert_raises(Wenmar::Error) { client.show_customer(999) }
      assert_equal "not_found", err.code
      assert_equal 404, err.status
    end

    def test_sets_bearer_auth_header
      stub = stub_request(:get, "#{@base_url}/customers")
             .with(headers: { "Authorization" => "Bearer my-token" })
             .to_return(status: 200, body: [].to_json, headers: { "Content-Type" => "application/json" })
      client = Client.new(token: "my-token", base_url: @base_url)
      client.list_customers
      assert_requested stub
    end

    def test_retry_after_is_honored_on_429
      # First request 429 + Retry-After, then success.
      stub_request(:get, "#{@base_url}/customers")
        .to_return({ status: 429, body: { error: "rate_limited" }.to_json, headers: { "Retry-After" => "0" } },
                   { status: 200, body: [].to_json })
      client = Client.new(token: "my-token", base_url: @base_url)
      result = client.list_customers
      assert_kind_of Array, result
    end

    def test_post_on_500_not_retried
      requests = []
      stub_request(:post, "#{@base_url}/customers").to_return do |_request|
        requests << 1
        {
          status: 500,
          body: { error: { code: "internal_error", message: "fail", details: {} } }.to_json,
          headers: { "Content-Type" => "application/json" }
        }
      end

      client = Client.new(token: "my-token", base_url: @base_url)
      assert_raises(Wenmar::Error) { client.create_customer(first_name: "Test") }
      assert_equal 1, requests.size, "POST must not retry on 500 (got #{requests.size} requests)"
    end

    def test_post_on_429_retried
      requests = []
      stub_request(:post, "#{@base_url}/customers").to_return do |_request|
        requests << 1
        if requests.size == 1
          {
            status: 429,
            body: { error: { code: "rate_limited", message: "slow", details: {} } }.to_json,
            headers: { "Content-Type" => "application/json", "Retry-After" => "0" }
          }
        else
          {
            status: 201,
            body: { id: 1 }.to_json,
            headers: { "Content-Type" => "application/json" }
          }
        end
      end

      client = Client.new(token: "my-token", base_url: @base_url)
      result = client.create_customer(first_name: "Test")
      assert_equal 2, requests.size, "429 on POST should retry (got #{requests.size})"
      assert_equal 1, result["id"]
    end

    def test_conditional_get_returns_cached_body_on_304
      body = { id: 1, full_name: "Jane" }.to_json
      stub_request(:get, "#{@base_url}/customers/1")
        .to_return(status: 200, body: body, headers: { "Content-Type" => "application/json", "ETag" => "\"abc123\"" })
      stub_request(:get, "#{@base_url}/customers/1")
        .with(headers: { "If-None-Match" => "\"abc123\"" })
        .to_return(status: 304, headers: { "Content-Type" => "application/json" })

      client = Client.new(token: "my-token", base_url: @base_url)
      first = client.show_customer(1)
      assert_equal 1, first["id"]

      second = client.show_customer(1)
      assert_equal 1, second["id"], "expected cached body returned on 304"
      assert_requested stub_request(:get, "#{@base_url}/customers/1")
        .with(headers: { "If-None-Match" => "\"abc123\"" })
    end

    def test_list_account
      stub_api(:get, "/account", { "id" => 1, "name" => "Main Shop" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.list_account
      assert_equal 1, result["id"]
      assert_equal "Main Shop", result["name"]
    end

    def test_show_location
      stub_api(:get, "/locations/1", { "id" => 1, "name" => "Bay 1" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.show_location("1")
      assert_equal 1, result["id"]
      assert_equal "Bay 1", result["name"]
    end

    private

    def stub_api(method, path, body, status: 200)
      stub_request(method, "#{@base_url}#{path}")
        .to_return(status: status, body: (body || "").to_json, headers: { "Content-Type" => "application/json" })
    end
  end
end
