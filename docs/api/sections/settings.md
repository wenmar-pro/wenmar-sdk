# Settings

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List tags

```
GET /settings/tags
```

List all tags, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `customer_tags` | array of object | Yes |
| `vehicle_tags` | array of object | Yes |

**Response 404** — no content.

**Example**

```json
{
  "customer_tags": [
    {
      "id": 1,
      "account_id": 1,
      "name": "VIP",
      "color": "blue",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "customers_count": 5,
      "metadata": {},
      "updated_at": "2026-08-27T12:00:00.000-04:00"
    },
    {
      "id": 2,
      "account_id": 1,
      "name": "FLEET",
      "color": "green",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "customers_count": 2,
      "metadata": {},
      "updated_at": "2026-08-27T12:00:00.000-04:00"
    }
  ],
  "vehicle_tags": [
    {
      "id": 3,
      "account_id": 1,
      "name": "FLEET",
      "color": "green",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "updated_at": "2026-08-27T12:00:00.000-04:00"
    },
    {
      "id": 4,
      "account_id": 1,
      "name": "LEASED",
      "color": "amber",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "updated_at": "2026-08-27T12:00:00.000-04:00"
    }
  ]
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/tags.json
```

## Update tags

```
PATCH /settings/tags
```

Update a tags by ID.

**Request body**:

| Field | Type | Required |
|---|---|---|
| `vehicle_tags` | array of object | No |
| `customer_tags` | array of object | No |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `customer_tags` | array of object | Yes |
| `vehicle_tags` | array of any | Yes |

**Response 403** — no content.

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/settings/tags.json
```

