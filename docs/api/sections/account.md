# Account

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List account

```
GET /account
```

List all account, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `slug` | string | Yes |
| `locations` | array of object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 401** — [Error](#error-schema) error envelope

**Response 403** — [Error](#error-schema) error envelope

**Example**

```json
{
  "id": 1,
  "name": "Main Shop",
  "slug": "main-shop",
  "locations": [
    {
      "id": 1,
      "name": "Bay 1",
      "url": "https://app.wenmarpro.com/locations/1.json",
      "app_url": "https://app.wenmarpro.com/locations/1"
    }
  ],
  "url": "https://app.wenmarpro.com/account.json",
  "app_url": "https://app.wenmarpro.com/account"
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

