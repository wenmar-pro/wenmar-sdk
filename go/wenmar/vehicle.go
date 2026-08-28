package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// CreateVehicleRequest is the hand-written input for creating a vehicle.
type CreateVehicleRequest struct {
	CustomerID int
	Make       string
	Model      string
	Year       int
}

// CreateVehicle creates a new vehicle.
func (c *Client) CreateVehicle(ctx context.Context, req CreateVehicleRequest) (*generated.CreateVehicleResponse, error) {
	body := generated.CreateVehicleJSONRequestBody{
		Vehicle: struct {
			CustomerId int    `json:"customer_id"`
			Make       string `json:"make"`
			Model      string `json:"model"`
			Year       int    `json:"year"`
		}{
			CustomerId: req.CustomerID,
			Make:       req.Make,
			Model:      req.Model,
			Year:       req.Year,
		},
	}
	resp, err := c.gen.CreateVehicleWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// UpdateVehicleRequest is the hand-written input for updating a vehicle.
type UpdateVehicleRequest struct {
	Make string
}

// UpdateVehicle updates a vehicle.
func (c *Client) UpdateVehicle(ctx context.Context, id int, req UpdateVehicleRequest) (*generated.UpdateVehicleResponse, error) {
	body := generated.UpdateVehicleJSONRequestBody{
		Vehicle: struct {
			Make string `json:"make"`
		}{
			Make: req.Make,
		},
	}
	resp, err := c.gen.UpdateVehicleWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
