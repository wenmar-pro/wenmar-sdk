# Wenmar Pro API — Compact Reference

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

Base URL: `https://app.wenmarpro.com`
Auth: `Authorization: Bearer <token>`
Responses: bare objects/arrays, no envelope. Errors: `{ "error": { code, message, details } }`.

## Account

- **DELETE /account** -> no content | 403: error envelope | 422: error envelope
- **GET /account** -> 200: object{id,name,slug,billing_email,website,business_type,tax_id,created_at,updated_at,url,app_url,formatted_join_code,station_login_url,deletion_scheduled_at,locations} | 401: error envelope
- **PATCH /account** -> 200: object{id,name,slug,billing_email,website,business_type,tax_id,created_at,updated_at,url,app_url,formatted_join_code,station_login_url,deletion_scheduled_at,locations} | 403: error envelope
- **GET /account/billing** -> 200: object{billing_email,subscription_status,next_billing_date,work_orders_this_month,url,app_url} | 403: error envelope
- **GET /account/capabilities** -> 200: object{tier,tier_display,capabilities,limits}
- **GET /account/driveon** -> 200: object{id,name,default_starting_float_cents,driveon_station_number,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url} | 403: error envelope
- **GET /account/payments** -> 200: object{processor_application_status,processor_onboarded_at,url,app_url} | 403: error envelope
- **GET /account/phone_numbers** -> 200: object{texting_phone,phones,url,app_url} | 403: error envelope
- **GET /account/quickbooks** -> 200: object{connected,qbo_company_id,qbo_sync_mode,url,app_url} | 403: error envelope
- **GET /account/station_link** -> 200: object{join_code,formatted_join_code,station_login_url,qr_code_url}
- **POST /account/station_link/regenerate** -> 200: object{id,name,slug,billing_email,website,business_type,tax_id,created_at,updated_at,url,app_url,formatted_join_code,station_login_url,deletion_scheduled_at,locations}

## Ai Suggestions

- **PATCH /ai_suggestions/{id}** ?id -> 200: object{id,status} | 403: error envelope | 422: error envelope

## Appointments

- **GET /appointments** ?per_page,q,status -> 200: array of object
- **POST /appointments** -> 201: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url} | 403: error envelope | 422: error envelope
- **GET /appointments/available_slots** ?date,duration_minutes -> 200: array of object
- **DELETE /appointments/{id}** ?id -> no content
- **GET /appointments/{id}** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url} | no content
- **PATCH /appointments/{id}** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url}
- **POST /appointments/{id}/approvals** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url}
- **POST /appointments/{id}/cancellations** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url}
- **POST /appointments/{id}/follow_ups** ?id -> 201: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url}
- **POST /appointments/{id}/rejections** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url,reconcile_vehicle_url}
- **POST /appointments/{id}/vehicle_reconciliations** ?id -> 200: object{id,type,status,appointment_type,appointment_source,intake_method,all_day,starts_at,ends_at,estimated_duration_minutes,customer_name,customer_email,customer_phone,customer_concern,follow_up_reason,year,make,model,submodel,vin,license_plate,url,app_url,customer,vehicle,service_advisor,work_order,location,created_at,updated_at,customer_confirmed,confirmation_sent_at,reminder_sent_at,customer_arrived_at,customer_initiated,rescheduled_from_id,latest_reschedule_id,messages_count,reschedules_count,approve_url,reject_url,cancel_url,follow_up_url,work_order_url}
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
- **POST /conversations** -> 201: object{id,status,reply_state,channel,from_number,from_email,last_message_preview,messages_count,has_failed_message,unread_count,last_message_at,customer_last_read_at,driver_last_read_at,oldest_unanswered_inbound_at,created_at,updated_at,customer,messages_url,url,app_url} | 422: error envelope
- **POST /conversations/bulk_mark_read** -> 200: object{ok,affected}
- **POST /conversations/{conversation_id}/customer_links** ?conversation_id -> 200: object{id,status,reply_state,channel,customer,unread_count,messages_url,created_at,updated_at,url,app_url,from_number,from_email,last_message_preview,messages_count,has_failed_message,last_message_at,customer_last_read_at,driver_last_read_at,oldest_unanswered_inbound_at}
- **POST /conversations/{conversation_id}/ignores** ?conversation_id -> 200: object{id,status,reply_state,channel,customer,unread_count,messages_url,created_at,updated_at,url,app_url,from_number,from_email,last_message_preview,messages_count,has_failed_message,last_message_at,customer_last_read_at,driver_last_read_at,oldest_unanswered_inbound_at}
- **GET /conversations/{conversation_id}/messages** ?conversation_id -> 200: array of object
- **POST /conversations/{conversation_id}/messages** ?conversation_id -> 201: object{id,conversation_id,direction,channel,status,body,sender,work_order_id,statement_id,appointment_id,read_at,created_at,recipient_phone,recipient_email,failure_reason,attachment_count,sent_at,delivered_at,failed_at,updated_at,conversation_url,url,app_url}
- **GET /conversations/{id}** ?id -> 200: object{id,status,reply_state,channel,customer,unread_count,messages_url,created_at,updated_at,url,app_url,from_number,from_email,last_message_preview,messages_count,has_failed_message,last_message_at,customer_last_read_at,driver_last_read_at,oldest_unanswered_inbound_at}
- **PATCH /conversations/{id}** ?id -> 200: object{id,status,reply_state,channel,customer,unread_count,messages_url,created_at,updated_at,url,app_url,from_number,from_email,last_message_preview,messages_count,has_failed_message,last_message_at,customer_last_read_at,driver_last_read_at,oldest_unanswered_inbound_at}

## Core Tax Rules

- **GET /core_tax_rules** -> 200: array of object | 403: error envelope
- **PATCH /core_tax_rules/{id}** ?id -> 200: object{id,province_code,tax_core_charge,tax_core_credit,notes,created_at,updated_at,url,app_url}

## Counter Sales

- **GET /counter_sales** -> 200: array of object
- **POST /counter_sales** -> 201: object{id,counter_sale_number,status,walk_in_name,notes,subtotal_cents,tax_total_cents,total_cents,paid_cents,remaining_cents,paid,currency,line_items_count,processed_by,location,created_at,updated_at,url,app_url,reopen_url} | 403: error envelope
- **GET /counter_sales/{counter_sale_id}/line_items/brands** ?counter_sale_id -> 403: error envelope
- **GET /counter_sales/{id}** ?id -> 200: object{id,counter_sale_number,status,walk_in_name,notes,subtotal_cents,tax_total_cents,total_cents,paid_cents,remaining_cents,paid,currency,line_items_count,processed_by,location,created_at,updated_at,url,app_url,reopen_url}
- **PATCH /counter_sales/{id}** ?id -> 200: object{id,counter_sale_number,status,walk_in_name,notes,subtotal_cents,tax_total_cents,total_cents,paid_cents,remaining_cents,paid,currency,line_items_count,processed_by,location,created_at,updated_at,url,app_url,reopen_url}
- **PATCH /counter_sales/{id}/reopen** ?id -> 200: object{id,counter_sale_number,status,walk_in_name,notes,subtotal_cents,tax_total_cents,total_cents,paid_cents,remaining_cents,paid,currency,line_items_count,processed_by,location,created_at,updated_at,url,app_url,reopen_url} | 403: error envelope | 422: error envelope

## Courtesy Cars

- **GET /courtesy_cars** -> 200: array of object | 403: error envelope
- **POST /courtesy_cars** -> 201: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,status,trashed_at,work_orders_url,customer,created_at,updated_at,url,app_url,location} | 403: error envelope | 422: error envelope
- **DELETE /courtesy_cars/{id}** ?id -> no content | 422: error envelope
- **GET /courtesy_cars/{id}** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,status,trashed_at,work_orders_url,customer,created_at,updated_at,url,app_url,location}
- **PATCH /courtesy_cars/{id}** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,status,trashed_at,work_orders_url,customer,created_at,updated_at,url,app_url,location}

## Current Location

- **GET /current_location** -> 200: object{id,name,slug,location_type,currency,time_zone,country,address,city,state,postal_code,contact_email,dock,url,app_url}
- **PATCH /current_location** -> 200: object{id,name,slug,location_type,currency,time_zone,country,address,city,state,postal_code,contact_email,dock,url,app_url} | 403: error envelope

## Customer Tags

- **GET /customer_tags** ?status -> 200: array of object
- **POST /customer_tags** -> 201: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at} | 403: error envelope
- **GET /customer_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at}
- **PATCH /customer_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at}
- **PATCH /customer_tags/{id}/archive** ?id -> 200: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at}
- **PATCH /customer_tags/{id}/restore** ?id -> 200: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at}
- **PATCH /customer_tags/{id}/trash** ?id -> 200: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at}

## Customers

- **GET /customers** ?customer_tag_id,has_balance,has_vehicle,last_visit_months,page,per_page,q,status,type -> 200: array of Customer | 401: error envelope
- **POST /customers** -> 201: Customer | 403: error envelope | 422: error envelope
- **GET /customers/1043910089/merge** -> no content
- **GET /customers/check_duplicate** ?email,first_name,last_name,phone -> 200: object{matches} | 403: error envelope
- **POST /customers/export** -> 200: Customer | 403: error envelope
- **GET /customers/export/{id}/download** ?id -> no content | 404: error envelope
- **GET /customers/lookup** ?id,query -> 200: array of Customer | 403: error envelope
- **GET /customers/{customer_id}/drivers** ?customer_id -> 200: array of Driver | 401: error envelope
- **POST /customers/{customer_id}/drivers** ?customer_id -> 201: Driver | 401: error envelope | 422: error envelope
- **DELETE /customers/{customer_id}/drivers/{id}** ?customer_id,id -> no content | 401: error envelope
- **GET /customers/{customer_id}/drivers/{id}** ?customer_id,id -> 200: Driver
- **PATCH /customers/{customer_id}/drivers/{id}** ?customer_id,id -> 200: Driver
- **GET /customers/{customer_id}/statements** ?customer_id -> 200: array of Statement
- **GET /customers/{customer_id}/vehicles** ?customer_id -> 200: array of Vehicle
- **GET /customers/{customer_id}/vehicles/{vehicle_id}/history** ?customer_id,vehicle_id -> 200: Vehicle | 404: error envelope
- **GET /customers/{customer_id}/work_orders** ?customer_id -> 200: array of WorkOrder
- **GET /customers/{id}** ?id -> 200: Customer | 404: error envelope
- **PATCH /customers/{id}** ?id -> 200: Customer
- **PATCH /customers/{id}/archive** ?id -> 200: object{type,id,full_name,company_name,first_name,last_name,fleet_identifier,marketing_opt_in,tax_exempt,vehicles_count,emails_count,phones_count,status,trashed_at,vehicles_url,work_orders_url,created_at,updated_at,url,app_url,location,emails,phones,addresses,outstanding_balance_cents,total_revenue_cents,store_credit_cents,last_visit_at,statements_count,currency}
- **POST /customers/{id}/merges** ?id -> 200: object{type,id,full_name,company_name,first_name,last_name,fleet_identifier,marketing_opt_in,tax_exempt,vehicles_count,emails_count,phones_count,status,trashed_at,vehicles_url,work_orders_url,created_at,updated_at,url,app_url,location,emails,phones,addresses,outstanding_balance_cents,total_revenue_cents,store_credit_cents,last_visit_at,statements_count,currency} | 403: error envelope | 422: error envelope
- **PATCH /customers/{id}/restore** ?id -> 200: object{type,id,full_name,company_name,first_name,last_name,fleet_identifier,marketing_opt_in,tax_exempt,vehicles_count,emails_count,phones_count,status,trashed_at,vehicles_url,work_orders_url,created_at,updated_at,url,app_url,location,emails,phones,addresses,outstanding_balance_cents,total_revenue_cents,store_credit_cents,last_visit_at,statements_count,currency}
- **PATCH /customers/{id}/trash** ?id -> 200: object{type,id,full_name,company_name,first_name,last_name,fleet_identifier,marketing_opt_in,tax_exempt,vehicles_count,emails_count,phones_count,status,trashed_at,vehicles_url,work_orders_url,created_at,updated_at,url,app_url,location,emails,phones,addresses,outstanding_balance_cents,total_revenue_cents,store_credit_cents,last_visit_at,statements_count,currency} | 403: error envelope | 422: error envelope

## Drivers

- **GET /drivers** ?filters[has_open_work_order],filters[q] -> 200: array of Driver | 401: error envelope

## Expenses

- **GET /expenses** -> 200: array of object | 403: error envelope
- **POST /expenses** -> 201: object{id,payee,category,description,amount_cents,amount_currency,expense_date,payment_method,recurring,recurrence_rule,creator,location,created_at,updated_at,url,app_url} | 403: error envelope
- **GET /expenses/data_transfer** -> 200: object{resource,template_url,export_url,validate_url} | 403: error envelope
- **POST /expenses/export** -> 200: object{export_log_id,status} | 403: error envelope
- **GET /expenses/export/{id}/download** ?id -> no content | 404: error envelope
- **POST /expenses/imports/commit** -> 422: error envelope
- **GET /expenses/imports/template** -> no content
- **POST /expenses/imports/validate** -> 200: object{valid_count,error_count,duplicate_count,total,errors} | 403: error envelope | 422: error envelope
- **DELETE /expenses/{id}** ?id -> no content
- **GET /expenses/{id}** ?id -> 200: object{id,payee,category,description,amount_cents,amount_currency,expense_date,payment_method,recurring,recurrence_rule,creator,location,created_at,updated_at,url,app_url} | 403: error envelope
- **PATCH /expenses/{id}** ?id -> 200: object{id,payee,category,description,amount_cents,amount_currency,expense_date,payment_method,recurring,recurrence_rule,creator,location,created_at,updated_at,url,app_url}

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
- **GET /inspections/{id}** ?id -> 200: Inspection
- **PATCH /inspections/{id}** ?id -> 200: Inspection
- **PATCH /inspections/{id}/archive** ?id -> 200: object{id,type,name,description,status,trashed_at,is_default,created_at,updated_at,url,app_url,location,groups} | 422: error envelope
- **PATCH /inspections/{id}/remove_default** ?id -> 200: object{id,type,name,description,status,trashed_at,is_default,created_at,updated_at,url,app_url,location,groups}
- **PATCH /inspections/{id}/restore** ?id -> 200: object{id,type,name,description,status,trashed_at,is_default,created_at,updated_at,url,app_url,location,groups}
- **PATCH /inspections/{id}/set_default** ?id -> 200: object{id,type,name,description,status,trashed_at,is_default,created_at,updated_at,url,app_url,location,groups}
- **PATCH /inspections/{id}/toggle** ?id -> 200: object{id,type,name,description,status,trashed_at,is_default,created_at,updated_at,url,app_url,location,groups}
- **PATCH /inspections/{id}/trash** ?id -> 200: object{id,type,name,description,status,trashed_at,is_default,created_at,updated_at,url,app_url,location,groups}

## Inventory Levels

- **GET /inventory_levels** ?brand,page,per_page,q,stock_status,stocked -> 200: array of object
- **POST /inventory_levels** -> 201: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url} | 403: error envelope | 422: error envelope
- **GET /inventory_levels/barcode_lookup** ?barcode -> 200: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url} | 404: error envelope
- **DELETE /inventory_levels/{id}** ?id -> no content
- **GET /inventory_levels/{id}** ?id -> 200: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url} | no content | 404: error envelope
- **PATCH /inventory_levels/{id}** ?id -> 200: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url}
- **PATCH /inventory_levels/{id}/stock** ?id -> 200: object{id,part_number,description,brand,part_type,barcode,stocked,cost_cents,sell_cents,taxable,vendor,on_hand,reorder_point,bin_location,created_at,updated_at,url,app_url}

## Labor Matrices

- **GET /labor_matrices** -> 200: array of object
- **POST /labor_matrices** -> 201: object{id,name,matrix_type,active,tiers,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /labor_matrices/{id}** ?id -> no content
- **PATCH /labor_matrices/{id}** ?id -> 200: object{id,name,matrix_type,active,tiers,created_at,updated_at,url,app_url}

## Labor Rates

- **GET /labor_rates** -> 200: array of LaborRate
- **POST /labor_rates** -> 201: LaborRate | 403: error envelope
- **PATCH /labor_rates/{id}/archive** ?id -> 200: object{id,name,rate_cents,cost_per_hour_cents,is_default,status,trashed_at,created_at,updated_at,url,app_url}
- **PATCH /labor_rates/{id}/restore** ?id -> 200: object{id,name,rate_cents,cost_per_hour_cents,is_default,status,trashed_at,created_at,updated_at,url,app_url}
- **PATCH /labor_rates/{id}/trash** ?id -> 200: object{id,name,rate_cents,cost_per_hour_cents,is_default,status,trashed_at,created_at,updated_at,url,app_url}

## Labor Templates

- **GET /labor_templates** -> 200: array of object
- **DELETE /labor_templates/{id}** ?id -> no content
- **PATCH /labor_templates/{id}** ?id -> 200: object{id,description,default_hours,usage_count,last_used_at,created_at,updated_at,url,app_url}

## Lead Sources

- **GET /lead_sources** -> 200: array of object
- **POST /lead_sources** -> 201: object{id,name,category,system_default,active,position,created_at,updated_at,url,app_url} | 403: error envelope
- **POST /lead_sources/seed_defaults** -> 200: array of object
- **DELETE /lead_sources/{id}** ?id -> no content
- **PATCH /lead_sources/{id}** ?id -> 200: object{id,name,category,system_default,active,position,created_at,updated_at,url,app_url}

## Locations

- **GET /locations/{id}** ?id -> 200: object{id,name,slug,location_type,currency,time_zone,country,address,city,state,postal_code,contact_email,dock,url,app_url} | 401: error envelope | 404: error envelope
- **PATCH /locations/{id}** ?id -> 200: object{id,name,slug,location_type,currency,time_zone,country,address,city,state,postal_code,contact_email,dock,url,app_url} | 404: error envelope
- **GET /locations/{id}/business_profile** ?id -> 200: object{location,tax_breakdown}
- **PATCH /locations/{id}/business_profile** ?id -> 200: object{location,tax_breakdown} | 403: error envelope
- **GET /locations/{id}/operations** ?id -> 200: object{location,available_vendors}
- **PATCH /locations/{id}/operations** ?id -> 200: object{location,available_vendors} | 403: error envelope
- **GET /locations/{location_id}/close_requirements** ?location_id -> 200: object{id,name,default_starting_float_cents,driveon_station_number,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url} | 403: error envelope
- **PATCH /locations/{location_id}/close_requirements** ?location_id -> 200: object{id,name,default_starting_float_cents,driveon_station_number,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /locations/{location_id}/courtesy_cars** ?location_id -> 200: object{courtesy_cars} | 403: error envelope
- **GET /locations/{location_id}/documents** ?location_id -> 200: object{document_settings,estimate_terms_text,terms_text,payment_instructions,url,app_url} | 401: error envelope | 403: error envelope
- **PATCH /locations/{location_id}/documents** ?location_id -> 200: object{document_settings,estimate_terms_text,terms_text,payment_instructions,url,app_url}
- **GET /locations/{location_id}/lead_source_requirements** ?location_id -> 200: object{id,name,default_starting_float_cents,driveon_station_number,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url} | 403: error envelope
- **PATCH /locations/{location_id}/lead_source_requirements** ?location_id -> 200: object{id,name,default_starting_float_cents,driveon_station_number,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /locations/{location_id}/reminders** ?location_id -> 200: object{id,name,default_starting_float_cents,driveon_station_number,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url} | 403: error envelope
- **PATCH /locations/{location_id}/reminders** ?location_id -> 200: object{id,name,default_starting_float_cents,driveon_station_number,oil_change_reminders_enabled,tire_swap_reminders_enabled,brake_inspection_reminders_enabled,battery_check_reminders_enabled,close_requirements,lead_source_requirements,url,app_url}
- **GET /locations/{location_id}/schedule_config** ?location_id -> 200: object{schedule_config}
- **PATCH /locations/{location_id}/schedule_config** ?location_id -> 200: object{schedule_config} | 403: error envelope

## Me

- **GET /me/api_tokens** -> 200: object{api_tokens}
- **POST /me/api_tokens** -> 201: object{api_token}
- **DELETE /me/api_tokens/{id}** ?id -> no content | 404: error envelope
- **GET /me/notifications** -> 200: object{email_fallback_enabled,url,app_url}
- **PATCH /me/notifications** -> 200: object{email_fallback_enabled,url,app_url}
- **GET /me/preferences** -> 200: object{preferences,url,app_url}
- **PATCH /me/preferences** -> 200: object{preferences,url,app_url}
- **GET /me/profile** -> 200: object{user}
- **PATCH /me/profile** -> 200: object{user} | 422: error envelope

## Messages

- **GET /messages** ?conversation_id -> 200: array of object
- **GET /messages/{id}** ?id -> 200: object{id,conversation_id,direction,channel,status,body,recipient_phone,recipient_email,work_order_id,statement_id,appointment_id,failure_reason,attachment_count,sender,sent_at,delivered_at,read_at,failed_at,created_at,updated_at,conversation_url,url,app_url}
- **POST /messages/{message_id}/resends** ?message_id -> 201: object{id,conversation_id,direction,channel,status,body,recipient_phone,recipient_email,work_order_id,statement_id,appointment_id,failure_reason,attachment_count,sender,sent_at,delivered_at,read_at,failed_at,created_at,updated_at,conversation_url,url,app_url}

## Notifications

- **GET /notifications** -> 200: array of object | 401: error envelope
- **POST /notifications/bulk_mark_read** -> 200: object{ok,affected}
- **GET /notifications/{id}** ?id -> 200: object{id,trigger_type,title,message_body,read,read_at,created_at,updated_at,url,app_url} | 404: error envelope
- **PATCH /notifications/{id}** ?id -> 200: object{id,trigger_type,title,message_body,read,read_at,created_at,updated_at,url,app_url}

## Orders

- **GET /orders/purchase_orders** ?page,per_page,vendor_id -> 200: array of PurchaseOrder
- **POST /orders/purchase_orders** -> 201: PurchaseOrder | 403: error envelope | 404: error envelope
- **DELETE /orders/purchase_orders/{id}** ?id -> no content
- **GET /orders/purchase_orders/{id}** ?id -> 200: PurchaseOrder | no content
- **PATCH /orders/purchase_orders/{id}** ?id -> 200: PurchaseOrder
- **POST /orders/purchase_orders/{purchase_order_id}/cancellations** ?purchase_order_id -> 200: PurchaseOrder
- **GET /orders/return_orders** -> 200: array of ReturnOrder | 403: error envelope
- **DELETE /orders/return_orders/{id}** ?id -> no content
- **GET /orders/return_orders/{id}** ?id -> 200: ReturnOrder
- **PATCH /orders/return_orders/{id}** ?id -> 200: ReturnOrder
- **POST /orders/return_orders/{return_order_id}/refund_completions** ?return_order_id -> 200: ReturnOrder
- **GET /orders/sublet_orders** -> 200: array of SubletOrder | 403: error envelope
- **PATCH /orders/sublet_orders/mark_payment_complete** -> 200: SubletOrder
- **GET /orders/sublet_orders/{id}** ?id -> 200: SubletOrder | 403: error envelope

## Packages

- **GET /packages** -> 200: array of Package
- **POST /packages** -> 201: Package | 403: error envelope
- **PATCH /packages/{id}** ?id -> 403: error envelope
- **PATCH /packages/{id}/archive** ?id -> 200: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,status,trashed_at,url,app_url,location}
- **POST /packages/{id}/duplicate** ?id -> 201: Package | 403: error envelope
- **PATCH /packages/{id}/restore** ?id -> 200: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,status,trashed_at,url,app_url,location}
- **PATCH /packages/{id}/trash** ?id -> 200: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,status,trashed_at,url,app_url,location}

## Parts Matrices

- **GET /parts_matrices** -> 200: array of object
- **POST /parts_matrices** -> 201: object{id,name,is_default,active,default_multiplier,max_markup_cents,tiers,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /parts_matrices/{id}** ?id -> no content
- **PATCH /parts_matrices/{id}** ?id -> 200: object{id,name,is_default,active,default_multiplier,max_markup_cents,tiers,created_at,updated_at,url,app_url}

## Payments

- **GET /payments** ?method -> 200: array of object | 403: error envelope
- **GET /payments/pending** -> 200: array of object
- **GET /payments/{id}** ?id -> 200: object{id,amount_cents,currency,method,processor_status,is_refund,is_adjustment,voided,voided_at,processed_at,reference,created_at,updated_at,work_order_id,customer_id,url,app_url,work_order,customer,processed_by,location}
- **POST /payments/{id}/cancellation** ?id -> 200: object{id,amount_cents,currency,method,processor_status,is_refund,is_adjustment,voided,voided_at,processed_at,reference,created_at,updated_at,work_order_id,customer_id,url,app_url,work_order,customer,processed_by,location}
- **POST /payments/{id}/confirmation** ?id -> 200: object{id,amount_cents,currency,method,processor_status,is_refund,is_adjustment,voided,voided_at,processed_at,reference,created_at,updated_at,work_order_id,customer_id,url,app_url,work_order,customer,processed_by,location}
- **POST /payments/{id}/failure** ?id -> 200: object{id,amount_cents,currency,method,processor_status,is_refund,is_adjustment,voided,voided_at,processed_at,reference,created_at,updated_at,work_order_id,customer_id,url,app_url,work_order,customer,processed_by,location}

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
- **PATCH /service_categories/{id}** ?id -> 200: ServiceCategory
- **PATCH /service_categories/{id}/archive** ?id -> 200: object{id,name,description,service_type,icon,color,status,trashed_at,position,canonical,canonical_key,job_count,url,app_url}
- **PATCH /service_categories/{id}/restore** ?id -> 200: object{id,name,description,service_type,icon,color,status,trashed_at,position,canonical,canonical_key,job_count,url,app_url}
- **PATCH /service_categories/{id}/trash** ?id -> 200: object{id,name,description,service_type,icon,color,status,trashed_at,position,canonical,canonical_key,job_count,url,app_url}

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
- **GET /statements/{statement_id}/payments** ?statement_id -> 200: array of Statement

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

## Tire Events

- **GET /tire_events** -> 200: array of object | 401: error envelope
- **POST /tire_events** -> 201: object{id,event_type,event_type_label,occurred_at,notes,vehicle,created_at,updated_at,url,app_url} | 422: error envelope
- **GET /tire_events/{id}** ?id -> 200: object{id,event_type,event_type_label,occurred_at,notes,vehicle,created_at,updated_at,url,app_url}

## Tire Storage Slots

- **GET /tire_storage_slots** -> 200: array of object
- **POST /tire_storage_slots** -> 201: object{id,slot_label,season,season_label,tire_set_description,stored_at,released_at,currently_stored,storage_fee_cents,fee_type,notes,customer,vehicle,location,created_at,updated_at,url,app_url} | 403: error envelope
- **POST /tire_storage_slots/export** -> 200: object{export_log_id,status}
- **GET /tire_storage_slots/export/{id}/download** ?id -> no content | 404: error envelope
- **DELETE /tire_storage_slots/{id}** ?id -> no content
- **GET /tire_storage_slots/{id}** ?id -> 200: object{id,slot_label,season,season_label,tire_set_description,stored_at,released_at,currently_stored,storage_fee_cents,fee_type,notes,customer,vehicle,location,created_at,updated_at,url,app_url}
- **PATCH /tire_storage_slots/{id}** ?id -> 200: object{id,slot_label,season,season_label,tire_set_description,stored_at,released_at,currently_stored,storage_fee_cents,fee_type,notes,customer,vehicle,location,created_at,updated_at,url,app_url}
- **POST /tire_storage_slots/{tire_storage_slot_id}/check_outs** ?tire_storage_slot_id -> 200: object{id,slot_label,season,season_label,tire_set_description,stored_at,released_at,currently_stored,storage_fee_cents,fee_type,notes,customer,vehicle,location,created_at,updated_at,url,app_url}

## Tires

- **GET /tires** -> 200: array of object
- **POST /tires** -> 201: object{id,position,position_label,status,size_raw,size_width,size_aspect_ratio,size_rim_diameter,brand,model,load_index,speed_rating,dot_serial,dot_registered,dot_registered_at,purchase_date,tread_depth_new_mm,tread_depth_mm,source,notes,vehicle,created_at,updated_at,url,app_url} | 403: error envelope
- **POST /tires/export** -> 200: object{export_log_id,status}
- **GET /tires/export/{id}/download** ?id -> no content | 404: error envelope
- **POST /tires/imports/commit** -> 422: error envelope
- **GET /tires/imports/template** -> no content
- **POST /tires/imports/validate** -> 200: object{valid_count,error_count,duplicate_count,total,errors} | 422: error envelope
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
- **POST /users/{id}/disable** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities}
- **POST /users/{id}/enable** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities}
- **GET /users/{id}/qr_code** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities,qr_card}
- **POST /users/{id}/reset_pin** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities,new_pin}
- **POST /users/{id}/send_confirmation** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities}
- **POST /users/{id}/send_password_reset** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities}
- **POST /users/{id}/unlock** ?id -> 200: object{id,type,full_name,email,initials,role,disabled,disabled_at,locked,confirmed,hourly_cost_cents,hourly_cost_currency,certification_number,certification_label,mfa_enabled,mfa_required,qr_token_generated_at,created_at,updated_at,url,app_url,location,locations,capabilities}

## Vehicles

- **GET /vehicles** ?customer_id,page,per_page,status -> 200: array of Vehicle | 403: error envelope
- **POST /vehicles** -> 201: Vehicle | 403: error envelope | 422: error envelope
- **GET /vehicles/autocomplete** ?type -> 403: error envelope
- **GET /vehicles/check_duplicate** ?vin -> 200: object{matches} | 403: error envelope
- **GET /vehicles/customer_vehicles** ?customer_id -> 403: error envelope
- **GET /vehicles/lookup** ?query -> 200: array of Vehicle | 403: error envelope
- **GET /vehicles/prefill** ?make,model,vin,year -> 200: Vehicle | 403: error envelope | 404: error envelope
- **GET /vehicles/vin_decode** ?vin -> 200: object{make,model,vin} | 403: error envelope | 404: error envelope
- **GET /vehicles/{id}** ?id -> 200: Vehicle | 404: error envelope
- **PATCH /vehicles/{id}** ?id -> 200: Vehicle
- **PATCH /vehicles/{id}/archive** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,status,trashed_at,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count}
- **POST /vehicles/{id}/merges** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,status,trashed_at,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count} | 422: error envelope
- **PATCH /vehicles/{id}/restore** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,status,trashed_at,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count}
- **POST /vehicles/{id}/transfers** ?id -> 201: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,status,trashed_at,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count} | 403: error envelope | 422: error envelope
- **PATCH /vehicles/{id}/trash** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,status,trashed_at,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count} | 422: error envelope
- **GET /vehicles/{vehicle_id}/work_orders** ?vehicle_id -> 200: array of WorkOrder

## Vendors

- **GET /vendors** -> 200: array of Vendor | 401: error envelope | 403: error envelope
- **POST /vendors** -> 201: Vendor | 403: error envelope
- **GET /vendors/{id}** ?id -> 200: Vendor | 404: error envelope
- **PATCH /vendors/{id}** ?id -> 200: Vendor
- **PATCH /vendors/{id}/archive** ?id -> 200: object{id,name,vendor_type,payment_terms,status,trashed_at,phone,email,website,account_number,notes,quick_order,order_url_template,catalog_url_template,location,created_at,updated_at,url,app_url}
- **PATCH /vendors/{id}/restore** ?id -> 200: object{id,name,vendor_type,payment_terms,status,trashed_at,phone,email,website,account_number,notes,quick_order,order_url_template,catalog_url_template,location,created_at,updated_at,url,app_url}
- **PATCH /vendors/{id}/trash** ?id -> 200: object{id,name,vendor_type,payment_terms,status,trashed_at,phone,email,website,account_number,notes,quick_order,order_url_template,catalog_url_template,location,created_at,updated_at,url,app_url}
- **GET /vendors/{vendor_id}/purchase_orders** ?vendor_id -> 200: array of PurchaseOrder

## Voice Commands

- **DELETE /voice_commands/{id}** ?id -> 200: object{status} | 404: error envelope

## Work Order Tags

- **GET /work_order_tags** -> 200: array of object
- **POST /work_order_tags** -> 201: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at,created_at,updated_at,url,app_url}
- **PATCH /work_order_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at,created_at,updated_at,url,app_url}
- **PATCH /work_order_tags/{id}/archive** ?id -> 200: object{id,name,color,color_hex,color_class,bubble_classes,status,trashed_at,created_at,updated_at,url,app_url}

## Work Orders

- **GET /work_orders** ?per_page -> 200: array of WorkOrder | 403: error envelope
- **POST /work_orders** -> 201: WorkOrder
- **DELETE /work_orders/{id}** ?id -> no content
- **GET /work_orders/{id}** ?id -> 200: WorkOrder | 401: error envelope
- **PATCH /work_orders/{id}** ?id -> 200: WorkOrder | 403: error envelope | 422: error envelope
- **POST /work_orders/{id}/account_postings** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}/close** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url,reopen_url}
- **PATCH /work_orders/{id}/close_as_declined** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}/close_as_voided** ?id -> 200: WorkOrder
- **DELETE /work_orders/{id}/courtesy_car_assignment** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}/courtesy_car_assignment** ?id -> 200: WorkOrder
- **POST /work_orders/{id}/courtesy_car_assignment** ?id -> 200: WorkOrder | 403: error envelope
- **GET /work_orders/{id}/declined_services** ?id -> 200: array of object
- **POST /work_orders/{id}/estimate_deliveries** ?id -> 200: WorkOrder
- **POST /work_orders/{id}/invoice_summaries** ?id -> 200: WorkOrder
- **POST /work_orders/{id}/reminders** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}/reopen** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url} | 403: error envelope
- **POST /work_orders/{id}/service_declines** ?id -> 200: WorkOrder
- **GET /work_orders/{id}/service_history** ?id -> 200: array of object
- **POST /work_orders/{work_order_id}/authorization_decisions** ?work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/authorizations** ?work_order_id -> 200: WorkOrder | 403: error envelope | 422: error envelope
- **GET /work_orders/{work_order_id}/concerns** ?work_order_id -> 200: array of WorkOrder
- **GET /work_orders/{work_order_id}/estimate** ?work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/fee_exclusions** ?work_order_id -> 201: WorkOrder | 403: error envelope | 422: error envelope
- **DELETE /work_orders/{work_order_id}/fee_exclusions/{id}** ?id,work_order_id -> no content
- **GET /work_orders/{work_order_id}/inspection** ?work_order_id -> 200: WorkOrder
- **GET /work_orders/{work_order_id}/parts** ?work_order_id -> 200: WorkOrder
- **GET /work_orders/{work_order_id}/payments** ?work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/payments** ?work_order_id -> 201: WorkOrder | 403: error envelope | 422: error envelope
- **DELETE /work_orders/{work_order_id}/payments/reverse_ar** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url,reopen_url}
- **POST /work_orders/{work_order_id}/payments/send_to_ar** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url,activity_url,vehicle_history_url,appointments_url,authorization_logs_url,reopen_url}
- **POST /work_orders/{work_order_id}/refunds** ?work_order_id -> 201: WorkOrder
- **POST /work_orders/{work_order_id}/services** ?work_order_id -> 201: WorkOrder | no content | 422: error envelope
- **PATCH /work_orders/{work_order_id}/services/reorder** ?work_order_id -> no content
- **DELETE /work_orders/{work_order_id}/services/{id}** ?id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{id}** ?id,work_order_id -> 200: WorkOrder | no content | 422: error envelope
- **GET /work_orders/{work_order_id}/services/{id}/adjust_time** ?id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{id}/adjust_time** ?id,work_order_id -> 200: WorkOrder
- **DELETE /work_orders/{work_order_id}/services/{id}/authorization** ?id,work_order_id -> 200: WorkOrder
- **DELETE /work_orders/{work_order_id}/services/{id}/completion** ?id,work_order_id -> 200: WorkOrder | 403: error envelope
- **POST /work_orders/{work_order_id}/services/{id}/completion** ?id,work_order_id -> 200: WorkOrder | 422: error envelope
- **POST /work_orders/{work_order_id}/services/{id}/copies** ?id,work_order_id -> 201: WorkOrder
- **POST /work_orders/{work_order_id}/services/{id}/packages** ?id,work_order_id -> 200: WorkOrder | no content
- **PATCH /work_orders/{work_order_id}/services/{id}/pause** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/publish** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/revive** ?id,work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/services/{id}/time_entries** ?id,work_order_id -> 200: WorkOrder | 403: error envelope
- **PATCH /work_orders/{work_order_id}/services/{id}/toggle_labor_completion** ?id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{id}/update_category** ?id,work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/services/{service_id}/line_items** ?service_id,work_order_id -> 201: WorkOrder | 422: error envelope
- **DELETE /work_orders/{work_order_id}/services/{service_id}/line_items/{id}** ?id,service_id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}** ?id,service_id,work_order_id -> 200: WorkOrder | 422: error envelope
- **POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/copies** ?id,service_id,work_order_id -> 201: WorkOrder
- **POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/inventory_additions** ?id,service_id,work_order_id -> 200: WorkOrder
- **POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/price_refreshes** ?id,service_id,work_order_id -> 200: WorkOrder | 422: error envelope
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/pull** ?id,service_id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_pull** ?id,service_id,work_order_id -> 200: WorkOrder
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_return** ?id,service_id,work_order_id -> 200: WorkOrder
- **GET /work_orders/{work_order_id}/summary/activity** ?category,work_order_id -> 200: array of WorkOrder | 403: error envelope
- **GET /work_orders/{work_order_id}/summary/appointments** ?work_order_id -> 200: array of WorkOrder
- **GET /work_orders/{work_order_id}/summary/authorization_logs** ?work_order_id -> 200: array of WorkOrder
- **GET /work_orders/{work_order_id}/summary/vehicle_history** ?work_order_id -> 200: array of WorkOrder
- **POST /work_orders/{work_order_id}/voids** ?work_order_id -> 201: WorkOrder
- **GET /work_orders/{work_order_id}/wip** ?work_order_id -> 200: WorkOrder | no content


## Schemas

- **WorkOrder**: {type*:string, id*:integer, work_order_number*:integer, status*:string, intake_method*:string, scheduled_for*:string, authorized*:boolean, paid*:boolean, created_at*:string, updated_at*:string, closed_at*:string, location_id*:integer, service_advisor_id*:integer, assigned_technician_id*:integer, sub_status_type_id*:integer, payer_customer_id*:integer, vehicle_arrived_at*:string, work_order_services_count*:integer, inspection_reports_count*:integer, customer*:object{id,full_name,url}, vehicle*:object{id,make,model,year,vin,url}, totals*:object{subtotal_cents,tax_cents,total_cents,paid_cents,remaining_cents,currency}, url*:string, app_url*:string, location*:object{id,name,url}, odometer_in*:integer, odometer_out*:integer, odometer_unit*:string, authorized_at*:string, authorized_total_cents*:integer, customer_notified*:boolean, customer_notified_ready*:boolean, ready_for_pickup_at*:string, completed_at*:string, declined_at*:string, decline_reason*:string, discount_cents*:integer, fees_cents*:integer, parts_cents*:integer, labor_cents*:integer, tires_cents*:integer, subcontracts_cents*:integer, credit_balance_cents*:integer, saved_for_later*:boolean, closure_reason*:string, closure_reason_notes*:string, notes*:string, purchase_order_number*:string, return_method*:string, return_method_notes*:string, vehicle_keys_location*:string, vehicle_location*:string, customer_visit_count*:integer, customer_total_spend_cents*:integer, average_ticket_cents*:integer, activity_total*:integer, recent_activities*:array of object, services_url*:string, payments_url*:string, wip_url*:string, inspection_url*:string, parts_url*:string, concerns_url*:string, service_history_url*:string, declined_services_url*:string, activity_url*:string, vehicle_history_url*:string, appointments_url*:string, authorization_logs_url*:string, payer_customer:object{id,full_name,url}}
- **BroadcastCampaign**: {id*:integer, type*:string, name*:string, status*:string, sms_body*:string, filters*:object, recipient_count*:integer, sent_count*:integer, failed_count*:integer, progress_percentage*:integer, sent_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, creator*:object{id,name,url}, location*:object{id,name,url}}
- **Customer**: {type*:string, id*:integer, full_name*:string, company_name*:string, first_name*:string, last_name*:string, fleet_identifier*:string, marketing_opt_in*:boolean, tax_exempt*:boolean, vehicles_count*:integer, emails_count*:integer, phones_count*:integer, status*:string, trashed_at*:string, vehicles_url*:string, work_orders_url*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, emails*:array of object, phones*:array of object, addresses*:array of object, outstanding_balance_cents*:integer, total_revenue_cents*:integer, store_credit_cents*:integer, last_visit_at*:string, statements_count*:integer, currency*:string}
- **Driver**: {id*:integer, full_name*:string, phone*:string, email*:string, customer*:object{id,full_name,url}, work_orders_count*:integer, work_orders_url*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Statement**: {id*:integer, amount_cents*:integer, currency*:string, method*:string, processor_status*:string, is_refund*:boolean, is_adjustment*:boolean, voided*:boolean, voided_at*:string, processed_at*:string, reference*:string, created_at*:string, updated_at*:string, work_order_id*:integer, customer_id*:integer, url*:string, app_url*:string, work_order*:object{id,url}, customer*:object{id,full_name,url}, processed_by*:object{id,full_name,url}, location*:object{id,name,url}}
- **Vehicle**: {type*:string, id*:integer, make*:string, model*:string, year*:integer, submodel*:string, body_style*:string, engine*:string, vin*:string, license_plate*:string, license_plate_state*:string, license_plate_country*:string, drivetrain*:string, transmission*:string, color*:string, vehicle_type*:string, unit_number*:string, fleet_identifier*:string, production_date*:string, annual_safety_expires_at*:string, notes*:string, odometer*:object{reading,unit}, work_orders_count*:integer, status*:string, trashed_at*:string, work_orders_url*:string, customer*:object{id,full_name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, last_serviced_at*:string, lifetime_revenue_cents*:integer, open_work_orders_count*:integer, appointments_count*:integer}
- **InspectionReport**: {id*:integer, type*:string, name*:string, status*:string, work_order_id*:integer, quick_finding*:boolean, published*:boolean, completed*:boolean, items_count*:integer, checked_count*:integer, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, groups*:array of any}
- **Inspection**: {id*:integer, type*:string, name*:string, description*:string, status*:string, trashed_at*:string, is_default*:boolean, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, groups*:array of object}
- **LaborRate**: {id*:integer, name*:string, rate_cents*:integer, cost_per_hour_cents*:integer, is_default*:boolean, status*:string, trashed_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **PurchaseOrder**: {id*:integer, type*:string, po_number*:integer, status*:string, order_method*:string, payment_method*:string, fulfillment_method*:string, tracking_number*:any, vendor_invoice_number*:any, vendor_invoice_received_at*:string, notes*:string, freight_cost_cents*:integer, freight_cost_currency*:string, subtotal_cents*:string, total_cents*:string, core_charges_cents*:integer, line_items_count*:integer, ordered_at*:string, received_at*:string, payment_due_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, vendor*:object{id,name,url}, creator*:object{id,name,url}, location*:object{id,name,url}, line_items*:array of object}
- **ReturnOrder**: {id*:integer, type*:string, return_number*:integer, status*:string, credit_method*:any, reason_code*:any, rma_number*:any, is_warranty_claim*:boolean, restocking_fee_cents*:integer, shipping_fee_cents*:integer, notes*:string, line_items_count*:integer, refund_completed_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, vendor*:object{id,name,url}, work_order*:object{id,number,url}, creator*:object{id,name,url}, location*:object{id,name,url}, line_items*:array of object}
- **SubletOrder**: {id*:integer, type*:string, sublet_number*:integer, title*:string, payment_status*:string, payment_method*:any, total_cents*:integer, total_cost_cents*:integer, margin_cents*:integer, margin_percentage*:integer, sent_to_ap*:boolean, vendor_paid_at*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, vendor*:object{id,name,url}, work_order*:object{id,number,url}, work_order_service*:object{id,name}, location*:object{id,name,url}}
- **Package**: {id*:integer, name*:string, description*:string, service_type*:string, category_id*:integer, category_name*:string, estimated_hours*:string, customer_notes*:string, show_tech_with_cert*:boolean, triggers_tire_storage*:boolean, price_cents*:integer, status*:string, trashed_at*:string, url*:string, app_url*:string, location*:object{id,name,url}}
- **ServiceCategory**: {id*:integer, name*:string, description*:string, service_type*:string, icon*:string, color*:string, status*:string, trashed_at*:string, position*:integer, canonical*:boolean, canonical_key*:string, job_count*:integer, url*:string, app_url*:string}
- **Vendor**: {id*:integer, name*:string, vendor_type*:string, payment_terms*:string, status*:string, trashed_at*:string, phone*:string, email*:string, website*:string, account_number*:string, notes*:string, quick_order*:boolean, order_url_template*:string, catalog_url_template*:string, location*:object{id,name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Error**: {code*:string, message*:string, field_errors*:object}
- **UpdateAccountRequest**: {account*:object{name}}
- **CreateAccountStationLinkRegenerateRequest**: {}
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
- **CreateConversationRequest**: {conversation*:object{customer_id}}
- **CreateConversationsBulkMarkReadRequest**: {}
- **CreateConversationsCustomerLinkRequest**: {customer_id*:integer}
- **CreateConversationsIgnoreRequest**: {}
- **CreateConversationsMessageRequest**: {message*:object{body,channel}}
- **UpdateConversationRequest**: {status*:string}
- **UpdateCoreTaxRuleRequest**: {core_tax_rule*:object{tax_core_charge}}
- **CreateCounterSaleRequest**: {}
- **UpdateCounterSaleRequest**: {counter_sale*:object{notes}}
- **UpdateCounterSalesReopenRequest**: {}
- **CreateCourtesyCarRequest**: {vehicle*:object{year,make,model,vin}}
- **UpdateCourtesyCarRequest**: {vehicle*:object{color}}
- **UpdateCurrentLocationRequest**: {location*:object{name}}
- **CreateCustomerTagRequest**: {name*:string}
- **UpdateCustomerTagRequest**: {name*:string}
- **UpdateCustomerTagsArchiveRequest**: {}
- **UpdateCustomerTagsRestoreRequest**: {}
- **UpdateCustomerTagsTrashRequest**: {}
- **CreateCustomerRequest**: {customer*:object{first_name,last_name,company_name,fleet_identifier,billing_terms,credit_limit_cents,tax_exempt,tax_exempt_number,notes,marketing_opt_in,discount_percent,po_required,customer_tag_id,emails_attributes,phones_attributes,addresses_attributes}}
- **CreateCustomersExportRequest**: {}
- **CreateDriverRequest**: {driver*:object{full_name,phone}}
- **UpdateDriverRequest**: {driver*:object{full_name}}
- **UpdateCustomerRequest**: {customer*:object{phones_attributes,emails_attributes,status,first_name,last_name,company_name,fleet_identifier,billing_terms,credit_limit_cents,tax_exempt,notes,marketing_opt_in,discount_percent,po_required,customer_tag_id,addresses_attributes}}
- **ArchiveCustomerRequest**: {}
- **MergeCustomerRequest**: {source_customer_id*:integer}
- **RestoreCustomerRequest**: {}
- **TrashCustomerRequest**: {}
- **CreateExpenseRequest**: {expense*:object{payee,category,description,amount,expense_date,payment_method}}
- **CreateExpensesExportRequest**: {}
- **UpdateExpenseRequest**: {expense*:object{receipt_id,description}}
- **CreateInspectionReportRequest**: {}
- **CompleteInspectionReportRequest**: {}
- **MarkAllInspectionReportRequest**: {status*:string, group_name*:string}
- **PublishInspectionReportRequest**: {}
- **ReassignInspectionReportRequest**: {user_id*:integer}
- **ReopenInspectionReportRequest**: {}
- **ResetInspectionReportRequest**: {}
- **RetryInspectionReportRecordingRequest**: {}
- **UnpublishInspectionReportRequest**: {}
- **CreateInspectionRequest**: {inspection*:object{name}}
- **UpdateInspectionRequest**: {inspection*:object{name}}
- **ArchiveInspectionRequest**: {}
- **RemoveDefaultInspectionRequest**: {}
- **RestoreInspectionRequest**: {}
- **SetDefaultInspectionRequest**: {}
- **ToggleInspectionRequest**: {}
- **TrashInspectionRequest**: {}
- **CreateInventoryLevelRequest**: {part*:object{part_number,description,brand,part_type,stocked,initial_quantity,cost,sell}}
- **UpdateInventoryLevelRequest**: {part*:object{description}}
- **UpdateInventoryLevelsStockRequest**: {inventory_level*:object{on_hand,reason}}
- **CreateLaborMatriceRequest**: {labor_matrix*:object{name,matrix_type,active}}
- **UpdateLaborMatriceRequest**: {labor_matrix*:object{name}}
- **CreateLaborRateRequest**: {labor_rate*:object{name,rate,is_default}}
- **ArchiveLaborRateRequest**: {}
- **RestoreLaborRateRequest**: {}
- **TrashLaborRateRequest**: {}
- **UpdateLaborTemplateRequest**: {labor_template*:object{default_hours}}
- **CreateLeadSourceRequest**: {lead_source*:object{name,category}}
- **CreateLeadSourcesSeedDefaultRequest**: {}
- **UpdateLeadSourceRequest**: {lead_source*:object{active}}
- **UpdateLocationRequest**: {location*:object{name}}
- **UpdateLocationsBusinessProfileRequest**: {location*:object{trade_name}}
- **UpdateLocationsOperationsRequest**: {location*:object{time_zone}}
- **UpdateLocationsCloseRequirementsRequest**: {close_requirements*:object{odometer_in,key_location}}
- **UpdateLocationsDocumentsRequest**: {location*:object{estimate_terms_text}}
- **UpdateLocationsLeadSourceRequirementsRequest**: {lead_source_requirements*:object{customer_lead_source,ro_marketing_source}}
- **UpdateLocationsRemindersRequest**: {location*:object{oil_change_reminders_enabled,tire_swap_reminders_enabled}}
- **UpdateLocationsScheduleConfigRequest**: {schedule_config*:object{slot_duration_minutes}}
- **CreateMeApiTokenRequest**: {api_token*:object{name}}
- **UpdateMeNotificationsRequest**: {user*:object{email_fallback_enabled}}
- **UpdateMePreferencesRequest**: {user*:object{preferences}}
- **UpdateMeProfileRequest**: {user*:object{full_name,email}}
- **CreateMessagesResendRequest**: {}
- **CreateNotificationsBulkMarkReadRequest**: {notification_ids:array of integer}
- **UpdateNotificationRequest**: {read*:boolean}
- **CreateOrdersPurchaseOrderRequest**: {purchase_order*:object{vendor_id,payment_method,fulfillment_method,line_items}}
- **UpdateOrdersPurchaseOrderRequest**: {purchase_order*:object{tracking_number}}
- **CreateOrdersPurchaseOrdersCancellationRequest**: {}
- **UpdateOrdersReturnOrderRequest**: {return_order*:object{rma_number}}
- **CreateOrdersReturnOrdersRefundCompletionRequest**: {}
- **UpdateOrdersSubletOrdersMarkPaymentCompleteRequest**: {sublet_order_ids*:array of integer, payment_method*:string}
- **CreatePackageRequest**: {package*:object{name,description}}
- **ArchivePackageRequest**: {}
- **CreatePackagesDuplicateRequest**: {}
- **RestorePackageRequest**: {}
- **TrashPackageRequest**: {}
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
- **UpdateServiceCategoryRequest**: {service_category*:object{position,name}}
- **ArchiveServiceCategoryRequest**: {}
- **RestoreServiceCategoryRequest**: {}
- **TrashServiceCategoryRequest**: {}
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
- **CreateTireEventRequest**: {tire_event*:object{vehicle_id,event_type,occurred_at,notes}}
- **CreateTireStorageSlotRequest**: {tire_storage_slot*:object{vehicle_id,customer_id,slot_label,season,stored_at}}
- **CreateTireStorageSlotsExportRequest**: {}
- **UpdateTireStorageSlotRequest**: {tire_storage_slot*:object{tire_set_description}}
- **CreateTireStorageSlotsCheckOutRequest**: {}
- **CreateTireRequest**: {tire*:object{vehicle_id,position,status,size_raw,brand}}
- **CreateTiresExportRequest**: {}
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
- **CreateVehicleRequest**: {vehicle*:object{customer_id,vin,year,make,model,submodel,body_style,engine,transmission,drivetrain,color,license_plate,license_plate_state,odometer_reading,odometer_unit,unit_number,fleet_identifier,notes,production_date}}
- **UpdateVehicleRequest**: {vehicle*:object{make,model,year,vin,submodel,body_style,engine,transmission,drivetrain,color,license_plate,license_plate_state,odometer_reading,odometer_unit,notes}}
- **ArchiveVehicleRequest**: {}
- **MergeVehicleRequest**: {source_vehicle_id*:integer}
- **RestoreVehicleRequest**: {}
- **TransferVehicleRequest**: {customer_id*:integer, mode*:string}
- **TrashVehicleRequest**: {}
- **CreateVendorRequest**: {vendor*:object{name,vendor_type,payment_terms}}
- **UpdateVendorRequest**: {vendor*:object{name}}
- **ArchiveVendorRequest**: {}
- **RestoreVendorRequest**: {}
- **TrashVendorRequest**: {}
- **CreateWorkOrderTagRequest**: {name*:string, color*:string}
- **UpdateWorkOrderTagRequest**: {name*:string}
- **UpdateWorkOrderTagsArchiveRequest**: {}
- **CreateWorkOrderRequest**: {work_order*:object{customer_id,vehicle_id}}
- **UpdateWorkOrderRequest**: {work_order*:object{intake_method,work_order_tag_id,sub_status_type_id,vehicle_arrived_at,saved_for_later,payer_customer_id,waiting_for_customer}}
- **CreateWorkOrdersAccountPostingRequest**: {}
- **CloseWorkOrderRequest**: {}
- **UpdateWorkOrdersCloseAsDeclinedRequest**: {closure_reason*:string}
- **UpdateWorkOrdersCloseAsVoidedRequest**: {closure_reason*:string}
- **UpdateWorkOrdersCourtesyCarAssignmentRequest**: {action_type*:string}
- **CreateWorkOrdersCourtesyCarAssignmentRequest**: {vehicle_id*:integer}
- **CreateWorkOrdersEstimateDeliverieRequest**: {}
- **CreateWorkOrdersInvoiceSummarieRequest**: {}
- **CreateWorkOrdersReminderRequest**: {}
- **ReopenWorkOrderRequest**: {}
- **CreateWorkOrdersServiceDeclineRequest**: {decline_reason*:string}
- **CreateWorkOrdersAuthorizationDecisionRequest**: {service_decision_reasons*:object{1047559673}}
- **CreateWorkOrdersAuthorizationRequest**: {authorization_method*:string, service_ids*:array of integer, service_decisions*:object{1047559673}}
- **CreateWorkOrdersFeeExclusionRequest**: {work_order_fee_exclusion*:object{shop_fee_config_id}}
- **CreateWorkOrderPaymentRequest**: {payment*:object{method,amount_cents}}
- **SendWorkOrderPaymentToArRequest**: {}
- **CreateWorkOrdersRefundRequest**: {refund*:object{payment_id,amount,reason}}
- **CreateWorkOrdersServiceRequest**: {work_order_service*:object{name,service_type}, package_id:integer}
- **UpdateWorkOrdersServicesReorderRequest**: {service_ids*:array of integer}
- **UpdateWorkOrdersServiceRequest**: {work_order_service*:object{name,technician_id,pricing_mode,position,labor_tax_enabled}}
- **UpdateWorkOrdersServicesAdjustTimeRequest**: {hours*:integer, minutes*:integer}
- **CreateWorkOrdersServicesCompletionRequest**: {}
- **CreateWorkOrdersServicesCopieRequest**: {}
- **CreateWorkOrdersServicesPackageRequest**: {package_id*:integer}
- **UpdateWorkOrdersServicesPauseRequest**: {}
- **UpdateWorkOrdersServicesPublishRequest**: {}
- **UpdateWorkOrdersServicesReviveRequest**: {}
- **CreateWorkOrdersServicesTimeEntrieRequest**: {}
- **UpdateWorkOrdersServicesToggleLaborCompletionRequest**: {line_item_id*:integer}
- **UpdateWorkOrdersServicesUpdateCategoryRequest**: {category_id*:integer}
- **CreateWorkOrdersServicesLineItemRequest**: {work_order_line_item*:object{item_type,description,hours,labor_rate_id,unit_price,quantity,total}}
- **UpdateWorkOrdersServicesLineItemRequest**: {work_order_line_item*:object{part_status,description}}
- **CreateWorkOrdersServicesLineItemsCopieRequest**: {}
- **CreateWorkOrdersServicesLineItemsInventoryAdditionRequest**: {}
- **CreateWorkOrdersServicesLineItemsPriceRefresheRequest**: {}
- **PullWorkOrderServiceLineItemRequest**: {}
- **UndoPullWorkOrderServiceLineItemRequest**: {}
- **UndoReturnWorkOrderServiceLineItemRequest**: {}
- **CreateWorkOrdersVoidRequest**: {payment_id*:integer}

