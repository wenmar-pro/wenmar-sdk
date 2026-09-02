# Reports

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List reports accounting

```
GET /reports/accounting
```

List all reports accounting, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `tax` | object | Yes |
| `payments` | object | Yes |
| `outstanding_balances` | array of any | Yes |
| `outstanding_total` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`tax` — object:
| Field | Type | Required |
|---|---|---|
| `tax_collected_by_type` | object | Yes |
| `tax_collected_by_province` | object | Yes |
| `total_tax_collected_cents` | integer | Yes |
| `taxable_vs_exempt` | object | Yes |

`tax_collected_by_type` — object:

`tax_collected_by_province` — object:

`taxable_vs_exempt` — object:
| Field | Type | Required |
|---|---|---|
| `taxable_cents` | integer | Yes |
| `exempt_cents` | integer | Yes |

`payments` — object:
| Field | Type | Required |
|---|---|---|
| `total_collected_cents` | integer | Yes |
| `collected_by_payment_method` | object | Yes |
| `payment_count` | integer | Yes |

`collected_by_payment_method` — object:

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/accounting.json
```

## List reports ar aging

```
GET /reports/ar_aging
```

List all reports ar aging, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `rows` | array of any | Yes |
| `totals` | object | Yes |
| `credit_rows` | array of any | Yes |
| `credits_total_cents` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`totals` — object:
| Field | Type | Required |
|---|---|---|
| `current_cents` | integer | Yes |
| `days_31_60_cents` | integer | Yes |
| `days_61_90_cents` | integer | Yes |
| `days_91_120_cents` | integer | Yes |
| `days_120_plus_cents` | integer | Yes |
| `total_cents` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/ar_aging.json
```

## List reports declined work

```
GET /reports/declined_work
```

List all reports declined work, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `rows` | array of any | Yes |
| `total_declined_value_cents` | integer | Yes |
| `declined_count` | integer | Yes |
| `category_breakdown` | object | Yes |
| `aging_buckets` | object | Yes |
| `saved_for_later` | array of any | Yes |
| `total_saved_for_later_cents` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`category_breakdown` — object:

`aging_buckets` — object:
| Field | Type | Required |
|---|---|---|
| `days_0_30` | integer | Yes |
| `days_31_60` | integer | Yes |
| `days_61_90` | integer | Yes |
| `days_90_plus` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/declined_work.json
```

## List reports end of day

```
GET /reports/end_of_day
```

List all reports end of day, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `date` | string | Yes |
| `total_sales_by_method` | object | Yes |
| `ar_postings_cents` | integer | Yes |
| `ar_payments_collected_cents` | integer | Yes |
| `store_credit_issued_cents` | integer | Yes |
| `store_credit_applied_cents` | integer | Yes |
| `over_short_cents` | integer | Yes |
| `total_cashiered_cents` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`total_sales_by_method` — object:

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/end_of_day.json
```

## List reports financial

```
GET /reports/financial
```

List all reports financial, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string \| null | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `tab` | string | Yes |
| `revenue` | object | Yes |
| `tax` | object | Yes |
| `payments` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`revenue` — object:
| Field | Type | Required |
|---|---|---|
| `total_revenue_cents` | integer | Yes |
| `revenue_by_category` | object | Yes |
| `revenue_by_location` | object | Yes |
| `average_repair_order_cents` | integer | Yes |
| `period_comparison` | object | Yes |

`revenue_by_category` — object:

`revenue_by_location` — object:

`period_comparison` — object:
| Field | Type | Required |
|---|---|---|
| `current_cents` | integer | Yes |
| `previous_cents` | integer | Yes |
| `difference_cents` | integer | Yes |

`tax` — object:
| Field | Type | Required |
|---|---|---|
| `total_tax_collected_cents` | integer | Yes |
| `tax_collected_by_type` | object | Yes |

`tax_collected_by_type` — object:

`payments` — object:
| Field | Type | Required |
|---|---|---|
| `total_collected_cents` | integer | Yes |
| `collected_by_payment_method` | object | Yes |
| `payment_count` | integer | Yes |

`collected_by_payment_method` — object:

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/financial.json
```

## List reports open work

```
GET /reports/open_work
```

List all reports open work, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `rows` | array of any | Yes |
| `total_estimated_value_cents` | integer | Yes |
| `status_summary` | object | Yes |
| `stuck_work` | array of any | Yes |
| `saved_for_later` | array of any | Yes |
| `total_saved_for_later_cents` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`status_summary` — object:

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/open_work.json
```

## List reports parts purchases

```
GET /reports/parts_purchases
```

List all reports parts purchases, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `rows` | array of any | Yes |
| `total_spend_cents` | integer | Yes |
| `po_count` | integer | Yes |
| `return_count` | integer | Yes |
| `net_units` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/parts_purchases.json
```

## List reports parts usage

```
GET /reports/parts_usage
```

List all reports parts usage, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `rows` | array of any | Yes |
| `total_cost_cents` | integer | Yes |
| `total_revenue_cents` | integer | Yes |
| `total_gp_cents` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/parts_usage.json
```

## List reports performance

```
GET /reports/performance
```

List all reports performance, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `view_type` | string | Yes |
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `kpi_summary` | object | Yes |
| `daily_target_cents` | integer | Yes |
| `revenue_series` | array of any | Yes |
| `car_count_series` | array of any | Yes |
| `aro_series` | array of any | Yes |
| `top_services` | array of any | Yes |
| `retention` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`kpi_summary` — object:
| Field | Type | Required |
|---|---|---|
| `total_gross_revenue_cents` | integer | Yes |
| `average_daily_revenue_cents` | integer | Yes |
| `active_days_count` | integer | Yes |
| `effective_labor_rate_cents` | integer | Yes |
| `close_ratio` | number | Yes |
| `hours_sold_per_ro` | number | Yes |

`retention` — object:
| Field | Type | Required |
|---|---|---|
| `new` | integer | Yes |
| `returning` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/performance.json
```

## List reports profit and loss

```
GET /reports/profit_and_loss
```

List all reports profit and loss, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `revenue` | object | Yes |
| `net_revenue_cents` | integer | Yes |
| `cogs` | object | Yes |
| `gross_profit_cents` | integer | Yes |
| `operating_expenses` | object | Yes |
| `total_operating_expenses_cents` | integer | Yes |
| `net_income_cents` | integer | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`revenue` — object:
| Field | Type | Required |
|---|---|---|
| `labor_cents` | integer | Yes |
| `parts_cents` | integer | Yes |
| `subcontracts_cents` | integer | Yes |
| `tires_cents` | integer | Yes |
| `fees_cents` | integer | Yes |
| `discount_cents` | integer | Yes |

`cogs` — object:
| Field | Type | Required |
|---|---|---|
| `parts_cost_cents` | integer | Yes |
| `subcontract_cost_cents` | integer | Yes |

`operating_expenses` — object:
| Field | Type | Required |
|---|---|---|
| `rent` | integer | Yes |
| `utilities` | integer | Yes |
| `payroll` | integer | Yes |
| `insurance` | integer | Yes |
| `supplies` | integer | Yes |
| `equipment` | integer | Yes |
| `marketing` | integer | Yes |
| `professional_services` | integer | Yes |
| `other` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/profit_and_loss.json
```

## List reports sales summary

```
GET /reports/sales_summary
```

List all reports sales summary, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `rows` | array of any | Yes |
| `totals` | object | Yes |
| `car_count` | integer | Yes |
| `invoice_count` | integer | Yes |
| `aro_cents` | integer | Yes |
| `segment_gp` | object | Yes |
| `kpi_summary` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`totals` — object:
| Field | Type | Required |
|---|---|---|
| `group_id` | integer \| null | Yes |
| `group_label` | string | Yes |
| `service_count` | integer | Yes |
| `billed_hours` | number | Yes |
| `actual_hours` | number | Yes |
| `fees_cents` | integer | Yes |
| `discount_cents` | integer | Yes |
| `est_taxes_cents` | integer | Yes |
| `job_total_cents` | integer | Yes |
| `labor_revenue_cents` | integer | Yes |
| `parts_revenue_cents` | integer | Yes |
| `stock_inv_cost_cents` | integer | Yes |
| `non_stock_inv_cost_cents` | integer | Yes |
| `sublet_cost_cents` | integer | Yes |
| `labor_cost_cents` | integer | Yes |
| `total_cost_cents` | integer | Yes |
| `gp_cents` | integer | Yes |
| `gp_percent` | integer | Yes |
| `average_job_cents` | integer | Yes |

`segment_gp` — object:
| Field | Type | Required |
|---|---|---|
| `labor` | object | Yes |
| `parts` | object | Yes |
| `sublet` | object | Yes |

`labor` — object:
| Field | Type | Required |
|---|---|---|
| `revenue_cents` | integer | Yes |
| `cost_cents` | integer | Yes |
| `gp_cents` | integer | Yes |
| `gp_percent` | integer | Yes |

`parts` — object:
| Field | Type | Required |
|---|---|---|
| `revenue_cents` | integer | Yes |
| `cost_cents` | integer | Yes |
| `gp_cents` | integer | Yes |
| `gp_percent` | integer | Yes |

`sublet` — object:
| Field | Type | Required |
|---|---|---|
| `revenue_cents` | integer | Yes |
| `cost_cents` | integer | Yes |
| `gp_cents` | integer | Yes |
| `gp_percent` | integer | Yes |

`kpi_summary` — object:
| Field | Type | Required |
|---|---|---|
| `total_gross_revenue_cents` | integer | Yes |
| `average_daily_revenue_cents` | integer | Yes |
| `active_days_count` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/sales_summary.json
```

## List reports service categories

```
GET /reports/service_categories
```

List all reports service categories, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `rows` | array of any | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/service_categories.json
```

## List reports store credit

```
GET /reports/store_credit
```

List all reports store credit, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `total_liability_cents` | integer | Yes |
| `average_balance_cents` | integer | Yes |
| `customers_with_balance` | array of any | Yes |
| `all_transactions` | array of any | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/store_credit.json
```

## List reports technician productivity

```
GET /reports/technician_productivity
```

List all reports technician productivity, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `rows` | array of any | Yes |
| `totals` | object | Yes |
| `period_comparison` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`period_comparison` — object:
| Field | Type | Required |
|---|---|---|
| `current` | object | Yes |
| `previous` | object | Yes |

`current` — object:
| Field | Type | Required |
|---|---|---|
| `clocked_hours` | integer | Yes |
| `billed_hours` | integer | Yes |
| `efficiency` | integer | Yes |
| `revenue` | integer | Yes |

`previous` — object:
| Field | Type | Required |
|---|---|---|
| `clocked_hours` | integer | Yes |
| `billed_hours` | integer | Yes |
| `efficiency` | integer | Yes |
| `revenue` | integer | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/technician_productivity.json
```

## List reports work order profitability

```
GET /reports/work_order_profitability
```

List all reports work order profitability, paginated via the Link header.

**Response 200**

| Field | Type | Required |
|---|---|---|
| `currency` | string | Yes |
| `start_date` | string | Yes |
| `end_date` | string | Yes |
| `rows` | array of any | Yes |
| `totals` | object | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

`totals` — object:
| Field | Type | Required |
|---|---|---|
| `work_order_count` | integer | Yes |
| `job_total_cents` | integer | Yes |
| `total_cost_cents` | integer | Yes |
| `gp_cents` | integer | Yes |
| `gp_percent` | integer | Yes |
| `avg_gp_cents` | integer | Yes |
| `billed_hours` | number | Yes |
| `actual_hours` | number | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/reports/work_order_profitability.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

