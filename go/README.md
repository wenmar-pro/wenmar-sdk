# Wenmar — Go SDK

Go client for the Wenmar Pro API.

```
go get github.com/wenmar-pro/wenmar-sdk/go
```

## Quick start

```go
import "github.com/wenmar-pro/wenmar-sdk/go/wenmar"

cfg := wenmar.DefaultConfig()
cfg.BaseURL = "https://app.wenmarpro.com"

client, err := wenmar.NewClient(cfg, wenmar.NewStaticTokenProvider("YOUR_API_TOKEN"))
if err != nil {
    // token provider is required
}

// List customers (paginated via the Link header)
resp, err := client.ListCustomers(ctx, nil)

// Show a customer
customer, err := client.ShowCustomer(ctx, 1)

// Create a customer (request body is nested under the resource key)
created, err := client.CreateCustomer(ctx, wenmar.CreateCustomerRequest{
    Customer: struct {
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
    }{
        FirstName: "Jane",
        LastName:  "Doe",
    },
})
```

The client uses the `context` package for cancellation and deadlines.

## Configuration

`wenmar.NewClient(cfg, tp)` takes a `Config` and a `TokenProvider`. The
`Config` supports a custom `HTTPClient`, `Timeout`, `MaxRetries`,
`CacheEnabled`, and `Hooks`. An empty token provider is an error.

## Location scoping

Use `ForLocation` to scope every request to a specific location. The parent
client is not mutated:

```go
shop := client.ForLocation("42")
resp, err := shop.ListCustomers(ctx, nil) // sends X-Wenmar-Location: 42
```

## API coverage

All 76 operations are generated into `operations.gen.go`. Key methods:

| Operation | Method |
|---|---|
| List customers | `ListCustomers(ctx, params *ListCustomersParams)` |
| Create customer | `CreateCustomer(ctx, body CreateCustomerRequest)` |
| Show customer | `ShowCustomer(ctx, id)` |
| Update customer | `UpdateCustomer(ctx, id, body UpdateCustomerRequest)` |
| List vehicles | `ListVehicles(ctx, params *ListVehiclesParams)` |
| Create vehicle | `CreateVehicle(ctx, body CreateVehicleRequest)` |
| Show vehicle | `ShowVehicle(ctx, id)` |
| Update vehicle | `UpdateVehicle(ctx, id, body UpdateVehicleRequest)` |
| Delete vehicle | `DeleteVehicle(ctx, id)` |
| Decode VIN | `DecodeVin(ctx, params *DecodeVinParams)` |
| Check duplicates | `CheckVehicleDuplicate(ctx, params *CheckVehicleDuplicateParams)` |
| List work orders | `ListWorkOrders(ctx)` |
| Create work order | `CreateWorkOrder(ctx, body CreateWorkOrderRequest)` |
| Show work order | `ShowWorkOrder(ctx, id)` |
| Update work order | `UpdateWorkOrder(ctx, id, body UpdateWorkOrderRequest)` |
| Delete work order | `DeleteWorkOrder(ctx, id)` |

Every paginated list also has a `GetAll*` variant that auto-paginates with a
1,000-item safety cap, e.g. `GetAllCustomers(ctx, nil, nil)`.

Every `List*` method also has a `ListXxxRaw` variant that returns the raw
oapi-codegen response envelope for callers who need headers or status codes.

## Pagination

List endpoints paginate via the RFC 5988 `Link` header. All `List*` methods
return a typed `*ListResult[T]`:

```go
result, err := client.ListCustomers(ctx, nil)
// result.Items is []Customer
// result.Meta.TotalCount, result.Meta.PerPage, result.Meta.HasMore
for result.HasNext() {
    result, err = result.Next(ctx)
    // result.Items holds the next page
}
```

Or collect everything with `GetAllCustomers`, configurable via `GetAllOptions`:

```go
items, err := client.GetAllCustomers(ctx, nil, nil) // nil opts = default 1000 cap
items, err := client.GetAllCustomers(ctx, nil, &wenmar.GetAllOptions{MaxItems: 50})
```

For raw access to the full response envelope (headers, status code), use
the `ListXxxRaw` variant:

```go
resp, err := client.ListCustomersRaw(ctx, nil)
// resp.JSON200, resp.HTTPResponse, resp.StatusCode()
```

## Errors

Non-2xx responses return a `*wenmar.APIError`:

```go
resp, err := client.ShowCustomer(ctx, 999)
if err != nil {
    apiErr := err.(*wenmar.APIError)
    apiErr.Code         // => "not_found"
    apiErr.StatusCode   // => 404
    apiErr.Message
    apiErr.FieldErrorsMap
}
```

See [docs/errors.md](../docs/api/errors.md) for the full error envelope and code table.

## Retry

The client retries 429/503/504 with exponential backoff + jitter (max 3
retries). It respects the `Retry-After` response header. Mutations are only
retried on 429 (the throttle response means the request was not processed).

## Documentation

- [API reference](../docs/api/api-reference.md)
- [Authentication](../docs/api/authentication.md)
- [Pagination](../docs/api/pagination.md)
- [Errors](../docs/api/errors.md)

## License

MIT
