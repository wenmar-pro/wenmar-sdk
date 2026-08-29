package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// MergeCustomer merges a source customer into the keeper (id).
func (c *Client) MergeCustomer(ctx context.Context, id int, req MergeCustomerRequest) (*MergeCustomerResponse, error) {
	body := req.ToGenerated()
	resp, err := c.gen.MergeCustomerWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// LookupCustomer searches customers by query.
func (c *Client) LookupCustomer(ctx context.Context, query string) (*LookupCustomerResponse, error) {
	resp, err := c.gen.LookupCustomerWithResponse(ctx, &generated.LookupCustomerParams{Query: &query})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// CheckCustomerDuplicate checks for duplicate customers.
func (c *Client) CheckCustomerDuplicate(ctx context.Context, params CheckCustomerDuplicateParams) (*CheckCustomerDuplicateResponse, error) {
	resp, err := c.gen.CheckCustomerDuplicateWithResponse(ctx, &params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}