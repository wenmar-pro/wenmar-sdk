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
	BaseURL string
	Token   string
	http    *http.Client
	gen     *generated.ClientWithResponses
}

// NewClient creates a Wenmar API client. The baseURL must use https unless
// the host is localhost or 127.0.0.1 (for development). The token is sent as
// a Bearer header on every request.
func NewClient(baseURL, token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("API token is required")
	}
	if err := requireHTTPS(baseURL); err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: newCachingTransport(newRetryTransport()),
		CheckRedirect: stripAuthOnCrossOriginRedirect,
	}

	gen, err := generated.NewClientWithResponses(baseURL,
		generated.WithHTTPClient(httpClient),
		generated.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Accept", "application/json")
			req.Header.Set("User-Agent", "wenmar-sdk-go")
			return nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &Client{
		BaseURL: baseURL,
		Token:   token,
		http:    httpClient,
		gen:     gen,
	}, nil
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
	method, path := "", ""
	if hr != nil && hr.Request != nil {
		method = hr.Request.Method
		path = hr.Request.URL.Path
	}
	return ParseErrorBodyWithRequest(body, statusCode, method, path)
}

func (c *Client) ListCustomers(ctx context.Context) (*generated.ListCustomersResponse, error) {
	resp, err := c.gen.ListCustomersWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
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

func (c *Client) CreateCustomer(ctx context.Context, body generated.CreateCustomerJSONRequestBody) (*generated.CreateCustomerResponse, error) {
	resp, err := c.gen.CreateCustomerWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) UpdateCustomer(ctx context.Context, id int, body generated.UpdateCustomerJSONRequestBody) (*generated.UpdateCustomerResponse, error) {
	resp, err := c.gen.UpdateCustomerWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) ListVehicles(ctx context.Context) (*generated.ListVehiclesResponse, error) {
	resp, err := c.gen.ListVehiclesWithResponse(ctx)
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

func (c *Client) CreateVehicle(ctx context.Context, body generated.CreateVehicleJSONRequestBody) (*generated.CreateVehicleResponse, error) {
	resp, err := c.gen.CreateVehicleWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) UpdateVehicle(ctx context.Context, id int, body generated.UpdateVehicleJSONRequestBody) (*generated.UpdateVehicleResponse, error) {
	resp, err := c.gen.UpdateVehicleWithResponse(ctx, id, body)
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

func (c *Client) CheckDuplicate(ctx context.Context, vin string) (*generated.CheckDuplicateResponse, error) {
	params := &generated.CheckDuplicateParams{Vin: &vin}
	resp, err := c.gen.CheckDuplicateWithResponse(ctx, params)
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

func (c *Client) CreateWorkOrder(ctx context.Context, body generated.CreateWorkOrderJSONRequestBody) (*generated.CreateWorkOrderResponse, error) {
	resp, err := c.gen.CreateWorkOrderWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

func (c *Client) UpdateWorkOrder(ctx context.Context, id int, body generated.UpdateWorkOrderJSONRequestBody) (*generated.UpdateWorkOrderResponse, error) {
	resp, err := c.gen.UpdateWorkOrderWithResponse(ctx, id, body)
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
