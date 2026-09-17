# Store Credits

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create store credits refund

```
POST /store_credits/{store_credit_id}/refunds
```

Create a store credits refund.

| Param | Type | Required |
|---|---|---|
| `store_credit_id` | integer | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `currency` | string | Yes |
| `method` | string | Yes |
| `processor_status` | string \| null | Yes |
| `is_refund` | boolean | Yes |
| `is_adjustment` | boolean | Yes |
| `voided` | boolean | Yes |
| `voided_at` | string \| null | Yes |
| `processed_at` | string | Yes |
| `reference` | string \| null | Yes |
| `notes` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `customer_id` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `customer` | object | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |

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

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/store_credits/{store_credit_id}/refunds.json
```

## Create store credits void

```
POST /store_credits/{store_credit_id}/voids
```

Create a store credits void.

| Param | Type | Required |
|---|---|---|
| `store_credit_id` | integer | Yes |

**Response 200** — [StoreCredit](#storecredit-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/store_credits/{store_credit_id}/voids.json
```

---

### CreateStoreCreditsRefundRequest schema {#createstorecreditsrefundrequest-schema}

| Field | Type | Required |
|---|---|---|
| `refund` | object | Yes |

`refund` — object:
| Field | Type | Required |
|---|---|---|
| `amount` | string | Yes |
| `method` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### StoreCredit schema {#storecredit-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `kind` | string | Yes |
| `amount_cents` | integer | Yes |
| `balance_cents` | integer | Yes |
| `reason` | string | Yes |
| `voided_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `voided_by` | any | Yes |
| `issued_by` | object | Yes |
| `source_payment` | any | Yes |
| `location` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `voids_url` | string | Yes |

`issued_by` — object:
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

