package wenmar

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestConditionalGet_ReturnsCachedBodyOn304(t *testing.T) {
	var calls int32
	var lastIfNoneMatch string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastIfNoneMatch = r.Header.Get("If-None-Match")
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("ETag", `"abc123"`)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":1,"full_name":"Jane"}`))
			return
		}
		if lastIfNoneMatch == `"abc123"` {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":1,"full_name":"Jane"}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")

	resp, err := c.ShowCustomer(context.Background(), 1)
	if err != nil {
		t.Fatalf("first call failed: %v", err)
	}
	if resp.JSON200 == nil || resp.JSON200.FullName != "Jane" {
		t.Fatalf("expected parsed data on first call, got %+v", resp.JSON200)
	}

	resp2, err := c.ShowCustomer(context.Background(), 1)
	if err != nil {
		t.Fatalf("second call failed: %v", err)
	}
	if resp2.JSON200 == nil || resp2.JSON200.FullName != "Jane" {
		t.Fatalf("expected cached body parsed on 304, got %+v", resp2.JSON200)
	}

	if lastIfNoneMatch != `"abc123"` {
		t.Errorf("expected If-None-Match header set, got %q", lastIfNoneMatch)
	}
	if calls != 2 {
		t.Errorf("expected 2 HTTP calls (200 + 304), got %d", calls)
	}
}

func containsStr(s, substr string) bool {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
