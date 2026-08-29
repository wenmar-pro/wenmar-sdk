# Vendors

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List vendors

```
GET /vendors
```

List all vendors, paginated via the Link header.

**Response 200** — array of [Vendor](#vendor-schema)

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

---

### Vendor schema {#vendor-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `vendor_type` | string | Yes |
| `payment_terms` | string | Yes |
| `active` | boolean | Yes |
| `phone` | string | Yes |
| `email` | string | Yes |
| `website` | string | Yes |
| `account_number` | string | Yes |
| `notes` | string | Yes |
| `quick_order` | boolean | Yes |
| `order_url_template` | any | Yes |
| `catalog_url_template` | any | Yes |
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

