package wenmar

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownloadResult_StreamsBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=\"customers.csv\"")
		w.WriteHeader(200)
		w.Write([]byte("id,name\n1,Jane\n2,John\n"))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test")
	result, err := c.DownloadCustomersExport(context.Background(), 1)
	if err != nil {
		t.Fatalf("DownloadCustomersExport failed: %v", err)
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("reading body failed: %v", err)
	}
	expected := "id,name\n1,Jane\n2,John\n"
	if string(body) != expected {
		t.Errorf("expected %q, got %q", expected, string(body))
	}
	if result.ContentType != "text/csv" {
		t.Errorf("expected text/csv, got %s", result.ContentType)
	}
	if result.Filename != "customers.csv" {
		t.Errorf("expected customers.csv, got %s", result.Filename)
	}
}

func TestDownloadResult_ErrorClosesBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.WriteHeader(404)
		w.Write([]byte(`{"error":{"code":"not_found","message":"x","field_errors":{}}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test")
	_, err := c.DownloadCustomersExport(context.Background(), 999)
	if err == nil {
		t.Fatal("expected an error for 404 download")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Code != "not_found" {
		t.Errorf("expected code not_found, got %s", apiErr.Code)
	}
}
