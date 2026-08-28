package wenmar

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zalando/go-keyring"
)

// CredentialStore persists the API token securely. It first tries the
// system keyring (macOS Keychain, Windows Credential Manager, Linux
// Secret Service). If the keyring is unavailable (headless Linux, CI),
// it falls back to a file at ~/.config/wenmar/credentials.json with
// 0600 permissions.
type CredentialStore struct {
	keyringService string
	keyringUser    string
	filePath       string
}

// NewCredentialStore creates a store with the default keyring service name
// and a file path at ~/.config/wenmar/credentials.json.
func NewCredentialStore() *CredentialStore {
	home, _ := os.UserHomeDir()
	return &CredentialStore{
		keyringService: "wenmar-sdk",
		keyringUser:    "default",
		filePath:       filepath.Join(home, ".config", "wenmar", "credentials.json"),
	}
}

// Save stores the token in the keyring (if available) and the file fallback.
func (s *CredentialStore) Save(token string) error {
	// Try keyring first; ignore errors (headless/unsupported).
	_ = keyring.Set(s.keyringService, s.keyringUser, token)

	// Always write the file fallback.
	if err := os.MkdirAll(filepath.Dir(s.filePath), 0700); err != nil {
		return fmt.Errorf("create credential dir: %w", err)
	}
	data, _ := json.Marshal(map[string]string{"token": token})
	return os.WriteFile(s.filePath, data, 0600)
}

// Token implements TokenProvider. It reads from the keyring first, then the
// file fallback.
func (s *CredentialStore) Token(_ context.Context) (string, error) {
	if tok, err := keyring.Get(s.keyringService, s.keyringUser); err == nil && tok != "" {
		return tok, nil
	}
	return s.readFromFile()
}

func (s *CredentialStore) readFromFile() (string, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return "", fmt.Errorf("read credentials file: %w", err)
	}
	var creds struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(data, &creds); err != nil {
		return "", fmt.Errorf("parse credentials file: %w", err)
	}
	if creds.Token == "" {
		return "", fmt.Errorf("credentials file has empty token")
	}
	return creds.Token, nil
}

// FileCredentialStore is a file-only credential store for testing and
// environments without a keyring.
type FileCredentialStore struct {
	path string
}

func NewFileCredentialStore(path string) *FileCredentialStore {
	return &FileCredentialStore{path: path}
}

func (s *FileCredentialStore) Save(token string) error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0700); err != nil {
		return err
	}
	data, _ := json.Marshal(map[string]string{"token": token})
	return os.WriteFile(s.path, data, 0600)
}

// Load reads the token from the file. It is an alias for Token for callers
// that prefer an explicit Load name.
func (s *FileCredentialStore) Load(ctx context.Context) (string, error) {
	return s.Token(ctx)
}

func (s *FileCredentialStore) Token(_ context.Context) (string, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return "", err
	}
	var creds struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(data, &creds); err != nil {
		return "", err
	}
	if creds.Token == "" {
		return "", fmt.Errorf("empty token in credentials file")
	}
	return creds.Token, nil
}
