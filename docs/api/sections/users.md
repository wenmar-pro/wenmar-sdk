# Users

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List users

```
GET /users
```

List all users, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `page` | integer | No |
| `per_page` | integer | No |

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/users.json
```

## Create user

```
POST /users
```

Create a user.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users.json
```

## List permission groups

```
GET /users/permission_groups
```

List all permission groups, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `role` | string \| null | Yes |
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/users/permission_groups.json
```

## Create permission group

```
POST /users/permission_groups
```

Create a permission group.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `role` | string \| null | Yes |
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/permission_groups.json
```

## Update permission group

```
PATCH /users/permission_groups/{id}
```

Update a permission group by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `role` | string \| null | Yes |
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/permission_groups/<id>.json
```

## Delete user

```
DELETE /users/{id}
```

Delete a user by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/users/<id>.json
```

## Show user

```
GET /users/{id}
```

Show a user by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/users/<id>.json
```

## Update user

```
PATCH /users/{id}
```

Update a user by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/<id>.json
```

## Create users disable

```
POST /users/{id}/disable
```

Create a users disable.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/<id>.json
```

## Create users enable

```
POST /users/{id}/enable
```

Create a users enable.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/<id>.json
```

## List users qr code

```
GET /users/{id}/qr_code
```

List all users qr code, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |
| `qr_card` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

`qr_card` — object:
| Field | Type | Required |
|---|---|---|
| `staff_name` | string | Yes |
| `valid_from` | string | Yes |
| `valid_until` | string | Yes |
| `instructions` | string | Yes |
| `security_notice` | string | Yes |
| `svg` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/users/<id>.json
```

## Create users reset pin

```
POST /users/{id}/reset_pin
```

Create a users reset pin.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |
| `new_pin` | string | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/<id>.json
```

## Create users send confirmation

```
POST /users/{id}/send_confirmation
```

Create a users send confirmation.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/<id>.json
```

## Create users send password reset

```
POST /users/{id}/send_password_reset
```

Create a users send password reset.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/<id>.json
```

## Create users unlock

```
POST /users/{id}/unlock
```

Create a users unlock.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `full_name` | string | Yes |
| `email` | string \| null | Yes |
| `initials` | string | Yes |
| `role` | string | Yes |
| `disabled` | boolean | Yes |
| `disabled_at` | string \| null | Yes |
| `locked` | boolean | Yes |
| `confirmed` | boolean | Yes |
| `hourly_cost_cents` | integer | Yes |
| `hourly_cost_currency` | string | Yes |
| `certification_number` | any | Yes |
| `certification_label` | any | Yes |
| `mfa_enabled` | boolean | Yes |
| `mfa_required` | boolean | Yes |
| `qr_token_generated_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `locations` | array of object | Yes |
| `capabilities` | object | Yes |

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/users/<id>.json
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

### CreateUserRequest schema {#createuserrequest-schema}

| Field | Type | Required |
|---|---|---|
| `user` | object | Yes |

`user` — object:
| Field | Type | Required |
|---|---|---|
| `full_name` | string | Yes |
| `email` | string | Yes |
| `role` | string | Yes |
| `home_location_id` | integer | Yes |
| `location_ids` | array of integer | Yes |
| `can_perform_work` | boolean | Yes |

---

### CreatePermissionGroupRequest schema {#createpermissiongrouprequest-schema}

| Field | Type | Required |
|---|---|---|
| `permission_group` | object | Yes |

`permission_group` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `can_perform_work` | boolean | Yes |

---

### UpdatePermissionGroupRequest schema {#updatepermissiongrouprequest-schema}

| Field | Type | Required |
|---|---|---|
| `permission_group` | object | Yes |

`permission_group` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### UpdateUserRequest schema {#updateuserrequest-schema}

| Field | Type | Required |
|---|---|---|
| `user` | object | Yes |

`user` — object:
| Field | Type | Required |
|---|---|---|
| `full_name` | string | Yes |

