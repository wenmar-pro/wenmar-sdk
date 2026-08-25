require "json"

module Wenmar
  class Error < StandardError
    attr_reader :code, :message, :details, :status

    def initialize(code:, message:, details:, status:)
      @code = code
      @message = message
      @details = details
      @status = status
      super("#{code}: #{message} (HTTP #{status})")
    end

    def self.from_response(response)
      body = JSON.parse(response.body)
      error = body["error"] || {}
      new(
        code: error["code"] || "unknown",
        message: error["message"] || "Unknown error",
        details: error["details"] || {},
        status: response.status
      )
    rescue JSON::ParserError
      new(code: "unknown", message: "Malformed error response", details: {}, status: response.status)
    end
  end
end
