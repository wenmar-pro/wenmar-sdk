# Messages

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List messages

```
GET /messages
```

List all messages, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `conversation_id` | integer | No |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `conversation_id` | integer | Yes |
| `direction` | string | Yes |
| `channel` | string | Yes |
| `status` | string | Yes |
| `body` | string | Yes |
| `recipient_phone` | string | Yes |
| `recipient_email` | any | Yes |
| `work_order_id` | integer \| null | Yes |
| `statement_id` | integer \| null | Yes |
| `appointment_id` | integer \| null | Yes |
| `failure_reason` | string | Yes |
| `attachment_count` | integer | Yes |
| `sender` | object | Yes |
| `sent_at` | string \| null | Yes |
| `delivered_at` | string \| null | Yes |
| `read_at` | string \| null | Yes |
| `failed_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `conversation_url` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`sender` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/messages.json
```

## Show message

```
GET /messages/{id}
```

Show a message by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `conversation_id` | integer | Yes |
| `direction` | string | Yes |
| `channel` | string | Yes |
| `status` | string | Yes |
| `body` | string | Yes |
| `recipient_phone` | string | Yes |
| `recipient_email` | any | Yes |
| `work_order_id` | integer \| null | Yes |
| `statement_id` | integer \| null | Yes |
| `appointment_id` | integer \| null | Yes |
| `failure_reason` | string | Yes |
| `attachment_count` | integer | Yes |
| `sender` | object | Yes |
| `sent_at` | string \| null | Yes |
| `delivered_at` | string \| null | Yes |
| `read_at` | string \| null | Yes |
| `failed_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `conversation_url` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`sender` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/messages/<id>.json
```

## Create messages resend

```
POST /messages/{message_id}/resends
```

Create a messages resend.

| Param | Type | Required |
|---|---|---|
| `message_id` | integer | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `conversation_id` | integer | Yes |
| `direction` | string | Yes |
| `channel` | string | Yes |
| `status` | string | Yes |
| `body` | string | Yes |
| `recipient_phone` | string | Yes |
| `recipient_email` | any | Yes |
| `work_order_id` | integer \| null | Yes |
| `statement_id` | integer \| null | Yes |
| `appointment_id` | integer \| null | Yes |
| `failure_reason` | string \| null | Yes |
| `attachment_count` | integer | Yes |
| `sender` | object | No |
| `sent_at` | string \| null | Yes |
| `delivered_at` | string \| null | Yes |
| `read_at` | string \| null | Yes |
| `failed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `conversation_url` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`sender` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/messages/{message_id}/resends.json
```

