# Notifications

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List notifications

```
GET /notifications
```

List all notifications, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `category` | string | No |
| `limit` | integer | No |
| `location_id` | integer | No |
| `read` | boolean | No |
| `since` | string | No |
| `trigger_type` | string | No |

**Response 200** — array of [Notification](#notification-schema)

**Response 401** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/notifications.json
```

## Create notifications bulk mark read

```
POST /notifications/bulk_mark_read
```

Create a notifications bulk mark read.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `ok` | boolean | Yes |
| `affected` | integer | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/notifications/bulk_mark_read.json
```

## Show notification

```
GET /notifications/{id}
```

Show a notification by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Notification](#notification-schema)

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/notifications/<id>.json
```

## Update notification

```
PATCH /notifications/{id}
```

Update a notification by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Notification](#notification-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/notifications/<id>.json
```

---

### Notification schema {#notification-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `trigger_type` | string | Yes |
| `category` | string | Yes |
| `title` | string | Yes |
| `message_body` | string | Yes |
| `read` | boolean | Yes |
| `read_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `location_id` | integer | Yes |
| `work_order_id` | integer \| null | Yes |
| `triggered_by_id` | integer \| null | Yes |
| `action_path` | any | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `metadata` | object | No |

`metadata` — object:
| Field | Type | Required |
|---|---|---|
| `inbound` | boolean | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### CreateNotificationsBulkMarkReadRequest schema {#createnotificationsbulkmarkreadrequest-schema}

| Field | Type | Required |
|---|---|---|
| `notification_ids` | array of integer | No |

---

### UpdateNotificationRequest schema {#updatenotificationrequest-schema}

| Field | Type | Required |
|---|---|---|
| `read` | boolean | Yes |

