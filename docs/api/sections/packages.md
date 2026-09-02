# Packages

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List packages

```
GET /packages
```

List all packages, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string | Yes |
| `service_type` | string | Yes |
| `category_id` | integer \| null | Yes |
| `category_name` | string \| null | Yes |
| `estimated_hours` | string | Yes |
| `customer_notes` | string \| null | Yes |
| `show_tech_with_cert` | boolean | Yes |
| `triggers_tire_storage` | boolean | Yes |
| `price_cents` | integer \| null | Yes |
| `active` | boolean | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/packages.json
```

## Create package

```
POST /packages
```

Create a package.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string | Yes |
| `service_type` | string | Yes |
| `category_id` | integer \| null | Yes |
| `category_name` | string \| null | Yes |
| `estimated_hours` | number \| null | Yes |
| `customer_notes` | string \| null | Yes |
| `show_tech_with_cert` | boolean | Yes |
| `triggers_tire_storage` | boolean | Yes |
| `price_cents` | integer \| null | Yes |
| `active` | boolean | Yes |
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
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/packages.json
```

## Update package

```
PATCH /packages/{id}
```

Update a package by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string | Yes |
| `service_type` | string | Yes |
| `category_id` | integer \| null | Yes |
| `category_name` | string \| null | Yes |
| `estimated_hours` | string | Yes |
| `customer_notes` | string \| null | Yes |
| `show_tech_with_cert` | boolean | Yes |
| `triggers_tire_storage` | boolean | Yes |
| `price_cents` | integer \| null | Yes |
| `active` | boolean | Yes |
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
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/packages/<id>.json
```

## Create packages duplicate

```
POST /packages/{id}/duplicate
```

Create a packages duplicate.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string | Yes |
| `service_type` | string | Yes |
| `category_id` | integer \| null | Yes |
| `category_name` | string \| null | Yes |
| `estimated_hours` | string | Yes |
| `customer_notes` | string \| null | Yes |
| `show_tech_with_cert` | boolean | Yes |
| `triggers_tire_storage` | boolean | Yes |
| `price_cents` | integer \| null | Yes |
| `active` | boolean | Yes |
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
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/packages/<id>.json
```

---

### CreatePackageRequest schema {#createpackagerequest-schema}

| Field | Type | Required |
|---|---|---|
| `package` | object | Yes |

`package` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `description` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdatePackageRequest schema {#updatepackagerequest-schema}

| Field | Type | Required |
|---|---|---|
| `package` | object | Yes |

`package` — object:
| Field | Type | Required |
|---|---|---|
| `active` | boolean | Yes |

