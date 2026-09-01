# Fleets

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List fleets

```
GET /fleets
```

List all fleets, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `company_name` | string | Yes |
| `first_name` | string | Yes |
| `last_name` | string | Yes |
| `fleet_identifier` | string | Yes |
| `marketing_opt_in` | boolean | Yes |
| `tax_exempt` | boolean | Yes |
| `vehicles_count` | integer | Yes |
| `emails_count` | integer | Yes |
| `phones_count` | integer | Yes |
| `vehicles_url` | string | Yes |
| `work_orders_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/fleets.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

