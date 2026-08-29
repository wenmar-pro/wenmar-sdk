# Vehicles

## List vehicles

```
GET /vehicles
```

Paginated via the `Link` header. Returns a bare array.

| Param | Type | Required |
|---|---|---|
| `page` | integer | No |

**Response 200** — bare array of vehicle objects.

## Create vehicle

```
POST /vehicles
```

**Request body** — `vehicle` object:

| Field | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |

**Response 201** — bare vehicle object (same shape as show).

## Show vehicle

```
GET /vehicles/{id}
```

**Response 200** — bare vehicle object:

```json
{
  "type": "Vehicle",
  "id": 1,
  "make": "Toyota",
  "model": "Camry",
  "year": 2020,
  "submodel": null,
  "body_style": null,
  "engine": null,
  "vin": "ABC123",
  "license_plate": "JANE1",
  "license_plate_state": "ON",
  "license_plate_country": "CA",
  "drivetrain": null,
  "transmission": null,
  "color": null,
  "vehicle_type": "customer",
  "unit_number": null,
  "fleet_identifier": null,
  "production_date": null,
  "annual_safety_expires_at": null,
  "notes": null,
  "odometer": { "reading": 50000, "unit": "km" },
  "work_orders_count": 1,
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?vehicle_id=1",
  "customer": { "id": 1, "full_name": "Jane Doe", "url": "https://app.wenmarpro.com/customers/1.json" },
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/vehicles/1.json",
  "last_serviced_at": null,
  "lifetime_revenue_cents": 50000,
  "open_work_orders_count": 0,
  "appointments_count": 0
}
```

## Update vehicle

```
PATCH /vehicles/{id}
```

**Request body** — `vehicle` object (any subset of fields).

**Response 200** — bare vehicle object.

## Delete vehicle

```
DELETE /vehicles/{id}
```

**Response 204** — no content.

## Decode VIN

```
GET /vehicles/vin_decode?vin={vin}
```

| Param | Type | Required |
|---|---|---|
| `vin` | string | Yes |

**Response 200** — decoded vehicle data:

```json
{ "make": "Honda", "model": "Civic", "year": 2020, "vin": "1HGCM82633A004352" }
```

## Check duplicates

```
GET /vehicles/check_duplicate?vin={vin}
```

| Param | Type | Required |
|---|---|---|
| `vin` | string | Yes |

**Response 200**:

```json
{ "matches": [ { "id": 1, "display_name": "Toyota Camry", "url": "/vehicles/1", "reasons": ["vin"] } ] }
```
