# Packages

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List packages

```
GET /packages
```

List all packages, paginated via the Link header.

**Response 200** — array of [Package](#package-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/packages.json
```

## Create package

```
POST /packages
```

Create a package.

**Response 201** — [Package](#package-schema)

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

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/packages/<id>.json
```

## Archive package

```
PATCH /packages/{id}/archive
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
| `description` | string | Yes |
| `service_type` | string | Yes |
| `category_id` | integer \| null | Yes |
| `category_name` | string \| null | Yes |
| `estimated_hours` | string | Yes |
| `customer_notes` | string \| null | Yes |
| `show_tech_with_cert` | boolean | Yes |
| `triggers_tire_storage` | boolean | Yes |
| `price_cents` | integer \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
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

**Response 201** — [Package](#package-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/packages/<id>.json
```

## Restore package

```
PATCH /packages/{id}/restore
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
| `description` | string | Yes |
| `service_type` | string | Yes |
| `category_id` | integer \| null | Yes |
| `category_name` | string \| null | Yes |
| `estimated_hours` | string | Yes |
| `customer_notes` | string \| null | Yes |
| `show_tech_with_cert` | boolean | Yes |
| `triggers_tire_storage` | boolean | Yes |
| `price_cents` | integer \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
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
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/packages/<id>.json
```

## Trash package

```
PATCH /packages/{id}/trash
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
| `description` | string | Yes |
| `service_type` | string | Yes |
| `category_id` | integer \| null | Yes |
| `category_name` | string \| null | Yes |
| `estimated_hours` | string | Yes |
| `customer_notes` | string \| null | Yes |
| `show_tech_with_cert` | boolean | Yes |
| `triggers_tire_storage` | boolean | Yes |
| `price_cents` | integer \| null | Yes |
| `status` | string | Yes |
| `trashed_at` | string | Yes |
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
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/packages/<id>.json
```

---

### Package schema {#package-schema}

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
| `status` | string | Yes |
| `trashed_at` | string \| null | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

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

