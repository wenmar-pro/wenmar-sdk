package wenmar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	gen "github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// TokenProvider resolves a bearer token for a request. Implementations may
// refresh or rotate the token; the provider is called per request.
type TokenProvider interface {
	Token(ctx context.Context) (string, error)
}

type staticToken string

func (s staticToken) Token(context.Context) (string, error) { return string(s), nil }

// StaticTokenProvider returns a fixed token. Suitable for simple scripts and
// tests.
type StaticTokenProvider struct {
	token string
}

// NewStaticTokenProvider creates a TokenProvider that always returns the
// given token.
func NewStaticTokenProvider(token string) *StaticTokenProvider {
	return &StaticTokenProvider{token: token}
}

func (p *StaticTokenProvider) Token(context.Context) (string, error) {
	if p.token == "" {
		return "", fmt.Errorf("token is empty")
	}
	return p.token, nil
}

// Client is the hand-written SDK entry point. All 76 operations are generated
// into operations.gen.go and share the core request pipeline defined here.
type Client struct {
	BaseURL  string
	cfg      Config
	tp       TokenProvider
	http     *http.Client
	gen      *gen.ClientWithResponses
	location string
	hooks    Hooks
}

// NewClient creates a Wenmar API client from the given Config and
// TokenProvider. The Config is deep-copied so callers cannot mutate the
// client's configuration after construction.
func NewClient(cfg Config, tp TokenProvider, opts ...ClientOption) (*Client, error) {
	if tp == nil {
		return nil, fmt.Errorf("token provider is required")
	}

	cfgCopy := cfg
	if cfgCopy.BaseURL == "" {
		cfgCopy.BaseURL = DefaultConfig().BaseURL
	}
	if err := requireHTTPS(cfgCopy.BaseURL); err != nil {
		return nil, err
	}
	cfgCopy.BaseURL = strings.TrimSuffix(cfgCopy.BaseURL, "/")

	httpClient := cfgCopy.HTTPClient
	if httpClient == nil {
		// Build the transport stack bottom-up: caching -> retry -> base.
		var transport http.RoundTripper = http.DefaultTransport
		if cfgCopy.CacheEnabled {
			transport = newCachingTransport(transport)
		}
		if cfgCopy.MaxRetries > 0 {
			transport = newRetryTransportWithRetries(cfgCopy.MaxRetries, transport)
		}
		httpClient = &http.Client{
			Transport:     transport,
			Timeout:       cfgCopy.Timeout,
			CheckRedirect: stripAuthOnCrossOriginRedirect,
		}
	}

	c := &Client{
		BaseURL: cfgCopy.BaseURL,
		cfg:     cfgCopy,
		tp:      tp,
		http:    httpClient,
		hooks:   NoopHooks{},
	}
	if cfgCopy.Hooks != nil {
		c.hooks = cfgCopy.Hooks
	}
	for _, opt := range opts {
		opt(c)
	}

	genClient, err := gen.NewClientWithResponses(cfgCopy.BaseURL,
		gen.WithHTTPClient(httpClient),
		gen.WithRequestEditorFn(c.requestEditor),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	c.gen = genClient
	return c, nil
}

// requestEditor is called by the generated client before every request. It
// resolves the token per request and sets common headers, including the
// X-Wenmar-Location header when the client is location-scoped.
func (c *Client) requestEditor(ctx context.Context, req *http.Request) error {
	token, err := c.tp.Token(ctx)
	if err != nil {
		return fmt.Errorf("token provider: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "wenmar-sdk-go/"+Version)
	if c.location != "" {
		req.Header.Set("X-Wenmar-Location", c.location)
	}
	return nil
}

// requireHTTPS rejects non-https URLs unless the host is localhost or 127.0.0.1.
func requireHTTPS(baseURL string) error {
	parsed, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid base URL: %w", err)
	}
	if strings.EqualFold(parsed.Scheme, "https") {
		return nil
	}
	host := parsed.Hostname()
	if host == "localhost" || host == "127.0.0.1" {
		return nil
	}
	return fmt.Errorf("base URL must use https (got %q); http is only allowed for localhost", baseURL)
}

// stripAuthOnCrossOriginRedirect drops the Authorization header when a
// redirect target has a different scheme or host, preventing credential
// leakage to third parties.
func stripAuthOnCrossOriginRedirect(req *http.Request, via []*http.Request) error {
	if len(via) == 0 {
		return nil
	}
	original := via[0]
	if !sameOrigin(req.URL.String(), original.URL.String()) {
		req.Header.Del("Authorization")
	}
	return nil
}

// ForLocation returns a scoped client that injects X-Wenmar-Location on every
// request. The parent client is not mutated.
func (c *Client) ForLocation(locationID string) *Client {
	if locationID == "" {
		return c
	}
	child := *c
	child.location = locationID
	// Re-create the generated client so the request editor closure reads the
	// child's location value.
	genClient, err := gen.NewClientWithResponses(child.BaseURL,
		gen.WithHTTPClient(child.http),
		gen.WithRequestEditorFn(child.requestEditor),
	)
	if err == nil {
		child.gen = genClient
	}
	return &child
}

// parseError converts a failed generated response into an *APIError using the
// already-drained body bytes (oapi-codegen drains the body into resp.Body).
func parseError(body []byte, statusCode int, hr *http.Response) error {
	method, path, requestID := "", "", ""
	if hr != nil && hr.Request != nil {
		method = hr.Request.Method
		path = hr.Request.URL.Path
	}
	if hr != nil {
		requestID = hr.Header.Get("X-Request-Id")
	}
	return ParseErrorBodyWithRequestAndID(body, statusCode, method, path, requestID)
}

// FetchPage fetches a pagination URL (from a Link header) through the same
// transport stack and same-origin validation as normal requests. It returns
// the raw body and the next Link header. It is used by conformance tests and
// advanced callers that follow Link headers manually.
func (c *Client) FetchPage(ctx context.Context, url string) ([]byte, string, error) {
	return c.fetchURL(ctx, url)
}

// PaginatorFromResponse builds a Paginator from a list response's Link header.
// It lets callers follow Link-header pages manually from a raw list response.
func (c *Client) PaginatorFromResponse(resp *http.Response) *Paginator {
	return newPaginatorFromResponse(resp, c)
}

// collectAll follows the Link header from a first page body, appending items
func collectAll[T any](ctx context.Context, c *Client, body []byte, link string, max int) ([]T, error) {
	items, err := parseListResponse[T](body)
	if err != nil {
		return nil, err
	}
	if max > 0 && len(items) >= max {
		return items[:max], nil
	}
	nextURL := parseLinkHeader(link, "next")
	for nextURL != "" {
		nextBody, nextLink, err := c.fetchURL(ctx, nextURL)
		if err != nil {
			return nil, err
		}
		nextItems, err := parseListResponse[T](nextBody)
		if err != nil {
			return nil, err
		}
		if max > 0 && len(items)+len(nextItems) > max {
			nextItems = nextItems[:max-len(items)]
		}
		items = append(items, nextItems...)
		if max > 0 && len(items) >= max {
			return items, nil
		}
		nextURL = parseLinkHeader(nextLink, "next")
	}
	return items, nil
}
