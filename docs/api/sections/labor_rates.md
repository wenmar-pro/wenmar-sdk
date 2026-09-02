# Labor Rates

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List labor rates

```
GET /labor_rates
```

List all labor rates, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `rate_cents` | integer | Yes |
| `cost_per_hour_cents` | integer | Yes |
| `is_default` | boolean | Yes |
| `active` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/labor_rates.json
```

## Create labor rate

```
POST /labor_rates
```

Create a labor rate.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `rate_cents` | integer | Yes |
| `cost_per_hour_cents` | integer | Yes |
| `is_default` | boolean | Yes |
| `active` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/labor_rates.json
```

## Update labor rate

```
PATCH /labor_rates/{id}
```

Update a labor rate by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `rate_cents` | integer | Yes |
| `cost_per_hour_cents` | integer | Yes |
| `is_default` | boolean | Yes |
| `active` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/labor_rates/<id>.json
```

---

### CreateLaborRateRequest schema {#createlaborraterequest-schema}

| Field | Type | Required |
|---|---|---|
| `labor_rate` | object | Yes |

`labor_rate` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `rate` | number | Yes |
| `is_default` | boolean | Yes |
| `active` | boolean | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateLaborRateRequest schema {#updatelaborraterequest-schema}

| Field | Type | Required |
|---|---|---|
| `labor_rate` | object | Yes |

`labor_rate` — object:
| Field | Type | Required |
|---|---|---|
| `active` | boolean | Yes |

