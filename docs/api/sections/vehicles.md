# Vehicles

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List vehicles

```
GET /vehicles
```

List all vehicles, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | No |
| `page` | integer | No |

**Response 200** — array of [Vehicle](#vehicle-schema)

**Response 403** — [Error](#error-schema) error envelope

**Example**

```json
[
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
    "odometer": {
      "reading": 50000,
      "unit": "km"
    },
    "work_orders_count": 1,
    "work_orders_url": "https://app.wenmarpro.com/work_orders.json?vehicle_id=1",
    "customer": {
      "id": 1,
      "full_name": "Jane Doe",
      "url": "https://app.wenmarpro.com/customers/1.json"
    },
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "app_url": "https://app.wenmarpro.com/vehicles/1",
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "last_serviced_at": null,
    "lifetime_revenue_cents": 0,
    "open_work_orders_count": 0,
    "appointments_count": 0
  },
  {
    "type": "Vehicle",
    "id": 2,
    "make": "Honda",
    "model": "Civic",
    "year": 2018,
    "submodel": null,
    "body_style": null,
    "engine": null,
    "vin": "XYZ789",
    "license_plate": "JOHN1",
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
    "odometer": {
      "reading": 72000,
      "unit": "km"
    },
    "work_orders_count": 0,
    "work_orders_url": "https://app.wenmarpro.com/work_orders.json?vehicle_id=2",
    "customer": {
      "id": 2,
      "full_name": "John Smith",
      "url": "https://app.wenmarpro.com/customers/2.json"
    },
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/vehicles/2.json",
    "app_url": "https://app.wenmarpro.com/vehicles/2",
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "last_serviced_at": null,
    "lifetime_revenue_cents": 0,
    "open_work_orders_count": 0,
    "appointments_count": 0
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicles.json
```

## Create vehicle

```
POST /vehicles
```

Create a vehicle.

**Request body** — wrapper key `vehicle`:

| Field | Type | Required |
|---|---|---|
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `customer_id` | integer | Yes |
| `vin` | string | No |
| `submodel` | string | No |
| `body_style` | string | No |
| `engine` | string | No |
| `transmission` | string | No |
| `drivetrain` | string | No |
| `color` | string | No |
| `license_plate` | string | No |
| `license_plate_state` | string | No |
| `odometer_reading` | integer | No |
| `odometer_unit` | string | No |
| `unit_number` | string | No |
| `fleet_identifier` | string | No |
| `notes` | string | No |
| `production_date` | string | No |
| `vehicle_tag_ids` | array of any | No |

**Response 201** — [Vehicle](#vehicle-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

**Example**

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
  "odometer": {
    "reading": 50000,
    "unit": "km"
  },
  "work_orders_count": 1,
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?vehicle_id=1",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/vehicles/1.json",
  "app_url": "https://app.wenmarpro.com/vehicles/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "last_serviced_at": null,
  "lifetime_revenue_cents": 50000,
  "open_work_orders_count": 0,
  "appointments_count": 0
}
```

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vehicles.json
```

## Check vehicle duplicate

```
GET /vehicles/check_duplicate
```

Check duplicate

| Param | Type | Required |
|---|---|---|
| `vin` | string | No |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `matches` | array of object | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicles/check_duplicate.json
```

## Lookup vehicle

```
GET /vehicles/lookup
```

Lookup

| Param | Type | Required |
|---|---|---|
| `query` | string | No |

**Response 200** — array of [Vehicle](#vehicle-schema)

**Example**

```json
[
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
    "odometer": {
      "reading": 50000,
      "unit": "km"
    },
    "work_orders_count": 1,
    "work_orders_url": "https://app.wenmarpro.com/vehicles/1/work_orders.json",
    "customer": {
      "id": 1,
      "full_name": "Jane Doe",
      "url": "https://app.wenmarpro.com/customers/1.json"
    },
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "app_url": "https://app.wenmarpro.com/vehicles/1",
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "last_serviced_at": null,
    "lifetime_revenue_cents": 50000,
    "open_work_orders_count": 0,
    "appointments_count": 0
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicles/lookup.json
```

## Prefill vehicle

```
GET /vehicles/prefill
```

Prefill

| Param | Type | Required |
|---|---|---|
| `make` | string | No |
| `model` | string | No |
| `vin` | string | No |
| `year` | integer | No |

**Response 200** — [Vehicle](#vehicle-schema)

**Response 404** — [Error](#error-schema) error envelope

**Example**

```json
{
  "type": "Vehicle",
  "id": 1,
  "make": "Toyota",
  "model": "Camry",
  "year": 2020,
  "submodel": "XSE",
  "body_style": "sedan",
  "engine": "2.5L 4-cyl",
  "vin": "ABC123",
  "license_plate": "JANE1",
  "license_plate_state": "ON",
  "license_plate_country": "CA",
  "drivetrain": "fwd",
  "transmission": "automatic",
  "color": "silver",
  "vehicle_type": "customer",
  "unit_number": null,
  "fleet_identifier": null,
  "production_date": "2019-11-01",
  "annual_safety_expires_at": null,
  "notes": null,
  "odometer": {
    "reading": 45000,
    "unit": "km"
  },
  "work_orders_count": 0,
  "work_orders_url": "https://app.wenmarpro.com/vehicles/1/work_orders.json",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/vehicles/1.json",
  "app_url": "https://app.wenmarpro.com/vehicles/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "last_serviced_at": null,
  "lifetime_revenue_cents": 0,
  "open_work_orders_count": 0,
  "appointments_count": 0
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicles/prefill.json
```

## Decode VIN

```
GET /vehicles/vin_decode
```

VIN decode

| Param | Type | Required |
|---|---|---|
| `vin` | string | No |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `make` | string | Yes |
| `model` | string | Yes |
| `vin` | string | Yes |

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicles/vin_decode.json
```

## Delete vehicle

```
DELETE /vehicles/{id}
```

Delete a vehicle by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 202** — [Vehicle](#vehicle-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicles/<id>.json
```

## Show vehicle

```
GET /vehicles/{id}
```

Show a vehicle by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Vehicle](#vehicle-schema)

**Response 404** — [Error](#error-schema) error envelope

**Example**

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
  "odometer": {
    "reading": 50000,
    "unit": "km"
  },
  "work_orders_count": 1,
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?vehicle_id=1",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/vehicles/1.json",
  "app_url": "https://app.wenmarpro.com/vehicles/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "last_serviced_at": null,
  "lifetime_revenue_cents": 50000,
  "open_work_orders_count": 0,
  "appointments_count": 0
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicles/<id>.json
```

## Update vehicle

```
PATCH /vehicles/{id}
```

Update a vehicle by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Request body** — wrapper key `vehicle`:

| Field | Type | Required |
|---|---|---|
| `make` | string | Yes |
| `model` | string | No |
| `year` | integer | No |
| `vin` | string | No |
| `submodel` | string | No |
| `body_style` | string | No |
| `engine` | string | No |
| `transmission` | string | No |
| `drivetrain` | string | No |
| `color` | string | No |
| `license_plate` | string | No |
| `license_plate_state` | string | No |
| `odometer_reading` | integer | No |
| `odometer_unit` | string | No |
| `notes` | string | No |

**Response 200** — [Vehicle](#vehicle-schema)

**Example**

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
  "odometer": {
    "reading": 50000,
    "unit": "km"
  },
  "work_orders_count": 1,
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?vehicle_id=1",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/vehicles/1.json",
  "app_url": "https://app.wenmarpro.com/vehicles/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "last_serviced_at": null,
  "lifetime_revenue_cents": 50000,
  "open_work_orders_count": 0,
  "appointments_count": 0
}
```

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vehicles/<id>.json
```

## Merge vehicle

```
POST /vehicles/{id}/merge
```

Perform merge

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Request body**:

| Field | Type | Required |
|---|---|---|
| `source_vehicle_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `submodel` | any | Yes |
| `body_style` | any | Yes |
| `engine` | any | Yes |
| `vin` | string | Yes |
| `license_plate` | string | Yes |
| `license_plate_state` | string | Yes |
| `license_plate_country` | string | Yes |
| `drivetrain` | any | Yes |
| `transmission` | any | Yes |
| `color` | any | Yes |
| `vehicle_type` | string | Yes |
| `unit_number` | any | Yes |
| `fleet_identifier` | any | Yes |
| `production_date` | any | Yes |
| `annual_safety_expires_at` | any | Yes |
| `notes` | any | Yes |
| `odometer` | object | Yes |
| `work_orders_count` | integer | Yes |
| `work_orders_url` | string | Yes |
| `customer` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `last_serviced_at` | any | Yes |
| `lifetime_revenue_cents` | integer | Yes |
| `open_work_orders_count` | integer | Yes |
| `appointments_count` | integer | Yes |

`odometer` — object:
| Field | Type | Required |
|---|---|---|
| `reading` | any | Yes |
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

**Response 422** — [Error](#error-schema) error envelope

**Example**

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
  "odometer": {
    "reading": 50000,
    "unit": "km"
  },
  "work_orders_count": 1,
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?vehicle_id=1",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/vehicles/1.json",
  "app_url": "https://app.wenmarpro.com/vehicles/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "last_serviced_at": null,
  "lifetime_revenue_cents": 50000,
  "open_work_orders_count": 0,
  "appointments_count": 0
}
```

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vehicles/<id>.json
```

## Transfer vehicle

```
PATCH /vehicles/{id}/transfer
```

Perform transfer

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Request body**:

| Field | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |
| `mode` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `submodel` | any | Yes |
| `body_style` | any | Yes |
| `engine` | any | Yes |
| `vin` | string | Yes |
| `license_plate` | string | Yes |
| `license_plate_state` | string | Yes |
| `license_plate_country` | string | Yes |
| `drivetrain` | any | Yes |
| `transmission` | any | Yes |
| `color` | any | Yes |
| `vehicle_type` | string | Yes |
| `unit_number` | any | Yes |
| `fleet_identifier` | any | Yes |
| `production_date` | any | Yes |
| `annual_safety_expires_at` | any | Yes |
| `notes` | any | Yes |
| `odometer` | object | Yes |
| `work_orders_count` | integer | Yes |
| `work_orders_url` | string | Yes |
| `customer` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `last_serviced_at` | any | Yes |
| `lifetime_revenue_cents` | integer | Yes |
| `open_work_orders_count` | integer | Yes |
| `appointments_count` | integer | Yes |

`odometer` — object:
| Field | Type | Required |
|---|---|---|
| `reading` | any | Yes |
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

**Response 422** — [Error](#error-schema) error envelope

**Example**

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
  "odometer": {
    "reading": 50000,
    "unit": "km"
  },
  "work_orders_count": 1,
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?vehicle_id=1",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/vehicles/1.json",
  "app_url": "https://app.wenmarpro.com/vehicles/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "last_serviced_at": null,
  "lifetime_revenue_cents": 50000,
  "open_work_orders_count": 0,
  "appointments_count": 0
}
```

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/vehicles/<id>.json
```

## List vehicles work orders

```
GET /vehicles/{vehicle_id}/work_orders
```

List all vehicles work orders, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `vehicle_id` | integer | Yes |

**Response 200** — array of [WorkOrder](#workorder-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/vehicles/{vehicle_id}/work_orders.json
```

---

### Vehicle schema {#vehicle-schema}

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
| `vin` | string \| null | Yes |
| `license_plate` | string \| null | Yes |
| `license_plate_state` | string \| null | Yes |
| `license_plate_country` | string | Yes |
| `drivetrain` | string \| null | Yes |
| `transmission` | string \| null | Yes |
| `color` | string \| null | Yes |
| `vehicle_type` | string | Yes |
| `unit_number` | any | Yes |
| `fleet_identifier` | any | Yes |
| `production_date` | string \| null | Yes |
| `annual_safety_expires_at` | any | Yes |
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
| `last_serviced_at` | any | Yes |
| `lifetime_revenue_cents` | integer | Yes |
| `open_work_orders_count` | integer | Yes |
| `appointments_count` | integer | Yes |

`odometer` — object:
| Field | Type | Required |
|---|---|---|
| `reading` | integer \| null | Yes |
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

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### WorkOrder schema {#workorder-schema}

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `status` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | any | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | any | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer | Yes |
| `assigned_technician_id` | any | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `odometer_in` | any | Yes |
| `odometer_out` | any | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | any | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `vehicle_arrived_at` | string | Yes |
| `ready_for_pickup_at` | any | Yes |
| `completed_at` | any | Yes |
| `declined_at` | any | Yes |
| `decline_reason` | any | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | any | Yes |
| `closure_reason_notes` | any | Yes |
| `notes` | any | Yes |
| `purchase_order_number` | any | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | any | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `summary_url` | string | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`totals` — object:
| Field | Type | Required |
|---|---|---|
| `subtotal_cents` | integer | Yes |
| `tax_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `paid_cents` | integer | Yes |
| `remaining_cents` | integer | Yes |
| `currency` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

