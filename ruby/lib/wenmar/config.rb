module Wenmar
  class Config
    attr_accessor :base_url, :timeout, :max_retries, :cache_enabled

    def initialize(base_url: "https://app.wenmarpro.com", timeout: 30, max_retries: 3, cache_enabled: true)
      @base_url = base_url
      @timeout = timeout
      @max_retries = max_retries
      @cache_enabled = cache_enabled
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
