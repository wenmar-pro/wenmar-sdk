# frozen_string_literal: true

module Wenmar
  class Error < StandardError
    attr_reader :code, :message, :field_errors, :status, :request_id

    def initialize(code:, message:, field_errors: {}, status: nil, request_id: nil)
      @code = code
      @message = message
      @field_errors = field_errors
      @status = status
      @request_id = request_id
      super("#{code}: #{message} (HTTP #{status})")
    end

    def self.from_response(response)
      body = response.body
      headers = response.respond_to?(:headers) ? (response.headers || {}) : {}

      # Empty bodies can't be parsed; fall back to the status code so callers
      # still get a meaningful error when the server omits the error envelope.
      if body.nil? || (body.is_a?(String) && body.strip.empty?)
        if (fallback = fallback_code_for_status(response.status))
          return new(code: fallback, message: "HTTP #{response.status}", status: response.status, request_id: headers["X-Request-Id"])
        end
        return new(code: "unknown", message: "HTTP #{response.status} with empty body", status: response.status)
      end

      body = JSON.parse(body) if body.is_a?(String)
      body = {} unless body.is_a?(Hash)

      if body["error"].is_a?(Hash)
        err = body["error"]
        new(
          code: err["code"] || "unknown",
          message: err["message"] || "API error",
          field_errors: field_errors_map(err["field_errors"] || err["details"] || {}),
          status: response.status,
          request_id: headers["X-Request-Id"]
        )
      elsif (fallback = fallback_code_for_status(response.status))
        new(code: fallback, message: "HTTP #{response.status}", status: response.status, request_id: headers["X-Request-Id"])
      else
        new(code: "unknown", message: body.to_s, status: response.status)
      end
    rescue JSON::ParserError
      new(code: "unknown", message: "Malformed error response", status: response.status)
    end

    def self.fallback_code_for_status(status)
      {
        401 => "unauthorized",
        403 => "forbidden",
        404 => "not_found",
        409 => "conflict",
        422 => "validation_failed",
        429 => "rate_limited",
        500 => "internal_error",
        502 => "bad_gateway",
        503 => "service_unavailable",
        504 => "gateway_timeout",
        507 => "limit_exceeded"
      }[status]
    end

    def self.field_errors_map(list)      case list
      when Hash
        list
      when Array
        list.each_with_object({}) do |item, memo|
          if item.is_a?(Hash)
            memo[item["field"] || item[:field]] = item["message"] || item[:message]
          end
        end
      else
        {}
      end
    end

    # Extracts validation field errors from the field_errors hash.
    # The Wenmar API sends: field_errors: { "first_name" => ["can't be blank"] }
    # Returns nil if there are no field errors.
    def field_errors_by_field
      return nil if @field_errors.nil? || @field_errors.empty?

      result = {}
      @field_errors.each do |field, raw|
        case raw
        when Array
          msgs = raw.map(&:to_s)
          result[field] = msgs if msgs.any?
        when String
          result[field] = [raw]
        end
      end
      result.empty? ? nil : result
    end

    # Reports whether the error is worth retrying.
    def retryable?
      return true if @code == "rate_limited"
      return false if @code == "limit_exceeded"

      @status >= 500 && @status != 507
    end
  end
end
