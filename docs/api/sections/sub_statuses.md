# Sub Statuses

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List sub statuses

```
GET /sub_statuses
```

List all sub statuses, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `status_scope` | string | Yes |
| `active` | boolean | Yes |
| `is_default` | boolean | Yes |
| `position` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/sub_statuses.json
```

## Create sub statuse

```
POST /sub_statuses
```

Create a sub statuse.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `status_scope` | string | Yes |
| `active` | boolean | Yes |
| `is_default` | boolean | Yes |
| `position` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/sub_statuses.json
```

## Update sub statuse

```
PATCH /sub_statuses/{id}
```

Update a sub statuse by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `status_scope` | string | Yes |
| `active` | boolean | Yes |
| `is_default` | boolean | Yes |
| `position` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/sub_statuses/<id>.json
```

---

### CreateSubStatuseRequest schema {#createsubstatuserequest-schema}

| Field | Type | Required |
|---|---|---|
| `sub_status_type` | object | Yes |

`sub_status_type` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `color` | string | Yes |
| `status_scope` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateSubStatuseRequest schema {#updatesubstatuserequest-schema}

| Field | Type | Required |
|---|---|---|
| `sub_status_type` | object | Yes |

`sub_status_type` — object:
| Field | Type | Required |
|---|---|---|
| `active` | boolean | Yes |

