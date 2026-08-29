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
          if tc.dig("expect", "fieldErrors")
            assert_equal tc.dig("expect", "fieldErrors"), e.field_errors_by_field, "[#{tc["name"]}] field errors"
          end
        end
      end

      assert_request_count(tc) if tc.dig("expect", "requestCount")
      assert_no_outbound_request(tc) if tc.dig("expect", "assertNoOutboundRequest")
    end

    def assert_no_outbound_request(tc)
      expected = tc["mockResponses"].length
      actual = WebMock::RequestRegistry.instance.requested_signatures.hash.values.sum
      assert_equal expected, actual, "[#{tc["name"]}] expected no outbound request beyond mocks (got #{actual} calls)"
    end

    def execute_operation(client, tc)
      case tc["operation"]
      when "list_customers"
        client.list_customers
      when "list_customers_paginated"
        result = client.list_customers
        paginator = result.paginator
        while paginator.has_next?
          result = paginator.next_page
        end
        result
      when "list_customers_with_params"
        client.list_customers(params: tc["query"])
      when "list_customers_with_params_paginated"
        result = client.list_customers(params: tc["query"])
        paginator = result.paginator
        while paginator.has_next?
          result = paginator.next_page
        end
        result
      when "show_customer"
        client.show_customer(tc.dig("pathParams", "id"))
      when "create_customer"
        client.create_customer(tc.dig("requestBody", "customer") || {})
      when "update_customer"
        client.update_customer(tc.dig("pathParams", "id"), tc.dig("requestBody", "customer") || {})
      when "list_vehicles"
        client.list_vehicles
      when "show_vehicle"
        client.show_vehicle(tc.dig("pathParams", "id"))
      when "create_vehicle"
        client.create_vehicle(tc.dig("requestBody", "vehicle") || {})
      when "update_vehicle"
        client.update_vehicle(tc.dig("pathParams", "id"), tc.dig("requestBody", "vehicle") || {})
      when "delete_vehicle"
        client.delete_vehicle(tc.dig("pathParams", "id"))
      when "decode_vin"
        client.decode_vin(tc.dig("query", "vin"))
      when "check_duplicate"
        client.check_duplicate(tc.dig("query", "vin"))
      when "list_work_orders"
        client.list_work_orders
      when "list_work_orders_paginated"
        result = client.list_work_orders
        paginator = result.paginator
        while paginator.has_next?
          result = paginator.next_page
        end
        result
      when "show_work_order"
        client.show_work_order(tc.dig("pathParams", "id"))
      when "create_work_order"
        client.create_work_order(tc.dig("requestBody", "work_order") || {})
      when "update_work_order"
        client.update_work_order(tc.dig("pathParams", "id"), tc.dig("requestBody", "work_order") || {})
      when "delete_work_order"
        client.delete_work_order(tc.dig("pathParams", "id"))
      when "list_account"
        client.list_account
      when "show_location"
        client.show_location(tc.dig("pathParams", "id"))
      when "list_customers_drivers"
        client.list_drivers(tc.dig("pathParams", "customer_id"))
      when "show_driver"
        client.show_driver(tc.dig("pathParams", "customer_id"), tc.dig("pathParams", "id"))
      when "create_driver"
        client.create_driver(tc.dig("pathParams", "customer_id"), tc.dig("requestBody", "driver") || {})
      when "update_driver"
        client.update_driver(tc.dig("pathParams", "customer_id"), tc.dig("pathParams", "id"), tc.dig("requestBody", "driver") || {})
      when "delete_driver"
        client.delete_driver(tc.dig("pathParams", "customer_id"), tc.dig("pathParams", "id"))
      when "list_customers_statements"
        client.list_statements(tc.dig("pathParams", "customer_id"))
      when "show_statement"
        client.show_statement(tc.dig("pathParams", "id"))
      when "list_vendors"
        client.list_vendors
      when "show_vendor"
        client.show_vendor(tc.dig("pathParams", "id"))
      when "show_work_order_estimate"
        client.show_work_order_estimate(tc.dig("pathParams", "work_order_id"))
      when "show_work_order_wip"
        client.show_work_order_wip(tc.dig("pathParams", "work_order_id"))
      when "show_work_order_inspection"
        client.show_work_order_inspection(tc.dig("pathParams", "work_order_id"))
      when "show_work_order_parts"
        client.show_work_order_parts(tc.dig("pathParams", "work_order_id"))
      when "show_work_order_payments"
        client.show_work_order_payments(tc.dig("pathParams", "work_order_id"))
      when "create_work_order_payment"
        client.create_work_order_payment(tc.dig("pathParams", "work_order_id"), tc.dig("requestBody", "payment") || {})
      when "merge_customer"
        client.merge_customer(tc.dig("pathParams", "id"), tc.dig("requestBody", "source_customer_id"))
      when "lookup_customer"
        client.lookup_customer(query: tc.dig("query", "query"))
      when "check_customer_duplicate"
        client.check_customer_duplicate(first_name: tc.dig("query", "first_name"), last_name: tc.dig("query", "last_name"))
      when "transfer_vehicle"
        client.transfer_vehicle(tc.dig("pathParams", "id"), customer_id: tc.dig("requestBody", "customer_id"), mode: tc.dig("requestBody", "mode") || "vehicle_only")
      when "merge_vehicle"
        client.merge_vehicle(tc.dig("pathParams", "id"), source_vehicle_id: tc.dig("requestBody", "source_vehicle_id"))
      when "prefill_vehicle"
        client.prefill_vehicle(vin: tc.dig("query", "vin"))
      when "lookup_vehicle"
        client.lookup_vehicle(query: tc.dig("query", "query"))
      when "check_vehicle_duplicate"
        client.check_vehicle_duplicate(vin: tc.dig("query", "vin"))
      when "list_customer_vehicles"
        client.list_customer_vehicles(tc.dig("pathParams", "customer_id"))
      when "list_customer_work_orders"
        client.list_customer_work_orders(tc.dig("pathParams", "customer_id"))
      when "list_vehicle_work_orders"
        client.list_vehicle_work_orders(tc.dig("pathParams", "vehicle_id"))
      when "list_tags"
        client.list_tags
      when "update_tags"
        client.update_tags(customer_tags: tc.dig("requestBody", "customer_tags"))
      when "create_customer_tag"
        client.create_customer_tag(name: tc.dig("requestBody", "name"))
      when "create_vehicle_tag"
        client.create_vehicle_tag(name: tc.dig("requestBody", "name"))
      when "list_service_categories"
        client.list_service_categories
      when "create_service_category"
        client.create_service_category(name: tc.dig("requestBody", "service_category", "name"), service_type: tc.dig("requestBody", "service_category", "service_type"), icon: tc.dig("requestBody", "service_category", "icon"))
      when "update_service_category"
        client.update_service_category(tc.dig("pathParams", "id"), name: tc.dig("requestBody", "service_category", "name"))
      when "delete_service_category"
        client.delete_service_category(tc.dig("pathParams", "id"))
      when "deactivate_service_category"
        client.deactivate_service_category(tc.dig("pathParams", "id"))
      when "reactivate_service_category"
        client.reactivate_service_category(tc.dig("pathParams", "id"))
      when "move_up_service_category"
        client.move_up_service_category(tc.dig("pathParams", "id"))
      when "move_down_service_category"
        client.move_down_service_category(tc.dig("pathParams", "id"))
      when "seed_defaults_service_categories"
        client.seed_defaults_service_categories
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
