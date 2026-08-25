package wenmar

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

type Client struct {
	BaseURL string
	Token   string
	http    *http.Client
	gen     *generated.ClientWithResponses
}

func NewClient(baseURL, token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("API token is required")
	}

	httpClient := &http.Client{
		Transport: newRetryTransport(),
	}

	gen, err := generated.NewClientWithResponses(baseURL,
		generated.WithHTTPClient(httpClient),
		generated.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
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

func (c *Client) ListCustomers(ctx context.Context, page *int) (*generated.ListCustomersResponse, error) {
	params := &generated.ListCustomersParams{}
	if page != nil {
		params.Page = page
	}
	resp, err := c.gen.ListCustomersWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, ParseErrorBody(resp.Body, resp.StatusCode())
	}
	return resp, nil
}

func (c *Client) ListCustomersWithPagination(ctx context.Context, page *int) (*generated.ListCustomersResponse, *Paginator, error) {
	resp, err := c.ListCustomers(ctx, page)
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
		return nil, ParseErrorBody(resp.Body, resp.StatusCode())
	}
	return resp, nil
}

func (c *Client) CreateCustomer(ctx context.Context, body generated.CreateCustomerJSONRequestBody) (*generated.CreateCustomerResponse, error) {
	resp, err := c.gen.CreateCustomerWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, ParseErrorBody(resp.Body, resp.StatusCode())
	}
	return resp, nil
}

func (c *Client) ShowVehicle(ctx context.Context, id int) (*generated.ShowVehicleResponse, error) {
	resp, err := c.gen.ShowVehicleWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, ParseErrorBody(resp.Body, resp.StatusCode())
	}
	return resp, nil
}

func (c *Client) ListWorkOrders(ctx context.Context, page *int) (*generated.ListWorkOrdersResponse, error) {
	params := &generated.ListWorkOrdersParams{}
	if page != nil {
		params.Page = page
	}
	resp, err := c.gen.ListWorkOrdersWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, ParseErrorBody(resp.Body, resp.StatusCode())
	}
	return resp, nil
}

func (c *Client) ListWorkOrdersWithPagination(ctx context.Context, page *int) (*generated.ListWorkOrdersResponse, *Paginator, error) {
	resp, err := c.ListWorkOrders(ctx, page)
	if err != nil {
		return nil, nil, err
	}
	paginator := newWorkOrdersPaginatorFromResponse(resp.HTTPResponse, c)
	return resp, paginator, nil
}

func (c *Client) ShowWorkOrder(ctx context.Context, id int) (*generated.ShowWorkOrderResponse, error) {
	resp, err := c.gen.ShowWorkOrderWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, ParseErrorBody(resp.Body, resp.StatusCode())
	}
	return resp, nil
}
