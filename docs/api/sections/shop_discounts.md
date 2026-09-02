# Shop Discounts

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List shop discounts

```
GET /shop_discounts
```

List all shop discounts, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `discount_type` | string | Yes |
| `amount_cents` | integer | Yes |
| `percentage` | number \| null | Yes |
| `active` | boolean | Yes |
| `category` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/shop_discounts.json
```

## Create shop discount

```
POST /shop_discounts
```

Create a shop discount.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `discount_type` | string | Yes |
| `amount_cents` | integer | Yes |
| `percentage` | string | Yes |
| `active` | boolean | Yes |
| `category` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/shop_discounts.json
```

## Update shop discount

```
PATCH /shop_discounts/{id}
```

Update a shop discount by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `discount_type` | string | Yes |
| `amount_cents` | integer | Yes |
| `percentage` | number \| null | Yes |
| `active` | boolean | Yes |
| `category` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/shop_discounts/<id>.json
```

---

### CreateShopDiscountRequest schema {#createshopdiscountrequest-schema}

| Field | Type | Required |
|---|---|---|
| `shop_discount_config` | object | Yes |

`shop_discount_config` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `discount_type` | string | Yes |
| `percentage` | integer | Yes |
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

### UpdateShopDiscountRequest schema {#updateshopdiscountrequest-schema}

| Field | Type | Required |
|---|---|---|
| `shop_discount_config` | object | Yes |

`shop_discount_config` — object:
| Field | Type | Required |
|---|---|---|
| `active` | boolean | Yes |

