package conformance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	wenmar "github.com/wenmar-pro/wenmar-sdk/go/wenmar"
	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

var ctx = context.Background()

type TestCase struct {
	Name          string            `json:"name"`
	Operation     string            `json:"operation"`
	Method        string            `json:"method"`
	Path          string            `json:"path"`
	PathParams    map[string]any    `json:"pathParams"`
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
	RequestCount  int            `json:"requestCount"`
	NoError       bool           `json:"noError"`
	ErrorCode     string         `json:"errorCode"`
	ErrorStatus   int            `json:"errorStatus"`
	ResponseBody  *BodyAssertion `json:"responseBody"`
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

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		idx := responseIndex
		responseIndex++

		var resp MockResponse
		if idx < len(tc.MockResponses) {
			resp = tc.MockResponses[idx]
		} else {
			resp = MockResponse{Status: 200, Body: json.RawMessage(`{"data":[]}`)}
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

	client, err := wenmar.NewClient(server.URL, "test-token")
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	body, err := executeOperation(client, tc)
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
}

// executeOperation dispatches the SDK call and returns the raw response body
// (as a decoded generic value) for body assertions.
func executeOperation(client *wenmar.Client, tc TestCase) (any, error) {
	switch tc.Operation {
	case "list_customers":
		resp, err := client.ListCustomers(ctx, nil)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_customers_paginated":
		resp, paginator, err := client.ListCustomersWithPagination(ctx, nil)
		if err != nil {
			return nil, err
		}
		for paginator.HasNext() {
			next, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			resp = next.(*generated.ListCustomersResponse)
		}
		return decodeBody(resp.Body)
	case "show_customer":
		id := intParam(tc.PathParams, "id")
		resp, err := client.ShowCustomer(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "create_customer":
		body := tc.RequestBody
		customer, _ := body["customer"].(map[string]any)
		reqBody := generated.CreateCustomerJSONBody{}
		if customer != nil {
			fullName, _ := customer["full_name"].(string)
			reqBody.Customer = &struct {
				Email    *string `json:"email,omitempty"`
				FullName string  `json:"full_name"`
				Phone    *string `json:"phone,omitempty"`
			}{FullName: fullName}
		}
		resp, err := client.CreateCustomer(ctx, generated.CreateCustomerJSONRequestBody(reqBody))
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_vehicle":
		id := intParam(tc.PathParams, "id")
		resp, err := client.ShowVehicle(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_work_orders":
		resp, err := client.ListWorkOrders(ctx, nil)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_work_orders_paginated":
		resp, paginator, err := client.ListWorkOrdersWithPagination(ctx, nil)
		if err != nil {
			return nil, err
		}
		for paginator.HasNext() {
			next, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			resp = next.(*generated.ListWorkOrdersResponse)
		}
		return decodeBody(resp.Body)
	case "show_work_order":
		id := intParam(tc.PathParams, "id")
		resp, err := client.ShowWorkOrder(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	default:
		return nil, fmt.Errorf("unknown operation: %s", tc.Operation)
	}
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

func intParam(params map[string]any, key string) int {
	if params == nil {
		return 0
	}
	switch v := params[key].(type) {
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
