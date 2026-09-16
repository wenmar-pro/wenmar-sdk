# Tech Assignments

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create work order tech assignment

```
POST /tech_assignments
```

Create a work order tech assignment.

| Param | Type | Required |
|---|---|---|
| `work_order_id` | integer | No |

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
| `assigned_technician_id` | integer | Yes |
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

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/tech_assignments.json
```

---

### CreateWorkOrderTechAssignmentRequest schema {#createworkordertechassignmentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `technician_id` | integer | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

