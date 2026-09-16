# Work Orders

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List work orders

```
GET /work_orders
```

List all work orders, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `per_page` | integer | No |

**Response 200** — array of [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Example**

```json
[
  {
    "type": "WorkOrder",
    "id": 1,
    "work_order_number": 1001,
    "stage": "in_progress",
    "intake_method": "drop_off",
    "scheduled_for": null,
    "authorized": false,
    "paid": false,
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "closed_at": null,
    "location_id": 1,
    "service_advisor_id": 5,
    "assigned_technician_id": null,
    "sub_status_type_id": null,
    "payer_customer_id": null,
    "work_order_services_count": 0,
    "inspection_reports_count": 0,
    "customer": {
      "id": 1,
      "full_name": "Jane Doe",
      "url": "https://app.wenmarpro.com/customers/1.json",
      "display_name": "Jane Doe"
    },
    "vehicle": {
      "id": 1,
      "make": "Toyota",
      "model": "Camry",
      "year": 2020,
      "vin": "ABC123",
      "url": "https://app.wenmarpro.com/vehicles/1.json",
      "display_name": "2020 Toyota Camry",
      "license_plate": "JANE1"
    },
    "totals": {
      "subtotal_cents": 5000,
      "tax_cents": 650,
      "total_cents": 5650,
      "paid_cents": 0,
      "remaining_cents": 5650,
      "currency": "CAD"
    },
    "url": "https://app.wenmarpro.com/work_orders/1.json",
    "app_url": "https://app.wenmarpro.com/work_orders/1",
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "services_url": "https://app.wenmarpro.com/work_orders/1/estimate.json",
    "payments_url": "https://app.wenmarpro.com/work_orders/1/payments.json",
    "wip_url": "https://app.wenmarpro.com/work_orders/1/wip.json",
    "inspection_url": "https://app.wenmarpro.com/work_orders/1/inspection.json",
    "parts_url": "https://app.wenmarpro.com/work_orders/1/parts.json",
    "concerns_url": "https://app.wenmarpro.com/work_orders/1/concerns.json",
    "odometer_in": null,
    "odometer_out": null,
    "odometer_unit": "km",
    "authorized_at": null,
    "authorized_total_cents": 0,
    "customer_notified": false,
    "customer_notified_ready": false,
    "vehicle_arrived_at": "2026-08-27T12:00:00.000-04:00",
    "ready_for_pickup_at": null,
    "completed_at": null,
    "declined_at": null,
    "decline_reason": null,
    "discount_cents": 0,
    "fees_cents": 0,
    "parts_cents": 0,
    "labor_cents": 0,
    "tires_cents": 0,
    "subcontracts_cents": 0,
    "credit_balance_cents": 0,
    "saved_for_later": false,
    "closure_reason": null,
    "closure_reason_notes": null,
    "notes": null,
    "purchase_order_number": null,
    "return_method": "customer_pickup",
    "return_method_notes": null,
    "vehicle_keys_location": "front_desk",
    "vehicle_location": "bay_2",
    "service_history_url": "https://app.wenmarpro.com/work_orders/1/service_history.json",
    "declined_services_url": "https://app.wenmarpro.com/work_orders/1/declined_services.json",
    "customer_visit_count": 3,
    "customer_total_spend_cents": 45200,
    "average_ticket_cents": 15067,
    "activity_total": 2,
    "recent_activities": [],
    "services_visible_to_customer": true,
    "inspections_visible_to_customer": true
  },
  {
    "type": "WorkOrder",
    "id": 2,
    "work_order_number": 1002,
    "stage": "open",
    "intake_method": "drive_in",
    "scheduled_for": null,
    "authorized": false,
    "paid": false,
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "closed_at": null,
    "location_id": 1,
    "service_advisor_id": 5,
    "assigned_technician_id": null,
    "sub_status_type_id": null,
    "payer_customer_id": null,
    "work_order_services_count": 0,
    "inspection_reports_count": 0,
    "customer": {
      "id": 2,
      "full_name": "John Smith",
      "url": "https://app.wenmarpro.com/customers/2.json",
      "display_name": "Jane Doe"
    },
    "vehicle": {
      "id": 2,
      "make": "Honda",
      "model": "Civic",
      "year": 2018,
      "vin": "XYZ789",
      "url": "https://app.wenmarpro.com/vehicles/2.json",
      "display_name": "2018 Honda Civic",
      "license_plate": "JANE1"
    },
    "totals": {
      "subtotal_cents": 0,
      "tax_cents": 0,
      "total_cents": 0,
      "paid_cents": 0,
      "remaining_cents": 0,
      "currency": "CAD"
    },
    "url": "https://app.wenmarpro.com/work_orders/2.json",
    "app_url": "https://app.wenmarpro.com/work_orders/2",
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "services_url": "https://app.wenmarpro.com/work_orders/2/estimate.json",
    "payments_url": "https://app.wenmarpro.com/work_orders/2/payments.json",
    "wip_url": "https://app.wenmarpro.com/work_orders/2/wip.json",
    "inspection_url": "https://app.wenmarpro.com/work_orders/2/inspection.json",
    "parts_url": "https://app.wenmarpro.com/work_orders/2/parts.json",
    "concerns_url": "https://app.wenmarpro.com/work_orders/2/concerns.json",
    "odometer_in": null,
    "odometer_out": null,
    "odometer_unit": "km",
    "authorized_at": null,
    "authorized_total_cents": 0,
    "customer_notified": false,
    "customer_notified_ready": false,
    "vehicle_arrived_at": "2026-08-27T12:00:00.000-04:00",
    "ready_for_pickup_at": null,
    "completed_at": null,
    "declined_at": null,
    "decline_reason": null,
    "discount_cents": 0,
    "fees_cents": 0,
    "parts_cents": 0,
    "labor_cents": 0,
    "tires_cents": 0,
    "subcontracts_cents": 0,
    "credit_balance_cents": 0,
    "saved_for_later": false,
    "closure_reason": null,
    "closure_reason_notes": null,
    "notes": null,
    "purchase_order_number": null,
    "return_method": "customer_pickup",
    "return_method_notes": null,
    "vehicle_keys_location": "front_desk",
    "vehicle_location": "bay_2",
    "service_history_url": "https://app.wenmarpro.com/work_orders/2/service_history.json",
    "declined_services_url": "https://app.wenmarpro.com/work_orders/2/declined_services.json",
    "customer_visit_count": 3,
    "customer_total_spend_cents": 45200,
    "average_ticket_cents": 15067,
    "activity_total": 2,
    "recent_activities": [],
    "services_visible_to_customer": true,
    "inspections_visible_to_customer": true
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders.json
```

## Create work order

```
POST /work_orders
```

Create a work order.

**Response 201** — [WorkOrder](#workorder-schema)

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "stage": "in_progress",
  "intake_method": "drop_off",
  "scheduled_for": null,
  "authorized": true,
  "paid": false,
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "closed_at": null,
  "location_id": 1,
  "service_advisor_id": 5,
  "assigned_technician_id": null,
  "sub_status_type_id": null,
  "payer_customer_id": null,
  "work_order_services_count": 2,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "display_name": "Jane Doe"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "display_name": "2020 Toyota Camry",
    "license_plate": "JANE1"
  },
  "totals": {
    "subtotal_cents": 5000,
    "tax_cents": 650,
    "total_cents": 5650,
    "paid_cents": 0,
    "remaining_cents": 5650,
    "currency": "CAD"
  },
  "url": "https://app.wenmarpro.com/work_orders/1.json",
  "app_url": "https://app.wenmarpro.com/work_orders/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "services_url": "https://app.wenmarpro.com/work_orders/1/estimate.json",
  "payments_url": "https://app.wenmarpro.com/work_orders/1/payments.json",
  "wip_url": "https://app.wenmarpro.com/work_orders/1/wip.json",
  "inspection_url": "https://app.wenmarpro.com/work_orders/1/inspection.json",
  "parts_url": "https://app.wenmarpro.com/work_orders/1/parts.json",
  "concerns_url": "https://app.wenmarpro.com/work_orders/1/concerns.json",
  "service_history_url": "https://app.wenmarpro.com/work_orders/1/service_history.json",
  "declined_services_url": "https://app.wenmarpro.com/work_orders/1/declined_services.json",
  "customer_visit_count": 3,
  "customer_total_spend_cents": 45200,
  "average_ticket_cents": 15067,
  "activity_total": 2,
  "recent_activities": [
    {
      "id": 101,
      "category": "services",
      "description": "Oil Change added",
      "created_at": "2026-08-27T12:00:00.000-04:00"
    },
    {
      "id": 102,
      "category": "contact",
      "description": "status changed from estimate to in_progress",
      "created_at": "2026-08-27T12:30:00.000-04:00"
    }
  ],
  "odometer_in": null,
  "odometer_out": null,
  "odometer_unit": "km",
  "authorized_at": "2026-08-27T12:30:00.000-04:00",
  "authorized_total_cents": 5650,
  "customer_notified": true,
  "customer_notified_ready": false,
  "vehicle_arrived_at": "2026-08-27T12:00:00.000-04:00",
  "ready_for_pickup_at": null,
  "completed_at": null,
  "declined_at": null,
  "decline_reason": null,
  "discount_cents": 0,
  "fees_cents": 0,
  "parts_cents": 2500,
  "labor_cents": 2500,
  "tires_cents": 0,
  "subcontracts_cents": 0,
  "credit_balance_cents": 0,
  "saved_for_later": false,
  "closure_reason": null,
  "closure_reason_notes": null,
  "notes": null,
  "purchase_order_number": null,
  "return_method": "customer_pickup",
  "return_method_notes": null,
  "vehicle_keys_location": "front_desk",
  "vehicle_location": "bay_2",
  "services_visible_to_customer": true,
  "inspections_visible_to_customer": true
}
```

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders.json
```

## Create work orders quick intake

```
POST /work_orders/quick_intake
```

Create a work orders quick intake.

**Response 201** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/quick_intake.json
```

## Show work order

```
GET /work_orders/{id}
```

Show a work order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 401** — [Error](#error-schema) error envelope

**Response 404** — [Error](#error-schema) error envelope

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "stage": "in_progress",
  "intake_method": "drop_off",
  "scheduled_for": null,
  "authorized": true,
  "paid": false,
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "closed_at": null,
  "location_id": 1,
  "service_advisor_id": 5,
  "assigned_technician_id": null,
  "sub_status_type_id": null,
  "payer_customer_id": null,
  "work_order_services_count": 2,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "display_name": "Jane Doe"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "display_name": "2020 Toyota Camry",
    "license_plate": "JANE1"
  },
  "totals": {
    "subtotal_cents": 5000,
    "tax_cents": 650,
    "total_cents": 5650,
    "paid_cents": 0,
    "remaining_cents": 5650,
    "currency": "CAD"
  },
  "url": "https://app.wenmarpro.com/work_orders/1.json",
  "app_url": "https://app.wenmarpro.com/work_orders/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "services_url": "https://app.wenmarpro.com/work_orders/1/estimate.json",
  "payments_url": "https://app.wenmarpro.com/work_orders/1/payments.json",
  "wip_url": "https://app.wenmarpro.com/work_orders/1/wip.json",
  "inspection_url": "https://app.wenmarpro.com/work_orders/1/inspection.json",
  "parts_url": "https://app.wenmarpro.com/work_orders/1/parts.json",
  "concerns_url": "https://app.wenmarpro.com/work_orders/1/concerns.json",
  "service_history_url": "https://app.wenmarpro.com/work_orders/1/service_history.json",
  "declined_services_url": "https://app.wenmarpro.com/work_orders/1/declined_services.json",
  "customer_visit_count": 3,
  "customer_total_spend_cents": 45200,
  "average_ticket_cents": 15067,
  "activity_total": 2,
  "recent_activities": [
    {
      "id": 101,
      "category": "services",
      "description": "Oil Change added",
      "created_at": "2026-08-27T12:00:00.000-04:00"
    },
    {
      "id": 102,
      "category": "contact",
      "description": "status changed from estimate to in_progress",
      "created_at": "2026-08-27T12:30:00.000-04:00"
    }
  ],
  "odometer_in": null,
  "odometer_out": null,
  "odometer_unit": "km",
  "authorized_at": "2026-08-27T12:30:00.000-04:00",
  "authorized_total_cents": 5650,
  "customer_notified": true,
  "customer_notified_ready": false,
  "vehicle_arrived_at": "2026-08-27T12:00:00.000-04:00",
  "ready_for_pickup_at": null,
  "completed_at": null,
  "declined_at": null,
  "decline_reason": null,
  "discount_cents": 0,
  "fees_cents": 0,
  "parts_cents": 2500,
  "labor_cents": 2500,
  "tires_cents": 0,
  "subcontracts_cents": 0,
  "credit_balance_cents": 0,
  "saved_for_later": false,
  "closure_reason": null,
  "closure_reason_notes": null,
  "notes": null,
  "purchase_order_number": null,
  "return_method": "customer_pickup",
  "return_method_notes": null,
  "vehicle_keys_location": "front_desk",
  "vehicle_location": "bay_2",
  "services_visible_to_customer": true,
  "inspections_visible_to_customer": true
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work order

```
PATCH /work_orders/{id}
```

Update a work order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "stage": "in_progress",
  "intake_method": "drop_off",
  "scheduled_for": null,
  "authorized": true,
  "paid": false,
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "closed_at": null,
  "location_id": 1,
  "service_advisor_id": 5,
  "assigned_technician_id": null,
  "sub_status_type_id": null,
  "payer_customer_id": null,
  "work_order_services_count": 2,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "display_name": "Jane Doe"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "display_name": "2020 Toyota Camry",
    "license_plate": "JANE1"
  },
  "totals": {
    "subtotal_cents": 5000,
    "tax_cents": 650,
    "total_cents": 5650,
    "paid_cents": 0,
    "remaining_cents": 5650,
    "currency": "CAD"
  },
  "url": "https://app.wenmarpro.com/work_orders/1.json",
  "app_url": "https://app.wenmarpro.com/work_orders/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "services_url": "https://app.wenmarpro.com/work_orders/1/estimate.json",
  "payments_url": "https://app.wenmarpro.com/work_orders/1/payments.json",
  "wip_url": "https://app.wenmarpro.com/work_orders/1/wip.json",
  "inspection_url": "https://app.wenmarpro.com/work_orders/1/inspection.json",
  "parts_url": "https://app.wenmarpro.com/work_orders/1/parts.json",
  "concerns_url": "https://app.wenmarpro.com/work_orders/1/concerns.json",
  "service_history_url": "https://app.wenmarpro.com/work_orders/1/service_history.json",
  "declined_services_url": "https://app.wenmarpro.com/work_orders/1/declined_services.json",
  "customer_visit_count": 3,
  "customer_total_spend_cents": 45200,
  "average_ticket_cents": 15067,
  "activity_total": 2,
  "recent_activities": [
    {
      "id": 101,
      "category": "services",
      "description": "Oil Change added",
      "created_at": "2026-08-27T12:00:00.000-04:00"
    },
    {
      "id": 102,
      "category": "contact",
      "description": "status changed from estimate to in_progress",
      "created_at": "2026-08-27T12:30:00.000-04:00"
    }
  ],
  "odometer_in": null,
  "odometer_out": null,
  "odometer_unit": "km",
  "authorized_at": "2026-08-27T12:30:00.000-04:00",
  "authorized_total_cents": 5650,
  "customer_notified": true,
  "customer_notified_ready": false,
  "vehicle_arrived_at": "2026-08-27T12:00:00.000-04:00",
  "ready_for_pickup_at": null,
  "completed_at": null,
  "declined_at": null,
  "decline_reason": null,
  "discount_cents": 0,
  "fees_cents": 0,
  "parts_cents": 2500,
  "labor_cents": 2500,
  "tires_cents": 0,
  "subcontracts_cents": 0,
  "credit_balance_cents": 0,
  "saved_for_later": false,
  "closure_reason": null,
  "closure_reason_notes": null,
  "notes": null,
  "purchase_order_number": null,
  "return_method": "customer_pickup",
  "return_method_notes": null,
  "vehicle_keys_location": "front_desk",
  "vehicle_location": "bay_2",
  "services_visible_to_customer": true,
  "inspections_visible_to_customer": true
}
```

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders activity

```
GET /work_orders/{id}/activity
```

List all work orders activity, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `category` | string | No |
| `id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `action` | string | Yes |
| `category` | string | Yes |
| `body` | string | Yes |
| `actor_name` | string | Yes |
| `actor_type` | string | Yes |
| `created_at` | string | Yes |
| `work_order_id` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders appointments

```
GET /work_orders/{id}/appointments
```

List all work orders appointments, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — array of [Appointment](#appointment-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders authorization logs

```
GET /work_orders/{id}/authorization_logs
```

List all work orders authorization logs, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `event_type` | string | Yes |
| `method` | string | Yes |
| `authorized_by_name` | string \| null | Yes |
| `authorized_by_phone` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `created_at` | string | Yes |
| `actor` | object | Yes |
| `work_order_service_id` | integer | Yes |
| `authorization_batch_id` | integer \| null | Yes |
| `snapshot` | object | Yes |

`actor` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |

`snapshot` — object:
| Field | Type | Required |
|---|---|---|
| `snapshot_version` | integer | Yes |
| `timestamp` | string | Yes |
| `work_order` | object | Yes |
| `services` | array of object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `stage` | string | Yes |
| `total_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `discount_cents` | integer | Yes |
| `subtotal_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `core_charges_cents` | integer | Yes |
| `labor_tax_cents` | integer | Yes |
| `parts_tax_cents` | integer | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Close work order

```
POST /work_orders/{id}/close
```

Close

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |
| `reopen_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Complete work order

```
POST /work_orders/{id}/complete
```

Complete

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer | Yes |
| `odometer_out` | integer | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |
| `reopen_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Delete work orders courtesy car assignment

```
DELETE /work_orders/{id}/courtesy_car_assignment
```

Delete a work orders courtesy car assignment by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders courtesy car assignment

```
PATCH /work_orders/{id}/courtesy_car_assignment
```

Update a work orders courtesy car assignment by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders courtesy car assignment

```
POST /work_orders/{id}/courtesy_car_assignment
```

Create a work orders courtesy car assignment.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Decline work order

```
POST /work_orders/{id}/decline
```

Decline

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |
| `reopen_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Show work order declined services

```
GET /work_orders/{id}/declined_services
```

Show a work order declined services by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `declined_at` | string | Yes |
| `work_order_number` | integer | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Post work order to account

```
POST /work_orders/{id}/post_to_account
```

Post to account

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |
| `reopen_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Reopen work order

```
POST /work_orders/{id}/reopen
```

Reopen

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Send work order estimate

```
POST /work_orders/{id}/send_estimate
```

Send estimate

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of any | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Send work order invoice summary

```
POST /work_orders/{id}/send_invoice_summary
```

Send invoice summary

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |
| `reopen_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Send work order reminder

```
POST /work_orders/{id}/send_reminder
```

Send reminder

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of any | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Show work order service history

```
GET /work_orders/{id}/service_history
```

Show a work order service history by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `authorization_status` | string | Yes |
| `completed_at` | string | Yes |
| `work_order_number` | integer | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Start work order

```
POST /work_orders/{id}/start
```

Start

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of any | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders vehicle history

```
GET /work_orders/{id}/vehicle_history
```

List all work orders vehicle history, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `created_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `services_summary` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Void work order

```
POST /work_orders/{id}/void
```

Void

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |
| `reopen_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders activity log

```
POST /work_orders/{work_order_id}/activity_logs
```

Create a work orders activity log.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/activity_logs.json
```

## Create work orders authorization decision

```
POST /work_orders/{work_order_id}/authorization_decisions
```

Create a work orders authorization decision.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/authorization_decisions.json
```

## Create work orders authorization

```
POST /work_orders/{work_order_id}/authorizations
```

Create a work orders authorization.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/authorizations.json
```

## List work orders concerns

```
GET /work_orders/{work_order_id}/concerns
```

List all work orders concerns, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `body` | string | Yes |
| `source` | string | Yes |
| `severity` | string \| null | Yes |
| `finding_notes` | string | Yes |
| `position` | integer | Yes |
| `converted` | boolean | Yes |
| `converted_service_id` | integer \| null | Yes |
| `creator` | object | Yes |
| `finding_added_by_name` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer \| null | Yes |
| `name` | string \| null | Yes |
| `url` | string \| null | Yes |

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/concerns.json
```

## Create work orders concern

```
POST /work_orders/{work_order_id}/concerns
```

Create a work orders concern.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/concerns.json
```

## Create work orders concerns copy all to estimate

```
POST /work_orders/{work_order_id}/concerns/copy_all_to_estimate
```

Create a work orders concerns copy all to estimate.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/concerns/copy_all_to_estimate.json
```

## Decline all work order concerns

```
POST /work_orders/{work_order_id}/concerns/decline_all
```

Decline all

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string | Yes |
| `decline_reason` | string | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/concerns/decline_all.json
```

## Delete work orders concern

```
DELETE /work_orders/{work_order_id}/concerns/{id}
```

Delete a work orders concern by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders concern

```
PATCH /work_orders/{work_order_id}/concerns/{id}
```

Update a work orders concern by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders concerns add package

```
POST /work_orders/{work_order_id}/concerns/{id}/add_package
```

Create a work orders concerns add package.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders concerns copy to estimate

```
POST /work_orders/{work_order_id}/concerns/{id}/copy_to_estimate
```

Create a work orders concerns copy to estimate.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders credit resolution

```
POST /work_orders/{work_order_id}/credit_resolution
```

Create a work orders credit resolution.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/credit_resolution.json
```

## Show work order estimate

```
GET /work_orders/{work_order_id}/estimate
```

Show a work order estimate by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "stage": "in_progress",
  "intake_method": "drop_off",
  "scheduled_for": null,
  "authorized": true,
  "paid": false,
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "closed_at": null,
  "location_id": 1,
  "service_advisor_id": 5,
  "assigned_technician_id": null,
  "sub_status_type_id": null,
  "payer_customer_id": null,
  "vehicle_arrived_at": null,
  "work_order_services_count": 1,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "display_name": "Jane Doe"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "display_name": "2020 Toyota Camry",
    "license_plate": "JANE1"
  },
  "totals": {
    "subtotal_cents": 5000,
    "tax_cents": 650,
    "total_cents": 5650,
    "paid_cents": 0,
    "remaining_cents": 5650,
    "currency": "CAD"
  },
  "url": "https://app.wenmarpro.com/work_orders/1.json",
  "app_url": "https://app.wenmarpro.com/work_orders/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "services": [
    {
      "id": 1,
      "name": "Brake Inspection",
      "service_type": "labor",
      "authorization_status": "approved",
      "pricing_mode": "flat",
      "technician_id": 2,
      "category_id": 1,
      "ordinal": 1,
      "discount_cents": 0,
      "labor_cents": 2500,
      "parts_cents": 2500,
      "fees_cents": 0,
      "sublet_cents": 0,
      "tires_cents": 0,
      "total_cents": 5000,
      "tax_total_cents": 650,
      "estimated_hours": "1.5",
      "customer_notes": null,
      "started_at": null,
      "completed_at": null,
      "authorized_at": "2026-08-27T12:30:00.000-04:00",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "updated_at": "2026-08-27T12:00:00.000-04:00",
      "line_items": []
    }
  ],
  "services_visible_to_customer": true,
  "inspections_visible_to_customer": true
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/estimate.json
```

## Create work orders fee exclusion

```
POST /work_orders/{work_order_id}/fee_exclusions
```

Create a work orders fee exclusion.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/fee_exclusions.json
```

## Delete work orders fee exclusion

```
DELETE /work_orders/{work_order_id}/fee_exclusions/{id}
```

Delete a work orders fee exclusion by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Show work order inspection

```
GET /work_orders/{work_order_id}/inspection
```

Show a work order inspection by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "stage": "in_progress",
  "intake_method": "drop_off",
  "scheduled_for": null,
  "authorized": true,
  "paid": false,
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "closed_at": null,
  "location_id": 1,
  "service_advisor_id": 5,
  "assigned_technician_id": null,
  "sub_status_type_id": null,
  "payer_customer_id": null,
  "vehicle_arrived_at": null,
  "work_order_services_count": 0,
  "inspection_reports_count": 1,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "display_name": "Jane Doe"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "display_name": "2020 Toyota Camry",
    "license_plate": "JANE1"
  },
  "totals": {
    "subtotal_cents": 5000,
    "tax_cents": 650,
    "total_cents": 5650,
    "paid_cents": 0,
    "remaining_cents": 5650,
    "currency": "CAD"
  },
  "url": "https://app.wenmarpro.com/work_orders/1.json",
  "app_url": "https://app.wenmarpro.com/work_orders/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "inspection_reports": [
    {
      "id": 1,
      "name": "Courtesy Check",
      "status": "active"
    }
  ],
  "services_visible_to_customer": true,
  "inspections_visible_to_customer": true
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/inspection.json
```

## Show work order parts

```
GET /work_orders/{work_order_id}/parts
```

Show a work order parts by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "stage": "in_progress",
  "intake_method": "drop_off",
  "scheduled_for": null,
  "authorized": true,
  "paid": false,
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "closed_at": null,
  "location_id": 1,
  "service_advisor_id": 5,
  "assigned_technician_id": null,
  "sub_status_type_id": null,
  "payer_customer_id": null,
  "vehicle_arrived_at": null,
  "work_order_services_count": 1,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "display_name": "Jane Doe"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "display_name": "2020 Toyota Camry",
    "license_plate": "JANE1"
  },
  "totals": {
    "subtotal_cents": 5000,
    "tax_cents": 650,
    "total_cents": 5650,
    "paid_cents": 0,
    "remaining_cents": 5650,
    "currency": "CAD"
  },
  "url": "https://app.wenmarpro.com/work_orders/1.json",
  "app_url": "https://app.wenmarpro.com/work_orders/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "services": [
    {
      "id": 1,
      "name": "Brake Inspection",
      "service_type": "labor",
      "authorization_status": "approved",
      "pricing_mode": "flat",
      "technician_id": 2,
      "category_id": 1,
      "ordinal": 1,
      "discount_cents": 0,
      "labor_cents": 2500,
      "parts_cents": 2500,
      "fees_cents": 0,
      "sublet_cents": 0,
      "tires_cents": 0,
      "total_cents": 5000,
      "tax_total_cents": 650,
      "estimated_hours": "1.5",
      "customer_notes": null,
      "started_at": null,
      "completed_at": null,
      "authorized_at": "2026-08-27T12:30:00.000-04:00",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "updated_at": "2026-08-27T12:00:00.000-04:00",
      "line_items": []
    }
  ],
  "services_visible_to_customer": true,
  "inspections_visible_to_customer": true
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/parts.json
```

## List work orders payment link

```
GET /work_orders/{work_order_id}/payment_link
```

List all work orders payment link, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/payment_link.json
```

## Create work orders payment link send

```
POST /work_orders/{work_order_id}/payment_link/send
```

Create a work orders payment link send.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/payment_link/send.json
```

## Show work order payments

```
GET /work_orders/{work_order_id}/payments
```

Show a work order payments by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "stage": "completed",
  "intake_method": "drop_off",
  "scheduled_for": null,
  "authorized": true,
  "paid": true,
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "closed_at": "2026-08-27T15:00:00.000-04:00",
  "location_id": 1,
  "service_advisor_id": 5,
  "assigned_technician_id": null,
  "sub_status_type_id": null,
  "payer_customer_id": null,
  "vehicle_arrived_at": null,
  "work_order_services_count": 1,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "display_name": "Jane Doe"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "display_name": "2020 Toyota Camry",
    "license_plate": "JANE1"
  },
  "totals": {
    "subtotal_cents": 5000,
    "tax_cents": 650,
    "total_cents": 5650,
    "paid_cents": 5650,
    "remaining_cents": 0,
    "currency": "CAD"
  },
  "url": "https://app.wenmarpro.com/work_orders/1.json",
  "app_url": "https://app.wenmarpro.com/work_orders/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "payments": [
    {
      "id": 1,
      "amount_cents": 5650,
      "method": "cash",
      "processor_status": "succeeded",
      "is_refund": false,
      "processed_at": "2026-08-27T15:00:00.000-04:00",
      "reference": null,
      "created_at": "2026-08-27T15:00:00.000-04:00",
      "updated_at": "2026-08-27T15:00:00.000-04:00",
      "work_order_id": 1,
      "customer_id": 1,
      "processed_by": {
        "id": 5,
        "full_name": "Alex Rivera"
      },
      "work_order": {
        "id": 1,
        "url": "https://app.wenmarpro.com/work_orders/1.json"
      },
      "customer": {
        "id": 1,
        "full_name": "Jane Doe",
        "url": "https://app.wenmarpro.com/customers/1.json"
      }
    }
  ],
  "services_visible_to_customer": true,
  "inspections_visible_to_customer": true
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/payments.json
```

## Create work order payment

```
POST /work_orders/{work_order_id}/payments
```

Create a work order payment.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/payments.json
```

## Reverse work order payment ar

```
DELETE /work_orders/{work_order_id}/payments/reverse_ar
```

Reverse ar

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |
| `reopen_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/payments/reverse_ar.json
```

## Send work order payment to ar

```
POST /work_orders/{work_order_id}/payments/send_to_ar
```

Send to ar

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | Yes |
| `odometer_out` | integer \| null | Yes |
| `odometer_unit` | string | Yes |
| `authorized_at` | string \| null | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `ready_for_pickup_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | string \| null | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string \| null | Yes |
| `closure_reason_notes` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `purchase_order_number` | string \| null | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | string \| null | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |
| `activity_url` | string | Yes |
| `vehicle_history_url` | string | Yes |
| `appointments_url` | string | Yes |
| `authorization_logs_url` | string | Yes |
| `reopen_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/payments/send_to_ar.json
```

## Create work orders purchase order

```
POST /work_orders/{work_order_id}/purchase_orders
```

Create a work orders purchase order.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [PurchaseOrder](#purchaseorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/purchase_orders.json
```

## Delete work orders purchase order

```
DELETE /work_orders/{work_order_id}/purchase_orders/{id}
```

Delete a work orders purchase order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders purchase order

```
PATCH /work_orders/{work_order_id}/purchase_orders/{id}
```

Update a work orders purchase order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [PurchaseOrder](#purchaseorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders purchase orders receive

```
POST /work_orders/{work_order_id}/purchase_orders/{id}/receive
```

Create a work orders purchase orders receive.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [PurchaseOrder](#purchaseorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders purchase orders return

```
POST /work_orders/{work_order_id}/purchase_orders/{id}/return
```

Create a work orders purchase orders return.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [PurchaseOrder](#purchaseorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders receipts

```
GET /work_orders/{work_order_id}/receipts
```

List all work orders receipts, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `currency` | string | Yes |
| `method` | string | Yes |
| `processor_status` | string | Yes |
| `is_refund` | boolean | Yes |
| `is_adjustment` | boolean | Yes |
| `voided` | boolean | Yes |
| `voided_at` | string \| null | Yes |
| `processed_at` | string | Yes |
| `reference` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `work_order_id` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/receipts.json
```

## Create work orders refund

```
POST /work_orders/{work_order_id}/refunds
```

Create a work orders refund.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/refunds.json
```

## List work order services

```
GET /work_orders/{work_order_id}/services
```

List all work order services, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `service_type` | string | Yes |
| `authorization_status` | string | Yes |
| `pricing_mode` | string | Yes |
| `technician_id` | integer \| null | Yes |
| `category_id` | integer \| null | Yes |
| `ordinal` | integer | Yes |
| `discount_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `sublet_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `estimated_hours` | string | Yes |
| `customer_notes` | string \| null | Yes |
| `started_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `authorized_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `line_items` | array of object | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/services.json
```

## Create work orders service

```
POST /work_orders/{work_order_id}/services
```

Create a work orders service.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

**Response 404** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/services.json
```

## Update work orders services reorder

```
PATCH /work_orders/{work_order_id}/services/reorder
```

Update a work orders services reorder by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `service_type` | string | Yes |
| `authorization_status` | string | Yes |
| `pricing_mode` | string | Yes |
| `technician_id` | integer \| null | Yes |
| `category_id` | integer \| null | Yes |
| `ordinal` | integer | Yes |
| `discount_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `sublet_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `estimated_hours` | string | Yes |
| `customer_notes` | string \| null | Yes |
| `started_at` | string \| null | Yes |
| `completed_at` | string \| null | Yes |
| `authorized_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `line_items` | array of object | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/services/reorder.json
```

## Delete work orders service

```
DELETE /work_orders/{work_order_id}/services/{id}
```

Delete a work orders service by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders service

```
PATCH /work_orders/{work_order_id}/services/{id}
```

Update a work orders service by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 403** — no content.

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders services adjust time

```
GET /work_orders/{work_order_id}/services/{id}/adjust_time
```

List all work orders services adjust time, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 406** — no content.

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders services adjust time

```
PATCH /work_orders/{work_order_id}/services/{id}/adjust_time
```

Update a work orders services adjust time by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Delete work orders services authorization

```
DELETE /work_orders/{work_order_id}/services/{id}/authorization
```

Delete a work orders services authorization by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services bulk pull

```
POST /work_orders/{work_order_id}/services/{id}/bulk_pull
```

Create a work orders services bulk pull.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Delete work orders services completion

```
DELETE /work_orders/{work_order_id}/services/{id}/completion
```

Delete a work orders services completion by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services completion

```
POST /work_orders/{work_order_id}/services/{id}/completion
```

Create a work orders services completion.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services copy

```
POST /work_orders/{work_order_id}/services/{id}/copies
```

Create a work orders services copy.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services package

```
POST /work_orders/{work_order_id}/services/{id}/packages
```

Create a work orders services package.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 404** — no content.

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders services pause

```
PATCH /work_orders/{work_order_id}/services/{id}/pause
```

Update a work orders services pause by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders services publish

```
PATCH /work_orders/{work_order_id}/services/{id}/publish
```

Update a work orders services publish by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders services revive

```
PATCH /work_orders/{work_order_id}/services/{id}/revive
```

Update a work orders services revive by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services time entry

```
POST /work_orders/{work_order_id}/services/{id}/time_entries
```

Create a work orders services time entry.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders services toggle labor completion

```
PATCH /work_orders/{work_order_id}/services/{id}/toggle_labor_completion
```

Update a work orders services toggle labor completion by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders services update category

```
PATCH /work_orders/{work_order_id}/services/{id}/update_category
```

Update a work orders services update category by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders services comments

```
GET /work_orders/{work_order_id}/services/{service_id}/comments
```

List all work orders services comments, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `body` | string | Yes |
| `user` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |

`user` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/services/{service_id}/comments.json
```

## Create work orders services comment

```
POST /work_orders/{work_order_id}/services/{service_id}/comments
```

Create a work orders services comment.

| Param | Type | Required |
|---|---|---|
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/services/{service_id}/comments.json
```

## Delete work orders services comment

```
DELETE /work_orders/{work_order_id}/services/{service_id}/comments/{id}
```

Delete a work orders services comment by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services extraction

```
POST /work_orders/{work_order_id}/services/{service_id}/extractions
```

Create a work orders services extraction.

| Param | Type | Required |
|---|---|---|
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 202** — [WorkOrder](#workorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 404** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/services/{service_id}/extractions.json
```

## List work orders services line items

```
GET /work_orders/{work_order_id}/services/{service_id}/line_items
```

List all work orders services line items, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `item_type` | string | Yes |
| `description` | string | Yes |
| `quantity` | string | Yes |
| `unit_price_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `tax_total_cents` | integer | Yes |
| `pricing_mode` | string | Yes |
| `is_taxable` | boolean | Yes |
| `is_warranty` | boolean | Yes |
| `completed` | boolean | Yes |
| `notes` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `part_type` | string | Yes |
| `part_status` | string | Yes |
| `vendor_id` | integer \| null | Yes |
| `vendor_part_number` | string \| null | Yes |
| `unit_cost_cents` | integer | Yes |
| `brand` | string \| null | Yes |
| `dot_code` | string \| null | Yes |
| `tire_size` | string \| null | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/services/{service_id}/line_items.json
```

## Create work orders services line item

```
POST /work_orders/{work_order_id}/services/{service_id}/line_items
```

Create a work orders services line item.

| Param | Type | Required |
|---|---|---|
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/services/{service_id}/line_items.json
```

## Delete work orders services line item

```
DELETE /work_orders/{work_order_id}/services/{service_id}/line_items/{id}
```

Delete a work orders services line item by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders services line item

```
PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}
```

Update a work orders services line item by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services line items copy

```
POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/copies
```

Create a work orders services line items copy.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services line items inventory addition

```
POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/inventory_additions
```

Create a work orders services line items inventory addition.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services line items price refresh

```
POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/price_refreshes
```

Create a work orders services line items price refresh.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Pull work order service line item

```
PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/pull
```

Pull

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services line items reorder

```
POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/reorder
```

Create a work orders services line items reorder.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Undo pull work order service line item

```
PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_pull
```

Undo pull

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Undo return work order service line item

```
PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_return
```

Undo return

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services sublet order

```
POST /work_orders/{work_order_id}/services/{service_id}/sublet_orders
```

Create a work orders services sublet order.

| Param | Type | Required |
|---|---|---|
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/services/{service_id}/sublet_orders.json
```

## Delete work orders services sublet order

```
DELETE /work_orders/{work_order_id}/services/{service_id}/sublet_orders/{id}
```

Delete a work orders services sublet order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders services sublet order

```
PATCH /work_orders/{work_order_id}/services/{service_id}/sublet_orders/{id}
```

Update a work orders services sublet order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders services sublet orders duplicate

```
POST /work_orders/{work_order_id}/services/{service_id}/sublet_orders/{id}/duplicate
```

Create a work orders services sublet orders duplicate.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `service_id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Show work orders signature

```
GET /work_orders/{work_order_id}/signatures/{id}
```

Show a work orders signature by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders sublet orders

```
GET /work_orders/{work_order_id}/sublet_orders
```

List all work orders sublet orders, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — array of [SubletOrder](#subletorder-schema)

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/sublet_orders.json
```

## Show work orders sublet order

```
GET /work_orders/{work_order_id}/sublet_orders/{id}
```

Show a work orders sublet order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [SubletOrder](#subletorder-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## List work orders tire storage

```
GET /work_orders/{work_order_id}/tire_storage
```

List all work orders tire storage, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `slot_label` | string | Yes |
| `season` | string | Yes |
| `season_label` | string | Yes |
| `tire_set_description` | string \| null | Yes |
| `stored_at` | string | Yes |
| `released_at` | string \| null | Yes |
| `currently_stored` | boolean | Yes |
| `storage_fee_cents` | integer \| null | Yes |
| `fee_type` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

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
| `name` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/tire_storage.json
```

## Create work orders tire storage

```
POST /work_orders/{work_order_id}/tire_storage
```

Create a work orders tire storage.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/tire_storage.json
```

## Delete work orders tire storage

```
DELETE /work_orders/{work_order_id}/tire_storage/{id}
```

Delete a work orders tire storage by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 204** — no content.

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Show work orders tire storage

```
GET /work_orders/{work_order_id}/tire_storage/{id}
```

Show a work orders tire storage by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
```

## Update work orders tire storage

```
PATCH /work_orders/{work_order_id}/tire_storage/{id}
```

Update a work orders tire storage by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work orders void

```
POST /work_orders/{work_order_id}/voids
```

Create a work orders void.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/voids.json
```

## Show work order WIP

```
GET /work_orders/{work_order_id}/wip
```

Show a work order wip by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — [WorkOrder](#workorder-schema)

**Response 404** — [Error](#error-schema) error envelope

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "stage": "in_progress",
  "intake_method": "drop_off",
  "scheduled_for": null,
  "authorized": true,
  "paid": false,
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "closed_at": null,
  "location_id": 1,
  "service_advisor_id": 5,
  "assigned_technician_id": null,
  "sub_status_type_id": null,
  "payer_customer_id": null,
  "vehicle_arrived_at": null,
  "work_order_services_count": 1,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "display_name": "Jane Doe"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json",
    "display_name": "2020 Toyota Camry",
    "license_plate": "JANE1"
  },
  "totals": {
    "subtotal_cents": 5000,
    "tax_cents": 650,
    "total_cents": 5650,
    "paid_cents": 0,
    "remaining_cents": 5650,
    "currency": "CAD"
  },
  "url": "https://app.wenmarpro.com/work_orders/1.json",
  "app_url": "https://app.wenmarpro.com/work_orders/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "services": [
    {
      "id": 1,
      "name": "Brake Inspection",
      "service_type": "labor",
      "authorization_status": "approved",
      "pricing_mode": "flat",
      "technician_id": 2,
      "category_id": 1,
      "ordinal": 1,
      "discount_cents": 0,
      "labor_cents": 2500,
      "parts_cents": 2500,
      "fees_cents": 0,
      "sublet_cents": 0,
      "tires_cents": 0,
      "total_cents": 5000,
      "tax_total_cents": 650,
      "estimated_hours": "1.5",
      "customer_notes": null,
      "started_at": null,
      "completed_at": null,
      "authorized_at": "2026-08-27T12:30:00.000-04:00",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "updated_at": "2026-08-27T12:00:00.000-04:00",
      "line_items": []
    }
  ],
  "services_visible_to_customer": true,
  "inspections_visible_to_customer": true
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/wip.json
```

---

### WorkOrder schema {#workorder-schema}

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `stage` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | string \| null | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `assigned_technician_id` | integer \| null | Yes |
| `sub_status_type_id` | integer \| null | Yes |
| `payer_customer_id` | integer \| null | Yes |
| `vehicle_arrived_at` | string \| null | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `services_visible_to_customer` | boolean | Yes |
| `inspections_visible_to_customer` | boolean | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `location` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `odometer_in` | integer \| null | No |
| `odometer_out` | integer \| null | No |
| `odometer_unit` | string | No |
| `authorized_at` | string \| null | No |
| `authorized_total_cents` | integer | No |
| `customer_notified` | boolean | No |
| `customer_notified_ready` | boolean | No |
| `ready_for_pickup_at` | string \| null | No |
| `completed_at` | string \| null | No |
| `declined_at` | string \| null | No |
| `decline_reason` | string \| null | No |
| `discount_cents` | integer | No |
| `fees_cents` | integer | No |
| `parts_cents` | integer | No |
| `labor_cents` | integer | No |
| `tires_cents` | integer | No |
| `subcontracts_cents` | integer | No |
| `credit_balance_cents` | integer | No |
| `saved_for_later` | boolean | No |
| `closure_reason` | string \| null | No |
| `closure_reason_notes` | string \| null | No |
| `notes` | string \| null | No |
| `purchase_order_number` | string \| null | No |
| `return_method` | string | No |
| `return_method_notes` | string \| null | No |
| `vehicle_keys_location` | string | No |
| `vehicle_location` | string | No |
| `customer_visit_count` | integer | No |
| `customer_total_spend_cents` | integer | No |
| `average_ticket_cents` | integer | No |
| `activity_total` | integer | No |
| `recent_activities` | array of object | No |
| `services_url` | string | No |
| `payments_url` | string | No |
| `wip_url` | string | No |
| `inspection_url` | string | No |
| `parts_url` | string | No |
| `concerns_url` | string | No |
| `service_history_url` | string | No |
| `declined_services_url` | string | No |
| `activity_url` | string | No |
| `vehicle_history_url` | string | No |
| `appointments_url` | string | No |
| `authorization_logs_url` | string | No |
| `payer_customer` | object | No |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `vin` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
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

`payer_customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
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

### CreateWorkOrderRequest schema {#createworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |
| `vehicle_id` | integer | Yes |

---

### CreateWorkOrdersQuickIntakeRequest schema {#createworkordersquickintakerequest-schema}

| Field | Type | Required |
|---|---|---|
| `customer` | object | Yes |
| `work_order` | object | Yes |
| `vehicle` | object | No |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `phone` | string | Yes |
| `first_name` | string | No |
| `last_name` | string | No |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `concern_description` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `year` | integer | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `license_plate` | string | Yes |
| `license_plate_state` | string | Yes |

---

### UpdateWorkOrderRequest schema {#updateworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `saved_for_later` | boolean | No |
| `vehicle_arrived_at` | string | No |
| `intake_method` | string | No |
| `waiting_for_customer` | boolean | No |
| `work_order_tag_id` | string | No |
| `payer_customer_id` | integer | No |
| `sub_status_type_id` | integer | No |
| `services_visible_to_customer` | boolean | No |

---

### Appointment schema {#appointment-schema}

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `scheduling_status` | string | Yes |
| `appointment_type` | string | Yes |
| `appointment_source` | string \| null | No |
| `intake_method` | string | Yes |
| `all_day` | boolean | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `estimated_duration_minutes` | integer | No |
| `customer_id` | integer \| null | No |
| `vehicle_id` | integer \| null | No |
| `service_advisor_id` | integer | Yes |
| `work_order_id` | integer \| null | Yes |
| `driver_id` | integer \| null | No |
| `marketing_source_id` | integer \| null | No |
| `customer_name` | string \| null | No |
| `customer_email` | string \| null | No |
| `customer_phone` | string \| null | No |
| `customer_concern` | string \| null | No |
| `follow_up_reason` | string \| null | No |
| `year` | any | No |
| `make` | any | No |
| `model` | string \| null | No |
| `submodel` | string \| null | No |
| `vin` | string \| null | No |
| `license_plate` | any | No |
| `customer_confirmed` | boolean | No |
| `confirmation_sent_at` | string \| null | No |
| `reminder_sent_at` | string \| null | No |
| `customer_arrived_at` | string \| null | No |
| `customer_initiated` | boolean | No |
| `rescheduled_from_id` | integer \| null | No |
| `messages_count` | integer | No |
| `reschedules_count` | integer | No |
| `display_name` | string | No |
| `url` | string | Yes |
| `app_url` | string | No |
| `created_at` | string | Yes |
| `updated_at` | string | No |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `service_advisor` | object | No |
| `work_order` | any | No |
| `location` | object | No |
| `latest_reschedule_id` | integer \| null | No |
| `approve_url` | string | No |
| `reject_url` | string | No |
| `cancel_url` | string | No |
| `follow_up_url` | string | No |
| `work_order_url` | string | No |
| `reconcile_vehicle_url` | string | No |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | No |
| `full_name` | string | No |
| `display_name` | string | No |
| `url` | string | No |
| `app_url` | string | No |
| `phones_count` | integer | No |
| `emails_count` | integer | No |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | No |
| `display_name` | string | No |
| `make` | string | No |
| `model` | string | No |
| `year` | integer | No |
| `license_plate` | string | No |
| `url` | string | No |
| `app_url` | string | No |

`service_advisor` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

---

### CompleteWorkOrderRequest schema {#completeworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order_completion_form` | object | Yes |

`work_order_completion_form` — object:
| Field | Type | Required |
|---|---|---|
| `odometer_in` | integer | Yes |
| `odometer_out` | integer | Yes |

---

### UpdateWorkOrdersCourtesyCarAssignmentRequest schema {#updateworkorderscourtesycarassignmentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `action_type` | string | Yes |

---

### CreateWorkOrdersCourtesyCarAssignmentRequest schema {#createworkorderscourtesycarassignmentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `vehicle_id` | integer | Yes |

---

### DeclineWorkOrderRequest schema {#declineworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `closure_reason` | string | Yes |

---

### VoidWorkOrderRequest schema {#voidworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `closure_reason` | string | Yes |

---

### CreateWorkOrdersActivityLogRequest schema {#createworkordersactivitylogrequest-schema}

| Field | Type | Required |
|---|---|---|
| `activity_log` | object | Yes |

`activity_log` — object:
| Field | Type | Required |
|---|---|---|
| `body` | string | Yes |

---

### CreateWorkOrdersAuthorizationDecisionRequest schema {#createworkordersauthorizationdecisionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `service_decision_reasons` | object | Yes |

`service_decision_reasons` — object:
| Field | Type | Required |
|---|---|---|
| `1047559673` | string | Yes |

---

### CreateWorkOrdersAuthorizationRequest schema {#createworkordersauthorizationrequest-schema}

| Field | Type | Required |
|---|---|---|
| `authorization_method` | string | Yes |
| `service_ids` | array of integer | Yes |
| `service_decisions` | object | Yes |

`service_decisions` — object:
| Field | Type | Required |
|---|---|---|
| `1047559673` | string | Yes |

---

### CreateWorkOrdersConcernRequest schema {#createworkordersconcernrequest-schema}

| Field | Type | Required |
|---|---|---|
| `concern` | object | Yes |

`concern` — object:
| Field | Type | Required |
|---|---|---|
| `body` | string | Yes |
| `source` | string | Yes |

---

### DeclineAllWorkOrderConcernsRequest schema {#declineallworkorderconcernsrequest-schema}

| Field | Type | Required |
|---|---|---|
| `decline_reason` | string | Yes |

---

### UpdateWorkOrdersConcernRequest schema {#updateworkordersconcernrequest-schema}

| Field | Type | Required |
|---|---|---|
| `concern` | object | Yes |

`concern` — object:
| Field | Type | Required |
|---|---|---|
| `body` | string | Yes |

---

### CreateWorkOrdersConcernsAddPackageRequest schema {#createworkordersconcernsaddpackagerequest-schema}

| Field | Type | Required |
|---|---|---|
| `package_id` | integer | Yes |

---

### CreateWorkOrdersCreditResolutionRequest schema {#createworkorderscreditresolutionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `resolution` | object | Yes |

`resolution` — object:
| Field | Type | Required |
|---|---|---|
| `action` | string | Yes |

---

### CreateWorkOrdersFeeExclusionRequest schema {#createworkordersfeeexclusionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order_fee_exclusion` | object | Yes |

`work_order_fee_exclusion` — object:
| Field | Type | Required |
|---|---|---|
| `shop_fee_config_id` | integer | Yes |

---

### CreateWorkOrdersPaymentLinkSendRequest schema {#createworkorderspaymentlinksendrequest-schema}

| Field | Type | Required |
|---|---|---|
| `message` | object | Yes |

`message` — object:
| Field | Type | Required |
|---|---|---|
| `channel` | string | Yes |
| `body` | string | Yes |

---

### CreateWorkOrderPaymentRequest schema {#createworkorderpaymentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `payment` | object | Yes |

`payment` — object:
| Field | Type | Required |
|---|---|---|
| `method` | string | Yes |
| `amount_cents` | number | Yes |

---

### CreateWorkOrdersPurchaseOrderRequest schema {#createworkorderspurchaseorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `purchase_order` | object | Yes |

`purchase_order` — object:
| Field | Type | Required |
|---|---|---|
| `vendor_id` | integer | Yes |
| `notes` | string | Yes |
| `line_items` | object | Yes |

`line_items` — object:
| Field | Type | Required |
|---|---|---|
| `0` | object | Yes |

`0` — object:
| Field | Type | Required |
|---|---|---|
| `part_id` | integer | Yes |
| `description` | string | Yes |
| `quantity_ordered` | integer | Yes |
| `unit_cost` | number | Yes |

---

### PurchaseOrder schema {#purchaseorder-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `po_number` | integer | Yes |
| `receiving_status` | string | Yes |
| `order_method` | string | Yes |
| `payment_method` | string | Yes |
| `fulfillment_method` | string | Yes |
| `tracking_number` | string \| null | Yes |
| `vendor_invoice_number` | string \| null | Yes |
| `vendor_invoice_received_at` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `freight_cost_cents` | integer | Yes |
| `freight_cost_currency` | string | Yes |
| `subtotal_cents` | integer | Yes |
| `total_cents` | integer | Yes |
| `core_charges_cents` | integer | Yes |
| `line_items_count` | integer | Yes |
| `ordered_at` | string | Yes |
| `received_at` | string \| null | Yes |
| `payment_due_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `vendor` | object | Yes |
| `work_order` | object | No |
| `location` | object | Yes |
| `creator` | object | No |
| `line_items` | array of object | No |

`vendor` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `number` | integer | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`creator` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

---

### UpdateWorkOrdersPurchaseOrderRequest schema {#updateworkorderspurchaseorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `purchase_order` | object | Yes |

`purchase_order` — object:
| Field | Type | Required |
|---|---|---|
| `notes` | string | Yes |

---

### CreateWorkOrdersPurchaseOrdersReceiveRequest schema {#createworkorderspurchaseordersreceiverequest-schema}

| Field | Type | Required |
|---|---|---|
| `lines` | object | Yes |

`lines` — object:
| Field | Type | Required |
|---|---|---|
| `20915845` | object | Yes |

`20915845` — object:
| Field | Type | Required |
|---|---|---|
| `selected` | string | Yes |
| `quantity_received` | string | Yes |

---

### CreateWorkOrdersPurchaseOrdersReturnRequest schema {#createworkorderspurchaseordersreturnrequest-schema}

| Field | Type | Required |
|---|---|---|
| `lines` | object | Yes |
| `return_order` | object | Yes |

`lines` — object:
| Field | Type | Required |
|---|---|---|
| `20915845` | object | Yes |

`20915845` — object:
| Field | Type | Required |
|---|---|---|
| `selected` | string | Yes |
| `quantity` | string | Yes |
| `purchase_order_line_item_id` | integer | Yes |
| `return_action` | string | Yes |
| `reason` | string | Yes |

`return_order` — object:
| Field | Type | Required |
|---|---|---|
| `credit_method` | string | Yes |

---

### CreateWorkOrdersRefundRequest schema {#createworkordersrefundrequest-schema}

| Field | Type | Required |
|---|---|---|
| `refund` | object | Yes |

`refund` — object:
| Field | Type | Required |
|---|---|---|
| `payment_id` | integer | Yes |
| `amount` | string | Yes |
| `reason` | string | Yes |

---

### CreateWorkOrdersServiceRequest schema {#createworkordersservicerequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order_service` | object | Yes |
| `package_id` | integer | No |

`work_order_service` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `service_type` | string | No |

---

### UpdateWorkOrdersServicesReorderRequest schema {#updateworkordersservicesreorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `service_ids` | array of integer | Yes |

---

### UpdateWorkOrdersServiceRequest schema {#updateworkordersservicerequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order_service` | object | Yes |

`work_order_service` — object:
| Field | Type | Required |
|---|---|---|
| `technician_id` | integer | No |
| `position` | integer | No |
| `name` | string | No |
| `labor_tax_enabled` | boolean | No |
| `pricing_mode` | string | No |

---

### UpdateWorkOrdersServicesAdjustTimeRequest schema {#updateworkordersservicesadjusttimerequest-schema}

| Field | Type | Required |
|---|---|---|
| `hours` | integer | Yes |
| `minutes` | integer | Yes |

---

### CreateWorkOrdersServicesPackageRequest schema {#createworkordersservicespackagerequest-schema}

| Field | Type | Required |
|---|---|---|
| `package_id` | integer | Yes |

---

### UpdateWorkOrdersServicesToggleLaborCompletionRequest schema {#updateworkordersservicestogglelaborcompletionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `line_item_id` | integer | Yes |

---

### UpdateWorkOrdersServicesUpdateCategoryRequest schema {#updateworkordersservicesupdatecategoryrequest-schema}

| Field | Type | Required |
|---|---|---|
| `category_id` | integer | Yes |

---

### CreateWorkOrdersServicesCommentRequest schema {#createworkordersservicescommentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `service_comment` | object | Yes |

`service_comment` — object:
| Field | Type | Required |
|---|---|---|
| `body` | string | Yes |

---

### CreateWorkOrdersServicesExtractionRequest schema {#createworkordersservicesextractionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `text` | string | Yes |
| `extraction_id` | string | Yes |

---

### CreateWorkOrdersServicesLineItemRequest schema {#createworkordersserviceslineitemrequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order_line_item` | object | Yes |

`work_order_line_item` — object:
| Field | Type | Required |
|---|---|---|
| `item_type` | string | Yes |
| `description` | string | Yes |
| `hours` | number | No |
| `labor_rate_id` | integer | No |
| `total` | string | No |
| `unit_price` | string | No |
| `quantity` | integer | No |

---

### UpdateWorkOrdersServicesLineItemRequest schema {#updateworkordersserviceslineitemrequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order_line_item` | object | Yes |

`work_order_line_item` — object:
| Field | Type | Required |
|---|---|---|
| `part_status` | string | No |
| `description` | string | No |

---

### CreateWorkOrdersServicesSubletOrderRequest schema {#createworkordersservicessubletorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `sublet_order` | object | Yes |

`sublet_order` — object:
| Field | Type | Required |
|---|---|---|
| `title` | string | Yes |

---

### UpdateWorkOrdersServicesSubletOrderRequest schema {#updateworkordersservicessubletorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `sublet_order` | object | Yes |

`sublet_order` — object:
| Field | Type | Required |
|---|---|---|
| `title` | string | Yes |

---

### SubletOrder schema {#subletorder-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `sublet_number` | integer | Yes |
| `title` | string | Yes |
| `payment_status` | string | Yes |
| `payment_method` | string \| null | Yes |
| `total_cents` | integer | Yes |
| `total_cost_cents` | integer | Yes |
| `margin_cents` | integer | Yes |
| `margin_percentage` | integer | Yes |
| `sent_to_ap` | boolean | Yes |
| `vendor_paid_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `vendor` | object | No |
| `work_order` | object | No |
| `work_order_service` | object | Yes |
| `location` | object | Yes |
| `fulfillment_status` | string | No |
| `notes` | string \| null | No |
| `vendor_invoice_number` | any | No |

`vendor` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `number` | integer | Yes |
| `url` | string | Yes |

`work_order_service` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

---

### CreateWorkOrdersTireStorageRequest schema {#createworkorderstirestoragerequest-schema}

| Field | Type | Required |
|---|---|---|
| `tire_storage_slot` | object | Yes |

`tire_storage_slot` — object:
| Field | Type | Required |
|---|---|---|
| `slot_label` | string | Yes |
| `season` | string | Yes |
| `tire_set_description` | string | Yes |
| `notes` | string | Yes |

---

### UpdateWorkOrdersTireStorageRequest schema {#updateworkorderstirestoragerequest-schema}

| Field | Type | Required |
|---|---|---|
| `tire_storage_slot` | object | Yes |

`tire_storage_slot` — object:
| Field | Type | Required |
|---|---|---|
| `tire_set_description` | string | Yes |

---

### CreateWorkOrdersVoidRequest schema {#createworkordersvoidrequest-schema}

| Field | Type | Required |
|---|---|---|
| `payment_id` | integer | Yes |

