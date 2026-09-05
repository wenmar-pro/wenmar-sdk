package wenmar

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
)

func TestParseLinkHeader_Next(t *testing.T) {
	header := `<https://api.example.com/customers?page=2>; rel="next"`
	next := parseLinkHeader(header, "next")
	if next != "https://api.example.com/customers?page=2" {
		t.Errorf("expected next URL, got '%s'", next)
	}
}

func TestParseLinkHeader_Prev(t *testing.T) {
	header := `<https://api.example.com/customers?page=1>; rel="prev", <https://api.example.com/customers?page=3>; rel="next"`
	prev := parseLinkHeader(header, "prev")
	if prev != "https://api.example.com/customers?page=1" {
		t.Errorf("expected prev URL, got '%s'", prev)
	}
}

func TestParseLinkHeader_Empty(t *testing.T) {
	result := parseLinkHeader("", "next")
	if result != "" {
		t.Errorf("expected empty string for empty header, got '%s'", result)
	}
}

func TestNewListResultFromResponse_ExtractsItemsAndMeta(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Link", "<http://api.example.com/customers?page=2>; rel=\"next\"")
	headers.Set("X-Total-Count", "42")
	headers.Set("X-Per-Page", "25")
	headers.Set("Content-Type", "application/json")

	body := []byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`)

	client := newTestClient(t, "https://api.example.com", "test")
	result := newListResultFromResponse[Customer](body, headers, client)

	if len(result.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(result.Items))
	}
	if result.Items[0].Id != 1 {
		t.Errorf("expected item id=1, got %d", result.Items[0].Id)
	}
	if !result.Meta.HasMore {
		t.Error("expected HasMore=true")
	}
	if result.Meta.TotalCount != 42 {
		t.Errorf("expected TotalCount=42, got %d", result.Meta.TotalCount)
	}
	if result.Meta.PerPage != 25 {
		t.Errorf("expected PerPage=25, got %d", result.Meta.PerPage)
	}
	if !result.HasNext() {
		t.Error("expected HasNext()=true")
	}
}

func TestGetAllWithOptions_RespectsMaxItems(t *testing.T) {
	var serverURL string
	var call int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&call, 1)
		w.Header().Set("Content-Type", "application/json")
		if atomic.LoadInt32(&call) == 1 {
			w.Header().Set("Link", fmt.Sprintf(`<%s/items?page=2>; rel="next"`, serverURL))
			w.Write([]byte(`[{"id":1},{"id":2},{"id":3}]`))
		} else {
			w.Write([]byte(`[{"id":4},{"id":5}]`))
		}
	}))
	defer ts.Close()
	serverURL = ts.URL

	c := newTestClient(t, ts.URL, "test")
	// Build a first ListResult manually
	headers := make(http.Header)
	headers.Set("Link", fmt.Sprintf(`<%s/items?page=2>; rel="next"`, serverURL))
	headers.Set("Content-Type", "application/json")
	first := newListResultFromResponse[map[string]any](
		[]byte(`[{"id":1},{"id":2},{"id":3}]`),
		headers,
		c,
	)
	items, truncated, err := getAll[map[string]any](context.Background(), first, &GetAllOptions{MaxItems: 2})
	if err != nil {
		t.Fatalf("getAll failed: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items (MaxItems cap), got %d", len(items))
	}
	if !truncated {
		t.Error("expected truncated=true")
	}
}

func TestGetAllCustomers_FollowsLinkHeader(t *testing.T) {
	var serverURL string
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		if calls == 1 {
			w.Header().Set("Link", fmt.Sprintf(`<%s/customers?page=2>; rel="next"`, serverURL))
		}
		w.Write([]byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
	}))
	defer ts.Close()
	serverURL = ts.URL

	c := newTestClient(t, ts.URL, "test-token")
	items, err := c.GetAllCustomers(ctx, nil, nil)
	if err != nil {
		t.Fatalf("GetAllCustomers failed: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items across 2 pages, got %d", len(items))
	}
	if calls != 2 {
		t.Errorf("expected 2 calls, got %d", calls)
	}
}

func TestGetAllCustomers_CapsAtMax(t *testing.T) {
	var serverURL string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Link", fmt.Sprintf(`<%s/customers?page=next>; rel="next"`, serverURL))
		w.Write([]byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"},{"id":2,"type":"Customer","first_name":"C","last_name":"D","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
	}))
	defer ts.Close()
	serverURL = ts.URL

	c := newTestClient(t, ts.URL, "test-token")
	items, err := collectAll[Customer](ctx, c, []byte(`[]`), fmt.Sprintf(`<%s/customers?page=1>; rel="next"`, serverURL), 2)
	if err != nil {
		t.Fatalf("collectAll failed: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items capped, got %d", len(items))
	}
}

func TestPaginationFollowsNextURL(t *testing.T) {
	var serverURL string
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		if requestCount == 1 {
			w.Header().Set("Link", `<`+serverURL+`/customers?page=2>; rel="next"`)
			w.WriteHeader(200)
			w.Write([]byte(`[{"id":1,"full_name":"Page1Customer"}]`))
			return
		}
		if r.URL.Query().Get("page") != "2" {
			t.Errorf("expected page=2 query param, got %q", r.URL.Query().Get("page"))
		}
		w.WriteHeader(200)
		w.Write([]byte(`[{"id":2,"full_name":"Page2Customer"}]`))
	}))
	defer server.Close()
	serverURL = server.URL

	client := newTestClient(t, server.URL, "test-token")

	resp, err := client.ListCustomersRaw(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var page1 []map[string]any
	json.Unmarshal(resp.Body, &page1)
	if len(page1) != 1 || page1[0]["id"] != float64(1) {
		t.Fatalf("expected page 1 customer, got %v", page1)
	}

	// Follow the Link header manually via fetchURL.
	nextURL := parseLinkHeader(resp.HTTPResponse.Header.Get("Link"), "next")
	if nextURL == "" {
		t.Fatal("expected Link header with next URL")
	}
	body, link, err := client.fetchURL(context.Background(), nextURL)
	if err != nil {
		t.Fatalf("unexpected error on fetchURL: %v", err)
	}
	var page2 []map[string]any
	json.Unmarshal(body, &page2)
	if len(page2) != 1 || page2[0]["id"] != float64(2) {
		t.Fatalf("expected page 2 customer (id=2), got %v", page2)
	}
	if parseLinkHeader(link, "next") != "" {
		t.Errorf("expected no further next link on page 2, got %q", link)
	}
}

func TestGetAllWorkOrdersConcerns_DoesNotReturnWorkOrders(t *testing.T) {
	var serverURL string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"id":1,"name":"brakes squealing","customer_complaint":"yes"}]`))
	}))
	defer ts.Close()
	serverURL = ts.URL

	c := newTestClient(t, ts.URL, "test-token")
	// list_work_orders_concerns returns concerns, not work orders.
	// The concern JSON has fields like "name" and "customer_complaint"
	// that don't exist on WorkOrder. If GetAllWorkOrdersConcerns returns
	// []WorkOrder, those fields are silently dropped.
	resp, err := c.ListWorkOrdersConcerns(context.Background(), 1)
	if err != nil {
		t.Fatalf("ListWorkOrdersConcerns failed: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode())
	}
	// The response body should contain the concern's "name" field.
	body := string(resp.Body)
	if !strings.Contains(body, "brakes squealing") {
		t.Errorf("response body should contain concern name, got: %s", body)
	}
	_ = serverURL // keep for future Link-header pagination test
}

func TestGetAllWorkOrdersConcerns_NotGenerated(t *testing.T) {
	// After the fix, GetAllWorkOrdersConcerns should NOT exist because
	// the response is not a WorkOrder. The broken version returned
	// []WorkOrder for a concerns endpoint, silently dropping fields.
	c := newTestClient(t, "http://localhost:9999", "test-token")
	// Use reflection to check the method does NOT exist.
	ty := reflect.TypeOf(c)
	_, exists := ty.MethodByName("GetAllWorkOrdersConcerns")
	if exists {
		t.Error("GetAllWorkOrdersConcerns should not exist — concerns are not WorkOrders")
	}
}

func TestPagination_RejectsCrossOriginNextURL(t *testing.T) {
	var outboundRequests int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&outboundRequests, 1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Link", `<https://attacker.example.com/customers?page=2>; rel="next"`)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1}]`))
	}))
	defer ts.Close()

	var attackerRequests int32
	var attackerGotAuth bool
	attacker := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attackerRequests, 1)
		if r.Header.Get("Authorization") != "" {
			attackerGotAuth = true
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer attacker.Close()

	c := newTestClient(t, ts.URL, "test-token")
	nextURL := "https://attacker.example.com/customers?page=2"
	_, _, err := c.fetchURL(ctx, nextURL)
	if err == nil {
		t.Fatal("expected error on cross-origin next URL, got nil")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.Code != "invalid_pagination" {
		t.Errorf("expected code 'invalid_pagination', got %q", apiErr.Code)
	}

	if attackerRequests != 0 {
		t.Errorf("expected 0 requests to attacker origin, got %d", attackerRequests)
	}
	if attackerGotAuth {
		t.Error("Authorization header leaked to cross-origin attacker server")
	}
}

func TestListCustomers_ReturnsTypedListResult(t *testing.T) {
	var serverURL string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Link", fmt.Sprintf(`<%s/customers?page=2>; rel="next"`, serverURL))
		w.Header().Set("X-Total-Count", "2")
		w.Write([]byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
	}))
	defer ts.Close()
	serverURL = ts.URL

	c := newTestClient(t, ts.URL, "test")
	result, err := c.ListCustomers(context.Background(), nil)
	if err != nil {
		t.Fatalf("ListCustomers failed: %v", err)
	}
	if len(result.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(result.Items))
	}
	if result.Items[0].Id != 1 {
		t.Errorf("expected id=1, got %d", result.Items[0].Id)
	}
	if !result.HasNext() {
		t.Error("expected HasNext()=true")
	}
	if result.Meta.TotalCount != 2 {
		t.Errorf("expected TotalCount=2, got %d", result.Meta.TotalCount)
	}
}

func TestListCustomers_NextPage(t *testing.T) {
	var serverURL string
	var call int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&call, 1)
		w.Header().Set("Content-Type", "application/json")
		if atomic.LoadInt32(&call) == 1 {
			w.Header().Set("Link", fmt.Sprintf(`<%s/customers?page=2>; rel="next"`, serverURL))
			w.Write([]byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
		} else {
			w.Write([]byte(`[{"id":2,"type":"Customer","first_name":"C","last_name":"D","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
		}
	}))
	defer ts.Close()
	serverURL = ts.URL

	c := newTestClient(t, ts.URL, "test")
	result, err := c.ListCustomers(context.Background(), nil)
	if err != nil {
		t.Fatalf("ListCustomers failed: %v", err)
	}
	if !result.HasNext() {
		t.Fatal("expected HasNext()=true")
	}
	page2, err := result.Next(context.Background())
	if err != nil {
		t.Fatalf("Next failed: %v", err)
	}
	if len(page2.Items) != 1 {
		t.Fatalf("expected 1 item on page 2, got %d", len(page2.Items))
	}
	if page2.Items[0].Id != 2 {
		t.Errorf("expected id=2, got %d", page2.Items[0].Id)
	}
	if page2.HasNext() {
		t.Error("expected HasNext()=false on page 2")
	}
}

func TestGetAllCustomers_WithMaxItemsOption(t *testing.T) {
	var serverURL string
	var call int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&call, 1)
		w.Header().Set("Content-Type", "application/json")
		if atomic.LoadInt32(&call) == 1 {
			w.Header().Set("Link", fmt.Sprintf(`<%s/customers?page=2>; rel="next"`, serverURL))
			w.Write([]byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
		} else {
			w.Write([]byte(`[{"id":2,"type":"Customer","first_name":"C","last_name":"D","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
		}
	}))
	defer ts.Close()
	serverURL = ts.URL

	c := newTestClient(t, ts.URL, "test")
	items, err := c.GetAllCustomers(context.Background(), nil, &GetAllOptions{MaxItems: 1})
	if err != nil {
		t.Fatalf("GetAllCustomers failed: %v", err)
	}
	if len(items) != 1 {
		t.Errorf("expected 1 item (MaxItems=1), got %d", len(items))
	}
}
