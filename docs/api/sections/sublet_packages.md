# Sublet Packages

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List sublet packages

```
GET /sublet_packages
```

List all sublet packages, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string | Yes |
| `active` | boolean | Yes |
| `default_fulfillment_status` | string | Yes |
| `default_payment_status` | string | Yes |
| `default_payment_method` | string \| null | Yes |
| `sublet_package_lines_count` | integer | Yes |
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/sublet_packages.json
```

## Create sublet package

```
POST /sublet_packages
```

Create a sublet package.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string | Yes |
| `active` | boolean | Yes |
| `default_fulfillment_status` | string | Yes |
| `default_payment_status` | string | Yes |
| `default_payment_method` | string \| null | Yes |
| `sublet_package_lines_count` | integer | Yes |
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

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/sublet_packages.json
```

## Delete sublet package

```
DELETE /sublet_packages/{id}
```

Delete a sublet package by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/sublet_packages/<id>.json
```

## Update sublet package

```
PATCH /sublet_packages/{id}
```

Update a sublet package by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string | Yes |
| `active` | boolean | Yes |
| `default_fulfillment_status` | string | Yes |
| `default_payment_status` | string | Yes |
| `default_payment_method` | string \| null | Yes |
| `sublet_package_lines_count` | integer | Yes |
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

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/sublet_packages/<id>.json
```

## Update sublet packages deactivate

```
PATCH /sublet_packages/{id}/deactivate
```

Update a sublet packages deactivate by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/sublet_packages/<id>.json
```

---

### CreateSubletPackageRequest schema {#createsubletpackagerequest-schema}

| Field | Type | Required |
|---|---|---|
| `sublet_package` | object | Yes |

`sublet_package` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `description` | string | Yes |
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

### UpdateSubletPackageRequest schema {#updatesubletpackagerequest-schema}

| Field | Type | Required |
|---|---|---|
| `sublet_package` | object | Yes |

`sublet_package` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

