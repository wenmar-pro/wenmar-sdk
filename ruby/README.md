# wenmar — Ruby SDK

Ruby SDK for the Wenmar Pro API.

## Installation

Add to your `Gemfile`:

```ruby
gem "wenmar"
```

then run `bundle install`. Or install it directly:

```sh
gem install wenmar
```

## Quick start

```ruby
require "wenmar"

client = Wenmar::Client.new(token: "YOUR_API_TOKEN")

# List customers (paginated)
customers = client.list_customers
customers # => [{ "id" => 1, "full_name" => "Jane Doe", ... }]

# Show a customer
customer = client.show_customer(1)
customer # => { "id" => 1, "full_name" => "Jane Doe", ... }

# Create a customer (request body is nested under the resource key)
created = client.create_customer(customer: { first_name: "Jane", last_name: "Doe" })
```

## Configuration

`Wenmar::Client.new` takes a token (required) and an optional base URL:

```ruby
client = Wenmar::Client.new(token: "YOUR_API_KEY", base_url: "https://app.wenmarpro.com")
```

## Location scoping

Use `for_location` to scope every request to a specific location. The parent
client is not mutated:

```ruby
shop = client.for_location("42")
shop.list_customers # sends X-Wenmar-Location: 42
```

## API coverage

The full surface is generated into `resources.rb` — hundreds of methods across
every tag. See the [generated API reference](../docs/api/api-reference.md) for
the live list. A representative sample:

| Operation | Method |
|---|---|
| List customers | `list_customers(customer_tag_id:, has_balance:, has_vehicle:, last_visit_months:, page:, per_page:, q:, status:, type:)` |
| Create customer | `create_customer(customer:)` |
| Show customer | `show_customer(id)` |
| Update customer | `update_customer(id, customer:)` |
| List vehicles | `list_vehicles(page:, per_page:, q:, ...)` |
| Create vehicle | `create_vehicle(vehicle:)` |
| Show vehicle | `show_vehicle(id)` |
| Update vehicle | `update_vehicle(id, vehicle:)` |
| Trash vehicle | `trash_vehicle(id)` |
| Decode VIN | `decode_vin(vin:)` |
| Check duplicates | `check_vehicle_duplicate(vin:)` |
| List work orders | `list_work_orders(page:, per_page:, q:, ...)` |
| Create work order | `create_work_order(work_order:)` |
| Show work order | `show_work_order(id)` |
| Update work order | `update_work_order(id, work_order:)` |
| Void work order | `void_work_order(id, closure_reason:)` |
| Reopen work order | `reopen_work_order(id)` |

Work orders use a domain workflow (`stage`) and are never hard-deleted — use
`void_work_order`/`reopen_work_order`. Lifecycle-managed resources (customers,
vehicles, vendors, …) support `trash_*`/`archive_*`/`restore_*` rather than
`delete`.

Every paginated list also has a `get_all_*` variant that auto-paginates with a
1,000-item safety cap, e.g. `get_all_customers`.

## Pagination

List endpoints paginate via the RFC 5988 `Link` header. Paginated list methods
return the page's items as an Array with a `.paginator` attached:

```ruby
result = client.list_customers

# Iterate all pages (page 1 included):
result.paginator.each { |customer| puts customer["full_name"] }

# Or collect everything with the generated helper:
all = client.get_all_customers
```

## Errors

All non-2xx responses raise `Wenmar::Error`:

```ruby
begin
  client.show_customer(999)
rescue Wenmar::Error => e
  e.code   # => "not_found"
  e.status # => 404
  e.message
  e.field_errors
end
```

See [docs/errors.md](../docs/api/errors.md) for the full error envelope and code table.

## Retry

The client retries 429/503/504 with exponential backoff (max 3 retries). It
respects the `Retry-After` response header. Mutations are only retried on 429 —
never on transport errors, which could duplicate a write.

## OAuth and credential storage

For OAuth (the browser `authorization_code` + PKCE flow is a CLI concern; the
SDK ships the token model, store, and refresh machinery):

```ruby
require "wenmar"

# Exchange a refresh token for a new access token.
token = Wenmar::OAuth.refresh(
  base_url: "https://app.wenmarpro.com",
  refresh_token: refresh_token
)

# Persist the full token (access + refresh + expiry) with 0600 permissions.
store = Wenmar::CredentialStore.new
store.save_token(token)

# Auto-refresh when the token is expired or within 5 minutes of expiry.
manager = Wenmar::AuthManager.new(store: store, oauth: { base_url: "https://app.wenmarpro.com" })
provider = Wenmar::CredentialStoreProvider.new(store: store, manager: manager)
client = Wenmar::Client.new(token_provider: provider)
```

`Wenmar::KeychainStore` is an alternative macOS keychain-backed store (it
requires the optional `ruby-keychain` gem, which is **not** a runtime
dependency — install it yourself if you want keychain storage).

## Documentation

- [API reference](../docs/api/api-reference.md)
- [Authentication](../docs/api/authentication.md)
- [Pagination](../docs/api/pagination.md)
- [Errors](../docs/api/errors.md)

## License

MIT
