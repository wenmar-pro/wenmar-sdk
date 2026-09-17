package wenmar

import (
	"context"
	"fmt"
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

func TestConditionalGet_CacheKeyIsLocationAware(t *testing.T) {
	var calls int32
	var lastLoc string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastLoc = r.Header.Get("X-Wenmar-Location")
		n := atomic.AddInt32(&calls, 1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("ETag", `"etag-loc"`)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"full_name":"Jane"}]`))
		_ = n
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")

	// Two scoped clients must not share cache entries across locations.
	scopedA, err := c.ForLocation("1")
	if err != nil {
		t.Fatalf("scoped A creation failed: %v", err)
	}
	scopedB, err := c.ForLocation("2")
	if err != nil {
		t.Fatalf("scoped B creation failed: %v", err)
	}

	if _, err := scopedA.ListCustomers(context.Background(), nil); err != nil {
		t.Fatalf("scoped A call failed: %v", err)
	}
	if lastLoc != "1" {
		t.Errorf("expected location header '1', got %q", lastLoc)
	}
	if _, err := scopedB.ListCustomers(context.Background(), nil); err != nil {
		t.Fatalf("scoped B call failed: %v", err)
	}
	if lastLoc != "2" {
		t.Errorf("expected location header '2', got %q", lastLoc)
	}
	// Each distinct location produced its own cache miss -> 2 outbound calls.
	if calls != 2 {
		t.Errorf("expected 2 calls (one per location), got %d", calls)
	}
}

func TestConditionalGet_304PreservesPaginationHeaders(t *testing.T) {
	var serverURL string
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("ETag", `"list-etag"`)
			w.Header().Set("Link", fmt.Sprintf(`<%s/customers?page=2>; rel="next"`, serverURL))
			w.Header().Set("X-Total-Count", "42")
			w.Header().Set("X-Per-Page", "25")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
			return
		}
		w.WriteHeader(http.StatusNotModified)
	}))
	defer ts.Close()
	serverURL = ts.URL

	c := newTestClient(t, ts.URL, "test-token")

	// First call caches the list (with its pagination headers).
	resp1, err := c.ListCustomersRaw(context.Background(), nil)
	if err != nil {
		t.Fatalf("first call failed: %v", err)
	}
	if resp1.HTTPResponse.Header.Get("Link") == "" {
		t.Fatal("expected Link header on first response")
	}

	// Second call hits the 304 and should return the cached body AND the
	// cached headers, so a PaginatorFromResponse client keeps paginating.
	resp2, err := c.ListCustomersRaw(context.Background(), nil)
	if err != nil {
		t.Fatalf("second call failed: %v", err)
	}
	if resp2.JSON200 == nil || len(*resp2.JSON200) != 1 {
		t.Fatalf("expected cached list parsed on 304, got %+v", resp2.JSON200)
	}

	h := resp2.HTTPResponse.Header
	if got := h.Get("Link"); got == "" || got != resp1.HTTPResponse.Header.Get("Link") {
		t.Errorf("expected Link header preserved on 304, got %q", got)
	}
	if got := h.Get("X-Total-Count"); got != "42" {
		t.Errorf("expected X-Total-Count preserved on 304, got %q", got)
	}
	if got := h.Get("X-Per-Page"); got != "25" {
		t.Errorf("expected X-Per-Page preserved on 304, got %q", got)
	}

	if calls != 2 {
		t.Errorf("expected 2 HTTP calls (200 + 304), got %d", calls)
	}

	// Driving a Paginator from the 304 response must still see the next page.
	paginator := c.PaginatorFromResponse(resp2.HTTPResponse)
	if !paginator.HasNext() {
		t.Error("expected PaginatorFromResponse on the 304 to have a next page (Link preserved)")
	}
}

func TestConditionalGet_CachesLastModifiedOnly(t *testing.T) {
	var calls int32
	var lastIMS string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastIMS = r.Header.Get("If-Modified-Since")
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Last-Modified", "Wed, 21 Oct 2015 07:28:00 GMT")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
			return
		}
		if lastIMS == "Wed, 21 Oct 2015 07:28:00 GMT" {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")

	if _, err := c.ListCustomersRaw(context.Background(), nil); err != nil {
		t.Fatalf("first call failed: %v", err)
	}
	if lastIMS != "" {
		t.Errorf("expected no If-Modified-Since on first call, got %q", lastIMS)
	}

	resp2, err := c.ListCustomersRaw(context.Background(), nil)
	if err != nil {
		t.Fatalf("second call failed: %v", err)
	}
	if lastIMS != "Wed, 21 Oct 2015 07:28:00 GMT" {
		t.Errorf("expected If-Modified-Since set from cached Last-Modified, got %q", lastIMS)
	}
	// A Last-Modified-only 200 should be cached, so the repeat returns the
	// cached body on 304 rather than a fresh (possibly empty) 200.
	if resp2.JSON200 == nil || len(*resp2.JSON200) != 1 {
		t.Errorf("expected cached body on 304 for Last-Modified-only response, got %+v", resp2.JSON200)
	}
	if calls != 2 {
		t.Errorf("expected 2 HTTP calls (200 + 304), got %d", calls)
	}
}
