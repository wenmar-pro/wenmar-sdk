# frozen_string_literal: true

require "minitest/autorun"
require "yaml"
require "json"
require "fileutils"
require_relative "generate_docs"

class GenerateDocsTest < Minitest::Test
  def setup
    @tmpdir = File.join(Dir.tmpdir, "generate_docs_test_#{Process.pid}_#{rand(1000)}")
    FileUtils.mkdir_p(@tmpdir)
  end

  def teardown
    FileUtils.rm_rf(@tmpdir)
  end

  # Operation extraction

  def test_extracts_operations_grouped_by_tag
    result = GenerateDocs.document(spec)

    assert_equal %w[account customers vehicles], result.keys.sort
    assert_equal 1, result["account"].length
    assert_equal 5, result["customers"].length
    assert_equal 6, result["vehicles"].length # list, create, show, update, delete, vin_decode
  end

  def test_operations_are_sorted_by_path_then_method_within_tag
    result = GenerateDocs.document(spec)
    customers = result["customers"]

    assert_equal ["/customers", "/customers", "/customers/{id}", "/customers/{id}", "/customers/{id}"],
      customers.map { |op| op[:path] }
    assert_equal %w[GET POST DELETE GET PATCH], customers.map { |op| op[:method] }
  end

  def test_extracts_operation_metadata
    result = GenerateDocs.document(spec)
    op = result["customers"].find { |o| o[:operation_id] == "list_customers" }

    assert_equal "list_customers", op[:operation_id]
    assert_equal "GET", op[:method]
    assert_equal "/customers", op[:path]
    assert_equal "List all customers, paginated via the Link header.", op[:description].strip
  end

  def test_falls_back_to_path_segment_when_tags_missing
    input = spec_with("/account" => { "get" => { "summary" => "show", "responses" => {} } })
    result = GenerateDocs.document(input)

    assert_includes result.keys, "account"
    assert_equal 1, result["account"].length
  end

  def test_ignores_non_http_keys_on_path_items
    input = spec_with("/customers" => {
      "parameters" => [{ "name" => "page", "in" => "query", "schema" => { "type" => "integer" } }],
      "get" => { "summary" => "index", "responses" => {} }
    })
    result = GenerateDocs.document(input)

    assert_equal 1, result["customers"].length
  end

  # Heading rendering

  def test_titleizes_operation_id_for_heading
    assert_equal "List customers", GenerateDocs.titleize_operation_id("list_customers")
    assert_equal "Show customer", GenerateDocs.titleize_operation_id("show_customer")
    assert_equal "Create work order", GenerateDocs.titleize_operation_id("create_work_order")
    assert_equal "Decode VIN", GenerateDocs.titleize_operation_id("decode_vin")
    assert_equal "Show work order WIP", GenerateDocs.titleize_operation_id("show_work_order_wip")
  end

  # Parameters

  def test_renders_parameter_table
    op = build_op(
      "get" => {
        "summary" => "check_duplicate",
        "parameters" => [
          { "name" => "email", "in" => "query", "required" => false, "schema" => { "type" => "string" } },
          { "name" => "phone", "in" => "query", "required" => true, "schema" => { "type" => "integer" } }
        ],
        "responses" => {}
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "| Param | Type | Required |"
    assert_includes markdown, "| `email` | string | No |"
    assert_includes markdown, "| `phone` | integer | Yes |"
  end

  def test_renders_no_parameter_table_when_absent
    op = build_op("get" => { "summary" => "show", "responses" => {} })

    markdown = GenerateDocs.render_operation(op)
    refute_includes markdown, "| Param |"
  end

  # Request body

  def test_renders_request_body_with_wrapper_key
    op = build_op(
      "post" => {
        "summary" => "create",
        "requestBody" => {
          "content" => {
            "application/json" => {
              "schema" => {
                "type" => "object",
                "properties" => {
                  "customer" => {
                    "type" => "object",
                    "properties" => {
                      "first_name" => { "type" => "string" },
                      "last_name" => { "type" => "string" }
                    }
                  }
                }
              }
            }
          }
        },
        "responses" => {}
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "**Request body**"
    assert_includes markdown, "wrapper key `customer`"
    assert_includes markdown, "| `first_name` | string | No |"
    assert_includes markdown, "| `last_name` | string | No |"
  end

  def test_renders_request_body_without_wrapper
    op = build_op(
      "post" => {
        "summary" => "seed_defaults",
        "requestBody" => {
          "content" => {
            "application/json" => {
              "schema" => {
                "type" => "object",
                "properties" => {
                  "tag_ids" => { "type" => "array", "items" => { "type" => "integer" } }
                }
              }
            }
          }
        },
        "responses" => {}
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "**Request body**"
    refute_includes markdown, "wrapper key"
    assert_includes markdown, "| `tag_ids` | array of integer | No |"
  end

  # Responses

  def test_renders_ref_response_with_schema_link
    op = build_op(
      "get" => {
        "summary" => "show",
        "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Customer" }) }
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "**Response 200**"
    assert_includes markdown, "[Customer](#customer-schema)"
  end

  def test_renders_array_of_ref_response
    op = build_op(
      "get" => {
        "summary" => "index",
        "responses" => {
          "200" => response_schema("type" => "array", "items" => { "$ref" => "#/components/schemas/Customer" })
        }
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "**Response 200**"
    assert_includes markdown, "array of [Customer](#customer-schema)"
  end

  def test_renders_inline_response_schema_table
    op = build_op(
      "post" => {
        "summary" => "seed_defaults",
        "responses" => {
          "200" => response_schema(
            "type" => "object",
            "properties" => {
              "created" => { "type" => "integer" },
              "message" => { "type" => "string" }
            }
          )
        }
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "**Response 200**"
    assert_includes markdown, "| Field | Type | Required |"
    assert_includes markdown, "| `created` | integer | No |"
  end

  def test_renders_no_content_response
    op = build_op(
      "delete" => {
        "summary" => "destroy",
        "responses" => { "204" => { "description" => "no content" } }
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "**Response 204**"
  end

  # cURL examples

  def test_renders_curl_example_from_spec
    op = build_op(
      "get" => {
        "summary" => "show",
        "x-curl-example" => 'curl -H "Authorization: Bearer $TOKEN" https://app.wenmarpro.com/customers.json',
        "responses" => {}
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "```bash"
    assert_includes markdown, 'curl -H "Authorization: Bearer $TOKEN" https://app.wenmarpro.com/customers.json'
  end

  def test_renders_curl_example_exactly_once
    op = build_op(
      "get" => {
        "summary" => "show",
        "x-curl-example" => "curl -H 'X: 1' https://example.com",
        "responses" => {}
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_equal 1, markdown.scan("curl -H 'X: 1'").length
  end

  def test_renders_error_ref_responses_as_detail_line_only
    op = build_op(
      "get" => {
        "summary" => "show",
        "responses" => {
          "404" => response_schema(
            "type" => "object",
            "properties" => { "error" => { "$ref" => "#/components/schemas/Error" } },
            "required" => ["error"]
          )
        }
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "**Response 404** — [Error](#error-schema)"
    refute_includes markdown, "| Field | Type | Required |\n|---|---|---|\n| `error` |"
  end

  def test_no_blank_line_runs_around_responses_and_curl
    op = build_op(
      "get" => {
        "summary" => "show",
        "x-curl-example" => "curl https://example.com",
        "responses" => {
          "200" => response_schema({ "$ref" => "#/components/schemas/Customer" }),
          "404" => response_schema(
            "type" => "object",
            "properties" => { "error" => { "$ref" => "#/components/schemas/Error" } }
          )
        }
      }
    )

    markdown = GenerateDocs.render_operation(op)
    refute_includes markdown, "\n\n\n", "expected no runs of 3+ newlines"
  end

  def test_blank_line_between_response_status_and_table
    op = build_op(
      "get" => {
        "summary" => "show",
        "responses" => {
          "200" => response_schema("type" => "object", "properties" => { "id" => { "type" => "integer" } })
        }
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "**Response 200**\n\n| Field |"
  end

  def test_description_falls_back_to_titleized_summary
    op = build_op(
      "get" => {
        "summary" => "check_duplicate",
        "responses" => {}
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "Check duplicate"
  end

  # Fixture examples

  def test_renders_fixture_example_when_available
    op = build_op(
      "get" => {
        "summary" => "show",
        "operationId" => "show_customer",
        "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Customer" }) }
      }
    )
    fixtures = { "show_customer" => { "id" => 1, "full_name" => "Jane Doe" } }

    markdown = GenerateDocs.render_operation(op, fixtures: fixtures)
    assert_includes markdown, "**Example**"
    assert_includes markdown, '"full_name": "Jane Doe"'
  end

  def test_omits_example_block_when_no_fixture
    op = build_op(
      "get" => {
        "summary" => "show",
        "operationId" => "show_customer",
        "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Customer" }) }
      }
    )

    markdown = GenerateDocs.render_operation(op, fixtures: {})
    refute_includes markdown, "**Example**"
  end

  def test_loads_fixtures_from_manifest_and_dir
    fixtures_dir = File.join(@tmpdir, "fixtures")
    FileUtils.mkdir_p(File.join(fixtures_dir, "customers"))
    File.write(File.join(fixtures_dir, "manifest.yaml"),
               YAML.dump({ "targets" => [{ "id" => "cust-get", "operation" => "show_customer", "fixture" => "customers/get.json" }] }))
    File.write(File.join(fixtures_dir, "customers", "get.json"), JSON.dump({ "id" => 7, "full_name" => "Jane" }))

    fixtures = GenerateDocs.load_fixtures(fixtures_dir)
    assert_equal({ "id" => 7, "full_name" => "Jane" }, fixtures["show_customer"])
  end

  def test_write_sections_embeds_fixtures
    fixtures = { "show_customer" => { "id" => 1, "full_name" => "Jane Doe" } }
    GenerateDocs.write_sections(@tmpdir, spec, fixtures: fixtures)

    content = File.read(File.join(@tmpdir, "sections", "customers.md"))
    assert_includes content, '"full_name": "Jane Doe"'
  end

  # Heading and path block

  def test_renders_heading_and_method_path_block
    op = build_op(
      "get" => {
        "summary" => "show",
        "operationId" => "show_customer",
        "responses" => {}
      }
    )

    markdown = GenerateDocs.render_operation(op)
    assert_includes markdown, "## Show customer\n"
    assert_includes markdown, "```\nGET /things\n```"
  end

  # Schema tables

  def test_renders_schema_table_with_required_marking
    schema = {
      "type" => "object",
      "properties" => {
        "id" => { "type" => "integer" },
        "full_name" => { "type" => "string" },
        "company_name" => { "type" => "string", "nullable" => true },
        "vehicles" => {
          "type" => "array",
          "items" => { "type" => "object", "properties" => { "id" => { "type" => "integer" } } }
        }
      },
      "required" => %w[id full_name]
    }

    markdown = GenerateDocs.render_schema_table(schema)
    assert_includes markdown, "`id` | integer | Yes |"
    assert_includes markdown, "`full_name` | string | Yes |"
    assert_includes markdown, "`company_name` | string \\| null | No |"
    assert_includes markdown, "`vehicles` | array of object | No |"
  end

  def test_renders_nested_object_as_sub_table
    schema = {
      "type" => "object",
      "properties" => {
        "totals" => {
          "type" => "object",
          "properties" => {
            "total_cents" => { "type" => "integer" },
            "currency" => { "type" => "string" }
          }
        }
      }
    }

    markdown = GenerateDocs.render_schema_table(schema)
    assert_includes markdown, "`totals` — object:"
    assert_includes markdown, "`total_cents` | integer | No |"
  end

  def test_renders_ref_field_as_schema_link
    schema = {
      "type" => "object",
      "properties" => { "customer" => { "$ref" => "#/components/schemas/Customer" } }
    }

    markdown = GenerateDocs.render_schema_table(schema)
    assert_includes markdown, "[Customer](#customer-schema)"
  end

  # api-reference.md endpoint table

  def test_builds_endpoint_rows_for_all_operations
    rows = GenerateDocs.endpoint_rows(spec)
    assert_equal 12, rows.length

    account_row = rows.find { |r| r[:operation_id] == "list_account" }
    refute_nil account_row
    assert_equal "GET", account_row[:method]
    assert_equal "/account", account_row[:path]
  end

  # File assembly

  def test_writes_section_files_grouped_by_tag
    GenerateDocs.write_sections(@tmpdir, spec)

    assert File.exist?(File.join(@tmpdir, "sections", "customers.md"))
    assert File.exist?(File.join(@tmpdir, "sections", "vehicles.md"))
    assert File.exist?(File.join(@tmpdir, "sections", "account.md"))

    content = File.read(File.join(@tmpdir, "sections", "customers.md"))
    assert_includes content, "# Customers"
    assert_includes content, "AUTO-GENERATED"
    assert_includes content, "## List customers"
    assert_includes content, "## Show customer"
    assert_includes content, "### Customer schema"
    refute_includes content, "\n\n\n", "section files must not contain 3+ newline runs"
  end

  def test_writes_api_reference_file
    GenerateDocs.write_api_reference(@tmpdir, spec)

    content = File.read(File.join(@tmpdir, "api-reference.md"))
    assert_includes content, "| GET | `/account` | `list_account` |"
    assert_includes content, "| GET | `/customers` | `list_customers` |"
    assert_includes content, "AUTO-GENERATED"
  end

  # llm-compact.md — single-file view for small context windows

  def test_writes_llm_compact_file
    GenerateDocs.write_llm_compact(@tmpdir, spec)

    content = File.read(File.join(@tmpdir, "llm-compact.md"))
    assert_includes content, "AUTO-GENERATED"
    # Every operation is present as a compact heading
    assert_includes content, "GET /customers"
    assert_includes content, "POST /customers"
    assert_includes content, "DELETE /customers/{id}"
    assert_includes content, "GET /vehicles/vin_decode"
  end

  def test_llm_compact_omits_curl_and_examples
    GenerateDocs.write_llm_compact(@tmpdir, spec)

    content = File.read(File.join(@tmpdir, "llm-compact.md"))
    refute_includes content, "```bash"
    refute_includes content, "**Example**"
  end

  def test_llm_compact_includes_response_shape_summary
    GenerateDocs.write_llm_compact(@tmpdir, spec)

    content = File.read(File.join(@tmpdir, "llm-compact.md"))
    # Ref responses show the schema name with status
    assert_includes content, "200: Customer"
    # Array of ref
    assert_includes content, "200: array of Customer"
    # No content
    assert_includes content, "-> no content"
  end

  def test_llm_compact_includes_param_names
    GenerateDocs.write_llm_compact(@tmpdir, spec)

    content = File.read(File.join(@tmpdir, "llm-compact.md"))
    assert_includes content, "vin"
  end

  private

  def spec
    input = spec_with(
      "/account" => {
        "get" => {
          "summary" => "show", "tags" => ["account"], "operationId" => "list_account",
          "responses" => { "200" => response_schema("type" => "object", "properties" => { "id" => { "type" => "integer" }, "name" => { "type" => "string" } }) }
        }
      },
      "/customers" => {
        "get" => {
          "summary" => "index", "tags" => ["customers"], "operationId" => "list_customers",
          "description" => "List all customers, paginated via the Link header.",
          "responses" => { "200" => response_schema("type" => "array", "items" => { "$ref" => "#/components/schemas/Customer" }) }
        },
        "post" => {
          "summary" => "create", "tags" => ["customers"], "operationId" => "create_customer",
          "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Customer" }) }
        }
      },
      "/customers/{id}" => {
        "get" => {
          "summary" => "show", "tags" => ["customers"], "operationId" => "show_customer",
          "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Customer" }) }
        },
        "patch" => {
          "summary" => "update", "tags" => ["customers"], "operationId" => "update_customer",
          "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Customer" }) }
        },
        "delete" => {
          "summary" => "destroy", "tags" => ["customers"], "operationId" => "delete_customer",
          "responses" => { "204" => { "description" => "no content" } }
        }
      },
      "/vehicles" => {
        "get" => {
          "summary" => "index", "tags" => ["vehicles"], "operationId" => "list_vehicles",
          "responses" => { "200" => response_schema("type" => "array", "items" => { "$ref" => "#/components/schemas/Vehicle" }) }
        },
        "post" => {
          "summary" => "create", "tags" => ["vehicles"], "operationId" => "create_vehicle",
          "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Vehicle" }) }
        }
      },
      "/vehicles/{id}" => {
        "get" => {
          "summary" => "show", "tags" => ["vehicles"], "operationId" => "show_vehicle",
          "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Vehicle" }) }
        },
        "patch" => {
          "summary" => "update", "tags" => ["vehicles"], "operationId" => "update_vehicle",
          "responses" => { "200" => response_schema({ "$ref" => "#/components/schemas/Vehicle" }) }
        },
        "delete" => {
          "summary" => "destroy", "tags" => ["vehicles"], "operationId" => "delete_vehicle",
          "responses" => { "204" => { "description" => "no content" } }
        }
      },
      "/vehicles/vin_decode" => {
        "get" => {
          "summary" => "vin_decode", "tags" => ["vehicles"], "operationId" => "decode_vin",
          "parameters" => [{ "name" => "vin", "in" => "query", "required" => true, "schema" => { "type" => "string" } }],
          "responses" => { "200" => response_schema("type" => "object", "properties" => { "make" => { "type" => "string" }, "model" => { "type" => "string" } }) }
        }
      }
    )

    input["components"] = {
      "schemas" => {
        "Customer" => {
          "type" => "object",
          "properties" => {
            "id" => { "type" => "integer" },
            "full_name" => { "type" => "string" },
            "emails" => { "type" => "array", "items" => { "type" => "object", "properties" => { "address" => { "type" => "string" } } } }
          },
          "required" => %w[id full_name]
        },
        "Vehicle" => {
          "type" => "object",
          "properties" => { "id" => { "type" => "integer" }, "make" => { "type" => "string" } },
          "required" => ["id"]
        }
      }
    }

    input
  end

  def spec_with(paths)
    {
      "openapi" => "3.0.0",
      "info" => { "title" => "Wenmar Pro API", "version" => "1.0.0" },
      "paths" => paths
    }
  end

  # Builds an operation by putting it through the real extraction pipeline,
  # so tests exercise extraction + rendering together.
  def build_op(methods)
    path = methods.delete("path") || "/things"
    input = spec_with(path => methods)
    GenerateDocs.document(input).values.flatten.first
  end

  def response_schema(schema)
    { "description" => "response", "content" => { "application/json" => { "schema" => schema } } }
  end
end