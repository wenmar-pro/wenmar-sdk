package wenmar

import (
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.BaseURL != "https://app.wenmarpro.com" {
		t.Errorf("expected default BaseURL, got %q", cfg.BaseURL)
	}
	if cfg.Timeout != 30*time.Second {
		t.Errorf("expected 30s timeout, got %v", cfg.Timeout)
	}
	if cfg.MaxRetries != 3 {
		t.Errorf("expected 3 max retries, got %d", cfg.MaxRetries)
	}
	if !cfg.CacheEnabled {
		t.Error("expected CacheEnabled=true by default")
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	t.Setenv("WENMAR_BASE_URL", "https://staging.wenmarpro.com")
	t.Setenv("WENMAR_TIMEOUT", "10s")
	t.Setenv("WENMAR_MAX_RETRIES", "5")
	t.Setenv("WENMAR_CACHE", "false")

	cfg := LoadConfigFromEnv()
	if cfg.BaseURL != "https://staging.wenmarpro.com" {
		t.Errorf("expected staging BaseURL, got %q", cfg.BaseURL)
	}
	if cfg.Timeout != 10*time.Second {
		t.Errorf("expected 10s timeout, got %v", cfg.Timeout)
	}
	if cfg.MaxRetries != 5 {
		t.Errorf("expected 5 retries, got %d", cfg.MaxRetries)
	}
	if cfg.CacheEnabled {
		t.Error("expected CacheEnabled=false")
	}
}

func TestConfig_DeepCopy(t *testing.T) {
	original := DefaultConfig()
	copy := original
	copy.BaseURL = "https://other.example.com"
	if original.BaseURL == copy.BaseURL {
		t.Error("deep copy failed: mutating copy affected original")
	}
}
