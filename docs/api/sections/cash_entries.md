# Cash Entries

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List cash entries

```
GET /cash_entries
```

List all cash entries, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `cash_drawer_session_id` | integer | No |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `cash_drawer_session_id` | integer | Yes |
| `payment_id` | integer \| null | Yes |
| `amount_cents` | integer | Yes |
| `direction` | string | Yes |
| `entry_type` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `reference` | string \| null | Yes |
| `recipient` | string | Yes |
| `processed_at` | string | Yes |
| `reversed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `entered_by` | object | Yes |
| `reversed_by` | any | Yes |
| `url` | string | Yes |

`entered_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/cash_entries.json
```

## Create cash entrie

```
POST /cash_entries
```

Create a cash entrie.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `cash_drawer_session_id` | integer | Yes |
| `payment_id` | integer \| null | Yes |
| `amount_cents` | integer | Yes |
| `direction` | string | Yes |
| `entry_type` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `reference` | string | Yes |
| `recipient` | string | Yes |
| `processed_at` | string | Yes |
| `reversed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `entered_by` | object | Yes |
| `reversed_by` | any | Yes |
| `url` | string | Yes |

`entered_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `cash_drawer_session_id` | integer | Yes |
| `payment_id` | integer \| null | Yes |
| `amount_cents` | integer | Yes |
| `direction` | string | Yes |
| `entry_type` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `reference` | string \| null | Yes |
| `recipient` | string | Yes |
| `processed_at` | string | Yes |
| `reversed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `entered_by` | object | Yes |
| `reversed_by` | any | Yes |
| `url` | string | Yes |

`entered_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/cash_entries.json
```

## Delete cash entrie

```
DELETE /cash_entries/{id}
```

Delete a cash entrie by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `cash_drawer_session_id` | integer | Yes |
| `payment_id` | integer \| null | Yes |
| `amount_cents` | integer | Yes |
| `direction` | string | Yes |
| `entry_type` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `reference` | string \| null | Yes |
| `recipient` | string | Yes |
| `processed_at` | string | Yes |
| `reversed_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `entered_by` | object | Yes |
| `reversed_by` | object | Yes |
| `url` | string | Yes |

`entered_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

`reversed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/cash_entries/<id>.json
```

## Show cash entrie

```
GET /cash_entries/{id}
```

Show a cash entrie by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `cash_drawer_session_id` | integer | Yes |
| `payment_id` | integer \| null | Yes |
| `amount_cents` | integer | Yes |
| `direction` | string | Yes |
| `entry_type` | string | Yes |
| `category` | string | Yes |
| `description` | string | Yes |
| `reference` | string \| null | Yes |
| `recipient` | string | Yes |
| `processed_at` | string | Yes |
| `reversed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `entered_by` | object | Yes |
| `reversed_by` | any | Yes |
| `url` | string | Yes |

`entered_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/cash_entries/<id>.json
```

---

### CreateCashEntrieRequest schema {#createcashentrierequest-schema}

| Field | Type | Required |
|---|---|---|
| `cash_entry` | object | Yes |

`cash_entry` — object:
| Field | Type | Required |
|---|---|---|
| `reference` | string | No |
| `amount_cents` | integer | Yes |
| `entry_type` | string | Yes |
| `description` | string | Yes |
| `recipient` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

