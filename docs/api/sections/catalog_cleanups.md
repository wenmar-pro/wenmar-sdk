# Catalog Cleanups

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create catalog cleanups application

```
POST /catalog_cleanups/{catalog_cleanup_id}/applications
```

Create a catalog cleanups application.

| Param | Type | Required |
|---|---|---|
| `catalog_cleanup_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |
| `applied_count` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/catalog_cleanups/{catalog_cleanup_id}/applications.json
```

---

### CreateCatalogCleanupsApplicationRequest schema {#createcatalogcleanupsapplicationrequest-schema}

| Field | Type | Required |
|---|---|---|
| `category` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

