# Campaigns

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List campaigns

```
GET /campaigns
```

List all campaigns, paginated via the Link header.

**Response 200** — array of [BroadcastCampaign](#broadcastcampaign-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/campaigns.json
```

## Create campaign

```
POST /campaigns
```

Create a campaign.

**Response 201** — [BroadcastCampaign](#broadcastcampaign-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/campaigns.json
```

## Show campaign

```
GET /campaigns/{id}
```

Show a campaign by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [BroadcastCampaign](#broadcastcampaign-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/campaigns/<id>.json
```

## Duplicate campaign

```
POST /campaigns/{id}/duplicate
```

Duplicate

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
| `sms_body` | string | Yes |
| `filters` | object | Yes |
| `recipient_count` | integer | Yes |
| `sent_count` | integer | Yes |
| `failed_count` | integer | Yes |
| `progress_percentage` | integer | Yes |
| `sent_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |

`filters` — object:

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/campaigns/<id>.json
```

## Send campaign

```
POST /campaigns/{id}/send_campaign
```

Send campaign

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
| `sms_body` | string | Yes |
| `filters` | object | Yes |
| `recipient_count` | integer | Yes |
| `sent_count` | integer | Yes |
| `failed_count` | integer | Yes |
| `progress_percentage` | integer | Yes |
| `sent_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |

`filters` — object:

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/campaigns/<id>.json
```

---

### BroadcastCampaign schema {#broadcastcampaign-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `status` | string | Yes |
| `sms_body` | string | Yes |
| `filters` | object | Yes |
| `recipient_count` | integer | Yes |
| `sent_count` | integer | Yes |
| `failed_count` | integer | Yes |
| `progress_percentage` | integer | Yes |
| `sent_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |

`filters` — object:

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

---

### CreateCampaignRequest schema {#createcampaignrequest-schema}

| Field | Type | Required |
|---|---|---|
| `broadcast_campaign` | object | Yes |

`broadcast_campaign` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `sms_body` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

