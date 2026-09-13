# frozen_string_literal: true

# Shared pure helpers + constants for the codegen pipeline. This is the single
# source of truth for identifier transforms and the conformance TEST_ALIASES so
# the generator scripts and parity gates never drift apart by hand.
#
# The module is required via require_relative, which resolves relative to this
# file regardless of the caller's CWD, so scripts work whether invoked from the
# repo root (`ruby scripts/foo.rb`) or from inside scripts/.
module GeneratorUtils
  # Test operations that aren't plain manifest IDs but reference a manifest
  # operation. Value is the manifest operation id they exercise. Single source
  # of truth — shared by generate_conformance_dispatch.rb (which emits dispatch
  # entries for them) and check_conformance_parity.rb (which validates them).
  TEST_ALIASES = {
    "list_customer_vehicles" => "list_customers_vehicles",
    "list_customer_work_orders" => "list_customers_work_orders",
    "list_vehicle_work_orders" => "list_vehicles_work_orders",
    "list_customers_with_params" => "list_customers",
    "list_customers_with_params_paginated" => "list_customers",
    "list_customers_paginated" => "list_customers",
    "list_work_orders_paginated" => "list_work_orders",
    "check_duplicate" => "check_vehicle_duplicate"
  }.freeze

  # Map an operation id (snake_case) to a Go exported method/type name.
  def pascal(identifier)
    identifier.split("_").reject(&:empty?).map { |p| p[0].upcase + p[1..] }.join
  end

  # Convert a snake_case path/query param name to a camelCase Go identifier.
  def go_param_name(name)
    parts = name.split("_")
    (parts.shift || "id") + parts.map { |p| p[0].upcase + p[1..] }.join
  end

  # Convert a query param name (e.g. "filters[has_open_work_order]") into a
  # valid Ruby keyword identifier (e.g. "filters_has_open_work_order").
  def ruby_param_name(name)
    name.gsub(/[\[\]]/, "_").gsub(/_+/, "_").sub(/_+\z/, "")
  end

  # Whether an operation gets a typed ListResult wrapper (and thus its raw
  # method is renamed to ListXxxRaw). Mirrored by generate_go_wrapper.rb and
  # generate_conformance_dispatch.rb.
  def typed_list?(op)
    op["paginated"] && op["responseSchema"] && op["id"].start_with?("list_")
  end

  # Normalize an OpenAPI tag to a filename/section key
  # (e.g. "Work Orders" -> "work_orders").
  def normalize_tag(tag)
    tag.downcase.gsub(/[^a-z0-9]+/, "_").gsub(/_+/, "_").sub(/_+$/, "")
  end

  extend self
end
