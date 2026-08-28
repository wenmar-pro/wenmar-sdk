package auth

import "time"

// Token represents an API access token, optionally with refresh and expiry
// metadata for OAuth flows.
type Token struct {
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	TokenType    string     `json:"token_type,omitempty"` // default "Bearer"
}

// IsExpired reports whether the token has already expired.
func (t *Token) IsExpired() bool {
	return t.ExpiresAt != nil && time.Now().After(*t.ExpiresAt)
}

// WillExpireWithin reports whether the token will expire within the given
// duration from now.
func (t *Token) WillExpireWithin(d time.Duration) bool {
	if t.ExpiresAt == nil {
		return false
	}
	return time.Now().Add(d).After(*t.ExpiresAt)
}
