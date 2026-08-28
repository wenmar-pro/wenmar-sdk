package auth

import (
	"context"
	"net/http"
	"testing"
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
