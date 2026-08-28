module Wenmar
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

      # Follow the actual next URL from the Link header, which includes
      # query params like ?page=2. This performs a raw GET against the URL.
      response = @client.send(:get_raw, @next_url)
      @next_url = self.class.parse_link_header(response.headers["Link"], "next")
      JSON.parse(response.body)
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
  end
end
