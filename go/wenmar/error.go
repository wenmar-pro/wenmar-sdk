package wenmar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIError struct {
	Code           string         `json:"code"`
	Message        string         `json:"message"`
	FieldErrorsMap map[string]any `json:"field_errors"`
	StatusCode     int            `json:"-"`
	Method         string         `json:"-"`
	Path           string         `json:"-"`
	RequestID      string         `json:"-"`
}

func (e *APIError) Error() string {
	s := fmt.Sprintf("%s: %s (HTTP %d)", e.Code, e.Message, e.StatusCode)
	if e.Method != "" && e.Path != "" {
		s = fmt.Sprintf("%s %s -> %s", e.Method, e.Path, s)
	}
	if e.RequestID != "" {
		s = fmt.Sprintf("%s [request_id=%s]", s, e.RequestID)
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

// ParseErrorBody parses the { "error": { code, message, field_errors } } envelope
// from an already-read response body. The generated oapi-codegen client drains
// the body into a byte slice, so callers pass that slice here.
func ParseErrorBody(body []byte, statusCode int) *APIError {
	return ParseErrorBodyWithRequest(body, statusCode, "", "")
}

// ParseErrorBodyWithRequest is like ParseErrorBody but also records the HTTP
// method and request path that produced the error, for richer diagnostics.
func ParseErrorBodyWithRequest(body []byte, statusCode int, method, path string) *APIError {
	return ParseErrorBodyWithRequestAndID(body, statusCode, method, path, "")
}

// ParseErrorBodyWithRequestAndID is like ParseErrorBodyWithRequest but also
// records the X-Request-Id header for support correlation.
func ParseErrorBodyWithRequestAndID(body []byte, statusCode int, method, path, requestID string) *APIError {
	apiErr := &APIError{StatusCode: statusCode, Method: method, Path: path, RequestID: requestID}

	if len(body) == 0 {
		apiErr.Code = statusFallbackCode(statusCode)
		if apiErr.Code == "" {
			apiErr.Code = "unknown"
		}
		apiErr.Message = fmt.Sprintf("HTTP %d with empty or unreadable body", statusCode)
		return apiErr
	}

	var envelope struct {
		Error struct {
			Code        string         `json:"code"`
			Message     string         `json:"message"`
			FieldErrors map[string]any `json:"field_errors"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &envelope); err != nil {
		apiErr.Code = statusFallbackCode(statusCode)
		if apiErr.Code == "" {
			apiErr.Code = "unknown"
		}
		apiErr.Message = fmt.Sprintf("HTTP %d with malformed body", statusCode)
		return apiErr
	}

	apiErr.Code = envelope.Error.Code
	apiErr.Message = envelope.Error.Message
	apiErr.FieldErrorsMap = envelope.Error.FieldErrors
	if apiErr.Code == "" {
		apiErr.Code = statusFallbackCode(statusCode)
		if apiErr.Code == "" {
			apiErr.Code = "unknown"
		}
	}

	return apiErr
}

// statusFallbackCode returns a code for a status when the body is empty
// or malformed. This lets the SDK recognize 507 limit_exceeded even when
// the server returns no body.
func statusFallbackCode(statusCode int) string {
	switch statusCode {
	case 507:
		return "limit_exceeded"
	default:
		return ""
	}
}

// FieldErrors extracts validation field errors from the FieldErrors map.
// The Wenmar API sends validation errors as:
//   field_errors: { "first_name": ["can't be blank"], "email": ["is invalid"] }
// This method coerces the loosely-typed JSON values into a
// map[string][]string for easy form-level error display.
// Returns nil if there are no field errors.
func (e *APIError) FieldErrors() map[string][]string {
	if e.FieldErrorsMap == nil {
		return nil
	}
	result := make(map[string][]string, len(e.FieldErrorsMap))
	for field, raw := range e.FieldErrorsMap {
		switch v := raw.(type) {
		case []any:
			msgs := make([]string, 0, len(v))
			for _, item := range v {
				if s, ok := item.(string); ok {
					msgs = append(msgs, s)
				}
			}
			if len(msgs) > 0 {
				result[field] = msgs
			}
		case string:
			result[field] = []string{v}
		case []string:
			result[field] = v
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// Retryable reports whether the error is worth retrying. Rate limits and
// 5xx server errors are transient; validation, auth, not-found, forbidden,
// and limit_exceeded (507) are terminal. The caller should still respect
// method semantics: mutations should not be retried even if Retryable is
// true (the retryTransport enforces this for 5xx).
func (e *APIError) Retryable() bool {
	switch e.Code {
	case "rate_limited":
		return true
	case "internal_error":
		return e.StatusCode >= 500
	case "limit_exceeded":
		return false
	}
	if e.StatusCode >= 500 && e.StatusCode != 507 {
		return true
	}
	return false
}
