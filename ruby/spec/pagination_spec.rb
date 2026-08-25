require_relative "spec_helper"

module Wenmar
  class PaginationTest < TestCase
    def test_parse_link_header_next
      header = '<https://api.example.com/api/customers?page=2>; rel="next"'
      assert_equal "https://api.example.com/api/customers?page=2",
                   Paginator.parse_link_header(header, "next")
    end

    def test_parse_link_header_prev
      header = '<https://api.example.com/api/customers?page=1>; rel="prev", <https://api.example.com/api/customers?page=3>; rel="next"'
      assert_equal "https://api.example.com/api/customers?page=1",
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
      stub_api(:get, "/api/customers", { data: [] }, link: '<https://api.example.com/api/customers?page=2>; rel="next"')
      client = Client.new(token: "test", base_url: @base_url)
      response = client.list_customers
      paginator = response.paginator
      assert paginator.has_next?
      assert_equal "https://api.example.com/api/customers?page=2", paginator.next_url
    end

    def test_next_page_follows_link
      stub_api(:get, "/api/customers", { data: [{ id: 1 }] }, link: '<https://api.example.com/api/customers?page=2>; rel="next"')
      stub_api(:get, "/api/customers?page=2", { data: [{ id: 2 }] })
      client = Client.new(token: "test", base_url: @base_url)
      paginator = Paginator.new(next_url: "#{@base_url}/api/customers?page=2", client: client)
      page = paginator.next_page
      assert_equal 2, page["data"].first["id"]
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
