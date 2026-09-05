# Tire Events

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List tire events

```
GET /tire_events
```

List all tire events, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `event_type` | string | Yes |
| `event_type_label` | string | Yes |
| `occurred_at` | string | Yes |
| `notes` | string \| null | Yes |
| `vehicle` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 401** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tire_events.json
```

## Create tire event

```
POST /tire_events
```

Create a tire event.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `event_type` | string | Yes |
| `event_type_label` | string | Yes |
| `occurred_at` | string | Yes |
| `notes` | string | Yes |
| `vehicle` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/tire_events.json
```

## Show tire event

```
GET /tire_events/{id}
```

Show a tire event by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `event_type` | string | Yes |
| `event_type_label` | string | Yes |
| `occurred_at` | string | Yes |
| `notes` | string | Yes |
| `vehicle` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tire_events/<id>.json
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

### CreateTireEventRequest schema {#createtireeventrequest-schema}

| Field | Type | Required |
|---|---|---|
| `tire_event` | object | Yes |

`tire_event` — object:
| Field | Type | Required |
|---|---|---|
| `vehicle_id` | integer | Yes |
| `event_type` | string | Yes |
| `occurred_at` | string | Yes |
| `notes` | string | Yes |

