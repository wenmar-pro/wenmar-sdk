package wenmar

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/bendangelo/wenmar-sdk/go/pkg/generated"
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
	client  *Client
}

func (p *Paginator) HasNext() bool {
	return p.nextURL != ""
}

func (p *Paginator) NextPage(ctx context.Context) (*generated.ListCustomersResponse, error) {
	if !p.HasNext() {
		return nil, nil
	}
	page := extractPageFromURL(p.nextURL)
	return p.client.ListCustomers(ctx, page)
}

func newPaginatorFromResponse(resp *http.Response, client *Client) *Paginator {
	next := parseLinkHeader(resp.Header.Get("Link"), "next")
	return &Paginator{nextURL: next, client: client}
}
