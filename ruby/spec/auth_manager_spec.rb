require_relative "spec_helper"
require "tmpdir"

class CredentialStoreProviderTest < Wenmar::TestCase
  def setup
    super
    @dir = Dir.mktmpdir("wenmar-auth")
    @store = Wenmar::CredentialStore.new(File.join(@dir, "credentials.json"))
  end

  def teardown
    FileUtils.remove_entry(@dir) if @dir && File.exist?(@dir)
  end

  def test_token_returns_valid_stored_token
    @store.save_token(Wenmar::Token.new(access_token: "valid", expires_at: Time.now + 3600))
    provider = Wenmar::CredentialStoreProvider.new(store: @store)
    assert_equal "valid", provider.token
  end

  def test_token_nil_when_nothing_stored
    provider = Wenmar::CredentialStoreProvider.new(store: @store)
    assert_nil provider.token
  end

  def test_token_refreshes_when_expired
    @store.save_token(Wenmar::Token.new(access_token: "old", refresh_token: "rt", expires_at: Time.now - 60))
    manager = Wenmar::AuthManager.new(store: @store)
    manager.set_refresh_fn do |refresh_token|
      raise "wrong token" unless refresh_token == "rt"

      Wenmar::Token.new(access_token: "new", refresh_token: "rt2", expires_at: Time.now + 3600)
    end
    provider = Wenmar::CredentialStoreProvider.new(store: @store, manager: manager)
    assert_equal "new", provider.token
    assert_equal "new", @store.get_token.access_token
  end

  def test_token_refreshes_within_window
    @store.save_token(Wenmar::Token.new(access_token: "old", refresh_token: "rt", expires_at: Time.now + 100))
    manager = Wenmar::AuthManager.new(store: @store)
    manager.set_refresh_fn { |_| Wenmar::Token.new(access_token: "new", refresh_token: "rt2", expires_at: Time.now + 3600) }
    provider = Wenmar::CredentialStoreProvider.new(store: @store, manager: manager, refresh_window: 300)
    assert_equal "new", provider.token
  end

  def test_token_does_not_refresh_when_far_from_expiry
    @store.save_token(Wenmar::Token.new(access_token: "fresh", refresh_token: "rt", expires_at: Time.now + 7200))
    manager = Wenmar::AuthManager.new(store: @store)
    manager.set_refresh_fn { |_| flunk "should not refresh" }
    provider = Wenmar::CredentialStoreProvider.new(store: @store, manager: manager, refresh_window: 300)
    assert_equal "fresh", provider.token
  end

  def test_credential_store_provider_refreshes_once_under_concurrency
    refresh_count = 0
    manager = Wenmar::AuthManager.new(store: @store)
    manager.set_refresh_fn do
      refresh_count += 1
      sleep 0.05
      Wenmar::Token.new(access_token: "fresh-#{refresh_count}", refresh_token: "r")
    end
    provider = Wenmar::CredentialStoreProvider.new(store: @store, manager: manager, refresh_window: 300)
    @store.save_token(Wenmar::Token.new(access_token: "old", refresh_token: "r",
                                        expires_at: Time.now + 1))

    threads = 8.times.map { Thread.new { provider.token } }
    threads.each(&:join)

    assert_equal 1, refresh_count, "expected a single refresh, got #{refresh_count}"
  end
end

class AuthManagerTest < Wenmar::TestCase
  def setup
    super
    @dir = Dir.mktmpdir("wenmar-manager")
    @store = Wenmar::CredentialStore.new(File.join(@dir, "credentials.json"))
  end

  def teardown
    FileUtils.remove_entry(@dir) if @dir && File.exist?(@dir)
  end

  def test_token_delegates_to_provider
    provider = Wenmar::StaticTokenProvider.new("static")
    manager = Wenmar::AuthManager.new(store: @store, provider: provider)
    assert_equal "static", manager.token
  end

  def test_token_reads_store_when_no_provider
    @store.save_token(Wenmar::Token.new(access_token: "stored"))
    manager = Wenmar::AuthManager.new(store: @store)
    assert_equal "stored", manager.token
  end

  def test_refresh_raises_when_no_refresh_token
    @store.save_token(Wenmar::Token.new(access_token: "a"))
    manager = Wenmar::AuthManager.new(store: @store)
    manager.set_refresh_fn { |_| Wenmar::Token.new(access_token: "b") }
    assert_raises(Wenmar::TokenError) { manager.refresh }
  end

  def test_refresh_uses_stored_refresh_token
    @store.save_token(Wenmar::Token.new(access_token: "a", refresh_token: "rt"))
    manager = Wenmar::AuthManager.new(store: @store)
    seen = nil
    manager.set_refresh_fn do |refresh_token|
      seen = refresh_token
      Wenmar::Token.new(access_token: "new", refresh_token: "rt2", expires_at: Time.now + 3600)
    end
    manager.refresh
    assert_equal "rt", seen
    assert_equal "new", @store.get_token.access_token
    assert_equal "rt2", @store.get_token.refresh_token
  end

  def test_logout_deletes_store
    @store.save_token(Wenmar::Token.new(access_token: "a"))
    manager = Wenmar::AuthManager.new(store: @store)
    manager.logout
    assert_nil @store.get_token
  end

  def test_oauth_default_refresh
    stub_request(:post, "https://api.example.com/oauth/token")
      .with(body: {grant_type: "refresh_token", client_id: "wenmar-cli", refresh_token: "rt"})
      .to_return(status: 200, body: {access_token: "refreshed", expires_in: 7200}.to_json)

    @store.save_token(Wenmar::Token.new(access_token: "a", refresh_token: "rt"))
    manager = Wenmar::AuthManager.new(store: @store, oauth: {base_url: "https://api.example.com", client_id: "wenmar-cli"})
    manager.refresh
    assert_equal "refreshed", @store.get_token.access_token
  end

  def test_refresh_raises_when_no_oauth_configured
    @store.save_token(Wenmar::Token.new(access_token: "a", refresh_token: "rt"))
    manager = Wenmar::AuthManager.new(store: @store)
    assert_raises(Wenmar::TokenError) { manager.refresh }
  end

  def test_oauth_defaults_client_id
    stub_request(:post, "https://api.example.com/oauth/token")
      .with(body: {grant_type: "refresh_token", client_id: "wenmar-cli", refresh_token: "rt"})
      .to_return(status: 200, body: {access_token: "refreshed", expires_in: 7200}.to_json)

    @store.save_token(Wenmar::Token.new(access_token: "a", refresh_token: "rt"))
    manager = Wenmar::AuthManager.new(store: @store, oauth: {base_url: "https://api.example.com"})
    manager.refresh
    assert_equal "refreshed", @store.get_token.access_token
  end
end
