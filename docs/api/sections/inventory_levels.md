# Inventory Levels

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

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

## Create inventory level extraction

```
POST /inventory_levels/extractions
```

Create a inventory level extraction.

**Response 202**

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |
| `stream_id` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inventory_levels/extractions.json
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
| `part_id` | integer | Yes |
| `location_id` | integer | Yes |
| `on_hand` | number | Yes |
| `available_quantity` | number | Yes |
| `quantity_on_order` | integer | Yes |
| `bin_location` | string | Yes |
| `reorder_point` | integer \| null | Yes |
| `max_stock` | any | Yes |
| `stock_status` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `part` | object | Yes |
| `recent_movements` | array of object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`part` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |

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
| `part_id` | integer | Yes |
| `location_id` | integer | Yes |
| `on_hand` | number | Yes |
| `available_quantity` | number | Yes |
| `quantity_on_order` | integer | Yes |
| `bin_location` | string | Yes |
| `reorder_point` | string | Yes |
| `max_stock` | string | Yes |
| `stock_status` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `part` | object | Yes |
| `recent_movements` | array of any | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`part` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inventory_levels/<id>.json
```

## Create inventory levels adjust

```
POST /inventory_levels/{id}/adjust
```

Create a inventory levels adjust.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `part_id` | integer | Yes |
| `location_id` | integer | Yes |
| `on_hand` | number | Yes |
| `available_quantity` | number | Yes |
| `quantity_on_order` | integer | Yes |
| `bin_location` | string | Yes |
| `reorder_point` | integer \| null | Yes |
| `max_stock` | any | Yes |
| `stock_status` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `part` | object | Yes |
| `recent_movements` | array of any | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`part` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inventory_levels/<id>.json
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

### CreateInventoryLevelExtractionRequest schema {#createinventorylevelextractionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `text` | string | Yes |
| `extraction_id` | string | Yes |

---

### UpdateInventoryLevelRequest schema {#updateinventorylevelrequest-schema}

| Field | Type | Required |
|---|---|---|
| `inventory_level` | object | Yes |

`inventory_level` — object:
| Field | Type | Required |
|---|---|---|
| `bin_location` | string | Yes |
| `reorder_point` | integer | Yes |
| `max_stock` | integer | Yes |

---

### CreateInventoryLevelsAdjustRequest schema {#createinventorylevelsadjustrequest-schema}

| Field | Type | Required |
|---|---|---|
| `adjustment` | object | Yes |

`adjustment` — object:
| Field | Type | Required |
|---|---|---|
| `quantity_delta` | integer | Yes |
| `reason` | string | Yes |

