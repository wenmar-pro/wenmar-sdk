# Customer Tags

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List customer tags

```
GET /customer_tags
```

List all customer tags, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customer_tags.json
```

## Create customer tag

```
POST /customer_tags
```

Create a customer tag.

**Request body**:

| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | No |

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customer_tags.json
```

## Delete customer tag

```
DELETE /customer_tags/{id}
```

Delete a customer tag by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customer_tags/<id>.json
```

## Update customer tag

```
PATCH /customer_tags/{id}
```

Update a customer tag by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Request body**:

| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customer_tags/<id>.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

