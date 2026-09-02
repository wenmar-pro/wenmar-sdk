# Core Tax Rules

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List core tax rules

```
GET /core_tax_rules
```

List all core tax rules, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `province_code` | string | Yes |
| `tax_core_charge` | boolean | Yes |
| `tax_core_credit` | boolean | Yes |
| `notes` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/core_tax_rules.json
```

## Update core tax rule

```
PATCH /core_tax_rules/{id}
```

Update a core tax rule by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `province_code` | string | Yes |
| `tax_core_charge` | boolean | Yes |
| `tax_core_credit` | boolean | Yes |
| `notes` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/core_tax_rules/<id>.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateCoreTaxRuleRequest schema {#updatecoretaxrulerequest-schema}

| Field | Type | Required |
|---|---|---|
| `core_tax_rule` | object | Yes |

`core_tax_rule` — object:
| Field | Type | Required |
|---|---|---|
| `tax_core_charge` | boolean | Yes |

