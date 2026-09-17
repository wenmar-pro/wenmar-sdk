package auth

import "testing"

func TestLoadConfigFromEnvURLPrecedence(t *testing.T) {
	t.Setenv("WENMAR_TOKEN", "tok")
	t.Setenv("WENMAR_LOCATION_ID", "42")

	t.Run("WENMAR_URL wins over WENMAR_BASE_URL", func(t *testing.T) {
		t.Setenv("WENMAR_URL", "https://url.example")
		t.Setenv("WENMAR_BASE_URL", "https://base.example")
		cfg := LoadConfigFromEnv()
		if cfg.BaseURL != "https://url.example" {
			t.Fatalf("expected WENMAR_URL to win, got %q", cfg.BaseURL)
		}
	})

	t.Run("WENMAR_BASE_URL used as fallback", func(t *testing.T) {
		t.Setenv("WENMAR_URL", "")
		t.Setenv("WENMAR_BASE_URL", "https://base.example")
		cfg := LoadConfigFromEnv()
		if cfg.BaseURL != "https://base.example" {
			t.Fatalf("expected WENMAR_BASE_URL fallback, got %q", cfg.BaseURL)
		}
	})

	t.Run("defaults when neither set", func(t *testing.T) {
		t.Setenv("WENMAR_URL", "")
		t.Setenv("WENMAR_BASE_URL", "")
		cfg := LoadConfigFromEnv()
		if cfg.BaseURL != DefaultConfig().BaseURL {
			t.Fatalf("expected default base URL, got %q", cfg.BaseURL)
		}
	})
}
