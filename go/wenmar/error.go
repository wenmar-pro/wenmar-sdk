package wenmar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIError struct {
	Code       string         `json:"code"`
	Message    string         `json:"message"`
	Details    map[string]any `json:"details"`
	StatusCode int            `json:"-"`
	Method     string         `json:"-"`
	Path       string         `json:"-"`
}

func (e *APIError) Error() string {
	s := fmt.Sprintf("%s: %s (HTTP %d)", e.Code, e.Message, e.StatusCode)
	if e.Method != "" && e.Path != "" {
		s = fmt.Sprintf("%s %s -> %s", e.Method, e.Path, s)
	}
	return s
}

// ParseError reads the error envelope from a failed response body.
func ParseError(resp *http.Response) *APIError {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIError{Code: "unknown", Message: "unreadable body", StatusCode: resp.StatusCode}
	}
	return ParseErrorBody(body, resp.StatusCode)
}

// ParseErrorBody parses the { "error": { code, message, details } } envelope
// from an already-read response body. The generated oapi-codegen client drains
// the body into a byte slice, so callers pass that slice here.
func ParseErrorBody(body []byte, statusCode int) *APIError {
	return ParseErrorBodyWithRequest(body, statusCode, "", "")
}

// ParseErrorBodyWithRequest is like ParseErrorBody but also records the HTTP
// method and request path that produced the error, for richer diagnostics.
func ParseErrorBodyWithRequest(body []byte, statusCode int, method, path string) *APIError {
	apiErr := &APIError{StatusCode: statusCode, Method: method, Path: path}

	if len(body) == 0 {
		apiErr.Code = "unknown"
		apiErr.Message = fmt.Sprintf("HTTP %d with empty or unreadable body", statusCode)
		return apiErr
	}

	var envelope struct {
		Error struct {
			Code    string         `json:"code"`
			Message string         `json:"message"`
			Details map[string]any `json:"details"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &envelope); err != nil {
		apiErr.Code = "unknown"
		apiErr.Message = fmt.Sprintf("HTTP %d with malformed body", statusCode)
		return apiErr
	}

	apiErr.Code = envelope.Error.Code
	apiErr.Message = envelope.Error.Message
	apiErr.Details = envelope.Error.Details
	if apiErr.Code == "" {
		apiErr.Code = "unknown"
	}

	return apiErr
}
