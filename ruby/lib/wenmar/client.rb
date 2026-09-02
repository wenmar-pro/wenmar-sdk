# frozen_string_literal: true

require "faraday"
require "faraday/retry"
require "json"
require "uri"
require_relative "version"
require_relative "error"
require_relative "auth"
require_relative "config"
require_relative "credentials"
require_relative "pagination"

module Wenmar
  class Client
    DEFAULT_BASE_URL = "https://app.wenmarpro.com"

    attr_reader :base_url, :location_id, :config

    def initialize(config = nil, token: nil, base_url: nil, token_provider: nil)
      @config = config.is_a?(Config) ? config : Config.new(config || {})
      @config.access_token = token if token
      @config.token_provider = token_provider if token_provider
      @config.base_url = base_url if base_url

      @base_url = (@config.base_url || DEFAULT_BASE_URL).to_s.sub(%r{/+\z}, "")
      @location_id = @config.location_id
      raise ArgumentError, "base_url must use https (http only allowed for localhost)" unless https_or_localhost?(@base_url)

      @read_connection = build_connection(retry_statuses: [429, 500, 502, 503, 504])
      @write_connection = build_connection(retry_statuses: [429], methods: %i[post patch delete])
      @cache = {}
    end

    def https_or_localhost?(url)
      parsed = URI.parse(url)
      return true if parsed.scheme == "https"

      %w[localhost 127.0.0.1].include?(parsed.host)
    rescue URI::InvalidURIError
      false
    end
    private :https_or_localhost?

    def token_provider
      @config.token_provider || StaticTokenProvider.new(@config.access_token)
    end

    def for_location(location_id)
      scoped = dup
      scoped.instance_variable_set(:@location_id, location_id)
      scoped.instance_variable_set(:@read_connection, build_connection(retry_statuses: [429, 500, 502, 503, 504], location_id: location_id))
      scoped.instance_variable_set(:@write_connection, build_connection(retry_statuses: [429], methods: %i[post patch delete], location_id: location_id))
      scoped
    end

    def get(path, params = {})
      cache_key = path_with_query(path, params)
      cached = @cache[cache_key]

      response = @read_connection.get(path, params) do |req|
        if cached
          req.headers["If-None-Match"] = cached[:etag] if cached[:etag]
          req.headers["If-Modified-Since"] = cached[:last_modified] if cached[:last_modified]
        end
      end

      if response.status == 304 && cached
        response = Faraday::Response.new(
          status: 200,
          response_headers: { "Content-Type" => "application/json" },
          body: cached[:body]
        )
      end

      if response.status == 200 && response.headers["ETag"]
        @cache[cache_key] = {
          etag: response.headers["ETag"],
          last_modified: response.headers["Last-Modified"],
          body: response.body
        }
      end

      handle_response(response)
    end

    def get_raw(url)
      response = @read_connection.get(url) do |req|
        req.headers["Accept"] = "application/json"
      end
      raise Wenmar::Error.from_response(response) if response.status >= 400

      response
    end

    def post(path, body = nil, params = {})
      response = @write_connection.post(path) do |req|
        req.params.merge!(params) unless params.empty?
        req.body = body.to_json if body
      end
      handle_response(response)
    end

    def patch(path, body = nil, params = {})
      response = @write_connection.patch(path) do |req|
        req.params.merge!(params) unless params.empty?
        req.body = body.to_json if body
      end
      handle_response(response)
    end

    def delete(path, params = {})
      response = @write_connection.delete(path) do |req|
        req.params.merge!(params) unless params.empty?
      end
      handle_response(response)
    end

    # paginator_to_a collects all pages from a paginated list result, up to
    # max items (default 1000).
    def paginator_to_a(result, max = 1000)
      return result unless result.respond_to?(:paginator)

      paginator = result.paginator
      return result unless paginator

      items = result.dup
      while paginator.has_next? && items.size < max
        page = paginator.next_page
        break if page.nil?

        items.concat(page)
      end
      items.first(max)
    end

    private

    def build_connection(retry_statuses:, methods: Faraday::Retry::Middleware::IDEMPOTENT_METHODS, location_id: @location_id)
      Faraday.new(url: @base_url) do |conn|
        conn.headers["Accept"] = "application/json"
        conn.headers["Content-Type"] = "application/json"
        conn.headers["User-Agent"] = "wenmar-sdk-ruby/#{Wenmar::VERSION}"
        conn.headers["X-Wenmar-Location"] = location_id if location_id
        conn.request :authorization, "Bearer", -> { resolve_token }
        conn.request :retry, max: @config.max_retries, interval: 0.1, backoff_factor: 2,
                    retry_statuses: retry_statuses,
                    methods: methods,
                    exceptions: [Faraday::Error, Faraday::ServerError]
        conn.adapter Faraday.default_adapter
      end
    end

    def resolve_token
      token_provider.token
    rescue TokenError => e
      raise Error.new(code: "auth_failed", message: e.message)
    end

    def path_with_query(path, params)
      return path if params.nil? || params.empty?

      "#{path}?#{URI.encode_www_form(params)}"
    end

    def handle_response(response)
      if response.status >= 400
        raise Wenmar::Error.from_response(response)
      end

      body = response.body
      return nil if body.nil? || body == ""

      body = JSON.parse(body) if body.is_a?(String)

      # Attach a paginator to list responses that carry a Link header.
      if body.is_a?(Array) && response.headers["Link"]
        client = self
        body.define_singleton_method(:paginator) do
          Paginator.from_response(response, client)
        end
      end
      body
    end
  end
end
