package wenmar

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.BaseURL != "https://app.wenmarpro.com" {
		t.Errorf("expected default BaseURL, got %q", cfg.BaseURL)
	}
	if cfg.Timeout != 30*time.Second {
		t.Errorf("expected 30s timeout, got %v", cfg.Timeout)
	}
	if cfg.MaxRetries != 3 {
		t.Errorf("expected 3 max retries, got %d", cfg.MaxRetries)
	}
	if !cfg.CacheEnabled {
		t.Error("expected CacheEnabled=true by default")
	}
}

func TestConfig_DeepCopy(t *testing.T) {
	original := DefaultConfig()
	copy := original
	copy.BaseURL = "https://other.example.com"
	if original.BaseURL == copy.BaseURL {
		t.Error("deep copy failed: mutating copy affected original")
	}
}

// capturesAuthToken builds a client whose single request returns the resolved
// Authorization header. It fails the test if no request reaches the server.
func capturesAuthToken(t *testing.T, cfg Config) string {
	t.Helper()
	var auth atomic.Value
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth.Store(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	cfg.BaseURL = ts.URL
	c, err := NewClient(cfg, nil)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	if _, err := c.ListCustomers(ctx, nil); err != nil {
		t.Fatalf("request failed: %v", err)
	}
	v := auth.Load()
	if v == nil {
		t.Fatal("no request reached the server")
	}
	return v.(string)
}

func TestNewClient_ProviderArgPrecedence(t *testing.T) {
	cfg := DefaultConfig()
	cfg.TokenProvider = NewStaticTokenProvider("from-cfg-provider")
	cfg.Token = "from-cfg-token"
	c, err := NewClient(cfg, NewStaticTokenProvider("from-arg"))
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	if c.tp == nil {
		t.Fatal("expected token provider to be set")
	}
	tok, err := c.tp.Token(ctx)
	if err != nil {
		t.Fatalf("Token failed: %v", err)
	}
	if tok != "from-arg" {
		t.Errorf("expected arg provider to win, got %q", tok)
	}
}

func TestNewClient_ConfigTokenProviderUsedWhenArgNil(t *testing.T) {
	cfg := DefaultConfig()
	cfg.TokenProvider = NewStaticTokenProvider("from-cfg-provider")
	got := capturesAuthToken(t, cfg)
	if got != "Bearer from-cfg-provider" {
		t.Errorf("expected cfg.TokenProvider to be used, got %q", got)
	}
}

func TestNewClient_ConfigTokenUsedWhenArgAndProviderNil(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Token = "from-cfg-token"
	got := capturesAuthToken(t, cfg)
	if got != "Bearer from-cfg-token" {
		t.Errorf("expected cfg.Token to be used, got %q", got)
	}
}

func TestNewClient_AllNilEmptyErrors(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Token = ""
	cfg.TokenProvider = nil
	_, err := NewClient(cfg, nil)
	if err == nil {
		t.Fatal("expected error when arg, cfg.TokenProvider and cfg.Token are all nil/empty")
	}
}
