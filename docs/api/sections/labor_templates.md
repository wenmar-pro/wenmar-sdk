# Labor Templates

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List labor templates

```
GET /labor_templates
```

List all labor templates, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `description` | string | Yes |
| `default_hours` | string | Yes |
| `usage_count` | integer | Yes |
| `last_used_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/labor_templates.json
```

## Delete labor template

```
DELETE /labor_templates/{id}
```

Delete a labor template by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/labor_templates/<id>.json
```

## Update labor template

```
PATCH /labor_templates/{id}
```

Update a labor template by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `description` | string | Yes |
| `default_hours` | string | Yes |
| `usage_count` | integer | Yes |
| `last_used_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/labor_templates/<id>.json
```

---

### UpdateLaborTemplateRequest schema {#updatelabortemplaterequest-schema}

| Field | Type | Required |
|---|---|---|
| `labor_template` | object | Yes |

`labor_template` — object:
| Field | Type | Required |
|---|---|---|
| `default_hours` | number | Yes |

