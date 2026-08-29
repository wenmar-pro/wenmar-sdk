package wenmar

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCustomer_WithTypedRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id":42,"first_name":"Jane","last_name":"Doe","type":"Customer","url":"https://example.com/customers/42.json","app_url":"https://example.com/customers/42","created_at":"2026-08-27T12:00:00.000Z","updated_at":"2026-08-27T12:00:00.000Z"}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	req := CreateCustomerRequest{FirstName: "Jane", LastName: "Doe"}
	resp, err := c.CreateCustomer(ctx, req)
	if err != nil {
		t.Fatalf("CreateCustomer failed: %v", err)
	}
	if resp.JSON201.Id != 42 {
		t.Errorf("expected id 42, got %d", resp.JSON201.Id)
	}
}