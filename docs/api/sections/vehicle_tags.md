# Vehicle Tags

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List vehicle tags

```
GET /vehicle_tags
```

List all vehicle tags, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicle_tags.json
```

## Create vehicle tag

```
POST /vehicle_tags
```

Create a vehicle tag.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vehicle_tags.json
```

## Delete vehicle tag

```
DELETE /vehicle_tags/{id}
```

Delete a vehicle tag by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicle_tags/<id>.json
```

## Update vehicle tag

```
PATCH /vehicle_tags/{id}
```

Update a vehicle tag by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `color` | string | Yes |
| `color_hex` | string | Yes |
| `color_class` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vehicle_tags/<id>.json
```

---

### CreateVehicleTagRequest schema {#createvehicletagrequest-schema}

| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### UpdateVehicleTagRequest schema {#updatevehicletagrequest-schema}

| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

