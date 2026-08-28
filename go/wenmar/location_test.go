package wenmar

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForLocation_HeaderInjection(t *testing.T) {
	var capturedLocation string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/locations/1" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":1,"name":"Bay 1"}`))
			return
		}
		capturedLocation = r.Header.Get("X-Wenmar-Location")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	lc, err := c.ForLocation(context.Background(), "1")
	if err != nil {
		t.Fatalf("ForLocation failed: %v", err)
	}
	if lc.LocationID() != "1" {
		t.Errorf("expected location '1', got %q", lc.LocationID())
	}

	if _, err := lc.ListCustomers(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedLocation != "1" {
		t.Errorf("expected X-Wenmar-Location '1', got %q", capturedLocation)
	}
}

func TestForLocation_EmptyID(t *testing.T) {
	c := newTestClient(t, "https://localhost", "test-token")
	if _, err := c.ForLocation(context.Background(), ""); err == nil {
		t.Error("expected error for empty location ID")
	}
}

func TestForLocation_VerifiesAccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":{"code":"not_found","message":"Location not found","details":{}}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	if _, err := c.ForLocation(context.Background(), "loc_missing"); err == nil {
		t.Error("expected error when location access fails")
	}
}
