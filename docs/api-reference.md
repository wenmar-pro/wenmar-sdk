# Wenmar Pro API Reference

Base URL: `https://app.wenmarpro.com`

All requests require a bearer token (see [authentication](authentication.md)).

- [List customers](#list-customers)
- [Create customer](#create-customer)
- [Show customer](#show-customer)
- [Show vehicle](#show-vehicle)
- [List work orders](#list-work-orders)
- [Show work order](#show-work-order)

## List customers

```
GET /api/customers
```

Paginated. Response carries a `Link` header for the next page.

| Param | Type | Required |
|---|---|---|
| `page` | integer | No |

**Response 200**

```json
{
  "data": [
    { "id": 1, "full_name": "Jane Doe", "email": "jane@example.com", "phone": "+15551234567" }
  ]
}
```

```bash
curl "https://app.wenmarpro.com/api/customers" \
  -H "Authorization: Bearer YOUR_API_TOKEN"
```

## Create customer

```
POST /api/customers
```

**Request body** — `customer` object:

| Field | Type | Required |
|---|---|---|
| `full_name` | string | Yes |
| `email` | string | No |
| `phone` | string | No |

**Response 201**

```json
{ "data": { "id": 1, "full_name": "Jane Doe", "email": "jane@example.com", "phone": "+15551234567" } }
```

```bash
curl "https://app.wenmarpro.com/api/customers" \
  -X POST -H "Authorization: Bearer YOUR_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"customer":{"full_name":"Jane Doe"}}'
```

**422 Validation error** — missing `full_name`

```json
{ "error": { "code": "validation_failed", "message": "Full name can't be blank", "details": { "full_name": ["can't be blank"] } } }
```

## Show customer

```
GET /api/customers/{id}
```

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

```json
{ "data": { "id": 1, "full_name": "Jane Doe", "email": "jane@example.com", "phone": "+15551234567" } }
```

**404** — customer not found

```json
{ "error": { "code": "not_found", "message": "Customer not found", "details": {} } }
```

## Show vehicle

```
GET /api/vehicles/{id}
```

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

```json
{ "data": { "id": 1, "make": "Toyota", "model": "Camry", "year": 2020, "vin": "ABC123" } }
```

## List work orders

```
GET /api/work_orders
```

Paginated via the `Link` header.

| Param | Type | Required |
|---|---|---|
| `page` | integer | No |

**Response 200**

```json
{
  "data": [
    {
      "id": 1,
      "status": "open",
      "customer": { "id": 1, "full_name": "Jane Doe" },
      "vehicle": { "id": 1, "make": "Toyota", "model": "Camry" }
    }
  ]
}
```

## Show work order

```
GET /api/work_orders/{id}
```

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

```json
{
  "data": {
    "id": 1,
    "status": "open",
    "customer": { "id": 1, "full_name": "Jane Doe" },
    "vehicle": { "id": 1, "make": "Toyota", "model": "Camry" }
  }
}
```

**404** — work order not found

```json
{ "error": { "code": "not_found", "message": "Work order not found", "details": {} } }
```
