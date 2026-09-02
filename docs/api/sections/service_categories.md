# Service Categories

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List service categories

```
GET /service_categories
```

List all service categories, paginated via the Link header.

**Response 200** — array of [ServiceCategory](#servicecategory-schema)

**Example**

```json
[
  {
    "id": 1,
    "name": "Maintenance",
    "description": "Routine maintenance services",
    "service_type": "maintenance",
    "icon": "wrench",
    "color": "blue",
    "active": true,
    "position": 1,
    "canonical": true,
    "canonical_key": "maintenance",
    "job_count": 0,
    "url": "https://app.wenmarpro.com/service_categories/1.json",
    "app_url": "https://app.wenmarpro.com/service_categories/1"
  },
  {
    "id": 2,
    "name": "Brakes",
    "description": "Brake system repairs and maintenance",
    "service_type": "repair",
    "icon": "disc",
    "color": "red",
    "active": true,
    "position": 2,
    "canonical": true,
    "canonical_key": "brakes",
    "job_count": 0,
    "url": "https://app.wenmarpro.com/service_categories/2.json",
    "app_url": "https://app.wenmarpro.com/service_categories/2"
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/service_categories.json
```

## Create service category

```
POST /service_categories
```

Create a service category.

**Response 201** — [ServiceCategory](#servicecategory-schema)

**Response 403** — [Error](#error-schema) error envelope

**Example**

```json
{
  "id": 1,
  "name": "Brakes",
  "description": "Brake system repairs and maintenance",
  "service_type": "repair",
  "icon": "disc",
  "color": "red",
  "active": true,
  "position": 2,
  "canonical": true,
  "canonical_key": "brakes",
  "job_count": 0,
  "url": "https://app.wenmarpro.com/service_categories/1.json",
  "app_url": "https://app.wenmarpro.com/service_categories/1"
}
```

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/service_categories.json
```

## Seed defaults service categories

```
POST /service_categories/seed_defaults
```

Seed defaults

**Response 200**

| Field | Type | Required |
|---|---|---|
| `created` | integer | Yes |
| `message` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/service_categories/seed_defaults.json
```

## Delete service category

```
DELETE /service_categories/{id}
```

Delete a service category by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [ServiceCategory](#servicecategory-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/service_categories/<id>.json
```

## Update service category

```
PATCH /service_categories/{id}
```

Update a service category by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [ServiceCategory](#servicecategory-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/service_categories/<id>.json
```

---

### ServiceCategory schema {#servicecategory-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `service_type` | string | Yes |
| `icon` | string \| null | Yes |
| `color` | string \| null | Yes |
| `active` | boolean | Yes |
| `position` | integer | Yes |
| `canonical` | boolean | Yes |
| `canonical_key` | string \| null | Yes |
| `job_count` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

---

### CreateServiceCategoryRequest schema {#createservicecategoryrequest-schema}

| Field | Type | Required |
|---|---|---|
| `service_category` | object | Yes |

`service_category` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `service_type` | string | Yes |
| `icon` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateServiceCategoryRequest schema {#updateservicecategoryrequest-schema}

| Field | Type | Required |
|---|---|---|
| `service_category` | object | Yes |

`service_category` — object:
| Field | Type | Required |
|---|---|---|
| `active` | boolean | No |
| `name` | string | No |
| `position` | integer | No |

