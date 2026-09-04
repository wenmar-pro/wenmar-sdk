# Account

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Delete account

```
DELETE /account
```

Delete a account by ID.

**Response 204** — no content.

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account.json
```

## List account

```
GET /account
```

List all account, paginated via the Link header.

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
| `formatted_join_code` | string | Yes |
| `station_login_url` | string | Yes |
| `deletion_scheduled_at` | string \| null | Yes |
| `locations` | array of object | Yes |

**Response 401** — [Error](#error-schema) error envelope

**Example**

```json
{
  "id": 1,
  "name": "Main Shop",
  "slug": "main-shop",
  "locations": [
    {
      "id": 1,
      "name": "Bay 1",
      "url": "https://app.wenmarpro.com/locations/1.json",
      "app_url": "https://app.wenmarpro.com/locations/1"
    }
  ],
  "url": "https://app.wenmarpro.com/account.json",
  "app_url": "https://app.wenmarpro.com/account"
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account.json
```

## Update account

```
PATCH /account
```

Update a account by ID.

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
| `formatted_join_code` | string | Yes |
| `station_login_url` | string | Yes |
| `deletion_scheduled_at` | string \| null | Yes |
| `locations` | array of object | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/account.json
```

## List account billing

```
GET /account/billing
```

List all account billing, paginated via the Link header.

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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account/billing.json
```

## List account capabilities

```
GET /account/capabilities
```

List all account capabilities, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `tier` | string \| null | Yes |
| `tier_display` | string \| null | Yes |
| `capabilities` | object | Yes |
| `limits` | object | Yes |

`capabilities` — object:
| Field | Type | Required |
|---|---|---|
| `appointments` | boolean | Yes |
| `payments` | boolean | Yes |
| `inventory` | boolean | Yes |
| `purchase_orders` | boolean | Yes |
| `counter_sales` | boolean | Yes |
| `cash_drawer` | boolean | Yes |
| `courtesy_cars` | boolean | Yes |
| `dealer_plates` | boolean | Yes |
| `tire_management` | boolean | Yes |
| `expenses` | boolean | Yes |
| `phone_numbers` | boolean | Yes |
| `integrations` | object | Yes |

`integrations` — object:
| Field | Type | Required |
|---|---|---|
| `quickbooks` | boolean | Yes |
| `driveon` | boolean | Yes |

`limits` — object:
| Field | Type | Required |
|---|---|---|
| `max_locations` | integer \| null | Yes |
| `max_users` | integer \| null | Yes |
| `max_work_orders_per_month` | any | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account/capabilities.json
```

## List account driveon

```
GET /account/driveon
```

List all account driveon, paginated via the Link header.

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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account/driveon.json
```

## List account payments

```
GET /account/payments
```

List all account payments, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `processor_application_status` | string \| null | Yes |
| `processor_onboarded_at` | string \| null | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account/payments.json
```

## List account phone numbers

```
GET /account/phone_numbers
```

List all account phone numbers, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `texting_phone` | string \| null | Yes |
| `phones` | array of any | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account/phone_numbers.json
```

## List account quickbooks

```
GET /account/quickbooks
```

List all account quickbooks, paginated via the Link header.

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
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account/quickbooks.json
```

## List account station link

```
GET /account/station_link
```

List all account station link, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `join_code` | string | Yes |
| `formatted_join_code` | string | Yes |
| `station_login_url` | string | Yes |
| `qr_code_url` | string | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/account/station_link.json
```

## Create account station link regenerate

```
POST /account/station_link/regenerate
```

Create a account station link regenerate.

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
| `formatted_join_code` | string | Yes |
| `station_login_url` | string | Yes |
| `deletion_scheduled_at` | string \| null | Yes |
| `locations` | array of object | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/account/station_link/regenerate.json
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

### UpdateAccountRequest schema {#updateaccountrequest-schema}

| Field | Type | Required |
|---|---|---|
| `account` | object | Yes |

`account` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

