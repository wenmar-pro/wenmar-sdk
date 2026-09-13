# frozen_string_literal: true

require "faraday"
require "json"
require "uri"

module Wenmar
  module OAuth
    DEFAULT_CLIENT_ID = "wenmar-cli"

    # POSTs a refresh_token grant to the Doorkeeper token endpoint and returns
    # a fresh Wenmar::Token. Redirects are not followed: a 3xx from the token
    # endpoint is a security risk, not a redirect to follow.
    def self.refresh(base_url:, refresh_token:, client_id: DEFAULT_CLIENT_ID)
      endpoint = "#{base_url.to_s.sub(%r{/+\z}, "")}/oauth/token"

      conn = Faraday.new(url: endpoint) do |f|
        f.options.timeout = 10
        f.adapter Faraday.default_adapter
      end

      response = conn.post do |req|
        req.headers["Content-Type"] = "application/x-www-form-urlencoded"
        req.headers["Accept"] = "application/json"
        req.body = URI.encode_www_form({ grant_type: "refresh_token", client_id: client_id, refresh_token: refresh_token })
      end

      unless response.success?
        error = begin
          JSON.parse(response.body)["error"]
        rescue StandardError
          nil
        end
        message = error ? "token refresh failed: #{error}" : "token refresh failed: HTTP #{response.status}"
        raise TokenError, message
      end

      data = JSON.parse(response.body)
      if data["access_token"].nil? || data["access_token"].empty?
        raise TokenError, "refresh response missing access_token"
      end

      Token.new(
        access_token: data["access_token"],
        refresh_token: data["refresh_token"],
        expires_at: data["expires_in"] ? Time.now + data["expires_in"] : nil,
        token_type: data["token_type"] || "Bearer"
      )
    rescue JSON::ParserError => e
      raise TokenError, "parse refresh response: #{e.message}"
    end
  end
end
