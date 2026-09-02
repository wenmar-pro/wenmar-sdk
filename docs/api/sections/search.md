# Search

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List search

```
GET /search
```

List all search, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `q` | string | No |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `html` | string | Yes |
| `query` | string | Yes |
| `announcement` | string | Yes |
| `total_count` | integer | Yes |
| `groups` | array of object | Yes |

**Response 401** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/search.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

