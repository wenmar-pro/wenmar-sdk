# Wenmar Pro API — Compact Reference

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

Base URL: `https://app.wenmarpro.com`
Auth: `Authorization: Bearer <token>`
Responses: bare objects/arrays, no envelope. Errors: `{ "error": { code, message, details } }`.

## Account

- **GET /account** -> 200: object{id,name,slug,locations,url,app_url} | 401: error envelope | 403: error envelope

## Ai Suggestions

- **PATCH /ai_suggestions/{id}** ?id -> 200: object{id,status} | 403: error envelope | 422: error envelope

## Appointments

- **GET /appointments** ?per_page,q,status -> 200: array of object | 500: object
- **POST /appointments** -> 201: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url} | 403: error envelope | 422: error envelope | 500: object
- **GET /appointments/available_slots** ?date,duration_minutes -> 200: array of object
- **DELETE /appointments/{id}** ?id -> no content
- **GET /appointments/{id}** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url} | no content | 500: object
- **PATCH /appointments/{id}** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url} | 500: object
- **POST /appointments/{id}/approvals** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url} | 500: object
- **POST /appointments/{id}/cancellations** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url} | 500: object
- **POST /appointments/{id}/follow_ups** ?id -> 201: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url}
- **POST /appointments/{id}/rejections** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url} | 500: object
- **POST /appointments/{id}/vehicle_reconciliations** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url} | 500: object
- **POST /appointments/{id}/work_orders** ?id -> 201: WorkOrder

## Campaigns

- **GET /campaigns** -> 200: array of BroadcastCampaign
- **POST /campaigns** -> 201: BroadcastCampaign | 403: error envelope
- **GET /campaigns/{id}** ?id -> 200: BroadcastCampaign
- **POST /campaigns/{id}/duplicate** ?id -> 200: object{id,type,name,status,sms_body,filters,recipient_count,sent_count,failed_count,progress_percentage,sent_at,created_at,updated_at,url,app_url,creator,location}
- **POST /campaigns/{id}/send_campaign** ?id -> 200: object{id,type,name,status,sms_body,filters,recipient_count,sent_count,failed_count,progress_percentage,sent_at,created_at,updated_at,url,app_url,creator,location}

## Cash Drawer Banner

- **PATCH /cash_drawer_banner** -> no content

## Catalog Cleanups

- **POST /catalog_cleanups/{catalog_cleanup_id}/applications** ?catalog_cleanup_id -> 200: object{status,applied_count} | 403: error envelope

## Conversations

- **GET /conversations** -> 200: array of object | no content
- **POST /conversations/bulk_mark_read** -> 200: object{ok,affected}
- **POST /conversations/{conversation_id}/customer_links** ?conversation_id -> 200: object{id,status,reply_state,channel,customer,unread_count,messages_url,created_at,updated_at,url,app_url}
- **POST /conversations/{conversation_id}/ignores** ?conversation_id -> 200: object{id,status,reply_state,channel,customer,unread_count,messages_url,created_at,updated_at,url,app_url}
- **GET /conversations/{conversation_id}/messages** ?conversation_id -> 200: array of object
- **POST /conversations/{conversation_id}/messages** ?conversation_id -> 201: object{id,conversation_id,direction,channel,status,body,sender,work_order_id,statement_id,appointment_id,read_at,created_at}
- **POST /conversations/{conversation_id}/messages/{id}/resends** ?conversation_id,id -> 200: object{id,conversation_id,direction,channel,status,body,work_order_id,statement_id,appointment_id,read_at,created_at}
- **GET /conversations/{id}** ?id -> 200: object{id,status,reply_state,channel,customer,unread_count,messages_url,created_at,updated_at,url,app_url}
- **PATCH /conversations/{id}** ?id -> 200: object{id,status,reply_state,channel,customer,unread_count,messages_url,created_at,updated_at,url,app_url}

## Core Tax Rules

- **GET /core_tax_rules** -> 200: array of object | 403: error envelope
- **PATCH /core_tax_rules/{id}** ?id -> 200: object{id,province_code,tax_core_charge,tax_core_credit,notes,created_at,updated_at,url,app_url}

## Counter Sales

- **GET /counter_sales** -> 200: array of object
- **POST /counter_sales** -> 201: object{id,counter_sale_number,status,walk_in_name,notes,subtotal_cents,tax_total_cents,total_cents,paid_cents,remaining_cents,paid,currency,line_items_count,processed_by,location,created_at,updated_at,url,app_url} | 403: error envelope
- **GET /counter_sales/{counter_sale_id}/line_items/brands** ?counter_sale_id -> 403: error envelope
- **GET /counter_sales/{id}** ?id -> 200: object{id,counter_sale_number,status,walk_in_name,notes,subtotal_cents,tax_total_cents,total_cents,paid_cents,remaining_cents,paid,currency,line_items_count,processed_by,location,created_at,updated_at,url,app_url}
- **PATCH /counter_sales/{id}** ?id -> 200: object{id,counter_sale_number,status,walk_in_name,notes,subtotal_cents,tax_total_cents,total_cents,paid_cents,remaining_cents,paid,currency,line_items_count,processed_by,location,created_at,updated_at,url,app_url}

## Courtesy Cars

- **GET /courtesy_cars** -> 200: array of object | 403: error envelope

## Customer Tags

- **GET /customer_tags** -> 200: array of object
- **POST /customer_tags** -> 201: object{id,name,color,color_hex,color_class} | 403: error envelope | 422: error envelope
- **DELETE /customer_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class}
- **PATCH /customer_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class}

## Customers

- **GET /customers** ?has_balance,has_vehicle,last_visit_months,page,per_page,q,tag_ids,type -> 200: array of Customer | 401: error envelope
- **POST /customers** -> 201: Customer | 403: error envelope | 422: error envelope
- **GET /customers/1043910089/merge** -> no content
- **GET /customers/check_duplicate** ?email,first_name,last_name,phone -> 200: object{matches} | 403: error envelope
- **POST /customers/export** -> 200: Customer | 403: error envelope
- **GET /customers/export/{id}/download** ?id -> no content | 404: error envelope
- **GET /customers/lookup** ?id,query -> 200: array of Customer | 403: error envelope
- **GET /customers/{customer_id}/drivers** ?customer_id -> 200: array of Driver
- **POST /customers/{customer_id}/drivers** ?customer_id -> 201: Driver | 422: error envelope
- **DELETE /customers/{customer_id}/drivers/{id}** ?customer_id,id -> no content
- **GET /customers/{customer_id}/drivers/{id}** ?customer_id,id -> 200: Driver
- **PATCH /customers/{customer_id}/drivers/{id}** ?customer_id,id -> 200: Driver
- **GET /customers/{customer_id}/statements** ?customer_id -> 200: array of Statement
- **GET /customers/{customer_id}/vehicles** ?customer_id -> 200: array of Vehicle
- **GET /customers/{customer_id}/vehicles/{vehicle_id}/history** ?customer_id,vehicle_id -> 200: Vehicle | 404: error envelope
- **GET /customers/{customer_id}/work_orders** ?customer_id -> 200: array of WorkOrder
- **DELETE /customers/{id}** ?id -> 202: Customer | 403: error envelope | 422: error envelope
- **GET /customers/{id}** ?id -> 200: Customer | 404: error envelope
- **PATCH /customers/{id}** ?id -> 200: Customer
- **POST /customers/{id}/merges** ?id -> 200: object{type,id,full_name,company_name,first_name,last_name,fleet_identifier,marketing_opt_in,tax_exempt,vehicles_count,emails_count,phones_count,vehicles_url,work_orders_url,created_at,updated_at,url,app_url,location,emails,phones,addresses,outstanding_balance_cents,total_revenue_cents,store_credit_cents,last_visit_at,statements_count,currency} | 403: error envelope | 422: error envelope

## Expenses

- **GET /expenses** -> 200: array of object | 403: error envelope | 500: object
- **POST /expenses** -> 201: object{id,payee,category,description,amount_cents,amount_currency,expense_date,payment_method,recurring,recurrence_rule,creator,location,created_at,updated_at,url,app_url} | 403: error envelope | 500: object
- **DELETE /expenses/{id}** ?id -> no content
- **PATCH /expenses/{id}** ?id -> 200: object{id,payee,category,description,amount_cents,amount_currency,expense_date,payment_method,recurring,recurrence_rule,creator,location,created_at,updated_at,url,app_url} | 500: object

## Fleets

- **GET /fleets** -> 200: array of object | 403: error envelope

## Inspection Reports

- **POST /inspection_reports** ?inspection_id,work_order_id -> 201: InspectionReport
- **DELETE /inspection_reports/{id}** ?id -> no content
- **GET /inspection_reports/{id}** ?id -> 200: InspectionReport | no content
- **PATCH /inspection_reports/{id}/complete** ?id -> 200: object{id,type,name,status,work_order_id,quick_finding,published,completed,items_count,checked_count,created_at,updated_at,url,app_url,location}
- **GET /inspection_reports/{id}/group** ?group_name,id -> 200: object{type,group_name,inspection_report_id,items}
- **POST /inspection_reports/{id}/mark_all** ?id -> 200: object{id,type,name,status,work_order_id,quick_finding,published,completed,items_count,checked_count,created_at,updated_at,url,app_url,location}
- **PATCH /inspection_reports/{id}/publish** ?id -> 200: object{id,type,name,status,work_order_id,quick_finding,published,completed,items_count,checked_count,created_at,updated_at,url,app_url,location}
- **PATCH /inspection_reports/{id}/reassign** ?id -> 200: object{id,type,name,status,work_order_id,quick_finding,published,completed,items_count,checked_count,created_at,updated_at,url,app_url,location}
- **PATCH /inspection_reports/{id}/reopen** ?id -> 200: object{id,type,name,status,work_order_id,quick_finding,published,completed,items_count,checked_count,created_at,updated_at,url,app_url,location}
- **PATCH /inspection_reports/{id}/reset** ?id -> 200: object{id,type,name,status,work_order_id,quick_finding,published,completed,items_count,checked_count,created_at,updated_at,url,app_url,location}
- **POST /inspection_reports/{id}/retry_recording** ?id,recording_id -> 200: object{id,status,error_message}
- **PATCH /inspection_reports/{id}/unpublish** ?id -> 200: object{id,type,name,status,work_order_id,quick_finding,published,completed,items_count,checked_count,created_at,updated_at,url,app_url,location}

## Inspections

- **GET /inspections** ?per_page -> 200: array of Inspection
- **POST /inspections** -> 201: Inspection | 403: error envelope
- **DELETE /inspections/{id}** ?id -> no content
- **GET /inspections/{id}** ?id -> 200: Inspection
- **PATCH /inspections/{id}** ?id -> 200: Inspection
- **PATCH /inspections/{id}/remove_default** ?id -> 200: object{id,type,name,description,active,is_default,created_at,updated_at,url,app_url,location,groups}
- **PATCH /inspections/{id}/set_default** ?id -> 200: object{id,type,name,description,active,is_default,created_at,updated_at,url,app_url,location,groups}
- **PATCH /inspections/{id}/toggle** ?id -> 200: object{id,type,name,description,active,is_default,created_at,updated_at,url,app_url,location,groups}

## Inventory Levels

- **GET /inventory_levels** ?brand,page,per_page,q,stock_status,stocked -> 200: array of object
- **POST /inventory_levels** -> 201: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url} | 403: error envelope | 422: error envelope
- **GET /inventory_levels/barcode_lookup** ?barcode -> 200: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url} | 404: error envelope
- **DELETE /inventory_levels/{id}** ?id -> no content
- **GET /inventory_levels/{id}** ?id -> 200: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url} | no content | 404: error envelope
- **PATCH /inventory_levels/{id}** ?id -> 200: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url}
- **PATCH /inventory_levels/{id}/stock** ?id -> 200: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url}

## Labels

- **GET /labels** -> 200: array of object
- **POST /labels** -> 201: object{id,name,color,active,created_at,updated_at,url,app_url} | 403: error envelope
- **PATCH /labels/{id}** ?id -> 200: object{id,name,color,active,created_at,updated_at,url,app_url}

## Labor Matrices

- **GET /labor_matrices** -> 200: array of object
- **POST /labor_matrices** -> 201: object{id,name,matrix_type,active,tiers,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /labor_matrices/{id}** ?id -> no content
- **PATCH /labor_matrices/{id}** ?id -> 200: object{id,name,matrix_type,active,tiers,created_at,updated_at,url,app_url}

## Labor Rates

- **GET /labor_rates** -> 200: array of object
- **POST /labor_rates** -> 201: object{id,name,rate_cents,cost_per_hour_cents,is_default,active,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /labor_rates/{id}** ?id -> no content
- **PATCH /labor_rates/{id}** ?id -> 200: object{id,name,rate_cents,cost_per_hour_cents,is_default,active,created_at,updated_at,url,app_url}

## Lead Sources

- **GET /lead_sources** -> 200: array of object
- **POST /lead_sources** -> 201: object{id,name,category,system_default,active,position,created_at,updated_at,url,app_url} | 403: error envelope
- **POST /lead_sources/seed_defaults** -> 200: array of object
- **DELETE /lead_sources/{id}** ?id -> no content
- **PATCH /lead_sources/{id}** ?id -> 200: object{id,name,category,system_default,active,position,created_at,updated_at,url,app_url}

## Locations

- **GET /locations/{id}** ?id -> 200: object{id,name,location_type,currency,dock,url,app_url} | 401: error envelope | 403: error envelope

## Notifications

- **GET /notifications** -> 200: array of object | 401: error envelope
- **POST /notifications/bulk_mark_read** -> 200: object{ok,affected}
- **GET /notifications/{id}** ?id -> 200: object{id,trigger_type,title,message_body,read,read_at,created_at,updated_at,url,app_url} | 404: error envelope
- **PATCH /notifications/{id}** ?id -> 200: object{id,trigger_type,title,message_body,read,read_at,created_at,updated_at,url,app_url}

## Orders

- **GET /orders/purchase_orders** ?page,per_page,vendor_id -> 200: array of PurchaseOrder
- **POST /orders/purchase_orders** -> 201: PurchaseOrder | 403: error envelope | 404: error envelope | 500: object
- **DELETE /orders/purchase_orders/{id}** ?id -> no content
- **GET /orders/purchase_orders/{id}** ?id -> 200: PurchaseOrder | no content
- **PATCH /orders/purchase_orders/{id}** ?id -> 200: PurchaseOrder
- **POST /orders/purchase_orders/{purchase_order_id}/cancellations** ?purchase_order_id -> 200: PurchaseOrder
- **GET /orders/return_orders** -> 200: array of ReturnOrder | 403: error envelope | 500: object
- **DELETE /orders/return_orders/{id}** ?id -> no content
- **GET /orders/return_orders/{id}** ?id -> 200: ReturnOrder | 500: object
- **PATCH /orders/return_orders/{id}** ?id -> 200: ReturnOrder | 500: object
- **POST /orders/return_orders/{return_order_id}/refund_completions** ?return_order_id -> 200: ReturnOrder | 500: object
- **GET /orders/sublet_orders** -> 200: array of SubletOrder | 403: error envelope
- **PATCH /orders/sublet_orders/mark_payment_complete** -> 200: SubletOrder
- **GET /orders/sublet_orders/{id}** ?id -> 200: SubletOrder | 403: error envelope

## Packages

- **GET /packages** -> 200: array of object
- **POST /packages** -> 201: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,active,url,app_url,location} | 403: error envelope
- **PATCH /packages/{id}** ?id -> 200: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,active,url,app_url,location} | 403: error envelope
- **POST /packages/{id}/duplicate** ?id -> 201: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,active,url,app_url,location} | 403: error envelope

## Parts Matrices

- **GET /parts_matrices** -> 200: array of object
- **POST /parts_matrices** -> 201: object{id,name,is_default,active,default_multiplier,max_markup_cents,tiers,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /parts_matrices/{id}** ?id -> no content
- **PATCH /parts_matrices/{id}** ?id -> 200: object{id,name,is_default,active,default_multiplier,max_markup_cents,tiers,created_at,updated_at,url,app_url}

## Payments

- **GET /payments** ?method -> 200: array of object | 403: error envelope | 500: object
- **GET /payments/pending** -> 200: array of object | 500: object
- **GET /payments/{id}** ?id -> 200: object{id,amount_cents,currency,method,processor_status,is_refund,is_adjustment,voided,voided_at,processed_at,reference,created_at,updated_at,work_order_id,customer_id,url,app_url,work_order,customer,processed_by,location} | 500: object
- **POST /payments/{id}/cancellation** ?id -> 200: object{id,amount_cents,currency,method,processor_status,is_refund,is_adjustment,voided,voided_at,processed_at,reference,created_at,updated_at,work_order_id,customer_id,url,app_url,work_order,customer,processed_by,location} | 500: object
- **POST /payments/{id}/confirmation** ?id -> 200: object{id,amount_cents,currency,method,processor_status,is_refund,is_adjustment,voided,voided_at,processed_at,reference,created_at,updated_at,work_order_id,customer_id,url,app_url,work_order,customer,processed_by,location} | 500: object
- **POST /payments/{id}/failure** ?id -> 200: object{id,amount_cents,currency,method,processor_status,is_refund,is_adjustment,voided,voided_at,processed_at,reference,created_at,updated_at,work_order_id,customer_id,url,app_url,work_order,customer,processed_by,location} | 500: object

## Preferences

- **GET /preferences** -> 200: object{preferences,url,app_url}

## Profile

- **GET /profile** -> 200: object{id,full_name,email,role,url,app_url}

## Recent Searches

- **GET /recent_searches** -> 200: object{recents}
- **POST /recent_searches** -> 200: object{recents}
- **DELETE /recent_searches/clear** -> no content
- **DELETE /recent_searches/{id}** ?id -> 200: object{recents}

## Reports

- **GET /reports/accounting** -> 200: object{currency,start_date,end_date,tax,payments,outstanding_balances,outstanding_total,url,app_url} | 403: error envelope
- **GET /reports/ar_aging** -> 200: object{rows,totals,credit_rows,credits_total_cents,url,app_url} | 403: error envelope
- **GET /reports/declined_work** -> 200: object{currency,start_date,end_date,rows,total_declined_value_cents,declined_count,category_breakdown,aging_buckets,saved_for_later,total_saved_for_later_cents,url,app_url} | 403: error envelope
- **GET /reports/end_of_day** -> 200: object{date,total_sales_by_method,ar_postings_cents,ar_payments_collected_cents,store_credit_issued_cents,store_credit_applied_cents,over_short_cents,total_cashiered_cents,url,app_url} | 403: error envelope
- **GET /reports/financial** -> 200: object{currency,start_date,end_date,tab,revenue,tax,payments,url,app_url} | 403: error envelope
- **GET /reports/open_work** -> 200: object{currency,start_date,end_date,rows,total_estimated_value_cents,status_summary,stuck_work,saved_for_later,total_saved_for_later_cents,url,app_url} | 403: error envelope
- **GET /reports/parts_purchases** -> 200: object{currency,start_date,end_date,rows,total_spend_cents,po_count,return_count,net_units,url,app_url} | 403: error envelope
- **GET /reports/parts_usage** -> 200: object{currency,start_date,end_date,rows,total_cost_cents,total_revenue_cents,total_gp_cents,url,app_url} | 403: error envelope
- **GET /reports/performance** -> 200: object{view_type,currency,start_date,end_date,kpi_summary,daily_target_cents,revenue_series,car_count_series,aro_series,top_services,retention,url,app_url} | 403: error envelope
- **GET /reports/profit_and_loss** -> 200: object{currency,start_date,end_date,revenue,net_revenue_cents,cogs,gross_profit_cents,operating_expenses,total_operating_expenses_cents,net_income_cents,url,app_url} | 403: error envelope
- **GET /reports/sales_summary** -> 200: object{currency,start_date,end_date,rows,totals,car_count,invoice_count,aro_cents,segment_gp,kpi_summary,url,app_url} | 403: error envelope
- **GET /reports/service_categories** -> 200: object{start_date,end_date,rows,url,app_url} | 403: error envelope
- **GET /reports/store_credit** -> 200: object{total_liability_cents,average_balance_cents,customers_with_balance,all_transactions,url,app_url} | 403: error envelope
- **GET /reports/technician_productivity** -> 200: object{currency,start_date,end_date,rows,totals,period_comparison,url,app_url} | 403: error envelope
- **GET /reports/work_order_profitability** -> 200: object{currency,start_date,end_date,rows,totals,url,app_url} | 403: error envelope

## Scan

- **POST /scan/lookups** -> 200: object{outcome,work_order_id,vehicle_id,customer_id,appointment_id}
- **POST /scan/started_work_orders** -> 200: object{status,work_order_id}
- **POST /scan/vehicles** -> 200: object{status,work_order_id}

## Search

- **GET /search** ?q -> 200: object{html,query,announcement,total_count,groups} | 401: error envelope

## Service Categories

- **GET /service_categories** -> 200: array of ServiceCategory
- **POST /service_categories** -> 201: ServiceCategory | 403: error envelope
- **POST /service_categories/seed_defaults** -> 200: object{created,message}
- **DELETE /service_categories/{id}** ?id -> 200: ServiceCategory | 422: error envelope
- **PATCH /service_categories/{id}** ?id -> 200: ServiceCategory

## Settings

- **GET /settings/account** -> 200: object{id,name,slug,billing_email,website,business_type,tax_id,created_at,updated_at,url,app_url} | 403: error envelope
- **PATCH /settings/account** -> 200: object{id,name,slug,billing_email,website,business_type,tax_id,created_at,updated_at,url,app_url}
- **GET /settings/billing** -> 200: object{billing_email,subscription_status,next_billing_date,work_orders_this_month,url,app_url} | 403: error envelope
- **GET /settings/cash_drawer** -> 200: object{id,name,cash_drawer_enabled,default_starting_float_cents,driveon_station_number,expenses_enabled,tire_management_enabled,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url} | 403: error envelope
- **GET /settings/close_requirements** -> 200: object{id,name,cash_drawer_enabled,default_starting_float_cents,driveon_station_number,expenses_enabled,tire_management_enabled,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /settings/documents** -> 200: object{document_settings,estimate_terms_text,terms_text,payment_instructions,url,app_url} | 401: error envelope
- **GET /settings/driveon** -> 200: object{id,name,cash_drawer_enabled,default_starting_float_cents,driveon_station_number,expenses_enabled,tire_management_enabled,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /settings/expenses** -> 200: object{id,name,cash_drawer_enabled,default_starting_float_cents,driveon_station_number,expenses_enabled,tire_management_enabled,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /settings/labor_templates** -> 200: array of object | 403: error envelope
- **DELETE /settings/labor_templates/{id}** ?id -> no content
- **PATCH /settings/labor_templates/{id}** ?id -> 200: object{id,description,default_hours,usage_count,last_used_at,created_at,updated_at,url,app_url}
- **GET /settings/lead_source_requirements** -> 200: object{id,name,cash_drawer_enabled,default_starting_float_cents,driveon_station_number,expenses_enabled,tire_management_enabled,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /settings/learning_preferences** -> 200: object{preferences,url,app_url}
- **GET /settings/notifications/edit** -> 200: object{email_fallback_enabled,url,app_url}
- **GET /settings/payments** -> 200: object{processor_application_status,processor_onboarded_at,url,app_url} | 403: error envelope
- **GET /settings/phone_numbers** -> 200: object{texting_phone,phones,url,app_url} | 403: error envelope
- **GET /settings/quickbooks** -> 200: object{connected,qbo_company_id,qbo_sync_mode,url,app_url} | 403: error envelope
- **GET /settings/reminders** -> 200: object{id,name,cash_drawer_enabled,default_starting_float_cents,driveon_station_number,expenses_enabled,tire_management_enabled,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /settings/tags** -> 200: object{customer_tags,vehicle_tags}
- **PATCH /settings/tags** -> 200: object{customer_tags,vehicle_tags} | 403: error envelope
- **GET /settings/tire_management** -> 200: object{id,name,cash_drawer_enabled,default_starting_float_cents,driveon_station_number,expenses_enabled,tire_management_enabled,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /settings/trust_levels** -> 200: object{trust_levels,url,app_url}

## Shop Discounts

- **GET /shop_discounts** -> 200: array of object
- **POST /shop_discounts** -> 201: object{id,name,discount_type,amount_cents,percentage,active,category,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /shop_discounts/{id}** ?id -> no content
- **PATCH /shop_discounts/{id}** ?id -> 200: object{id,name,discount_type,amount_cents,percentage,active,category,created_at,updated_at,url,app_url}
- **POST /shop_discounts/{id}/duplicate** ?id -> 201: object{id,name,discount_type,amount_cents,percentage,active,category,created_at,updated_at,url,app_url}

## Shop Fees

- **GET /shop_fees** -> 200: array of object
- **POST /shop_fees** -> 201: object{id,name,fee_type,amount_cents,percentage,active,applies_to,is_taxable,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /shop_fees/{id}** ?id -> no content
- **PATCH /shop_fees/{id}** ?id -> 200: object{id,name,fee_type,amount_cents,percentage,active,applies_to,is_taxable,created_at,updated_at,url,app_url}
- **POST /shop_fees/{id}/duplicate** ?id -> 201: object{id,name,fee_type,amount_cents,percentage,active,applies_to,is_taxable,created_at,updated_at,url,app_url}

## Statements

- **GET /statements/{id}** ?id -> 200: Statement | 404: error envelope
- **GET /statements/{statement_id}/payments** ?statement_id -> 200: array of Statement | 500: object

## Store Credits

- **POST /store_credits/{store_credit_id}/voids** ?store_credit_id -> 200: object{status} | 403: error envelope | 422: error envelope

## Sub Statuses

- **GET /sub_statuses** -> 200: array of object
- **POST /sub_statuses** -> 201: object{id,name,color,status_scope,active,is_default,position,created_at,updated_at,url,app_url} | 403: error envelope
- **PATCH /sub_statuses/{id}** ?id -> 200: object{id,name,color,status_scope,active,is_default,position,created_at,updated_at,url,app_url}

## Sublet Packages

- **GET /sublet_packages** -> 200: array of object
- **POST /sublet_packages** -> 201: object{id,name,description,active,default_fulfillment_status,default_payment_status,default_payment_method,sublet_package_lines_count,location,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /sublet_packages/{id}** ?id -> no content
- **PATCH /sublet_packages/{id}** ?id -> 200: object{id,name,description,active,default_fulfillment_status,default_payment_status,default_payment_method,sublet_package_lines_count,location,created_at,updated_at,url,app_url} | 403: error envelope
- **PATCH /sublet_packages/{id}/deactivate** ?id -> 403: error envelope

## Time Entries

- **POST /time_entries** -> 201: object{status} | 401: error envelope | 422: error envelope
- **PATCH /time_entries/{id}** ?id -> 200: object{status}

## Tire Storage Slots

- **GET /tire_storage_slots** -> 200: array of object
- **POST /tire_storage_slots** -> 201: object{id,slot_label,season,season_label,tire_set_description,stored_at,released_at,currently_stored,storage_fee_cents,fee_type,notes,customer,vehicle,location,created_at,updated_at,url,app_url} | 403: error envelope
- **GET /tire_storage_slots/{id}** ?id -> 200: object{id,slot_label,season,season_label,tire_set_description,stored_at,released_at,currently_stored,storage_fee_cents,fee_type,notes,customer,vehicle,location,created_at,updated_at,url,app_url}
- **POST /tire_storage_slots/{tire_storage_slot_id}/check_outs** ?tire_storage_slot_id -> 200: object{id,slot_label,season,season_label,tire_set_description,stored_at,released_at,currently_stored,storage_fee_cents,fee_type,notes,customer,vehicle,location,created_at,updated_at,url,app_url}

## Tires

- **GET /tires** -> 200: array of object
- **POST /tires** -> 201: object{id,position,position_label,status,size_raw,size_width,size_aspect_ratio,size_rim_diameter,brand,model,load_index,speed_rating,dot_serial,dot_registered,dot_registered_at,purchase_date,tread_depth_new_mm,tread_depth_mm,source,notes,vehicle,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /tires/{id}** ?id -> no content
- **GET /tires/{id}** ?id -> 200: object{id,position,position_label,status,size_raw,size_width,size_aspect_ratio,size_rim_diameter,brand,model,load_index,speed_rating,dot_serial,dot_registered,dot_registered_at,purchase_date,tread_depth_new_mm,tread_depth_mm,source,notes,vehicle,created_at,updated_at,url,app_url}
- **PATCH /tires/{id}** ?id -> 200: object{id,position,position_label,status,size_raw,size_width,size_aspect_ratio,size_rim_diameter,brand,model,load_index,speed_rating,dot_serial,dot_registered,dot_registered_at,purchase_date,tread_depth_new_mm,tread_depth_mm,source,notes,vehicle,created_at,updated_at,url,app_url}

## Users

- **GET /users** ?page,per_page -> 200: array of object | 403: error envelope
- **POST /users** -> 201: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities}
- **GET /users/permission_groups** -> 200: array of object | 403: error envelope
- **POST /users/permission_groups** -> 201: object{id,name,description,role,can_perform_work,can_dispatch_work,can_message_customers,can_manage_technicians,can_override_inspections,can_perform_inspections,can_view_all_active_work_orders,can_close_reopen_work_orders,can_hard_delete_work_orders,can_edit_permissions,can_view_job_board,can_view_metrics,can_view_activity_feed,created_at,updated_at,url,app_url}
- **PATCH /users/permission_groups/{id}** ?id -> 200: object{id,name,description,role,can_perform_work,can_dispatch_work,can_message_customers,can_manage_technicians,can_override_inspections,can_perform_inspections,can_view_all_active_work_orders,can_close_reopen_work_orders,can_hard_delete_work_orders,can_edit_permissions,can_view_job_board,can_view_metrics,can_view_activity_feed,created_at,updated_at,url,app_url}
- **DELETE /users/{id}** ?id -> no content | 403: error envelope
- **GET /users/{id}** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities}
- **PATCH /users/{id}** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities}
- **POST /users/{id}/disable** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities} | 415: object
- **POST /users/{id}/enable** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities} | 415: object
- **GET /users/{id}/qr_code** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities,qr_card}
- **POST /users/{id}/reset_pin** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities,new_pin} | 415: object
- **POST /users/{id}/send_confirmation** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities} | 415: object
- **POST /users/{id}/send_password_reset** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities} | 415: object
- **POST /users/{id}/unlock** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities} | 415: object

## Vehicle Tags

- **GET /vehicle_tags** -> 200: array of object
- **POST /vehicle_tags** -> 201: object{id,name,color,color_hex,color_class}
- **DELETE /vehicle_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class}
- **PATCH /vehicle_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class}

## Vehicles

- **GET /vehicles** ?customer_id,page,per_page -> 200: array of Vehicle | 403: error envelope
- **POST /vehicles** -> 201: Vehicle | 403: error envelope | 422: error envelope
- **GET /vehicles/autocomplete** ?type -> 403: error envelope
- **GET /vehicles/check_duplicate** ?vin -> 200: object{matches} | 403: error envelope
- **GET /vehicles/customer_vehicles** ?customer_id -> 403: error envelope
- **GET /vehicles/lookup** ?query -> 200: array of Vehicle | 403: error envelope
- **GET /vehicles/prefill** ?make,model,vin,year -> 200: Vehicle | 403: error envelope | 404: error envelope
- **GET /vehicles/vin_decode** ?vin -> 200: object{make,model,vin} | 403: error envelope | 404: error envelope
- **DELETE /vehicles/{id}** ?id -> 202: Vehicle | 422: error envelope
- **GET /vehicles/{id}** ?id -> 200: Vehicle | 404: error envelope
- **PATCH /vehicles/{id}** ?id -> 200: Vehicle
- **POST /vehicles/{id}/merges** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count} | 422: error envelope
- **POST /vehicles/{id}/transfers** ?id -> 201: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count} | 403: error envelope | 422: error envelope
- **GET /vehicles/{vehicle_id}/work_orders** ?vehicle_id -> 200: array of WorkOrder

## Vendors

- **GET /vendors** -> 200: array of Vendor | 401: error envelope | 403: error envelope
- **POST /vendors** -> 201: Vendor | 403: error envelope
- **DELETE /vendors/{id}** ?id -> no content
- **GET /vendors/{id}** ?id -> 200: Vendor | 404: error envelope
- **PATCH /vendors/{id}** ?id -> 200: Vendor
- **GET /vendors/{vendor_id}/purchase_orders** ?vendor_id -> 200: array of PurchaseOrder

## Voice Commands

- **DELETE /voice_commands/{id}** ?id -> 200: object{status} | 404: error envelope

## Work Orders

- **GET /work_orders** ?per_page -> 200: array of WorkOrder | 403: error envelope
- **POST /work_orders** -> 201: WorkOrder
- **DELETE /work_orders/{id}** ?id -> no content
- **GET /work_orders/{id}** ?id -> 200: WorkOrder | 401: error envelope
- **PATCH /work_orders/{id}** ?id -> 200: WorkOrder | 422: error envelope
- **PATCH /work_orders/{id}/close** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url} | 422: error envelope
- **PATCH /work_orders/{id}/close_as_paid** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url}
- **PATCH /work_orders/{id}/close_zero** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url}
- **DELETE /work_orders/{id}/courtesy_car_assignment** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}/courtesy_car_assignment** ?id -> 200: WorkOrder
- **POST /work_orders/{id}/courtesy_car_assignment** ?id -> 200: WorkOrder | 403: error envelope
- **PATCH /work_orders/{id}/decline_all** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url}
- **GET /work_orders/{id}/declined_services** ?id -> 200: array of object
- **DELETE /work_orders/{id}/hard_delete** ?id -> 202: WorkOrder | 422: error envelope
- **PATCH /work_orders/{id}/post_to_account** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}/reopen** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url} | 403: error envelope
- **PATCH /work_orders/{id}/return_to_board** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url}
- **PATCH /work_orders/{id}/save_for_later** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url}
- **PATCH /work_orders/{id}/send_estimate** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}/send_invoice_summary** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}/send_reminder** ?id -> 200: WorkOrder
- **GET /work_orders/{id}/service_history** ?id -> 200: array of object
- **PATCH /work_orders/{id}/toggle_waiting_for_customer** ?id -> 200: WorkOrder | no content
- **POST /work_orders/{work_order_id}/authorization** ?work_order_id -> 200: WorkOrder | 403: error envelope | 422: error envelope
- **POST /work_orders/{work_order_id}/authorization/update_decisions** ?work_order_id -> 200: WorkOrder
- **GET /work_orders/{work_order_id}/concerns** ?work_order_id -> 200: array of WorkOrder
- **GET /work_orders/{work_order_id}/estimate** ?work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/fee_exclusions** ?work_order_id -> 201: WorkOrder | 403: error envelope | 422: error envelope
- **DELETE /work_orders/{work_order_id}/fee_exclusions/{id}** ?id,work_order_id -> no content
- **GET /work_orders/{work_order_id}/inspection** ?work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/labels** ?work_order_id -> 200: WorkOrder | 403: error envelope | no content
- **DELETE /work_orders/{work_order_id}/labels/{id}** ?id,work_order_id -> 200: WorkOrder
- **GET /work_orders/{work_order_id}/parts** ?work_order_id -> 200: WorkOrder
- **GET /work_orders/{work_order_id}/payments** ?work_order_id -> 200: WorkOrder | 500: object
- **POST /work_orders/{work_order_id}/payments** ?work_order_id -> 201: WorkOrder | 403: error envelope | 422: error envelope | 500: object
- **DELETE /work_orders/{work_order_id}/payments/reverse_ar** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url}
- **POST /work_orders/{work_order_id}/payments/send_to_ar** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url}
- **POST /work_orders/{work_order_id}/refunds** ?work_order_id -> 201: WorkOrder | 500: object
- **POST /work_orders/{work_order_id}/services** ?work_order_id -> 201: WorkOrder | no content | 422: error envelope
- **PATCH /work_orders/{work_order_id}/services/reorder** ?work_order_id -> no content
- **DELETE /work_orders/{work_order_id}/services/{id}** ?id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{id}** ?id,work_order_id -> 200: WorkOrder | no content
- **PATCH /work_orders/{work_order_id}/services/{id}/acknowledge_parts** ?id,work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/services/{id}/add_line_item** ?id,work_order_id -> 201: WorkOrder | 422: error envelope
- **POST /work_orders/{work_order_id}/services/{id}/add_package** ?id,work_order_id -> 200: WorkOrder | no content
- **GET /work_orders/{work_order_id}/services/{id}/adjust_time** ?id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{id}/adjust_time** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/apply_discount** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/complete_service** ?id,work_order_id -> 200: WorkOrder | 422: error envelope
- **POST /work_orders/{work_order_id}/services/{id}/duplicate** ?id,work_order_id -> 201: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/pause** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/publish** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/reset_completion** ?id,work_order_id -> 200: WorkOrder | 403: error envelope
- **PATCH /work_orders/{work_order_id}/services/{id}/revive** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/start** ?id,work_order_id -> 200: WorkOrder | 403: error envelope
- **PATCH /work_orders/{work_order_id}/services/{id}/toggle_labor_completion** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/toggle_labor_tax** ?id,work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/services/{id}/unauthorize** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/update_category** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/update_pricing_mode** ?id,work_order_id -> 200: WorkOrder | 422: error envelope
- **PATCH /work_orders/{work_order_id}/services/{id}/update_service_technician** ?id,work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/services/{service_id}/line_items** ?service_id,work_order_id -> 201: WorkOrder
- **DELETE /work_orders/{work_order_id}/services/{service_id}/line_items/{id}** ?id,service_id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}** ?id,service_id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/add_to_inventory** ?id,service_id,work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/duplicate** ?id,service_id,work_order_id -> 201: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/pull** ?id,service_id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/refresh_price** ?id,service_id,work_order_id -> 200: WorkOrder | 422: error envelope
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_pull** ?id,service_id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_return** ?id,service_id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/update_part_status** ?id,service_id,work_order_id -> 200: WorkOrder | 422: error envelope
- **GET /work_orders/{work_order_id}/summary/activity** ?category,work_order_id -> 200: array of WorkOrder | 403: error envelope
- **GET /work_orders/{work_order_id}/summary/appointments** ?work_order_id -> 200: array of WorkOrder
- **GET /work_orders/{work_order_id}/summary/authorization_logs** ?work_order_id -> 200: array of WorkOrder
- **GET /work_orders/{work_order_id}/summary/vehicle_history** ?work_order_id -> 200: array of WorkOrder
- **POST /work_orders/{work_order_id}/voids** ?work_order_id -> 201: WorkOrder | 500: object
- **GET /work_orders/{work_order_id}/wip** ?work_order_id -> 200: WorkOrder | no content


## Schemas

- **WorkOrder**: {type*:string, id*:integer, work_order_number*:integer, status*:string, intake_method*:string, scheduled_for*:string, authorized*:boolean, paid*:boolean, created_at*:string, updated_at*:string, closed_at*:string, location_id*:integer, service_advisor_id*:integer, assigned_technician_id*:integer, sub_status_type_id*:integer, payer_customer_id*:integer, vehicle_arrived_at*:string, work_order_services_count*:integer, inspection_reports_count*:integer, customer*:object{id,full_name,url}, vehicle*:object{id,make,model,year,vin,url}, totals*:object{subtotal_cents,tax_cents,total_cents,paid_cents,remaining_cents,currency}, url*:string, app_url*:string, location*:object{id,name,url}, odometer_in*:integer, odometer_out*:integer, odometer_unit*:string, authorized_at*:string, authorized_total_cents*:integer, customer_notified*:boolean, customer_notified_ready*:boolean, ready_for_pickup_at*:string, completed_at*:string, declined_at*:string, decline_reason*:string, discount_cents*:integer, fees_cents*:integer, parts_cents*:integer, labor_cents*:integer, tires_cents*:integer, subcontracts_cents*:integer, credit_balance_cents*:integer, saved_for_later*:boolean, closure_reason*:string, closure_reason_notes*:string, notes*:string, purchase_order_number*:string, return_method*:string, return_method_notes*:string, vehicle_keys_location*:string, vehicle_location*:string, customer_visit_count*:integer, customer_total_spend_cents*:integer, average_ticket_cents*:integer, activity_total*:integer, recent_activities*:array of any, services_url*:string, payments_url*:string, wip_url*:string, inspection_url*:string, parts_url*:string, concerns_url*:string, service_history_url*:string, declined_services_url*:string, activity_url*:string, vehicle_history_url*:string, appointments_url*:string, authorization_logs_url*:string, payer_customer:object{id,full_name,url}}
- **BroadcastCampaign**: {id*:integer, type*:string, name*:string, status*:string, sms_body*:string, filters*:object, recipient_count*:integer, sent_count*:integer, failed_count*:integer, progress_percentage*:integer, sent_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, creator*:object{id,name,url}, location*:object{id,name,url}}
- **Customer**: {type*:string, id*:integer, full_name*:string, company_name*:string, first_name*:string, last_name*:string, fleet_identifier*:string, marketing_opt_in*:boolean, tax_exempt*:boolean, vehicles_count*:integer, emails_count*:integer, phones_count*:integer, vehicles_url*:string, work_orders_url*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, emails*:array of object, phones*:array of object, addresses*:array of object, outstanding_balance_cents*:integer, total_revenue_cents*:integer, store_credit_cents*:integer, last_visit_at*:string, statements_count*:integer, currency*:string}
- **Driver**: {id*:integer, full_name*:string, phone*:string, email*:string, customer*:object{id,full_name,url}, work_orders_count*:integer, work_orders_url*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Statement**: {id*:integer, amount_cents*:integer, currency*:string, method*:string, processor_status*:string, is_refund*:boolean, is_adjustment*:boolean, voided*:boolean, voided_at*:string, processed_at*:string, reference*:string, created_at*:string, updated_at*:string, work_order_id*:integer, customer_id*:integer, url*:string, app_url*:string, work_order*:object{id,url}, customer*:object{id,full_name,url}, processed_by*:object{id,full_name,url}, location*:object{id,name,url}}
- **Vehicle**: {type*:string, id*:integer, make*:string, model*:string, year*:integer, submodel*:string, body_style*:string, engine*:string, vin*:string, license_plate*:string, license_plate_state*:string, license_plate_country*:string, drivetrain*:string, transmission*:string, color*:string, vehicle_type*:string, unit_number*:string, fleet_identifier*:string, production_date*:string, annual_safety_expires_at*:string, notes*:string, odometer*:object{reading,unit}, work_orders_count*:integer, work_orders_url*:string, customer*:object{id,full_name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, last_serviced_at*:string, lifetime_revenue_cents*:integer, open_work_orders_count*:integer, appointments_count*:integer}
- **InspectionReport**: {id*:integer, type*:string, name*:string, status*:string, work_order_id*:integer, quick_finding*:boolean, published*:boolean, completed*:boolean, items_count*:integer, checked_count*:integer, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, groups*:array of any}
- **Inspection**: {id*:integer, type*:string, name*:string, description*:string, active*:boolean, is_default*:boolean, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, groups*:array of object}
- **PurchaseOrder**: {id*:integer, type*:string, po_number*:integer, status*:string, order_method*:string, payment_method*:string, fulfillment_method*:string, tracking_number*:any, vendor_invoice_number*:any, vendor_invoice_received_at*:string, notes*:string, freight_cost_cents*:integer, freight_cost_currency*:string, subtotal_cents*:string, total_cents*:string, core_charges_cents*:integer, line_items_count*:integer, ordered_at*:string, received_at*:string, payment_due_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, vendor*:object{id,name,url}, creator*:object{id,name,url}, location*:object{id,name,url}, line_items*:array of object}
- **ReturnOrder**: {id*:integer, type*:string, return_number*:integer, status*:string, credit_method*:any, reason_code*:any, rma_number*:any, is_warranty_claim*:boolean, restocking_fee_cents*:integer, shipping_fee_cents*:integer, notes*:string, line_items_count*:integer, refund_completed_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, vendor*:object{id,name,url}, work_order*:object{id,number,url}, creator*:object{id,name,url}, location*:object{id,name,url}, line_items*:array of object}
- **SubletOrder**: {id*:integer, type*:string, sublet_number*:integer, title*:string, payment_status*:string, payment_method*:any, total_cents*:integer, total_cost_cents*:integer, margin_cents*:integer, margin_percentage*:integer, sent_to_ap*:boolean, vendor_paid_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, vendor*:object{id,name,url}, work_order*:object{id,number,url}, work_order_service*:object{id,name}, location*:object{id,name,url}}
- **ServiceCategory**: {id*:integer, name*:string, description*:string, service_type*:string, icon*:string, color*:string, active*:boolean, position*:integer, canonical*:boolean, canonical_key*:string, job_count*:integer, url*:string, app_url*:string}
- **Vendor**: {id*:integer, name*:string, vendor_type*:string, payment_terms*:string, active*:boolean, phone*:string, email*:string, website*:string, account_number*:string, notes*:string, quick_order*:boolean, order_url_template*:string, catalog_url_template*:string, location*:object{id,name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Error**: {code*:string, message*:string, field_errors*:object}
- **UpdateAiSuggestionRequest**: {status*:string}
- **CreateAppointmentRequest**: {appointment*:object{starts_at,customer_id,vehicle_id,intake_method}}
- **UpdateAppointmentRequest**: {appointment*:object{customer_concern}}
- **CreateAppointmentsApprovalRequest**: {}
- **CreateAppointmentsCancellationRequest**: {}
- **CreateAppointmentsFollowUpRequest**: {appointment*:object{starts_at,follow_up_reason}}
- **CreateAppointmentsRejectionRequest**: {}
- **CreateAppointmentsVehicleReconciliationRequest**: {vehicle_action*:string, vehicle_id*:integer}
- **CreateAppointmentsWorkOrderRequest**: {}
- **CreateCampaignRequest**: {broadcast_campaign*:object{name,sms_body}}
- **DuplicateCampaignRequest**: {}
- **SendCampaignRequest**: {}
- **UpdateCashDrawerBannerRequest**: {dismissed*:boolean}
- **CreateCatalogCleanupsApplicationRequest**: {category*:string}
- **CreateConversationsBulkMarkReadRequest**: {}
- **CreateConversationsCustomerLinkRequest**: {customer_id*:integer}
- **CreateConversationsIgnoreRequest**: {}
- **CreateConversationsMessageRequest**: {message*:object{body,channel}}
- **CreateConversationsMessagesResendRequest**: {}
- **UpdateConversationRequest**: {status*:string}
- **UpdateCoreTaxRuleRequest**: {core_tax_rule*:object{tax_core_charge}}
- **CreateCounterSaleRequest**: {}
- **UpdateCounterSaleRequest**: {counter_sale*:object{notes}}
- **CreateCustomerTagRequest**: {name*:string}
- **UpdateCustomerTagRequest**: {name*:string}
- **CreateCustomerRequest**: {customer*:object{first_name,last_name,company_name,fleet_identifier,billing_terms,credit_limit_cents,tax_exempt,tax_exempt_number,notes,marketing_opt_in,discount_percent,po_required,tag_ids,emails_attributes,phones_attributes,addresses_attributes}}
- **CreateCustomersExportRequest**: {}
- **CreateDriverRequest**: {driver*:object{full_name,phone}}
- **UpdateDriverRequest**: {driver*:object{full_name}}
- **UpdateCustomerRequest**: {customer*:object{emails_attributes,first_name,last_name,company_name,fleet_identifier,billing_terms,credit_limit_cents,tax_exempt,notes,marketing_opt_in,discount_percent,po_required,phones_attributes,addresses_attributes}}
- **MergeCustomerRequest**: {source_customer_id*:integer}
- **CreateExpenseRequest**: {expense*:object{payee,category,description,amount,expense_date,payment_method}}
- **UpdateExpenseRequest**: {expense*:object{description}}
- **CreateInspectionReportRequest**: {}
- **CompleteInspectionReportRequest**: {}
- **MarkAllInspectionReportRequest**: {status*:string, group_name*:string}
- **PublishInspectionReportRequest**: {}
- **ReassignInspectionReportRequest**: {user_id*:integer}
- **ReopenInspectionReportRequest**: {}
- **ResetInspectionReportRequest**: {}
- **RetryInspectionReportRecordingRequest**: {}
- **UnpublishInspectionReportRequest**: {}
- **CreateInspectionRequest**: {inspection*:object{name,active}}
- **UpdateInspectionRequest**: {inspection*:object{name}}
- **RemoveDefaultInspectionRequest**: {}
- **SetDefaultInspectionRequest**: {}
- **ToggleInspectionRequest**: {}
- **CreateInventoryLevelRequest**: {part*:object{part_number,description,brand,part_type,stocked,initial_quantity,cost,sell}}
- **UpdateInventoryLevelRequest**: {part*:object{description}}
- **UpdateInventoryLevelsStockRequest**: {inventory_level*:object{on_hand,reason}}
- **CreateLabelRequest**: {label*:object{name,color}}
- **UpdateLabelRequest**: {label*:object{active}}
- **CreateLaborMatriceRequest**: {labor_matrix*:object{name,matrix_type,active}}
- **UpdateLaborMatriceRequest**: {labor_matrix*:object{name}}
- **CreateLaborRateRequest**: {labor_rate*:object{name,rate,is_default,active}}
- **UpdateLaborRateRequest**: {labor_rate*:object{active}}
- **CreateLeadSourceRequest**: {lead_source*:object{name,category}}
- **CreateLeadSourcesSeedDefaultRequest**: {}
- **UpdateLeadSourceRequest**: {lead_source*:object{active}}
- **CreateNotificationsBulkMarkReadRequest**: {notification_ids:array of integer}
- **UpdateNotificationRequest**: {read*:boolean}
- **CreateOrdersPurchaseOrderRequest**: {purchase_order*:object{vendor_id,payment_method,fulfillment_method,line_items}}
- **UpdateOrdersPurchaseOrderRequest**: {purchase_order*:object{tracking_number}}
- **CreateOrdersPurchaseOrdersCancellationRequest**: {}
- **UpdateOrdersReturnOrderRequest**: {return_order*:object{rma_number}}
- **CreateOrdersReturnOrdersRefundCompletionRequest**: {}
- **UpdateOrdersSubletOrdersMarkPaymentCompleteRequest**: {sublet_order_ids*:array of integer, payment_method*:string}
- **CreatePackageRequest**: {package*:object{name,description}}
- **UpdatePackageRequest**: {package*:object{active}}
- **CreatePackagesDuplicateRequest**: {}
- **CreatePartsMatriceRequest**: {parts_matrix*:object{name,is_default,active}}
- **UpdatePartsMatriceRequest**: {parts_matrix*:object{name}}
- **CreatePaymentsCancellationRequest**: {}
- **CreatePaymentsConfirmationRequest**: {}
- **CreatePaymentsFailureRequest**: {}
- **CreateRecentSearcheRequest**: {recent_search*:object{query,result_type,result_id,label}}
- **CreateScanLookupRequest**: {code*:string, type*:string}
- **CreateScanStartedWorkOrderRequest**: {outcome*:string, vehicle_id*:integer}
- **CreateScanVehicleRequest**: {customer*:object{first_name,last_name,phone}, vehicle*:object{vin,year,make,model}}
- **CreateServiceCategoryRequest**: {service_category*:object{name,service_type,icon}}
- **SeedDefaultsServiceCategoriesRequest**: {}
- **UpdateServiceCategoryRequest**: {service_category*:object{active,name,position}}
- **UpdateSettingsAccountRequest**: {account*:object{name}}
- **UpdateSettingsLaborTemplateRequest**: {labor_template*:object{default_hours}}
- **UpdateTagsRequest**: {customer_tags:array of object, vehicle_tags:array of object}
- **CreateShopDiscountRequest**: {shop_discount_config*:object{name,discount_type,percentage,active}}
- **UpdateShopDiscountRequest**: {shop_discount_config*:object{active}}
- **CreateShopDiscountsDuplicateRequest**: {}
- **CreateShopFeeRequest**: {shop_fee_config*:object{name,fee_type,amount,active}}
- **UpdateShopFeeRequest**: {shop_fee_config*:object{active}}
- **CreateShopFeesDuplicateRequest**: {}
- **CreateStoreCreditsVoidRequest**: {}
- **CreateSubStatuseRequest**: {sub_status_type*:object{name,color,status_scope}}
- **UpdateSubStatuseRequest**: {sub_status_type*:object{active}}
- **CreateSubletPackageRequest**: {sublet_package*:object{name,description,active}}
- **UpdateSubletPackageRequest**: {sublet_package*:object{name}}
- **CreateTimeEntrieRequest**: {type*:string, work_order_service_id*:integer}
- **UpdateTimeEntrieRequest**: {status*:string}
- **CreateTireStorageSlotRequest**: {tire_storage_slot*:object{vehicle_id,customer_id,slot_label,season,stored_at}}
- **CreateTireStorageSlotsCheckOutRequest**: {}
- **CreateTireRequest**: {tire*:object{vehicle_id,position,status,size_raw,brand}}
- **UpdateTireRequest**: {tire*:object{brand}}
- **CreateUserRequest**: {user*:object{full_name,email,role,home_location_id,location_ids,can_perform_work}}
- **CreatePermissionGroupRequest**: {permission_group*:object{name,can_perform_work}}
- **UpdatePermissionGroupRequest**: {permission_group*:object{name}}
- **UpdateUserRequest**: {user*:object{full_name}}
- **CreateUsersDisableRequest**: {}
- **CreateUsersEnableRequest**: {}
- **CreateUsersResetPinRequest**: {}
- **CreateUsersSendConfirmationRequest**: {}
- **CreateUsersSendPasswordResetRequest**: {}
- **CreateUsersUnlockRequest**: {}
- **CreateVehicleTagRequest**: {name*:string}
- **UpdateVehicleTagRequest**: {name*:string}
- **CreateVehicleRequest**: {vehicle*:object{customer_id,vin,year,make,model,submodel,body_style,engine,transmission,drivetrain,color,license_plate,license_plate_state,odometer_reading,odometer_unit,unit_number,fleet_identifier,notes,production_date,vehicle_tag_ids}}
- **UpdateVehicleRequest**: {vehicle*:object{make,model,year,vin,submodel,body_style,engine,transmission,drivetrain,color,license_plate,license_plate_state,odometer_reading,odometer_unit,notes}}
- **MergeVehicleRequest**: {source_vehicle_id*:integer}
- **TransferVehicleRequest**: {customer_id*:integer, mode*:string}
- **CreateVendorRequest**: {vendor*:object{name,vendor_type,payment_terms}}
- **UpdateVendorRequest**: {vendor*:object{name}}
- **CreateWorkOrderRequest**: {work_order*:object{customer_id,vehicle_id}}
- **UpdateWorkOrderRequest**: {work_order*:object{vehicle_arrived_at,intake_method,sub_status_type_id,payer_customer_id}}
- **CloseWorkOrderRequest**: {closure_type*:string, closure_reason*:string}
- **CloseWorkOrderAsPaidRequest**: {}
- **CloseWorkOrderZeroRequest**: {}
- **UpdateWorkOrdersCourtesyCarAssignmentRequest**: {action_type*:string}
- **CreateWorkOrdersCourtesyCarAssignmentRequest**: {vehicle_id*:integer}
- **DeclineAllWorkOrderServicesRequest**: {decline_reason*:string}
- **UpdateWorkOrdersPostToAccountRequest**: {}
- **ReopenWorkOrderRequest**: {}
- **ReturnWorkOrderToBoardRequest**: {}
- **SaveWorkOrderForLaterRequest**: {}
- **UpdateWorkOrdersSendEstimateRequest**: {}
- **UpdateWorkOrdersSendInvoiceSummaryRequest**: {}
- **UpdateWorkOrdersSendReminderRequest**: {}
- **UpdateWorkOrdersToggleWaitingForCustomerRequest**: {}
- **CreateWorkOrderAuthorizationRequest**: {authorization_method*:string, service_ids*:array of integer, service_decisions*:object{1047559673}}
- **UpdateWorkOrderAuthorizationDecisionsRequest**: {service_decision_reasons*:object{1047559673}}
- **CreateWorkOrdersFeeExclusionRequest**: {work_order_fee_exclusion*:object{shop_fee_config_id}}
- **CreateWorkOrdersLabelRequest**: {label_id*:integer}
- **CreateWorkOrderPaymentRequest**: {payment*:object{method,amount_cents}}
- **SendWorkOrderPaymentToArRequest**: {}
- **CreateWorkOrdersRefundRequest**: {refund*:object{payment_id,amount,reason}}
- **CreateWorkOrdersServiceRequest**: {work_order_service*:object{name,service_type}, package_id:integer}
- **UpdateWorkOrdersServicesReorderRequest**: {service_ids*:array of integer}
- **UpdateWorkOrdersServiceRequest**: {work_order_service*:object{pricing_mode,name,position}}
- **UpdateWorkOrdersServicesAcknowledgePartsRequest**: {}
- **CreateWorkOrdersServicesAddLineItemRequest**: {item_type*:string, name*:string, amount_cents*:integer}
- **CreateWorkOrdersServicesAddPackageRequest**: {package_id*:integer}
- **UpdateWorkOrdersServicesAdjustTimeRequest**: {hours*:integer, minutes*:integer}
- **UpdateWorkOrdersServicesApplyDiscountRequest**: {discount*:object{type,scope,value_cents}}
- **UpdateWorkOrdersServicesCompleteServiceRequest**: {}
- **CreateWorkOrdersServicesDuplicateRequest**: {}
- **UpdateWorkOrdersServicesPauseRequest**: {}
- **UpdateWorkOrdersServicesPublishRequest**: {}
- **UpdateWorkOrdersServicesResetCompletionRequest**: {}
- **UpdateWorkOrdersServicesReviveRequest**: {}
- **UpdateWorkOrdersServicesStartRequest**: {}
- **UpdateWorkOrdersServicesToggleLaborCompletionRequest**: {line_item_id*:integer}
- **UpdateWorkOrdersServicesToggleLaborTaxRequest**: {}
- **CreateWorkOrdersServicesUnauthorizeRequest**: {}
- **UpdateWorkOrdersServicesUpdateCategoryRequest**: {category_id*:integer}
- **UpdateWorkOrdersServicesUpdatePricingModeRequest**: {work_order_service*:object{pricing_mode}}
- **UpdateWorkOrdersServicesUpdateServiceTechnicianRequest**: {work_order_service*:object{technician_id}}
- **CreateWorkOrdersServicesLineItemRequest**: {work_order_line_item*:object{item_type,description,unit_price,quantity,hours,labor_rate_id}}
- **UpdateWorkOrdersServicesLineItemRequest**: {work_order_line_item*:object{description}}
- **AddWorkOrderServiceLineItemToInventoryRequest**: {}
- **DuplicateWorkOrderServiceLineItemRequest**: {}
- **PullWorkOrderServiceLineItemRequest**: {}
- **RefreshWorkOrderServiceLineItemPriceRequest**: {}
- **UndoPullWorkOrderServiceLineItemRequest**: {}
- **UndoReturnWorkOrderServiceLineItemRequest**: {}
- **UpdateWorkOrderServiceLineItemPartStatusRequest**: {part_status*:string}
- **CreateWorkOrdersVoidRequest**: {payment_id*:integer}

