# Labor Matrices

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List labor matrices

```
GET /labor_matrices
```

List all labor matrices, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `matrix_type` | string | Yes |
| `active` | boolean | Yes |
| `tiers` | array of object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/labor_matrices.json
```

## Create labor matrice

```
POST /labor_matrices
```

Create a labor matrice.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `matrix_type` | string | Yes |
| `active` | boolean | Yes |
| `tiers` | array of any | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/labor_matrices.json
```

## Delete labor matrice

```
DELETE /labor_matrices/{id}
```

Delete a labor matrice by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/labor_matrices/<id>.json
```

## Update labor matrice

```
PATCH /labor_matrices/{id}
```

Update a labor matrice by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `matrix_type` | string | Yes |
| `active` | boolean | Yes |
| `tiers` | array of object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/labor_matrices/<id>.json
```

---

### CreateLaborMatriceRequest schema {#createlabormatricerequest-schema}

| Field | Type | Required |
|---|---|---|
| `labor_matrix` | object | Yes |

`labor_matrix` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `matrix_type` | string | Yes |
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

### UpdateLaborMatriceRequest schema {#updatelabormatricerequest-schema}

| Field | Type | Required |
|---|---|---|
| `labor_matrix` | object | Yes |

`labor_matrix` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

