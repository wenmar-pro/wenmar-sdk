package wenmar

import (
	"context"
	"net/http"
	"regexp"

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

type Paginator struct {
	nextURL string
	fetch   func(ctx context.Context) (any, error)
}

func (p *Paginator) HasNext() bool {
	return p.nextURL != ""
}

func (p *Paginator) NextPage(ctx context.Context) (any, error) {
	if !p.HasNext() {
		return nil, nil
	}
	resp, err := p.fetch(ctx)
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
	return &Paginator{nextURL: next, fetch: func(ctx context.Context) (any, error) {
		return client.ListCustomers(ctx)
	}}
}

func newWorkOrdersPaginatorFromResponse(resp *http.Response, client *Client) *Paginator {
	next := parseLinkHeader(resp.Header.Get("Link"), "next")
	return &Paginator{nextURL: next, fetch: func(ctx context.Context) (any, error) {
		return client.ListWorkOrders(ctx)
	}}
}
