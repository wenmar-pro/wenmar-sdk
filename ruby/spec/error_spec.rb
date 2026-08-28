require_relative "spec_helper"

module Wenmar
  class ErrorTest < TestCase
    def test_field_errors_extracts_string_arrays
      error = Error.new(
        code: "validation_failed",
        message: "Validation failed",
        details: { "first_name" => ["can't be blank"], "email" => ["is invalid", "already taken"] },
        status: 422
      )
      fe = error.field_errors
      assert_equal ["can't be blank"], fe["first_name"]
      assert_equal ["is invalid", "already taken"], fe["email"]
    end

    def test_field_errors_handles_single_string
      error = Error.new(
        code: "validation_failed",
        message: "x",
        details: { "last_name" => "can't be blank" },
        status: 422
      )
      assert_equal ["can't be blank"], error.field_errors["last_name"]
    end

    def test_field_errors_returns_nil_when_empty
      error = Error.new(code: "not_found", message: "x", details: {}, status: 404)
      assert_nil error.field_errors
    end

    def test_retryable_rate_limited
      error = Error.new(code: "rate_limited", message: "x", details: {}, status: 429)
      assert error.retryable?
    end

    def test_retryable_5xx
      error = Error.new(code: "internal_error", message: "x", details: {}, status: 500)
      assert error.retryable?
    end

    def test_not_retryable_validation
      error = Error.new(code: "validation_failed", message: "x", details: {}, status: 422)
      refute error.retryable?
    end

    def test_not_retryable_not_found
      error = Error.new(code: "not_found", message: "x", details: {}, status: 404)
      refute error.retryable?
    end

    def test_not_retryable_limit_exceeded
      error = Error.new(code: "limit_exceeded", message: "x", details: {}, status: 507)
      refute error.retryable?
    end

    def test_from_response_507
      response = Faraday::Response.new(
        status: 507,
        body: { error: { code: "limit_exceeded", message: "Account limit reached", details: {} } }.to_json
      )
      error = Error.from_response(response)
      assert_equal "limit_exceeded", error.code
      assert_equal 507, error.status
      refute error.retryable?
    end
  end
end
