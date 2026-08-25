package wenmar

import (
	"net/http/httptest"
	"testing"
)

func TestParseLinkHeader_Next(t *testing.T) {
	header := `<https://api.example.com/api/customers?page=2>; rel="next"`
	next := parseLinkHeader(header, "next")
	if next != "https://api.example.com/api/customers?page=2" {
		t.Errorf("expected next URL, got '%s'", next)
	}
}

func TestParseLinkHeader_Prev(t *testing.T) {
	header := `<https://api.example.com/api/customers?page=1>; rel="prev", <https://api.example.com/api/customers?page=3>; rel="next"`
	prev := parseLinkHeader(header, "prev")
	if prev != "https://api.example.com/api/customers?page=1" {
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
	p := &Paginator{nextURL: "https://api.example.com/api/customers?page=2"}
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
	resp.Header().Set("Link", `<https://api.example.com/api/customers?page=2>; rel="next"`)
	r := resp.Result()

	p := newPaginatorFromResponse(r, nil)
	if !p.HasNext() {
		t.Error("expected paginator to have next")
	}
	if p.nextURL != "https://api.example.com/api/customers?page=2" {
		t.Errorf("expected next URL, got '%s'", p.nextURL)
	}
}
