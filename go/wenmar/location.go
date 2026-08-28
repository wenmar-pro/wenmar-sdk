package wenmar

import (
	"context"
	"fmt"
)

// LocationClient scopes all requests to a specific location via the
// X-Wenmar-Location header. All Client methods are available on
// LocationClient via embedding.
type LocationClient struct {
	*Client
	locationID string
}

// ForLocation verifies access to the given location and returns a client
// scoped to it. The returned client injects the X-Wenmar-Location header on
// every request.
func (c *Client) ForLocation(ctx context.Context, locationID string) (*LocationClient, error) {
	if locationID == "" {
		return nil, fmt.Errorf("location ID is required")
	}

	// Verify access to the location before returning a scoped client.
	if _, err := c.ShowLocation(ctx, locationID); err != nil {
		return nil, fmt.Errorf("could not access location %q: %w", locationID, err)
	}

	// Set the shared location holder so the request editor injects the header.
	c.location.id = locationID

	return &LocationClient{Client: c, locationID: locationID}, nil
}

// LocationID returns the location this client is scoped to.
func (lc *LocationClient) LocationID() string {
	return lc.locationID
}
