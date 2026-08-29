# Customers

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List customers

```
GET /customers
```

List all customers, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `has_balance` | boolean | No |
| `has_vehicle` | boolean | No |
| `last_visit_months` | integer | No |
| `page` | integer | No |
| `per_page` | integer | No |
| `query` | string | No |
| `tag_ids` | array of string | No |
| `type` | string | No |

**Response 200** — array of [Customer](#customer-schema)

**Response 401** — [Error](#error-schema) error envelope

**Example**

```json
[
  {
    "type": "Customer",
    "id": 1,
    "full_name": "Jane Doe",
    "company_name": null,
    "first_name": "Jane",
    "last_name": "Doe",
    "fleet_identifier": null,
    "marketing_opt_in": false,
    "tax_exempt": false,
    "vehicles_count": 2,
    "emails_count": 1,
    "phones_count": 1,
    "emails": [
      {
        "id": 1,
        "address": "jane@example.com",
        "label": "personal",
        "primary": true
      }
    ],
    "phones": [
      {
        "id": 1,
        "number": "+15551234567",
        "label": "mobile",
        "primary": true
      }
    ],
    "addresses": [],
    "vehicles_url": "https://app.wenmarpro.com/vehicles.json?customer_id=1",
    "work_orders_url": "https://app.wenmarpro.com/work_orders.json?customer_id=1",
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "app_url": "https://app.wenmarpro.com/customers/1",
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "outstanding_balance_cents": 0,
    "total_revenue_cents": 50000,
    "store_credit_cents": 0,
    "last_visit_at": "2026-08-20T10:00:00.000-04:00",
    "statements_count": 2,
    "currency": "CAD"
  },
  {
    "type": "Customer",
    "id": 2,
    "full_name": "John Smith",
    "company_name": null,
    "first_name": "John",
    "last_name": "Smith",
    "fleet_identifier": null,
    "marketing_opt_in": false,
    "tax_exempt": false,
    "vehicles_count": 1,
    "emails_count": 1,
    "phones_count": 1,
    "emails": [
      {
        "id": 2,
        "address": "john@example.com",
        "label": "personal",
        "primary": true
      }
    ],
    "phones": [
      {
        "id": 2,
        "number": "+15559876543",
        "label": "mobile",
        "primary": true
      }
    ],
    "addresses": [],
    "vehicles_url": "https://app.wenmarpro.com/vehicles.json?customer_id=2",
    "work_orders_url": "https://app.wenmarpro.com/work_orders.json?customer_id=2",
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/customers/2.json",
    "app_url": "https://app.wenmarpro.com/customers/2",
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "outstanding_balance_cents": 0,
    "total_revenue_cents": 0,
    "store_credit_cents": 0,
    "last_visit_at": null,
    "statements_count": 0,
    "currency": "CAD"
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers.json
```

## Create customer

```
POST /customers
```

Create a customer.

**Request body** — wrapper key `customer`:

| Field | Type | Required |
|---|---|---|
| `first_name` | string | Yes |
| `last_name` | string | Yes |
| `company_name` | string | No |
| `fleet_identifier` | string | No |
| `billing_terms` | string | No |
| `credit_limit_cents` | string | No |
| `tax_exempt` | boolean | No |
| `tax_exempt_number` | string | No |
| `notes` | string | No |
| `marketing_opt_in` | boolean | No |
| `discount_percent` | string | No |
| `po_required` | boolean | No |
| `tag_ids` | array of any | No |
| `emails_attributes` | array of object | No |
| `phones_attributes` | array of object | No |
| `addresses_attributes` | array of object | No |

**Response 201** — [Customer](#customer-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

**Example**

```json
{
  "type": "Customer",
  "id": 1,
  "full_name": "Jane Doe",
  "company_name": null,
  "first_name": "Jane",
  "last_name": "Doe",
  "fleet_identifier": null,
  "marketing_opt_in": false,
  "tax_exempt": false,
  "vehicles_count": 2,
  "emails_count": 1,
  "phones_count": 1,
  "emails": [
    {
      "id": 1,
      "address": "jane@example.com",
      "label": "personal",
      "primary": true
    }
  ],
  "phones": [
    {
      "id": 1,
      "number": "+15551234567",
      "label": "mobile",
      "primary": true
    }
  ],
  "addresses": [],
  "vehicles_url": "https://app.wenmarpro.com/vehicles.json?customer_id=1",
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?customer_id=1",
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/customers/1.json",
  "app_url": "https://app.wenmarpro.com/customers/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "outstanding_balance_cents": 0,
  "total_revenue_cents": 50000,
  "store_credit_cents": 0,
  "last_visit_at": "2026-08-20T10:00:00.000-04:00",
  "statements_count": 2,
  "currency": "CAD"
}
```

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customers.json
```

## Check customer duplicate

```
GET /customers/check_duplicate
```

Check duplicate

| Param | Type | Required |
|---|---|---|
| `email` | string | No |
| `first_name` | string | No |
| `last_name` | string | No |
| `phone` | integer | No |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `matches` | array of object | Yes |

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/check_duplicate.json
```

## Lookup customer

```
GET /customers/lookup
```

Lookup

| Param | Type | Required |
|---|---|---|
| `id` | integer | No |
| `query` | string | No |

**Response 200** — array of [Customer](#customer-schema)

**Example**

```json
[
  {
    "type": "Customer",
    "id": 1,
    "full_name": "Jane Doe",
    "company_name": null,
    "first_name": "Jane",
    "last_name": "Doe",
    "fleet_identifier": null,
    "marketing_opt_in": false,
    "tax_exempt": false,
    "vehicles_count": 2,
    "emails_count": 1,
    "phones_count": 1,
    "emails": [
      {
        "id": 1,
        "address": "jane@example.com",
        "label": "personal",
        "primary": true
      }
    ],
    "phones": [
      {
        "id": 1,
        "number": "+15551234567",
        "label": "mobile",
        "primary": true
      }
    ],
    "addresses": [],
    "vehicles_url": "https://app.wenmarpro.com/customers/1/vehicles.json",
    "work_orders_url": "https://app.wenmarpro.com/customers/1/work_orders.json",
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/customers/1.json",
    "app_url": "https://app.wenmarpro.com/customers/1",
    "location": {
      "id": 1,
      "name": "Main Shop",
      "url": "https://app.wenmarpro.com/locations/1.json"
    },
    "outstanding_balance_cents": 0,
    "total_revenue_cents": 50000,
    "store_credit_cents": 0,
    "last_visit_at": "2026-08-20T10:00:00.000-04:00",
    "statements_count": 2,
    "currency": "CAD"
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/lookup.json
```

## List customers drivers

```
GET /customers/{customer_id}/drivers
```

List all customers drivers, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |

**Response 200** — array of [Driver](#driver-schema)

**Example**

```json
[
  {
    "id": 1,
    "full_name": "Jane Doe",
    "phone": "+1 416 555 0132",
    "email": "jane@acmeauto.com",
    "customer": {
      "id": 1,
      "full_name": "Jane Doe",
      "url": "https://app.wenmarpro.com/customers/1.json"
    },
    "work_orders_count": 3,
    "work_orders_url": "https://app.wenmarpro.com/customers/1/work_orders.json",
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/customers/1/drivers/1.json",
    "app_url": "https://app.wenmarpro.com/customers/1/drivers/1"
  },
  {
    "id": 2,
    "full_name": "John Smith",
    "phone": "+1 416 555 0199",
    "email": "john@acmeauto.com",
    "customer": {
      "id": 1,
      "full_name": "Jane Doe",
      "url": "https://app.wenmarpro.com/customers/1.json"
    },
    "work_orders_count": 1,
    "work_orders_url": "https://app.wenmarpro.com/customers/1/work_orders.json",
    "created_at": "2026-08-27T12:00:00.000-04:00",
    "updated_at": "2026-08-27T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/customers/1/drivers/2.json",
    "app_url": "https://app.wenmarpro.com/customers/1/drivers/2"
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/{customer_id}/drivers.json
```

## Create driver

```
POST /customers/{customer_id}/drivers
```

Create a driver.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |

**Request body** — wrapper key `driver`:

| Field | Type | Required |
|---|---|---|
| `full_name` | string | Yes |
| `phone` | string | Yes |

**Response 201** — [Driver](#driver-schema)

**Response 422** — [Error](#error-schema) error envelope

**Example**

```json
{
  "id": 1,
  "full_name": "Jane Doe",
  "phone": "+1 416 555 0132",
  "email": "jane@acmeauto.com",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "work_orders_count": 3,
  "work_orders_url": "https://app.wenmarpro.com/customers/1/work_orders.json",
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/customers/1/drivers/1.json",
  "app_url": "https://app.wenmarpro.com/customers/1/drivers/1"
}
```

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customers/{customer_id}/drivers.json
```

## Delete driver

```
DELETE /customers/{customer_id}/drivers/{id}
```

Delete a driver by ID.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/<id>.json
```

## Show driver

```
GET /customers/{customer_id}/drivers/{id}
```

Show a driver by ID.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |
| `id` | integer | Yes |

**Response 200** — [Driver](#driver-schema)

**Example**

```json
{
  "id": 1,
  "full_name": "Jane Doe",
  "phone": "+1 416 555 0132",
  "email": "jane@acmeauto.com",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "work_orders_count": 3,
  "work_orders_url": "https://app.wenmarpro.com/customers/1/work_orders.json",
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/customers/1/drivers/1.json",
  "app_url": "https://app.wenmarpro.com/customers/1/drivers/1"
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/<id>.json
```

## Update driver

```
PATCH /customers/{customer_id}/drivers/{id}
```

Update a driver by ID.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |
| `id` | integer | Yes |

**Request body** — wrapper key `driver`:

| Field | Type | Required |
|---|---|---|
| `full_name` | string | Yes |

**Response 200** — [Driver](#driver-schema)

**Example**

```json
{
  "id": 1,
  "full_name": "Jane Doe",
  "phone": "+1 416 555 0132",
  "email": "jane@acmeauto.com",
  "customer": {
    "id": 1,
    "full_name": "Jane Doe",
    "url": "https://app.wenmarpro.com/customers/1.json"
  },
  "work_orders_count": 3,
  "work_orders_url": "https://app.wenmarpro.com/customers/1/work_orders.json",
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/customers/1/drivers/1.json",
  "app_url": "https://app.wenmarpro.com/customers/1/drivers/1"
}
```

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customers/<id>.json
```

## List customers statements

```
GET /customers/{customer_id}/statements
```

List all customers statements, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |

**Response 200** — array of [Statement](#statement-schema)

**Example**

```json
[
  {
    "id": 1,
    "statement_number": "ST-0001",
    "status": "sent",
    "statement_date": "2026-08-01",
    "start_date": "2026-07-01",
    "end_date": "2026-07-31",
    "due_date": "2026-08-15",
    "totals": {
      "previous_balance_cents": 0,
      "new_charges_cents": 12000,
      "payments_received_cents": 0,
      "credits_cents": 0,
      "balance_due_cents": 12000,
      "currency": "CAD"
    },
    "sent_at": "2026-08-01T12:00:00.000-04:00",
    "viewed_at": null,
    "customer": {
      "id": 1,
      "full_name": "Jane Doe",
      "url": "https://app.wenmarpro.com/customers/1.json"
    },
    "created_at": "2026-08-01T12:00:00.000-04:00",
    "updated_at": "2026-08-01T12:00:00.000-04:00",
    "url": "https://app.wenmarpro.com/statements/1.json",
    "app_url": "https://app.wenmarpro.com/statements/1"
  }
]
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/{customer_id}/statements.json
```

## List customers vehicles

```
GET /customers/{customer_id}/vehicles
```

List all customers vehicles, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |

**Response 200** — array of [Vehicle](#vehicle-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/{customer_id}/vehicles.json
```

## Get customers vehicle history

```
GET /customers/{customer_id}/vehicles/{vehicle_id}/history
```

Show

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |
| `vehicle_id` | integer | Yes |

**Response 200** — [Vehicle](#vehicle-schema)

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/{customer_id}/vehicles/{vehicle_id}/history.json
```

## List customers work orders

```
GET /customers/{customer_id}/work_orders
```

List all customers work orders, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `customer_id` | integer | Yes |

**Response 200** — array of [WorkOrder](#workorder-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/{customer_id}/work_orders.json
```

## Delete customer

```
DELETE /customers/{id}
```

Delete a customer by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 202** — [Customer](#customer-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 422** — [Error](#error-schema) error envelope

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/<id>.json
```

## Show customer

```
GET /customers/{id}
```

Show a customer by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [Customer](#customer-schema)

**Response 404** — [Error](#error-schema) error envelope

**Example**

```json
{
  "type": "Customer",
  "id": 1,
  "full_name": "Jane Doe",
  "company_name": null,
  "first_name": "Jane",
  "last_name": "Doe",
  "fleet_identifier": null,
  "marketing_opt_in": false,
  "tax_exempt": false,
  "vehicles_count": 2,
  "emails_count": 1,
  "phones_count": 1,
  "emails": [
    {
      "id": 1,
      "address": "jane@example.com",
      "label": "personal",
      "primary": true
    }
  ],
  "phones": [
    {
      "id": 1,
      "number": "+15551234567",
      "label": "mobile",
      "primary": true
    }
  ],
  "addresses": [],
  "vehicles_url": "https://app.wenmarpro.com/vehicles.json?customer_id=1",
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?customer_id=1",
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/customers/1.json",
  "app_url": "https://app.wenmarpro.com/customers/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "outstanding_balance_cents": 0,
  "total_revenue_cents": 50000,
  "store_credit_cents": 0,
  "last_visit_at": "2026-08-20T10:00:00.000-04:00",
  "statements_count": 2,
  "currency": "CAD"
}
```

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/customers/<id>.json
```

## Update customer

```
PATCH /customers/{id}
```

Update a customer by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Request body** — wrapper key `customer`:

| Field | Type | Required |
|---|---|---|
| `phones_attributes` | array of object | No |
| `first_name` | string | No |
| `last_name` | string | No |
| `company_name` | string | No |
| `fleet_identifier` | string | No |
| `billing_terms` | string | No |
| `credit_limit_cents` | string | No |
| `tax_exempt` | boolean | No |
| `notes` | string | No |
| `marketing_opt_in` | boolean | No |
| `discount_percent` | string | No |
| `po_required` | boolean | No |
| `emails_attributes` | array of object | No |
| `addresses_attributes` | array of object | No |

**Response 200** — [Customer](#customer-schema)

**Example**

```json
{
  "type": "Customer",
  "id": 1,
  "full_name": "Jane Doe",
  "company_name": null,
  "first_name": "Jane",
  "last_name": "Doe",
  "fleet_identifier": null,
  "marketing_opt_in": false,
  "tax_exempt": false,
  "vehicles_count": 2,
  "emails_count": 1,
  "phones_count": 1,
  "emails": [
    {
      "id": 1,
      "address": "jane@example.com",
      "label": "personal",
      "primary": true
    }
  ],
  "phones": [
    {
      "id": 1,
      "number": "+15551234567",
      "label": "mobile",
      "primary": true
    }
  ],
  "addresses": [],
  "vehicles_url": "https://app.wenmarpro.com/vehicles.json?customer_id=1",
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?customer_id=1",
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/customers/1.json",
  "app_url": "https://app.wenmarpro.com/customers/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "outstanding_balance_cents": 0,
  "total_revenue_cents": 50000,
  "store_credit_cents": 0,
  "last_visit_at": "2026-08-20T10:00:00.000-04:00",
  "statements_count": 2,
  "currency": "CAD"
}
```

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customers/<id>.json
```

## Merge customer

```
POST /customers/{id}/merge
```

Perform merge

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Request body**:

| Field | Type | Required |
|---|---|---|
| `source_customer_id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `company_name` | any | Yes |
| `first_name` | string | Yes |
| `last_name` | string | Yes |
| `fleet_identifier` | any | Yes |
| `marketing_opt_in` | boolean | Yes |
| `tax_exempt` | boolean | Yes |
| `vehicles_count` | integer | Yes |
| `emails_count` | integer | Yes |
| `phones_count` | integer | Yes |
| `vehicles_url` | string | Yes |
| `work_orders_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `emails` | array of object | Yes |
| `phones` | array of object | Yes |
| `addresses` | array of any | Yes |
| `outstanding_balance_cents` | integer | Yes |
| `total_revenue_cents` | integer | Yes |
| `store_credit_cents` | integer | Yes |
| `last_visit_at` | any | Yes |
| `statements_count` | integer | Yes |
| `currency` | string | Yes |

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
  "type": "Customer",
  "id": 1,
  "full_name": "Jane Doe",
  "company_name": null,
  "first_name": "Jane",
  "last_name": "Doe",
  "fleet_identifier": null,
  "marketing_opt_in": false,
  "tax_exempt": false,
  "vehicles_count": 2,
  "emails_count": 1,
  "phones_count": 1,
  "emails": [
    {
      "id": 1,
      "address": "jane@example.com",
      "label": "personal",
      "primary": true
    }
  ],
  "phones": [
    {
      "id": 1,
      "number": "+15551234567",
      "label": "mobile",
      "primary": true
    }
  ],
  "addresses": [],
  "vehicles_url": "https://app.wenmarpro.com/vehicles.json?customer_id=1",
  "work_orders_url": "https://app.wenmarpro.com/work_orders.json?customer_id=1",
  "created_at": "2026-08-27T12:00:00.000-04:00",
  "updated_at": "2026-08-27T12:00:00.000-04:00",
  "url": "https://app.wenmarpro.com/customers/1.json",
  "app_url": "https://app.wenmarpro.com/customers/1",
  "location": {
    "id": 1,
    "name": "Main Shop",
    "url": "https://app.wenmarpro.com/locations/1.json"
  },
  "outstanding_balance_cents": 0,
  "total_revenue_cents": 50000,
  "store_credit_cents": 0,
  "last_visit_at": "2026-08-20T10:00:00.000-04:00",
  "statements_count": 2,
  "currency": "CAD"
}
```

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/customers/<id>.json
```

---

### Customer schema {#customer-schema}

| Field | Type | Required |
|---|---|---|
| `type` | string | Yes |
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `company_name` | string \| null | Yes |
| `first_name` | string | Yes |
| `last_name` | string | Yes |
| `fleet_identifier` | string \| null | Yes |
| `marketing_opt_in` | boolean | Yes |
| `tax_exempt` | boolean | Yes |
| `vehicles_count` | integer | Yes |
| `emails_count` | integer | Yes |
| `phones_count` | integer | Yes |
| `vehicles_url` | string | Yes |
| `work_orders_url` | string | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `location` | object | Yes |
| `emails` | array of object | Yes |
| `phones` | array of object | Yes |
| `addresses` | array of object | Yes |
| `outstanding_balance_cents` | integer | Yes |
| `total_revenue_cents` | integer | Yes |
| `store_credit_cents` | integer | Yes |
| `last_visit_at` | any | Yes |
| `statements_count` | integer | Yes |
| `currency` | string | Yes |

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

### Driver schema {#driver-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `phone` | any | Yes |
| `email` | any | Yes |
| `customer` | object | Yes |
| `work_orders_count` | integer | Yes |
| `work_orders_url` | string | Yes |
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

---

### Statement schema {#statement-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `statement_number` | string | Yes |
| `status` | string | Yes |
| `statement_date` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `due_date` | string | Yes |
| `totals` | object | Yes |
| `sent_at` | any | Yes |
| `viewed_at` | any | Yes |
| `customer` | object | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`totals` — object:
| Field | Type | Required |
|---|---|---|
| `previous_balance_cents` | integer | Yes |
| `new_charges_cents` | integer | Yes |
| `payments_received_cents` | integer | Yes |
| `credits_cents` | integer | Yes |
| `balance_due_cents` | integer | Yes |
| `currency` | string | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `full_name` | string | Yes |
| `url` | string | Yes |

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

