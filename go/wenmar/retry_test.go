package wenmar

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestRetry_On500Then200(t *testing.T) {
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":{"code":"internal_error","message":"fail","details":{}}}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.ListCustomers(ctx, nil)
	if err != nil {
		t.Fatalf("expected success after retry, got error: %v", err)
	}
	if calls != 3 {
		t.Errorf("expected 3 calls (2 failures + 1 success), got %d", calls)
	}
}

func TestRetry_MaxRetriesExceeded(t *testing.T) {
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":{"code":"internal_error","message":"fail","details":{}}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.ListCustomers(ctx, nil)
	if err == nil {
		t.Fatal("expected error after max retries")
	}
	if calls != 4 {
		t.Errorf("expected 4 calls (initial + 3 retries), got %d", calls)
	}
}

func TestRetry_PostOn500NotRetried(t *testing.T) {
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":{"code":"internal_error","message":"fail","details":{}}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.CreateCustomer(ctx, CreateCustomerRequest{})
	if err == nil {
		t.Fatal("expected error on POST 500")
	}
	if calls != 1 {
		t.Errorf("expected 1 call (POST must not retry on 500), got %d", calls)
	}
}

func TestRetry_PostOn429Retried(t *testing.T) {
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error":{"code":"rate_limited","message":"slow","details":{}}}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"customer":{"id":1}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.CreateCustomer(ctx, CreateCustomerRequest{})
	if err != nil {
		t.Fatalf("expected success after 429 retry, got: %v", err)
	}
	if calls != 2 {
		t.Errorf("expected 2 calls (1 throttle + 1 success), got %d", calls)
	}
}

func TestRetry_POST429RewindsBody(t *testing.T) {
	var calls int32
	var bodies []string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(strings.Builder)
		io.Copy(buf, r.Body)
		bodies = append(bodies, buf.String())
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error":{"code":"rate_limited","message":"slow"}}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"customer":{"id":1}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	req := CreateCustomerRequest{}
	req.Customer.FirstName = "Jane"
	req.Customer.LastName = "Doe"
	if _, err := c.CreateCustomer(ctx, req); err != nil {
		t.Fatalf("expected success after 429 retry, got: %v", err)
	}
	if calls != 2 {
		t.Fatalf("expected 2 calls, got %d", calls)
	}
	if len(bodies) != 2 || bodies[0] != bodies[1] {
		t.Errorf("expected the request body to be rewound on retry, got %v", bodies)
	}
	if !strings.Contains(bodies[0], "Jane") {
		t.Errorf("expected body to contain the customer first_name, got %q", bodies[0])
	}
}

func TestRetry_RetryAfterHTTPDate(t *testing.T) {
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.Header().Set("Retry-After", "Wed, 21 Oct 2015 07:28:00 GMT")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error":{"code":"rate_limited","message":"slow","details":{}}}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.ListCustomers(ctx, nil)
	if err != nil {
		t.Fatalf("expected success after HTTP-date Retry-After, got: %v", err)
	}
	if calls != 2 {
		t.Errorf("expected 2 calls, got %d", calls)
	}
}

func TestRetry_DrainedBodyNotRetried(t *testing.T) {
	// A request with a non-nil Body but nil GetBody cannot be safely replayed:
	// a retry would send a drained/empty body. The retry transport must refuse
	// to retry and call the inner transport exactly once.
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error":{"code":"rate_limited","message":"slow"}}`))
	}))
	defer ts.Close()

	inner := &callCountingTransport{next: http.DefaultTransport}
	rt := newRetryTransportWithRetries(3, inner)
	client := &http.Client{Transport: rt}

	req, err := http.NewRequest(http.MethodPost, ts.URL, strings.NewReader("payload"))
	if err != nil {
		t.Fatalf("NewRequest failed: %v", err)
	}
	// Clear GetBody so the body is not replayable (a real caller that sets a
	// plain strings.Reader body gets a nil GetBody in Go's http.NewRequest).
	req.GetBody = nil

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("expected 429 returned without retry, got %d", resp.StatusCode)
	}
	if calls != 1 {
		t.Errorf("expected exactly 1 call (drained body must not retry), got %d", calls)
	}
}

func TestRetry_NoRetryOn4xx(t *testing.T) {
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":{"code":"not_found","message":"not found","details":{}}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.ListCustomers(ctx, nil)
	if err == nil {
		t.Fatal("expected error on 404")
	}
	if calls != 1 {
		t.Errorf("expected 1 call (no retry on 4xx), got %d", calls)
	}
}

func TestRetry_BackoffCapsRetryAfter(t *testing.T) {
	rt := newRetryTransportWithRetries(3, http.DefaultTransport)
	resp := &http.Response{Header: http.Header{}}
	resp.Header.Set("Retry-After", "999999")
	got := rt.backoff(0, resp)
	if got != maxRetryAfterDelay {
		t.Errorf("expected Retry-After capped at %v, got %v", maxRetryAfterDelay, got)
	}
}

func TestRetry_CtxCancelDuringBackoffReturnsQuickly(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err := c.ListCustomers(ctx, nil)
	if err == nil {
		t.Fatal("expected context error, got nil")
	}
	if time.Since(start) > 2*time.Second {
		t.Errorf("cancellation honored too slowly: %v", time.Since(start))
	}
}
