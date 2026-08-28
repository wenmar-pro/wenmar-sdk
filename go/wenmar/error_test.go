package wenmar

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseError_NotFound(t *testing.T) {
	body := `{"error":{"code":"not_found","message":"Customer not found","details":{}}}`
	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusNotFound)
	resp.Body.WriteString(body)
	r := &http.Response{StatusCode: http.StatusNotFound, Body: resp.Result().Body}

	apiErr := ParseError(r)
	if apiErr.Code != "not_found" {
		t.Errorf("expected code 'not_found', got '%s'", apiErr.Code)
	}
	if apiErr.Message != "Customer not found" {
		t.Errorf("expected message 'Customer not found', got '%s'", apiErr.Message)
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", apiErr.StatusCode)
	}
}

func TestParseError_ValidationFailed(t *testing.T) {
	body := `{"error":{"code":"validation_failed","message":"Full name can't be blank","details":{"full_name":["can't be blank"]}}}`
	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusUnprocessableEntity)
	resp.Body.WriteString(body)
	r := &http.Response{StatusCode: http.StatusUnprocessableEntity, Body: resp.Result().Body}

	apiErr := ParseError(r)
	if apiErr.Code != "validation_failed" {
		t.Errorf("expected code 'validation_failed', got '%s'", apiErr.Code)
	}
	if apiErr.Details["full_name"] == nil {
		t.Error("expected details to have full_name key")
	}
}

func TestParseError_Unauthorized(t *testing.T) {
	body := `{"error":{"code":"unauthorized","message":"Invalid or missing API token","details":{}}}`
	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusUnauthorized)
	resp.Body.WriteString(body)
	r := &http.Response{StatusCode: http.StatusUnauthorized, Body: resp.Result().Body}

	apiErr := ParseError(r)
	if apiErr.Code != "unauthorized" {
		t.Errorf("expected code 'unauthorized', got '%s'", apiErr.Code)
	}
}

func TestAPIError_ErrorString(t *testing.T) {
	err := &APIError{Code: "not_found", Message: "Customer not found", StatusCode: 404}
	expected := "not_found: Customer not found (HTTP 404)"
	if err.Error() != expected {
		t.Errorf("expected '%s', got '%s'", expected, err.Error())
	}
}

func TestAPIError_ErrorStringWithRequest(t *testing.T) {
	err := &APIError{Code: "not_found", Message: "Customer not found", StatusCode: 404, Method: "GET", Path: "/customers/999"}
	expected := "GET /customers/999 -> not_found: Customer not found (HTTP 404)"
	if err.Error() != expected {
		t.Errorf("expected '%s', got '%s'", expected, err.Error())
	}
}

func TestParseErrorBodyWithRequest_SetsMethodAndPath(t *testing.T) {
	body := `{"error":{"code":"unauthorized","message":"Invalid or missing API token","details":{}}}`
	apiErr := ParseErrorBodyWithRequest([]byte(body), http.StatusUnauthorized, "GET", "/customers")
	if apiErr.Method != "GET" {
		t.Errorf("expected method 'GET', got '%s'", apiErr.Method)
	}
	if apiErr.Path != "/customers" {
		t.Errorf("expected path '/customers', got '%s'", apiErr.Path)
	}
	if apiErr.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", apiErr.StatusCode)
	}
}

func TestParseErrorBody_LeavesMethodAndPathEmpty(t *testing.T) {
	body := `{"error":{"code":"not_found","message":"Customer not found","details":{}}}`
	apiErr := ParseErrorBody([]byte(body), http.StatusNotFound)
	if apiErr.Method != "" {
		t.Errorf("expected empty method, got '%s'", apiErr.Method)
	}
	if apiErr.Path != "" {
		t.Errorf("expected empty path, got '%s'", apiErr.Path)
	}
}

func TestParseError_MalformedBody(t *testing.T) {
	body := `not json`
	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusInternalServerError)
	resp.Body.WriteString(body)
	r := &http.Response{StatusCode: http.StatusInternalServerError, Body: resp.Result().Body}

	apiErr := ParseError(r)
	if apiErr.Code != "unknown" {
		t.Errorf("expected code 'unknown' for malformed body, got '%s'", apiErr.Code)
	}
	if apiErr.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", apiErr.StatusCode)
	}
}

func TestAPIError_FieldErrors(t *testing.T) {
	apiErr := &APIError{
		Code:       "validation_failed",
		StatusCode: 422,
		Details: map[string]any{
			"first_name": []any{"can't be blank"},
			"last_name":  "can't be blank",
			"email":      []any{"is invalid", "already taken"},
		},
	}

	fe := apiErr.FieldErrors()
	if len(fe) != 3 {
		t.Fatalf("expected 3 field errors, got %d: %v", len(fe), fe)
	}
	if msgs, ok := fe["first_name"]; !ok || len(msgs) != 1 || msgs[0] != "can't be blank" {
		t.Errorf("first_name mismatch: %v", fe["first_name"])
	}
	if msgs, ok := fe["email"]; !ok || len(msgs) != 2 {
		t.Errorf("email mismatch: %v", fe["email"])
	}
}

func TestAPIError_Retryable(t *testing.T) {
	tests := []struct {
		code     string
		status   int
		expected bool
	}{
		{"rate_limited", 429, true},
		{"internal_error", 500, true},
		{"internal_error", 502, true},
		{"validation_failed", 422, false},
		{"not_found", 404, false},
		{"forbidden", 403, false},
		{"unauthorized", 401, false},
		{"limit_exceeded", 507, false},
	}
	for _, tt := range tests {
		apiErr := &APIError{Code: tt.code, StatusCode: tt.status}
		if got := apiErr.Retryable(); got != tt.expected {
			t.Errorf("Code=%q Status=%d: expected Retryable=%v, got %v", tt.code, tt.status, tt.expected, got)
		}
	}
}

func TestAPIError_RequestID(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-Id", "req-abc-123")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":{"code":"not_found","message":"x","details":{}}}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test-token")
	_, err := c.ShowCustomer(ctx, 999)
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.RequestID != "req-abc-123" {
		t.Errorf("expected RequestID 'req-abc-123', got %q", apiErr.RequestID)
	}
}

func TestParseError_507LimitExceeded(t *testing.T) {
	body := []byte(`{"error":{"code":"limit_exceeded","message":"Account limit reached","details":{}}}`)
	apiErr := ParseErrorBody(body, 507)
	if apiErr.Code != "limit_exceeded" {
		t.Errorf("expected code 'limit_exceeded', got %q", apiErr.Code)
	}
	if apiErr.StatusCode != 507 {
		t.Errorf("expected status 507, got %d", apiErr.StatusCode)
	}
	if apiErr.Retryable() {
		t.Error("limit_exceeded must not be retryable")
	}
}

func TestParseError_507StatusFallback(t *testing.T) {
	apiErr := ParseErrorBody([]byte{}, 507)
	if apiErr.Code != "limit_exceeded" {
		t.Errorf("expected code 'limit_exceeded' from status fallback, got %q", apiErr.Code)
	}
	if apiErr.Retryable() {
		t.Error("limit_exceeded must not be retryable")
	}
}
