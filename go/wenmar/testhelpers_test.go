package wenmar

import "testing"

func newTestClient(t *testing.T, baseURL, token string) *Client {
	t.Helper()
	cfg := DefaultConfig()
	cfg.BaseURL = baseURL
	c, err := NewClient(cfg, NewStaticTokenProvider(token))
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	return c
}
