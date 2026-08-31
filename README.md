# Wenmar SDK

Public SDKs and OpenAPI spec for the Wenmar Pro automotive shop management API.

- [Go SDK](go/README.md) — `github.com/wenmar-pro/wenmar-sdk/go`
- [Ruby SDK](ruby/README.md) — the `wenmar` gem
- [API spec & docs](docs/api/README.md) — the canonical OpenAPI spec and human-readable docs

## Quick start

**Go**

```go
import "github.com/wenmar-pro/wenmar-sdk/go/wenmar"

cfg := wenmar.DefaultConfig()
cfg.BaseURL = "https://app.wenmarpro.com"
client, err := wenmar.NewClient(cfg, wenmar.NewStaticTokenProvider("YOUR_API_TOKEN"))
resp, err := client.ListCustomers(ctx, nil)
```

**Ruby**

```ruby
require "wenmar"

client = Wenmar::Client.new(token: "YOUR_API_TOKEN")
customers = client.list_customers
```

## API coverage

The full endpoint table is [generated from the OpenAPI spec](docs/api/api-reference.md) — see it for the live list. Resources are grouped per-tag in `docs/api/sections/`.

## Spec

The canonical OpenAPI spec lives in [`spec/openapi.yaml`](spec/openapi.yaml).
It is pushed by wenmar-pro's CI on every merge to `main` and is read-only here.

## Repository structure

- `spec/` — OpenAPI 3.0 spec + shared JSON fixtures
- `docs/api/` — API docs: hand-written narrative + generated reference (CC BY-SA)
- `go/` — Go SDK (generated client + hand-written layer)
- `ruby/` — Ruby gem (hand-written Faraday client)
- `conformance/` — shared behavioral tests run against both SDKs
- `scripts/` — spec enrichment, doc generation, fixture-coverage checker

## License

Code is MIT. API docs in `docs/api/` are [CC BY-SA 4.0](docs/api/LICENSE).
