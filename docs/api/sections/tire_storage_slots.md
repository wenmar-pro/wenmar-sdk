# Tire Storage Slots

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List tire storage slots

```
GET /tire_storage_slots
```

List all tire storage slots, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `slot_label` | string | Yes |
| `season` | string | Yes |
| `season_label` | string | Yes |
| `tire_set_description` | string \| null | Yes |
| `stored_at` | string | Yes |
| `released_at` | string \| null | Yes |
| `currently_stored` | boolean | Yes |
| `storage_fee_cents` | integer \| null | Yes |
| `fee_type` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tire_storage_slots.json
```

## Create tire storage slot

```
POST /tire_storage_slots
```

Create a tire storage slot.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `slot_label` | string | Yes |
| `season` | string | Yes |
| `season_label` | string | Yes |
| `tire_set_description` | string \| null | Yes |
| `stored_at` | string | Yes |
| `released_at` | string \| null | Yes |
| `currently_stored` | boolean | Yes |
| `storage_fee_cents` | integer \| null | Yes |
| `fee_type` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
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
     -d '{"...":"..."}' https://app.wenmarpro.com/tire_storage_slots.json
```

## Create tire storage slots export

```
POST /tire_storage_slots/export
```

Create a tire storage slots export.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `export_log_id` | integer | Yes |
| `status` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/tire_storage_slots/export.json
```

## List tire storage slots export download

```
GET /tire_storage_slots/export/{id}/download
```

List all tire storage slots export download, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — no content.

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tire_storage_slots/export/<id>.json
```

## Delete tire storage slot

```
DELETE /tire_storage_slots/{id}
```

Delete a tire storage slot by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tire_storage_slots/<id>.json
```

## Show tire storage slot

```
GET /tire_storage_slots/{id}
```

Show a tire storage slot by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `slot_label` | string | Yes |
| `season` | string | Yes |
| `season_label` | string | Yes |
| `tire_set_description` | string \| null | Yes |
| `stored_at` | string | Yes |
| `released_at` | string \| null | Yes |
| `currently_stored` | boolean | Yes |
| `storage_fee_cents` | integer \| null | Yes |
| `fee_type` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tire_storage_slots/<id>.json
```

## Update tire storage slot

```
PATCH /tire_storage_slots/{id}
```

Update a tire storage slot by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `slot_label` | string | Yes |
| `season` | string | Yes |
| `season_label` | string | Yes |
| `tire_set_description` | string | Yes |
| `stored_at` | string | Yes |
| `released_at` | string \| null | Yes |
| `currently_stored` | boolean | Yes |
| `storage_fee_cents` | integer \| null | Yes |
| `fee_type` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
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
     -d '{"...":"..."}' https://app.wenmarpro.com/tire_storage_slots/<id>.json
```

## Create tire storage slots check out

```
POST /tire_storage_slots/{tire_storage_slot_id}/check_outs
```

Create a tire storage slots check out.

| Param | Type | Required |
|---|---|---|
| `tire_storage_slot_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `slot_label` | string | Yes |
| `season` | string | Yes |
| `season_label` | string | Yes |
| `tire_set_description` | string \| null | Yes |
| `stored_at` | string | Yes |
| `released_at` | string | Yes |
| `currently_stored` | boolean | Yes |
| `storage_fee_cents` | integer \| null | Yes |
| `fee_type` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
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
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/tire_storage_slots/{tire_storage_slot_id}/check_outs.json
```

---

### CreateTireStorageSlotRequest schema {#createtirestorageslotrequest-schema}

| Field | Type | Required |
|---|---|---|
| `tire_storage_slot` | object | Yes |

`tire_storage_slot` — object:
| Field | Type | Required |
|---|---|---|
| `vehicle_id` | integer | Yes |
| `customer_id` | integer | Yes |
| `slot_label` | string | Yes |
| `season` | string | Yes |
| `stored_at` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateTireStorageSlotRequest schema {#updatetirestorageslotrequest-schema}

| Field | Type | Required |
|---|---|---|
| `tire_storage_slot` | object | Yes |

`tire_storage_slot` — object:
| Field | Type | Required |
|---|---|---|
| `tire_set_description` | string | Yes |

