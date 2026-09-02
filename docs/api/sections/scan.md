# Scan

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create scan lookup

```
POST /scan/lookups
```

Create a scan lookup.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `outcome` | string | Yes |
| `work_order_id` | integer \| null | Yes |
| `vehicle_id` | integer \| null | Yes |
| `customer_id` | integer \| null | Yes |
| `appointment_id` | integer \| null | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/scan/lookups.json
```

## Create scan started work order

```
POST /scan/started_work_orders
```

Create a scan started work order.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |
| `work_order_id` | integer | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/scan/started_work_orders.json
```

## Create scan vehicle

```
POST /scan/vehicles
```

Create a scan vehicle.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `status` | string | Yes |
| `work_order_id` | integer | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/scan/vehicles.json
```

---

### CreateScanLookupRequest schema {#createscanlookuprequest-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `type` | string | Yes |

---

### CreateScanStartedWorkOrderRequest schema {#createscanstartedworkorderrequest-schema}

| Field | Type | Required |
|---|---|---|
| `outcome` | string | Yes |
| `vehicle_id` | integer | Yes |

---

### CreateScanVehicleRequest schema {#createscanvehiclerequest-schema}

| Field | Type | Required |
|---|---|---|
| `customer` | object | Yes |
| `vehicle` | object | Yes |

`customer` — object:
| Field | Type | Required |
|---|---|---|
| `first_name` | string | Yes |
| `last_name` | string | Yes |
| `phone` | string | Yes |

`vehicle` — object:
| Field | Type | Required |
|---|---|---|
| `vin` | string | Yes |
| `year` | string | Yes |
| `make` | string | Yes |
| `model` | string | Yes |

