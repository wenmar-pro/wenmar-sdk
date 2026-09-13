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

**Response 200** — array of [Appointment](#appointment-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/appointments.json
```

## Create appointment

```
POST /appointments
```

Create a appointment.

**Response 201** — [Appointment](#appointment-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

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

**Response 200** — [Appointment](#appointment-schema)

**Response 304** — no content.

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

**Response 200** — [Appointment](#appointment-schema)

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

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `status` | string | Yes |
| `appointment_type` | string | Yes |
| `appointment_source` | any | Yes |
| `intake_method` | string | Yes |
| `all_day` | boolean | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `estimated_duration_minutes` | integer | Yes |
| `customer_id` | integer \| null | Yes |
| `vehicle_id` | integer \| null | Yes |
| `service_advisor_id` | integer | Yes |
| `work_order_id` | integer \| null | Yes |
| `driver_id` | integer \| null | Yes |
| `marketing_source_id` | integer \| null | Yes |
| `customer_name` | string | Yes |
| `customer_email` | string | Yes |
| `customer_phone` | string | Yes |
| `customer_concern` | any | Yes |
| `follow_up_reason` | string \| null | Yes |
| `year` | any | Yes |
| `make` | any | Yes |
| `model` | string \| null | Yes |
| `submodel` | string \| null | Yes |
| `vin` | string \| null | Yes |
| `license_plate` | any | Yes |
| `customer_confirmed` | boolean | Yes |
| `confirmation_sent_at` | string \| null | Yes |
| `reminder_sent_at` | string \| null | Yes |
| `customer_arrived_at` | string \| null | Yes |
| `customer_initiated` | boolean | Yes |
| `rescheduled_from_id` | integer \| null | Yes |
| `messages_count` | integer | Yes |
| `reschedules_count` | integer | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `service_advisor` | object | Yes |
| `work_order` | any | Yes |
| `location` | object | Yes |
| `latest_reschedule_id` | integer \| null | Yes |
| `approve_url` | string | Yes |
| `reject_url` | string | Yes |
| `cancel_url` | string | Yes |
| `follow_up_url` | string | Yes |
| `work_order_url` | string | Yes |
| `reconcile_vehicle_url` | string | Yes |

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

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `status` | string | Yes |
| `appointment_type` | string | Yes |
| `appointment_source` | any | Yes |
| `intake_method` | string | Yes |
| `all_day` | boolean | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `estimated_duration_minutes` | integer | Yes |
| `customer_id` | integer \| null | Yes |
| `vehicle_id` | integer \| null | Yes |
| `service_advisor_id` | integer | Yes |
| `work_order_id` | integer \| null | Yes |
| `driver_id` | integer \| null | Yes |
| `marketing_source_id` | integer \| null | Yes |
| `customer_name` | string | Yes |
| `customer_email` | string | Yes |
| `customer_phone` | string | Yes |
| `customer_concern` | any | Yes |
| `follow_up_reason` | string \| null | Yes |
| `year` | any | Yes |
| `make` | any | Yes |
| `model` | string \| null | Yes |
| `submodel` | string \| null | Yes |
| `vin` | string \| null | Yes |
| `license_plate` | any | Yes |
| `customer_confirmed` | boolean | Yes |
| `confirmation_sent_at` | string \| null | Yes |
| `reminder_sent_at` | string \| null | Yes |
| `customer_arrived_at` | string \| null | Yes |
| `customer_initiated` | boolean | Yes |
| `rescheduled_from_id` | integer \| null | Yes |
| `messages_count` | integer | Yes |
| `reschedules_count` | integer | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `service_advisor` | object | Yes |
| `work_order` | any | Yes |
| `location` | object | Yes |
| `latest_reschedule_id` | integer \| null | Yes |
| `approve_url` | string | Yes |
| `reject_url` | string | Yes |
| `cancel_url` | string | Yes |
| `follow_up_url` | string | Yes |
| `work_order_url` | string | Yes |
| `reconcile_vehicle_url` | string | Yes |

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
| `type` | string | Yes |
| `id` | integer | Yes |
| `status` | string | Yes |
| `appointment_type` | string | Yes |
| `appointment_source` | string | Yes |
| `intake_method` | string | Yes |
| `all_day` | boolean | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `estimated_duration_minutes` | integer | Yes |
| `customer_id` | integer | Yes |
| `vehicle_id` | integer | Yes |
| `service_advisor_id` | integer \| null | Yes |
| `work_order_id` | integer | Yes |
| `driver_id` | integer \| null | Yes |
| `marketing_source_id` | integer \| null | Yes |
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
| `customer_confirmed` | boolean | Yes |
| `confirmation_sent_at` | string \| null | Yes |
| `reminder_sent_at` | string \| null | Yes |
| `customer_arrived_at` | string \| null | Yes |
| `customer_initiated` | boolean | Yes |
| `rescheduled_from_id` | integer \| null | Yes |
| `messages_count` | integer | Yes |
| `reschedules_count` | integer | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `service_advisor` | any | Yes |
| `work_order` | object | Yes |
| `location` | object | Yes |
| `latest_reschedule_id` | integer \| null | Yes |
| `approve_url` | string | Yes |
| `reject_url` | string | Yes |
| `cancel_url` | string | Yes |
| `follow_up_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `phones_count` | integer | Yes |
| `emails_count` | integer | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`work_order` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `status` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

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

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `status` | string | Yes |
| `appointment_type` | string | Yes |
| `appointment_source` | any | Yes |
| `intake_method` | string | Yes |
| `all_day` | boolean | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `estimated_duration_minutes` | integer | Yes |
| `customer_id` | integer \| null | Yes |
| `vehicle_id` | integer \| null | Yes |
| `service_advisor_id` | integer | Yes |
| `work_order_id` | integer \| null | Yes |
| `driver_id` | integer \| null | Yes |
| `marketing_source_id` | integer \| null | Yes |
| `customer_name` | string | Yes |
| `customer_email` | string | Yes |
| `customer_phone` | string | Yes |
| `customer_concern` | any | Yes |
| `follow_up_reason` | string \| null | Yes |
| `year` | any | Yes |
| `make` | any | Yes |
| `model` | string \| null | Yes |
| `submodel` | string \| null | Yes |
| `vin` | string \| null | Yes |
| `license_plate` | any | Yes |
| `customer_confirmed` | boolean | Yes |
| `confirmation_sent_at` | string \| null | Yes |
| `reminder_sent_at` | string \| null | Yes |
| `customer_arrived_at` | string \| null | Yes |
| `customer_initiated` | boolean | Yes |
| `rescheduled_from_id` | integer \| null | Yes |
| `messages_count` | integer | Yes |
| `reschedules_count` | integer | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `service_advisor` | object | Yes |
| `work_order` | any | Yes |
| `location` | object | Yes |
| `latest_reschedule_id` | integer \| null | Yes |
| `approve_url` | string | Yes |
| `reject_url` | string | Yes |
| `cancel_url` | string | Yes |
| `follow_up_url` | string | Yes |
| `work_order_url` | string | Yes |
| `reconcile_vehicle_url` | string | Yes |

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

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `status` | string | Yes |
| `appointment_type` | string | Yes |
| `appointment_source` | any | Yes |
| `intake_method` | string | Yes |
| `all_day` | boolean | Yes |
| `starts_at` | string | Yes |
| `ends_at` | string | Yes |
| `estimated_duration_minutes` | integer | Yes |
| `customer_id` | integer | Yes |
| `vehicle_id` | integer | Yes |
| `service_advisor_id` | integer | Yes |
| `work_order_id` | integer \| null | Yes |
| `driver_id` | integer \| null | Yes |
| `marketing_source_id` | integer \| null | Yes |
| `customer_name` | string | Yes |
| `customer_email` | string | Yes |
| `customer_phone` | string | Yes |
| `customer_concern` | any | Yes |
| `follow_up_reason` | string \| null | Yes |
| `year` | any | Yes |
| `make` | any | Yes |
| `model` | string \| null | Yes |
| `submodel` | string \| null | Yes |
| `vin` | string \| null | Yes |
| `license_plate` | any | Yes |
| `customer_confirmed` | boolean | Yes |
| `confirmation_sent_at` | string \| null | Yes |
| `reminder_sent_at` | string \| null | Yes |
| `customer_arrived_at` | string \| null | Yes |
| `customer_initiated` | boolean | Yes |
| `rescheduled_from_id` | integer \| null | Yes |
| `messages_count` | integer | Yes |
| `reschedules_count` | integer | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `customer` | object | Yes |
| `vehicle` | object | Yes |
| `service_advisor` | object | Yes |
| `work_order` | any | Yes |
| `location` | object | Yes |
| `latest_reschedule_id` | integer \| null | Yes |
| `approve_url` | string | Yes |
| `reject_url` | string | Yes |
| `cancel_url` | string | Yes |
| `follow_up_url` | string | Yes |
| `work_order_url` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `display_name` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `phones_count` | integer | Yes |
| `emails_count` | integer | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `display_name` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |
| `year` | integer | Yes |
| `license_plate` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

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

### Appointment schema {#appointment-schema}

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `status` | string | Yes |
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

### CreateAppointmentRequest schema {#createappointmentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `appointment` | object | Yes |

`appointment` — object:
| Field | Type | Required |
|---|---|---|
| `starts_at` | string | Yes |
| `customer_id` | integer | Yes |
| `vehicle_id` | integer | Yes |
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

### UpdateAppointmentRequest schema {#updateappointmentrequest-schema}

| Field | Type | Required |
|---|---|---|
| `appointment` | object | Yes |

`appointment` — object:
| Field | Type | Required |
|---|---|---|
| `customer_concern` | string | Yes |

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

### CreateAppointmentsVehicleReconciliationRequest schema {#createappointmentsvehiclereconciliationrequest-schema}

| Field | Type | Required |
|---|---|---|
| `vehicle_action` | string | Yes |
| `vehicle_id` | integer | Yes |

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

