# Customers

## List customers

```
GET /customers
```

Paginated via the `Link` header. Returns a bare array.

| Param | Type | Required |
|---|---|---|
| `page` | integer | No |

**Response 200**

```json
[
  { "type": "Customer", "id": 1, "full_name": "Jane Doe", "emails": [...], "phones": [...] }
]
```

```bash
curl "https://app.wenmarpro.com/customers" \
  -H "Authorization: Bearer YOUR_API_TOKEN"
```

## Create customer

```
POST /customers
```

**Request body** — `customer` object:

| Field | Type | Required |
|---|---|---|
| `first_name` | string | Yes |
| `last_name` | string | Yes |

**Response 201** — bare customer object (same shape as show).

```bash
curl "https://app.wenmarpro.com/customers" \
  -X POST -H "Authorization: Bearer YOUR_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"customer":{"first_name":"Jane","last_name":"Doe"}}'
```

## Show customer

```
GET /customers/{id}
```

**Response 200** — bare customer object:

```json
{
  "type": "Customer",
  "id": 1,
  "full_name": "Jane Doe",
  "company_name": null,
  "first_name": "Jane",
  "last_name": "Doe",
  "fleet_identifier": null,
  "marketing_opt_in": false,
  "tax_exempt": false,
  "vehicles_count": 2,
  "emails": [
    { "id": 1, "address": "jane@example.com", "label": "personal", "primary": true }
  ],
  "phones": [
    { "id": 1, "number": "+15551234567", "label": "mobile", "primary": true }
  ],
  "vehicles_url": "https://app.wenmarpro.com/vehicles.json?customer_id=1",
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?customer_id=1",
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/customers/1.json",
  "outstanding_balance_cents": 0,
  "total_revenue_cents": 50000,
  "store_credit_cents": 0,
  "last_visit_at": "2026-08-20T10:00:00.000-04:00",
  "statements_count": 2,
  "currency": "CAD"
}
```

## Update customer

```
PATCH /customers/{id}
```

**Request body** — `customer` object (any subset of fields).

**Response 200** — bare customer object (same shape as show).
