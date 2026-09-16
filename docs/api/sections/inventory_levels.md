# Inventory Levels

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

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

### CreateInventoryLevelExtractionRequest schema {#createinventorylevelextractionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `text` | string | Yes |
| `extraction_id` | string | Yes |

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

