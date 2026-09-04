# Customer Tags

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List customer tags

```
GET /customer_tags
```

List all customer tags, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `status` | string | No |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customer_tags.json
```

## Create customer tag

```
POST /customer_tags
```

Create a customer tag.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customer_tags.json
```

## Show customer tag

```
GET /customer_tags/{id}
```

Show a customer tag by ID.

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
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customer_tags/<id>.json
```

## Update customer tag

```
PATCH /customer_tags/{id}
```

Update a customer tag by ID.

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
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customer_tags/<id>.json
```

## Update customer tags archive

```
PATCH /customer_tags/{id}/archive
```

Update a customer tags archive by ID.

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
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customer_tags/<id>.json
```

## Update customer tags restore

```
PATCH /customer_tags/{id}/restore
```

Update a customer tags restore by ID.

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
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customer_tags/<id>.json
```

## Update customer tags trash

```
PATCH /customer_tags/{id}/trash
```

Update a customer tags trash by ID.

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
| `bubble_classes` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customer_tags/<id>.json
```

---

### CreateCustomerTagRequest schema {#createcustomertagrequest-schema}

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

### UpdateCustomerTagRequest schema {#updatecustomertagrequest-schema}

| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

