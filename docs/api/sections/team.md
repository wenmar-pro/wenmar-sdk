# Team

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List team

```
GET /team/members
```

List all team, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `email` | any | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `home_location` | object | Yes |
| `locations` | array of object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |

`home_location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/team/members.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

