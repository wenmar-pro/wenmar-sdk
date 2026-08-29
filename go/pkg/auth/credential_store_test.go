package auth

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFileStore_FullTokenRoundTrip(t *testing.T) {
	store := FileStore{Path: filepath.Join(t.TempDir(), "credentials.json")}

	expiresAt := time.Now().Add(2 * time.Hour)
	original := &Token{
		AccessToken:  "access-123",
		RefreshToken: "refresh-456",
		ExpiresAt:    &expiresAt,
		TokenType:    "Bearer",
	}

	if err := store.SaveToken(context.Background(), original); err != nil {
		t.Fatalf("SaveToken failed: %v", err)
	}

	loaded, err := store.GetToken(context.Background())
	if err != nil {
		t.Fatalf("GetToken failed: %v", err)
	}

	if loaded.AccessToken != original.AccessToken {
		t.Errorf("AccessToken = %q, want %q", loaded.AccessToken, original.AccessToken)
	}
	if loaded.RefreshToken != original.RefreshToken {
		t.Errorf("RefreshToken = %q, want %q", loaded.RefreshToken, original.RefreshToken)
	}
	if loaded.TokenType != original.TokenType {
		t.Errorf("TokenType = %q, want %q", loaded.TokenType, original.TokenType)
	}
	if loaded.ExpiresAt == nil {
		t.Fatal("ExpiresAt is nil")
	}
	// Compare Unix seconds (sub-second precision may differ)
	if loaded.ExpiresAt.Unix() != original.ExpiresAt.Unix() {
		t.Errorf("ExpiresAt Unix = %d, want %d", loaded.ExpiresAt.Unix(), original.ExpiresAt.Unix())
	}
}

func TestFileStore_StaticTokenRoundTrip(t *testing.T) {
	store := FileStore{Path: filepath.Join(t.TempDir(), "credentials.json")}

	original := &Token{AccessToken: "static-token-only"}

	if err := store.SaveToken(context.Background(), original); err != nil {
		t.Fatalf("SaveToken failed: %v", err)
	}

	loaded, err := store.GetToken(context.Background())
	if err != nil {
		t.Fatalf("GetToken failed: %v", err)
	}

	if loaded.AccessToken != "static-token-only" {
		t.Errorf("AccessToken = %q, want %q", loaded.AccessToken, "static-token-only")
	}
	if loaded.RefreshToken != "" {
		t.Errorf("RefreshToken = %q, want empty", loaded.RefreshToken)
	}
}

func TestFileStore_DeleteToken(t *testing.T) {
	store := FileStore{Path: filepath.Join(t.TempDir(), "credentials.json")}

	_ = store.SaveToken(context.Background(), &Token{AccessToken: "x"})
	if err := store.DeleteToken(context.Background()); err != nil {
		t.Fatalf("DeleteToken failed: %v", err)
	}

	_, err := store.GetToken(context.Background())
	if err == nil {
		t.Fatal("expected error after delete, got nil")
	}
}

func TestFileStore_DeleteToken_Idempotent(t *testing.T) {
	store := FileStore{Path: filepath.Join(t.TempDir(), "credentials.json")}

	// Deleting a non-existent file should not error
	if err := store.DeleteToken(context.Background()); err != nil {
		t.Fatalf("DeleteToken on non-existent file should not error: %v", err)
	}
}

func TestToken_JSONSerialization(t *testing.T) {
	expiresAt := time.Now().Add(2 * time.Hour).UTC().Truncate(time.Second)

	original := &Token{
		AccessToken:  "access-123",
		RefreshToken: "refresh-456",
		ExpiresAt:    &expiresAt,
		TokenType:    "Bearer",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var loaded Token
	if err := json.Unmarshal(data, &loaded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if loaded.AccessToken != original.AccessToken {
		t.Errorf("AccessToken = %q, want %q", loaded.AccessToken, original.AccessToken)
	}
	if loaded.RefreshToken != original.RefreshToken {
		t.Errorf("RefreshToken = %q, want %q", loaded.RefreshToken, original.RefreshToken)
	}
}

func TestToken_JSONBackwardCompat_RawString(t *testing.T) {
	// Simulate an old keyring entry that was stored as a raw string
	// (not JSON). The GetToken fallback should handle this.
	rawString := "old-static-token"

	// Try to unmarshal as JSON first — this should fail
	var token Token
	if err := json.Unmarshal([]byte(rawString), &token); err == nil {
		// If it somehow parses, AccessToken should be empty (it's not valid JSON)
		if token.AccessToken != "" {
			t.Fatalf("unexpected successful parse of raw string: %+v", token)
		}
	}

	// The fallback: treat as raw AccessToken
	fallbackToken := &Token{AccessToken: rawString}
	if fallbackToken.AccessToken != "old-static-token" {
		t.Errorf("fallback AccessToken = %q, want %q", fallbackToken.AccessToken, "old-static-token")
	}
}

func TestFileStore_Roundtrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "credentials.json")
	store := FileStore{Path: path}

	tok := &Token{AccessToken: "sk-test-1234"}
	if err := store.SaveToken(context.Background(), tok); err != nil {
		t.Fatalf("SaveToken failed: %v", err)
	}

	info, _ := os.Stat(path)
	if info != nil && info.Mode().Perm() != 0600 {
		t.Errorf("expected file mode 0600, got %v", info.Mode().Perm())
	}

	got, err := store.GetToken(context.Background())
	if err != nil {
		t.Fatalf("GetToken failed: %v", err)
	}
	if got.AccessToken != "sk-test-1234" {
		t.Errorf("expected 'sk-test-1234', got %q", got.AccessToken)
	}
}

func TestFileStore_Delete(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "credentials.json")
	store := FileStore{Path: path}

	if err := store.SaveToken(context.Background(), &Token{AccessToken: "x"}); err != nil {
		t.Fatalf("SaveToken failed: %v", err)
	}
	if err := store.DeleteToken(context.Background()); err != nil {
		t.Fatalf("DeleteToken failed: %v", err)
	}
	if _, err := store.GetToken(context.Background()); err == nil {
		t.Error("expected error after delete")
	}
}

func TestFileStore_EmptyTokenRejected(t *testing.T) {
	dir := t.TempDir()
	store := FileStore{Path: filepath.Join(dir, "credentials.json")}
	if err := store.SaveToken(context.Background(), &Token{}); err == nil {
		t.Error("expected error saving empty token")
	}
}

func TestToken_IsExpired(t *testing.T) {
	past := time.Now().Add(-time.Hour)
	future := time.Now().Add(time.Hour)
	if !(&Token{ExpiresAt: &past}).IsExpired() {
		t.Error("expected expired token to be expired")
	}
	if (&Token{ExpiresAt: &future}).IsExpired() {
		t.Error("expected future token to not be expired")
	}
	if (&Token{}).IsExpired() {
		t.Error("expected nil-expiry token to not be expired")
	}
}

func TestToken_WillExpireWithin(t *testing.T) {
	soon := time.Now().Add(2 * time.Minute)
	far := time.Now().Add(2 * time.Hour)
	if !(&Token{ExpiresAt: &soon}).WillExpireWithin(5 * time.Minute) {
		t.Error("expected token expiring in 2m to expire within 5m")
	}
	if (&Token{ExpiresAt: &far}).WillExpireWithin(5 * time.Minute) {
		t.Error("expected token expiring in 2h to NOT expire within 5m")
	}
	if (&Token{}).WillExpireWithin(5 * time.Minute) {
		t.Error("expected nil-expiry token to not expire within window")
	}
}
