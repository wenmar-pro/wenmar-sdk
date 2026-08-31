# frozen_string_literal: true

module Wenmar
  class TokenError < StandardError; end

  module TokenProvider
    def token
      raise NotImplementedError
    end
  end

  class StaticTokenProvider
    include TokenProvider

    def initialize(token)
      raise ArgumentError, "token is required" if token.nil? || token.empty?

      @token = token
    end

    def token
      @token
    end
  end
end
