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
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s (HTTP %d)", e.Code, e.Message, e.StatusCode)
}

func ParseError(resp *http.Response) *APIError {
	apiErr := &APIError{StatusCode: resp.StatusCode}

	body, err := io.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		apiErr.Code = "unknown"
		apiErr.Message = fmt.Sprintf("HTTP %d with empty or unreadable body", resp.StatusCode)
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
		apiErr.Message = fmt.Sprintf("HTTP %d with malformed body", resp.StatusCode)
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
