# Wenmar — Go SDK

Go client for the Wenmar Pro API.

```
go get github.com/wenmar-pro/wenmar-sdk/go
```

## Quick start

```go
import "github.com/wenmar-pro/wenmar-sdk/go/wenmar"

client, err := wenmar.NewClient("https://app.wenmarpro.com", "YOUR_API_TOKEN")
if err != nil {
    // token is required
}

// List customers (paginated)
resp, err := client.ListCustomers(ctx, nil)

// Show a customer
customer, err := client.ShowCustomer(ctx, 1)

// Create a customer
created, err := client.CreateCustomer(ctx, generated.CreateCustomerJSONRequestBody{
    Customer: &struct {
        Email    *string `json:"email,omitempty"`
        FullName string  `json:"full_name"`
        Phone    *string `json:"phone,omitempty"`
    }{FullName: "Jane Doe"},
})
```

The client uses the `context` package for cancellation and deadlines.

## Configuration

`wenmar.NewClient(baseURL, token)` returns a `*Client` with retry, pagination,
error mapping, and bearer-token auth built in. An empty token is an error.

## API coverage

| Operation | Method |
|---|---|
| List customers | `ListCustomers(ctx, page *int)` |
| Create customer | `CreateCustomer(ctx, body)` |
| Show customer | `ShowCustomer(ctx, id)` |
| Show vehicle | `ShowVehicle(ctx, id)` |
| List work orders | `ListWorkOrders(ctx, page *int)` |
| Show work order | `ShowWorkOrder(ctx, id)` |

The `*generated.*Response` values carry `.Body` (raw JSON), `.HTTPResponse`,
and parsed `.JSON200` / `.JSON404` fields.

## Pagination

List endpoints paginate via the RFC 5988 `Link` header. Use the paginated
helpers to walk pages:

```go
resp, paginator, err := client.ListCustomersWithPagination(ctx, nil)
for paginator.HasNext() {
    resp, err = paginator.NextPage(ctx)
    // resp holds the next page
}
```

## Errors

Non-2xx responses return a `*wenmar.APIError`:

```go
resp, err := client.ShowCustomer(ctx, 999)
if err != nil {
    apiErr := err.(*wenmar.APIError)
    apiErr.Code       // => "not_found"
    apiErr.StatusCode // => 404
    apiErr.Message
    apiErr.Details
}
```

See [docs/errors.md](../docs/errors.md) for the full error envelope and code table.

## Retry

The client retries 5xx responses with exponential backoff + jitter (max 3
retries). It respects the `Retry-After` response header and never retries 4xx
errors.

## Documentation

- [API reference](../docs/api-reference.md)
- [Authentication](../docs/authentication.md)
- [Pagination](../docs/pagination.md)
- [Errors](../docs/errors.md)

## License

MIT
