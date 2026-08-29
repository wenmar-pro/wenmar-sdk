package wenmar

import "github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"

// --- Customer requests ---

// EmailAttribute is a nested email for customer create/update.
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

// CreateCustomerRequest is the typed input for creating a customer.
type CreateCustomerRequest struct {
	FirstName        string
	LastName         string
	CompanyName      *string
	FleetIdentifier  *string
	BillingTerms     *string
	CreditLimitCents *string
	TaxExempt        *bool
	TaxExemptNumber  *string
	Notes            *string
	MarketingOptIn   *bool
	DiscountPercent  *string
	PoRequired      *bool
	TagIds           *[]interface{}
	Emails           *[]EmailAttribute
	Phones           *[]PhoneAttribute
	Addresses        *[]AddressAttribute
}

// ToGenerated converts to the generated request body type.
func (r CreateCustomerRequest) ToGenerated() generated.CreateCustomerJSONRequestBody {
	body := generated.CreateCustomerJSONRequestBody{
		Customer: struct {
			AddressesAttributes *[]struct {
				Address1   string `json:"address1"`
				City       string `json:"city"`
				Country    string `json:"country"`
				IsBilling  bool   `json:"is_billing"`
				PostalCode string `json:"postal_code"`
				State      string `json:"state"`
			} `json:"addresses_attributes,omitempty"`
			BillingTerms     *string `json:"billing_terms,omitempty"`
			CompanyName      *string `json:"company_name,omitempty"`
			CreditLimitCents *string `json:"credit_limit_cents,omitempty"`
			DiscountPercent  *string `json:"discount_percent,omitempty"`
			EmailsAttributes *[]struct {
				Email   string `json:"email"`
				Label   string `json:"label"`
				Primary bool   `json:"primary"`
			} `json:"emails_attributes,omitempty"`
			FirstName        string  `json:"first_name"`
			FleetIdentifier  *string `json:"fleet_identifier,omitempty"`
			LastName         string  `json:"last_name"`
			MarketingOptIn   *bool   `json:"marketing_opt_in,omitempty"`
			Notes            *string `json:"notes,omitempty"`
			PhonesAttributes *[]struct {
				Label   string `json:"label"`
				Number  string `json:"number"`
				Primary bool   `json:"primary"`
			} `json:"phones_attributes,omitempty"`
			PoRequired      *bool          `json:"po_required,omitempty"`
			TagIds          *[]interface{} `json:"tag_ids,omitempty"`
			TaxExempt       *bool          `json:"tax_exempt,omitempty"`
			TaxExemptNumber *string        `json:"tax_exempt_number,omitempty"`
		}{
			FirstName:        r.FirstName,
			LastName:         r.LastName,
			CompanyName:      r.CompanyName,
			FleetIdentifier:  r.FleetIdentifier,
			BillingTerms:     r.BillingTerms,
			CreditLimitCents:  r.CreditLimitCents,
			TaxExempt:        r.TaxExempt,
			TaxExemptNumber:  r.TaxExemptNumber,
			Notes:           r.Notes,
			MarketingOptIn:  r.MarketingOptIn,
			DiscountPercent: r.DiscountPercent,
			PoRequired:      r.PoRequired,
			TagIds:          r.TagIds,
		},
	}
	if r.Emails != nil {
		emails := make([]struct {
			Email   string `json:"email"`
			Label   string `json:"label"`
			Primary bool   `json:"primary"`
		}, len(*r.Emails))
		for i, e := range *r.Emails {
			emails[i] = struct {
				Email   string `json:"email"`
				Label   string `json:"label"`
				Primary bool   `json:"primary"`
			}{Email: e.Email, Label: e.Label, Primary: e.Primary}
		}
		body.Customer.EmailsAttributes = &emails
	}
	if r.Phones != nil {
		phones := make([]struct {
			Label   string `json:"label"`
			Number  string `json:"number"`
			Primary bool   `json:"primary"`
		}, len(*r.Phones))
		for i, p := range *r.Phones {
			phones[i] = struct {
				Label   string `json:"label"`
				Number  string `json:"number"`
				Primary bool   `json:"primary"`
			}{Label: p.Label, Number: p.Number, Primary: p.Primary}
		}
		body.Customer.PhonesAttributes = &phones
	}
	if r.Addresses != nil {
		addresses := make([]struct {
			Address1   string `json:"address1"`
			City       string `json:"city"`
			Country    string `json:"country"`
			IsBilling  bool   `json:"is_billing"`
			PostalCode string `json:"postal_code"`
			State      string `json:"state"`
		}, len(*r.Addresses))
		for i, a := range *r.Addresses {
			addresses[i] = struct {
				Address1   string `json:"address1"`
				City       string `json:"city"`
				Country    string `json:"country"`
				IsBilling  bool   `json:"is_billing"`
				PostalCode string `json:"postal_code"`
				State      string `json:"state"`
			}{Address1: a.Address1, City: a.City, Country: a.Country, IsBilling: a.IsBilling, PostalCode: a.PostalCode, State: a.State}
		}
		body.Customer.AddressesAttributes = &addresses
	}
	return body
}

// UpdateCustomerRequest is the typed input for updating a customer.
type UpdateCustomerRequest struct {
	FirstName        *string
	LastName         *string
	CompanyName      *string
	FleetIdentifier  *string
	BillingTerms     *string
	CreditLimitCents *string
	TaxExempt        *bool
	Notes            *string
	MarketingOptIn   *bool
	DiscountPercent  *string
	PoRequired      *bool
	Emails           *[]EmailUpdateAttribute
	Phones           *[]PhoneUpdateAttribute
}

// ToGenerated converts to the generated request body type.
func (r UpdateCustomerRequest) ToGenerated() generated.UpdateCustomerJSONRequestBody {
	body := generated.UpdateCustomerJSONRequestBody{
		Customer: struct {
			AddressesAttributes *[]struct {
				Address1   string `json:"address1"`
				City       string `json:"city"`
				Country    string `json:"country"`
				IsBilling  bool   `json:"is_billing"`
				PostalCode string `json:"postal_code"`
				State      string `json:"state"`
			} `json:"addresses_attributes,omitempty"`
			BillingTerms     *string `json:"billing_terms,omitempty"`
			CompanyName      *string `json:"company_name,omitempty"`
			CreditLimitCents *string `json:"credit_limit_cents,omitempty"`
			DiscountPercent  *string `json:"discount_percent,omitempty"`
			EmailsAttributes *[]struct {
				Email string  `json:"email"`
				Id    *int    `json:"id,omitempty"`
				Label *string `json:"label,omitempty"`
			} `json:"emails_attributes,omitempty"`
			FirstName        *string `json:"first_name,omitempty"`
			FleetIdentifier  *string `json:"fleet_identifier,omitempty"`
			LastName         *string `json:"last_name,omitempty"`
			MarketingOptIn   *bool   `json:"marketing_opt_in,omitempty"`
			Notes            *string `json:"notes,omitempty"`
			PhonesAttributes *[]struct {
				UnderscoreDestroy *bool   `json:"_destroy,omitempty"`
				Id                *int    `json:"id,omitempty"`
				Label             *string `json:"label,omitempty"`
				Number            *string `json:"number,omitempty"`
				Primary           *bool   `json:"primary,omitempty"`
			} `json:"phones_attributes,omitempty"`
			PoRequired *bool `json:"po_required,omitempty"`
			TaxExempt  *bool `json:"tax_exempt,omitempty"`
		}{
			FirstName:        r.FirstName,
			LastName:         r.LastName,
			CompanyName:      r.CompanyName,
			FleetIdentifier:  r.FleetIdentifier,
			BillingTerms:     r.BillingTerms,
			CreditLimitCents:  r.CreditLimitCents,
			TaxExempt:        r.TaxExempt,
			Notes:           r.Notes,
			MarketingOptIn:  r.MarketingOptIn,
			DiscountPercent: r.DiscountPercent,
			PoRequired:      r.PoRequired,
		},
	}
	if r.Emails != nil {
		emails := make([]struct {
			Email string  `json:"email"`
			Id    *int    `json:"id,omitempty"`
			Label *string `json:"label,omitempty"`
		}, len(*r.Emails))
		for i, e := range *r.Emails {
			emails[i] = struct {
				Email string  `json:"email"`
				Id    *int    `json:"id,omitempty"`
				Label *string `json:"label,omitempty"`
			}{Email: e.Email, Id: e.Id, Label: e.Label}
		}
		body.Customer.EmailsAttributes = &emails
	}
	if r.Phones != nil {
		phones := make([]struct {
			UnderscoreDestroy *bool   `json:"_destroy,omitempty"`
			Id                *int    `json:"id,omitempty"`
			Label             *string `json:"label,omitempty"`
			Number            *string `json:"number,omitempty"`
			Primary           *bool   `json:"primary,omitempty"`
		}, len(*r.Phones))
		for i, p := range *r.Phones {
			phones[i] = struct {
				UnderscoreDestroy *bool   `json:"_destroy,omitempty"`
				Id                *int    `json:"id,omitempty"`
				Label             *string `json:"label,omitempty"`
				Number            *string `json:"number,omitempty"`
				Primary           *bool   `json:"primary,omitempty"`
			}{UnderscoreDestroy: p.UnderscoreDestroy, Id: p.Id, Label: p.Label, Number: p.Number, Primary: p.Primary}
		}
		body.Customer.PhonesAttributes = &phones
	}
	return body
}

// MergeCustomerRequest is the typed input for merging a customer.
type MergeCustomerRequest struct {
	SourceCustomerID int
}

// ToGenerated converts to the generated request body type.
func (r MergeCustomerRequest) ToGenerated() generated.MergeCustomerJSONRequestBody {
	return generated.MergeCustomerJSONRequestBody{SourceCustomerId: r.SourceCustomerID}
}

// --- Vehicle requests ---

// CreateVehicleRequest is the typed input for creating a vehicle.
type CreateVehicleRequest struct {
	CustomerID        int
	Make               string
	Model              string
	Year               int
	Vin                *string
	Submodel           *string
	BodyStyle          *string
	Engine             *string
	Transmission       *string
	Drivetrain         *string
	Color              *string
	LicensePlate       *string
	LicensePlateState  *string
	OdometerReading    *int
	OdometerUnit       *string
	UnitNumber         *string
	FleetIdentifier    *string
	ProductionDate     *string
	Notes              *string
	VehicleTagIDs      *[]interface{}
}

// ToGenerated converts to the generated request body type.
func (r CreateVehicleRequest) ToGenerated() generated.CreateVehicleJSONRequestBody {
	return generated.CreateVehicleJSONRequestBody{
		Vehicle: struct {
			BodyStyle         *string        `json:"body_style,omitempty"`
			Color             *string        `json:"color,omitempty"`
			CustomerId        int            `json:"customer_id"`
			Drivetrain        *string        `json:"drivetrain,omitempty"`
			Engine            *string        `json:"engine,omitempty"`
			FleetIdentifier   *string        `json:"fleet_identifier,omitempty"`
			LicensePlate      *string        `json:"license_plate,omitempty"`
			LicensePlateState *string        `json:"license_plate_state,omitempty"`
			Make              string         `json:"make"`
			Model             string         `json:"model"`
			Notes             *string        `json:"notes,omitempty"`
			OdometerReading   *int           `json:"odometer_reading,omitempty"`
			OdometerUnit      *string        `json:"odometer_unit,omitempty"`
			ProductionDate    *string        `json:"production_date,omitempty"`
			Submodel          *string        `json:"submodel,omitempty"`
			Transmission      *string        `json:"transmission,omitempty"`
			UnitNumber        *string        `json:"unit_number,omitempty"`
			VehicleTagIds     *[]interface{} `json:"vehicle_tag_ids,omitempty"`
			Vin               *string        `json:"vin,omitempty"`
			Year              int            `json:"year"`
		}{
			CustomerId:        r.CustomerID,
			Make:              r.Make,
			Model:             r.Model,
			Year:              r.Year,
			Vin:               r.Vin,
			Submodel:          r.Submodel,
			BodyStyle:         r.BodyStyle,
			Engine:            r.Engine,
			Transmission:      r.Transmission,
			Drivetrain:        r.Drivetrain,
			Color:             r.Color,
			LicensePlate:      r.LicensePlate,
			LicensePlateState: r.LicensePlateState,
			OdometerReading:   r.OdometerReading,
			OdometerUnit:      r.OdometerUnit,
			UnitNumber:        r.UnitNumber,
			FleetIdentifier:   r.FleetIdentifier,
			ProductionDate:    r.ProductionDate,
			Notes:             r.Notes,
			VehicleTagIds:     r.VehicleTagIDs,
		},
	}
}

// UpdateVehicleRequest is the typed input for updating a vehicle.
type UpdateVehicleRequest struct {
	Make              string
	Vin               *string
	Submodel           *string
	BodyStyle          *string
	Engine             *string
	Transmission       *string
	Drivetrain         *string
	Color              *string
	LicensePlate       *string
	LicensePlateState  *string
	OdometerReading    *int
	OdometerUnit       *string
	Notes              *string
}

// ToGenerated converts to the generated request body type.
func (r UpdateVehicleRequest) ToGenerated() generated.UpdateVehicleJSONRequestBody {
	return generated.UpdateVehicleJSONRequestBody{
		Vehicle: struct {
			BodyStyle         *string `json:"body_style,omitempty"`
			Color             *string `json:"color,omitempty"`
			Drivetrain        *string `json:"drivetrain,omitempty"`
			Engine            *string `json:"engine,omitempty"`
			LicensePlate      *string `json:"license_plate,omitempty"`
			LicensePlateState *string `json:"license_plate_state,omitempty"`
			Make              string  `json:"make"`
			Model             *string `json:"model,omitempty"`
			Notes             *string `json:"notes,omitempty"`
			OdometerReading   *int    `json:"odometer_reading,omitempty"`
			OdometerUnit      *string `json:"odometer_unit,omitempty"`
			Submodel          *string `json:"submodel,omitempty"`
			Transmission      *string `json:"transmission,omitempty"`
			Vin               *string `json:"vin,omitempty"`
			Year              *int    `json:"year,omitempty"`
		}{
			Make:              r.Make,
			Vin:               r.Vin,
			Submodel:          r.Submodel,
			BodyStyle:         r.BodyStyle,
			Engine:            r.Engine,
			Transmission:      r.Transmission,
			Drivetrain:        r.Drivetrain,
			Color:             r.Color,
			LicensePlate:      r.LicensePlate,
			LicensePlateState: r.LicensePlateState,
			OdometerReading:   r.OdometerReading,
			OdometerUnit:      r.OdometerUnit,
			Notes:             r.Notes,
		},
	}
}

// TransferVehicleRequest is the typed input for transferring a vehicle.
type TransferVehicleRequest struct {
	CustomerID int
	Mode       string
}

// ToGenerated converts to the generated request body type.
func (r TransferVehicleRequest) ToGenerated() generated.TransferVehicleJSONRequestBody {
	return generated.TransferVehicleJSONRequestBody{
		CustomerId: r.CustomerID,
		Mode:       r.Mode,
	}
}

// MergeVehicleRequest is the typed input for merging a vehicle.
type MergeVehicleRequest struct {
	SourceVehicleID int
}

// ToGenerated converts to the generated request body type.
func (r MergeVehicleRequest) ToGenerated() generated.MergeVehicleJSONRequestBody {
	return generated.MergeVehicleJSONRequestBody{SourceVehicleId: r.SourceVehicleID}
}

// --- Tag requests ---

// CustomerTagUpdate is a tag entry for bulk tag updates.
type CustomerTagUpdate struct {
	UnderscoreDestroy *string
	Id                int
	Name              *string
}

// VehicleTagUpdate is a tag entry for bulk tag updates.
type VehicleTagUpdate struct {
	UnderscoreDestroy string
	Id                int
}

// UpdateTagsRequest is the typed input for bulk updating tags.
type UpdateTagsRequest struct {
	CustomerTags []CustomerTagUpdate
	VehicleTags  *[]VehicleTagUpdate
}

// ToGenerated converts to the generated request body type.
func (r UpdateTagsRequest) ToGenerated() generated.UpdateTagsJSONRequestBody {
	body := generated.UpdateTagsJSONRequestBody{}
	if r.CustomerTags != nil {
		body.CustomerTags = make([]struct {
			UnderscoreDestroy *string `json:"_destroy,omitempty"`
			Id                int     `json:"id"`
			Name              *string `json:"name,omitempty"`
		}, len(r.CustomerTags))
		for i, t := range r.CustomerTags {
			body.CustomerTags[i] = struct {
				UnderscoreDestroy *string `json:"_destroy,omitempty"`
				Id                int     `json:"id"`
				Name              *string `json:"name,omitempty"`
			}{UnderscoreDestroy: t.UnderscoreDestroy, Id: t.Id, Name: t.Name}
		}
	}
	if r.VehicleTags != nil {
		vt := make([]struct {
			UnderscoreDestroy string `json:"_destroy"`
			Id                int    `json:"id"`
		}, len(*r.VehicleTags))
		for i, t := range *r.VehicleTags {
			vt[i] = struct {
				UnderscoreDestroy string `json:"_destroy"`
				Id                int    `json:"id"`
			}{UnderscoreDestroy: t.UnderscoreDestroy, Id: t.Id}
		}
		body.VehicleTags = &vt
	}
	return body
}

// CreateCustomerTagRequest is the typed input for creating a customer tag.
type CreateCustomerTagRequest struct {
	Name string
}

// ToGenerated converts to the generated request body type.
func (r CreateCustomerTagRequest) ToGenerated() generated.CreateCustomerTagJSONRequestBody {
	return generated.CreateCustomerTagJSONRequestBody{Name: r.Name}
}

// CreateVehicleTagRequest is the typed input for creating a vehicle tag.
type CreateVehicleTagRequest struct {
	Name string
}

// ToGenerated converts to the generated request body type.
func (r CreateVehicleTagRequest) ToGenerated() generated.CreateVehicleTagJSONRequestBody {
	return generated.CreateVehicleTagJSONRequestBody{Name: r.Name}
}