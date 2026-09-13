# Sub Statuses

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List sub statuses

```
GET /sub_statuses
```

List all sub statuses, paginated via the Link header.

**Response 200** — array of [SubStatusType](#substatustype-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/sub_statuses.json
```

## Create sub status

```
POST /sub_statuses
```

Create a sub status.

**Response 201** — [SubStatusType](#substatustype-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/sub_statuses.json
```

## Update sub status

```
PATCH /sub_statuses/{id}
```

Update a sub status by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [SubStatusType](#substatustype-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/sub_statuses/<id>.json
```

---

### SubStatusType schema {#substatustype-schema}

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

---

### CreateSubStatusRequest schema {#createsubstatusrequest-schema}

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

### UpdateSubStatusRequest schema {#updatesubstatusrequest-schema}

| Field | Type | Required |
|---|---|---|
| `sub_status_type` | object | Yes |

`sub_status_type` — object:
| Field | Type | Required |
|---|---|---|
| `active` | boolean | Yes |

