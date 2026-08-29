package wenmar

import (
	"context"
)

// ListCustomersTyped returns a typed, paginated list of customers.
func (c *Client) ListCustomersTyped(ctx context.Context) (*ListResult[Customer], error) {
	resp, err := c.ListCustomers(ctx)
	if err != nil {
		return nil, err
	}
	items, err := parseListResponse[Customer](resp.Body)
	if err != nil {
		return nil, err
	}
	meta, nextURL := extractPaginationMeta(resp.HTTPResponse)
	result := &ListResult[Customer]{
		Items: items,
		Meta:  meta,
	}
	if nextURL != "" {
		result.Next = func(ctx context.Context) (*ListResult[Customer], error) {
			return c.fetchNextPage[Customer](ctx, nextURL)
		}
	}
	return result, nil
}

// GetAllCustomers auto-paginates and returns all customers, up to
// MaxItems (default 1000 safety cap).
func (c *Client) GetAllCustomers(ctx context.Context, opts *GetAllOptions) ([]Customer, bool, error) {
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

// ListCustomersTypedWithParams returns a typed, paginated list of customers
// filtered by the given params (query, type, has_vehicle, has_balance,
// tag_ids) with pagination (page, per_page).
func (c *Client) ListCustomersTypedWithParams(ctx context.Context, params ListCustomersParams) (*ListResult[Customer], error) {
	resp, err := c.ListCustomersWithParams(ctx, params)
	if err != nil {
		return nil, err
	}
	items, err := parseListResponse[Customer](resp.Body)
	if err != nil {
		return nil, err
	}
	meta, nextURL := extractPaginationMeta(resp.HTTPResponse)
	result := &ListResult[Customer]{
		Items: items,
		Meta:  meta,
	}
	if nextURL != "" {
		result.Next = func(ctx context.Context) (*ListResult[Customer], error) {
			return c.fetchNextPage[Customer](ctx, nextURL)
		}
	}
	return result, nil
}

// GetAllCustomersWithParams auto-paginates filtered customers, up to
// MaxItems (default 1000 safety cap).
func (c *Client) GetAllCustomersWithParams(ctx context.Context, params ListCustomersParams, opts *GetAllOptions) ([]Customer, bool, error) {
	if opts == nil {
		opts = &GetAllOptions{MaxItems: 1000}
	}
	if opts.MaxItems == 0 {
		opts.MaxItems = 1000
	}
	first, err := c.ListCustomersTypedWithParams(ctx, params)
	if err != nil {
		return nil, false, err
	}
	return getAll(ctx, first, opts)
}

// CreateCustomer creates a new customer.
func (c *Client) CreateCustomer(ctx context.Context, req CreateCustomerRequest) (*CreateCustomerResponse, error) {
	body := req.ToGenerated()
	resp, err := c.gen.CreateCustomerWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// UpdateCustomer updates a customer.
func (c *Client) UpdateCustomer(ctx context.Context, id int, req UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	body := req.ToGenerated()
	resp, err := c.gen.UpdateCustomerWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}