package wenmar

import (
	"context"

	"github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"
)

// CreateCustomerRequest is the hand-written input for creating a customer.
// Callers use this instead of the generated CreateCustomerJSONRequestBody.
type CreateCustomerRequest struct {
	FirstName string
	LastName  string
}

// CreateCustomer creates a new customer.
func (c *Client) CreateCustomer(ctx context.Context, req CreateCustomerRequest) (*generated.CreateCustomerResponse, error) {
	body := generated.CreateCustomerJSONRequestBody{
		Customer: struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}{
			FirstName: req.FirstName,
			LastName:  req.LastName,
		},
	}
	resp, err := c.gen.CreateCustomerWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}

// UpdateCustomerRequest is the hand-written input for updating a customer.
type UpdateCustomerRequest struct {
	Emails []EmailInput
	Phones []PhoneInput
}

type EmailInput struct {
	Email string
	ID    *int
	Label *string
}

type PhoneInput struct {
	ID      int
	Destroy bool
}

// UpdateCustomer updates a customer.
func (c *Client) UpdateCustomer(ctx context.Context, id int, req UpdateCustomerRequest) (*generated.UpdateCustomerResponse, error) {
	body := generated.UpdateCustomerJSONRequestBody{
		Customer: struct {
			EmailsAttributes *[]struct {
				Email string  `json:"email"`
				Id    *int    `json:"id,omitempty"`
				Label *string `json:"label,omitempty"`
			} `json:"emails_attributes,omitempty"`
			PhonesAttributes *[]struct {
				UnderscoreDestroy bool `json:"_destroy"`
				Id                int  `json:"id"`
			} `json:"phones_attributes,omitempty"`
		}{},
	}
	if len(req.Emails) > 0 {
		emails := make([]struct {
			Email string  `json:"email"`
			Id    *int    `json:"id,omitempty"`
			Label *string `json:"label,omitempty"`
		}, len(req.Emails))
		for i, e := range req.Emails {
			emails[i] = struct {
				Email string  `json:"email"`
				Id    *int    `json:"id,omitempty"`
				Label *string `json:"label,omitempty"`
			}{Email: e.Email, Id: e.ID, Label: e.Label}
		}
		body.Customer.EmailsAttributes = &emails
	}
	if len(req.Phones) > 0 {
		phones := make([]struct {
			UnderscoreDestroy bool `json:"_destroy"`
			Id                int  `json:"id"`
		}, len(req.Phones))
		for i, p := range req.Phones {
			phones[i] = struct {
				UnderscoreDestroy bool `json:"_destroy"`
				Id                int  `json:"id"`
			}{UnderscoreDestroy: p.Destroy, Id: p.ID}
		}
		body.Customer.PhonesAttributes = &phones
	}

	resp, err := c.gen.UpdateCustomerWithResponse(ctx, id, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 400 {
		return nil, parseError(resp.Body, resp.StatusCode(), resp.HTTPResponse)
	}
	return resp, nil
}
