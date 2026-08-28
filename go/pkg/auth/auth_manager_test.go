package auth

import (
	"context"
	"errors"
	"testing"
	"time"
)

// memoryStore is a simple in-memory CredentialStore for tests.
type memoryStore struct {
	token *Token
}

func (m *memoryStore) GetToken(_ context.Context) (*Token, error) {
	if m.token == nil {
		return nil, errors.New("no token stored")
	}
	return m.token, nil
}

func (m *memoryStore) SaveToken(_ context.Context, t *Token) error {
	m.token = t
	return nil
}

func (m *memoryStore) DeleteToken(_ context.Context) error {
	m.token = nil
	return nil
}

func TestAuthManager_Token_FromProvider(t *testing.T) {
	store := &memoryStore{token: &Token{AccessToken: "stored"}}
	provider := NewStaticTokenProvider("from-provider")
	m := NewAuthManager(store, provider)

	tok, err := m.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok != "from-provider" {
		t.Errorf("expected provider token, got %q", tok)
	}
}

func TestAuthManager_Token_FromStore(t *testing.T) {
	store := &memoryStore{token: &Token{AccessToken: "stored"}}
	m := NewAuthManager(store, nil)

	tok, err := m.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok != "stored" {
		t.Errorf("expected stored token, got %q", tok)
	}
}

func TestAuthManager_Refresh_NotImplemented(t *testing.T) {
	store := &memoryStore{token: &Token{AccessToken: "stored", RefreshToken: "refresh"}}
	m := NewAuthManager(store, nil)

	err := m.Refresh(context.Background())
	if !errors.Is(err, ErrOAuthNotImplemented) {
		t.Errorf("expected ErrOAuthNotImplemented, got %v", err)
	}
}

func TestAuthManager_Refresh_NoRefreshToken(t *testing.T) {
	store := &memoryStore{token: &Token{AccessToken: "stored"}}
	m := NewAuthManager(store, nil)

	err := m.Refresh(context.Background())
	if !errors.Is(err, ErrOAuthNotImplemented) {
		t.Errorf("expected ErrOAuthNotImplemented, got %v", err)
	}
}

func TestAuthManager_Refresh_CustomFn(t *testing.T) {
	store := &memoryStore{token: &Token{AccessToken: "old", RefreshToken: "refresh"}}
	m := NewAuthManager(store, nil)
	m.SetRefreshFn(func(_ context.Context, rt string) (*Token, error) {
		if rt != "refresh" {
			t.Errorf("expected refresh token 'refresh', got %q", rt)
		}
		return &Token{AccessToken: "new-token"}, nil
	})

	if err := m.Refresh(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if store.token.AccessToken != "new-token" {
		t.Errorf("expected stored token 'new-token', got %q", store.token.AccessToken)
	}
}

func TestAuthManager_Logout(t *testing.T) {
	store := &memoryStore{token: &Token{AccessToken: "stored"}}
	m := NewAuthManager(store, nil)

	if err := m.Logout(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if store.token != nil {
		t.Error("expected token to be cleared")
	}
}

func TestCredentialStoreProvider_RefreshesExpired(t *testing.T) {
	expired := time.Now().Add(-time.Hour)
	store := &memoryStore{token: &Token{AccessToken: "old", RefreshToken: "refresh", ExpiresAt: &expired}}
	m := NewAuthManager(store, nil)
	m.SetRefreshFn(func(_ context.Context, _ string) (*Token, error) {
		return &Token{AccessToken: "new-token"}, nil
	})
	p := &CredentialStoreProvider{Store: store, Manager: m}

	tok, err := p.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok != "new-token" {
		t.Errorf("expected refreshed token 'new-token', got %q", tok)
	}
}

func TestCredentialStoreProvider_NoRefreshToken(t *testing.T) {
	expired := time.Now().Add(-time.Hour)
	store := &memoryStore{token: &Token{AccessToken: "old", ExpiresAt: &expired}}
	m := NewAuthManager(store, nil)
	p := &CredentialStoreProvider{Store: store, Manager: m}

	_, err := p.Token(context.Background())
	if !errors.Is(err, ErrOAuthNotImplemented) {
		t.Errorf("expected ErrOAuthNotImplemented, got %v", err)
	}
}
