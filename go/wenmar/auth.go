package wenmar

import (
	"context"
	"fmt"
)

// TokenProvider supplies an API token. Implementations may read from a static
// value, the system keyring, or a credential file. The token is fetched per
// request so providers can rotate or refresh tokens.
type TokenProvider interface {
	Token(ctx context.Context) (string, error)
}

// StaticTokenProvider returns a fixed token. Suitable for simple scripts and
// tests. For production use, prefer CredentialStore which persists the token
// securely.
type StaticTokenProvider struct {
	token string
}

// NewStaticTokenProvider creates a TokenProvider that always returns the
// given token.
func NewStaticTokenProvider(token string) *StaticTokenProvider {
	return &StaticTokenProvider{token: token}
}

func (p *StaticTokenProvider) Token(_ context.Context) (string, error) {
	if p.token == "" {
		return "", fmt.Errorf("token is empty")
	}
	return p.token, nil
}
