# Lead Sources

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List lead sources

```
GET /lead_sources
```

List all lead sources, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `category` | string | Yes |
| `system_default` | boolean | Yes |
| `active` | boolean | Yes |
| `position` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/lead_sources.json
```

## Create lead source

```
POST /lead_sources
```

Create a lead source.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `category` | string | Yes |
| `system_default` | boolean | Yes |
| `active` | boolean | Yes |
| `position` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/lead_sources.json
```

## Create lead sources seed default

```
POST /lead_sources/seed_defaults
```

Create a lead sources seed default.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `category` | string | Yes |
| `system_default` | boolean | Yes |
| `active` | boolean | Yes |
| `position` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/lead_sources/seed_defaults.json
```

## Delete lead source

```
DELETE /lead_sources/{id}
```

Delete a lead source by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/lead_sources/<id>.json
```

## Update lead source

```
PATCH /lead_sources/{id}
```

Update a lead source by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `category` | string | Yes |
| `system_default` | boolean | Yes |
| `active` | boolean | Yes |
| `position` | integer | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/lead_sources/<id>.json
```

---

### CreateLeadSourceRequest schema {#createleadsourcerequest-schema}

| Field | Type | Required |
|---|---|---|
| `lead_source` | object | Yes |

`lead_source` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `category` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateLeadSourceRequest schema {#updateleadsourcerequest-schema}

| Field | Type | Required |
|---|---|---|
| `lead_source` | object | Yes |

`lead_source` — object:
| Field | Type | Required |
|---|---|---|
| `active` | boolean | Yes |

