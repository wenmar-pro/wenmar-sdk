# Inspections

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List inspections

```
GET /inspections
```

List all inspections, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `per_page` | integer | No |

**Response 200** — array of [Inspection](#inspection-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inspections.json
```

## Create inspection

```
POST /inspections
```

Create a inspection.

**Response 201** — [Inspection](#inspection-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections.json
```

## Show inspection

```
GET /inspections/{id}
```

Show a inspection by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Inspection](#inspection-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inspections/<id>.json
```

## Update inspection

```
PATCH /inspections/{id}
```

Update a inspection by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Inspection](#inspection-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

## Archive inspection

```
PATCH /inspections/{id}/archive
```

Archive

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `is_default` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `groups` | array of any | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

## Remove default inspection

```
PATCH /inspections/{id}/remove_default
```

Remove default

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `is_default` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `groups` | array of any | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

## Restore inspection

```
PATCH /inspections/{id}/restore
```

Restore

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `is_default` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `groups` | array of any | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

## Set default inspection

```
PATCH /inspections/{id}/set_default
```

Set default

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `is_default` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `groups` | array of any | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

## Trash inspection

```
PATCH /inspections/{id}/trash
```

Trash

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string | Yes |
| `is_default` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `groups` | array of any | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

## Create inspections group

```
POST /inspections/{inspection_id}/groups
```

Create a inspections group.

| Param | Type | Required |
|---|---|---|
| `inspection_id` | integer | Yes |

**Response 201** — [Inspection](#inspection-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/{inspection_id}/groups.json
```

## Delete inspections group

```
DELETE /inspections/{inspection_id}/groups/{id}
```

Delete a inspections group by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `inspection_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inspections/<id>.json
```

## Update inspections group

```
PATCH /inspections/{inspection_id}/groups/{id}
```

Update a inspections group by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `inspection_id` | integer | Yes |

**Response 200** — [Inspection](#inspection-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

## Create inspections item

```
POST /inspections/{inspection_id}/items
```

Create a inspections item.

| Param | Type | Required |
|---|---|---|
| `inspection_id` | integer | Yes |

**Response 201** — [Inspection](#inspection-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/{inspection_id}/items.json
```

## Delete inspections item

```
DELETE /inspections/{inspection_id}/items/{id}
```

Delete a inspections item by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `inspection_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inspections/<id>.json
```

## Update inspections item

```
PATCH /inspections/{inspection_id}/items/{id}
```

Update a inspections item by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `inspection_id` | integer | Yes |

**Response 200** — [Inspection](#inspection-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

## Create inspections preset

```
POST /inspections/{inspection_id}/presets
```

Create a inspections preset.

| Param | Type | Required |
|---|---|---|
| `inspection_id` | integer | Yes |
| `item_id` | integer | No |

**Response 201** — [Inspection](#inspection-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/{inspection_id}/presets.json
```

## Delete inspections preset

```
DELETE /inspections/{inspection_id}/presets/{id}
```

Delete a inspections preset by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `inspection_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inspections/<id>.json
```

## Update inspections preset

```
PATCH /inspections/{inspection_id}/presets/{id}
```

Update a inspections preset by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `inspection_id` | integer | Yes |

**Response 200** — [Inspection](#inspection-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspections/<id>.json
```

---

### Inspection schema {#inspection-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `is_default` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `groups` | array of object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | No |
| `name` | string | No |
| `url` | string | No |

---

### CreateInspectionRequest schema {#createinspectionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `inspection` | object | Yes |

`inspection` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateInspectionRequest schema {#updateinspectionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `inspection` | object | Yes |

`inspection` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### CreateInspectionsGroupRequest schema {#createinspectionsgrouprequest-schema}

| Field | Type | Required |
|---|---|---|
| `inspection_group` | object | Yes |

`inspection_group` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### UpdateInspectionsGroupRequest schema {#updateinspectionsgrouprequest-schema}

| Field | Type | Required |
|---|---|---|
| `inspection_group` | object | Yes |

`inspection_group` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### CreateInspectionsItemRequest schema {#createinspectionsitemrequest-schema}

| Field | Type | Required |
|---|---|---|
| `inspection_item` | object | Yes |
| `group_id` | integer | Yes |

`inspection_item` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `requires_measurement` | boolean | Yes |
| `measurement_unit` | string | Yes |

---

### UpdateInspectionsItemRequest schema {#updateinspectionsitemrequest-schema}

| Field | Type | Required |
|---|---|---|
| `inspection_item` | object | Yes |

`inspection_item` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### CreateInspectionsPresetRequest schema {#createinspectionspresetrequest-schema}

| Field | Type | Required |
|---|---|---|
| `inspection_preset` | object | Yes |

`inspection_preset` — object:
| Field | Type | Required |
|---|---|---|
| `title` | string | Yes |
| `color_rating` | string | Yes |

---

### UpdateInspectionsPresetRequest schema {#updateinspectionspresetrequest-schema}

| Field | Type | Required |
|---|---|---|
| `inspection_preset` | object | Yes |

`inspection_preset` — object:
| Field | Type | Required |
|---|---|---|
| `title` | string | Yes |

