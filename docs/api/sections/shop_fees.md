# Shop Fees

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List shop fees

```
GET /shop_fees
```

List all shop fees, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `fee_type` | string | Yes |
| `amount_cents` | integer | Yes |
| `percentage` | string \| null | Yes |
| `active` | boolean | Yes |
| `applies_to` | string \| null | Yes |
| `is_taxable` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/shop_fees.json
```

## Create shop fee

```
POST /shop_fees
```

Create a shop fee.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `fee_type` | string | Yes |
| `amount_cents` | integer | Yes |
| `percentage` | number \| null | Yes |
| `active` | boolean | Yes |
| `applies_to` | string | Yes |
| `is_taxable` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/shop_fees.json
```

## Update shop fee

```
PATCH /shop_fees/{id}
```

Update a shop fee by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `fee_type` | string | Yes |
| `amount_cents` | integer | Yes |
| `percentage` | number \| null | Yes |
| `active` | boolean | Yes |
| `applies_to` | string | Yes |
| `is_taxable` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/shop_fees/<id>.json
```

---

### CreateShopFeeRequest schema {#createshopfeerequest-schema}

| Field | Type | Required |
|---|---|---|
| `shop_fee_config` | object | Yes |

`shop_fee_config` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `fee_type` | string | Yes |
| `amount` | number | Yes |
| `active` | boolean | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateShopFeeRequest schema {#updateshopfeerequest-schema}

| Field | Type | Required |
|---|---|---|
| `shop_fee_config` | object | Yes |

`shop_fee_config` — object:
| Field | Type | Required |
|---|---|---|
| `active` | boolean | Yes |

