package wenmar

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var ctx = context.Background()

func strPtr(s string) *string { return &s }

func intPtr(i int) *int { return &i }

func TestNewClient_WithConfigAndTokenProvider(t *testing.T) {
	cfg := DefaultConfig()
	cfg.BaseURL = "https://localhost" // will be overwritten by httptest below
	c, err := NewClient(cfg, NewStaticTokenProvider("test-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.BaseURL != "https://localhost" {
		t.Errorf("expected BaseURL 'https://localhost', got %q", c.BaseURL)
	}
}

func TestNewClient_SetsBaseURL(t *testing.T) {
	c, err := NewClient(DefaultConfig(), NewStaticTokenProvider("test-token"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.BaseURL != "https://app.wenmarpro.com" {
		t.Errorf("expected BaseURL 'https://app.wenmarpro.com', got '%s'", c.BaseURL)
	}
}

func TestNewClient_EmptyToken(t *testing.T) {
	c, err := NewClient(DefaultConfig(), NewStaticTokenProvider(""))
	if err != nil {
		t.Fatalf("expected construction to succeed (token resolved per request), got: %v", err)
	}
	// The empty token surfaces as a transport error on the first request.
	if _, err := c.ListCustomers(ctx, nil); err == nil {
		t.Error("expected error for empty token on request")
	}
}

func TestNewClient_RequiresTokenProvider(t *testing.T) {
	_, err := NewClient(Config{BaseURL: "https://app.wenmarpro.com"}, nil)
	if err == nil {
		t.Error("expected error when no token provider is given")
	}
}

func TestClient_AuthHeader(t *testing.T) {
	var capturedAuth string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "my-token")
	_, err := c.ListCustomers(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedAuth != "Bearer my-token" {
		t.Errorf("expected 'Bearer my-token', got '%s'", capturedAuth)
	}
}

func TestClient_ListCustomers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"full_name":"Jane Doe"}]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	resp, err := c.ListCustomers(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.JSON200 == nil || len(*resp.JSON200) != 1 {
		t.Fatalf("expected 1 customer, got %+v", resp.JSON200)
	}
	if (*resp.JSON200)[0].FullName != "Jane Doe" {
		t.Errorf("expected full_name 'Jane Doe', got %q", (*resp.JSON200)[0].FullName)
	}
}

func TestClient_CreateVehicle(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/vehicles" {
			t.Errorf("expected POST /vehicles, got %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id":1,"make":"Honda","model":"Civic","year":2020}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	req := CreateVehicleRequest{}
	req.Vehicle.Make = "Honda"
	req.Vehicle.Model = "Civic"
	req.Vehicle.Year = 2020

	resp, err := c.CreateVehicle(ctx, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.JSON201 == nil || resp.JSON201.Make != "Honda" {
		t.Errorf("expected vehicle make Honda, got %+v", resp.JSON201)
	}
}

func TestClient_UpdateVehicle(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch || r.URL.Path != "/vehicles/1" {
			t.Errorf("expected PATCH /vehicles/1, got %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":1,"make":"Toyota","model":"Camry","year":2020}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	req := UpdateVehicleRequest{}
	req.Vehicle.Make = "Toyota"

	resp, err := c.UpdateVehicle(ctx, 1, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.JSON200 == nil || resp.JSON200.Make != "Toyota" {
		t.Errorf("expected vehicle make Toyota, got %+v", resp.JSON200)
	}
}

func TestClient_TrashVehicle(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch || r.URL.Path != "/vehicles/1/trash" {
			t.Errorf("expected PATCH /vehicles/1/trash, got %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id":1,"status":"trashed"}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.TrashVehicle(ctx, 1, TrashVehicleRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClient_DecodeVin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vehicles/vin_decode" {
			t.Errorf("expected /vehicles/vin_decode, got %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("vin"); got != "1HGCM82633A004352" {
			t.Errorf("expected vin query param, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"make":"Honda","model":"Civic","vin":"1HGCM82633A004352"}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	resp, err := c.DecodeVin(ctx, &DecodeVinParams{Vin: strPtr("1HGCM82633A004352")})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.JSON200 == nil || resp.JSON200.Make != "Honda" {
		t.Errorf("expected decoded make Honda, got %+v", resp.JSON200)
	}
}

func TestClient_CheckDuplicate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/vehicles/check_duplicate" {
			t.Errorf("expected /vehicles/check_duplicate, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"matches":[{"id":1,"display_name":"Toyota Camry","url":"/vehicles/1","reasons":["vin"]}]}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	resp, err := c.CheckVehicleDuplicate(ctx, &CheckVehicleDuplicateParams{Vin: strPtr("ABC123")})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.JSON200 == nil || len(resp.JSON200.Matches) != 1 {
		t.Fatalf("expected 1 match, got %+v", resp.JSON200)
	}
}

func TestClient_UpdateCustomer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch || r.URL.Path != "/customers/1" {
			t.Errorf("expected PATCH /customers/1, got %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":1,"full_name":"Jane Doe"}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	req := UpdateCustomerRequest{}

	resp, err := c.UpdateCustomer(ctx, 1, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.JSON200 == nil || resp.JSON200.Id != 1 {
		t.Errorf("expected updated customer id 1, got %+v", resp.JSON200)
	}
}

func TestClient_ListVehicles(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"make":"Honda","model":"Civic","year":2020}]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	resp, err := c.ListVehicles(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.JSON200 == nil || len(*resp.JSON200) != 1 {
		t.Fatalf("expected 1 vehicle, got %+v", resp.JSON200)
	}
}

func TestListAccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/account" {
			t.Errorf("expected /account, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("expected bearer token, got %s", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"id":1,"name":"Main Shop"}`))
	}))
	defer server.Close()

	client := newTestClient(t, server.URL, "test-token")

	resp, err := client.ListAccount(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode())
	}
}

func TestShowLocation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/locations/1" {
			t.Errorf("expected /locations/1, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"id":1,"name":"Bay 1"}`))
	}))
	defer server.Close()

	client := newTestClient(t, server.URL, "test-token")

	resp, err := client.ShowLocation(context.Background(), "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode())
	}
}

func TestClient_ErrorMapping(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":{"code":"not_found","message":"Customer not found","field_errors":{}}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.ShowCustomer(ctx, 999)
	if err == nil {
		t.Fatal("expected error on 404")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected APIError, got %T", err)
	}
	if apiErr.Code != "not_found" || apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("expected not_found/404, got %s/%d", apiErr.Code, apiErr.StatusCode)
	}
}

func TestNewClient_RejectsHTTP(t *testing.T) {
	cfg := DefaultConfig()
	cfg.BaseURL = "http://app.wenmarpro.com"
	_, err := NewClient(cfg, NewStaticTokenProvider("test-token"))
	if err == nil {
		t.Fatal("expected error for http:// base URL, got nil")
	}
	if !strings.Contains(err.Error(), "https") {
		t.Errorf("expected error mentioning https, got: %v", err)
	}
}

func TestNewClient_AcceptsHTTPS(t *testing.T) {
	cfg := DefaultConfig()
	cfg.BaseURL = "https://app.wenmarpro.com"
	_, err := NewClient(cfg, NewStaticTokenProvider("test-token"))
	if err != nil {
		t.Fatalf("expected success for https://, got: %v", err)
	}
}

func TestNewClient_AcceptsLocalhostHTTP(t *testing.T) {
	cfg := DefaultConfig()
	cfg.BaseURL = "http://localhost:3000"
	_, err := NewClient(cfg, NewStaticTokenProvider("test-token"))
	if err != nil {
		t.Fatalf("expected success for http://localhost, got: %v", err)
	}
	cfg.BaseURL = "http://127.0.0.1:3000"
	_, err = NewClient(cfg, NewStaticTokenProvider("test-token"))
	if err != nil {
		t.Fatalf("expected success for http://127.0.0.1, got: %v", err)
	}
}

func TestClient_StripsAuthOnCrossOriginRedirect(t *testing.T) {
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "" {
			t.Errorf("Authorization header leaked to cross-origin redirect target: %q", auth)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer target.Close()

	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusFound)
	}))
	defer origin.Close()

	c := newTestClient(t, origin.URL, "test-token")
	_, err := c.ListCustomers(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestForLocation_InjectsHeader(t *testing.T) {
	var capturedLoc string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedLoc = r.Header.Get("X-Wenmar-Location")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	scoped := c.ForLocation("42")
	if _, err := scoped.ListCustomers(ctx, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedLoc != "42" {
		t.Errorf("expected X-Wenmar-Location '42', got %q", capturedLoc)
	}
}

func TestForLocation_DoesNotMutateParent(t *testing.T) {
	var capturedLoc string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedLoc = r.Header.Get("X-Wenmar-Location")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_ = c.ForLocation("42")
	if _, err := c.ListCustomers(ctx, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedLoc != "" {
		t.Errorf("parent client should not carry location header, got %q", capturedLoc)
	}
}
