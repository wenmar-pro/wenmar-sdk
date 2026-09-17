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

      fetch_next
    end

    def each
      return to_enum unless block_given?

      @data.each { |item| yield item }
      while has_next?
        data = fetch_next
        data.each { |item| yield item }
      end
    end

    def to_a(max = 1000)
      result = @data.dup
      result.concat(fetch_next) while has_next? && result.size < max
      result.first(max)
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
      body = response.body
      parsed = body.is_a?(String) ? JSON.parse(body) : body
      data = parsed.is_a?(Array) ? parsed : []
      new(client, {"data" => data, "links" => {"next" => next_url}, "meta" => {}})
    end

    private

    def fetch_next
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
      body = response.body
      parsed = body.is_a?(String) ? JSON.parse(body) : body
      parsed.is_a?(Array) ? parsed : []
    end

    def same_origin?(url, base_url)
      parsed = URI.parse(url)
      base = URI.parse(base_url)
      parsed.scheme == base.scheme && parsed.host == base.host && parsed.port == base.port
    rescue URI::InvalidURIError
      false
    end
  end
end
