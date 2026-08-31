# frozen_string_literal: true

module Wenmar
  class Config
    attr_accessor :base_url, :access_token, :token_provider, :location_id,
                  :timeout, :max_retries, :cache_enabled, :cache_store, :retry_options

    def initialize(attrs = {})
      @base_url = attrs[:base_url] || "https://app.wenmarpro.com"
      @access_token = attrs[:access_token]
      @token_provider = attrs[:token_provider]
      @location_id = attrs[:location_id]
      @timeout = attrs[:timeout] || 30
      @max_retries = attrs[:max_retries] || 3
      @cache_enabled = attrs.fetch(:cache_enabled, true)
      @cache_store = attrs[:cache_store] || {}
      @retry_options = attrs[:retry_options] || {
        max: @max_retries,
        interval: 0.1,
        interval_randomness: 0.5,
        backoff_factor: 2,
        max_interval: 30
      }
    end

    def self.from_env
      new(
        base_url: ENV.fetch("WENMAR_BASE_URL", "https://app.wenmarpro.com"),
        timeout: ENV["WENMAR_TIMEOUT"]&.to_i || 30,
        max_retries: ENV["WENMAR_MAX_RETRIES"]&.to_i || 3,
        cache_enabled: ENV["WENMAR_CACHE"] != "false"
      )
    end
  end
end
