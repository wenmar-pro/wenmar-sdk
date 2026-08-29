package wenmar

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
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
	_, err := c.ListCustomers(ctx)
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
	_, err := c.ListCustomers(ctx)
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
	_, err := c.CreateCustomer(ctx, generated.CreateCustomerJSONRequestBody{})
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
	_, err := c.CreateCustomer(ctx, generated.CreateCustomerJSONRequestBody{})
	if err != nil {
		t.Fatalf("expected success after 429 retry, got: %v", err)
	}
	if calls != 2 {
		t.Errorf("expected 2 calls (1 throttle + 1 success), got %d", calls)
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
	_, err := c.ListCustomers(ctx)
	if err != nil {
		t.Fatalf("expected success after HTTP-date Retry-After, got: %v", err)
	}
	if calls != 2 {
		t.Errorf("expected 2 calls, got %d", calls)
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
	_, err := c.ListCustomers(ctx)
	if err == nil {
		t.Fatal("expected error on 404")
	}
	if calls != 1 {
		t.Errorf("expected 1 call (no retry on 4xx), got %d", calls)
	}
}
