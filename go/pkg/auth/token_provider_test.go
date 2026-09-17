package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestStaticTokenProvider(t *testing.T) {
	p := NewStaticTokenProvider("my-token")
	tok, err := p.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok != "my-token" {
		t.Errorf("expected 'my-token', got %q", tok)
	}
}

func TestStaticTokenProvider_Empty(t *testing.T) {
	p := NewStaticTokenProvider("")
	if _, err := p.Token(context.Background()); err == nil {
		t.Error("expected error for empty token")
	}
}

func TestBearerAuth_Authenticate(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	auth := &BearerAuth{Provider: NewStaticTokenProvider("secret")}
	if err := auth.Authenticate(req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := req.Header.Get("Authorization"); got != "Bearer secret" {
		t.Errorf("expected 'Bearer secret', got %q", got)
	}
}

func TestBearerAuth_NilProvider(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	auth := &BearerAuth{}
	if err := auth.Authenticate(req); err == nil {
		t.Error("expected error for nil provider")
	}
}

func TestCredentialStoreProvider_ConcurrentRefreshSingleFlight(t *testing.T) {
	var refreshCalls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&refreshCalls, 1)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"access_token":"fresh-token","refresh_token":"rt2","expires_in":3600}`)
	}))
	defer srv.Close()

	dir := t.TempDir()
	store := FileStore{Path: filepath.Join(dir, "credentials.json")}
	nearExpiry := time.Now().Add(1 * time.Minute)
	if err := store.SaveToken(context.Background(), &Token{
		AccessToken: "stale", RefreshToken: "rt1", ExpiresAt: &nearExpiry,
	}); err != nil {
		t.Fatalf("seed: %v", err)
	}

	m := NewAuthManagerWithOAuth(store, nil, srv.URL, "wenmar-cli")
	p := &CredentialStoreProvider{Store: store, Manager: m}

	var wg sync.WaitGroup
	tokens := make([]string, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			tok, err := p.Token(context.Background())
			if err != nil {
				t.Errorf("Token: %v", err)
				return
			}
			tokens[i] = tok
		}(i)
	}
	wg.Wait()

	if n := atomic.LoadInt32(&refreshCalls); n != 1 {
		t.Errorf("expected exactly 1 refresh call, got %d", n)
	}
	for _, tok := range tokens {
		if tok != "fresh-token" {
			t.Errorf("expected refreshed token, got %q", tok)
		}
	}
}
