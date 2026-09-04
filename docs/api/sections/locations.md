# Locations

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Show location

```
GET /locations/{id}
```

Show a location by ID.

| Param | Type | Required |
|---|---|---|
| `id` | string | Yes |

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

**Response 401** — [Error](#error-schema) error envelope

**Response 404** — [Error](#error-schema) error envelope

**Example**

```json
{
  "id": 1,
  "name": "Bay 1",
  "location_type": "shop",
  "currency": "CAD",
  "dock": [
    {
      "name": "work_orders",
      "enabled": true,
      "url": "https://app.wenmarpro.com/work_orders.json"
    }
  ],
  "url": "https://app.wenmarpro.com/locations/1.json",
  "app_url": "https://app.wenmarpro.com/locations/1"
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/<id>.json
```

## Update location

```
PATCH /locations/{id}
```

Update a location by ID.

| Param | Type | Required |
|---|---|---|
| `id` | string | Yes |

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

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/locations/<id>.json
```

## List locations business profile

```
GET /locations/{id}/business_profile
```

List all locations business profile, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |
| `tax_breakdown` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `trade_name` | string \| null | Yes |
| `location_type` | string | Yes |
| `gst_number` | any | Yes |
| `address` | string | Yes |
| `city` | string | Yes |
| `state` | string | Yes |
| `postal_code` | string | Yes |
| `country` | string | Yes |
| `contact_email` | any | Yes |
| `daily_hours` | object | Yes |
| `logo_url` | any | Yes |

`daily_hours` — object:
| Field | Type | Required |
|---|---|---|
| `1` | object | Yes |
| `2` | object | Yes |
| `3` | object | Yes |
| `4` | object | Yes |
| `5` | object | Yes |

`1` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`2` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`3` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`4` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`5` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`tax_breakdown` — object:
| Field | Type | Required |
|---|---|---|
| `state` | string | Yes |
| `gst_rate` | any | Yes |
| `pst_rate` | any | Yes |
| `hst_rate` | number | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/<id>.json
```

## Update locations business profile

```
PATCH /locations/{id}/business_profile
```

Update a locations business profile by ID.

| Param | Type | Required |
|---|---|---|
| `id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |
| `tax_breakdown` | any | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `trade_name` | string | Yes |
| `location_type` | string | Yes |
| `gst_number` | any | Yes |
| `address` | string | Yes |
| `city` | string | Yes |
| `state` | string | Yes |
| `postal_code` | string | Yes |
| `country` | string | Yes |
| `contact_email` | any | Yes |
| `daily_hours` | object | Yes |
| `logo_url` | any | Yes |

`daily_hours` — object:
| Field | Type | Required |
|---|---|---|
| `1` | object | Yes |
| `2` | object | Yes |
| `3` | object | Yes |
| `4` | object | Yes |
| `5` | object | Yes |

`1` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`2` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`3` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`4` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

`5` — object:
| Field | Type | Required |
|---|---|---|
| `open` | string | Yes |
| `close` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/locations/<id>.json
```

## List locations operations

```
GET /locations/{id}/operations
```

List all locations operations, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |
| `available_vendors` | array of any | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `time_zone` | string | Yes |
| `currency` | string | Yes |
| `daily_revenue_target` | object | Yes |
| `default_vendor_id` | integer \| null | Yes |
| `default_vendor_name` | string \| null | Yes |

`daily_revenue_target` — object:
| Field | Type | Required |
|---|---|---|
| `cents` | integer | Yes |
| `currency_iso` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/<id>.json
```

## Update locations operations

```
PATCH /locations/{id}/operations
```

Update a locations operations by ID.

| Param | Type | Required |
|---|---|---|
| `id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |
| `available_vendors` | array of any | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `time_zone` | string | Yes |
| `currency` | string | Yes |
| `daily_revenue_target` | object | Yes |
| `default_vendor_id` | integer \| null | Yes |
| `default_vendor_name` | string \| null | Yes |

`daily_revenue_target` — object:
| Field | Type | Required |
|---|---|---|
| `cents` | integer | Yes |
| `currency_iso` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/locations/<id>.json
```

## List locations close requirements

```
GET /locations/{location_id}/close_requirements
```

List all locations close requirements, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `oil_change_reminders_enabled` | boolean | Yes |
| `tire_swap_reminders_enabled` | boolean | Yes |
| `brake_inspection_reminders_enabled` | boolean | Yes |
| `battery_check_reminders_enabled` | boolean | Yes |
| `close_requirements` | object | Yes |
| `lead_source_requirements` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`close_requirements` — object:

`lead_source_requirements` — object:

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/{location_id}/close_requirements.json
```

## Update locations close requirements

```
PATCH /locations/{location_id}/close_requirements
```

Update a locations close requirements by ID.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `oil_change_reminders_enabled` | boolean | Yes |
| `tire_swap_reminders_enabled` | boolean | Yes |
| `brake_inspection_reminders_enabled` | boolean | Yes |
| `battery_check_reminders_enabled` | boolean | Yes |
| `close_requirements` | object | Yes |
| `lead_source_requirements` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`close_requirements` — object:
| Field | Type | Required |
|---|---|---|
| `odometer_in` | string | Yes |
| `key_location` | string | Yes |

`lead_source_requirements` — object:

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/locations/{location_id}/close_requirements.json
```

## List locations courtesy cars

```
GET /locations/{location_id}/courtesy_cars
```

List all locations courtesy cars, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `courtesy_cars` | object | Yes |

`courtesy_cars` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `courtesy_car_fee_enabled` | boolean | Yes |
| `courtesy_car_fee_cents` | integer | Yes |
| `courtesy_car_terms_text` | any | Yes |
| `url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/{location_id}/courtesy_cars.json
```

## List locations documents

```
GET /locations/{location_id}/documents
```

List all locations documents, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `document_settings` | object | Yes |
| `estimate_terms_text` | string \| null | Yes |
| `terms_text` | string | Yes |
| `payment_instructions` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`document_settings` — object:

**Response 401** — [Error](#error-schema) error envelope

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/{location_id}/documents.json
```

## Update locations documents

```
PATCH /locations/{location_id}/documents
```

Update a locations documents by ID.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `document_settings` | object | Yes |
| `estimate_terms_text` | string | Yes |
| `terms_text` | any | Yes |
| `payment_instructions` | any | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`document_settings` — object:

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/locations/{location_id}/documents.json
```

## List locations lead source requirements

```
GET /locations/{location_id}/lead_source_requirements
```

List all locations lead source requirements, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `oil_change_reminders_enabled` | boolean | Yes |
| `tire_swap_reminders_enabled` | boolean | Yes |
| `brake_inspection_reminders_enabled` | boolean | Yes |
| `battery_check_reminders_enabled` | boolean | Yes |
| `close_requirements` | object | Yes |
| `lead_source_requirements` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`close_requirements` — object:

`lead_source_requirements` — object:

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/{location_id}/lead_source_requirements.json
```

## Update locations lead source requirements

```
PATCH /locations/{location_id}/lead_source_requirements
```

Update a locations lead source requirements by ID.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `oil_change_reminders_enabled` | boolean | Yes |
| `tire_swap_reminders_enabled` | boolean | Yes |
| `brake_inspection_reminders_enabled` | boolean | Yes |
| `battery_check_reminders_enabled` | boolean | Yes |
| `close_requirements` | object | Yes |
| `lead_source_requirements` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`close_requirements` — object:

`lead_source_requirements` — object:
| Field | Type | Required |
|---|---|---|
| `customer_lead_source` | string | Yes |
| `ro_marketing_source` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/locations/{location_id}/lead_source_requirements.json
```

## List locations reminders

```
GET /locations/{location_id}/reminders
```

List all locations reminders, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `oil_change_reminders_enabled` | boolean | Yes |
| `tire_swap_reminders_enabled` | boolean | Yes |
| `brake_inspection_reminders_enabled` | boolean | Yes |
| `battery_check_reminders_enabled` | boolean | Yes |
| `close_requirements` | object | Yes |
| `lead_source_requirements` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`close_requirements` — object:

`lead_source_requirements` — object:

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/{location_id}/reminders.json
```

## Update locations reminders

```
PATCH /locations/{location_id}/reminders
```

Update a locations reminders by ID.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `oil_change_reminders_enabled` | boolean | Yes |
| `tire_swap_reminders_enabled` | boolean | Yes |
| `brake_inspection_reminders_enabled` | boolean | Yes |
| `battery_check_reminders_enabled` | boolean | Yes |
| `close_requirements` | object | Yes |
| `lead_source_requirements` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`close_requirements` — object:

`lead_source_requirements` — object:

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/locations/{location_id}/reminders.json
```

## List locations schedule config

```
GET /locations/{location_id}/schedule_config
```

List all locations schedule config, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `schedule_config` | object | Yes |

`schedule_config` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `slot_duration_minutes` | integer | Yes |
| `max_appointments_per_slot` | integer | Yes |
| `booking_window_days` | integer | Yes |
| `blocked_time_color` | string | Yes |
| `drop_off_color` | string | Yes |
| `waiter_color` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/locations/{location_id}/schedule_config.json
```

## Update locations schedule config

```
PATCH /locations/{location_id}/schedule_config
```

Update a locations schedule config by ID.

| Param | Type | Required |
|---|---|---|
| `location_id` | string | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `schedule_config` | object | Yes |

`schedule_config` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `slot_duration_minutes` | integer | Yes |
| `max_appointments_per_slot` | integer | Yes |
| `booking_window_days` | integer | Yes |
| `blocked_time_color` | string | Yes |
| `drop_off_color` | string | Yes |
| `waiter_color` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/locations/{location_id}/schedule_config.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateLocationRequest schema {#updatelocationrequest-schema}

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### UpdateLocationsBusinessProfileRequest schema {#updatelocationsbusinessprofilerequest-schema}

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `trade_name` | string | Yes |

---

### UpdateLocationsOperationsRequest schema {#updatelocationsoperationsrequest-schema}

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `time_zone` | string | Yes |

---

### UpdateLocationsCloseRequirementsRequest schema {#updatelocationscloserequirementsrequest-schema}

| Field | Type | Required |
|---|---|---|
| `close_requirements` | object | Yes |

`close_requirements` — object:
| Field | Type | Required |
|---|---|---|
| `odometer_in` | string | Yes |
| `key_location` | string | Yes |

---

### UpdateLocationsDocumentsRequest schema {#updatelocationsdocumentsrequest-schema}

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `estimate_terms_text` | string | Yes |

---

### UpdateLocationsLeadSourceRequirementsRequest schema {#updatelocationsleadsourcerequirementsrequest-schema}

| Field | Type | Required |
|---|---|---|
| `lead_source_requirements` | object | Yes |

`lead_source_requirements` — object:
| Field | Type | Required |
|---|---|---|
| `customer_lead_source` | string | Yes |
| `ro_marketing_source` | string | Yes |

---

### UpdateLocationsRemindersRequest schema {#updatelocationsremindersrequest-schema}

| Field | Type | Required |
|---|---|---|
| `location` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `oil_change_reminders_enabled` | boolean | Yes |
| `tire_swap_reminders_enabled` | boolean | Yes |

---

### UpdateLocationsScheduleConfigRequest schema {#updatelocationsscheduleconfigrequest-schema}

| Field | Type | Required |
|---|---|---|
| `schedule_config` | object | Yes |

`schedule_config` — object:
| Field | Type | Required |
|---|---|---|
| `slot_duration_minutes` | integer | Yes |

