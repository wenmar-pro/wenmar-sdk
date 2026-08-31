# frozen_string_literal: true

require "uri"

module Wenmar
  class Paginator
    include Enumerable

    attr_reader :data, :links, :meta, :next_url, :client

    def initialize(client, page)
      @client = client
      @data = page.fetch("data", page.fetch(:data, []))
      @links = page.fetch("links", page.fetch(:links, {}))
      @meta = page.fetch("meta", page.fetch(:meta, {}))
      @page = page
      @next_url = @links["next"] || @links[:next]
    end

    def has_next?
      !@next_url.nil? && !@next_url.empty?
    end

    def next_page
      return nil unless has_next?

      unless same_origin?(@next_url, @client.base_url)
        raise Wenmar::Error.new(
          code: "invalid_pagination",
          message: "pagination next URL is not same-origin as base URL",
          field_errors: {},
          status: 0
        )
      end

      response = @client.send(:get_raw, @next_url)
      @next_url = self.class.parse_link_header(response.headers["Link"], "next")
      JSON.parse(response.body)
    end

    def each
      return to_enum unless block_given?

      page = self
      loop do
        page.data.each { |item| yield item }
        page = page.next_page
        break if page.nil?
      end
    end

    def to_a(max = 1000)
      result = []
      page = self
      loop do
        result.concat(page.data)
        break if result.size >= max

        page = page.next_page
        break if page.nil?
      end
      result.first(max)
    end

    def same_origin?(url, base_url)
      parsed = URI.parse(url)
      base = URI.parse(base_url)
      parsed.scheme == base.scheme && parsed.host == base.host && parsed.port == base.port
    rescue URI::InvalidURIError
      false
    end
    private :same_origin?

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
      new(client, { "data" => [], "links" => { "next" => next_url }, "meta" => {} })
    end
  end
end
