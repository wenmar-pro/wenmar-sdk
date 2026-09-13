package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func TestNewAuthManagerWithOAuth_Refreshes(t *testing.T) {
	const (
		clientID = "wenmar-cli"
		rt       = "stored-refresh"
	)

	var (
		method  string
		path    string
		ctype   string
		formVal url.Values
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		path = r.URL.Path
		ctype = r.Header.Get("Content-Type")
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		formVal = r.Form
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token":  "new",
			"refresh_token": "rotated",
			"token_type":    "Bearer",
			"expires_in":    7200,
		})
	}))
	defer srv.Close()

	store := &memoryStore{token: &Token{AccessToken: "old", RefreshToken: rt}}
	m := NewAuthManagerWithOAuth(store, nil, srv.URL, clientID)

	if err := m.Refresh(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if method != http.MethodPost {
		t.Errorf("expected method POST, got %q", method)
	}
	if path != "/oauth/token" {
		t.Errorf("expected path /oauth/token, got %q", path)
	}
	if !strings.HasPrefix(ctype, "application/x-www-form-urlencoded") {
		t.Errorf("expected form Content-Type, got %q", ctype)
	}
	if got := formVal.Get("grant_type"); got != "refresh_token" {
		t.Errorf("expected grant_type refresh_token, got %q", got)
	}
	if got := formVal.Get("client_id"); got != clientID {
		t.Errorf("expected client_id %q, got %q", clientID, got)
	}
	if got := formVal.Get("refresh_token"); got != rt {
		t.Errorf("expected refresh_token %q, got %q", rt, got)
	}

	if store.token.AccessToken != "new" {
		t.Errorf("expected stored access token 'new', got %q", store.token.AccessToken)
	}
	if store.token.RefreshToken != "rotated" {
		t.Errorf("expected rotated refresh token 'rotated', got %q", store.token.RefreshToken)
	}
	if store.token.ExpiresAt == nil {
		t.Error("expected non-nil ExpiresAt")
	} else if !store.token.ExpiresAt.After(time.Now()) {
		t.Errorf("expected ExpiresAt in the future, got %v", store.token.ExpiresAt)
	}
}

func TestNewAuthManagerWithOAuth_EmptyBaseURL(t *testing.T) {
	store := &memoryStore{token: &Token{AccessToken: "old", RefreshToken: "refresh"}}
	m := NewAuthManagerWithOAuth(store, nil, "", "wenmar-cli")

	err := m.Refresh(context.Background())
	if !errors.Is(err, ErrOAuthNotImplemented) {
		t.Errorf("expected ErrOAuthNotImplemented, got %v", err)
	}
}

func TestNewAuthManagerWithOAuth_EmptyClientID(t *testing.T) {
	store := &memoryStore{token: &Token{AccessToken: "old", RefreshToken: "refresh"}}
	m := NewAuthManagerWithOAuth(store, nil, "https://example.com", "")

	err := m.Refresh(context.Background())
	if !errors.Is(err, ErrOAuthNotImplemented) {
		t.Errorf("expected ErrOAuthNotImplemented, got %v", err)
	}
}

func TestNewAuthManagerWithOAuth_TrimsTrailingSlash(t *testing.T) {
	var path string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token":  "new",
			"refresh_token": "rotated",
			"token_type":    "Bearer",
			"expires_in":    7200,
		})
	}))
	defer srv.Close()

	store := &memoryStore{token: &Token{AccessToken: "old", RefreshToken: "refresh"}}
	m := NewAuthManagerWithOAuth(store, nil, srv.URL+"/", "wenmar-cli")

	if err := m.Refresh(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if path != "/oauth/token" {
		t.Errorf("expected path /oauth/token, got %q", path)
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
