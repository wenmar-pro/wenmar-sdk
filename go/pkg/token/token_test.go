package token

import (
	"context"
	"testing"
)

func TestStaticTokenProvider(t *testing.T) {
	p := NewStatic("my-token")
	tok, err := p.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok != "my-token" {
		t.Errorf("expected 'my-token', got %q", tok)
	}
}

func TestStaticTokenProvider_Empty(t *testing.T) {
	p := NewStatic("")
	if _, err := p.Token(context.Background()); err == nil {
		t.Error("expected error for empty token")
	}
}
