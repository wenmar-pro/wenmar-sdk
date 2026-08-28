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
