# Inspection Reports

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create inspection report

```
POST /inspection_reports
```

Create a inspection report.

| Param | Type | Required |
|---|---|---|
| `inspection_id` | integer | No |
| `work_order_id` | integer | No |

**Response 201** — [InspectionReport](#inspectionreport-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports.json
```

## Delete inspection report

```
DELETE /inspection_reports/{id}
```

Delete a inspection report by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Show inspection report

```
GET /inspection_reports/{id}
```

Show a inspection report by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [InspectionReport](#inspectionreport-schema)

**Response 302** — no content.

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Complete inspection report

```
PATCH /inspection_reports/{id}/complete
```

Complete

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `work_order_id` | integer | Yes |
| `quick_finding` | boolean | Yes |
| `published` | boolean | Yes |
| `completed` | boolean | Yes |
| `items_count` | integer | Yes |
| `checked_count` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Show inspection report group

```
GET /inspection_reports/{id}/group
```

Show a inspection report group by ID.

| Param | Type | Required |
|---|---|---|
| `group_name` | string | No |
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `group_name` | string | Yes |
| `inspection_report_id` | integer | Yes |
| `items` | array of object | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Mark all inspection report

```
POST /inspection_reports/{id}/mark_all
```

Mark all

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `work_order_id` | integer | Yes |
| `quick_finding` | boolean | Yes |
| `published` | boolean | Yes |
| `completed` | boolean | Yes |
| `items_count` | integer | Yes |
| `checked_count` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Publish inspection report

```
PATCH /inspection_reports/{id}/publish
```

Publish

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `work_order_id` | integer | Yes |
| `quick_finding` | boolean | Yes |
| `published` | boolean | Yes |
| `completed` | boolean | Yes |
| `items_count` | integer | Yes |
| `checked_count` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Reassign inspection report

```
PATCH /inspection_reports/{id}/reassign
```

Reassign

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `work_order_id` | integer | Yes |
| `quick_finding` | boolean | Yes |
| `published` | boolean | Yes |
| `completed` | boolean | Yes |
| `items_count` | integer | Yes |
| `checked_count` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Reopen inspection report

```
PATCH /inspection_reports/{id}/reopen
```

Reopen

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `work_order_id` | integer | Yes |
| `quick_finding` | boolean | Yes |
| `published` | boolean | Yes |
| `completed` | boolean | Yes |
| `items_count` | integer | Yes |
| `checked_count` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Reset inspection report

```
PATCH /inspection_reports/{id}/reset
```

Reset

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `work_order_id` | integer | Yes |
| `quick_finding` | boolean | Yes |
| `published` | boolean | Yes |
| `completed` | boolean | Yes |
| `items_count` | integer | Yes |
| `checked_count` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Retry inspection report recording

```
POST /inspection_reports/{id}/retry_recording
```

Retry recording

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `recording_id` | integer | No |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `status` | string | Yes |
| `error_message` | any | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports/<id>.json
```

## Unpublish inspection report

```
PATCH /inspection_reports/{id}/unpublish
```

Unpublish

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `work_order_id` | integer | Yes |
| `quick_finding` | boolean | Yes |
| `published` | boolean | Yes |
| `completed` | boolean | Yes |
| `items_count` | integer | Yes |
| `checked_count` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/inspection_reports/<id>.json
```

---

### InspectionReport schema {#inspectionreport-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `work_order_id` | integer | Yes |
| `quick_finding` | boolean | Yes |
| `published` | boolean | Yes |
| `completed` | boolean | Yes |
| `items_count` | integer | Yes |
| `checked_count` | integer | Yes |
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

---

### MarkAllInspectionReportRequest schema {#markallinspectionreportrequest-schema}

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |
| `group_name` | string | Yes |

---

### ReassignInspectionReportRequest schema {#reassigninspectionreportrequest-schema}

| Field | Type | Required |
|---|---|---|
| `user_id` | integer | Yes |

