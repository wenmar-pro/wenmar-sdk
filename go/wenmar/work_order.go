package wenmar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// CreateWorkOrderRequest is the hand-written input for creating a work order.
type CreateWorkOrderRequest struct {
	CustomerID int
	VehicleID  int
}

// CreateWorkOrder creates a new work order.
func (c *Client) CreateWorkOrder(ctx context.Context, req CreateWorkOrderRequest) (*CreateWorkOrderResponse, error) {
	body := generated.CreateWorkOrderJSONRequestBody{
		WorkOrder: struct {
			CustomerId int `json:"customer_id"`
			VehicleId  int `json:"vehicle_id"`
		}{
			CustomerId: req.CustomerID,
			VehicleId:  req.VehicleID,
		},
	}
	resp, err := c.gen.CreateWorkOrderWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// UpdateWorkOrderRequest is the hand-written input for updating a work order.
type UpdateWorkOrderRequest struct {
	IntakeMethod string
}

// UpdateWorkOrder updates a work order.
func (c *Client) UpdateWorkOrder(ctx context.Context, id int, req UpdateWorkOrderRequest) (*UpdateWorkOrderResponse, error) {
	body := generated.UpdateWorkOrderJSONRequestBody{
		WorkOrder: struct {
			IntakeMethod string `json:"intake_method"`
		}{
			IntakeMethod: req.IntakeMethod,
		},
	}
	resp, err := c.gen.UpdateWorkOrderWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// lifecycleAction performs a raw PATCH against a work order lifecycle route.
// The lifecycle routes are not yet in the generated client (the spec does not
// include them); they are added here as raw calls until the spec is updated
// (Track D). TODO(Track D): replace with generated client calls once the
// lifecycle routes are captured in the spec.
func (c *Client) lifecycleAction(ctx context.Context, id int, action string, body map[string]any) error {
	url := fmt.Sprintf("%s/work_orders/%d/lifecycle/%s", c.BaseURL, id, action)
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "wenmar-sdk-go")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return parseError(bodyBytes, resp.StatusCode, resp)
	}
	return nil
}

// StageTransition moves a work order to a new stage.
func (c *Client) StageTransition(ctx context.Context, id int, status string) error {
	return c.lifecycleAction(ctx, id, "stage_transition", map[string]any{"status": status})
}

// Close closes a work order.
func (c *Client) Close(ctx context.Context, id int) error {
	return c.lifecycleAction(ctx, id, "close", nil)
}

// CloseAsPaid closes a work order as paid.
func (c *Client) CloseAsPaid(ctx context.Context, id int) error {
	return c.lifecycleAction(ctx, id, "close_as_paid", nil)
}

// Reopen reopens a closed work order.
func (c *Client) Reopen(ctx context.Context, id int) error {
	return c.lifecycleAction(ctx, id, "reopen", nil)
}

// ShowWorkOrderEstimate returns the estimate tab (services) for a work order.
func (c *Client) ShowWorkOrderEstimate(ctx context.Context, id int) (*ShowWorkOrderEstimateResponse, error) {
	resp, err := c.gen.ShowWorkOrderEstimateWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ShowWorkOrderWip returns the work-in-progress tab (services) for a work order.
func (c *Client) ShowWorkOrderWip(ctx context.Context, id int) (*ShowWorkOrderWipResponse, error) {
	resp, err := c.gen.ShowWorkOrderWipWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ShowWorkOrderInspection returns the inspection tab (inspection reports) for a work order.
func (c *Client) ShowWorkOrderInspection(ctx context.Context, id int) (*ShowWorkOrderInspectionResponse, error) {
	resp, err := c.gen.ShowWorkOrderInspectionWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ShowWorkOrderParts returns the parts tab (services) for a work order.
func (c *Client) ShowWorkOrderParts(ctx context.Context, id int) (*ShowWorkOrderPartsResponse, error) {
	resp, err := c.gen.ShowWorkOrderPartsWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ShowWorkOrderPayments returns the payments tab (payments) for a work order.
func (c *Client) ShowWorkOrderPayments(ctx context.Context, id int) (*ShowWorkOrderPaymentsResponse, error) {
	resp, err := c.gen.ShowWorkOrderPaymentsWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// CreateWorkOrderPaymentRequest is the hand-written input for creating a payment.
type CreateWorkOrderPaymentRequest struct {
	AmountCents string
	Method      string
}

// CreateWorkOrderPayment creates a payment against a work order.
func (c *Client) CreateWorkOrderPayment(ctx context.Context, id int, req CreateWorkOrderPaymentRequest) (*CreateWorkOrderPaymentResponse, error) {
	body := generated.CreateWorkOrderPaymentJSONRequestBody{
		Payment: struct {
			AmountCents string `json:"amount_cents"`
			Method      string `json:"method"`
		}{
			AmountCents: req.AmountCents,
			Method:      req.Method,
		},
	}
	resp, err := c.gen.CreateWorkOrderPaymentWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
