package wenmar

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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
	client  *Client
	// fetchNext is called with the raw next URL to fetch the next page.
	// It returns the response body and the Link header from the next response.
	fetchNext func(ctx context.Context, url string) (body []byte, linkHeader string, err error)
}

// PaginationMeta holds metadata about a paginated list result.
type PaginationMeta struct {
	HasMore    bool
	TotalCount int // from X-Total-Count header if available, 0 if unknown
	PerPage    int // from X-Per-Page header if available, 0 if unknown
}

// ListResult is a typed paginated result. Items is the current page's data.
// Next returns a new ListResult for the next page, or nil if there is no
// next page.
type ListResult[T any] struct {
	Items []T
	Meta  PaginationMeta
	Next  func(ctx context.Context) (*ListResult[T], error)
}

// HasMore reports whether there is a next page.
func (r *ListResult[T]) HasNext() bool {
	return r.Next != nil
}

// GetAllOptions controls auto-pagination behavior.
type GetAllOptions struct {
	MaxItems int // Stop after collecting this many items (0 = no cap)
	MaxPages int // Stop after this many pages (0 = no cap)
}

// getAll auto-paginates and collects all items into a single slice.
// If MaxItems or MaxPages is hit, truncated is true.
func getAll[T any](ctx context.Context, first *ListResult[T], opts *GetAllOptions) (items []T, truncated bool, err error) {
	if opts == nil {
		opts = &GetAllOptions{}
	}
	items = append(items, first.Items...)
	pages := 1
	current := first
	for current.Next != nil {
		if opts.MaxPages > 0 && pages >= opts.MaxPages {
			return items, true, nil
		}
		if opts.MaxItems > 0 && len(items) >= opts.MaxItems {
			return items[:opts.MaxItems], true, nil
		}
		current, err = current.Next(ctx)
		if err != nil {
			return items, false, err
		}
		items = append(items, current.Items...)
		pages++
	}
	return items, false, nil
}

// parseListResponse decodes a JSON array into a slice of T.
func parseListResponse[T any](body []byte) ([]T, error) {
	var items []T
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// extractPaginationMeta reads X-Total-Count, X-Per-Page, and the Link header.
func extractPaginationMeta(resp *http.Response) (PaginationMeta, string) {
	meta := PaginationMeta{}
	if v := resp.Header.Get("X-Total-Count"); v != "" {
		meta.TotalCount, _ = strconv.Atoi(v)
	}
	if v := resp.Header.Get("X-Per-Page"); v != "" {
		meta.PerPage, _ = strconv.Atoi(v)
	}
	nextURL := parseLinkHeader(resp.Header.Get("Link"), "next")
	meta.HasMore = nextURL != ""
	return meta, nextURL
}

// fetchNextPage fetches the next page (same-origin validated by fetchURL)
// and decodes it into a typed ListResult.
func (c *Client) fetchNextPage[T any](ctx context.Context, url string) (*ListResult[T], error) {
	body, linkHeader, err := c.fetchURL(ctx, url)
	if err != nil {
		return nil, err
	}
	items, err := parseListResponse[T](body)
	if err != nil {
		return nil, err
	}
	nextURL := parseLinkHeader(linkHeader, "next")
	result := &ListResult[T]{
		Items: items,
		Meta:  PaginationMeta{HasMore: nextURL != ""},
	}
	if nextURL != "" {
		result.Next = func(ctx context.Context) (*ListResult[T], error) {
			return c.fetchNextPage[T](ctx, nextURL)
		}
	}
	return result, nil
}

func (p *Paginator) HasNext() bool {
	return p.nextURL != ""
}

// NextPage fetches the next page by following the actual next URL from the
// Link header. The response is a generic decoded value (array or object).
func (p *Paginator) NextPage(ctx context.Context) (any, error) {
	if !p.HasNext() {
		return nil, nil
	}

	body, linkHeader, err := p.fetchNext(ctx, p.nextURL)
	if err != nil {
		return nil, err
	}

	// Advance to the next link from the response, if any.
	p.nextURL = parseLinkHeader(linkHeader, "next")

	// Decode the body as a generic value for the caller.
	var result any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func newPaginatorFromResponse(resp *http.Response, client *Client) *Paginator {
	next := parseLinkHeader(resp.Header.Get("Link"), "next")
	return &Paginator{
		nextURL: next,
		client:  client,
		fetchNext: func(ctx context.Context, url string) ([]byte, string, error) {
			return client.fetchURL(ctx, url)
		},
	}
}

// sameOrigin reports whether rawURL has the same scheme, host, and port
// as baseURL. This prevents the SDK from sending credentials to a
// different host if the API returns a malicious Link header.
func sameOrigin(rawURL, baseURL string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return false
	}
	return strings.EqualFold(parsed.Scheme, base.Scheme) &&
		strings.EqualFold(parsed.Host, base.Host)
}

// fetchURL performs a raw GET against the given URL (which may include query
// params like ?page=2) and returns the response body and Link header.
// It validates the URL is same-origin as the client's BaseURL before
// attaching credentials, preventing token exfiltration via malicious
// Link headers.
func (c *Client) fetchURL(ctx context.Context, url string) ([]byte, string, error) {
	if !sameOrigin(url, c.BaseURL) {
		return nil, "", &APIError{
			Code:       "invalid_pagination",
			Message:    "pagination next URL is not same-origin as base URL",
			StatusCode: 0,
			Method:     "GET",
			Path:       url,
		}
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "wenmar-sdk-go")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, "", ParseErrorBodyWithRequest(body, resp.StatusCode, "GET", url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return body, resp.Header.Get("Link"), nil
}
