# Pagination

List endpoints — `GET /customers` and `GET /work_orders` — paginate
results using [RFC 5988](https://www.rfc-editor.org/rfc/rfc5988) `Link`
headers.

## How it works

Each response includes a `Link` header pointing to the next page:

```
Link: <https://app.wenmarpro.com/customers?page=2>; rel="next"
```

When there is no next page, the `Link` header is absent (or carries no
`rel="next"`). Page numbers are 1-based; page 1 is the default.

The `next` URL includes query params (e.g. `?page=2`) that must be
followed — the SDK follows the actual URL rather than re-calling the
base endpoint.

## Go

Use the paginated helpers to walk all pages:

```go
resp, paginator, err := client.ListCustomersWithPagination(ctx)
for paginator.HasNext() {
    next, err := paginator.NextPage(ctx)
    // next is the decoded response from the next URL
}
```

The same pattern applies to `ListWorkOrdersWithPagination`.

## Ruby

List methods return the parsed body (a bare array). Call `.paginator`
on it to walk pages:

```ruby
result = client.list_customers
while result.paginator.has_next?
  result = result.paginator.next_page
  # result is the next page (a bare array of records)
end
```

## Response shape

List endpoints return bare arrays — no envelope:

```json
[
  { "id": 1, "full_name": "Jane Doe" },
  { "id": 2, "full_name": "John Smith" }
]
```
