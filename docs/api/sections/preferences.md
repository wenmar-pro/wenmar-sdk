# Preferences

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List preferences

```
GET /preferences
```

List all preferences, paginated via the Link header.

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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/preferences.json
```

