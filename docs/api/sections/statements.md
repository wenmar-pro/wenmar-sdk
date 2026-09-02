# Statements

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Show statement

```
GET /statements/{id}
```

Show a statement by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Statement](#statement-schema)

**Response 404** — [Error](#error-schema) error envelope

**Example**

```json
{
  "id": 1,
  "statement_number": "ST-0001",
  "status": "sent",
  "statement_date": "2026-08-01",
  "start_date": "2026-07-01",
  "end_date": "2026-07-31",
  "due_date": "2026-08-15",
  "totals": {
    "previous_balance_cents": 0,
    "new_charges_cents": 12000,
    "payments_received_cents": 0,
    "credits_cents": 0,
    "balance_due_cents": 12000,
    "currency": "CAD"
  },
  "sent_at": "2026-08-01T12:00:00.000-04:00",
  "viewed_at": null,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "created_at": "2026-08-01T12:00:00.000-04:00",
  "updated_at": "2026-08-01T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/statements/1.json",
  "app_url": "https://app.wenmarpro.com/statements/1"
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/statements/<id>.json
```

## List statements payments

```
GET /statements/{statement_id}/payments
```

List all statements payments, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `statement_id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/statements/{statement_id}/payments.json
```

---

### Statement schema {#statement-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `statement_number` | string | Yes |
| `status` | string | Yes |
| `statement_date` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `due_date` | string | Yes |
| `totals` | object | Yes |
| `sent_at` | string \| null | Yes |
| `viewed_at` | string \| null | Yes |
| `customer` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`totals` — object:
| Field | Type | Required |
|---|---|---|
| `previous_balance_cents` | integer | Yes |
| `new_charges_cents` | integer | Yes |
| `payments_received_cents` | integer | Yes |
| `credits_cents` | integer | Yes |
| `balance_due_cents` | integer | Yes |
| `currency` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

