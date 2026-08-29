package wenmar

import (
	"context"
)

// CreateVehicle creates a new vehicle.
func (c *Client) CreateVehicle(ctx context.Context, req CreateVehicleRequest) (*CreateVehicleResponse, error) {
	body := req.ToGenerated()
	resp, err := c.gen.CreateVehicleWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// UpdateVehicle updates a vehicle.
func (c *Client) UpdateVehicle(ctx context.Context, id int, req UpdateVehicleRequest) (*UpdateVehicleResponse, error) {
	body := req.ToGenerated()
	resp, err := c.gen.UpdateVehicleWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}