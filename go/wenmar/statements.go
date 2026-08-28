package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// ListStatementsTyped returns a typed, paginated list of statements for a customer.
func (c *Client) ListStatementsTyped(ctx context.Context, customerID int) (*ListResult[generated.Statement], error) {
	resp, err := c.ListStatements(ctx, customerID)
	if err != nil {
		return nil, err
	}
	items, err := parseListResponse[generated.Statement](resp.Body)
	if err != nil {
		return nil, err
	}
	meta, nextURL := extractPaginationMeta(resp.HTTPResponse)
	result := &ListResult[generated.Statement]{
		Items: items,
		Meta:  meta,
	}
	if nextURL != "" {
		result.Next = func(ctx context.Context) (*ListResult[generated.Statement], error) {
			return c.fetchNextPage[generated.Statement](ctx, nextURL)
		}
	}
	return result, nil
}

// GetAllStatements auto-paginates and returns all statements for a customer,
// up to MaxItems (default 1000 safety cap).
func (c *Client) GetAllStatements(ctx context.Context, customerID int, opts *GetAllOptions) ([]generated.Statement, bool, error) {
	if opts == nil {
		opts = &GetAllOptions{MaxItems: 1000}
	}
	if opts.MaxItems == 0 {
		opts.MaxItems = 1000
	}
	first, err := c.ListStatementsTyped(ctx, customerID)
	if err != nil {
		return nil, false, err
	}
	return getAll(ctx, first, opts)
}

// ListStatements lists statements for a customer.
func (c *Client) ListStatements(ctx context.Context, customerID int) (*generated.ListCustomersStatementsResponse, error) {
	resp, err := c.gen.ListCustomersStatementsWithResponse(ctx, customerID)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ShowStatement shows a single statement.
func (c *Client) ShowStatement(ctx context.Context, id int) (*generated.ShowStatementResponse, error) {
	resp, err := c.gen.ShowStatementWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
