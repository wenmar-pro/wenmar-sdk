# Labor Matrices

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List labor matrices

```
GET /labor_matrices
```

List all labor matrices, paginated via the Link header.

**Response 200** — array of [LaborMatrix](#labormatrix-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/labor_matrices.json
```

## Create labor matrix

```
POST /labor_matrices
```

Create a labor matrix.

**Response 201** — [LaborMatrix](#labormatrix-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/labor_matrices.json
```

## Delete labor matrix

```
DELETE /labor_matrices/{id}
```

Delete a labor matrix by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/labor_matrices/<id>.json
```

## Update labor matrix

```
PATCH /labor_matrices/{id}
```

Update a labor matrix by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [LaborMatrix](#labormatrix-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/labor_matrices/<id>.json
```

---

### LaborMatrix schema {#labormatrix-schema}

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

---

### CreateLaborMatrixRequest schema {#createlabormatrixrequest-schema}

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

### UpdateLaborMatrixRequest schema {#updatelabormatrixrequest-schema}

| Field | Type | Required |
|---|---|---|
| `labor_matrix` | object | Yes |

`labor_matrix` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

