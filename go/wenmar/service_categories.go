package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// ListServiceCategoriesTyped returns a typed list of service categories.
// Service categories are not paginated (the index returns all categories
// ordered by position), so there is no Next closure.
func (c *Client) ListServiceCategoriesTyped(ctx context.Context) (*ListResult[ServiceCategory], error) {
	resp, err := c.ListServiceCategories(ctx)
	if err != nil {
		return nil, err
	}
	items, err := parseListResponse[ServiceCategory](resp.Body)
	if err != nil {
		return nil, err
	}
	meta, nextURL := extractPaginationMeta(resp.HTTPResponse)
	result := &ListResult[ServiceCategory]{
		Items: items,
		Meta:  meta,
	}
	if nextURL != "" {
		result.Next = func(ctx context.Context) (*ListResult[ServiceCategory], error) {
			return c.fetchNextPage[ServiceCategory](ctx, nextURL)
		}
	}
	return result, nil
}

// GetAllServiceCategories auto-paginates and returns all service categories,
// up to MaxItems (default 1000 safety cap).
func (c *Client) GetAllServiceCategories(ctx context.Context, opts *GetAllOptions) ([]ServiceCategory, bool, error) {
	if opts == nil {
		opts = &GetAllOptions{MaxItems: 1000}
	}
	if opts.MaxItems == 0 {
		opts.MaxItems = 1000
	}
	first, err := c.ListServiceCategoriesTyped(ctx)
	if err != nil {
		return nil, false, err
	}
	return getAll(ctx, first, opts)
}

// ListServiceCategories lists service categories.
func (c *Client) ListServiceCategories(ctx context.Context) (*ListServiceCategoriesResponse, error) {
	resp, err := c.gen.ListServiceCategoriesWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// CreateServiceCategory creates a service category.
func (c *Client) CreateServiceCategory(ctx context.Context, req CreateServiceCategoryRequest) (*CreateServiceCategoryResponse, error) {
	body := req.ToGenerated()
	resp, err := c.gen.CreateServiceCategoryWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// UpdateServiceCategory updates a service category.
func (c *Client) UpdateServiceCategory(ctx context.Context, id int, req UpdateServiceCategoryRequest) (*UpdateServiceCategoryResponse, error) {
	body := req.ToGenerated()
	resp, err := c.gen.UpdateServiceCategoryWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// DeleteServiceCategory deletes a service category. Returns 422 if the
// category has jobs assigned.
func (c *Client) DeleteServiceCategory(ctx context.Context, id int) (*DeleteServiceCategoryResponse, error) {
	resp, err := c.gen.DeleteServiceCategoryWithResponse(ctx, id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// DeactivateServiceCategory deactivates a service category (sets active=false).
func (c *Client) DeactivateServiceCategory(ctx context.Context, id int) (*DeactivateServiceCategoryResponse, error) {
	resp, err := c.gen.DeactivateServiceCategoryWithResponse(ctx, id, generated.DeactivateServiceCategoryJSONRequestBody{})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// ReactivateServiceCategory reactivates a service category (sets active=true).
func (c *Client) ReactivateServiceCategory(ctx context.Context, id int) (*ReactivateServiceCategoryResponse, error) {
	resp, err := c.gen.ReactivateServiceCategoryWithResponse(ctx, id, generated.ReactivateServiceCategoryJSONRequestBody{})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// MoveUpServiceCategory moves a service category up one position. Returns
// the full ordered list of categories.
func (c *Client) MoveUpServiceCategory(ctx context.Context, id int) (*MoveUpServiceCategoryResponse, error) {
	resp, err := c.gen.MoveUpServiceCategoryWithResponse(ctx, id, generated.MoveUpServiceCategoryJSONRequestBody{})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// MoveDownServiceCategory moves a service category down one position. Returns
// the full ordered list of categories.
func (c *Client) MoveDownServiceCategory(ctx context.Context, id int) (*MoveDownServiceCategoryResponse, error) {
	resp, err := c.gen.MoveDownServiceCategoryWithResponse(ctx, id, generated.MoveDownServiceCategoryJSONRequestBody{})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// SeedDefaultsServiceCategories seeds default service categories for the
// account. Returns a summary with the count of categories created.
func (c *Client) SeedDefaultsServiceCategories(ctx context.Context) (*SeedDefaultsServiceCategoriesResponse, error) {
	resp, err := c.gen.SeedDefaultsServiceCategoriesWithResponse(ctx, generated.SeedDefaultsServiceCategoriesJSONRequestBody{})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}