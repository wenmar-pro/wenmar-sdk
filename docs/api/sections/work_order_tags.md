# Work Order Tags

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List work order tags

```
GET /work_order_tags
```

List all work order tags, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_order_tags.json
```

## Create work order tag

```
POST /work_order_tags
```

Create a work order tag.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_order_tags.json
```

## Update work order tag

```
PATCH /work_order_tags/{id}
```

Update a work order tag by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_order_tags/<id>.json
```

## Update work order tags archive

```
PATCH /work_order_tags/{id}/archive
```

Update a work order tags archive by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_order_tags/<id>.json
```

---

### CreateWorkOrderTagRequest schema {#createworkordertagrequest-schema}

| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `color` | string | Yes |

---

### UpdateWorkOrderTagRequest schema {#updateworkordertagrequest-schema}

| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

