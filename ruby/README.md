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

# Create a customer
created = client.create_customer(full_name: "Jane Doe")
```

## Configuration

`Wenmar::Client.new` takes a token (required) and an optional base URL:

```ruby
client = Wenmar::Client.new(token: "YOUR_API_KEY", base_url: "https://app.wenmarpro.com")
```

## API coverage

| Operation | Method |
|---|---|
| List account | `list_account` |
| List customers | `list_customers(page: nil)` |
| Create customer | `create_customer(full_name: ..., email: ..., phone: ...)` |
| Show customer | `show_customer(id)` |
| Update customer | `update_customer(id, attrs)` |
| List vehicles | `list_vehicles` |
| Create vehicle | `create_vehicle(attrs)` |
| Show vehicle | `show_vehicle(id)` |
| Update vehicle | `update_vehicle(id, attrs)` |
| Delete vehicle | `delete_vehicle(id)` |
| Decode VIN | `decode_vin(vin)` |
| Check duplicates | `check_duplicate(vin)` |
| List work orders | `list_work_orders(page: nil)` |
| Create work order | `create_work_order(attrs)` |
| Show work order | `show_work_order(id)` |
| Update work order | `update_work_order(id, attrs)` |
| Delete work order | `delete_work_order(id)` |
| Show location | `show_location(id)` |

All methods return the parsed JSON response body as a `Hash`.

## Pagination

List endpoints paginate via the RFC 5988 `Link` header. Call `.paginator` on
the result to walk pages:

```ruby
result = client.list_customers
while result.paginator.has_next?
  result = result.paginator.next_page
  result # => next page of records
end
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

The client retries 5xx responses with exponential backoff (max 3 retries). It
respects the `Retry-After` response header and never retries 4xx errors.

## Documentation

- [API reference](../docs/api/api-reference.md)
- [Authentication](../docs/api/authentication.md)
- [Pagination](../docs/api/pagination.md)
- [Errors](../docs/api/errors.md)

## License

MIT
