require_relative "spec_helper"

module Wenmar
  class ResourcesTest < TestCase
    def test_all_generated_methods_exist
      client = Client.new(token: "test", base_url: @base_url)
      expected = %i[
        list_account list_customer_tags create_customer_tag delete_customer_tag update_customer_tag
        list_customers create_customer check_customer_duplicate lookup_customer
        list_customers_drivers create_driver delete_driver show_driver update_driver
        list_customers_statements list_customers_vehicles list_customers_vehicles_history
        list_customers_work_orders delete_customer show_customer update_customer merge_customer
        show_location list_service_categories create_service_category seed_defaults_service_categories
        delete_service_category update_service_category
        list_tags update_tags show_statement list_users list_permission_groups list_vehicle_tags
        create_vehicle_tag delete_vehicle_tag update_vehicle_tag
        list_vehicles create_vehicle check_vehicle_duplicate lookup_vehicle prefill_vehicle decode_vin
        delete_vehicle show_vehicle update_vehicle merge_vehicle transfer_vehicle
        list_vehicles_work_orders list_vendors show_vendor
        list_work_orders create_work_order delete_work_order show_work_order update_work_order
        show_work_order_declined_services show_work_order_service_history
        create_work_order_authorization update_work_order_authorization_decisions
        show_work_order_estimate show_work_order_inspection
        close_work_order close_work_order_as_paid decline_all_work_order_services
        reopen_work_order return_work_order_to_board save_work_order_for_later
        show_work_order_parts show_work_order_payments create_work_order_payment show_work_order_wip
      ]
      expected.each do |m|
        assert_respond_to client, m, "expected generated method ##{m}"
      end
    end

    def test_get_all_aliases_exist
      client = Client.new(token: "test", base_url: @base_url)
      %i[get_all_customers get_all_vehicles get_all_work_orders get_all_vendors get_all_service_categories].each do |m|
        assert_respond_to client, m, "expected generated method ##{m}"
      end
    end

    def test_create_customer_wraps_body
      stub_request(:post, "#{@base_url}/customers")
        .with(body: { customer: { first_name: "Jane" } }.to_json)
        .to_return(status: 201, body: { id: 1 }.to_json, headers: { "Content-Type" => "application/json" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.create_customer(customer: { first_name: "Jane" })
      assert_equal 1, result["id"]
    end

    def test_merge_customer_wraps_body
      stub_request(:post, "#{@base_url}/customers/1/merges")
        .with(body: { source_customer_id: 2 }.to_json)
        .to_return(status: 200, body: { id: 1 }.to_json, headers: { "Content-Type" => "application/json" })
      client = Client.new(token: "test", base_url: @base_url)
      result = client.merge_customer(1, source_customer_id: 2)
      assert_equal 1, result["id"]
    end
  end
end
