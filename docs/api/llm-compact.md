# Wenmar Pro API — Compact Reference

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

Base URL: `https://app.wenmarpro.com`
Auth: `Authorization: Bearer <token>`
Responses: bare objects/arrays, no envelope. Errors: `{ "error": { code, message, details } }`.

## Account

- **GET /account** -> 200: object{id,name,slug,locations,url,app_url} | 403: error envelope

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

## Counter Sales

- **GET /counter_sales** -> 200: array of object
- **POST /counter_sales** -> 201: object{id,counter_sale_number,status,walk_in_name,notes,subtotal_cents,tax_total_cents,total_cents,paid_cents,remaining_cents,paid,currency,line_items_count,processed_by,location,created_at,updated_at,url,app_url} | 403: error envelope
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

- **GET /customers** ?has_balance,has_vehicle,last_visit_months,page,per_page,query,tag_ids,type -> 200: array of Customer | 401: error envelope
- **POST /customers** -> 201: Customer | 403: error envelope | 422: error envelope
- **GET /customers/check_duplicate** ?email,first_name,last_name,phone -> 200: object{matches}
- **GET /customers/lookup** ?id,query -> 200: array of Customer
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

- **GET /expenses** -> 200: array of object
- **POST /expenses** -> 201: object{id,payee,category,description,amount_cents,amount_currency,expense_date,payment_method,recurring,recurrence_rule,location,created_at,updated_at,url,app_url,creator} | 403: error envelope
- **DELETE /expenses/{id}** ?id -> no content
- **PATCH /expenses/{id}** ?id -> 200: object{id,payee,category,description,amount_cents,amount_currency,expense_date,payment_method,recurring,recurrence_rule,location,created_at,updated_at,url,app_url,creator}

## Fleets

- **GET /fleets** -> 200: array of object | 403: error envelope

## Locations

- **GET /locations/{id}** ?id -> 200: object{id,name,location_type,currency,dock,url,app_url} | 403: error envelope

## Notifications

- **GET /notifications** -> 200: array of object | 401: error envelope
- **POST /notifications/bulk_mark_read** -> 200: object{ok,affected}
- **GET /notifications/{id}** ?id -> 200: object{id,trigger_type,title,message_body,read,read_at,created_at,updated_at,url,app_url} | 404: error envelope
- **PATCH /notifications/{id}** ?id -> 200: object{id,trigger_type,title,message_body,read,read_at,created_at,updated_at,url,app_url}

## Packages

- **GET /packages** -> 200: array of object
- **POST /packages** -> 201: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,active,url,app_url,location} | 403: error envelope
- **PATCH /packages/{id}** ?id -> 200: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,active,url,app_url,location}
- **POST /packages/{id}/duplicate** ?id -> 201: object{id,name,description,service_type,category_id,category_name,estimated_hours,customer_notes,show_tech_with_cert,triggers_tire_storage,price_cents,active,url,app_url,location}

## Service Categories

- **GET /service_categories** -> 200: array of ServiceCategory
- **POST /service_categories** -> 201: ServiceCategory | 403: error envelope
- **POST /service_categories/seed_defaults** -> 200: object{created,message}
- **DELETE /service_categories/{id}** ?id -> 200: ServiceCategory | 422: error envelope
- **PATCH /service_categories/{id}** ?id -> 200: ServiceCategory

## Settings

- **GET /settings/tags** -> 200: object{customer_tags,vehicle_tags}
- **PATCH /settings/tags** -> 200: object{customer_tags,vehicle_tags} | no content

## Statements

- **GET /statements/{id}** ?id -> 200: Statement | 404: error envelope

## Sublet Packages

- **GET /sublet_packages** -> 200: array of object
- **POST /sublet_packages** -> 201: object{id,name,description,active,default_fulfillment_status,default_payment_status,default_payment_method,sublet_package_lines_count,location,created_at,updated_at,url,app_url} | 403: error envelope
- **DELETE /sublet_packages/{id}** ?id -> no content
- **PATCH /sublet_packages/{id}** ?id -> 200: object{id,name,description,active,default_fulfillment_status,default_payment_status,default_payment_method,sublet_package_lines_count,location,created_at,updated_at,url,app_url}

## Team

- **GET /team/members** -> 200: array of object | 403: error envelope

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

## Vehicle Tags

- **GET /vehicle_tags** -> 200: array of object
- **POST /vehicle_tags** -> 201: object{id,name,color,color_hex,color_class}
- **DELETE /vehicle_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class}
- **PATCH /vehicle_tags/{id}** ?id -> 200: object{id,name,color,color_hex,color_class}

## Vehicles

- **GET /vehicles** ?customer_id,page -> 200: array of Vehicle | 403: error envelope
- **POST /vehicles** -> 201: Vehicle | 403: error envelope | 422: error envelope
- **GET /vehicles/check_duplicate** ?vin -> 200: object{matches}
- **GET /vehicles/lookup** ?query -> 200: array of Vehicle
- **GET /vehicles/prefill** ?make,model,vin,year -> 200: Vehicle | 404: error envelope
- **GET /vehicles/vin_decode** ?vin -> 200: object{make,model,vin} | 404: error envelope
- **DELETE /vehicles/{id}** ?id -> 202: Vehicle | 422: error envelope
- **GET /vehicles/{id}** ?id -> 200: Vehicle | 404: error envelope
- **PATCH /vehicles/{id}** ?id -> 200: Vehicle
- **POST /vehicles/{id}/merges** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count} | 422: error envelope
- **POST /vehicles/{id}/transfers** ?id -> 201: object{type,id,vehicle_id,from_customer_id,to_customer_id,status,mode} | 403: error envelope | 422: error envelope
- **GET /vehicles/{vehicle_id}/work_orders** ?vehicle_id -> 200: array of WorkOrder

## Vendors

- **GET /vendors** -> 200: array of Vendor
- **POST /vendors** -> 201: Vendor | 403: error envelope
- **DELETE /vendors/{id}** ?id -> no content
- **GET /vendors/{id}** ?id -> 200: Vendor | 404: error envelope
- **PATCH /vendors/{id}** ?id -> 200: Vendor
- **GET /vendors/{vendor_id}/purchase_orders** ?vendor_id -> 200: array of object

## Work Orders

- **GET /work_orders** -> 200: array of WorkOrder | 403: error envelope
- **POST /work_orders** -> 201: WorkOrder
- **DELETE /work_orders/{id}** ?id -> no content
- **GET /work_orders/{id}** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}** ?id -> 200: WorkOrder | 422: error envelope
- **PATCH /work_orders/{id}/close** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url} | 422: error envelope
- **PATCH /work_orders/{id}/close_as_paid** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **PATCH /work_orders/{id}/close_zero** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **POST /work_orders/{id}/courtesy_car_assignment** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **PATCH /work_orders/{id}/decline_all** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **GET /work_orders/{id}/declined_services** ?id -> 200: array of object
- **DELETE /work_orders/{id}/hard_delete** ?id -> 202: object{status,resource} | 422: error envelope
- **PATCH /work_orders/{id}/post_to_account** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **PATCH /work_orders/{id}/reopen** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url} | 403: error envelope
- **PATCH /work_orders/{id}/return_to_board** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **PATCH /work_orders/{id}/save_for_later** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **PATCH /work_orders/{id}/send_estimate** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **PATCH /work_orders/{id}/send_invoice_summary** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **PATCH /work_orders/{id}/send_reminder** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **GET /work_orders/{id}/service_history** ?id -> 200: array of object
- **PATCH /work_orders/{id}/toggle_waiting_for_customer** ?id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **POST /work_orders/{work_order_id}/authorization** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url} | 403: error envelope | 422: error envelope
- **POST /work_orders/{work_order_id}/authorization/update_decisions** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **GET /work_orders/{work_order_id}/concerns** ?work_order_id -> 200: array of object
- **GET /work_orders/{work_order_id}/estimate** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,services}
- **GET /work_orders/{work_order_id}/inspection** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,inspection_reports}
- **POST /work_orders/{work_order_id}/labels** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **DELETE /work_orders/{work_order_id}/labels/{id}** ?id,work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **GET /work_orders/{work_order_id}/parts** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,services}
- **GET /work_orders/{work_order_id}/payments** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,payments}
- **POST /work_orders/{work_order_id}/payments** ?work_order_id -> 201: object{id,amount_cents,method,processor_status,is_refund,processed_at,reference,created_at,updated_at,work_order_id,customer_id,work_order,customer,processed_by,location} | 403: error envelope | 422: error envelope
- **DELETE /work_orders/{work_order_id}/payments/reverse_ar** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **POST /work_orders/{work_order_id}/payments/send_to_ar** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,odometer_in,odometer_out,odometer_unit,authorized_at,authorized_total_cents,customer_notified,customer_notified_ready,ready_for_pickup_at,completed_at,declined_at,decline_reason,discount_cents,fees_cents,parts_cents,labor_cents,tires_cents,subcontracts_cents,credit_balance_cents,saved_for_later,closure_reason,closure_reason_notes,notes,purchase_order_number,return_method,return_method_notes,vehicle_keys_location,vehicle_location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities,services_url,payments_url,wip_url,inspection_url,parts_url,concerns_url,service_history_url,declined_services_url}
- **POST /work_orders/{work_order_id}/services** ?work_order_id -> 201: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items} | no content | 422: error envelope
- **DELETE /work_orders/{work_order_id}/services/{id}** ?id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{id}** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/acknowledge_parts** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **POST /work_orders/{work_order_id}/services/{id}/add_line_item** ?id,work_order_id -> 201: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,fee_type,fee_on_labor,fee_on_parts,fee_on_sublets} | 422: error envelope
- **POST /work_orders/{work_order_id}/services/{id}/add_package** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items} | no content
- **GET /work_orders/{work_order_id}/services/{id}/adjust_time** ?id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{id}/adjust_time** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/apply_discount** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/complete_service** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items} | 422: error envelope
- **POST /work_orders/{work_order_id}/services/{id}/duplicate** ?id,work_order_id -> 201: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/pause** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/publish** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/reset_completion** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items} | 403: error envelope
- **PATCH /work_orders/{work_order_id}/services/{id}/revive** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/start** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items} | 403: error envelope | 415: object
- **PATCH /work_orders/{work_order_id}/services/{id}/toggle_labor_completion** ?id,work_order_id -> 200: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,hours,rate_cents,technician_id,labor_rate_id}
- **PATCH /work_orders/{work_order_id}/services/{id}/toggle_labor_tax** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **POST /work_orders/{work_order_id}/services/{id}/unauthorize** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/update_category** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **PATCH /work_orders/{work_order_id}/services/{id}/update_pricing_mode** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items} | 422: error envelope
- **PATCH /work_orders/{work_order_id}/services/{id}/update_service_technician** ?id,work_order_id -> 200: object{id,name,service_type,authorization_status,pricing_mode,technician_id,technician,category_id,ordinal,discount_cents,labor_cents,parts_cents,fees_cents,sublet_cents,tires_cents,total_cents,tax_total_cents,estimated_hours,customer_notes,started_at,completed_at,authorized_at,created_at,updated_at,line_items}
- **POST /work_orders/{work_order_id}/services/{service_id}/line_items** ?service_id,work_order_id -> 201: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,hours,rate_cents,technician_id,labor_rate_id,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size}
- **DELETE /work_orders/{work_order_id}/services/{service_id}/line_items/{id}** ?id,service_id,work_order_id -> no content
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}** ?id,service_id,work_order_id -> 200: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size}
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/add_to_inventory** ?id,service_id,work_order_id -> 200: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size}
- **POST /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/duplicate** ?id,service_id,work_order_id -> 201: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size}
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/pull** ?id,service_id,work_order_id -> 200: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size}
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/refresh_price** ?id,service_id,work_order_id -> 200: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size} | 422: error envelope
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_pull** ?id,service_id,work_order_id -> 200: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size}
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_return** ?id,service_id,work_order_id -> 200: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size}
- **PATCH /work_orders/{work_order_id}/services/{service_id}/line_items/{id}/update_part_status** ?id,service_id,work_order_id -> 200: object{id,item_type,description,quantity,unit_price_cents,total_cents,tax_total_cents,pricing_mode,is_taxable,is_warranty,completed,notes,created_at,updated_at,part_type,part_status,vendor_id,vendor_part_number,unit_cost_cents,brand,dot_code,tire_size} | 422: error envelope
- **GET /work_orders/{work_order_id}/wip** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,sub_status_type_id,payer_customer_id,vehicle_arrived_at,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,services}


## Schemas

- **Customer**: {type*:string, id*:integer, full_name*:string, company_name*:string, first_name*:string, last_name*:string, fleet_identifier*:string, marketing_opt_in*:boolean, tax_exempt*:boolean, vehicles_count*:integer, emails_count*:integer, phones_count*:integer, vehicles_url*:string, work_orders_url*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, emails*:array of object, phones*:array of object, addresses*:array of object, outstanding_balance_cents*:integer, total_revenue_cents*:integer, store_credit_cents*:integer, last_visit_at*:any, statements_count*:integer, currency*:string}
- **Driver**: {id*:integer, full_name*:string, phone*:any, email*:any, customer*:object{id,full_name,url}, work_orders_count*:integer, work_orders_url*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Statement**: {id*:integer, statement_number*:string, status*:string, statement_date*:string, start_date*:string, end_date*:string, due_date*:string, totals*:object{previous_balance_cents,new_charges_cents,payments_received_cents,credits_cents,balance_due_cents,currency}, sent_at*:any, viewed_at*:any, customer*:object{id,full_name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Vehicle**: {type*:string, id*:integer, make*:string, model*:string, year*:integer, submodel*:string, body_style*:string, engine*:string, vin*:string, license_plate*:string, license_plate_state*:string, license_plate_country*:string, drivetrain*:string, transmission*:string, color*:string, vehicle_type*:string, unit_number*:any, fleet_identifier*:any, production_date*:string, annual_safety_expires_at*:any, notes*:string, odometer*:object{reading,unit}, work_orders_count*:integer, work_orders_url*:string, customer*:object{id,full_name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, last_serviced_at*:any, lifetime_revenue_cents*:integer, open_work_orders_count*:integer, appointments_count*:integer}
- **WorkOrder**: {type*:string, id*:integer, work_order_number*:integer, status*:string, intake_method*:string, scheduled_for*:any, authorized*:boolean, paid*:boolean, created_at*:string, updated_at*:string, closed_at*:any, location_id*:integer, service_advisor_id*:any, assigned_technician_id*:any, sub_status_type_id*:integer, payer_customer_id*:integer, vehicle_arrived_at*:string, work_order_services_count*:integer, inspection_reports_count*:integer, customer*:object{id,full_name,url}, payer_customer:object{id,full_name,url}, vehicle*:object{id,make,model,year,vin,url}, totals*:object{subtotal_cents,tax_cents,total_cents,paid_cents,remaining_cents,currency}, url*:string, app_url*:string, location*:object{id,name,url}, odometer_in*:any, odometer_out*:any, odometer_unit*:string, authorized_at*:any, authorized_total_cents*:integer, customer_notified*:boolean, customer_notified_ready*:boolean, ready_for_pickup_at*:any, completed_at*:any, declined_at*:any, decline_reason*:any, discount_cents*:integer, fees_cents*:integer, parts_cents*:integer, labor_cents*:integer, tires_cents*:integer, subcontracts_cents*:integer, credit_balance_cents*:integer, saved_for_later*:boolean, closure_reason*:any, closure_reason_notes*:any, notes*:any, purchase_order_number*:any, return_method*:string, return_method_notes*:any, vehicle_keys_location*:string, vehicle_location*:string, customer_visit_count*:integer, customer_total_spend_cents*:integer, average_ticket_cents*:integer, activity_total*:integer, recent_activities*:array of any, services_url*:string, payments_url*:string, wip_url*:string, inspection_url*:string, parts_url*:string, concerns_url*:string, service_history_url*:string, declined_services_url*:string}
- **ServiceCategory**: {id*:integer, name*:string, description*:any, service_type*:string, icon*:string, color*:string, active*:boolean, position*:integer, canonical*:boolean, canonical_key*:string, job_count*:integer, url*:string, app_url*:string}
- **Vendor**: {id*:integer, name*:string, vendor_type*:string, payment_terms*:string, active*:boolean, phone*:string, email*:string, website*:string, account_number*:string, notes*:string, quick_order*:boolean, order_url_template*:any, catalog_url_template*:any, location*:object{id,name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Error**: {code*:string, message*:string, field_errors*:object}
- **CreateConversationRequest**: {}
- **PostConversationsCustomerLinkRequest**: {customer_id*:integer}
- **PostConversationsIgnoreRequest**: {}
- **PostConversationsMessageRequest**: {message*:object{body,channel}}
- **PostConversationsMessageResendsRequest**: {}
- **UpdateConversationRequest**: {status*:string}
- **CreateCounterSaleRequest**: {}
- **UpdateCounterSaleRequest**: {counter_sale*:object{notes}}
- **CreateCustomerTagRequest**: {name*:string}
- **UpdateCustomerTagRequest**: {name*:string}
- **CreateCustomerRequest**: {customer*:object{first_name,last_name,company_name,fleet_identifier,billing_terms,credit_limit_cents,tax_exempt,tax_exempt_number,notes,marketing_opt_in,discount_percent,po_required,tag_ids,emails_attributes,phones_attributes,addresses_attributes}}
- **CreateDriverRequest**: {driver*:object{full_name,phone}}
- **UpdateDriverRequest**: {driver*:object{full_name}}
- **UpdateCustomerRequest**: {customer*:object{first_name,last_name,company_name,fleet_identifier,billing_terms,credit_limit_cents,tax_exempt,notes,marketing_opt_in,discount_percent,po_required,emails_attributes,phones_attributes,addresses_attributes}}
- **MergeCustomerRequest**: {source_customer_id*:integer}
- **CreateExpenseRequest**: {expense*:object{payee,category,description,amount,expense_date,payment_method}}
- **UpdateExpenseRequest**: {expense*:object{description}}
- **CreateNotificationRequest**: {notification_ids:array of integer}
- **UpdateNotificationRequest**: {read*:boolean}
- **CreatePackageRequest**: {package*:object{name,description}}
- **UpdatePackageRequest**: {package*:object{active}}
- **PostPackagesDuplicateRequest**: {}
- **CreateServiceCategoryRequest**: {service_category*:object{name,service_type,icon}}
- **SeedDefaultsServiceCategoriesRequest**: {}
- **UpdateServiceCategoryRequest**: {service_category*:object{name,active,position}}
- **UpdateTagsRequest**: {customer_tags:array of object, vehicle_tags:array of object}
- **CreateSubletPackageRequest**: {sublet_package*:object{name,description,active}}
- **UpdateSubletPackageRequest**: {sublet_package*:object{name}}
- **CreateTireStorageSlotRequest**: {tire_storage_slot*:object{vehicle_id,customer_id,slot_label,season,stored_at}}
- **PostTireStorageSlotsCheckOutRequest**: {}
- **CreateTireRequest**: {tire*:object{vehicle_id,position,status,size_raw,brand}}
- **UpdateTireRequest**: {tire*:object{brand}}
- **CreateVehicleTagRequest**: {name*:string}
- **UpdateVehicleTagRequest**: {name*:string}
- **CreateVehicleRequest**: {vehicle*:object{customer_id,vin,year,make,model,submodel,body_style,engine,transmission,drivetrain,color,license_plate,license_plate_state,odometer_reading,odometer_unit,unit_number,fleet_identifier,notes,production_date,vehicle_tag_ids}}
- **UpdateVehicleRequest**: {vehicle*:object{make,model,year,vin,submodel,body_style,engine,transmission,drivetrain,color,license_plate,license_plate_state,odometer_reading,odometer_unit,notes}}
- **MergeVehicleRequest**: {source_vehicle_id*:integer}
- **TransferVehicleRequest**: {customer_id*:integer, mode*:string}
- **CreateVendorRequest**: {vendor*:object{name,vendor_type,payment_terms}}
- **UpdateVendorRequest**: {vendor*:object{name}}
- **CreateWorkOrderRequest**: {work_order*:object{customer_id,vehicle_id}}
- **UpdateWorkOrderRequest**: {work_order*:object{payer_customer_id,intake_method,sub_status_type_id,vehicle_arrived_at}}
- **CloseWorkOrderRequest**: {closure_type*:string, closure_reason*:string}
- **CloseWorkOrderAsPaidRequest**: {}
- **CloseWorkOrderZeroRequest**: {}
- **PostWorkOrdersCourtesyCarAssignmentRequest**: {vehicle_id*:integer}
- **DeclineAllWorkOrderServicesRequest**: {decline_reason*:string}
- **PatchWorkOrdersPostToAccountRequest**: {}
- **ReopenWorkOrderRequest**: {}
- **ReturnWorkOrderToBoardRequest**: {}
- **SaveWorkOrderForLaterRequest**: {}
- **PatchWorkOrdersSendEstimateRequest**: {}
- **PatchWorkOrdersSendInvoiceSummaryRequest**: {}
- **PatchWorkOrdersSendReminderRequest**: {}
- **PatchWorkOrdersToggleWaitingForCustomerRequest**: {}
- **CreateWorkOrderAuthorizationRequest**: {authorization_method*:string, service_ids*:array of integer, service_decisions*:object{1047559673}}
- **UpdateWorkOrderAuthorizationDecisionsRequest**: {service_decision_reasons*:object{1047559673}}
- **PostWorkOrdersLabelRequest**: {label_id*:integer}
- **CreateWorkOrderPaymentRequest**: {payment*:object{method,amount_cents}}
- **SendWorkOrderPaymentToArRequest**: {}
- **PostWorkOrdersServiceRequest**: {work_order_service*:object{name,service_type}, package_id:integer}
- **PatchWorkOrdersServiceRequest**: {work_order_service*:object{name,pricing_mode}}
- **PatchWorkOrdersServiceAcknowledgePartsRequest**: {}
- **PostWorkOrdersServiceAddLineItemRequest**: {item_type*:string, name*:string, amount_cents*:integer}
- **PostWorkOrdersServiceAddPackageRequest**: {package_id*:integer}
- **PatchWorkOrdersServiceAdjustTimeRequest**: {hours*:integer, minutes*:integer}
- **PatchWorkOrdersServiceApplyDiscountRequest**: {discount*:object{type,scope,value_cents}}
- **PatchWorkOrdersServiceCompleteServiceRequest**: {}
- **PostWorkOrdersServiceDuplicateRequest**: {}
- **PatchWorkOrdersServicePauseRequest**: {}
- **PatchWorkOrdersServicePublishRequest**: {}
- **PatchWorkOrdersServiceResetCompletionRequest**: {}
- **PatchWorkOrdersServiceReviveRequest**: {}
- **PatchWorkOrdersServiceStartRequest**: {}
- **PatchWorkOrdersServiceToggleLaborCompletionRequest**: {line_item_id*:integer}
- **PatchWorkOrdersServiceToggleLaborTaxRequest**: {}
- **PostWorkOrdersServiceUnauthorizeRequest**: {}
- **PatchWorkOrdersServiceUpdateCategoryRequest**: {category_id*:integer}
- **PatchWorkOrdersServiceUpdatePricingModeRequest**: {work_order_service*:object{pricing_mode}}
- **PatchWorkOrdersServiceUpdateServiceTechnicianRequest**: {work_order_service*:object{technician_id}}
- **PostWorkOrdersServiceLineItemsRequest**: {work_order_line_item*:object{item_type,description,hours,labor_rate_id,unit_price,quantity}}
- **PatchWorkOrdersServiceLineItemsRequest**: {work_order_line_item*:object{description}}
- **AddWorkOrderServiceLineItemToInventoryRequest**: {}
- **DuplicateWorkOrderServiceLineItemRequest**: {}
- **PullWorkOrderServiceLineItemRequest**: {}
- **RefreshWorkOrderServiceLineItemPriceRequest**: {}
- **UndoPullWorkOrderServiceLineItemRequest**: {}
- **UndoReturnWorkOrderServiceLineItemRequest**: {}
- **UpdateWorkOrderServiceLineItemPartStatusRequest**: {part_status*:string}

