# Inventory Levels

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List inventory levels

```
GET /inventory_levels
```

List all inventory levels, paginated via the Link header.

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
| `on_hand` | integer | Yes |
| `reorder_point` | string \| null | Yes |
| `bin_location` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inventory_levels.json
```

## Create inventory level

```
POST /inventory_levels
```

Create a inventory level.

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
| `on_hand` | integer | Yes |
| `reorder_point` | integer \| null | Yes |
| `bin_location` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inventory_levels.json
```

## List inventory levels barcode lookup

```
GET /inventory_levels/barcode_lookup
```

List all inventory levels barcode lookup, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `barcode` | string | No |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `part_number` | string | Yes |
| `description` | string | Yes |
| `brand` | string | Yes |
| `part_type` | string | Yes |
| `barcode` | string | Yes |
| `stocked` | boolean | Yes |
| `cost_cents` | integer | Yes |
| `sell_cents` | integer | Yes |
| `taxable` | boolean | Yes |
| `vendor` | object | Yes |
| `on_hand` | integer | Yes |
| `reorder_point` | integer \| null | Yes |
| `bin_location` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inventory_levels/barcode_lookup.json
```

## Delete inventory level

```
DELETE /inventory_levels/{id}
```

Delete a inventory level by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inventory_levels/<id>.json
```

## Show inventory level

```
GET /inventory_levels/{id}
```

Show a inventory level by ID.

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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inventory_levels/<id>.json
```

## Update inventory level

```
PATCH /inventory_levels/{id}
```

Update a inventory level by ID.

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
| `on_hand` | integer | Yes |
| `reorder_point` | integer \| null | Yes |
| `bin_location` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inventory_levels/<id>.json
```

## Update inventory levels stock

```
PATCH /inventory_levels/{id}/stock
```

Update a inventory levels stock by ID.

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
| `on_hand` | integer | Yes |
| `reorder_point` | integer \| null | Yes |
| `bin_location` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inventory_levels/<id>.json
```

---

### CreateInventoryLevelRequest schema {#createinventorylevelrequest-schema}

| Field | Type | Required |
|---|---|---|
| `part` | object | Yes |

`part` — object:
| Field | Type | Required |
|---|---|---|
| `part_number` | string | Yes |
| `description` | string | Yes |
| `brand` | string | Yes |
| `part_type` | string | Yes |
| `stocked` | boolean | Yes |
| `initial_quantity` | integer | Yes |
| `cost` | string | Yes |
| `sell` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateInventoryLevelRequest schema {#updateinventorylevelrequest-schema}

| Field | Type | Required |
|---|---|---|
| `part` | object | Yes |

`part` — object:
| Field | Type | Required |
|---|---|---|
| `description` | string | Yes |

---

### UpdateInventoryLevelsStockRequest schema {#updateinventorylevelsstockrequest-schema}

| Field | Type | Required |
|---|---|---|
| `inventory_level` | object | Yes |

`inventory_level` — object:
| Field | Type | Required |
|---|---|---|
| `on_hand` | integer | Yes |
| `reason` | string | Yes |

