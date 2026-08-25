# Pagination

List endpoints — `GET /api/customers` and `GET /api/work_orders` — paginate
results using [RFC 5988](https://www.rfc-editor.org/rfc/rfc5988) `Link`
headers.

## How it works

Each response includes a `Link` header pointing to the next page:

```
Link: <https://app.wenmarpro.com/api/customers?page=2>; rel="next"
```

When there is no next page, the `Link` header is absent (or carries no
`rel="next"`). Page numbers are 1-based; page 1 is the default.

## Go

Use the paginated helpers to walk all pages:

```go
resp, paginator, err := client.ListCustomersWithPagination(ctx, nil)
for paginator.HasNext() {
    resp, err = paginator.NextPage(ctx)
    // resp holds the next page of customers
}
```

The same pattern applies to `ListWorkOrdersWithPagination`.

## Ruby

List methods return the parsed body. Call `.paginator` on it to walk pages:

```ruby
result = client.list_customers
while result.paginator.has_next?
  result = result.paginator.next_page
  result["data"] # => next page of records
end
```
