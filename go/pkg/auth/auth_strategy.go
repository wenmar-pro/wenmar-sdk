package auth

import (
	"fmt"
	"net/http"
)

// AuthStrategy authenticates an outgoing HTTP request. This replaces the
// hardcoded request editor in the client.
type AuthStrategy interface {
	Authenticate(req *http.Request) error
}

// BearerAuth sets the Authorization header to "Bearer <token>" using a
// TokenProvider.
type BearerAuth struct {
	Provider TokenProvider
}

// Authenticate sets the Authorization header on the request.
func (b *BearerAuth) Authenticate(req *http.Request) error {
	if b.Provider == nil {
		return fmt.Errorf("bearer auth requires a token provider")
	}
	token, err := b.Provider.Token(req.Context())
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}
