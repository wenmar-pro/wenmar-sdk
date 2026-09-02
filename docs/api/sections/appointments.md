# Appointments

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List appointments

```
GET /appointments
```

List all appointments, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `per_page` | integer | No |
| `q` | string | No |
| `status` | string | No |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/appointments.json
```

## Create appointment

```
POST /appointments
```

Create a appointment.

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/appointments.json
```

## List appointments available slots

```
GET /appointments/available_slots
```

List all appointments available slots, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `date` | string | No |
| `duration_minutes` | integer | No |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `time` | string | Yes |
| `booked` | integer | Yes |
| `capacity` | integer | Yes |
| `available` | boolean | Yes |
| `blocked` | boolean | Yes |
| `slots_needed` | integer | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/appointments/available_slots.json
```

## Delete appointment

```
DELETE /appointments/{id}
```

Delete a appointment by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/appointments/<id>.json
```

## Show appointment

```
GET /appointments/{id}
```

Show a appointment by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/appointments/<id>.json
```

## Update appointment

```
PATCH /appointments/{id}
```

Update a appointment by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/appointments/<id>.json
```

## Create appointments approval

```
POST /appointments/{id}/approvals
```

Create a appointments approval.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/appointments/<id>.json
```

## Create appointments cancellation

```
POST /appointments/{id}/cancellations
```

Create a appointments cancellation.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/appointments/<id>.json
```

## Create appointments follow up

```
POST /appointments/{id}/follow_ups
```

Create a appointments follow up.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `status` | string | Yes |
| `appointment_type` | string | Yes |
| `appointment_source` | string | Yes |
| `intake_method` | string | Yes |
| `all_day` | boolean | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `estimated_duration_minutes` | integer | Yes |
| `customer_name` | string | Yes |
| `customer_email` | string | Yes |
| `customer_phone` | string | Yes |
| `customer_concern` | any | Yes |
| `follow_up_reason` | string | Yes |
| `year` | integer | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `submodel` | string \| null | Yes |
| `vin` | string | Yes |
| `license_plate` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `service_advisor` | any | Yes |
| `work_order` | object | Yes |
| `location` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `customer_confirmed` | boolean | Yes |
| `confirmation_sent_at` | string \| null | Yes |
| `reminder_sent_at` | string \| null | Yes |
| `customer_arrived_at` | string \| null | Yes |
| `customer_initiated` | boolean | Yes |
| `rescheduled_from_id` | integer \| null | Yes |
| `latest_reschedule_id` | integer \| null | Yes |
| `messages_count` | integer | Yes |
| `reschedules_count` | integer | Yes |
| `approve_url` | string | Yes |
| `reject_url` | string | Yes |
| `cancel_url` | string | Yes |
| `follow_up_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `phones_count` | integer | Yes |
| `emails_count` | integer | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `status` | string | Yes |
| `url` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/appointments/<id>.json
```

## Create appointments rejection

```
POST /appointments/{id}/rejections
```

Create a appointments rejection.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/appointments/<id>.json
```

## Create appointments vehicle reconciliation

```
POST /appointments/{id}/vehicle_reconciliations
```

Create a appointments vehicle reconciliation.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/appointments/<id>.json
```

## Create appointments work order

```
POST /appointments/{id}/work_orders
```

Create a appointments work order.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 201** — [WorkOrder](#workorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/appointments/<id>.json
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

### CreateAppointmentsFollowUpRequest schema {#createappointmentsfollowuprequest-schema}

| Field | Type | Required |
|---|---|---|
| `appointment` | object | Yes |

`appointment` — object:
| Field | Type | Required |
|---|---|---|
| `starts_at` | string | Yes |
| `follow_up_reason` | string | Yes |

---

### WorkOrder schema {#workorder-schema}

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `work_order_number` | integer | Yes |
| `status` | string | Yes |
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
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
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
| `payer_customer` | object | No |

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

`payer_customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

