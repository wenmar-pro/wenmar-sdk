package wenmar

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
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

	c, _ := NewClient(ts.URL, "test-token")
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

	c, _ := NewClient(ts.URL, "test-token")
	_, err := c.ListCustomers(ctx)
	if err == nil {
		t.Fatal("expected error after max retries")
	}
	if calls != 4 {
		t.Errorf("expected 4 calls (initial + 3 retries), got %d", calls)
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

	c, _ := NewClient(ts.URL, "test-token")
	_, err := c.ListCustomers(ctx)
	if err == nil {
		t.Fatal("expected error on 404")
	}
	if calls != 1 {
		t.Errorf("expected 1 call (no retry on 4xx), got %d", calls)
	}
}
