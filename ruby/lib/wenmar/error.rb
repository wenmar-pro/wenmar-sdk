require "json"

module Wenmar
  class Error < StandardError
    attr_reader :code, :message, :details, :status, :request_id

    def initialize(code:, message:, details:, status:, request_id: nil)
      @code = code
      @message = message
      @details = details
      @status = status
      @request_id = request_id
      super("#{code}: #{message} (HTTP #{status})")
    end

    def self.from_response(response)
      body = JSON.parse(response.body)
      error = body["error"] || {}
      headers = response.respond_to?(:headers) ? (response.headers || {}) : {}
      new(
        code: error["code"] || "unknown",
        message: error["message"] || "Unknown error",
        details: error["details"] || {},
        status: response.status,
        request_id: headers["X-Request-Id"]
      )
    rescue JSON::ParserError
      new(code: "unknown", message: "Malformed error response", details: {}, status: response.status)
    end

    # Extracts validation field errors from the details hash.
    # The Wenmar API sends: details: { "first_name" => ["can't be blank"] }
    # Returns nil if there are no field errors.
    def field_errors
      return nil if @details.nil? || @details.empty?

      result = {}
      @details.each do |field, raw|
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
