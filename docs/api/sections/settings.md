# Settings

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List settings account

```
GET /settings/account
```

List all settings account, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `slug` | string | Yes |
| `billing_email` | string | Yes |
| `website` | string \| null | Yes |
| `business_type` | string \| null | Yes |
| `tax_id` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/account.json
```

## Update settings account

```
PATCH /settings/account
```

Update a settings account by ID.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `slug` | string | Yes |
| `billing_email` | string | Yes |
| `website` | string \| null | Yes |
| `business_type` | string \| null | Yes |
| `tax_id` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/settings/account.json
```

## List settings billing

```
GET /settings/billing
```

List all settings billing, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `billing_email` | string | Yes |
| `subscription_status` | string | Yes |
| `next_billing_date` | string | Yes |
| `work_orders_this_month` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/billing.json
```

## List settings cash drawer

```
GET /settings/cash_drawer
```

List all settings cash drawer, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `cash_drawer_enabled` | boolean | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `expenses_enabled` | boolean | Yes |
| `tire_management_enabled` | boolean | Yes |
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/cash_drawer.json
```

## List settings close requirements

```
GET /settings/close_requirements
```

List all settings close requirements, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `cash_drawer_enabled` | boolean | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `expenses_enabled` | boolean | Yes |
| `tire_management_enabled` | boolean | Yes |
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/close_requirements.json
```

## List settings documents

```
GET /settings/documents
```

List all settings documents, paginated via the Link header.

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

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/documents.json
```

## List settings driveon

```
GET /settings/driveon
```

List all settings driveon, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `cash_drawer_enabled` | boolean | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `expenses_enabled` | boolean | Yes |
| `tire_management_enabled` | boolean | Yes |
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/driveon.json
```

## List settings expenses

```
GET /settings/expenses
```

List all settings expenses, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `cash_drawer_enabled` | boolean | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `expenses_enabled` | boolean | Yes |
| `tire_management_enabled` | boolean | Yes |
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/expenses.json
```

## List settings labor templates

```
GET /settings/labor_templates
```

List all settings labor templates, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `description` | string | Yes |
| `default_hours` | string | Yes |
| `usage_count` | integer | Yes |
| `last_used_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/labor_templates.json
```

## Delete settings labor template

```
DELETE /settings/labor_templates/{id}
```

Delete a settings labor template by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/labor_templates/<id>.json
```

## Update settings labor template

```
PATCH /settings/labor_templates/{id}
```

Update a settings labor template by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `description` | string | Yes |
| `default_hours` | string | Yes |
| `usage_count` | integer | Yes |
| `last_used_at` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/settings/labor_templates/<id>.json
```

## List settings lead source requirements

```
GET /settings/lead_source_requirements
```

List all settings lead source requirements, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `cash_drawer_enabled` | boolean | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `expenses_enabled` | boolean | Yes |
| `tire_management_enabled` | boolean | Yes |
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/lead_source_requirements.json
```

## List settings learning preferences

```
GET /settings/learning_preferences
```

List all settings learning preferences, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `preferences` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`preferences` — object:
| Field | Type | Required |
|---|---|---|
| `auto_create_packages` | object | Yes |
| `suggest_pricing_changes` | object | Yes |
| `proactive_campaigns` | object | Yes |
| `auto_refine_inspections` | object | Yes |
| `enable_cross_shop_learning` | object | Yes |
| `auto_catalog_cleanup` | object | Yes |

`auto_create_packages` — object:
| Field | Type | Required |
|---|---|---|
| `label` | string | Yes |
| `description` | string | Yes |
| `default` | boolean | Yes |

`suggest_pricing_changes` — object:
| Field | Type | Required |
|---|---|---|
| `label` | string | Yes |
| `description` | string | Yes |
| `default` | boolean | Yes |

`proactive_campaigns` — object:
| Field | Type | Required |
|---|---|---|
| `label` | string | Yes |
| `description` | string | Yes |
| `default` | boolean | Yes |

`auto_refine_inspections` — object:
| Field | Type | Required |
|---|---|---|
| `label` | string | Yes |
| `description` | string | Yes |
| `default` | boolean | Yes |

`enable_cross_shop_learning` — object:
| Field | Type | Required |
|---|---|---|
| `label` | string | Yes |
| `description` | string | Yes |
| `default` | boolean | Yes |

`auto_catalog_cleanup` — object:
| Field | Type | Required |
|---|---|---|
| `label` | string | Yes |
| `description` | string | Yes |
| `default` | boolean | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/learning_preferences.json
```

## List settings notifications edit

```
GET /settings/notifications/edit
```

List all settings notifications edit, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `email_fallback_enabled` | boolean | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/notifications/edit.json
```

## List settings payments

```
GET /settings/payments
```

List all settings payments, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `processor_application_status` | string \| null | Yes |
| `processor_onboarded_at` | string \| null | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/payments.json
```

## List settings phone numbers

```
GET /settings/phone_numbers
```

List all settings phone numbers, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `texting_phone` | string \| null | Yes |
| `phones` | array of any | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/phone_numbers.json
```

## List settings quickbooks

```
GET /settings/quickbooks
```

List all settings quickbooks, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `connected` | boolean | Yes |
| `qbo_company_id` | string \| null | Yes |
| `qbo_sync_mode` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/quickbooks.json
```

## List settings reminders

```
GET /settings/reminders
```

List all settings reminders, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `cash_drawer_enabled` | boolean | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `expenses_enabled` | boolean | Yes |
| `tire_management_enabled` | boolean | Yes |
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/reminders.json
```

## List tags

```
GET /settings/tags
```

List all tags, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `customer_tags` | array of object | Yes |
| `vehicle_tags` | array of object | Yes |

**Example**

```json
{
  "customer_tags": [
    {
      "id": 1,
      "account_id": 1,
      "name": "VIP",
      "color": "blue",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "customers_count": 5,
      "metadata": {},
      "updated_at": "2026-08-27T12:00:00.000-04:00"
    },
    {
      "id": 2,
      "account_id": 1,
      "name": "FLEET",
      "color": "green",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "customers_count": 2,
      "metadata": {},
      "updated_at": "2026-08-27T12:00:00.000-04:00"
    }
  ],
  "vehicle_tags": [
    {
      "id": 3,
      "account_id": 1,
      "name": "FLEET",
      "color": "green",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "updated_at": "2026-08-27T12:00:00.000-04:00"
    },
    {
      "id": 4,
      "account_id": 1,
      "name": "LEASED",
      "color": "amber",
      "created_at": "2026-08-27T12:00:00.000-04:00",
      "updated_at": "2026-08-27T12:00:00.000-04:00"
    }
  ]
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/tags.json
```

## Update tags

```
PATCH /settings/tags
```

Update a tags by ID.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `customer_tags` | array of object | Yes |
| `vehicle_tags` | array of any | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/settings/tags.json
```

## List settings tire management

```
GET /settings/tire_management
```

List all settings tire management, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `cash_drawer_enabled` | boolean | Yes |
| `default_starting_float_cents` | integer | Yes |
| `driveon_station_number` | string \| null | Yes |
| `expenses_enabled` | boolean | Yes |
| `tire_management_enabled` | boolean | Yes |
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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/tire_management.json
```

## List settings trust levels

```
GET /settings/trust_levels
```

List all settings trust levels, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `trust_levels` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`trust_levels` — object:
| Field | Type | Required |
|---|---|---|
| `create_work_order` | object | Yes |
| `update_work_order` | object | Yes |
| `add_service_to_work_order` | object | Yes |
| `book_appointment` | object | Yes |
| `create_customer` | object | Yes |
| `send_estimate` | object | Yes |
| `send_message` | object | Yes |
| `send_invoice` | object | Yes |
| `send_follow_up` | object | Yes |

`create_work_order` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

`update_work_order` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

`add_service_to_work_order` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

`book_appointment` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

`create_customer` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

`send_estimate` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

`send_message` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

`send_invoice` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

`send_follow_up` — object:
| Field | Type | Required |
|---|---|---|
| `level` | string | Yes |
| `escalation` | object | Yes |

`escalation` — object:
| Field | Type | Required |
|---|---|---|
| `total` | integer | Yes |
| `unedited` | integer | Yes |
| `edited` | integer | Yes |
| `consecutive_unedited` | integer | Yes |
| `eligible_for_promotion` | boolean | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/settings/trust_levels.json
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

### UpdateSettingsAccountRequest schema {#updatesettingsaccountrequest-schema}

| Field | Type | Required |
|---|---|---|
| `account` | object | Yes |

`account` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

---

### UpdateSettingsLaborTemplateRequest schema {#updatesettingslabortemplaterequest-schema}

| Field | Type | Required |
|---|---|---|
| `labor_template` | object | Yes |

`labor_template` — object:
| Field | Type | Required |
|---|---|---|
| `default_hours` | number | Yes |

---

### UpdateTagsRequest schema {#updatetagsrequest-schema}

| Field | Type | Required |
|---|---|---|
| `customer_tags` | array of object | No |
| `vehicle_tags` | array of object | No |

