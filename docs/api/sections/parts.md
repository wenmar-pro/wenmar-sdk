# Parts

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List parts

```
GET /parts
```

List all parts, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `brand` | string | No |
| `page` | integer | No |
| `per_page` | integer | No |
| `q` | string | No |
| `stock_status` | string | No |
| `stocked` | boolean | No |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `part_number` | string | Yes |
| `description` | string | Yes |
| `brand` | string | Yes |
| `part_type` | string | Yes |
| `barcode` | string \| null | Yes |
| `stocked` | boolean | Yes |
| `cost_cents` | integer | Yes |
| `sell_cents` | integer | Yes |
| `taxable` | boolean | Yes |
| `vendor` | object | Yes |
| `vendor_id` | integer \| null | Yes |
| `core_charge_cents` | integer | Yes |
| `on_hand` | integer | Yes |
| `reorder_point` | string \| null | Yes |
| `bin_location` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/parts.json
```

## Create part

```
POST /parts
```

Create a part.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `part_number` | string | Yes |
| `description` | string | Yes |
| `brand` | string | Yes |
| `part_type` | string | Yes |
| `barcode` | string \| null | Yes |
| `stocked` | boolean | Yes |
| `cost_cents` | integer | Yes |
| `sell_cents` | integer | Yes |
| `taxable` | boolean | Yes |
| `vendor` | object | Yes |
| `vendor_id` | integer \| null | Yes |
| `core_charge_cents` | integer | Yes |
| `on_hand` | integer | Yes |
| `reorder_point` | integer \| null | Yes |
| `bin_location` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`vendor` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/parts.json
```

## Delete part

```
DELETE /parts/{id}
```

Delete a part by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/parts/<id>.json
```

## Show part

```
GET /parts/{id}
```

Show a part by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `part_number` | string | Yes |
| `description` | string | Yes |
| `brand` | string | Yes |
| `part_type` | string | Yes |
| `barcode` | string \| null | Yes |
| `stocked` | boolean | Yes |
| `cost_cents` | integer | Yes |
| `sell_cents` | integer | Yes |
| `taxable` | boolean | Yes |
| `vendor` | object | Yes |
| `vendor_id` | integer \| null | Yes |
| `core_charge_cents` | integer | Yes |
| `on_hand` | integer | Yes |
| `reorder_point` | string | Yes |
| `bin_location` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 304** — no content.

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/parts/<id>.json
```

## Update part

```
PATCH /parts/{id}
```

Update a part by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `part_number` | string | Yes |
| `description` | string | Yes |
| `brand` | string | Yes |
| `part_type` | string | Yes |
| `barcode` | string \| null | Yes |
| `stocked` | boolean | Yes |
| `cost_cents` | integer | Yes |
| `sell_cents` | integer | Yes |
| `taxable` | boolean | Yes |
| `vendor` | object | Yes |
| `vendor_id` | integer \| null | Yes |
| `core_charge_cents` | integer | Yes |
| `on_hand` | integer | Yes |
| `reorder_point` | integer \| null | Yes |
| `bin_location` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/parts/<id>.json
```

---

### CreatePartRequest schema {#createpartrequest-schema}

| Field | Type | Required |
|---|---|---|
| `part` | object | Yes |

`part` — object:
| Field | Type | Required |
|---|---|---|
| `part_number` | string | Yes |
| `description` | string | Yes |
| `brand` | string | Yes |
| `part_type` | string | No |
| `stocked` | boolean | Yes |
| `initial_quantity` | integer | Yes |
| `cost` | string | Yes |
| `sell` | string | Yes |
| `taxable` | boolean | No |
| `vendor_id` | integer | No |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdatePartRequest schema {#updatepartrequest-schema}

| Field | Type | Required |
|---|---|---|
| `part` | object | Yes |

`part` — object:
| Field | Type | Required |
|---|---|---|
| `description` | string | Yes |

