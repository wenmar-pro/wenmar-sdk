# Migrating from v0.1

This document describes the breaking changes introduced in the v0.2 SDK
restructure. The SDK is pre-release, so no version bump was made — but the
public API changed and existing callers must update.

## Migrating from v0.2 to v0.3

v0.3.0 is a breaking release that tracks the wenmar-pro API's pre-v1
foundation fixes. Two breaking changes affect existing callers.

### Error envelope: `details` → `field_errors`

The API renamed the error envelope's field-error key from `details` to
`field_errors`. Validation errors now arrive as:

```json
{ "error": { "code": "validation_failed", "message": "...", "field_errors": { "first_name": ["can't be blank"] } } }
```

**Go:** `APIError.Details` is renamed to `APIError.FieldErrorsMap`. The
`FieldErrors()` accessor is unchanged and still returns
`map[string][]string`.

**Ruby:** `Error#details` is renamed to `Error#field_errors`. The
`field_errors_by_field` accessor (formerly `field_errors`) returns the
coerced `{ field => [messages] }` hash.

### Work order show: embedded arrays → `_url` links

`GET /work_orders/{id}` no longer embeds `services[]` and `payments[]`
arrays. They are replaced by `services_url`, `payments_url`, `wip_url`,
`inspection_url`, and `parts_url` link fields. Fetch each sub-collection
via its dedicated endpoint:

- `GET /work_orders/{id}/estimate` → `services[]`
- `GET /work_orders/{id}/wip` → `services[]`
- `GET /work_orders/{id}/inspection` → `inspection_reports[]`
- `GET /work_orders/{id}/parts` → `services[]`
- `GET /work_orders/{id}/payments` → `payments[]`

### Drivers: ad-hoc → full resource

Drivers were previously returned as `{ id, full_name }`. They are now full
resources with `phone`, `email`, `customer` stub, `work_orders_count`,
`work_orders_url`, timestamps, `url`, and `app_url`. Full CRUD is available
under `/customers/{customer_id}/drivers`.

## New features (v0.3)

- Drivers full CRUD: `ListDrivers`, `ShowDriver`, `CreateDriver`,
  `UpdateDriver`, `DeleteDriver` (Go) / `list_drivers`, `show_driver`,
  `create_driver`, `update_driver`, `delete_driver` (Ruby).
- Statements: `ListStatements`, `ShowStatement` (Go) / `list_statements`,
  `show_statement` (Ruby).
- Vendors: `ListVendors`, `ShowVendor` (Go) / `list_vendors`,
  `show_vendor` (Ruby).
- Work order sub-collections: `ShowWorkOrderEstimate`, `ShowWorkOrderWip`,
  `ShowWorkOrderInspection`, `ShowWorkOrderParts`, `ShowWorkOrderPayments`,
  `CreateWorkOrderPayment` (Go) / `show_work_order_estimate`,
  `show_work_order_wip`, `show_work_order_inspection`,
  `show_work_order_parts`, `show_work_order_payments`,
  `create_work_order_payment` (Ruby).

## Go

### `NewClient` signature

**Before:**

```go
client, err := wenmar.NewClient("https://app.wenmarpro.com", "YOUR_API_TOKEN")
```

**After:**

```go
cfg := wenmar.DefaultConfig()
cfg.BaseURL = "https://app.wenmarpro.com"
client, err := wenmar.NewClient(cfg, wenmar.NewStaticTokenProvider("YOUR_API_TOKEN"))
```

`NewClient` now takes a `Config` struct and a `TokenProvider`. The `Config`
supports env loading (`LoadConfigFromEnv`), timeout, retry count, and a cache
toggle. The `TokenProvider` can be a static token, the system keyring, or a
credential file.

### Hand-written request structs

The generated request body types are hidden from the public API. Callers no
longer import the `generated` package for requests.

**Before:**

```go
body := generated.CreateCustomerJSONRequestBody{}
body.Customer.FirstName = "Jane"
body.Customer.LastName = "Doe"
resp, err := client.CreateCustomer(ctx, body)
```

**After:**

```go
resp, err := client.CreateCustomer(ctx, wenmar.CreateCustomerRequest{
    FirstName: "Jane",
    LastName:  "Doe",
})
```

### Typed pagination

The old `Paginator.NextPage() (any, error)` is replaced by a typed
`ListResult[T]`.

**Before:**

```go
resp, paginator, err := client.ListCustomersWithPagination(ctx)
next, err := paginator.NextPage(ctx)
```

**After:**

```go
result, err := client.ListCustomersTyped(ctx)
for result.HasNext() {
    result, err = result.Next(ctx)
}
```

`GetAllCustomers` auto-paginates with a `max_items` safety cap.

### Response types

**Known limitation:** the generated response types (e.g.
`*generated.CreateCustomerResponse`) remain in the public API. Full
response-type hiding (mapping `generated.Customer` → `wenmar.Customer` on
every method) is a follow-up. The request types — the leakiest abstraction —
are hidden now.

## Ruby

### `Client.new`

**Before:**

```ruby
client = Wenmar::Client.new(token: "YOUR_API_TOKEN")
```

**After:**

```ruby
client = Wenmar::Client.new(
  config: Wenmar::Config.from_env,
  token_provider: Wenmar::StaticTokenProvider.new("YOUR_API_TOKEN")
)
```

The `token:` keyword still works as a backwards-compatible shortcut.

## New features

- `Config` / `Config.from_env` — reads `WENMAR_*` environment variables.
- `CredentialStore` — persists the token in the system keyring with a file
  fallback at `~/.config/wenmar/credentials.json`.
- `for_location` / `ForLocation` — a sub-client for explicit location context
  (SDK-side guard; server-side per-request location selection is a follow-up).
- `Hooks` (Go) — observability interface for logging, tracing, and metrics.
- `X-Request-Id` — captured on errors for support correlation.
- Versioned `User-Agent` — `wenmar-sdk-{lang}/{version}`.
