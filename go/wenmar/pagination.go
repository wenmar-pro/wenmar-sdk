package wenmar

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
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
