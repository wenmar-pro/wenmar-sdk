# Orders

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List orders purchase orders

```
GET /orders/purchase_orders
```

List all orders purchase orders, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `page` | integer | No |
| `per_page` | integer | No |
| `vendor_id` | integer | No |

**Response 200** — array of [PurchaseOrder](#purchaseorder-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/orders/purchase_orders.json
```

## Create orders purchase order

```
POST /orders/purchase_orders
```

Create a orders purchase order.

**Response 201** — [PurchaseOrder](#purchaseorder-schema)

**Response 403** — [Error](#error-schema) error envelope

**Response 404** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/orders/purchase_orders.json
```

## Delete orders purchase order

```
DELETE /orders/purchase_orders/{id}
```

Delete a orders purchase order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/orders/purchase_orders/<id>.json
```

## Show orders purchase order

```
GET /orders/purchase_orders/{id}
```

Show a orders purchase order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [PurchaseOrder](#purchaseorder-schema)

**Response 304** — no content.

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/orders/purchase_orders/<id>.json
```

## Update orders purchase order

```
PATCH /orders/purchase_orders/{id}
```

Update a orders purchase order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [PurchaseOrder](#purchaseorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/orders/purchase_orders/<id>.json
```

## Create orders purchase orders cancellation

```
POST /orders/purchase_orders/{purchase_order_id}/cancellations
```

Create a orders purchase orders cancellation.

| Param | Type | Required |
|---|---|---|
| `purchase_order_id` | integer | Yes |

**Response 200** — [PurchaseOrder](#purchaseorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/orders/purchase_orders/{purchase_order_id}/cancellations.json
```

## List orders return orders

```
GET /orders/return_orders
```

List all orders return orders, paginated via the Link header.

**Response 200** — array of [ReturnOrder](#returnorder-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/orders/return_orders.json
```

## Delete orders return order

```
DELETE /orders/return_orders/{id}
```

Delete a orders return order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 204** — no content.

```bash
curl -X DELETE -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/orders/return_orders/<id>.json
```

## Show orders return order

```
GET /orders/return_orders/{id}
```

Show a orders return order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [ReturnOrder](#returnorder-schema)

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/orders/return_orders/<id>.json
```

## Update orders return order

```
PATCH /orders/return_orders/{id}
```

Update a orders return order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [ReturnOrder](#returnorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/orders/return_orders/<id>.json
```

## Create orders return orders refund completion

```
POST /orders/return_orders/{return_order_id}/refund_completions
```

Create a orders return orders refund completion.

| Param | Type | Required |
|---|---|---|
| `return_order_id` | integer | Yes |

**Response 200** — [ReturnOrder](#returnorder-schema)

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/orders/return_orders/{return_order_id}/refund_completions.json
```

## List orders sublet orders

```
GET /orders/sublet_orders
```

List all orders sublet orders, paginated via the Link header.

**Response 200** — array of [SubletOrder](#subletorder-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/orders/sublet_orders.json
```

## Update orders sublet orders mark payment complete

```
PATCH /orders/sublet_orders/mark_payment_complete
```

Update a orders sublet orders mark payment complete by ID.

**Response 200** — [SubletOrder](#subletorder-schema)

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/orders/sublet_orders/mark_payment_complete.json
```

## Show orders sublet order

```
GET /orders/sublet_orders/{id}
```

Show a orders sublet order by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200** — [SubletOrder](#subletorder-schema)

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/orders/sublet_orders/<id>.json
```

---

### PurchaseOrder schema {#purchaseorder-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `po_number` | integer | Yes |
| `status` | string | Yes |
| `order_method` | string | Yes |
| `payment_method` | string | Yes |
| `fulfillment_method` | string | Yes |
| `tracking_number` | any | Yes |
| `vendor_invoice_number` | any | Yes |
| `vendor_invoice_received_at` | string \| null | Yes |
| `notes` | string \| null | Yes |
| `freight_cost_cents` | integer | Yes |
| `freight_cost_currency` | string | Yes |
| `subtotal_cents` | string | Yes |
| `total_cents` | string | Yes |
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
| `creator` | object | Yes |
| `location` | object | Yes |
| `line_items` | array of object | Yes |

`vendor` — object:
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

`location` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `url` | string | Yes |

---

### CreateOrdersPurchaseOrderRequest schema {#createorderspurchaseorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `purchase_order` | object | Yes |

`purchase_order` — object:
| Field | Type | Required |
|---|---|---|
| `vendor_id` | integer | Yes |
| `payment_method` | string | Yes |
| `fulfillment_method` | string | Yes |
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
| `unit_cost` | string | Yes |

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### UpdateOrdersPurchaseOrderRequest schema {#updateorderspurchaseorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `purchase_order` | object | Yes |

`purchase_order` — object:
| Field | Type | Required |
|---|---|---|
| `tracking_number` | string | Yes |

---

### ReturnOrder schema {#returnorder-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `return_number` | integer | Yes |
| `status` | string | Yes |
| `credit_method` | any | Yes |
| `reason_code` | any | Yes |
| `rma_number` | any | Yes |
| `is_warranty_claim` | boolean | Yes |
| `restocking_fee_cents` | integer | Yes |
| `shipping_fee_cents` | integer | Yes |
| `notes` | string \| null | Yes |
| `line_items_count` | integer | Yes |
| `refund_completed_at` | string \| null | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |
| `vendor` | object | Yes |
| `work_order` | object | Yes |
| `creator` | object | Yes |
| `location` | object | Yes |
| `line_items` | array of object | Yes |

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

`creator` — object:
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

---

### UpdateOrdersReturnOrderRequest schema {#updateordersreturnorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `return_order` | object | Yes |

`return_order` — object:
| Field | Type | Required |
|---|---|---|
| `rma_number` | string | Yes |

---

### SubletOrder schema {#subletorder-schema}

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `type` | string | Yes |
| `sublet_number` | integer | Yes |
| `title` | string | Yes |
| `payment_status` | string | Yes |
| `payment_method` | any | Yes |
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
| `vendor` | object | Yes |
| `work_order` | object | Yes |
| `work_order_service` | object | Yes |
| `location` | object | Yes |

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

### UpdateOrdersSubletOrdersMarkPaymentCompleteRequest schema {#updateorderssubletordersmarkpaymentcompleterequest-schema}

| Field | Type | Required |
|---|---|---|
| `sublet_order_ids` | array of integer | Yes |
| `payment_method` | string | Yes |

