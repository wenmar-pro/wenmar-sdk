require "json"
require "fileutils"

module Wenmar
  class CredentialStore
    include TokenProvider

    KEYRING_SERVICE = "wenmar-sdk"

    def initialize(path: default_path)
      @path = path
    end

    def save(token)
      save_to_keyring(token)
      save_to_file(token)
    end

    def token
      from_keyring || from_file
    end

    private

    def default_path
      File.join(Dir.home, ".config", "wenmar", "credentials.json")
    end

    def save_to_keyring(token)
      require "keychain"
      keychain = Keychain.default
      item = keychain.generic_password_item(
        service: KEYRING_SERVICE,
        account: "default",
        password: token
      )
      item.save!
    rescue LoadError, StandardError
      # keychain gem not installed or no keyring available — file fallback
    end

    def save_to_file(token)
      FileUtils.mkdir_p(File.dirname(@path))
      File.write(@path, { token: token }.to_json, perm: 0600)
    end

    def from_keyring
      require "keychain"
      item = Keychain.default.generic_password_item(service: KEYRING_SERVICE, account: "default")
      item.password if item
    rescue LoadError, StandardError
      nil
    end

    def from_file
      data = JSON.parse(File.read(@path))
      data["token"] || raise("empty token in credentials file")
    rescue Errno::ENOENT
      raise "credentials file not found at #{@path}"
    end
  end
end
