package wenmar

import (
	"context"
	"testing"
)

func TestStaticTokenProvider(t *testing.T) {
	tp := NewStaticTokenProvider("my-token")
	tok, err := tp.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok != "my-token" {
		t.Errorf("expected 'my-token', got %q", tok)
	}
}

func TestStaticTokenProvider_Empty(t *testing.T) {
	tp := NewStaticTokenProvider("")
	_, err := tp.Token(context.Background())
	if err == nil {
		t.Error("expected error for empty token")
	}
}
