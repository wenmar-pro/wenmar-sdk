require_relative "spec_helper"

module Wenmar
  class PaginationTest < TestCase
    def test_parse_link_header_next
      header = '<https://api.example.com/customers?page=2>; rel="next"'
      assert_equal "https://api.example.com/customers?page=2",
                   Paginator.parse_link_header(header, "next")
    end

    def test_parse_link_header_prev
      header = '<https://api.example.com/customers?page=1>; rel="prev", <https://api.example.com/customers?page=3>; rel="next"'
      assert_equal "https://api.example.com/customers?page=1",
                   Paginator.parse_link_header(header, "prev")
    end

    def test_parse_link_header_empty
      assert_nil Paginator.parse_link_header("", "next")
    end

    def test_has_next
      paginator = Paginator.new(next_url: "https://api.example.com?page=2", client: nil)
      assert paginator.has_next?

      paginator = Paginator.new(next_url: nil, client: nil)
      refute paginator.has_next?
    end

    def test_from_response
      stub_api(:get, "/customers", [], link: '<https://api.example.com/customers?page=2>; rel="next"')
      client = Client.new(token: "test", base_url: @base_url)
      response = client.list_customers
      paginator = response.paginator
      assert paginator.has_next?
      assert_equal "https://api.example.com/customers?page=2", paginator.next_url
    end

    def test_next_page_follows_actual_url
      # First response: page 1 data + Link header pointing to ?page=2
      stub_request(:get, "#{@base_url}/customers")
        .to_return(
          status: 200,
          body: [{ "id" => 1, "name" => "Page1" }].to_json,
          headers: { "Content-Type" => "application/json", "Link" => "<#{@base_url}/customers?page=2>; rel=\"next\"" }
        )

      # Second response: page 2 data (must be a DIFFERENT stub matching ?page=2)
      stub_request(:get, "#{@base_url}/customers?page=2")
        .to_return(
          status: 200,
          body: [{ "id" => 2, "name" => "Page2" }].to_json,
          headers: { "Content-Type" => "application/json" }
        )

      client = Client.new(token: "test", base_url: @base_url)
      result = client.list_customers
      assert_equal 1, result.first["id"]

      page2 = result.paginator.next_page
      assert_equal 2, page2.first["id"], "expected page 2 data (id=2), got #{page2.inspect}"
    end

    def test_next_page_rejects_cross_origin_url
      stub_request(:get, "#{@base_url}/customers").to_return(
        status: 200,
        body: [{ id: 1 }].to_json,
        headers: {
          "Content-Type" => "application/json",
          "Link" => '<https://attacker.example.com/customers?page=2>; rel="next"'
        }
      )

      client = Client.new(token: "my-token", base_url: @base_url)
      result = client.list_customers
      assert result.paginator.has_next?

      assert_raises(Wenmar::Error) { result.paginator.next_page }

      assert_not_requested(:get, %r{https://attacker\.example\.com})
    end

    private

    def stub_api(method, path, body, status: 200, link: nil)
      headers = { "Content-Type" => "application/json" }
      headers["Link"] = link if link
      stub_request(method, "#{@base_url}#{path}")
        .to_return(status: status, body: body.to_json, headers: headers)
    end
  end
end
