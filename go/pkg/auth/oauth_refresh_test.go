package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRefreshToken_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/token" {
			t.Errorf("path = %q, want /oauth/token", r.URL.Path)
		}
		if err := r.ParseForm(); err != nil {
			t.Fatalf("ParseForm: %v", err)
		}
		if r.PostForm.Get("grant_type") != "refresh_token" {
			t.Errorf("grant_type = %q, want refresh_token", r.PostForm.Get("grant_type"))
		}
		if r.PostForm.Get("client_id") != "wenmar-cli" {
			t.Errorf("client_id = %q, want wenmar-cli", r.PostForm.Get("client_id"))
		}
		if r.PostForm.Get("refresh_token") != "old-refresh-token" {
			t.Errorf("refresh_token = %q, want old-refresh-token", r.PostForm.Get("refresh_token"))
		}

		json.NewEncoder(w).Encode(map[string]any{
			"access_token":  "new-access-token",
			"refresh_token": "new-refresh-token",
			"token_type":    "Bearer",
			"expires_in":    7200,
		})
	}))
	defer server.Close()

	token, err := RefreshToken(context.Background(), server.URL+"/oauth/token", "wenmar-cli", "old-refresh-token")
	if err != nil {
		t.Fatalf("RefreshToken failed: %v", err)
	}

	if token.AccessToken != "new-access-token" {
		t.Errorf("AccessToken = %q, want new-access-token", token.AccessToken)
	}
	if token.RefreshToken != "new-refresh-token" {
		t.Errorf("RefreshToken = %q, want new-refresh-token", token.RefreshToken)
	}
	if token.TokenType != "Bearer" {
		t.Errorf("TokenType = %q, want Bearer", token.TokenType)
	}
	if token.ExpiresAt == nil {
		t.Fatal("ExpiresAt is nil")
	}
	// ExpiresAt should be ~2 hours from now
	expectedExpiry := time.Now().Add(2 * time.Hour)
	if token.ExpiresAt.Before(expectedExpiry.Add(-1*time.Minute)) || token.ExpiresAt.After(expectedExpiry.Add(1*time.Minute)) {
		t.Errorf("ExpiresAt = %v, want ~%v", token.ExpiresAt, expectedExpiry)
	}
}

func TestRefreshToken_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid_grant",
		})
	}))
	defer server.Close()

	_, err := RefreshToken(context.Background(), server.URL+"/oauth/token", "wenmar-cli", "bad-token")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRefreshToken_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err := RefreshToken(context.Background(), server.URL+"/oauth/token", "wenmar-cli", "token")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRefreshToken_DoesNotFollowRedirects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth/token" {
			http.Redirect(w, r, "/evil", http.StatusFound)
			return
		}
		// If the client follows the redirect, it hits here
		t.Errorf("client followed redirect to %q — should not follow", r.URL.Path)
	}))
	defer server.Close()

	_, err := RefreshToken(context.Background(), server.URL+"/oauth/token", "wenmar-cli", "token")
	if err == nil {
		t.Fatal("expected error for redirect response, got nil")
	}
}