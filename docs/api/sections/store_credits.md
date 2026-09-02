# Store Credits

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create store credits void

```
POST /store_credits/{store_credit_id}/voids
```

Create a store credits void.

| Param | Type | Required |
|---|---|---|
| `store_credit_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/store_credits/{store_credit_id}/voids.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

