require_relative "spec_helper"

class OAuthTest < Wenmar::TestCase
  def test_refresh_posts_form_and_parses_token
    stub_request(:post, "https://api.example.com/oauth/token")
      .with(
        body: { grant_type: "refresh_token", client_id: "wenmar-cli", refresh_token: "rt" },
        headers: {
          "Content-Type" => "application/x-www-form-urlencoded",
          "Accept" => "application/json"
        }
      )
      .to_return(
        status: 200,
        headers: { "Content-Type" => "application/json" },
        body: { access_token: "new-access", refresh_token: "new-refresh", token_type: "Bearer", expires_in: 7200 }.to_json
      )

    token = Wenmar::OAuth.refresh(base_url: "https://api.example.com", refresh_token: "rt")
    assert_equal "new-access", token.access_token
    assert_equal "new-refresh", token.refresh_token
    assert_equal "Bearer", token.token_type
    assert token.expires_at > Time.now
    assert token.expires_at < Time.now + 7201
  end

  def test_refresh_trims_trailing_slash
    stub_request(:post, "https://api.example.com/oauth/token")
      .with(body: { grant_type: "refresh_token", client_id: "wenmar-cli", refresh_token: "rt" })
      .to_return(status: 200, body: { access_token: "new-access", expires_in: 7200 }.to_json)

    token = Wenmar::OAuth.refresh(base_url: "https://api.example.com/", refresh_token: "rt")
    assert_equal "new-access", token.access_token
  end

  def test_refresh_defaults_token_type_to_bearer
    stub_request(:post, "https://api.example.com/oauth/token")
      .to_return(status: 200, body: { access_token: "new-access", expires_in: 7200 }.to_json)

    token = Wenmar::OAuth.refresh(base_url: "https://api.example.com", refresh_token: "rt")
    assert_equal "Bearer", token.token_type
  end

  def test_refresh_does_not_follow_redirect
    stub_request(:post, "https://api.example.com/oauth/token")
      .to_return(status: 302, headers: { "Location" => "https://evil.example.com/token" }, body: "")
    stub_request(:post, "https://evil.example.com/token").to_return(status: 200, body: { access_token: "hacked" }.to_json)

    assert_raises(Wenmar::TokenError) do
      Wenmar::OAuth.refresh(base_url: "https://api.example.com", refresh_token: "rt")
    end
  end

  def test_refresh_error_json_raises_with_error_field
    stub_request(:post, "https://api.example.com/oauth/token")
      .to_return(status: 400, headers: { "Content-Type" => "application/json" }, body: { error: "invalid_grant" }.to_json)

    error = assert_raises(Wenmar::TokenError) do
      Wenmar::OAuth.refresh(base_url: "https://api.example.com", refresh_token: "bad")
    end
    assert_includes error.message, "invalid_grant"
  end

  def test_refresh_non_200_raises
    stub_request(:post, "https://api.example.com/oauth/token")
      .to_return(status: 500, body: "boom")

    error = assert_raises(Wenmar::TokenError) do
      Wenmar::OAuth.refresh(base_url: "https://api.example.com", refresh_token: "rt")
    end
    assert_includes error.message, "500"
  end

  def test_refresh_missing_access_token_raises
    stub_request(:post, "https://api.example.com/oauth/token")
      .to_return(status: 200, body: { refresh_token: "x", expires_in: 7200 }.to_json)

    error = assert_raises(Wenmar::TokenError) do
      Wenmar::OAuth.refresh(base_url: "https://api.example.com", refresh_token: "rt")
    end
    assert_includes error.message, "access_token"
  end
end
