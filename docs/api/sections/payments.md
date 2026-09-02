# Payments

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List payments

```
GET /payments
```

List all payments, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `method` | string | No |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `currency` | string | Yes |
| `method` | string | Yes |
| `processor_status` | string | Yes |
| `is_refund` | boolean | Yes |
| `is_adjustment` | boolean | Yes |
| `voided` | boolean | Yes |
| `voided_at` | string \| null | Yes |
| `processed_at` | string | Yes |
| `reference` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `work_order_id` | integer | Yes |
| `customer_id` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `work_order` | object | Yes |
| `customer` | object | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/payments.json
```

## List payments pending

```
GET /payments/pending
```

List all payments pending, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `currency` | string | Yes |
| `method` | string | Yes |
| `processor_status` | string | Yes |
| `is_refund` | boolean | Yes |
| `is_adjustment` | boolean | Yes |
| `voided` | boolean | Yes |
| `voided_at` | string \| null | Yes |
| `processed_at` | string | Yes |
| `reference` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `work_order_id` | integer | Yes |
| `customer_id` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `work_order` | object | Yes |
| `customer` | object | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/payments/pending.json
```

## Show payment

```
GET /payments/{id}
```

Show a payment by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `currency` | string | Yes |
| `method` | string | Yes |
| `processor_status` | string | Yes |
| `is_refund` | boolean | Yes |
| `is_adjustment` | boolean | Yes |
| `voided` | boolean | Yes |
| `voided_at` | string \| null | Yes |
| `processed_at` | string | Yes |
| `reference` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `work_order_id` | integer | Yes |
| `customer_id` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `work_order` | object | Yes |
| `customer` | object | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/payments/<id>.json
```

## Create payments cancellation

```
POST /payments/{id}/cancellation
```

Create a payments cancellation.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `currency` | string | Yes |
| `method` | string | Yes |
| `processor_status` | string | Yes |
| `is_refund` | boolean | Yes |
| `is_adjustment` | boolean | Yes |
| `voided` | boolean | Yes |
| `voided_at` | string | Yes |
| `processed_at` | string | Yes |
| `reference` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `work_order_id` | integer | Yes |
| `customer_id` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `work_order` | object | Yes |
| `customer` | object | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/payments/<id>.json
```

## Create payments confirmation

```
POST /payments/{id}/confirmation
```

Create a payments confirmation.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `currency` | string | Yes |
| `method` | string | Yes |
| `processor_status` | string | Yes |
| `is_refund` | boolean | Yes |
| `is_adjustment` | boolean | Yes |
| `voided` | boolean | Yes |
| `voided_at` | string \| null | Yes |
| `processed_at` | string | Yes |
| `reference` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `work_order_id` | integer | Yes |
| `customer_id` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `work_order` | object | Yes |
| `customer` | object | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/payments/<id>.json
```

## Create payments failure

```
POST /payments/{id}/failure
```

Create a payments failure.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `currency` | string | Yes |
| `method` | string | Yes |
| `processor_status` | string | Yes |
| `is_refund` | boolean | Yes |
| `is_adjustment` | boolean | Yes |
| `voided` | boolean | Yes |
| `voided_at` | string \| null | Yes |
| `processed_at` | string | Yes |
| `reference` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `work_order_id` | integer | Yes |
| `customer_id` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `work_order` | object | Yes |
| `customer` | object | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/payments/<id>.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

