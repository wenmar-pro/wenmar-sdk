package wenmar

// Nested attribute types used when building create/update request bodies.
// These mirror the generated request structs' nested fields so callers can
// construct requests without importing pkg/generated.

// EmailAttribute is a nested email for customer create.
type EmailAttribute struct {
	Email   string `json:"email"`
	Label   string `json:"label"`
	Primary bool   `json:"primary"`
}

// EmailUpdateAttribute is a nested email for customer update (with id).
type EmailUpdateAttribute struct {
	Email string  `json:"email"`
	Id    *int    `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
}

// PhoneAttribute is a nested phone for customer create.
type PhoneAttribute struct {
	Label   string `json:"label"`
	Number  string `json:"number"`
	Primary bool   `json:"primary"`
}

// PhoneUpdateAttribute is a nested phone for customer update.
type PhoneUpdateAttribute struct {
	UnderscoreDestroy *bool   `json:"_destroy,omitempty"`
	Id                *int    `json:"id,omitempty"`
	Label             *string `json:"label,omitempty"`
	Number            *string `json:"number,omitempty"`
	Primary           *bool   `json:"primary,omitempty"`
}

// AddressAttribute is a nested address for customer create/update.
type AddressAttribute struct {
	Address1   string `json:"address1"`
	City       string `json:"city"`
	Country    string `json:"country"`
	IsBilling  bool   `json:"is_billing"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}
