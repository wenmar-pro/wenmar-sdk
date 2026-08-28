require "faraday"
require "faraday/retry"
require "json"
require "uri"

module Wenmar
  class Client
    attr_reader :base_url, :token

    def initialize(token:, base_url: "https://app.wenmarpro.com")
      raise ArgumentError, "API token is required" if token.nil? || token.empty?

      @token = token
      @base_url = base_url
      @connection = build_connection
      @cache = {}
    end

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

    private

    def wrap_with_paginator(response)
      result = handle_response(response)
      client = self
      result.define_singleton_method(:paginator) do
        Paginator.from_response(response, client)
      end
      result
    end

    def build_connection
      Faraday.new(url: @base_url) do |conn|
        conn.headers["Authorization"] = "Bearer #{@token}"
        conn.headers["Content-Type"] = "application/json"
        conn.request :retry, max: 3, interval: 0.5, backoff_factor: 2,
                    retry_statuses: [429, 500, 502, 503, 504],
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

      response = @connection.get(path, params) do |req|
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

    def path_with_query(path, params)
      return path if params.empty?

      "#{path}?#{URI.encode_www_form(params)}"
    end

    def post(path, body)
      @connection.post(path) do |req|
        req.body = body.to_json
      end
    end

    def patch(path, body)
      @connection.patch(path) do |req|
        req.body = body.to_json
      end
    end

    def delete(path)
      @connection.delete(path)
    end

    def handle_response(response)
      if response.status >= 400
        raise Wenmar::Error.from_response(response)
      end

      JSON.parse(response.body)
    end
  end
end
