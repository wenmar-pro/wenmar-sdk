package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// TransferVehicle transfers a vehicle to a new customer.
func (c *Client) TransferVehicle(ctx context.Context, id int, body generated.TransferVehicleJSONRequestBody) (*generated.TransferVehicleResponse, error) {
	resp, err := c.gen.TransferVehicleWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// MergeVehicle merges a source vehicle into the keeper (id).
func (c *Client) MergeVehicle(ctx context.Context, id int, body generated.MergeVehicleJSONRequestBody) (*generated.MergeVehicleResponse, error) {
	resp, err := c.gen.MergeVehicleWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// PrefillVehicle decodes VIN/plate/year-make-model to prefill a vehicle form.
func (c *Client) PrefillVehicle(ctx context.Context, params generated.PrefillVehicleParams) (*generated.PrefillVehicleResponse, error) {
	resp, err := c.gen.PrefillVehicleWithResponse(ctx, &params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// LookupVehicle searches vehicles by query.
func (c *Client) LookupVehicle(ctx context.Context, query string) (*generated.LookupVehicleResponse, error) {
	resp, err := c.gen.LookupVehicleWithResponse(ctx, &generated.LookupVehicleParams{Query: &query})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// CheckVehicleDuplicate checks for duplicate vehicles by VIN/plate.
func (c *Client) CheckVehicleDuplicate(ctx context.Context, params generated.CheckVehicleDuplicateParams) (*generated.CheckVehicleDuplicateResponse, error) {
	resp, err := c.gen.CheckVehicleDuplicateWithResponse(ctx, &params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
