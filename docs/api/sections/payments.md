# Payments

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List payments

```
GET /payments
```

List all payments, paginated via the Link header.

| Param | Type | Required |
|---|---|---|
| `method` | string | No |

**Response 403** — [Error](#error-schema) error envelope

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/payments.json
```

## List payments pending

```
GET /payments/pending
```

List all payments pending, paginated via the Link header.

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/payments/pending.json
```

## Show payment

```
GET /payments/{id}
```

Show a payment by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/payments/<id>.json
```

## Create payments cancellation

```
POST /payments/{id}/cancellation
```

Create a payments cancellation.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/payments/<id>.json
```

## Create payments confirmation

```
POST /payments/{id}/confirmation
```

Create a payments confirmation.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/payments/<id>.json
```

## Create payments failure

```
POST /payments/{id}/failure
```

Create a payments failure.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 500** — [Error](#error-schema) error envelope

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/payments/<id>.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

