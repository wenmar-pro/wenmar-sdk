package wenmar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

type Client struct {
	BaseURL  string
	Token    string
	cfg      Config
	tp       TokenProvider
	http     *http.Client
	gen      *generated.ClientWithResponses
	location *locationHolder
	hooks    Hooks
}

// locationHolder is a shared pointer so a scoped LocationClient and its
// underlying Client inject the same X-Wenmar-Location header.
type locationHolder struct {
	id string
}

// NewClient creates a Wenmar API client from the given Config and
// TokenProvider. The Config is deep-copied so callers cannot mutate the
// client's configuration after construction.
func NewClient(cfg Config, tp TokenProvider, opts ...ClientOption) (*Client, error) {
	if tp == nil {
		return nil, fmt.Errorf("token provider is required")
	}

	// Deep-copy config to prevent post-construction mutation.
	cfgCopy := cfg
	if cfgCopy.BaseURL == "" {
		cfgCopy.BaseURL = DefaultConfig().BaseURL
	}
	if err := requireHTTPS(cfgCopy.BaseURL); err != nil {
		return nil, err
	}

	// Resolve the token once at construction (for fetchURL/pagination).
	// A TokenProvider that rotates tokens per-request is supported by
	// a request-editor wrapper, but the pagination path needs the token
	// eagerly. For now, StaticTokenProvider and CredentialStore both
	// return a stable token.
	token, err := tp.Token(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to resolve token: %w", err)
	}

	var transport http.RoundTripper = http.DefaultTransport
	if cfgCopy.MaxRetries > 0 {
		transport = newRetryTransportWithRetries(cfgCopy.MaxRetries)
	}
	if cfgCopy.CacheEnabled {
		transport = newCachingTransport(transport)
	}

	httpClient := cfgCopy.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Transport:     transport,
			Timeout:       cfgCopy.Timeout,
			CheckRedirect: stripAuthOnCrossOriginRedirect,
		}
	}

	loc := &locationHolder{}
	gen, err := generated.NewClientWithResponses(cfgCopy.BaseURL,
		generated.WithHTTPClient(httpClient),
		generated.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Accept", "application/json")
			req.Header.Set("User-Agent", fmt.Sprintf("wenmar-sdk-go/%s", Version))
			if loc.id != "" {
				req.Header.Set("X-Wenmar-Location", loc.id)
			}
			return nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	c := &Client{
		BaseURL:  cfgCopy.BaseURL,
		Token:    token,
		cfg:      cfgCopy,
		tp:       tp,
		http:     httpClient,
		gen:      gen,
		location: loc,
		hooks:    NoopHooks{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
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

// stripAuthOnCrossOriginRedirect is the CheckRedirect policy: it preserves
// all headers for same-origin redirects but drops the Authorization header
// when the redirect target has a different scheme or host, preventing
// credential leakage to third parties.
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

// parseError converts a failed response into an *APIError, capturing the
// request method and path for diagnostics.
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

func (c *Client) ListCustomers(ctx context.Context) (*generated.ListCustomersResponse, error) {
	ctx = c.hooks.OnOperationStart(ctx, OperationInfo{Operation: "ListCustomers"})
	resp, err := c.gen.ListCustomersWithResponse(ctx)
	if err != nil {
		c.hooks.OnOperationEnd(ctx, OperationInfo{Operation: "ListCustomers"}, OperationResult{Operation: "ListCustomers", Err: err})
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		perr := parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
		c.hooks.OnOperationEnd(ctx, OperationInfo{Operation: "ListCustomers"}, OperationResult{Operation: "ListCustomers", Err: perr})
		return nil, perr
	}
	c.hooks.OnOperationEnd(ctx, OperationInfo{Operation: "ListCustomers"}, OperationResult{Operation: "ListCustomers"})
	return resp, nil
}

func (c *Client) ListCustomersWithPagination(ctx context.Context) (*generated.ListCustomersResponse, *Paginator, error) {
	resp, err := c.ListCustomers(ctx)
	if err != nil {
		return nil, nil, err
	}
	paginator := newPaginatorFromResponse(resp.HTTPResponse, c)
	return resp, paginator, nil
}

func (c *Client) ShowCustomer(ctx context.Context, id int) (*generated.ShowCustomerResponse, error) {
	resp, err := c.gen.ShowCustomerWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) ListVehicles(ctx context.Context) (*generated.ListVehiclesResponse, error) {
	resp, err := c.gen.ListVehiclesWithResponse(ctx, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) ShowVehicle(ctx context.Context, id int) (*generated.ShowVehicleResponse, error) {
	resp, err := c.gen.ShowVehicleWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) DeleteVehicle(ctx context.Context, id int) (*generated.DeleteVehicleResponse, error) {
	resp, err := c.gen.DeleteVehicleWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) DecodeVin(ctx context.Context, vin string) (*generated.DecodeVinResponse, error) {
	params := &generated.DecodeVinParams{Vin: &vin}
	resp, err := c.gen.DecodeVinWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) ListWorkOrders(ctx context.Context) (*generated.ListWorkOrdersResponse, error) {
	resp, err := c.gen.ListWorkOrdersWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) ListWorkOrdersWithPagination(ctx context.Context) (*generated.ListWorkOrdersResponse, *Paginator, error) {
	resp, err := c.ListWorkOrders(ctx)
	if err != nil {
		return nil, nil, err
	}
	paginator := newPaginatorFromResponse(resp.HTTPResponse, c)
	return resp, paginator, nil
}

func (c *Client) ShowWorkOrder(ctx context.Context, id int) (*generated.ShowWorkOrderResponse, error) {
	resp, err := c.gen.ShowWorkOrderWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) DeleteWorkOrder(ctx context.Context, id int) (*generated.DeleteWorkOrderResponse, error) {
	resp, err := c.gen.DeleteWorkOrderWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) ListAccount(ctx context.Context) (*generated.ListAccountResponse, error) {
	resp, err := c.gen.ListAccountWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) ShowLocation(ctx context.Context, id string) (*generated.ShowLocationResponse, error) {
	resp, err := c.gen.ShowLocationWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
