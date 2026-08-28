package auth

import "os"

// Config holds authentication-related configuration.
type Config struct {
	BaseURL    string `json:"base_url" yaml:"base_url"`
	Token      string `json:"token,omitempty" yaml:"token,omitempty"` // static token (legacy/file fallback)
	AuthMethod string `json:"auth_method" yaml:"auth_method"`         // "static" | "oauth"
	LocationID string `json:"location_id,omitempty" yaml:"location_id,omitempty"`
}

// DefaultConfig returns a Config with production defaults.
func DefaultConfig() Config {
	return Config{
		BaseURL:    "https://app.wenmarpro.com",
		AuthMethod: "static",
	}
}

// LoadConfigFromEnv reads WENMAR_TOKEN, WENMAR_URL, and WENMAR_LOCATION_ID
// environment variables.
func LoadConfigFromEnv() Config {
	cfg := DefaultConfig()
	if v := os.Getenv("WENMAR_TOKEN"); v != "" {
		cfg.Token = v
	}
	if v := os.Getenv("WENMAR_URL"); v != "" {
		cfg.BaseURL = v
	}
	if v := os.Getenv("WENMAR_LOCATION_ID"); v != "" {
		cfg.LocationID = v
	}
	return cfg
}
