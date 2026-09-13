package wenmar

import (
	"net/http"
	"sync/atomic"
	"testing"
)

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

// callCountingTransport counts every RoundTrip call and delegates to next.
type callCountingTransport struct {
	next  http.RoundTripper
	calls int32
}

func (t *callCountingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	atomic.AddInt32(&t.calls, 1)
	return t.next.RoundTrip(req)
}
