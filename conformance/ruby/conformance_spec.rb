require "minitest/autorun"
require "json"
require "webmock/minitest"
require "wenmar"
require_relative "dispatch.gen"

WebMock.disable_net_connect!

module Conformance
  # paginate_ruby follows Link-header pages from a list result and returns
  # the final page's data. Used by the generated dispatch lambdas.
  def self.paginate_ruby(client, result)
    paginator = result.respond_to?(:paginator) ? result.paginator : nil
    return result unless paginator

    while paginator.has_next?
      result = paginator.next_page
    end
    result
  end

  class ConformanceTest < Minitest::Test
    TESTS_DIR = File.expand_path("../tests", __dir__)
    BASE_URL = "https://api.example.com"

    def test_conformance
      load_cases.each do |tc|
        run_case(tc)
      end
    end

    private

    def load_cases
      Dir.glob(File.join(TESTS_DIR, "*.json")).flat_map do |file|
        JSON.parse(File.read(file))
      end
    end

    def run_case(tc)
      WebMock.reset!
      captured_headers = {}
      stub_responses(tc, captured_headers)
      client = Wenmar::Client.new(token: "test-token", base_url: BASE_URL)

      fn = DISPATCH[tc["operation"]]
      flunk "[#{tc["name"]}] operation #{tc["operation"]} not in dispatch" unless fn

      args = {
        "pathParams" => tc["pathParams"] || {},
        "query" => tc["query"] || {},
        "requestBody" => tc["requestBody"] || {}
      }

      begin
        result = fn.call(client, args)
        if tc.dig("expect", "noError")
          assert_body_path(result, tc.dig("expect", "responseBody"), tc["name"]) if tc.dig("expect", "responseBody")
        else
          flunk "[#{tc["name"]}] expected error, got success"
        end
      rescue Wenmar::Error => e
        if tc.dig("expect", "noError")
          flunk "[#{tc["name"]}] expected success, got error: #{e.message}"
        else
          assert_equal tc.dig("expect", "errorCode"), e.code, "[#{tc["name"]}] error code" if tc.dig("expect", "errorCode")
          assert_equal tc.dig("expect", "errorStatus"), e.status, "[#{tc["name"]}] error status" if tc.dig("expect", "errorStatus")
          if tc.dig("expect", "fieldErrors")
            assert_equal tc.dig("expect", "fieldErrors"), e.field_errors_by_field, "[#{tc["name"]}] field errors"
          end
        end
      end

      assert_request_count(tc) if tc.dig("expect", "requestCount")
      assert_no_outbound_request(tc) if tc.dig("expect", "assertNoOutboundRequest")
      assert_request_headers(tc, captured_headers) if tc.dig("expect", "requestHeaders")
    end

    # buildArgs is used by the generated dispatch lambdas to reconstruct the
    # SDK call arguments from the conformance test case.
    def buildArgs(client, args)
      args
    end

    def assert_no_outbound_request(tc)
      expected = tc["mockResponses"].length
      actual = WebMock::RequestRegistry.instance.requested_signatures.hash.values.sum
      assert_equal expected, actual, "[#{tc["name"]}] expected no outbound request beyond mocks (got #{actual} calls)"
    end

    def stub_responses(tc, captured_headers = nil)
      responses = tc["mockResponses"].map do |resp|
        headers = { "Content-Type" => "application/json" }
        (resp["headers"] || {}).each do |k, v|
          headers[k] = v.gsub("{server}", BASE_URL)
        end
        body = resp["body"].nil? ? "" : resp["body"].to_json
        { status: resp["status"], body: body, headers: headers }
      end

      stub = stub_request(tc["method"].downcase.to_sym, /#{Regexp.escape(BASE_URL)}#{Regexp.escape(tc["path"])}(\?.*)?\z/)
        .to_return(responses)
      if captured_headers
        stub.with { |request| captured_headers.merge!(request.headers); true }
      end
    end

    def assert_request_count(tc)
      expected = tc.dig("expect", "requestCount")
      actual = WebMock::RequestRegistry.instance.requested_signatures.hash.values.sum
      assert_equal expected, actual, "[#{tc["name"]}] expected #{expected} requests, got #{actual}"
    end

    def assert_request_headers(tc, captured_headers)
      expected = tc.dig("expect", "requestHeaders")
      expected.each do |k, v|
        actual = captured_headers[k] || captured_headers[k.to_s]
        assert_equal v, actual, "[#{tc["name"]}] expected header #{k}=#{v}, got #{actual}"
      end
    end

    def assert_body_path(result, assertion, name)
      value = navigate_path(result, assertion["path"])
      assert_equal assertion["equals"], value, "[#{name}] expected #{assertion["path"]} to equal #{assertion["equals"]}"
    end

    def navigate_path(obj, path)
      path.split(".").reduce(obj) do |current, part|
        if part =~ /\A\d+\z/
          current[part.to_i]
        else
          current[part]
        end
      end
    end
  end
end
