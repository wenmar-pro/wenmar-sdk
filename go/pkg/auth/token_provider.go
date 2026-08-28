package auth

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

// StaticTokenProvider wraps a fixed string. Used for the --token flag and
// WENMAR_TOKEN env var.
type StaticTokenProvider struct {
	value string
}

// NewStaticTokenProvider creates a provider that always returns the given token.
func NewStaticTokenProvider(token string) *StaticTokenProvider {
	return &StaticTokenProvider{value: token}
}

// Token returns the fixed token.
func (p *StaticTokenProvider) Token(_ context.Context) (string, error) {
	if p.value == "" {
		return "", fmt.Errorf("token is empty")
	}
	return p.value, nil
}

// CredentialStoreProvider reads from a CredentialStore, auto-refreshing via
// the AuthManager when the token is expired or near-expiry (5 min window).
type CredentialStoreProvider struct {
	Store   CredentialStore
	Manager *AuthManager
}

// Token returns the current access token, refreshing it if it is expired or
// about to expire.
func (p *CredentialStoreProvider) Token(ctx context.Context) (string, error) {
	tok, err := p.Store.GetToken(ctx)
	if err != nil {
		return "", err
	}
	if tok == nil || tok.AccessToken == "" {
		return "", fmt.Errorf("no token stored")
	}
	if tok.IsExpired() || tok.WillExpireWithin(refreshWindow) {
		if p.Manager != nil {
			if err := p.Manager.Refresh(ctx); err != nil {
				return "", err
			}
			tok, err = p.Store.GetToken(ctx)
			if err != nil {
				return "", err
			}
		}
	}
	if tok == nil || tok.AccessToken == "" {
		return "", fmt.Errorf("no token stored")
	}
	return tok.AccessToken, nil
}
