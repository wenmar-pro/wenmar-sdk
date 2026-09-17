# frozen_string_literal: true

require "json"
require "fileutils"

module Wenmar
  class CredentialStore
    TOKEN_KEY = "access_token"
    LEGACY_TOKEN_KEY = "token"

    def initialize(path = default_path)
      @path = path
    end

    def load
      return {} unless File.exist?(@path)

      JSON.parse(File.read(@path))
    rescue JSON::ParserError
      {}
    end

    # Legacy: returns the access token string or nil. Prefer get_token.
    def token
      data = load
      data[TOKEN_KEY] || data[LEGACY_TOKEN_KEY]
    end

    # Legacy: wraps a token string in a Wenmar::Token and persists it.
    def save(token)
      save_token(Token.new(access_token: token))
    end

    def get_token
      data = load
      return nil if data.empty?

      return Token.new(access_token: data[LEGACY_TOKEN_KEY]) if data[LEGACY_TOKEN_KEY] && !data[TOKEN_KEY]

      Token.from_h(data)
    rescue TokenError
      nil
    end

    def save_token(token)
      token = Token.from_h(token) if token.is_a?(Hash)
      dir = File.dirname(@path)
      FileUtils.mkdir_p(dir)
      tmp = File.join(dir, ".#{File.basename(@path)}.tmp.#{Process.pid}.#{Thread.current.object_id}")
      begin
        File.write(tmp, JSON.pretty_generate(token.to_h), perm: 0o600)
        File.rename(tmp, @path)
      ensure
        File.delete(tmp) if File.exist?(tmp)
      end
    end

    def delete
      FileUtils.rm_f(@path)
    end

    private

    def default_path
      File.join(Dir.home, ".config", "wenmar", "credentials.json")
    end
  end

  class KeychainStore
    SERVICE = "wenmar"
    LEGACY_SERVICE = "wenmar-cli"
    ACCOUNT = "token"

    # The ruby-keychain gem is macOS-only and optional. It must not be a hard
    # runtime dependency, so we require it lazily and only when a keychain
    # operation is attempted. Keychain users must `gem install ruby-keychain`
    # themselves.
    def initialize
      require_keychain
    end

    def get_token
      password = read_password
      return nil unless password

      Token.from_h(JSON.parse(password))
    rescue TokenError, JSON::ParserError
      # A legacy plain-token entry has no JSON envelope; treat it as a bare
      # access token.
      bare = read_password
      bare ? Token.new(access_token: bare) : nil
    end

    def save_token(token)
      token = Token.from_h(token) if token.is_a?(Hash)
      write_password(JSON.generate(token.to_h))
    end

    def delete
      require_keychain
      Keychain.generic_passwords.where(service: SERVICE, account: ACCOUNT).first&.destroy
    end

    private

    def require_keychain
      require "keychain"
    rescue LoadError
      raise TokenError, "The 'ruby-keychain' gem is required to use KeychainStore. Run `gem install ruby-keychain`."
    end

    def read_password
      require_keychain
      item = Keychain.generic_passwords.where(service: SERVICE, account: ACCOUNT).first
      return item.password if item

      legacy = Keychain.generic_passwords.where(service: LEGACY_SERVICE, account: ACCOUNT).first
      return nil unless legacy

      # Migrate a legacy entry into the current service.
      write_password(legacy.password)
      legacy.password
    end

    def write_password(password)
      require_keychain
      Keychain.generic_passwords.where(service: SERVICE, account: ACCOUNT).first&.destroy
      Keychain.generic_passwords.create(service: SERVICE, account: ACCOUNT, password: password)
    end
  end
end
