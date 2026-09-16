require_relative "spec_helper"

class TokenTest < Wenmar::TestCase
  def test_defaults_to_bearer
    token = Wenmar::Token.new(access_token: "abc")
    assert_equal "abc", token.access_token
    assert_equal "Bearer", token.token_type
    assert_nil token.refresh_token
    assert_nil token.expires_at
  end

  def test_expired_false_when_no_expiry
    refute Wenmar::Token.new(access_token: "abc").expired?
  end

  def test_expired_false_when_in_future
    token = Wenmar::Token.new(access_token: "abc", expires_at: Time.now + 3600)
    refute token.expired?
  end

  def test_expired_true_when_in_past
    token = Wenmar::Token.new(access_token: "abc", expires_at: Time.now - 1)
    assert token.expired?
  end

  def test_will_expire_within_false_when_no_expiry
    refute Wenmar::Token.new(access_token: "abc").will_expire_within?(5)
  end

  def test_will_expire_within_true_when_within_window
    token = Wenmar::Token.new(access_token: "abc", expires_at: Time.now + 60)
    assert token.will_expire_within?(300)
  end

  def test_will_expire_within_false_when_beyond_window
    token = Wenmar::Token.new(access_token: "abc", expires_at: Time.now + 3600)
    refute token.will_expire_within?(300)
  end

  def test_will_expire_within_boundary_uses_time_now
    expires = Time.now + 300
    token = Wenmar::Token.new(access_token: "abc", expires_at: expires)
    assert token.will_expire_within?(300)
  end

  def test_to_h_omits_nil
    hash = Wenmar::Token.new(access_token: "abc").to_h
    assert_equal "abc", hash["access_token"]
    refute hash.key?("refresh_token")
    refute hash.key?("expires_at")
    assert_equal "Bearer", hash["token_type"]
  end

  def test_to_h_includes_expiry_as_iso8601
    expires = Time.at(1_700_000_000)
    hash = Wenmar::Token.new(access_token: "abc", refresh_token: "rt", expires_at: expires).to_h
    assert_equal "rt", hash["refresh_token"]
    assert_equal expires.utc.iso8601, hash["expires_at"]
  end

  def test_from_h_round_trip
    expires = Time.at(1_700_000_000)
    original = Wenmar::Token.new(access_token: "abc", refresh_token: "rt", expires_at: expires, token_type: "bearer")
    round = Wenmar::Token.from_h(original.to_h)
    assert_equal "abc", round.access_token
    assert_equal "rt", round.refresh_token
    assert_equal "bearer", round.token_type
    assert_in_delta expires.to_f, round.expires_at.to_f, 1
  end

  def test_from_h_tolerates_symbol_keys
    token = Wenmar::Token.from_h({access_token: "abc", refresh_token: "rt"})
    assert_equal "abc", token.access_token
    assert_equal "rt", token.refresh_token
  end

  def test_from_h_empty_hash_raises
    assert_raises(Wenmar::TokenError) { Wenmar::Token.from_h({}) }
  end

  def test_from_h_nil_raises
    assert_raises(Wenmar::TokenError) { Wenmar::Token.from_h(nil) }
  end

  def test_from_h_invalid_expires_at_raises_token_error
    assert_raises(Wenmar::TokenError) do
      Wenmar::Token.from_h("access_token" => "abc", "expires_at" => "not-a-time")
    end
  end
end
