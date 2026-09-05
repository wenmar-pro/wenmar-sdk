package conformance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	wenmar "github.com/wenmar-pro/wenmar-sdk/go/wenmar"
)

var ctx = context.Background()

type TestCase struct {
	Name          string            `json:"name"`
	Operation     string            `json:"operation"`
	Method        string            `json:"method"`
	Path          string            `json:"path"`
	PathParams    map[string]any    `json:"pathParams"`
	Query         map[string]any    `json:"query"`
	RequestBody   map[string]any    `json:"requestBody"`
	MockResponses []MockResponse    `json:"mockResponses"`
	Expect        Expectation       `json:"expect"`
}

type MockResponse struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body"`
}

type Expectation struct {
	RequestCount            int            `json:"requestCount"`
	NoError                 bool           `json:"noError"`
	ErrorCode               string         `json:"errorCode"`
	ErrorStatus             int            `json:"errorStatus"`
	FieldErrors             map[string][]string `json:"fieldErrors"`
	AssertNoOutboundRequest bool           `json:"assertNoOutboundRequest"`
	ResponseBody            *BodyAssertion `json:"responseBody"`
	RequestHeaders          map[string]string `json:"requestHeaders"`
}

type BodyAssertion struct {
	Path   string `json:"path"`
	Equals any    `json:"equals"`
}

func TestConformance(t *testing.T) {
	cases := loadTestCases(t)

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runCase(t, tc)
		})
	}
}

func runCase(t *testing.T, tc TestCase) {
	requestCount := 0
	responseIndex := 0
	var serverURL string
	var lastRequestHeaders http.Header

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		lastRequestHeaders = r.Header.Clone()
		idx := responseIndex
		responseIndex++

		var resp MockResponse
		if idx < len(tc.MockResponses) {
			resp = tc.MockResponses[idx]
		} else {
			resp = MockResponse{Status: 200, Body: json.RawMessage(`[]`)}
		}

		w.Header().Set("Content-Type", "application/json")
		for k, v := range resp.Headers {
			w.Header().Set(k, strings.ReplaceAll(v, "{server}", serverURL))
		}
		w.WriteHeader(resp.Status)
		if len(resp.Body) > 0 {
			w.Write(resp.Body)
		}
	}))
	defer server.Close()
	serverURL = server.URL

	cfg := wenmar.DefaultConfig()
	cfg.BaseURL = server.URL
	client, err := wenmar.NewClient(cfg, wenmar.NewStaticTokenProvider("test-token"))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	args := map[string]interface{}{
		"pathParams":  tc.PathParams,
		"query":       tc.Query,
		"requestBody": tc.RequestBody,
	}

	fn, ok := dispatch[tc.Operation]
	if !ok {
		t.Fatalf("operation %q not in dispatch", tc.Operation)
	}

	body, err := fn(ctx, t, client, args)
	if err != nil {
		if tc.Expect.NoError {
			t.Fatalf("expected success, got error: %v", err)
		}
		apiErr, ok := err.(*wenmar.APIError)
		if !ok {
			t.Fatalf("expected APIError, got %T: %v", err, err)
		}
		if tc.Expect.ErrorCode != "" && apiErr.Code != tc.Expect.ErrorCode {
			t.Errorf("expected error code '%s', got '%s'", tc.Expect.ErrorCode, apiErr.Code)
		}
		if tc.Expect.ErrorStatus != 0 && apiErr.StatusCode != tc.Expect.ErrorStatus {
			t.Errorf("expected status %d, got %d", tc.Expect.ErrorStatus, apiErr.StatusCode)
		}
		if tc.Expect.FieldErrors != nil {
			got := apiErr.FieldErrors()
			if len(got) != len(tc.Expect.FieldErrors) {
				t.Errorf("expected %d field errors, got %d: %v", len(tc.Expect.FieldErrors), len(got), got)
			}
			for field, msgs := range tc.Expect.FieldErrors {
				if !reflect.DeepEqual(got[field], msgs) {
					t.Errorf("field %q: expected %v, got %v", field, msgs, got[field])
				}
			}
		}
	} else {
		if !tc.Expect.NoError {
			t.Fatalf("expected error, got success")
		}
		if tc.Expect.ResponseBody != nil {
			assertBodyPath(t, body, tc.Expect.ResponseBody)
		}
	}

	if tc.Expect.RequestCount != 0 && requestCount != tc.Expect.RequestCount {
		t.Errorf("expected %d requests, got %d", tc.Expect.RequestCount, requestCount)
	}

	if tc.Expect.AssertNoOutboundRequest {
		if requestCount > len(tc.MockResponses) {
			t.Errorf("expected no outbound request beyond mocks, got %d calls", requestCount)
		}
	}

	if tc.Expect.RequestHeaders != nil {
		for k, v := range tc.Expect.RequestHeaders {
			if lastRequestHeaders.Get(k) != v {
				t.Errorf("expected header %s=%q, got %q", k, v, lastRequestHeaders.Get(k))
			}
		}
	}
}

// --- dispatch helpers (used by dispatch.gen.go) ---

// intArg extracts an integer path/query param from args.
func intArg(m map[string]interface{}, key string) int {
	if m == nil {
		return 0
	}
	switch v := m[key].(type) {
	case float64:
		return int(v)
	case int:
		return v
	case string:
		n, _ := strconv.Atoi(v)
		return n
	}
	return 0
}

// strArg extracts a string path/query param from args.
func strArg(m map[string]interface{}, key string) string {
	if m == nil {
		return ""
	}
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

// strPtr returns a *string for a query param, or nil if absent.
func strPtr(m map[string]interface{}, key string) *string {
	if m == nil {
		return nil
	}
	if v, ok := m[key].(string); ok {
		return &v
	}
	return nil
}

// boolPtr returns a *bool for a query param, or nil if absent.
func boolPtr(m map[string]interface{}, key string) *bool {
	if m == nil {
		return nil
	}
	if v, ok := m[key].(bool); ok {
		return &v
	}
	return nil
}

// intPtr returns a *int for a query param, or nil if absent.
func intPtr(m map[string]interface{}, key string) *int {
	if m == nil {
		return nil
	}
	switch v := m[key].(type) {
	case float64:
		i := int(v)
		return &i
	case int:
		return &v
	}
	return nil
}

// flatVal extracts a scalar value for a flat request body field.
func flatVal(m map[string]interface{}, key string) interface{} {
	if m == nil {
		return nil
	}
	return m[key]
}

// buildWrapper builds a request struct whose single required field is a
// wrapper (e.g. CreateCustomerRequest.Customer). It sets the wrapper field
// from the requestBody map.
func buildWrapper[T any](wrapper string, body map[string]interface{}) T {
	var out T
	// The wrapper field is the first (and only) field of the struct. We set
	// it by marshalling the body and unmarshalling into the struct, which
	// populates the nested field via its json tag.
	if body != nil {
		data, _ := json.Marshal(map[string]interface{}{wrapper: body})
		_ = json.Unmarshal(data, &out)
	}
	return out
}

// buildFlat builds a flat request struct (e.g. TransferVehicleRequest) from
// the request body map via JSON round-trip.
func buildFlat[T any](body map[string]interface{}) T {
	var out T
	if body != nil {
		data, _ := json.Marshal(body)
		_ = json.Unmarshal(data, &out)
	}
	return out
}

// paginateBody follows Link-header pages up to maxPages and returns the final
// decoded body. It uses the SDK's FetchPage so same-origin validation and the
// transport stack (retry/cache) apply.
func paginateBody(c *wenmar.Client, body []byte, link string, maxPages int) (interface{}, error) {
	for i := 0; i < maxPages; i++ {
		next := parseLink(link, "next")
		if next == "" {
			break
		}
		nextBody, nextLink, err := c.FetchPage(ctx, next)
		if err != nil {
			return nil, err
		}
		body = nextBody
		link = nextLink
	}
	return decodeBody(body)
}

func parseLink(header, rel string) string {
	if header == "" {
		return ""
	}
	for _, part := range strings.Split(header, ",") {
		part = strings.TrimSpace(part)
		if strings.Contains(part, `rel="`+rel+`"`) {
			start := strings.Index(part, "<")
			end := strings.Index(part, ">")
			if start >= 0 && end > start {
				return part[start+1 : end]
			}
		}
	}
	return ""
}

func decodeBody(body []byte) (any, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var v any
	if err := json.Unmarshal(body, &v); err != nil {
		return nil, err
	}
	return v, nil
}

func assertBodyPath(t *testing.T, body any, assertion *BodyAssertion) {
	value, ok := navigatePath(body, assertion.Path)
	if !ok {
		t.Fatalf("response body path '%s' not found", assertion.Path)
	}
	if !valuesEqual(value, assertion.Equals) {
		t.Errorf("expected %s to equal %v, got %v", assertion.Path, assertion.Equals, value)
	}
}

func navigatePath(body any, path string) (any, bool) {
	parts := strings.Split(path, ".")
	var current any = body
	for _, part := range parts {
		if idx, err := strconv.Atoi(part); err == nil {
			arr, ok := current.([]any)
			if !ok || idx >= len(arr) {
				return nil, false
			}
			current = arr[idx]
		} else {
			obj, ok := current.(map[string]any)
			if !ok {
				return nil, false
			}
			current, ok = obj[part]
			if !ok {
				return nil, false
			}
		}
	}
	return current, true
}

func valuesEqual(a, b any) bool {
	af, aok := a.(float64)
	bf, bok := b.(float64)
	if aok && bok {
		return af == bf
	}
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func loadTestCases(t *testing.T) []TestCase {
	files, err := filepath.Glob("../tests/*.json")
	if err != nil {
		t.Fatalf("failed to glob test files: %v", err)
	}
	var cases []TestCase
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("failed to read %s: %v", f, err)
		}
		var fileCases []TestCase
		if err := json.Unmarshal(data, &fileCases); err != nil {
			t.Fatalf("failed to parse %s: %v", f, err)
		}
		cases = append(cases, fileCases...)
	}
	return cases
}
