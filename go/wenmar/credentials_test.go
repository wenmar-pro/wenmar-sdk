package wenmar

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestFileCredentialStore_SaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "credentials.json")
	store := NewFileCredentialStore(path)

	err := store.Save("my-api-token")
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	info, _ := os.Stat(path)
	if info != nil && info.Mode().Perm() != 0600 {
		t.Errorf("expected file mode 0600, got %v", info.Mode().Perm())
	}

	tok, err := store.Load(context.Background())
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if tok != "my-api-token" {
		t.Errorf("expected 'my-api-token', got %q", tok)
	}
}

func TestFileCredentialStore_NotFound(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nonexistent", "credentials.json")
	store := NewFileCredentialStore(path)

	_, err := store.Load(context.Background())
	if err == nil {
		t.Error("expected error when credentials file does not exist")
	}
}
