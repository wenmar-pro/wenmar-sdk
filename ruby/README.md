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
customers["data"] # => [{ "id" => 1, "full_name" => "Jane Doe", ... }]

# Show a customer
customer = client.show_customer(1)
customer["data"] # => { "id" => 1, "full_name" => "Jane Doe", ... }

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
| List customers | `list_customers(page: nil)` |
| Create customer | `create_customer(full_name: ..., email: ..., phone: ...)` |
| Show customer | `show_customer(id)` |
| Show vehicle | `show_vehicle(id)` |
| List work orders | `list_work_orders(page: nil)` |
| Show work order | `show_work_order(id)` |

All methods return the parsed JSON response body as a `Hash`.

## Pagination

List endpoints paginate via the RFC 5988 `Link` header. Call `.paginator` on
the result to walk pages:

```ruby
result = client.list_customers
while result.paginator.has_next?
  result = result.paginator.next_page
  result["data"] # => next page of records
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
  e.details
end
```

See [docs/errors.md](../docs/errors.md) for the full error envelope and code table.

## Retry

The client retries 5xx responses with exponential backoff (max 3 retries). It
respects the `Retry-After` response header and never retries 4xx errors.

## Documentation

- [API reference](../docs/api-reference.md)
- [Authentication](../docs/authentication.md)
- [Pagination](../docs/pagination.md)
- [Errors](../docs/errors.md)

## License

MIT
