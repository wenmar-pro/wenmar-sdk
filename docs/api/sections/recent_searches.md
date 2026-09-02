# Recent Searches

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List recent searches

```
GET /recent_searches
```

List all recent searches, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `recents` | array of object | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/recent_searches.json
```

## Create recent searche

```
POST /recent_searches
```

Create a recent searche.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `recents` | array of object | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/recent_searches.json
```

## Delete recent searches clear

```
DELETE /recent_searches/clear
```

Delete a recent searches clear by ID.

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/recent_searches/clear.json
```

## Delete recent searche

```
DELETE /recent_searches/{id}
```

Delete a recent searche by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `recents` | array of any | Yes |

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/recent_searches/<id>.json
```

---

### CreateRecentSearcheRequest schema {#createrecentsearcherequest-schema}

| Field | Type | Required |
|---|---|---|
| `recent_search` | object | Yes |

`recent_search` — object:
| Field | Type | Required |
|---|---|---|
| `query` | string | Yes |
| `result_type` | string | Yes |
| `result_id` | string | Yes |
| `label` | string | Yes |

