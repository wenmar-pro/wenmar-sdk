package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// CreateVehicle creates a new vehicle. Accepts the generated request body
// type directly — all fields are typed from the OpenAPI spec.
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

// UpdateVehicle updates a vehicle. Accepts the generated request body type
// directly.
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
