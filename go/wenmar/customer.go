package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// ListCustomersTyped returns a typed, paginated list of customers.
func (c *Client) ListCustomersTyped(ctx context.Context) (*ListResult[generated.Customer], error) {
	resp, err := c.ListCustomers(ctx)
	if err != nil {
		return nil, err
	}
	items, err := parseListResponse[generated.Customer](resp.Body)
	if err != nil {
		return nil, err
	}
	meta, nextURL := extractPaginationMeta(resp.HTTPResponse)
	result := &ListResult[generated.Customer]{
		Items: items,
		Meta:  meta,
	}
	if nextURL != "" {
		result.Next = func(ctx context.Context) (*ListResult[generated.Customer], error) {
			return c.fetchNextPage[generated.Customer](ctx, nextURL)
		}
	}
	return result, nil
}

// GetAllCustomers auto-paginates and returns all customers, up to
// MaxItems (default 1000 safety cap).
func (c *Client) GetAllCustomers(ctx context.Context, opts *GetAllOptions) ([]generated.Customer, bool, error) {
	if opts == nil {
		opts = &GetAllOptions{MaxItems: 1000}
	}
	if opts.MaxItems == 0 {
		opts.MaxItems = 1000
	}
	first, err := c.ListCustomersTyped(ctx)
	if err != nil {
		return nil, false, err
	}
	return getAll(ctx, first, opts)
}

// CreateCustomer creates a new customer. Accepts the generated request body
// type directly — all fields are typed from the OpenAPI spec.
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

// UpdateCustomer updates a customer. Accepts the generated request body type
// directly.
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
