package wenmar

import "github.com/wenmar-pro/wenmar-sdk/go/pkg/generated"

// Model type aliases. These expose the generated model types under the
// wenmar package so callers never need to import pkg/generated.
type (
	Customer  = generated.Customer
	Vehicle   = generated.Vehicle
	WorkOrder = generated.WorkOrder
	Driver    = generated.Driver
	Statement = generated.Statement
	Vendor    = generated.Vendor
	Error     = generated.Error
)

// Response type aliases. Each is identical to the generated response
// struct — callers can access .JSON200, .HTTPResponse, .Body, etc.
// without importing pkg/generated.
type (
	ListAccountResponse               = generated.ListAccountResponse
	ListCustomersResponse             = generated.ListCustomersResponse
	ShowCustomerResponse              = generated.ShowCustomerResponse
	CreateCustomerResponse            = generated.CreateCustomerResponse
	UpdateCustomerResponse            = generated.UpdateCustomerResponse
	MergeCustomerResponse             = generated.MergeCustomerResponse
	LookupCustomerResponse            = generated.LookupCustomerResponse
	CheckCustomerDuplicateResponse    = generated.CheckCustomerDuplicateResponse
	DeleteCustomerResponse            = generated.DeleteCustomerResponse
	ListCustomersVehiclesResponse     = generated.ListCustomersVehiclesResponse
	GetCustomersVehicleHistoryResponse = generated.GetCustomersVehicleHistoryResponse
	ListCustomersWorkOrdersResponse   = generated.ListCustomersWorkOrdersResponse
	ListCustomersStatementsResponse   = generated.ListCustomersStatementsResponse
	ListCustomersDriversResponse      = generated.ListCustomersDriversResponse
	ShowDriverResponse                = generated.ShowDriverResponse
	CreateDriverResponse              = generated.CreateDriverResponse
	UpdateDriverResponse               = generated.UpdateDriverResponse
	DeleteDriverResponse               = generated.DeleteDriverResponse
	ShowLocationResponse              = generated.ShowLocationResponse
	ListTagsResponse                  = generated.ListTagsResponse
	UpdateTagsResponse                = generated.UpdateTagsResponse
	CreateCustomerTagResponse         = generated.CreateCustomerTagResponse
	CreateVehicleTagResponse          = generated.CreateVehicleTagResponse
	ShowStatementResponse             = generated.ShowStatementResponse
	ListVehiclesResponse              = generated.ListVehiclesResponse
	ShowVehicleResponse               = generated.ShowVehicleResponse
	CreateVehicleResponse             = generated.CreateVehicleResponse
	UpdateVehicleResponse             = generated.UpdateVehicleResponse
	DeleteVehicleResponse             = generated.DeleteVehicleResponse
	MergeVehicleResponse              = generated.MergeVehicleResponse
	TransferVehicleResponse           = generated.TransferVehicleResponse
	DecodeVinResponse                 = generated.DecodeVinResponse
	LookupVehicleResponse             = generated.LookupVehicleResponse
	CheckVehicleDuplicateResponse     = generated.CheckVehicleDuplicateResponse
	PrefillVehicleResponse            = generated.PrefillVehicleResponse
	ListVehiclesWorkOrdersResponse    = generated.ListVehiclesWorkOrdersResponse
	ListVendorsResponse               = generated.ListVendorsResponse
	ShowVendorResponse                = generated.ShowVendorResponse
	ListWorkOrdersResponse            = generated.ListWorkOrdersResponse
	ShowWorkOrderResponse             = generated.ShowWorkOrderResponse
	CreateWorkOrderResponse           = generated.CreateWorkOrderResponse
	UpdateWorkOrderResponse           = generated.UpdateWorkOrderResponse
	DeleteWorkOrderResponse           = generated.DeleteWorkOrderResponse
	ShowWorkOrderEstimateResponse     = generated.ShowWorkOrderEstimateResponse
	ShowWorkOrderWipResponse          = generated.ShowWorkOrderWipResponse
	ShowWorkOrderInspectionResponse   = generated.ShowWorkOrderInspectionResponse
	ShowWorkOrderPartsResponse        = generated.ShowWorkOrderPartsResponse
	ShowWorkOrderPaymentsResponse     = generated.ShowWorkOrderPaymentsResponse
	CreateWorkOrderPaymentResponse    = generated.CreateWorkOrderPaymentResponse
	GetWorkOrdersSummaryResponse      = generated.GetWorkOrdersSummaryResponse
	ListTeamResponse                  = generated.ListTeamResponse
)

// Params type aliases for query-parameter types used in public method
// signatures.
type (
	CheckCustomerDuplicateParams = generated.CheckCustomerDuplicateParams
	CheckVehicleDuplicateParams  = generated.CheckVehicleDuplicateParams
	PrefillVehicleParams         = generated.PrefillVehicleParams
	DecodeVinParams              = generated.DecodeVinParams
	ListVehiclesParams           = generated.ListVehiclesParams
)