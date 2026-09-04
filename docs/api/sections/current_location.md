# Current Location

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List current location

```
GET /current_location
```

List all current location, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `slug` | string | Yes |
| `location_type` | string | Yes |
| `currency` | string | Yes |
| `time_zone` | string | Yes |
| `country` | string | Yes |
| `address` | string | Yes |
| `city` | string | Yes |
| `state` | string | Yes |
| `postal_code` | string | Yes |
| `contact_email` | any | Yes |
| `dock` | array of object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/current_location.json
```

## Update current location

```
PATCH /current_location
```

Update a current location by ID.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `slug` | string | Yes |
| `location_type` | string | Yes |
| `currency` | string | Yes |
| `time_zone` | string | Yes |
| `country` | string | Yes |
| `address` | string | Yes |
| `city` | string | Yes |
| `state` | string | Yes |
| `postal_code` | string | Yes |
| `contact_email` | any | Yes |
| `dock` | array of object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/current_location.json
```

---

### UpdateCurrentLocationRequest schema {#updatecurrentlocationrequest-schema}

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

