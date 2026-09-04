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
| `recurrence_rule` | string \| null | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

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
| `recurrence_rule` | string \| null | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
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

## List expenses data transfer

```
GET /expenses/data_transfer
```

List all expenses data transfer, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `resource` | string | Yes |
| `template_url` | string | Yes |
| `export_url` | string | Yes |
| `validate_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/expenses/data_transfer.json
```

## Create expenses export

```
POST /expenses/export
```

Create a expenses export.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `export_log_id` | integer | Yes |
| `status` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/expenses/export.json
```

## List expenses export download

```
GET /expenses/export/{id}/download
```

List all expenses export download, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — no content.

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/expenses/export/<id>.json
```

## Create expenses imports commit

```
POST /expenses/imports/commit
```

Create a expenses imports commit.

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/expenses/imports/commit.json
```

## List expenses imports template

```
GET /expenses/imports/template
```

List all expenses imports template, paginated via the Link header.

**Response 200** — no content.

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/expenses/imports/template.json
```

## Create expenses imports validate

```
POST /expenses/imports/validate
```

Create a expenses imports validate.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `valid_count` | integer | Yes |
| `error_count` | integer | Yes |
| `duplicate_count` | integer | Yes |
| `total` | integer | Yes |
| `errors` | array of any | Yes |

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/expenses/imports/validate.json
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

## Show expense

```
GET /expenses/{id}
```

Show a expense by ID.

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
| `recurrence_rule` | string \| null | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/expenses/<id>.json
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
| `recurrence_rule` | string \| null | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
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

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

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

### UpdateExpenseRequest schema {#updateexpenserequest-schema}

| Field | Type | Required |
|---|---|---|
| `expense` | object | Yes |

`expense` — object:
| Field | Type | Required |
|---|---|---|
| `receipt_id` | string | No |
| `description` | string | No |

