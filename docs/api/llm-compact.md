# Wenmar Pro API — Compact Reference

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

Base URL: `https://app.wenmarpro.com`
Auth: `Authorization: Bearer <token>`
Responses: bare objects/arrays, no envelope. Errors: `{ "error": { code, message, details } }`.

## Account

- **GET /account** -> 200: object{id,name,slug,locations,url,app_url} | 403: error envelope

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
- **POST /customers/{id}/merge** ?id -> 200: object{type,id,full_name,company_name,first_name,last_name,fleet_identifier,marketing_opt_in,tax_exempt,vehicles_count,emails_count,phones_count,vehicles_url,work_orders_url,created_at,updated_at,url,app_url,location,emails,phones,addresses,outstanding_balance_cents,total_revenue_cents,store_credit_cents,last_visit_at,statements_count,currency} | 403: error envelope | 422: error envelope

## Locations

- **GET /locations/{id}** ?id -> 200: object{id,name,location_type,currency,dock,url,app_url} | 403: error envelope

## Service Categories

- **GET /service_categories** -> 200: array of ServiceCategory
- **POST /service_categories** -> 201: ServiceCategory | 403: error envelope
- **POST /service_categories/seed_defaults** -> 200: object{created,message}
- **DELETE /service_categories/{id}** ?id -> 200: ServiceCategory | 422: error envelope
- **PATCH /service_categories/{id}** ?id -> 200: ServiceCategory
- **PATCH /service_categories/{id}/deactivate** ?id -> 200: object{id,name,description,service_type,icon,color,active,position,canonical,canonical_key,job_count,url,app_url}
- **PATCH /service_categories/{id}/move_down** ?id -> 200: array of object
- **PATCH /service_categories/{id}/move_up** ?id -> 200: array of object
- **PATCH /service_categories/{id}/reactivate** ?id -> 200: object{id,name,description,service_type,icon,color,active,position,canonical,canonical_key,job_count,url,app_url}

## Settings

- **GET /settings/tags** -> 200: object{customer_tags,vehicle_tags} | no content
- **PATCH /settings/tags** -> 200: object{customer_tags,vehicle_tags} | no content

## Statements

- **GET /statements/{id}** ?id -> 200: Statement | 404: error envelope

## Team

- **GET /team/members** -> 200: array of object | 403: error envelope

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
- **POST /vehicles/{id}/merge** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count} | 422: error envelope
- **PATCH /vehicles/{id}/transfer** ?id -> 200: object{type,id,make,model,year,submodel,body_style,engine,vin,license_plate,license_plate_state,license_plate_country,drivetrain,transmission,color,vehicle_type,unit_number,fleet_identifier,production_date,annual_safety_expires_at,notes,odometer,work_orders_count,work_orders_url,customer,created_at,updated_at,url,app_url,location,last_serviced_at,lifetime_revenue_cents,open_work_orders_count,appointments_count} | 403: error envelope | 422: error envelope
- **GET /vehicles/{vehicle_id}/work_orders** ?vehicle_id -> 200: array of WorkOrder

## Vendors

- **GET /vendors** -> 200: array of Vendor
- **GET /vendors/{id}** ?id -> 200: Vendor | 404: error envelope

## Work Orders

- **GET /work_orders** -> 200: array of WorkOrder
- **POST /work_orders** -> 201: WorkOrder
- **DELETE /work_orders/{id}** ?id -> no content
- **GET /work_orders/{id}** ?id -> 200: WorkOrder
- **PATCH /work_orders/{id}** ?id -> 200: WorkOrder | 422: error envelope
- **GET /work_orders/{work_order_id}/estimate** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,services}
- **GET /work_orders/{work_order_id}/inspection** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,inspection_reports}
- **GET /work_orders/{work_order_id}/parts** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,services}
- **GET /work_orders/{work_order_id}/payments** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,payments}
- **POST /work_orders/{work_order_id}/payments** ?work_order_id -> 201: object{id,amount_cents,method,processor_status,is_refund,processed_at,reference,created_at,updated_at,work_order_id,customer_id,work_order,customer,processed_by} | 403: error envelope | 422: error envelope
- **GET /work_orders/{work_order_id}/summary** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,customer_visit_count,customer_total_spend_cents,average_ticket_cents,activity_total,recent_activities}
- **GET /work_orders/{work_order_id}/wip** ?work_order_id -> 200: object{type,id,work_order_number,status,intake_method,scheduled_for,authorized,paid,created_at,updated_at,closed_at,location_id,service_advisor_id,assigned_technician_id,work_order_services_count,inspection_reports_count,customer,vehicle,totals,url,app_url,location,services}


## Schemas

- **Customer**: {type*:string, id*:integer, full_name*:string, company_name*:string, first_name*:string, last_name*:string, fleet_identifier*:string, marketing_opt_in*:boolean, tax_exempt*:boolean, vehicles_count*:integer, emails_count*:integer, phones_count*:integer, vehicles_url*:string, work_orders_url*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, emails*:array of object, phones*:array of object, addresses*:array of object, outstanding_balance_cents*:integer, total_revenue_cents*:integer, store_credit_cents*:integer, last_visit_at*:any, statements_count*:integer, currency*:string}
- **Driver**: {id*:integer, full_name*:string, phone*:any, email*:any, customer*:object{id,full_name,url}, work_orders_count*:integer, work_orders_url*:string, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Statement**: {id*:integer, statement_number*:string, status*:string, statement_date*:string, start_date*:string, end_date*:string, due_date*:string, totals*:object{previous_balance_cents,new_charges_cents,payments_received_cents,credits_cents,balance_due_cents,currency}, sent_at*:any, viewed_at*:any, customer*:object{id,full_name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Vehicle**: {type*:string, id*:integer, make*:string, model*:string, year*:integer, submodel*:string, body_style*:string, engine*:string, vin*:string, license_plate*:string, license_plate_state*:string, license_plate_country*:string, drivetrain*:string, transmission*:string, color*:string, vehicle_type*:string, unit_number*:any, fleet_identifier*:any, production_date*:string, annual_safety_expires_at*:any, notes*:string, odometer*:object{reading,unit}, work_orders_count*:integer, work_orders_url*:string, customer*:object{id,full_name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string, location*:object{id,name,url}, last_serviced_at*:any, lifetime_revenue_cents*:integer, open_work_orders_count*:integer, appointments_count*:integer}
- **WorkOrder**: {type*:string, id*:integer, work_order_number*:integer, status*:string, intake_method*:string, scheduled_for*:any, authorized*:boolean, paid*:boolean, created_at*:string, updated_at*:string, closed_at*:any, location_id*:integer, service_advisor_id*:integer, assigned_technician_id*:any, work_order_services_count*:integer, inspection_reports_count*:integer, customer*:object{id,full_name,url}, vehicle*:object{id,make,model,year,vin,url}, totals*:object{subtotal_cents,tax_cents,total_cents,paid_cents,remaining_cents,currency}, url*:string, app_url*:string, location*:object{id,name,url}, odometer_in*:any, odometer_out*:any, odometer_unit*:string, authorized_at*:any, authorized_total_cents*:integer, customer_notified*:boolean, customer_notified_ready*:boolean, vehicle_arrived_at*:string, ready_for_pickup_at*:any, completed_at*:any, declined_at*:any, decline_reason*:any, discount_cents*:integer, fees_cents*:integer, parts_cents*:integer, labor_cents*:integer, tires_cents*:integer, subcontracts_cents*:integer, credit_balance_cents*:integer, saved_for_later*:boolean, closure_reason*:any, closure_reason_notes*:any, notes*:any, purchase_order_number*:any, return_method*:string, return_method_notes*:any, vehicle_keys_location*:string, vehicle_location*:string, summary_url*:string, services_url*:string, payments_url*:string, wip_url*:string, inspection_url*:string, parts_url*:string, concerns_url*:string}
- **ServiceCategory**: {id*:integer, name*:string, description*:any, service_type*:string, icon*:string, color*:string, active*:boolean, position*:integer, canonical*:boolean, canonical_key*:string, job_count*:integer, url*:string, app_url*:string}
- **Vendor**: {id*:integer, name*:string, vendor_type*:string, payment_terms*:string, active*:boolean, phone*:string, email*:string, website*:string, account_number*:string, notes*:string, quick_order*:boolean, order_url_template*:any, catalog_url_template*:any, location*:object{id,name,url}, created_at*:string, updated_at*:string, url*:string, app_url*:string}
- **Error**: {code*:string, message*:string, field_errors*:object}

