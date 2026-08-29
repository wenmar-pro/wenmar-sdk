package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// ListTags returns all customer and vehicle tags.
func (c *Client) ListTags(ctx context.Context) (*generated.ListTagsResponse, error) {
	resp, err := c.gen.ListTagsWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// UpdateTags does a bulk update/destroy of customer and vehicle tags.
func (c *Client) UpdateTags(ctx context.Context, body generated.UpdateTagsJSONRequestBody) (*generated.UpdateTagsResponse, error) {
	resp, err := c.gen.UpdateTagsWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// CreateCustomerTag creates a single customer tag.
func (c *Client) CreateCustomerTag(ctx context.Context, body generated.CreateCustomerTagJSONRequestBody) (*generated.CreateCustomerTagResponse, error) {
	resp, err := c.gen.CreateCustomerTagWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// CreateVehicleTag creates a single vehicle tag.
func (c *Client) CreateVehicleTag(ctx context.Context, body generated.CreateVehicleTagJSONRequestBody) (*generated.CreateVehicleTagResponse, error) {
	resp, err := c.gen.CreateVehicleTagWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
