# Me

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List me notifications

```
GET /me/notifications
```

List all me notifications, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `email_fallback_enabled` | boolean | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/me/notifications.json
```

## Update me notifications

```
PATCH /me/notifications
```

Update a me notifications by ID.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `email_fallback_enabled` | boolean | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/me/notifications.json
```

## List me preferences

```
GET /me/preferences
```

List all me preferences, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `preferences` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`preferences` — object:
| Field | Type | Required |
|---|---|---|
| `email_notifications` | object | Yes |
| `in_app_notifications` | object | Yes |

`email_notifications` — object:
| Field | Type | Required |
|---|---|---|
| `customer_message_received` | string | Yes |
| `portal_work_authorized` | string | Yes |
| `inspection_completed` | string | Yes |
| `team_mention` | string | Yes |
| `assigned_to_work_order` | string | Yes |
| `technician_update` | string | Yes |
| `online_booking_created` | string | Yes |
| `inspection_acknowledged` | string | Yes |
| `stale_purchase_order` | string | Yes |
| `portal_work_declined` | string | Yes |
| `qbo_sync_failed` | string | Yes |

`in_app_notifications` — object:
| Field | Type | Required |
|---|---|---|
| `customer_message_received` | string | Yes |
| `portal_work_authorized` | string | Yes |
| `inspection_completed` | string | Yes |
| `team_mention` | string | Yes |
| `assigned_to_work_order` | string | Yes |
| `technician_update` | string | Yes |
| `online_booking_created` | string | Yes |
| `inspection_acknowledged` | string | Yes |
| `stale_purchase_order` | string | Yes |
| `portal_work_declined` | string | Yes |
| `qbo_sync_failed` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/me/preferences.json
```

## Update me preferences

```
PATCH /me/preferences
```

Update a me preferences by ID.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `preferences` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`preferences` — object:
| Field | Type | Required |
|---|---|---|
| `email_notifications` | object | Yes |
| `in_app_notifications` | object | Yes |
| `message_sound` | boolean | Yes |

`email_notifications` — object:
| Field | Type | Required |
|---|---|---|
| `customer_message_received` | string | Yes |
| `portal_work_authorized` | string | Yes |
| `inspection_completed` | string | Yes |
| `team_mention` | string | Yes |
| `assigned_to_work_order` | string | Yes |
| `technician_update` | string | Yes |
| `online_booking_created` | string | Yes |
| `inspection_acknowledged` | string | Yes |
| `stale_purchase_order` | string | Yes |
| `portal_work_declined` | string | Yes |
| `qbo_sync_failed` | string | Yes |

`in_app_notifications` — object:
| Field | Type | Required |
|---|---|---|
| `customer_message_received` | string | Yes |
| `portal_work_authorized` | string | Yes |
| `inspection_completed` | string | Yes |
| `team_mention` | string | Yes |
| `assigned_to_work_order` | string | Yes |
| `technician_update` | string | Yes |
| `online_booking_created` | string | Yes |
| `inspection_acknowledged` | string | Yes |
| `stale_purchase_order` | string | Yes |
| `portal_work_declined` | string | Yes |
| `qbo_sync_failed` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/me/preferences.json
```

## List me profile

```
GET /me/profile
```

List all me profile, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `user` | object | Yes |

`user` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `role` | string | Yes |
| `confirmed` | boolean | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/me/profile.json
```

## Update me profile

```
PATCH /me/profile
```

Update a me profile by ID.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `user` | object | Yes |

`user` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `email` | string | Yes |
| `role` | string | Yes |
| `confirmed` | boolean | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/me/profile.json
```

---

### UpdateMeNotificationsRequest schema {#updatemenotificationsrequest-schema}

| Field | Type | Required |
|---|---|---|
| `user` | object | Yes |

`user` — object:
| Field | Type | Required |
|---|---|---|
| `email_fallback_enabled` | boolean | Yes |

---

### UpdateMePreferencesRequest schema {#updatemepreferencesrequest-schema}

| Field | Type | Required |
|---|---|---|
| `user` | object | Yes |

`user` — object:
| Field | Type | Required |
|---|---|---|
| `preferences` | object | Yes |

`preferences` — object:
| Field | Type | Required |
|---|---|---|
| `message_sound` | boolean | Yes |

---

### UpdateMeProfileRequest schema {#updatemeprofilerequest-schema}

| Field | Type | Required |
|---|---|---|
| `user` | object | Yes |

`user` — object:
| Field | Type | Required |
|---|---|---|
| `full_name` | string | Yes |
| `email` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

