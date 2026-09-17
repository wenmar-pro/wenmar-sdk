# Migrating from v0.1

This document describes the breaking changes introduced in the v0.2 SDK
restructure. The SDK is pre-release, so no version bump was made — but the
public API changed and existing callers must update.

## Migrating from v0.8.1 to v0.9.0

v0.9.0 is a production-readiness release for the Go SDK. It removes unused
public surface, tightens the `GetAll*` default, and centralizes the token
provider. The Ruby gem API is unchanged. Breaking changes:

1. `wenmar.DefaultGetAllOptions` is removed. It was a mutable package-level
   `var` that `GetAll*` reused when you passed `nil` opts. It is replaced by an
   unexported constant (1,000-item safety cap). The `GetAll*` signature and
   behavior are unchanged — `nil` opts still caps at 1,000 and returns the
   `truncated` flag — but callers who mutated the old var can no longer do so.

   ```go
   // Before: wenmar.DefaultGetAllOptions.MaxItems = 5000
   // After:  pass explicit opts per call
   items, truncated, err := client.GetAllCustomers(ctx, nil, &wenmar.GetAllOptions{MaxItems: 5000})
   // Want unlimited? Pass an explicit &GetAllOptions{} (MaxItems stays 0 = unlimited).
   items, _, err := client.GetAllCustomers(ctx, nil, &wenmar.GetAllOptions{})
   ```

2. `wenmar.ParseErrorBody` now takes full request context:
   `ParseErrorBody(body, statusCode, method, path, requestID)`.
   `ParseErrorBodyWithRequest` and `ParseErrorBodyWithRequestAndID` are
   removed.

   ```go
   // Before: wenmar.ParseErrorBody(body, statusCode)
   // After:
   apiErr := wenmar.ParseErrorBody(body, statusCode, method, path, requestID)
   ```

3. `wenmar.LoadConfigFromEnv` is removed. The Go SDK no longer reads
   `WENMAR_*` environment variables; build the `Config` in code via
   `wenmar.DefaultConfig()` and override the fields you need. (Ruby still
   reads env vars through `Config.from_env`.)

   ```go
   // Before: cfg := wenmar.LoadConfigFromEnv()
   // After:
   cfg := wenmar.DefaultConfig()
   cfg.BaseURL = "https://app.wenmarpro.com"
   ```

4. `pkg/auth.Config`, `pkg/auth.DefaultConfig`, and
   `pkg/auth.LoadConfigFromEnv` are removed (unused). Consumers read env vars
   themselves.

5. The canonical `TokenProvider` interface now lives in `pkg/token`
   (`token.Provider`, `token.NewStatic`). `wenmar.TokenProvider` and
   `wenmar.NewStaticTokenProvider` remain and work unchanged via type aliases;
   `auth.TokenProvider` is aliased to the same interface. No caller changes
   required, but new code may import `github.com/wenmar-pro/wenmar-sdk/go/pkg/token`
   directly.

Behavioral fixes shipped with this release (no API change): token refresh is
single-flighted; `Retry-After` is capped at 2m and the RoundTripper error
contract is honored; the conditional-GET cache is capped at 128 entries (FIFO
eviction); caller-supplied `http.Client` is copied before wrapping; and partial
response bodies are never cached or replayed. Both SDKs report `0.9.0`.

## Migrating from v0.8 to v0.8.1

v0.8.1 is a behavior-fix release with two breaking Go signature changes
(pre-1.0; the Ruby gem API is unchanged):

1. `GetAll*` now returns the truncation flag so callers can detect when the
   `MaxItems`/`MaxPages` cap stopped collection:

   ```go
   // Before: items, err := client.GetAllCustomers(ctx, nil, nil)
   items, truncated, err := client.GetAllCustomers(ctx, nil, nil)
   ```

2. `Client.ForLocation` returns `(*Client, error)` instead of swallowing the
   client-construction error:

   ```go
   // Before: scoped := client.ForLocation("42")
   scoped, err := client.ForLocation("42")
   ```

Behavioral fixes: typed `List*` methods now return an error when the response
body cannot be parsed (previously they returned an empty result silently);
OAuth refresh wraps transport failures in `TokenError` (Ruby); concurrent
token refresh is single-flighted (both SDKs); credential writes are atomic
(both SDKs). Both SDKs report `0.8.1`.

## Migrating from v0.7 to v0.8

v0.8 renames the domain-workflow `status` field on six resources to their
reserved domain names (Rule 17). `status` is now reserved for lifecycle state
(active/archived/trashed) and processing state (queued/sent) — no resource's
business-process field uses it any more. The HTTP paths are unchanged; only the
JSON field names, filter params, and generated SDK field/param names change.

| Before (v0.7) | After (v0.8) |
|---|---|
| `WorkOrder.status` | `WorkOrder.stage` |
| `Statement.status` | `Statement.billing_status` |
| `PurchaseOrder.status` | `PurchaseOrder.receiving_status` |
| `CounterSale.status` | `CounterSale.sale_status` |
| `Appointment.status` | `Appointment.scheduling_status` |
| `ReturnOrder.status` | `ReturnOrder.return_status` |

Filter params follow the field names:

| Before (v0.7) | After (v0.8) |
|---|---|
| `filter[status]` on `/reports/statements` | `filter[billing_status]` |
| `?status=` on `/appointments` | `?scheduling_status=` |

This is a breaking rename of generated struct fields, request/response bodies,
and query params. Both SDKs report `0.8.0`.

## Migrating from v0.6 to v0.7

v0.7 fixes a noun-singularization bug in the operationId generator. Several
generated method names were mis-singularized (the generator naively stripped a
trailing `s`), producing names like `create_cash_entrie` and
`update_labor_matrice`. These are now correct. This is a breaking rename of the
generated SDK methods (and the corresponding conformance dispatch entries); the
HTTP paths are unchanged.

| Before (v0.6) | After (v0.7) |
|---|---|
| `CreateCashEntrie` / `create_cash_entrie` | `CreateCashEntry` / `create_cash_entry` |
| `ShowCashEntrie` / `show_cash_entrie` | `ShowCashEntry` / `show_cash_entry` |
| `DeleteCashEntrie` / `delete_cash_entrie` | `DeleteCashEntry` / `delete_cash_entry` |
| `CreateLaborMatrice` / `create_labor_matrice` | `CreateLaborMatrix` / `create_labor_matrix` |
| `UpdateLaborMatrice` / `update_labor_matrice` | `UpdateLaborMatrix` / `update_labor_matrix` |
| `DeleteLaborMatrice` / `delete_labor_matrice` | `DeleteLaborMatrix` / `delete_labor_matrix` |
| `CreatePartsMatrice` / `create_parts_matrice` | `CreatePartsMatrix` / `create_parts_matrix` |
| `UpdatePartsMatrice` / `update_parts_matrice` | `UpdatePartsMatrix` / `update_parts_matrix` |
| `DeletePartsMatrice` / `delete_parts_matrice` | `DeletePartsMatrix` / `delete_parts_matrix` |
| `CreateSubStatuse` / `create_sub_statuse` | `CreateSubStatus` / `create_sub_status` |
| `UpdateSubStatuse` / `update_sub_statuse` | `UpdateSubStatus` / `update_sub_status` |
| `CreateWorkOrdersServicesCopie` / `create_work_orders_services_copie` | `CreateWorkOrdersServicesCopy` / `create_work_orders_services_copy` |
| `CreateWorkOrdersServicesTimeEntrie` / `create_work_orders_services_time_entrie` | `CreateWorkOrdersServicesTimeEntry` / `create_work_orders_services_time_entry` |
| `CreateWorkOrdersServicesLineItemsCopie` / `create_work_orders_services_line_items_copie` | `CreateWorkOrdersServicesLineItemsCopy` / `create_work_orders_services_line_items_copy` |
| `CreateWorkOrdersServicesLineItemsPriceRefreshe` / `create_work_orders_services_line_items_price_refreshe` | `CreateWorkOrdersServicesLineItemsPriceRefresh` / `create_work_orders_services_line_items_price_refresh` |

Other v0.7 changes:

- The Ruby gem version is now sourced from `Wenmar::VERSION` (was hardcoded and
  had drifted); both SDKs report `0.7.0`.
- The Ruby SDK gains OAuth refresh + credential storage (`Wenmar::Token`,
  `Wenmar::OAuth`, `Wenmar::CredentialStore`, `Wenmar::AuthManager`,
  `Wenmar::CredentialStoreProvider`). The Go SDK gains
  `auth.NewAuthManagerWithOAuth` for Doorkeeper refresh.
- The documented 1,000-item auto-pagination cap is now actually enforced; the
  default `GetAll*` behavior for `nil` options truncates at 1,000 (a `truncated`
  flag is reported internally).

## Migrating from v0.3 to v0.4.1

v0.4.1 is a breaking release that moves both SDKs onto a generated operation
layer driven by `spec/operations.json`. Request bodies are now nested under a
resource key, list methods take params structs, and several method names
changed to match the API's operation IDs.

### Request struct nesting

Create/update request bodies are now wrapped under a resource key instead of
being flat.

**Go — before:**

```go
resp, err := client.CreateCustomer(ctx, wenmar.CreateCustomerRequest{
    FirstName: "Jane",
    LastName:  "Doe",
})
```

**Go — after:**

```go
resp, err := client.CreateCustomer(ctx, wenmar.CreateCustomerRequest{
    Customer: struct {
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
    }{
        FirstName: "Jane",
        LastName:  "Doe",
    },
})
```

The same applies to `CreateVehicleRequest` (`.Vehicle`), `CreateWorkOrderRequest`
(`.WorkOrder`), `CreateDriverRequest` (`.Driver`), and
`CreateServiceCategoryRequest` (`.ServiceCategory`).

**Ruby — before:**

```ruby
client.create_customer(full_name: "Jane Doe")
```

**Ruby — after:**

```ruby
client.create_customer(customer: { first_name: "Jane", last_name: "Doe" })
```

### Method renames

Method names now align with the API operation IDs:

| Old | New |
|---|---|
| `ListDrivers` / `list_drivers` | `ListCustomersDrivers` / `list_customers_drivers` |
| `ListStatements` / `list_statements` | `ListCustomersStatements` / `list_customers_statements` |
| `ListCustomerVehicles` / `list_customer_vehicles` | `ListCustomersVehicles` / `list_customers_vehicles` |
| `ListCustomerWorkOrders` / `list_customer_work_orders` | `ListCustomersWorkOrders` / `list_customers_work_orders` |
| `ListVehicleWorkOrders` / `list_vehicle_work_orders` | `ListVehiclesWorkOrders` / `list_vehicles_work_orders` |

### List methods take params structs

**Go:** `ListCustomers(ctx, nil)` and `ListVehicles(ctx, nil)` now take a
`*ListCustomersParams` / `*ListVehiclesParams` (pass `nil` for defaults).
`DecodeVin` and `LookupVehicle` take a params struct instead of a bare string.

**Ruby:** list methods accept keyword params, e.g.
`list_customers(query: "Acme", page: 2)`.

### `Token` field removal

The `Config.Token` field is still accepted for convenience, but the preferred
path is a `TokenProvider`. `NewClient` now requires a non-nil `TokenProvider`
argument:

```go
client, err := wenmar.NewClient(cfg, wenmar.NewStaticTokenProvider("YOUR_API_TOKEN"))
```

### Credential JSON shape

The credential file at `~/.config/wenmar/credentials.json` now stores the
token under `access_token` (previously `token`). The old key is still read for
backwards compatibility. The keyring service name changed from `wenmar-cli`
to `wenmar`; legacy entries are read and migrated on first access.

### Hook signature change

`OnRequestStart` and `OnRequestEnd` now return / receive a `context.Context`,
so request-level OTel child spans and hook chaining are supported. See
`SPEC.md` for the updated signatures.

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
