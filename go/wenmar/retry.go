package wenmar

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Retry policy (derived from bc3's documented error handling):
//
//   429 Too Many Requests → read Retry-After header, wait, retry (max 3)
//   500/502/503/504      → exponential backoff with jitter, retry (max 3)
//   404 Not Found         → do NOT retry (deleted, inaccessible, or insufficient permissions)
//   304 Not Modified     → not an error; return cached body
//
// 404 is terminal because retrying a deleted/inaccessible resource will
// never succeed. 429 and 5xx are transient and may recover.
type retryTransport struct {
	transport   http.RoundTripper
	maxRetries  int
	baseDelay   time.Duration
}

func newRetryTransport() *retryTransport {
	return &retryTransport{
		transport:  http.DefaultTransport,
		maxRetries: 3,
		baseDelay:  500 * time.Millisecond,
	}
}

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt <= t.maxRetries; attempt++ {
		resp, err = t.transport.RoundTrip(req)
		if err != nil || !isRetryableStatus(resp.StatusCode) {
			return resp, err
		}

		if attempt < t.maxRetries {
			delay := t.backoff(attempt, resp)
			select {
			case <-req.Context().Done():
				return resp, req.Context().Err()
			case <-time.After(delay):
			}
			// Close the response body before retrying
			resp.Body.Close()
		}
	}

	return resp, err
}

func (t *retryTransport) backoff(attempt int, resp *http.Response) time.Duration {
	// Respect Retry-After header if present
	if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
		if seconds, err := strconv.Atoi(retryAfter); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}

	// Exponential backoff with jitter: delay = base * 2^attempt + jitter
	delay := float64(t.baseDelay) * math.Pow(2, float64(attempt))
	jitter := rand.Float64() * float64(t.baseDelay)
	return time.Duration(delay + jitter)
}

func isRetryableStatus(code int) bool {
	switch code {
	case http.StatusTooManyRequests, // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	default:
		return false
	}
}
