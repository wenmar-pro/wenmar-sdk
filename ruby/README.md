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

All 76 operations are generated into `resources.rb`. Key methods:

| Operation | Method |
|---|---|
| List customers | `list_customers(query: nil, page: nil, ...)` |
| Create customer | `create_customer(customer:)` |
| Show customer | `show_customer(id)` |
| Update customer | `update_customer(id, customer:)` |
| List vehicles | `list_vehicles(customer_id: nil, page: nil)` |
| Create vehicle | `create_vehicle(vehicle:)` |
| Show vehicle | `show_vehicle(id)` |
| Update vehicle | `update_vehicle(id, vehicle:)` |
| Delete vehicle | `delete_vehicle(id)` |
| Decode VIN | `decode_vin(vin:)` |
| Check duplicates | `check_vehicle_duplicate(vin:)` |
| List work orders | `list_work_orders` |
| Create work order | `create_work_order(work_order:)` |
| Show work order | `show_work_order(id)` |
| Update work order | `update_work_order(id, work_order:)` |
| Delete work order | `delete_work_order(id)` |

Every paginated list also has a `get_all_*` variant that auto-paginates with a
1,000-item safety cap, e.g. `get_all_customers`.

## Pagination

List endpoints paginate via the RFC 5988 `Link` header. Paginated list methods
return a `Wenmar::Paginator`:

```ruby
result = client.list_customers
result.each { |customer| puts customer["full_name"] }
```

Or collect everything with `get_all_customers`.

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
respects the `Retry-After` response header. Mutations are only retried on 429.

## Documentation

- [API reference](../docs/api/api-reference.md)
- [Authentication](../docs/api/authentication.md)
- [Pagination](../docs/api/pagination.md)
- [Errors](../docs/api/errors.md)

## License

MIT
