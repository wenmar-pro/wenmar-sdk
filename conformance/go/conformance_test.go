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
	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
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
}

// executeOperation dispatches the SDK call and returns the raw response body
// (as a decoded generic value) for body assertions.
func executeOperation(client *wenmar.Client, tc TestCase) (any, error) {
	switch tc.Operation {
	case "list_customers":
		resp, err := client.ListCustomers(ctx)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_customers_paginated":
		resp, paginator, err := client.ListCustomersWithPagination(ctx)
		if err != nil {
			return nil, err
		}
		body, err := decodeBody(resp.Body)
		if err != nil {
			return nil, err
		}
		for paginator.HasNext() {
			next, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			body = next
		}
		return body, nil
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
		var req generated.CreateCustomerJSONRequestBody
		if customer != nil {
			req.Customer.FirstName = stringVal(customer["first_name"])
			req.Customer.LastName = stringVal(customer["last_name"])
		}
		resp, err := client.CreateCustomer(ctx, req)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "update_customer":
		id := intParam(tc.PathParams, "id")
		resp, err := client.UpdateCustomer(ctx, id, generated.UpdateCustomerJSONRequestBody{})
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_vehicles":
		resp, err := client.ListVehicles(ctx)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "create_vehicle":
		body := tc.RequestBody
		vehicle, _ := body["vehicle"].(map[string]any)
		var req generated.CreateVehicleJSONRequestBody
		if vehicle != nil {
			req.Vehicle.Make = stringVal(vehicle["make"])
			req.Vehicle.Model = stringVal(vehicle["model"])
			req.Vehicle.Year = intParamValue(vehicle, "year")
			req.Vehicle.CustomerId = intParamValue(vehicle, "customer_id")
		}
		resp, err := client.CreateVehicle(ctx, req)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "update_vehicle":
		id := intParam(tc.PathParams, "id")
		body := tc.RequestBody
		vehicle, _ := body["vehicle"].(map[string]any)
		var req generated.UpdateVehicleJSONRequestBody
		if vehicle != nil {
			req.Vehicle.Make = stringVal(vehicle["make"])
		}
		resp, err := client.UpdateVehicle(ctx, id, req)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "delete_vehicle":
		id := intParam(tc.PathParams, "id")
		resp, err := client.DeleteVehicle(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "decode_vin":
		vin := stringParam(tc.Query, "vin")
		resp, err := client.DecodeVin(ctx, vin)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "check_duplicate":
		vin := stringParam(tc.Query, "vin")
		resp, err := client.CheckVehicleDuplicate(ctx, generated.CheckVehicleDuplicateParams{Vin: &vin})
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
		resp, err := client.ListWorkOrders(ctx)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_work_orders_paginated":
		resp, paginator, err := client.ListWorkOrdersWithPagination(ctx)
		if err != nil {
			return nil, err
		}
		body, err := decodeBody(resp.Body)
		if err != nil {
			return nil, err
		}
		for paginator.HasNext() {
			next, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			body = next
		}
		return body, nil
	case "show_work_order":
		id := intParam(tc.PathParams, "id")
		resp, err := client.ShowWorkOrder(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_account":
		resp, err := client.ListAccount(ctx)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_location":
		id := stringParam(tc.PathParams, "id")
		resp, err := client.ShowLocation(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_customers_drivers":
		customerID := intParam(tc.PathParams, "customer_id")
		resp, err := client.ListDrivers(ctx, customerID)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_driver":
		customerID := intParam(tc.PathParams, "customer_id")
		id := intParam(tc.PathParams, "id")
		resp, err := client.ShowDriver(ctx, customerID, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "create_driver":
		customerID := intParam(tc.PathParams, "customer_id")
		body := tc.RequestBody
		driver, _ := body["driver"].(map[string]any)
		req := wenmar.CreateDriverRequest{}
		if driver != nil {
			req.FullName, _ = driver["full_name"].(string)
			req.Phone, _ = driver["phone"].(string)
		}
		resp, err := client.CreateDriver(ctx, customerID, req)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "update_driver":
		customerID := intParam(tc.PathParams, "customer_id")
		id := intParam(tc.PathParams, "id")
		resp, err := client.UpdateDriver(ctx, customerID, id, wenmar.UpdateDriverRequest{})
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "delete_driver":
		customerID := intParam(tc.PathParams, "customer_id")
		id := intParam(tc.PathParams, "id")
		resp, err := client.DeleteDriver(ctx, customerID, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_customers_statements":
		customerID := intParam(tc.PathParams, "customer_id")
		resp, err := client.ListStatements(ctx, customerID)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_customer_vehicles":
		customerID := intParam(tc.PathParams, "customer_id")
		resp, err := client.ListCustomerVehicles(ctx, customerID)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_customer_work_orders":
		customerID := intParam(tc.PathParams, "customer_id")
		resp, err := client.ListCustomerWorkOrders(ctx, customerID)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_vehicle_work_orders":
		vehicleID := intParam(tc.PathParams, "vehicle_id")
		resp, err := client.ListVehicleWorkOrders(ctx, vehicleID)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "merge_customer":
		id := intParam(tc.PathParams, "id")
		body := tc.RequestBody
		req := generated.MergeCustomerJSONRequestBody{SourceCustomerId: intParamValue(body, "source_customer_id")}
		resp, err := client.MergeCustomer(ctx, id, req)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "transfer_vehicle":
		id := intParam(tc.PathParams, "id")
		body := tc.RequestBody
		req := generated.TransferVehicleJSONRequestBody{
			CustomerId: intParamValue(body, "customer_id"),
			Mode:       stringVal(body["mode"]),
		}
		resp, err := client.TransferVehicle(ctx, id, req)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "merge_vehicle":
		id := intParam(tc.PathParams, "id")
		body := tc.RequestBody
		req := generated.MergeVehicleJSONRequestBody{SourceVehicleId: intParamValue(body, "source_vehicle_id")}
		resp, err := client.MergeVehicle(ctx, id, req)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_tags":
		resp, err := client.ListTags(ctx)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "update_tags":
		resp, err := client.UpdateTags(ctx, generated.UpdateTagsJSONRequestBody{})
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "create_customer_tag":
		resp, err := client.CreateCustomerTag(ctx, generated.CreateCustomerTagJSONRequestBody{Name: stringParam(tc.RequestBody, "name")})
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "create_vehicle_tag":
		resp, err := client.CreateVehicleTag(ctx, generated.CreateVehicleTagJSONRequestBody{Name: stringParam(tc.RequestBody, "name")})
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "lookup_customer":
		resp, err := client.LookupCustomer(ctx, stringParam(tc.Query, "query"))
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "lookup_vehicle":
		resp, err := client.LookupVehicle(ctx, stringParam(tc.Query, "query"))
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "prefill_vehicle":
		vin := stringParam(tc.Query, "vin")
		resp, err := client.PrefillVehicle(ctx, generated.PrefillVehicleParams{Vin: &vin})
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "check_customer_duplicate":
		fn := stringParam(tc.Query, "first_name")
		ln := stringParam(tc.Query, "last_name")
		resp, err := client.CheckCustomerDuplicate(ctx, generated.CheckCustomerDuplicateParams{
			FirstName: &fn,
			LastName:  &ln,
		})
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "check_vehicle_duplicate":
		vin := stringParam(tc.Query, "vin")
		resp, err := client.CheckVehicleDuplicate(ctx, generated.CheckVehicleDuplicateParams{Vin: &vin})
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_statement":
		id := intParam(tc.PathParams, "id")
		resp, err := client.ShowStatement(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "list_vendors":
		resp, err := client.ListVendors(ctx)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_vendor":
		id := intParam(tc.PathParams, "id")
		resp, err := client.ShowVendor(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_work_order_estimate":
		id := intParam(tc.PathParams, "work_order_id")
		resp, err := client.ShowWorkOrderEstimate(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_work_order_wip":
		id := intParam(tc.PathParams, "work_order_id")
		resp, err := client.ShowWorkOrderWip(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_work_order_inspection":
		id := intParam(tc.PathParams, "work_order_id")
		resp, err := client.ShowWorkOrderInspection(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_work_order_parts":
		id := intParam(tc.PathParams, "work_order_id")
		resp, err := client.ShowWorkOrderParts(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "show_work_order_payments":
		id := intParam(tc.PathParams, "work_order_id")
		resp, err := client.ShowWorkOrderPayments(ctx, id)
		if err != nil {
			return nil, err
		}
		return decodeBody(resp.Body)
	case "create_work_order_payment":
		id := intParam(tc.PathParams, "work_order_id")
		body := tc.RequestBody
		payment, _ := body["payment"].(map[string]any)
		req := wenmar.CreateWorkOrderPaymentRequest{}
		if payment != nil {
			req.AmountCents, _ = payment["amount_cents"].(string)
			req.Method, _ = payment["method"].(string)
		}
		resp, err := client.CreateWorkOrderPayment(ctx, id, req)
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

func intParamValue(params map[string]any, key string) int {
	return intParam(params, key)
}

func stringVal(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

func mustStr(s string) string { return s }

func stringParam(params map[string]any, key string) string {
	if params == nil {
		return ""
	}
	if v, ok := params[key].(string); ok {
		return v
	}
	return ""
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
