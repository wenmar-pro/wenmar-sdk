# frozen_string_literal: true

require "minitest/autorun"
require_relative "generator_utils"

# Unit tests for the shared codegen helpers. These are the single source of
# truth for identifier transforms used by every generator and the parity gates,
# so a regression here silently renames the public SDK surface.
class GeneratorUtilsTest < Minitest::Test
  include GeneratorUtils

  def test_pascal_converts_snake_to_exported
    assert_equal "ListCustomers", pascal("list_customers")
    assert_equal "CreateCashEntry", pascal("create_cash_entry")
    assert_equal "Location", pascal("location")
    assert_equal "ShowWorkOrdersServicesLineItems", pascal("show_work_orders_services_line_items")
  end

  def test_pascal_ignores_empty_segments
    assert_equal "ABC", pascal("a_b__c")
  end

  def test_go_param_name
    assert_equal "id", go_param_name("id")
    assert_equal "customerId", go_param_name("customer_id")
  end

  def test_go_param_name_defaults_empty_to_id
    # An empty param name falls back to "id" (parts.shift || "id"). Defensive.
    assert_equal "id", go_param_name("")
  end

  def test_ruby_param_name_strips_brackets
    assert_equal "filters_has_open_work_order", ruby_param_name("filters[has_open_work_order]")
    assert_equal "status", ruby_param_name("status")
  end

  def test_normalize_tag
    assert_equal "work_orders", normalize_tag("Work Orders")
    assert_equal "cash_entries", normalize_tag("Cash Entries")
    assert_equal "account", normalize_tag("account")
  end

  def test_typed_list_requires_paginated_named_list_op
    assert typed_list?("id" => "list_customers", "paginated" => true, "responseSchema" => "Customer")
    refute typed_list?("id" => "list_customers", "paginated" => false, "responseSchema" => "Customer")
    refute typed_list?("id" => "show_customer", "paginated" => true, "responseSchema" => "Customer")
    refute typed_list?("id" => "list_customers", "paginated" => true, "responseSchema" => nil)
  end

  def test_test_aliases_target_known_manifest_ids
    # The aliases must point at ids that exist in the committed manifest;
    # otherwise the dispatch generator emits a lambda for a missing operation.
    require "json"
    root = File.expand_path("..", __dir__)
    manifest = JSON.parse(File.read(File.join(root, "spec", "operations.json")))
    ids = manifest.fetch("operations").map { |o| o.fetch("id") }
    TEST_ALIASES.each_value do |target|
      assert_includes ids, target, "TEST_ALIASES target #{target} not in operations.json"
    end
  end

  def test_path_param_type_returns_recorded_type
    op = { "id" => "show_customer", "method" => "get", "path" => "/customers/{id}",
           "pathParamTypes" => { "id" => "integer" } }
    assert_equal "integer", path_param_type(op, "id")
  end

  def test_path_param_type_fails_loudly_when_missing
    op = { "id" => "show_customer", "method" => "get", "path" => "/customers/{id}",
           "pathParamTypes" => {} }
    error = assert_raises(RuntimeError) { path_param_type(op, "id") }
    assert_includes error.message, "has no recorded type"
  end
end
