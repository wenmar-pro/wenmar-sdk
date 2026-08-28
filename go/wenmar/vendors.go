package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// ListVendorsTyped returns a typed, paginated list of vendors.
func (c *Client) ListVendorsTyped(ctx context.Context) (*ListResult[generated.Vendor], error) {
	resp, err := c.ListVendors(ctx)
	if err != nil {
		return nil, err
	}
	items, err := parseListResponse[generated.Vendor](resp.Body)
	if err != nil {
		return nil, err
	}
	meta, nextURL := extractPaginationMeta(resp.HTTPResponse)
	result := &ListResult[generated.Vendor]{
		Items: items,
		Meta:  meta,
	}
	if nextURL != "" {
		result.Next = func(ctx context.Context) (*ListResult[generated.Vendor], error) {
			return c.fetchNextPage[generated.Vendor](ctx, nextURL)
		}
	}
	return result, nil
}

// GetAllVendors auto-paginates and returns all vendors, up to MaxItems
// (default 1000 safety cap).
func (c *Client) GetAllVendors(ctx context.Context, opts *GetAllOptions) ([]generated.Vendor, bool, error) {
	if opts == nil {
		opts = &GetAllOptions{MaxItems: 1000}
	}
	if opts.MaxItems == 0 {
		opts.MaxItems = 1000
	}
	first, err := c.ListVendorsTyped(ctx)
	if err != nil {
		return nil, false, err
	}
	return getAll(ctx, first, opts)
}

// ListVendors lists vendors.
func (c *Client) ListVendors(ctx context.Context) (*generated.ListVendorsResponse, error) {
	resp, err := c.gen.ListVendorsWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ShowVendor shows a single vendor.
func (c *Client) ShowVendor(ctx context.Context, id int) (*generated.ShowVendorResponse, error) {
	resp, err := c.gen.ShowVendorWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
