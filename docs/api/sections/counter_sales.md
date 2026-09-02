# Counter Sales

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List counter sales

```
GET /counter_sales
```

List all counter sales, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `counter_sale_number` | integer | Yes |
| `status` | string | Yes |
| `walk_in_name` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `subtotal_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `paid_cents` | integer | Yes |
| `remaining_cents` | integer | Yes |
| `paid` | boolean | Yes |
| `currency` | string | Yes |
| `line_items_count` | integer | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/counter_sales.json
```

## Create counter sale

```
POST /counter_sales
```

Create a counter sale.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `counter_sale_number` | integer | Yes |
| `status` | string | Yes |
| `walk_in_name` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `subtotal_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `paid_cents` | integer | Yes |
| `remaining_cents` | integer | Yes |
| `paid` | boolean | Yes |
| `currency` | string | Yes |
| `line_items_count` | integer | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/counter_sales.json
```

## List counter sales line items brands

```
GET /counter_sales/{counter_sale_id}/line_items/brands
```

List all counter sales line items brands, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `counter_sale_id` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/counter_sales/{counter_sale_id}/line_items/brands.json
```

## Show counter sale

```
GET /counter_sales/{id}
```

Show a counter sale by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `counter_sale_number` | integer | Yes |
| `status` | string | Yes |
| `walk_in_name` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `subtotal_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `paid_cents` | integer | Yes |
| `remaining_cents` | integer | Yes |
| `paid` | boolean | Yes |
| `currency` | string | Yes |
| `line_items_count` | integer | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/counter_sales/<id>.json
```

## Update counter sale

```
PATCH /counter_sales/{id}
```

Update a counter sale by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `counter_sale_number` | integer | Yes |
| `status` | string | Yes |
| `walk_in_name` | string \| null | Yes |
| `notes` | string | Yes |
| `subtotal_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `paid_cents` | integer | Yes |
| `remaining_cents` | integer | Yes |
| `paid` | boolean | Yes |
| `currency` | string | Yes |
| `line_items_count` | integer | Yes |
| `processed_by` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/counter_sales/<id>.json
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

### UpdateCounterSaleRequest schema {#updatecountersalerequest-schema}

| Field | Type | Required |
|---|---|---|
| `counter_sale` | object | Yes |

`counter_sale` — object:
| Field | Type | Required |
|---|---|---|
| `notes` | string | Yes |

