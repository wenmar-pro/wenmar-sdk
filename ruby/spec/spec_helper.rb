require "wenmar_pro"
require "minitest/autorun"
require "webmock/minitest"

WebMock.disable_net_connect!

module WenmarPro
  class TestCase < Minitest::Test
    def setup
      @base_url = "https://api.example.com"
      WebMock.reset!
    end
  end
end
