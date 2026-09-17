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
`CacheEnabled`, `Token`/`TokenProvider`, and `Hooks`.

When the `TokenProvider` argument is nil, `NewClient` falls back to
`cfg.TokenProvider`, then to `cfg.Token` via a static provider, and only errors
if all are empty:

```go
cfg := wenmar.DefaultConfig()
cfg.Token = "YOUR_API_TOKEN" // used when no provider is passed
client, err := wenmar.NewClient(cfg, nil)
```

### Environment variables

| Variable | Read by | Purpose |
|---|---|---|
| `WENMAR_URL` | CLI + `auth.LoadConfigFromEnv` | API base URL (also accepts `WENMAR_BASE_URL` as fallback) |
| `WENMAR_BASE_URL` | `wenmar.LoadConfigFromEnv`, Ruby `Config.from_env` | API base URL |
| `WENMAR_TOKEN` | CLI + `auth.LoadConfigFromEnv` | static bearer token |
| `WENMAR_LOCATION_ID` | CLI + `auth.LoadConfigFromEnv`, Ruby `Config.from_env` | default location scope |
| `WENMAR_TIMEOUT` / `WENMAR_MAX_RETRIES` / `WENMAR_CACHE` | `wenmar.LoadConfigFromEnv`, Ruby `Config.from_env` | client tuning |

## Location scoping

Use `ForLocation` to scope every request to a specific location. The parent
client is not mutated:

```go
shop, err := client.ForLocation("42")
if err != nil {
    return err
}
resp, err := shop.ListCustomers(ctx, nil) // sends X-Wenmar-Location: 42
```

## API coverage

The full surface is generated into `operations.gen.go` — hundreds of methods
across every tag. See the [generated API reference](../docs/api/api-reference.md)
for the live list. A representative sample:

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
| Trash vehicle | `TrashVehicle(ctx, id)` |
| Archive vehicle | `ArchiveVehicle(ctx, id)` |
| Restore vehicle | `RestoreVehicle(ctx, id)` |
| Decode VIN | `DecodeVin(ctx, params *DecodeVinParams)` |
| Check duplicates | `CheckVehicleDuplicate(ctx, params *CheckVehicleDuplicateParams)` |
| List work orders | `ListWorkOrders(ctx)` |
| Create work order | `CreateWorkOrder(ctx, body CreateWorkOrderRequest)` |
| Show work order | `ShowWorkOrder(ctx, id)` |
| Update work order | `UpdateWorkOrder(ctx, id, body UpdateWorkOrderRequest)` |

Work orders use a domain workflow (`stage`) and are never hard-deleted — use the
lifecycle/workflow actions (e.g. `VoidWorkOrder`, `ReopenWorkOrder`) instead.
Lifecycle-managed resources (customers, vehicles, vendors, …) support
`Trash*`/`Archive*`/`Restore*` rather than `DELETE`.

Every paginated list also has a `GetAll*` variant that auto-paginates with a
1,000-item safety cap, e.g. `GetAllCustomers(ctx, nil, nil)`. `GetAll*` returns
`(items, truncated, err)` so callers can detect when the cap stopped collection.

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
items, truncated, err := client.GetAllCustomers(ctx, nil, nil) // nil opts = default 1000 cap
items, truncated, err := client.GetAllCustomers(ctx, nil, &wenmar.GetAllOptions{MaxItems: 50})
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
