# Time Entries

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create time entrie

```
POST /time_entries
```

Create a time entrie.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |

**Response 401** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/time_entries.json
```

## Update time entrie

```
PATCH /time_entries/{id}
```

Update a time entrie by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/time_entries/<id>.json
```

---

### CreateTimeEntrieRequest schema {#createtimeentrierequest-schema}

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `work_order_service_id` | integer | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateTimeEntrieRequest schema {#updatetimeentrierequest-schema}

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |

