# Courtesy Cars

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List courtesy cars

```
GET /courtesy_cars
```

List all courtesy cars, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `submodel` | string \| null | Yes |
| `body_style` | string \| null | Yes |
| `engine` | string \| null | Yes |
| `vin` | string | Yes |
| `license_plate` | string | Yes |
| `license_plate_state` | string | Yes |
| `license_plate_country` | string | Yes |
| `drivetrain` | string \| null | Yes |
| `transmission` | string \| null | Yes |
| `color` | string \| null | Yes |
| `vehicle_type` | string | Yes |
| `unit_number` | string \| null | Yes |
| `fleet_identifier` | string \| null | Yes |
| `production_date` | string \| null | Yes |
| `annual_safety_expires_at` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `odometer` | object | Yes |
| `work_orders_count` | integer | Yes |
| `work_orders_url` | string | Yes |
| `customer` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |

`odometer` — object:
| Field | Type | Required |
|---|---|---|
| `reading` | number \| null | Yes |
| `unit` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/courtesy_cars.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

