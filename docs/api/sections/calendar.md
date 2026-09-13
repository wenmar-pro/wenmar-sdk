# Calendar

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create calendar blocked time

```
POST /calendar/blocked_times
```

Create a calendar blocked time.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `reason` | string | Yes |
| `recurring` | boolean | Yes |
| `recurrence_rule` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/calendar/blocked_times.json
```

## Delete calendar blocked time

```
DELETE /calendar/blocked_times/{id}
```

Delete a calendar blocked time by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/calendar/blocked_times/<id>.json
```

## Update calendar blocked time

```
PATCH /calendar/blocked_times/{id}
```

Update a calendar blocked time by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `reason` | string | Yes |
| `recurring` | boolean | Yes |
| `recurrence_rule` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/calendar/blocked_times/<id>.json
```

---

### CreateCalendarBlockedTimeRequest schema {#createcalendarblockedtimerequest-schema}

| Field | Type | Required |
|---|---|---|
| `blocked_time` | object | Yes |

`blocked_time` — object:
| Field | Type | Required |
|---|---|---|
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `reason` | string | Yes |

---

### UpdateCalendarBlockedTimeRequest schema {#updatecalendarblockedtimerequest-schema}

| Field | Type | Required |
|---|---|---|
| `blocked_time` | object | Yes |

`blocked_time` — object:
| Field | Type | Required |
|---|---|---|
| `reason` | string | Yes |

