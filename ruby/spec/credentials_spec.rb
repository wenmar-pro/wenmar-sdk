require_relative "spec_helper"
require "tmpdir"

class CredentialStoreTest < Wenmar::TestCase
  def setup
    super
    @dir = Dir.mktmpdir("wenmar-credentials")
    @path = File.join(@dir, "credentials.json")
    @store = Wenmar::CredentialStore.new(@path)
  end

  def teardown
    FileUtils.remove_entry(@dir) if @dir && File.exist?(@dir)
  end

  def test_save_token_round_trip
    token = Wenmar::Token.new(access_token: "a", refresh_token: "r", expires_at: Time.now + 7200)
    @store.save_token(token)

    stored = @store.get_token
    assert_equal "a", stored.access_token
    assert_equal "r", stored.refresh_token
    assert_equal "Bearer", stored.token_type
    assert_in_delta token.expires_at.to_f, stored.expires_at.to_f, 1
  end

  def test_save_token_writes_0600
    @store.save_token(Wenmar::Token.new(access_token: "a"))
    assert_equal 0o600, File.stat(@path).mode & 0o777
  end

  def test_save_token_leaves_no_tmp_file
    @store.save_token(Wenmar::Token.new(access_token: "abc"))
    leftover = Dir.glob(File.join(File.dirname(@path), "*.tmp.*"), File::FNM_DOTMATCH)
    refute leftover.any?, "temporary credential files must be cleaned up: #{leftover.inspect}"
  end

  def test_save_token_accepts_hash
    @store.save_token({"access_token" => "a", "refresh_token" => "r"})
    token = @store.get_token
    assert_equal "a", token.access_token
    assert_equal "r", token.refresh_token
  end

  def test_get_token_nil_when_missing
    assert_nil @store.get_token
  end

  def test_migrates_legacy_access_token_format
    File.write(@path, JSON.generate({"access_token" => "legacy"}))
    token = @store.get_token
    assert_equal "legacy", token.access_token
    assert_nil token.refresh_token
  end

  def test_migrates_legacy_token_format
    File.write(@path, JSON.generate({"token" => "legacy"}))
    token = @store.get_token
    assert_equal "legacy", token.access_token
  end

  def test_delete_removes_file
    @store.save_token(Wenmar::Token.new(access_token: "a"))
    @store.delete
    refute File.exist?(@path)
    assert_nil @store.get_token
  end

  def test_delete_idempotent_when_absent
    @store.delete
    @store.delete
    refute File.exist?(@path)
  end

  def test_save_shim_wraps_in_token
    @store.save("legacy")
    token = @store.get_token
    assert_equal "legacy", token.access_token
  end

  def test_token_shim_returns_access_token_string
    @store.save_token(Wenmar::Token.new(access_token: "abc", refresh_token: "r"))
    assert_equal "abc", @store.token
  end

  def test_load_returns_hash
    @store.save_token(Wenmar::Token.new(access_token: "a", refresh_token: "r"))
    data = @store.load
    assert_equal "a", data["access_token"]
    assert_equal "r", data["refresh_token"]
  end
end

class KeychainStoreTest < Wenmar::TestCase
  def test_constructing_without_keychain_gem_raises_clear_error
    original = defined?(::Keychain) ? Object.send(:remove_const, :Keychain) : nil

    begin
      require "keychain"
    rescue LoadError
      # gem absent
    end

    begin
      Wenmar::KeychainStore.new
      flunk "expected TokenError when keychain gem is unavailable"
    rescue Wenmar::TokenError => e
      assert_includes e.message, "ruby-keychain"
    ensure
      Object.const_set(:Keychain, original) if original
    end
  end
end
