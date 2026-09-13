# Parts Matrices

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List parts matrices

```
GET /parts_matrices
```

List all parts matrices, paginated via the Link header.

**Response 200** — array of [PartsMatrix](#partsmatrix-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/parts_matrices.json
```

## Create parts matrix

```
POST /parts_matrices
```

Create a parts matrix.

**Response 201** — [PartsMatrix](#partsmatrix-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/parts_matrices.json
```

## Delete parts matrix

```
DELETE /parts_matrices/{id}
```

Delete a parts matrix by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/parts_matrices/<id>.json
```

## Update parts matrix

```
PATCH /parts_matrices/{id}
```

Update a parts matrix by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [PartsMatrix](#partsmatrix-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/parts_matrices/<id>.json
```

---

### PartsMatrix schema {#partsmatrix-schema}

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

---

### CreatePartsMatrixRequest schema {#createpartsmatrixrequest-schema}

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

### UpdatePartsMatrixRequest schema {#updatepartsmatrixrequest-schema}

| Field | Type | Required |
|---|---|---|
| `parts_matrix` | object | Yes |

`parts_matrix` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

