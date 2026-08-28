package auth

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

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
