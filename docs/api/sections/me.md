# Me

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List me API tokens

```
GET /me/api_tokens
```

List all me api tokens, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `api_tokens` | array of object | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/me/api_tokens.json
```

## Create me API token

```
POST /me/api_tokens
```

Create a me api token.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `api_token` | object | Yes |

`api_token` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `created_at` | string | Yes |
| `expires_at` | string \| null | Yes |
| `token` | string | Yes |
| `token_preview` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/me/api_tokens.json
```

## Delete me API token

```
DELETE /me/api_tokens/{id}
```

Delete a me api token by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/me/api_tokens/<id>.json
```

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

### CreateMeApiTokenRequest schema {#createmeapitokenrequest-schema}

| Field | Type | Required |
|---|---|---|
| `api_token` | object | Yes |

`api_token` — object:
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

