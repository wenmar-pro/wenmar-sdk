# Cash Drawer Session

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## Create cash drawer session

```
POST /cash_drawer_session
```

Create a cash drawer session.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `state` | string | Yes |
| `opened_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `starting_float_cents` | integer | Yes |
| `expected_cash_cents` | integer | Yes |
| `actual_cash_cents` | integer | Yes |
| `expected_cheques_cents` | integer | Yes |
| `actual_cheques_cents` | integer | Yes |
| `variance_cents` | integer | Yes |
| `bank_deposit_cents` | integer | Yes |
| `leave_behind_cents` | integer | Yes |
| `variance_explanation` | any | Yes |
| `opened_by` | object | Yes |
| `closed_by` | any | Yes |
| `summary` | object | Yes |
| `ledger_url` | string | Yes |
| `close_url` | string | Yes |
| `url` | string | Yes |

`opened_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

`summary` — object:
| Field | Type | Required |
|---|---|---|
| `opening_balance_cents` | integer | Yes |
| `payments_cash_cents` | integer | Yes |
| `payments_cheque_cents` | integer | Yes |
| `entries_paid_in_cents` | integer | Yes |
| `entries_paid_out_cents` | integer | Yes |
| `closing_balance_cents` | integer | Yes |

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `state` | string | Yes |
| `opened_at` | string | Yes |
| `closed_at` | string \| null | Yes |
| `location_id` | integer | Yes |
| `starting_float_cents` | integer | Yes |
| `expected_cash_cents` | integer | Yes |
| `actual_cash_cents` | integer | Yes |
| `expected_cheques_cents` | integer | Yes |
| `actual_cheques_cents` | integer | Yes |
| `variance_cents` | integer | Yes |
| `bank_deposit_cents` | integer | Yes |
| `leave_behind_cents` | integer | Yes |
| `variance_explanation` | any | Yes |
| `opened_by` | object | Yes |
| `closed_by` | any | Yes |
| `summary` | object | Yes |
| `ledger_url` | string | Yes |
| `close_url` | string | Yes |
| `url` | string | Yes |

`opened_by` — object:
| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |

`summary` — object:
| Field | Type | Required |
|---|---|---|
| `opening_balance_cents` | integer | Yes |
| `payments_cash_cents` | integer | Yes |
| `payments_cheque_cents` | integer | Yes |
| `entries_paid_in_cents` | integer | Yes |
| `entries_paid_out_cents` | integer | Yes |
| `closing_balance_cents` | integer | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/cash_drawer_session.json
```

---

### CreateCashDrawerSessionRequest schema {#createcashdrawersessionrequest-schema}

| Field | Type | Required |
|---|---|---|
| `cash_drawer_session` | object | Yes |

`cash_drawer_session` — object:
| Field | Type | Required |
|---|---|---|
| `starting_float_cents` | integer | Yes |

