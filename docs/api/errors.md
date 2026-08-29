# Errors

Every error response uses a consistent envelope:

```json
{
  "error": {
    "code": "not_found",
    "message": "Customer not found",
    "details": {}
  }
}
```

## Error codes

| HTTP status | `code` | Description |
|---|---|---|
| 401 | `unauthorized` | Missing or invalid API token |
| 403 | `forbidden` | Authorization failure |
| 404 | `not_found` | Resource not found |
| 422 | `validation_failed` | Request body failed validation |
| 429 | `rate_limited` | Too many requests — retry after delay |
| 500+ | `internal_error` | Server-side failure |

The `details` object carries per-field validation errors, e.g.
`{ "full_name": ["can't be blank"] }`.

## Retries

The SDKs retry 5xx responses with exponential backoff (max 3 retries) and
respect the `Retry-After` response header. 4xx errors are **not** retried.

## Go

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

## Ruby

Non-2xx responses raise `Wenmar::Error`:

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
