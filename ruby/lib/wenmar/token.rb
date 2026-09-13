# frozen_string_literal: true

require "time"

module Wenmar
  class Token
    attr_reader :access_token, :refresh_token, :expires_at, :token_type

    def initialize(access_token:, refresh_token: nil, expires_at: nil, token_type: "Bearer")
      @access_token = access_token
      @refresh_token = refresh_token
      @expires_at = expires_at
      @token_type = token_type
    end

    def expired?
      @expires_at && @expires_at < Time.now
    end

    def will_expire_within?(seconds)
      @expires_at && Time.now + seconds > @expires_at
    end

    def to_h
      hash = { "access_token" => @access_token, "token_type" => @token_type }
      hash["refresh_token"] = @refresh_token if @refresh_token
      hash["expires_at"] = @expires_at.utc.iso8601 if @expires_at
      hash
    end

    def self.from_h(hash)
      hash ||= {}
      hash = hash.transform_keys(&:to_s)
      access_token = hash["access_token"]
      raise TokenError, "token hash missing access_token" if access_token.nil? || access_token.empty?

      expires_at = hash["expires_at"]
      expires_at = parse_time(expires_at) if expires_at.is_a?(String)
      new(
        access_token: access_token,
        refresh_token: hash["refresh_token"],
        expires_at: expires_at,
        token_type: hash.fetch("token_type", "Bearer")
      )
    end

    def self.parse_time(value)
      Time.parse(value)
    rescue ArgumentError
      raise TokenError, "invalid expires_at: #{value.inspect}"
    end
    private_class_method :parse_time
  end
end
