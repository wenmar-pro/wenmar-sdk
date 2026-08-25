require "minitest/autorun"
require "json"
require "webmock/minitest"
require "wenmar"

WebMock.disable_net_connect!

module Conformance
  class ConformanceTest < Minitest::Test
    TESTS_DIR = File.expand_path("../tests", __dir__)
    BASE_URL = "https://api.example.com"

    def test_conformance
      load_cases.each do |tc|
        run_case(tc)
      end
    end

    private

    def load_cases
      Dir.glob(File.join(TESTS_DIR, "*.json")).flat_map do |file|
        JSON.parse(File.read(file))
      end
    end

    def run_case(tc)
      WebMock.reset!
      stub_responses(tc)
      client = Wenmar::Client.new(token: "test-token", base_url: BASE_URL)

      begin
        result = execute_operation(client, tc)
        if tc.dig("expect", "noError")
          assert_body_path(result, tc.dig("expect", "responseBody"), tc["name"]) if tc.dig("expect", "responseBody")
        else
          flunk "[#{tc["name"]}] expected error, got success"
        end
      rescue Wenmar::Error => e
        if tc.dig("expect", "noError")
          flunk "[#{tc["name"]}] expected success, got error: #{e.message}"
        else
          assert_equal tc.dig("expect", "errorCode"), e.code, "[#{tc["name"]}] error code" if tc.dig("expect", "errorCode")
          assert_equal tc.dig("expect", "errorStatus"), e.status, "[#{tc["name"]}] error status" if tc.dig("expect", "errorStatus")
        end
      end

      assert_request_count(tc) if tc.dig("expect", "requestCount")
    end

    def execute_operation(client, tc)
      case tc["operation"]
      when "list_customers"
        client.list_customers
      when "list_customers_paginated"
        result = client.list_customers
        while result.paginator.has_next?
          result = result.paginator.next_page
        end
        result
      when "show_customer"
        client.show_customer(tc.dig("pathParams", "id"))
      when "create_customer"
        client.create_customer(tc.dig("requestBody", "customer") || {})
      when "show_vehicle"
        client.show_vehicle(tc.dig("pathParams", "id"))
      when "list_work_orders"
        client.list_work_orders
      when "list_work_orders_paginated"
        result = client.list_work_orders
        while result.paginator.has_next?
          result = result.paginator.next_page
        end
        result
      when "show_work_order"
        client.show_work_order(tc.dig("pathParams", "id"))
      else
        raise "unknown operation: #{tc["operation"]}"
      end
    end

    def stub_responses(tc)
      responses = tc["mockResponses"].map do |resp|
        headers = { "Content-Type" => "application/json" }
        (resp["headers"] || {}).each do |k, v|
          headers[k] = v.gsub("{server}", BASE_URL)
        end
        body = resp["body"].nil? ? "" : resp["body"].to_json
        { status: resp["status"], body: body, headers: headers }
      end

      # Match the path with any query string (pagination appends ?page=N).
      stub_request(tc["method"].downcase.to_sym, /#{Regexp.escape(BASE_URL)}#{Regexp.escape(tc["path"])}(\?.*)?\z/)
        .to_return(responses)
    end

    def assert_request_count(tc)
      expected = tc.dig("expect", "requestCount")
      actual = WebMock::RequestRegistry.instance.requested_signatures.hash.values.sum
      assert_equal expected, actual, "[#{tc["name"]}] expected #{expected} requests, got #{actual}"
    end

    def assert_body_path(result, assertion, name)
      value = navigate_path(result, assertion["path"])
      assert_equal assertion["equals"], value, "[#{name}] expected #{assertion["path"]} to equal #{assertion["equals"]}"
    end

    def navigate_path(obj, path)
      path.split(".").reduce(obj) do |current, part|
        if part =~ /\A\d+\z/
          current[part.to_i]
        else
          current[part]
        end
      end
    end
  end
end
