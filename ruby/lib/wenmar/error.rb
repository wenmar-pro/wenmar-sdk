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
      body = JSON.parse(body) if body.is_a?(String)
      body = {} unless body.is_a?(Hash)
      headers = response.respond_to?(:headers) ? (response.headers || {}) : {}

      if body["error"].is_a?(Hash)
        err = body["error"]
        new(
          code: err["code"] || "unknown",
          message: err["message"] || "API error",
          field_errors: field_errors_map(err["field_errors"] || err["details"] || {}),
          status: response.status,
          request_id: headers["X-Request-Id"]
        )
      elsif response.status == 507
        new(code: "limit_exceeded", message: "storage limit exceeded", status: response.status)
      else
        new(code: "unknown", message: body.to_s, status: response.status)
      end
    rescue JSON::ParserError
      new(code: "unknown", message: "Malformed error response", status: response.status)
    end

    def self.field_errors_map(list)
      case list
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
