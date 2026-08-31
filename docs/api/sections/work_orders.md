# Work Orders

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List work orders

```
GET /work_orders
```

List all work orders, paginated via the Link header.

**Response 200** — array of [WorkOrder](#workorder-schema)

**Example**

```json
[
  {
    "type": "WorkOrder",
    "id": 1,
    "work_order_number": 1001,
    "status": "in_progress",
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
    "work_order_services_count": 0,
    "inspection_reports_count": 0,
    "customer": {
      "id": 1,
      "full_name": "Jane Doe",
      "url": "https://app.wenmarpro.com/customers/1.json"
    },
    "vehicle": {
      "id": 1,
      "make": "Toyota",
      "model": "Camry",
      "year": 2020,
      "vin": "ABC123",
      "url": "https://app.wenmarpro.com/vehicles/1.json"
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
    "recent_activities": []
  },
  {
    "type": "WorkOrder",
    "id": 2,
    "work_order_number": 1002,
    "status": "open",
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
    "work_order_services_count": 0,
    "inspection_reports_count": 0,
    "customer": {
      "id": 2,
      "full_name": "John Smith",
      "url": "https://app.wenmarpro.com/customers/2.json"
    },
    "vehicle": {
      "id": 2,
      "make": "Honda",
      "model": "Civic",
      "year": 2018,
      "vin": "XYZ789",
      "url": "https://app.wenmarpro.com/vehicles/2.json"
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
    "recent_activities": []
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
  "status": "in_progress",
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
  "work_order_services_count": 2,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json"
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
  "vehicle_location": "bay_2"
}
```

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders.json
```

## Delete work order

```
DELETE /work_orders/{id}
```

Delete a work order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/<id>.json
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

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "status": "in_progress",
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
  "work_order_services_count": 2,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json"
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
  "vehicle_location": "bay_2"
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

**Response 422** — [Error](#error-schema) error envelope

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "status": "in_progress",
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
  "work_order_services_count": 2,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json"
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
  "vehicle_location": "bay_2"
}
```

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/<id>.json
```

## Create work order authorization

```
POST /work_orders/{work_order_id}/authorization
```

Create a work order authorization.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
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
| `authorized_at` | string | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `vehicle_arrived_at` | any | Yes |
| `ready_for_pickup_at` | any | Yes |
| `completed_at` | any | Yes |
| `declined_at` | string \| null | Yes |
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
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/authorization.json
```

## Update work order authorization decisions

```
POST /work_orders/{work_order_id}/authorization/update_decisions
```

Update a work order authorization decisions by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
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
| `authorized_at` | string | Yes |
| `authorized_total_cents` | integer | Yes |
| `customer_notified` | boolean | Yes |
| `customer_notified_ready` | boolean | Yes |
| `vehicle_arrived_at` | any | Yes |
| `ready_for_pickup_at` | any | Yes |
| `completed_at` | any | Yes |
| `declined_at` | string | Yes |
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
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/authorization/update_decisions.json
```

## Show work order declined services

```
GET /work_orders/{work_order_id}/declined_services
```

Show a work order declined services by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `name` | string | Yes |
| `declined_at` | string | Yes |
| `work_order_number` | integer | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/declined_services.json
```

## Show work order estimate

```
GET /work_orders/{work_order_id}/estimate
```

Show a work order estimate by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
| `assigned_technician_id` | any | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `services` | array of object | Yes |

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

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "status": "in_progress",
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
  "work_order_services_count": 1,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json"
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
  ]
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/estimate.json
```

## Show work order inspection

```
GET /work_orders/{work_order_id}/inspection
```

Show a work order inspection by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
| `assigned_technician_id` | any | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `inspection_reports` | array of object | Yes |

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

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "status": "in_progress",
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
  "work_order_services_count": 0,
  "inspection_reports_count": 1,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json"
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
  ]
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/inspection.json
```

## Close work order

```
PATCH /work_orders/{work_order_id}/lifecycle/close
```

Close

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `closed_at` | string | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | any | Yes |
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
| `vehicle_arrived_at` | any | Yes |
| `ready_for_pickup_at` | any | Yes |
| `completed_at` | any | Yes |
| `declined_at` | string \| null | Yes |
| `decline_reason` | any | Yes |
| `discount_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `labor_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `credit_balance_cents` | integer | Yes |
| `saved_for_later` | boolean | Yes |
| `closure_reason` | string | Yes |
| `closure_reason_notes` | any | Yes |
| `notes` | any | Yes |
| `purchase_order_number` | any | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | any | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/lifecycle/close.json
```

## Close work order as paid

```
PATCH /work_orders/{work_order_id}/lifecycle/close_as_paid
```

Close as paid

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `closed_at` | string | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | any | Yes |
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
| `vehicle_arrived_at` | any | Yes |
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
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/lifecycle/close_as_paid.json
```

## Decline all work order services

```
PATCH /work_orders/{work_order_id}/lifecycle/decline_all
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
| `status` | string | Yes |
| `intake_method` | string | Yes |
| `scheduled_for` | any | Yes |
| `authorized` | boolean | Yes |
| `paid` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `closed_at` | any | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | any | Yes |
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
| `vehicle_arrived_at` | any | Yes |
| `ready_for_pickup_at` | any | Yes |
| `completed_at` | any | Yes |
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
| `closure_reason` | any | Yes |
| `closure_reason_notes` | any | Yes |
| `notes` | any | Yes |
| `purchase_order_number` | any | Yes |
| `return_method` | string | Yes |
| `return_method_notes` | any | Yes |
| `vehicle_keys_location` | string | Yes |
| `vehicle_location` | string | Yes |
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/lifecycle/decline_all.json
```

## Reopen work order

```
PATCH /work_orders/{work_order_id}/lifecycle/reopen
```

Reopen

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
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
| `vehicle_arrived_at` | any | Yes |
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
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/lifecycle/reopen.json
```

## Return work order to board

```
PATCH /work_orders/{work_order_id}/lifecycle/return_to_board
```

Return to board

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
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
| `vehicle_arrived_at` | any | Yes |
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
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of any | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/lifecycle/return_to_board.json
```

## Save work order for later

```
PATCH /work_orders/{work_order_id}/lifecycle/save_for_later
```

Save for later

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
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
| `vehicle_arrived_at` | any | Yes |
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
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of any | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/lifecycle/save_for_later.json
```

## Stage transition work order

```
PATCH /work_orders/{work_order_id}/lifecycle/stage_transition
```

Stage transition

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `closed_at` | string | Yes |
| `location_id` | integer | Yes |
| `service_advisor_id` | any | Yes |
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
| `vehicle_arrived_at` | any | Yes |
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
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of object | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/lifecycle/stage_transition.json
```

## Show work order parts

```
GET /work_orders/{work_order_id}/parts
```

Show a work order parts by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
| `assigned_technician_id` | any | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `services` | array of object | Yes |

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

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "status": "in_progress",
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
  "work_order_services_count": 1,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json"
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
  ]
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/parts.json
```

## Show work order payments

```
GET /work_orders/{work_order_id}/payments
```

Show a work order payments by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
| `assigned_technician_id` | any | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `payments` | array of object | Yes |

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

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "status": "completed",
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
  "work_order_services_count": 1,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json"
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
  ]
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

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `amount_cents` | integer | Yes |
| `method` | string | Yes |
| `processor_status` | any | Yes |
| `is_refund` | boolean | Yes |
| `processed_at` | string | Yes |
| `reference` | any | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `work_order_id` | integer | Yes |
| `customer_id` | integer | Yes |
| `work_order` | object | Yes |
| `customer` | object | Yes |
| `processed_by` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

`processed_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/work_orders/{work_order_id}/payments.json
```

## Show work order service history

```
GET /work_orders/{work_order_id}/service_history
```

Show a work order service history by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/work_orders/{work_order_id}/service_history.json
```

## Show work order WIP

```
GET /work_orders/{work_order_id}/wip
```

Show a work order wip by ID.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | Yes |

**Response 200**

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
| `service_advisor_id` | any | Yes |
| `assigned_technician_id` | any | Yes |
| `work_order_services_count` | integer | Yes |
| `inspection_reports_count` | integer | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `services` | array of object | Yes |

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

**Example**

```json
{
  "type": "WorkOrder",
  "id": 1,
  "work_order_number": 1001,
  "status": "in_progress",
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
  "work_order_services_count": 1,
  "inspection_reports_count": 0,
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "vehicle": {
    "id": 1,
    "make": "Toyota",
    "model": "Camry",
    "year": 2020,
    "vin": "ABC123",
    "url": "https://app.wenmarpro.com/vehicles/1.json"
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
  ]
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
| `services_url` | string | Yes |
| `payments_url` | string | Yes |
| `wip_url` | string | Yes |
| `inspection_url` | string | Yes |
| `parts_url` | string | Yes |
| `concerns_url` | string | Yes |
| `customer_visit_count` | integer | Yes |
| `customer_total_spend_cents` | integer | Yes |
| `average_ticket_cents` | integer | Yes |
| `activity_total` | integer | Yes |
| `recent_activities` | array of any | Yes |
| `service_history_url` | string | Yes |
| `declined_services_url` | string | Yes |

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

### UpdateWorkOrderRequest schema {#updateworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `work_order` | object | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `intake_method` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### CreateWorkOrderAuthorizationRequest schema {#createworkorderauthorizationrequest-schema}

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

### UpdateWorkOrderAuthorizationDecisionsRequest schema {#updateworkorderauthorizationdecisionsrequest-schema}

| Field | Type | Required |
|---|---|---|
| `service_decision_reasons` | object | Yes |

`service_decision_reasons` — object:
| Field | Type | Required |
|---|---|---|
| `1047559673` | string | Yes |

---

### CloseWorkOrderRequest schema {#closeworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `closure_type` | string | Yes |
| `closure_reason` | string | Yes |

---

### DeclineAllWorkOrderServicesRequest schema {#declineallworkorderservicesrequest-schema}

| Field | Type | Required |
|---|---|---|
| `decline_reason` | string | Yes |

---

### StageTransitionWorkOrderRequest schema {#stagetransitionworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `target` | string | Yes |

---

### CreateWorkOrderPaymentRequest schema {#createworkorderpaymentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `payment` | object | Yes |

`payment` — object:
| Field | Type | Required |
|---|---|---|
| `method` | string | Yes |
| `amount_cents` | string | Yes |

