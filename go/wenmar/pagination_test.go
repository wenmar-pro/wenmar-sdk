package wenmar

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var serverURLForTest string

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

func TestPaginator_HasNext(t *testing.T) {
	p := &Paginator{nextURL: "https://api.example.com/customers?page=2"}
	if !p.HasNext() {
		t.Error("expected HasNext to be true")
	}

	p = &Paginator{nextURL: ""}
	if p.HasNext() {
		t.Error("expected HasNext to be false")
	}
}

func TestNewPaginatorFromResponse(t *testing.T) {
	resp := httptest.NewRecorder()
	resp.Header().Set("Link", `<https://api.example.com/customers?page=2>; rel="next"`)
	r := resp.Result()

	p := newPaginatorFromResponse(r, nil)
	if !p.HasNext() {
		t.Error("expected paginator to have next")
	}
	if p.nextURL != "https://api.example.com/customers?page=2" {
		t.Errorf("expected next URL, got '%s'", p.nextURL)
	}
}

func TestPaginationFollowsNextURL(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		if requestCount == 1 {
			// First page: return Link header pointing to ?page=2
			w.Header().Set("Link", `<`+serverURLForTest+`/customers?page=2>; rel="next"`)
			w.WriteHeader(200)
			w.Write([]byte(`[{"id":1,"full_name":"Page1Customer"}]`))
			return
		}
		// Page 2: verify the request actually has ?page=2
		if r.URL.Query().Get("page") != "2" {
			t.Errorf("expected page=2 query param, got %q", r.URL.Query().Get("page"))
		}
		w.WriteHeader(200)
		w.Write([]byte(`[{"id":2,"full_name":"Page2Customer"}]`))
	}))
	defer server.Close()
	serverURLForTest = server.URL

	client, err := NewClient(server.URL, "test-token")
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	resp, paginator, err := client.ListCustomersWithPagination(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify page 1 data
	var page1 []map[string]any
	json.Unmarshal(resp.Body, &page1)
	if len(page1) != 1 || page1[0]["id"] != float64(1) {
		t.Fatalf("expected page 1 customer, got %v", page1)
	}

	if !paginator.HasNext() {
		t.Fatal("expected paginator to have next page")
	}

	// Fetch page 2
	next, err := paginator.NextPage(context.Background())
	if err != nil {
		t.Fatalf("unexpected error on NextPage: %v", err)
	}

	page2, ok := next.([]any)
	if !ok {
		t.Fatalf("expected []any from NextPage, got %T", next)
	}
	if len(page2) != 1 || page2[0].(map[string]any)["id"] != float64(2) {
		t.Fatalf("expected page 2 customer (id=2), got %v", page2)
	}
}
