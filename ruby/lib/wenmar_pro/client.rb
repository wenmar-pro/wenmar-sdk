require "faraday"
require "faraday/retry"
require "json"

module WenmarPro
  class Client
    attr_reader :base_url, :token

    def initialize(token:, base_url: "https://app.wenmarpro.com")
      raise ArgumentError, "API token is required" if token.nil? || token.empty?

      @token = token
      @base_url = base_url
      @connection = build_connection
    end

    def list_customers(page: nil)
      params = {}
      params[:page] = page if page
      response = get("/api/customers", params)
      result = handle_response(response)
      client = self
      result.define_singleton_method(:paginator) do
        Paginator.from_response(response, client)
      end
      result
    end

    def show_customer(id)
      response = get("/api/customers/#{id}")
      handle_response(response)
    end

    def create_customer(attrs)
      response = post("/api/customers", customer: attrs)
      handle_response(response)
    end

    def show_vehicle(id)
      response = get("/api/vehicles/#{id}")
      handle_response(response)
    end

    def list_work_orders(page: nil)
      params = {}
      params[:page] = page if page
      response = get("/api/work_orders", params)
      result = handle_response(response)
      client = self
      result.define_singleton_method(:paginator) do
        Paginator.from_response(response, client)
      end
      result
    end

    def show_work_order(id)
      response = get("/api/work_orders/#{id}")
      handle_response(response)
    end

    private

    def build_connection
      Faraday.new(url: @base_url) do |conn|
        conn.headers["Authorization"] = "Bearer #{@token}"
        conn.headers["Content-Type"] = "application/json"
        conn.request :retry, max: 3, interval: 0.5, backoff_factor: 2,
                    retry_statuses: [500, 502, 503, 504],
                    exceptions: [Faraday::Error, Faraday::ServerError]
      end
    end

    def get(path, params = {})
      @connection.get(path, params)
    end

    def post(path, body)
      @connection.post(path) do |req|
        req.body = body.to_json
      end
    end

    def handle_response(response)
      if response.status >= 400
        raise WenmarPro::Error.from_response(response)
      end

      JSON.parse(response.body)
    end
  end
end
