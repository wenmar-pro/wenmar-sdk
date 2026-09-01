# Tires

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List tires

```
GET /tires
```

List all tires, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `position` | string | Yes |
| `position_label` | string | Yes |
| `status` | string | Yes |
| `size_raw` | string | Yes |
| `size_width` | integer | Yes |
| `size_aspect_ratio` | integer | Yes |
| `size_rim_diameter` | string | Yes |
| `brand` | string | Yes |
| `model` | string | Yes |
| `load_index` | any | Yes |
| `speed_rating` | any | Yes |
| `dot_serial` | any | Yes |
| `dot_registered` | boolean | Yes |
| `dot_registered_at` | any | Yes |
| `purchase_date` | any | Yes |
| `tread_depth_new_mm` | string | Yes |
| `tread_depth_mm` | any | Yes |
| `source` | string | Yes |
| `notes` | any | Yes |
| `vehicle` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tires.json
```

## Create tire

```
POST /tires
```

Create a tire.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `position` | string | Yes |
| `position_label` | string | Yes |
| `status` | string | Yes |
| `size_raw` | string | Yes |
| `size_width` | integer | Yes |
| `size_aspect_ratio` | integer | Yes |
| `size_rim_diameter` | string | Yes |
| `brand` | string | Yes |
| `model` | any | Yes |
| `load_index` | any | Yes |
| `speed_rating` | any | Yes |
| `dot_serial` | any | Yes |
| `dot_registered` | boolean | Yes |
| `dot_registered_at` | any | Yes |
| `purchase_date` | any | Yes |
| `tread_depth_new_mm` | string | Yes |
| `tread_depth_mm` | any | Yes |
| `source` | string | Yes |
| `notes` | any | Yes |
| `vehicle` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/tires.json
```

## Delete tire

```
DELETE /tires/{id}
```

Delete a tire by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tires/<id>.json
```

## Show tire

```
GET /tires/{id}
```

Show a tire by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `position` | string | Yes |
| `position_label` | string | Yes |
| `status` | string | Yes |
| `size_raw` | string | Yes |
| `size_width` | integer | Yes |
| `size_aspect_ratio` | integer | Yes |
| `size_rim_diameter` | string | Yes |
| `brand` | string | Yes |
| `model` | string | Yes |
| `load_index` | any | Yes |
| `speed_rating` | any | Yes |
| `dot_serial` | any | Yes |
| `dot_registered` | boolean | Yes |
| `dot_registered_at` | any | Yes |
| `purchase_date` | any | Yes |
| `tread_depth_new_mm` | string | Yes |
| `tread_depth_mm` | any | Yes |
| `source` | string | Yes |
| `notes` | any | Yes |
| `vehicle` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/tires/<id>.json
```

## Update tire

```
PATCH /tires/{id}
```

Update a tire by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `position` | string | Yes |
| `position_label` | string | Yes |
| `status` | string | Yes |
| `size_raw` | string | Yes |
| `size_width` | integer | Yes |
| `size_aspect_ratio` | integer | Yes |
| `size_rim_diameter` | string | Yes |
| `brand` | string | Yes |
| `model` | string | Yes |
| `load_index` | any | Yes |
| `speed_rating` | any | Yes |
| `dot_serial` | any | Yes |
| `dot_registered` | boolean | Yes |
| `dot_registered_at` | any | Yes |
| `purchase_date` | any | Yes |
| `tread_depth_new_mm` | string | Yes |
| `tread_depth_mm` | any | Yes |
| `source` | string | Yes |
| `notes` | any | Yes |
| `vehicle` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/tires/<id>.json
```

---

### CreateTireRequest schema {#createtirerequest-schema}

| Field | Type | Required |
|---|---|---|
| `tire` | object | Yes |

`tire` — object:
| Field | Type | Required |
|---|---|---|
| `vehicle_id` | integer | Yes |
| `position` | string | Yes |
| `status` | string | Yes |
| `size_raw` | string | Yes |
| `brand` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateTireRequest schema {#updatetirerequest-schema}

| Field | Type | Required |
|---|---|---|
| `tire` | object | Yes |

`tire` — object:
| Field | Type | Required |
|---|---|---|
| `brand` | string | Yes |

