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

    attr_reader :token
  end

  # Reads a Wenmar::Token from a store, auto-refreshing via an AuthManager when
  # the stored token is expired or within a refresh window of expiry.
  class CredentialStoreProvider
    DEFAULT_REFRESH_WINDOW = 300

    def initialize(store:, manager: nil, refresh_window: DEFAULT_REFRESH_WINDOW)
      @store = store
      @manager = manager
      @refresh_window = refresh_window
    end

    def token
      token = @store.get_token
      return nil if token.nil? || token.access_token.nil? || token.access_token.empty?

      if token.expired? || token.will_expire_within?(@refresh_window)
        if @manager
          @manager.refresh
          token = @store.get_token
        end
      end
      token&.access_token
    end
  end

  # Coordinates token storage, retrieval, and refresh. When configured with an
  # oauth hash (base_url and optional client_id), #refresh exchanges the stored
  # refresh token at the Doorkeeper token endpoint.
  class AuthManager
    def initialize(store:, provider: nil, oauth: nil)
      @store = store
      @provider = provider
      @refresh_fn = oauth ? build_oauth_refresh(oauth) : nil
    end

    # Overrides the refresh function used by #refresh. Useful for tests and
    # custom refresh flows.
    def set_refresh_fn(fn = nil, &block)
      @refresh_fn = fn || block
    end

    def token
      return @provider.token if @provider

      token = @store.get_token
      token&.access_token
    end

    def refresh
      token = @store.get_token
      raise TokenError, "no token stored" if token.nil?
      raise TokenError, "no refresh token stored" if token.refresh_token.nil? || token.refresh_token.empty?
      raise TokenError, "OAuth refresh is not configured" if @refresh_fn.nil?

      new_token = @refresh_fn.call(token.refresh_token)
      @store.save_token(new_token)
      new_token
    end

    def logout
      @store.delete
    end

    private

    def build_oauth_refresh(oauth)
      oauth = {base_url: oauth} if oauth.is_a?(String)
      base_url = oauth[:base_url]
      client_id = oauth[:client_id] || OAuth::DEFAULT_CLIENT_ID
      return nil if base_url.nil? || base_url.empty?

      ->(refresh_token) { OAuth.refresh(base_url: base_url, refresh_token: refresh_token, client_id: client_id) }
    end
  end
end
