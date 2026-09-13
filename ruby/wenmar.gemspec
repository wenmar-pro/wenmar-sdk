require_relative "lib/wenmar/version"

Gem::Specification.new do |spec|
  spec.name          = "wenmar"
  spec.version       = Wenmar::VERSION
  spec.summary       = "Ruby SDK for the Wenmar Pro API"
  spec.description   = "Idiomatic Ruby client for the Wenmar Pro automotive shop management API."
  spec.authors       = ["Ben D'Angelo"]
  spec.email         = ["ben@wenmarpro.com"]
  spec.homepage      = "https://github.com/wenmar-pro/wenmar-sdk"
  spec.license       = "MIT"
  spec.required_ruby_version = ">= 3.1.0"

  spec.files = Dir.glob("lib/**/*") + ["README.md", "LICENSE"]
  spec.require_paths = ["lib"]

  spec.metadata = {
    "source_code_uri" => "https://github.com/wenmar-pro/wenmar-sdk",
    "changelog_uri"   => "https://github.com/wenmar-pro/wenmar-sdk/blob/main/MIGRATING.md"
  }

  spec.add_dependency "faraday", "~> 2.0"
  spec.add_dependency "faraday-retry", "~> 2.0"

  spec.add_development_dependency "ruby-keychain", "~> 0.3"
end
