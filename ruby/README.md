# wenmar

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

## License

MIT
