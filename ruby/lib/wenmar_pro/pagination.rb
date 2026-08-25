require "uri"

module WenmarPro
  class Paginator
    attr_reader :next_url, :client

    def initialize(next_url:, client:)
      @next_url = next_url
      @client = client
    end

    def has_next?
      !@next_url.nil? && !@next_url.empty?
    end

    def next_page
      return nil unless has_next?

      page = extract_page_param(@next_url)
      if @next_url.include?("/api/customers")
        @client.list_customers(page: page)
      elsif @next_url.include?("/api/work_orders")
        @client.list_work_orders(page: page)
      end
    end

    def self.parse_link_header(header, rel)
      return nil if header.nil? || header.empty?

      header.split(",").each do |part|
        match = part.match(/<([^>]+)>;\s*rel="#{rel}"/)
        return match[1] if match
      end
      nil
    end

    def self.from_response(response, client)
      link_header = response.headers["Link"]
      next_url = parse_link_header(link_header, "next")
      new(next_url: next_url, client: client)
    end

    private

    def extract_page_param(url)
      uri = URI(url)
      params = URI.decode_www_form(uri.query.to_s).to_h
      params["page"]&.to_i
    end
  end
end
