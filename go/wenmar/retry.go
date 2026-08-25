package wenmar

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

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
		if err != nil || resp.StatusCode < 500 {
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
