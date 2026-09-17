package wenmar

import (
	"net/http"
	"time"
)

// Config holds all SDK configuration. Pass it to NewClient. The struct is
// deep-copied by NewClient so callers cannot mutate the client's config
// after construction.
type Config struct {
	// BaseURL is the API root, e.g. "https://app.wenmarpro.com".
	BaseURL string
	// Timeout for HTTP requests.
	Timeout time.Duration
	// MaxRetries is the maximum number of retry attempts (excluding the
	// initial request). 0 disables retry.
	MaxRetries int
	// CacheEnabled controls ETag/Last-Modified conditional GET caching.
	CacheEnabled bool
	// HTTPClient is the underlying http.Client. If nil, a new one is
	// built with the Timeout and transport stack (retry + caching).
	HTTPClient *http.Client
	// Token is a static bearer token. Prefer TokenProvider for rotation.
	Token string
	// TokenProvider resolves a token per request. Takes precedence over Token.
	TokenProvider TokenProvider
	// Hooks receives observability callbacks. Defaults to NoopHooks.
	Hooks Hooks
}

// DefaultConfig returns a Config with production defaults.
func DefaultConfig() Config {
	return Config{
		BaseURL:      "https://app.wenmarpro.com",
		Timeout:      30 * time.Second,
		MaxRetries:   3,
		CacheEnabled: true,
	}
}
