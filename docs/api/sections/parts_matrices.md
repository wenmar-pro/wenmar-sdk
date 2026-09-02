# Parts Matrices

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List parts matrices

```
GET /parts_matrices
```

List all parts matrices, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `is_default` | boolean | Yes |
| `active` | boolean | Yes |
| `default_multiplier` | string | Yes |
| `max_markup_cents` | integer \| null | Yes |
| `tiers` | array of object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/parts_matrices.json
```

## Create parts matrice

```
POST /parts_matrices
```

Create a parts matrice.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `is_default` | boolean | Yes |
| `active` | boolean | Yes |
| `default_multiplier` | string | Yes |
| `max_markup_cents` | integer \| null | Yes |
| `tiers` | array of any | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/parts_matrices.json
```

## Delete parts matrice

```
DELETE /parts_matrices/{id}
```

Delete a parts matrice by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/parts_matrices/<id>.json
```

## Update parts matrice

```
PATCH /parts_matrices/{id}
```

Update a parts matrice by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `is_default` | boolean | Yes |
| `active` | boolean | Yes |
| `default_multiplier` | string | Yes |
| `max_markup_cents` | integer \| null | Yes |
| `tiers` | array of object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/parts_matrices/<id>.json
```

---

### CreatePartsMatriceRequest schema {#createpartsmatricerequest-schema}

| Field | Type | Required |
|---|---|---|
| `parts_matrix` | object | Yes |

`parts_matrix` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
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

### UpdatePartsMatriceRequest schema {#updatepartsmatricerequest-schema}

| Field | Type | Required |
|---|---|---|
| `parts_matrix` | object | Yes |

`parts_matrix` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

