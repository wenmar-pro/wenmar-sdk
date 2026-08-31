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

    def token
      data = load
      data[TOKEN_KEY] || data[LEGACY_TOKEN_KEY]
    end

    def save(token)
      FileUtils.mkdir_p(File.dirname(@path))
      File.write(@path, JSON.pretty_generate({ TOKEN_KEY => token }), perm: 0600)
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

    def initialize
      require "keychain"
    end

    def token
      item = Keychain.generic_passwords.where(service: SERVICE, account: ACCOUNT).first
      return item.password if item

      legacy = Keychain.generic_passwords.where(service: LEGACY_SERVICE, account: ACCOUNT).first
      return nil unless legacy

      save(legacy.password)
      legacy.password
    end

    def save(token)
      existing = Keychain.generic_passwords.where(service: SERVICE, account: ACCOUNT).first
      existing.destroy if existing
      Keychain.generic_passwords.create(service: SERVICE, account: ACCOUNT, password: token)
    end
  end
end
