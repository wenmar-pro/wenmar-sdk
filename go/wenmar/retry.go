package wenmar

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Retry policy:
//
//   429 Too Many Requests → retry all methods (idempotent throttle; honor Retry-After)
//   500/502/503/504      → retry GET only (exponential backoff with jitter; max 3)
//   507 Insufficient Storage → do NOT retry (account/plan limit; not transient)
//   404 Not Found         → do NOT retry (deleted, inaccessible, or forbidden)
//   304 Not Modified     → not an error; returned cached body by cachingTransport
//
// Mutations (POST/PATCH/DELETE) are NOT retried on 5xx because the server
// may have processed the request before the response was lost, and retrying
// would duplicate the side effect. 429 is safe to retry because the
// throttle response means the request was NOT processed.
type retryTransport struct {
	transport   http.RoundTripper
	maxRetries  int
	baseDelay   time.Duration
}

func newRetryTransportWithRetries(maxRetries int) *retryTransport {
	return &retryTransport{
		transport:  http.DefaultTransport,
		maxRetries: maxRetries,
		baseDelay:  500 * time.Millisecond,
	}
}

// Keep the old constructor for backwards compatibility within the package.
func newRetryTransport() *retryTransport {
	return newRetryTransportWithRetries(3)
}

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt <= t.maxRetries; attempt++ {
		resp, err = t.transport.RoundTrip(req)
		if err != nil || !isRetryableStatus(req.Method, resp.StatusCode) {
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
	if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
		if seconds, err := strconv.Atoi(retryAfter); err == nil {
			return time.Duration(seconds) * time.Second
		}
		if httpDate, err := http.ParseTime(retryAfter); err == nil {
			delay := time.Until(httpDate)
			if delay > 0 {
				return delay
			}
			return 0
		}
	}

	// Exponential backoff with jitter: delay = base * 2^attempt + jitter
	delay := float64(t.baseDelay) * math.Pow(2, float64(attempt))
	jitter := rand.Float64() * float64(t.baseDelay)
	return time.Duration(delay + jitter)
}

func isRetryableStatus(method string, code int) bool {
	// 429 (Too Many Requests) is retryable for all methods — it is an
	// idempotent throttle response that does not execute the request.
	if code == http.StatusTooManyRequests {
		return true
	}

	// 5xx server errors are only retried for safe (GET) methods.
	// Retrying POST/PATCH/DELETE on 5xx risks duplicating mutations
	// if the server processed the request but the response was lost.
	if method == http.MethodGet {
		switch code {
		case http.StatusInternalServerError,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout:
			return true
		}
	}

	return false
}
