// Package token defines the canonical bearer-token provider interface
// shared by the wenmar client and the auth package. It lives in a leaf
// package so the client does not inherit the auth package's keyring
// dependencies.
package token

import (
	"context"
	"fmt"
)

// Provider resolves a bearer token for a request. Implementations may
// refresh or rotate the token; the provider is called per request.
type Provider interface {
	Token(ctx context.Context) (string, error)
}

// Static is a Provider that returns a fixed token.
type Static struct{ value string }

// NewStatic creates a Provider that always returns the given token.
func NewStatic(value string) *Static { return &Static{value: value} }

func (p *Static) Token(context.Context) (string, error) {
	if p.value == "" {
		return "", fmt.Errorf("token is empty")
	}
	return p.value, nil
}
