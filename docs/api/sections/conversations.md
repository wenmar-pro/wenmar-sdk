# Conversations

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List conversations

```
GET /conversations
```

List all conversations, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `status` | string | Yes |
| `reply_state` | string | Yes |
| `channel` | string | Yes |
| `customer` | object | Yes |
| `unread_count` | integer | Yes |
| `messages_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 302** — no content.

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/conversations.json
```

## Create conversations bulk mark read

```
POST /conversations/bulk_mark_read
```

Create a conversations bulk mark read.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `ok` | boolean | Yes |
| `affected` | integer | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/conversations/bulk_mark_read.json
```

## Create conversations customer link

```
POST /conversations/{conversation_id}/customer_links
```

Create a conversations customer link.

| Param | Type | Required |
|---|---|---|
| `conversation_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `status` | string | Yes |
| `reply_state` | string | Yes |
| `channel` | string | Yes |
| `customer` | object | Yes |
| `unread_count` | integer | Yes |
| `messages_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/conversations/{conversation_id}/customer_links.json
```

## Create conversations ignore

```
POST /conversations/{conversation_id}/ignores
```

Create a conversations ignore.

| Param | Type | Required |
|---|---|---|
| `conversation_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `status` | string | Yes |
| `reply_state` | string | Yes |
| `channel` | string | Yes |
| `customer` | object | Yes |
| `unread_count` | integer | Yes |
| `messages_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/conversations/{conversation_id}/ignores.json
```

## List conversations messages

```
GET /conversations/{conversation_id}/messages
```

List all conversations messages, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `conversation_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `conversation_id` | integer | Yes |
| `direction` | string | Yes |
| `channel` | string | Yes |
| `status` | string | Yes |
| `body` | string | Yes |
| `work_order_id` | integer \| null | Yes |
| `statement_id` | integer \| null | Yes |
| `appointment_id` | integer \| null | Yes |
| `read_at` | string \| null | Yes |
| `created_at` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/conversations/{conversation_id}/messages.json
```

## Create conversations message

```
POST /conversations/{conversation_id}/messages
```

Create a conversations message.

| Param | Type | Required |
|---|---|---|
| `conversation_id` | integer | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `conversation_id` | integer | Yes |
| `direction` | string | Yes |
| `channel` | string | Yes |
| `status` | string | Yes |
| `body` | string | Yes |
| `sender` | object | Yes |
| `work_order_id` | integer \| null | Yes |
| `statement_id` | integer \| null | Yes |
| `appointment_id` | integer \| null | Yes |
| `read_at` | string \| null | Yes |
| `created_at` | string | Yes |

`sender` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/conversations/{conversation_id}/messages.json
```

## Create conversations messages resend

```
POST /conversations/{conversation_id}/messages/{id}/resends
```

Create a conversations messages resend.

| Param | Type | Required |
|---|---|---|
| `conversation_id` | integer | Yes |
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
| `work_order_id` | integer \| null | Yes |
| `statement_id` | integer \| null | Yes |
| `appointment_id` | integer \| null | Yes |
| `read_at` | string \| null | Yes |
| `created_at` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/conversations/<id>.json
```

## Show conversation

```
GET /conversations/{id}
```

Show a conversation by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `status` | string | Yes |
| `reply_state` | string | Yes |
| `channel` | string | Yes |
| `customer` | object | Yes |
| `unread_count` | integer | Yes |
| `messages_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/conversations/<id>.json
```

## Update conversation

```
PATCH /conversations/{id}
```

Update a conversation by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `status` | string | Yes |
| `reply_state` | string | Yes |
| `channel` | string | Yes |
| `customer` | object | Yes |
| `unread_count` | integer | Yes |
| `messages_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/conversations/<id>.json
```

---

### CreateConversationsCustomerLinkRequest schema {#createconversationscustomerlinkrequest-schema}

| Field | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |

---

### CreateConversationsMessageRequest schema {#createconversationsmessagerequest-schema}

| Field | Type | Required |
|---|---|---|
| `message` | object | Yes |

`message` — object:
| Field | Type | Required |
|---|---|---|
| `body` | string | Yes |
| `channel` | string | Yes |

---

### UpdateConversationRequest schema {#updateconversationrequest-schema}

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |

