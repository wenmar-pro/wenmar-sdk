package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// ListCustomerVehicles returns a customer's vehicles.
func (c *Client) ListCustomerVehicles(ctx context.Context, customerID int) (*generated.ListCustomersVehiclesResponse, error) {
	resp, err := c.gen.ListCustomersVehiclesWithResponse(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ListCustomerWorkOrders returns a customer's work orders.
func (c *Client) ListCustomerWorkOrders(ctx context.Context, customerID int) (*generated.ListCustomersWorkOrdersResponse, error) {
	resp, err := c.gen.ListCustomersWorkOrdersWithResponse(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ListVehicleWorkOrders returns a vehicle's work orders.
func (c *Client) ListVehicleWorkOrders(ctx context.Context, vehicleID int) (*generated.ListVehiclesWorkOrdersResponse, error) {
	resp, err := c.gen.ListVehiclesWorkOrdersWithResponse(ctx, vehicleID)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ListCustomerStatements returns a customer's statements.
func (c *Client) ListCustomerStatements(ctx context.Context, customerID int) (*generated.ListCustomersStatementsResponse, error) {
	resp, err := c.gen.ListCustomersStatementsWithResponse(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ListCustomerDrivers returns a customer's drivers.
func (c *Client) ListCustomerDrivers(ctx context.Context, customerID int) (*generated.ListCustomersDriversResponse, error) {
	resp, err := c.gen.ListCustomersDriversWithResponse(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
