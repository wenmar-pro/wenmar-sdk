module Wenmar
  module TokenProvider
    def token
      raise NotImplementedError
    end
  end

  class StaticTokenProvider
    include TokenProvider
    attr_reader :token

    def initialize(token)
      raise ArgumentError, "token is required" if token.nil? || token.empty?
      @token = token
    end
  end
end
