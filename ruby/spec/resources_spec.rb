require "json"
require_relative "spec_helper"

module Wenmar
  class ResourcesTest < TestCase
    MANIFEST_PATH = File.expand_path("../../spec/operations.json", __dir__)
    GENERATED_OPERATION_IDS = JSON.parse(File.read(MANIFEST_PATH)).fetch("operations").map { |op| op.fetch("id").to_sym }.freeze

    def test_all_generated_methods_exist
      client = Client.new(token: "test", base_url: @base_url)
      GENERATED_OPERATION_IDS.each do |m|
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
