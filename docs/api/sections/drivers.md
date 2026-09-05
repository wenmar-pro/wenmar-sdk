# Drivers

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List drivers

```
GET /drivers
```

List all drivers, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `filters[has_open_work_order]` | boolean | No |
| `filters[q]` | string | No |

**Response 200** — array of [Driver](#driver-schema)

**Response 401** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/drivers.json
```

---

### Driver schema {#driver-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `phone` | string \| null | Yes |
| `email` | string \| null | Yes |
| `customer` | object | Yes |
| `work_orders_count` | integer | Yes |
| `work_orders_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

