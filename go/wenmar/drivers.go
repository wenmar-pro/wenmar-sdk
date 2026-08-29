package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// ListDriversTyped returns a typed, paginated list of drivers for a customer.
func (c *Client) ListDriversTyped(ctx context.Context, customerID int) (*ListResult[Driver], error) {
	resp, err := c.ListDrivers(ctx, customerID)
	if err != nil {
		return nil, err
	}
	items, err := parseListResponse[Driver](resp.Body)
	if err != nil {
		return nil, err
	}
	meta, nextURL := extractPaginationMeta(resp.HTTPResponse)
	result := &ListResult[Driver]{
		Items: items,
		Meta:  meta,
	}
	if nextURL != "" {
		result.Next = func(ctx context.Context) (*ListResult[Driver], error) {
			return c.fetchNextPage[Driver](ctx, nextURL)
		}
	}
	return result, nil
}

// GetAllDrivers auto-paginates and returns all drivers for a customer, up to
// MaxItems (default 1000 safety cap).
func (c *Client) GetAllDrivers(ctx context.Context, customerID int, opts *GetAllOptions) ([]Driver, bool, error) {
	if opts == nil {
		opts = &GetAllOptions{MaxItems: 1000}
	}
	if opts.MaxItems == 0 {
		opts.MaxItems = 1000
	}
	first, err := c.ListDriversTyped(ctx, customerID)
	if err != nil {
		return nil, false, err
	}
	return getAll(ctx, first, opts)
}

// ListDrivers lists drivers for a customer.
func (c *Client) ListDrivers(ctx context.Context, customerID int) (*ListCustomersDriversResponse, error) {
	resp, err := c.gen.ListCustomersDriversWithResponse(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ShowDriver shows a single driver.
func (c *Client) ShowDriver(ctx context.Context, customerID, id int) (*ShowDriverResponse, error) {
	resp, err := c.gen.ShowDriverWithResponse(ctx, customerID, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// CreateDriverRequest is the hand-written input for creating a driver.
type CreateDriverRequest struct {
	FullName string
	Phone    string
}

// CreateDriver creates a new driver for a customer.
func (c *Client) CreateDriver(ctx context.Context, customerID int, req CreateDriverRequest) (*CreateDriverResponse, error) {
	body := generated.CreateDriverJSONRequestBody{
		Driver: struct {
			FullName string `json:"full_name"`
			Phone    string `json:"phone"`
		}{
			FullName: req.FullName,
			Phone:    req.Phone,
		},
	}
	resp, err := c.gen.CreateDriverWithResponse(ctx, customerID, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// UpdateDriverRequest is the hand-written input for updating a driver.
type UpdateDriverRequest struct {
	FullName string
}

// UpdateDriver updates a driver.
func (c *Client) UpdateDriver(ctx context.Context, customerID, id int, req UpdateDriverRequest) (*UpdateDriverResponse, error) {
	body := generated.UpdateDriverJSONRequestBody{
		Driver: struct {
			FullName string `json:"full_name"`
		}{
			FullName: req.FullName,
		},
	}
	resp, err := c.gen.UpdateDriverWithResponse(ctx, customerID, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// DeleteDriver deletes a driver.
func (c *Client) DeleteDriver(ctx context.Context, customerID, id int) (*DeleteDriverResponse, error) {
	resp, err := c.gen.DeleteDriverWithResponse(ctx, customerID, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
