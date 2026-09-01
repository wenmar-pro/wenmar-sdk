# Expenses

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List expenses

```
GET /expenses
```

List all expenses, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `payee` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `amount_cents` | integer | Yes |
| `amount_currency` | string | Yes |
| `expense_date` | string | Yes |
| `payment_method` | string | Yes |
| `recurring` | boolean | Yes |
| `recurrence_rule` | any | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `creator` | object | No |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/expenses.json
```

## Create expense

```
POST /expenses
```

Create a expense.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `payee` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `amount_cents` | integer | Yes |
| `amount_currency` | string | Yes |
| `expense_date` | string | Yes |
| `payment_method` | string | Yes |
| `recurring` | boolean | Yes |
| `recurrence_rule` | any | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `creator` | object | No |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/expenses.json
```

## Delete expense

```
DELETE /expenses/{id}
```

Delete a expense by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/expenses/<id>.json
```

## Update expense

```
PATCH /expenses/{id}
```

Update a expense by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `payee` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `amount_cents` | integer | Yes |
| `amount_currency` | string | Yes |
| `expense_date` | string | Yes |
| `payment_method` | string | Yes |
| `recurring` | boolean | Yes |
| `recurrence_rule` | any | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `creator` | object | No |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/expenses/<id>.json
```

---

### CreateExpenseRequest schema {#createexpenserequest-schema}

| Field | Type | Required |
|---|---|---|
| `expense` | object | Yes |

`expense` — object:
| Field | Type | Required |
|---|---|---|
| `payee` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `amount` | integer | Yes |
| `expense_date` | string | Yes |
| `payment_method` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateExpenseRequest schema {#updateexpenserequest-schema}

| Field | Type | Required |
|---|---|---|
| `expense` | object | Yes |

`expense` — object:
| Field | Type | Required |
|---|---|---|
| `description` | string | Yes |

