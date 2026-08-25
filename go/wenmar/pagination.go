package wenmar

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

var linkRelRE = regexp.MustCompile(`<([^>]+)>;\s*rel="([^"]+)"`)

func parseLinkHeader(header, rel string) string {
	if header == "" {
		return ""
	}
	matches := linkRelRE.FindAllStringSubmatch(header, -1)
	for _, m := range matches {
		if len(m) >= 3 && m[2] == rel {
			return m[1]
		}
	}
	return ""
}

func extractPageFromURL(url string) *int {
	idx := strings.Index(url, "page=")
	if idx == -1 {
		return nil
	}
	pageStr := url[idx+5:]
	if ampIdx := strings.Index(pageStr, "&"); ampIdx != -1 {
		pageStr = pageStr[:ampIdx]
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil
	}
	return &page
}

type Paginator struct {
	nextURL string
	fetch   func(ctx context.Context, page *int) (any, error)
}

func (p *Paginator) HasNext() bool {
	return p.nextURL != ""
}

func (p *Paginator) NextPage(ctx context.Context) (any, error) {
	if !p.HasNext() {
		return nil, nil
	}
	page := extractPageFromURL(p.nextURL)
	resp, err := p.fetch(ctx, page)
	if err != nil {
		return nil, err
	}
	// Advance to the next link from the response, if any.
	switch r := resp.(type) {
	case *generated.ListCustomersResponse:
		p.nextURL = parseLinkHeader(r.HTTPResponse.Header.Get("Link"), "next")
	case *generated.ListWorkOrdersResponse:
		p.nextURL = parseLinkHeader(r.HTTPResponse.Header.Get("Link"), "next")
	default:
		p.nextURL = ""
	}
	return resp, nil
}

func newPaginatorFromResponse(resp *http.Response, client *Client) *Paginator {
	next := parseLinkHeader(resp.Header.Get("Link"), "next")
	return &Paginator{nextURL: next, fetch: func(ctx context.Context, page *int) (any, error) {
		return client.ListCustomers(ctx, page)
	}}
}

func newWorkOrdersPaginatorFromResponse(resp *http.Response, client *Client) *Paginator {
	next := parseLinkHeader(resp.Header.Get("Link"), "next")
	return &Paginator{nextURL: next, fetch: func(ctx context.Context, page *int) (any, error) {
		return client.ListWorkOrders(ctx, page)
	}}
}
