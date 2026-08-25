# Wenmar SDK

Public SDKs and OpenAPI spec for the Wenmar Pro automotive shop management API.

- [Go SDK](go/README.md) — `github.com/wenmar-pro/wenmar-sdk/go`
- [Ruby SDK](ruby/README.md) — the `wenmar` gem
- [API documentation](https://wenmar-pro.github.io/wenmar-sdk/) — full reference, auto-generated from the spec
- [API reference](docs/api-reference.md) — all endpoints, request/response examples
- [Authentication](docs/authentication.md)
- [Pagination](docs/pagination.md)
- [Errors](docs/errors.md)

## Quick start

**Go**

```go
import "github.com/wenmar-pro/wenmar-sdk/go/wenmar"

client, err := wenmar.NewClient("https://app.wenmarpro.com", "YOUR_API_TOKEN")
resp, err := client.ListCustomers(ctx, nil)
```

**Ruby**

```ruby
require "wenmar"

client = Wenmar::Client.new(token: "YOUR_API_TOKEN")
customers = client.list_customers
```

## API coverage

| Resource | Operations |
|---|---|
| Customers | List, Create, Show |
| Vehicles | Show |
| Work orders | List, Show |

## Spec

The canonical OpenAPI spec lives in [`spec/openapi.yaml`](spec/openapi.yaml).
It is synced from the private wenmar-pro repo and is read-only here.

## Repository structure

- `spec/` — OpenAPI 3.0 spec + shared JSON fixtures
- `go/` — Go SDK (generated client + hand-written layer)
- `ruby/` — Ruby gem (hand-written Faraday client)
- `conformance/` — shared behavioral tests run against both SDKs
- `scripts/` — spec enrichment and fixture-coverage checker

## License

MIT
