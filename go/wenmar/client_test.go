package wenmar

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ctx = context.Background()

func TestNewClient_SetsBaseURL(t *testing.T) {
	c, err := NewClient("https://api.example.com", "test-token")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.BaseURL != "https://api.example.com" {
		t.Errorf("expected BaseURL 'https://api.example.com', got '%s'", c.BaseURL)
	}
	if c.Token != "test-token" {
		t.Errorf("expected Token 'test-token', got '%s'", c.Token)
	}
}

func TestNewClient_EmptyToken(t *testing.T) {
	_, err := NewClient("https://api.example.com", "")
	if err == nil {
		t.Error("expected error for empty token")
	}
}

func TestClient_AuthHeader(t *testing.T) {
	var capturedAuth string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":[]}`))
	}))
	defer ts.Close()

	c, _ := NewClient(ts.URL, "my-token")
	_, err := c.ListCustomers(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedAuth != "Bearer my-token" {
		t.Errorf("expected 'Bearer my-token', got '%s'", capturedAuth)
	}
}
