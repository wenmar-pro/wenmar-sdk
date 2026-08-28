require_relative "spec_helper"

class ConfigTest < Wenmar::TestCase
  def test_default_config
    config = Wenmar::Config.new
    assert_equal "https://app.wenmarpro.com", config.base_url
    assert_equal 30, config.timeout
    assert_equal 3, config.max_retries
    assert config.cache_enabled
  end

  def test_from_env
    ENV["WENMAR_BASE_URL"] = "https://staging.wenmarpro.com"
    ENV["WENMAR_TIMEOUT"] = "10"
    ENV["WENMAR_MAX_RETRIES"] = "5"
    ENV["WENMAR_CACHE"] = "false"
    config = Wenmar::Config.from_env
    assert_equal "https://staging.wenmarpro.com", config.base_url
    assert_equal 10, config.timeout
    assert_equal 5, config.max_retries
    refute config.cache_enabled
  ensure
    ENV.delete("WENMAR_BASE_URL")
    ENV.delete("WENMAR_TIMEOUT")
    ENV.delete("WENMAR_MAX_RETRIES")
    ENV.delete("WENMAR_CACHE")
  end
end

class AuthTest < Wenmar::TestCase
  def test_static_token_provider
    provider = Wenmar::StaticTokenProvider.new("my-token")
    assert_equal "my-token", provider.token
  end

  def test_static_token_provider_empty
    assert_raises(ArgumentError) { Wenmar::StaticTokenProvider.new("") }
  end
end
