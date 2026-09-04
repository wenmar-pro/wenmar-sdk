# Vendors

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List vendors

```
GET /vendors
```

List all vendors, paginated via the Link header.

**Response 200** — array of [Vendor](#vendor-schema)

**Response 401** — [Error](#error-schema) error envelope

**Response 403** — [Error](#error-schema) error envelope

**Example**

```json
[
  {
    "id": 1,
    "name": "Acme Auto Parts",
    "vendor_type": "parts",
    "payment_terms": "net_30",
    "active": true,
    "phone": "+1 416 555 0100",
    "email": "orders@acmeparts.com",
    "website": "https://acmeparts.com",
    "account_number": "ACC-1001",
    "notes": "Preferred brake parts supplier",
    "quick_order": true,
    "order_url_template": null,
    "catalog_url_template": null,
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/vendors/1.json",
    "app_url": "https://app.wenmarpro.com/vendors/1"
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vendors.json
```

## Create vendor

```
POST /vendors
```

Create a vendor.

**Response 201** — [Vendor](#vendor-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vendors.json
```

## Show vendor

```
GET /vendors/{id}
```

Show a vendor by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Vendor](#vendor-schema)

**Response 404** — [Error](#error-schema) error envelope

**Example**

```json
{
  "id": 1,
  "name": "Acme Auto Parts",
  "vendor_type": "parts",
  "payment_terms": "net_30",
  "active": true,
  "phone": "+1 416 555 0100",
  "email": "orders@acmeparts.com",
  "website": "https://acmeparts.com",
  "account_number": "ACC-1001",
  "notes": "Preferred brake parts supplier",
  "quick_order": true,
  "order_url_template": null,
  "catalog_url_template": null,
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/vendors/1.json",
  "app_url": "https://app.wenmarpro.com/vendors/1"
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vendors/<id>.json
```

## Update vendor

```
PATCH /vendors/{id}
```

Update a vendor by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Vendor](#vendor-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vendors/<id>.json
```

## Archive vendor

```
PATCH /vendors/{id}/archive
```

Archive

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `vendor_type` | string | Yes |
| `payment_terms` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `phone` | string | Yes |
| `email` | string | Yes |
| `website` | string | Yes |
| `account_number` | string | Yes |
| `notes` | string | Yes |
| `quick_order` | boolean | Yes |
| `order_url_template` | string \| null | Yes |
| `catalog_url_template` | string \| null | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vendors/<id>.json
```

## Restore vendor

```
PATCH /vendors/{id}/restore
```

Restore

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `vendor_type` | string | Yes |
| `payment_terms` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `phone` | string | Yes |
| `email` | string | Yes |
| `website` | string | Yes |
| `account_number` | string | Yes |
| `notes` | string | Yes |
| `quick_order` | boolean | Yes |
| `order_url_template` | string \| null | Yes |
| `catalog_url_template` | string \| null | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vendors/<id>.json
```

## Trash vendor

```
PATCH /vendors/{id}/trash
```

Trash

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `vendor_type` | string | Yes |
| `payment_terms` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string | Yes |
| `phone` | string | Yes |
| `email` | string | Yes |
| `website` | string | Yes |
| `account_number` | string | Yes |
| `notes` | string | Yes |
| `quick_order` | boolean | Yes |
| `order_url_template` | string \| null | Yes |
| `catalog_url_template` | string \| null | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vendors/<id>.json
```

## List vendors purchase orders

```
GET /vendors/{vendor_id}/purchase_orders
```

List all vendors purchase orders, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `vendor_id` | integer | Yes |

**Response 200** — array of [PurchaseOrder](#purchaseorder-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vendors/{vendor_id}/purchase_orders.json
```

---

### Vendor schema {#vendor-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `vendor_type` | string | Yes |
| `payment_terms` | string | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `phone` | string | Yes |
| `email` | string | Yes |
| `website` | string | Yes |
| `account_number` | string | Yes |
| `notes` | string | Yes |
| `quick_order` | boolean | Yes |
| `order_url_template` | string \| null | Yes |
| `catalog_url_template` | string \| null | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

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

### CreateVendorRequest schema {#createvendorrequest-schema}

| Field | Type | Required |
|---|---|---|
| `vendor` | object | Yes |

`vendor` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `vendor_type` | string | Yes |
| `payment_terms` | string | Yes |

---

### UpdateVendorRequest schema {#updatevendorrequest-schema}

| Field | Type | Required |
|---|---|---|
| `vendor` | object | Yes |

`vendor` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### PurchaseOrder schema {#purchaseorder-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `po_number` | integer | Yes |
| `status` | string | Yes |
| `order_method` | string | Yes |
| `payment_method` | string | Yes |
| `fulfillment_method` | string | Yes |
| `tracking_number` | any | Yes |
| `vendor_invoice_number` | any | Yes |
| `vendor_invoice_received_at` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `freight_cost_cents` | integer | Yes |
| `freight_cost_currency` | string | Yes |
| `subtotal_cents` | string | Yes |
| `total_cents` | string | Yes |
| `core_charges_cents` | integer | Yes |
| `line_items_count` | integer | Yes |
| `ordered_at` | string | Yes |
| `received_at` | string \| null | Yes |
| `payment_due_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `vendor` | object | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |
| `line_items` | array of object | Yes |

`vendor` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

