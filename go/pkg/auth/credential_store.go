package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zalando/go-keyring"
)

// CredentialStore persists the API token securely. Implementations may use
// the system keyring, a file, or an in-memory store for tests.
type CredentialStore interface {
	GetToken(ctx context.Context) (*Token, error)
	SaveToken(ctx context.Context, token *Token) error
	DeleteToken(ctx context.Context) error
}

const (
	keyringService = "wenmar-cli"
	keyringUser    = "token"
)

// KeyringStore uses the system keyring (macOS Keychain, Windows Credential
// Manager, Linux Secret Service) via zalando/go-keyring.
type KeyringStore struct{}

// GetToken reads the token from the keyring.
func (KeyringStore) GetToken(_ context.Context) (*Token, error) {
	raw, err := keyring.Get(keyringService, keyringUser)
	if err != nil {
		return nil, err
	}
	if raw == "" {
		return nil, fmt.Errorf("keyring returned an empty token")
	}
	return &Token{AccessToken: raw}, nil
}

// SaveToken writes the token to the keyring.
func (KeyringStore) SaveToken(_ context.Context, token *Token) error {
	if token == nil || token.AccessToken == "" {
		return fmt.Errorf("cannot save an empty token")
	}
	return keyring.Set(keyringService, keyringUser, token.AccessToken)
}

// DeleteToken removes the token from the keyring.
func (KeyringStore) DeleteToken(_ context.Context) error {
	return keyring.Delete(keyringService, keyringUser)
}

// FileStore persists the token to a JSON file with 0600 permissions. It is
// the fallback for headless Linux / CI where the keyring is unavailable.
type FileStore struct {
	Path string
}

// GetToken reads the token from the file.
func (s FileStore) GetToken(_ context.Context) (*Token, error) {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, err
	}
	var token Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("parse credentials file: %w", err)
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("credentials file has an empty token")
	}
	return &token, nil
}

// SaveToken writes the token to the file with 0600 permissions.
func (s FileStore) SaveToken(_ context.Context, token *Token) error {
	if token == nil || token.AccessToken == "" {
		return fmt.Errorf("cannot save an empty token")
	}
	if err := os.MkdirAll(filepath.Dir(s.Path), 0700); err != nil {
		return fmt.Errorf("create credential dir: %w", err)
	}
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0600)
}

// DeleteToken removes the credentials file.
func (s FileStore) DeleteToken(_ context.Context) error {
	if err := os.Remove(s.Path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// NewCredentialStore tries the keyring first and falls back to a file at
// ~/.config/wenmar/credentials.json. On headless Linux without D-Bus, the
// keyring returns ErrUnavailable, which triggers the file fallback.
func NewCredentialStore() CredentialStore {
	home, _ := os.UserHomeDir()
	return &fallbackStore{
		primary: KeyringStore{},
		fallback: FileStore{
			Path: filepath.Join(home, ".config", "wenmar", "credentials.json"),
		},
	}
}

// fallbackStore tries the primary (keyring) store and falls back to the file
// store when the keyring is unavailable.
type fallbackStore struct {
	primary  CredentialStore
	fallback CredentialStore
}

func (s *fallbackStore) GetToken(ctx context.Context) (*Token, error) {
	tok, err := s.primary.GetToken(ctx)
	if err == nil {
		return tok, nil
	}
	return s.fallback.GetToken(ctx)
}

func (s *fallbackStore) SaveToken(ctx context.Context, token *Token) error {
	// Always write the file fallback so headless environments can read it.
	if err := s.fallback.SaveToken(ctx, token); err != nil {
		return err
	}
	// Best-effort keyring write; ignore errors (headless/unsupported).
	_ = s.primary.SaveToken(ctx, token)
	return nil
}

func (s *fallbackStore) DeleteToken(ctx context.Context) error {
	_ = s.primary.DeleteToken(ctx)
	return s.fallback.DeleteToken(ctx)
}
