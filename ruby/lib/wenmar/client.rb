require "faraday"
require "faraday/retry"
require "json"
require "uri"

module Wenmar
  class Client
    attr_reader :base_url, :token, :config

    def initialize(config: Config.new, token_provider: nil, token: nil, base_url: nil)
      # Backwards-compatible: if token: is given directly, wrap in StaticTokenProvider.
      if token_provider.nil?
        t = token || raise(ArgumentError, "token or token_provider is required")
        token_provider = StaticTokenProvider.new(t)
      end
      @token_provider = token_provider
      @config = config
      @base_url = base_url || config.base_url
      @token = token_provider.token
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

    def list_customers
      response = get("/customers")
      wrap_with_paginator(response)
    end

    def show_customer(id)
      response = get("/customers/#{id}")
      handle_response(response)
    end

    def create_customer(attrs)
      response = post("/customers", customer: attrs)
      handle_response(response)
    end

    def update_customer(id, attrs)
      response = patch("/customers/#{id}", customer: attrs)
      handle_response(response)
    end

    def list_vehicles
      response = get("/vehicles")
      handle_response(response)
    end

    def show_vehicle(id)
      response = get("/vehicles/#{id}")
      handle_response(response)
    end

    def create_vehicle(attrs)
      response = post("/vehicles", vehicle: attrs)
      handle_response(response)
    end

    def update_vehicle(id, attrs)
      response = patch("/vehicles/#{id}", vehicle: attrs)
      handle_response(response)
    end

    def delete_vehicle(id)
      response = delete("/vehicles/#{id}")
      raise Wenmar::Error.from_response(response) if response.status >= 400

      nil
    end

    def decode_vin(vin)
      response = get("/vehicles/vin_decode", { vin: vin })
      handle_response(response)
    end

    def check_duplicate(vin)
      response = get("/vehicles/check_duplicate", { vin: vin })
      handle_response(response)
    end

    def list_work_orders
      response = get("/work_orders")
      wrap_with_paginator(response)
    end

    def show_work_order(id)
      response = get("/work_orders/#{id}")
      handle_response(response)
    end

    def create_work_order(attrs)
      response = post("/work_orders", work_order: attrs)
      handle_response(response)
    end

    def update_work_order(id, attrs)
      response = patch("/work_orders/#{id}", work_order: attrs)
      handle_response(response)
    end

    def delete_work_order(id)
      response = delete("/work_orders/#{id}")
      raise Wenmar::Error.from_response(response) if response.status >= 400

      nil
    end

    def list_account
      response = get("/account")
      handle_response(response)
    end

    def show_location(id)
      response = get("/locations/#{id}")
      handle_response(response)
    end

    def for_location(location_id)
      LocationClient.new(self, location_id)
    end

    class LocationClient
      attr_reader :parent, :location_id

      def initialize(parent, location_id)
        @parent = parent
        @location_id = location_id
      end

      # Delegate all operations to the parent client. The location_id is
      # documented in SPEC.md as a known limitation — the server currently
      # pins a token to one location. This sub-client is a guard/documentation
      # surface for callers who want to be explicit about the location context.
      def method_missing(name, *args, **kwargs, &block)
        @parent.send(name, *args, **kwargs, &block)
      end

      def respond_to_missing?(name, include_private = false)
        @parent.respond_to?(name, include_private) || super
      end
    end

    private

    def wrap_with_paginator(response)
      result = handle_response(response)
      client = self
      result.define_singleton_method(:paginator) do
        Paginator.from_response(response, client)
      end
      result.define_singleton_method(:meta) do
        next_url = Paginator.parse_link_header(response.headers["Link"], "next")
        {
          total_count: response.headers["X-Total-Count"]&.to_i,
          per_page: response.headers["X-Per-Page"]&.to_i,
          has_more: !next_url.nil? && !next_url.empty?
        }
      end
      result
    end

    def build_connection(retry_statuses:, methods: Faraday::Retry::Middleware::IDEMPOTENT_METHODS)
      Faraday.new(url: @base_url) do |conn|
        conn.headers["Authorization"] = "Bearer #{@token}"
        conn.headers["Content-Type"] = "application/json"
        conn.headers["User-Agent"] = "wenmar-sdk-ruby"
        conn.request :retry, max: 3, interval: 0.5, backoff_factor: 2,
                    retry_statuses: retry_statuses,
                    methods: methods,
                    exceptions: [Faraday::Error, Faraday::ServerError],
                    retry_block: lambda { |env:, **_kwargs|
                      if env.response && env.response.headers["Retry-After"]
                        delay = env.response.headers["Retry-After"].to_i
                        sleep(delay) if delay.positive?
                      end
                    }
      end
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
        return Faraday::Response.new(
          status: 304,
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

      response
    end

    def get_raw(url)
      response = @read_connection.get(url) do |req|
        req.headers["Accept"] = "application/json"
      end
      raise Wenmar::Error.from_response(response) if response.status >= 400

      response
    end

    def path_with_query(path, params)
      return path if params.empty?

      "#{path}?#{URI.encode_www_form(params)}"
    end

    def post(path, body)
      @write_connection.post(path) do |req|
        req.body = body.to_json
      end
    end

    def patch(path, body)
      @write_connection.patch(path) do |req|
        req.body = body.to_json
      end
    end

    def delete(path)
      @write_connection.delete(path)
    end

    def handle_response(response)
      if response.status >= 400
        raise Wenmar::Error.from_response(response)
      end

      JSON.parse(response.body)
    end
  end
end
