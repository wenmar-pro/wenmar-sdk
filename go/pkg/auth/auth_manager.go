package auth

import (
	"context"
	"errors"
	"time"
)

// refreshWindow is how far before expiry the provider attempts a refresh.
const refreshWindow = 5 * time.Minute

// ErrOAuthNotImplemented is returned by the default refresh function. OAuth
// token refresh is not yet implemented; users must re-run `wenmar auth login`.
var ErrOAuthNotImplemented = errors.New("OAuth token refresh is not yet implemented. Re-run `wenmar auth login` to get a new token")

// AuthManager coordinates token storage, retrieval, and refresh. When OAuth
// lands, only refreshFn changes — no caller code breaks.
type AuthManager struct {
	Store    CredentialStore
	Provider TokenProvider
	refreshFn func(ctx context.Context, refreshToken string) (*Token, error)
}

// NewAuthManager creates an AuthManager with the default (OAuth-stub) refresh
// function.
func NewAuthManager(store CredentialStore, provider TokenProvider) *AuthManager {
	return &AuthManager{
		Store:    store,
		Provider: provider,
		refreshFn: func(context.Context, string) (*Token, error) {
			return nil, ErrOAuthNotImplemented
		},
	}
}

// SetRefreshFn overrides the refresh function. Used when OAuth is implemented.
func (m *AuthManager) SetRefreshFn(fn func(ctx context.Context, refreshToken string) (*Token, error)) {
	if fn != nil {
		m.refreshFn = fn
	}
}

// Token returns the current access token, refreshing if needed.
func (m *AuthManager) Token(ctx context.Context) (string, error) {
	if m.Provider != nil {
		return m.Provider.Token(ctx)
	}
	tok, err := m.Store.GetToken(ctx)
	if err != nil {
		return "", err
	}
	if tok == nil {
		return "", errors.New("no token stored")
	}
	return tok.AccessToken, nil
}

// Refresh attempts to refresh the stored token via refreshFn.
func (m *AuthManager) Refresh(ctx context.Context) error {
	tok, err := m.Store.GetToken(ctx)
	if err != nil {
		return err
	}
	if tok == nil || tok.RefreshToken == "" {
		return ErrOAuthNotImplemented
	}
	newTok, err := m.refreshFn(ctx, tok.RefreshToken)
	if err != nil {
		return err
	}
	return m.Store.SaveToken(ctx, newTok)
}

// Logout clears the stored credentials.
func (m *AuthManager) Logout(ctx context.Context) error {
	return m.Store.DeleteToken(ctx)
}
