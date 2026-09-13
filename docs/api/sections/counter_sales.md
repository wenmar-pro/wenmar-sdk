# Counter Sales

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List counter sales

```
GET /counter_sales
```

List all counter sales, paginated via the Link header.

**Response 200** — array of [CounterSale](#countersale-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/counter_sales.json
```

## Create counter sale

```
POST /counter_sales
```

Create a counter sale.

**Response 201** — [CounterSale](#countersale-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/counter_sales.json
```

## Create counter sales line item

```
POST /counter_sales/{counter_sale_id}/line_items
```

Create a counter sales line item.

| Param | Type | Required |
|---|---|---|
| `counter_sale_id` | integer | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `counter_sale_id` | integer | Yes |
| `description` | string | Yes |
| `quantity` | number | Yes |
| `unit_price_cents` | integer | Yes |
| `unit_cost_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `item_type` | string | Yes |
| `part_type` | string | Yes |
| `source_id` | integer \| null | Yes |
| `source_type` | any | Yes |
| `vendor_id` | integer \| null | Yes |
| `part_number` | string \| null | Yes |
| `brand` | string \| null | Yes |
| `is_taxable` | boolean | Yes |
| `core_charge_cents` | integer | Yes |
| `discount_cents` | integer | Yes |
| `notes` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/counter_sales/{counter_sale_id}/line_items.json
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

## Delete counter sales line item

```
DELETE /counter_sales/{counter_sale_id}/line_items/{id}
```

Delete a counter sales line item by ID.

| Param | Type | Required |
|---|---|---|
| `counter_sale_id` | integer | Yes |
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/counter_sales/<id>.json
```

## Update counter sales line item

```
PATCH /counter_sales/{counter_sale_id}/line_items/{id}
```

Update a counter sales line item by ID.

| Param | Type | Required |
|---|---|---|
| `counter_sale_id` | integer | Yes |
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `counter_sale_id` | integer | Yes |
| `description` | string | Yes |
| `quantity` | number | Yes |
| `unit_price_cents` | integer | Yes |
| `unit_cost_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `item_type` | string | Yes |
| `part_type` | string | Yes |
| `source_id` | integer \| null | Yes |
| `source_type` | any | Yes |
| `vendor_id` | integer \| null | Yes |
| `part_number` | string \| null | Yes |
| `brand` | string \| null | Yes |
| `is_taxable` | boolean | Yes |
| `core_charge_cents` | integer | Yes |
| `discount_cents` | integer | Yes |
| `notes` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/counter_sales/<id>.json
```

## Create counter sales payment

```
POST /counter_sales/{counter_sale_id}/payments
```

Create a counter sales payment.

| Param | Type | Required |
|---|---|---|
| `counter_sale_id` | integer | Yes |

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
| `reopen_url` | string | Yes |

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
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/counter_sales/{counter_sale_id}/payments.json
```

## Show counter sale

```
GET /counter_sales/{id}
```

Show a counter sale by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [CounterSale](#countersale-schema)

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

**Response 200** — [CounterSale](#countersale-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/counter_sales/<id>.json
```

## Update counter sales reopen

```
PATCH /counter_sales/{id}/reopen
```

Update a counter sales reopen by ID.

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
| `reopen_url` | string | Yes |

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

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/counter_sales/<id>.json
```

---

### CounterSale schema {#countersale-schema}

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
| `reopen_url` | string | Yes |

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

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### CreateCounterSalesLineItemRequest schema {#createcountersaleslineitemrequest-schema}

| Field | Type | Required |
|---|---|---|
| `counter_sale_line_items` | array of object | No |
| `counter_sale_line_item` | object | No |

`counter_sale_line_item` — object:
| Field | Type | Required |
|---|---|---|
| `description` | string | Yes |
| `quantity` | integer | Yes |
| `unit_price` | number | Yes |
| `item_type` | string | Yes |

---

### UpdateCounterSalesLineItemRequest schema {#updatecountersaleslineitemrequest-schema}

| Field | Type | Required |
|---|---|---|
| `counter_sale_line_item` | object | Yes |

`counter_sale_line_item` — object:
| Field | Type | Required |
|---|---|---|
| `description` | string | Yes |

---

### CreateCounterSalesPaymentRequest schema {#createcountersalespaymentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `payment` | object | Yes |

`payment` — object:
| Field | Type | Required |
|---|---|---|
| `method` | string | Yes |
| `amount_cents` | number | Yes |

---

### UpdateCounterSaleRequest schema {#updatecountersalerequest-schema}

| Field | Type | Required |
|---|---|---|
| `counter_sale` | object | Yes |

`counter_sale` — object:
| Field | Type | Required |
|---|---|---|
| `notes` | string | Yes |

