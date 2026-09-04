# Wenmar Pro API Reference

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

Base URL: `https://app.wenmarpro.com`

For detailed per-resource docs, see [sections/](sections/).

## Endpoints

| Method | Path | Operation | Description |
|---|---|---|---|
| DELETE | `/account` | `delete_account` | Delete a account by ID. |
| GET | `/account` | `list_account` | List all account, paginated via the Link header. |
| PATCH | `/account` | `update_account` | Update a account by ID. |
| GET | `/account/billing` | `list_account_billing` | List all account billing, paginated via the Link header. |
| GET | `/account/capabilities` | `list_account_capabilities` | List all account capabilities, paginated via the Link header. |
| GET | `/account/driveon` | `list_account_driveon` | List all account driveon, paginated via the Link header. |
| GET | `/account/payments` | `list_account_payments` | List all account payments, paginated via the Link header. |
| GET | `/account/phone_numbers` | `list_account_phone_numbers` | List all account phone numbers, paginated via the Link header. |
| GET | `/account/quickbooks` | `list_account_quickbooks` | List all account quickbooks, paginated via the Link header. |
| GET | `/account/station_link` | `list_account_station_link` | List all account station link, paginated via the Link header. |
| POST | `/account/station_link/regenerate` | `create_account_station_link_regenerate` | Create a account station link regenerate. |
| PATCH | `/ai_suggestions/{id}` | `update_ai_suggestion` | Update a ai suggestion by ID. |
| GET | `/appointments` | `list_appointments` | List all appointments, paginated via the Link header. |
| POST | `/appointments` | `create_appointment` | Create a appointment. |
| GET | `/appointments/available_slots` | `list_appointments_available_slots` | List all appointments available slots, paginated via the Link header. |
| DELETE | `/appointments/{id}` | `delete_appointment` | Delete a appointment by ID. |
| GET | `/appointments/{id}` | `show_appointment` | Show a appointment by ID. |
| PATCH | `/appointments/{id}` | `update_appointment` | Update a appointment by ID. |
| POST | `/appointments/{id}/approvals` | `create_appointments_approval` | Create a appointments approval. |
| POST | `/appointments/{id}/cancellations` | `create_appointments_cancellation` | Create a appointments cancellation. |
| POST | `/appointments/{id}/follow_ups` | `create_appointments_follow_up` | Create a appointments follow up. |
| POST | `/appointments/{id}/rejections` | `create_appointments_rejection` | Create a appointments rejection. |
| POST | `/appointments/{id}/vehicle_reconciliations` | `create_appointments_vehicle_reconciliation` | Create a appointments vehicle reconciliation. |
| POST | `/appointments/{id}/work_orders` | `create_appointments_work_order` | Create a appointments work order. |
| GET | `/campaigns` | `list_campaigns` | List all campaigns, paginated via the Link header. |
| POST | `/campaigns` | `create_campaign` | Create a campaign. |
| GET | `/campaigns/{id}` | `show_campaign` | Show a campaign by ID. |
| POST | `/campaigns/{id}/duplicate` | `duplicate_campaign` | Duplicate |
| POST | `/campaigns/{id}/send_campaign` | `send_campaign` | Send campaign |
| PATCH | `/cash_drawer_banner` | `update_cash_drawer_banner` | Update a cash drawer banner by ID. |
| POST | `/catalog_cleanups/{catalog_cleanup_id}/applications` | `create_catalog_cleanups_application` | Create a catalog cleanups application. |
| GET | `/conversations` | `list_conversations` | List all conversations, paginated via the Link header. |
| POST | `/conversations/bulk_mark_read` | `create_conversations_bulk_mark_read` | Create a conversations bulk mark read. |
| POST | `/conversations/{conversation_id}/customer_links` | `create_conversations_customer_link` | Create a conversations customer link. |
| POST | `/conversations/{conversation_id}/ignores` | `create_conversations_ignore` | Create a conversations ignore. |
| GET | `/conversations/{conversation_id}/messages` | `list_conversations_messages` | List all conversations messages, paginated via the Link header. |
| POST | `/conversations/{conversation_id}/messages` | `create_conversations_message` | Create a conversations message. |
| POST | `/conversations/{conversation_id}/messages/{id}/resends` | `create_conversations_messages_resend` | Create a conversations messages resend. |
| GET | `/conversations/{id}` | `show_conversation` | Show a conversation by ID. |
| PATCH | `/conversations/{id}` | `update_conversation` | Update a conversation by ID. |
| GET | `/core_tax_rules` | `list_core_tax_rules` | List all core tax rules, paginated via the Link header. |
| PATCH | `/core_tax_rules/{id}` | `update_core_tax_rule` | Update a core tax rule by ID. |
| GET | `/counter_sales` | `list_counter_sales` | List all counter sales, paginated via the Link header. |
| POST | `/counter_sales` | `create_counter_sale` | Create a counter sale. |
| GET | `/counter_sales/{counter_sale_id}/line_items/brands` | `list_counter_sales_line_items_brands` | List all counter sales line items brands, paginated via the Link header. |
| GET | `/counter_sales/{id}` | `show_counter_sale` | Show a counter sale by ID. |
| PATCH | `/counter_sales/{id}` | `update_counter_sale` | Update a counter sale by ID. |
| GET | `/courtesy_cars` | `list_courtesy_cars` | List all courtesy cars, paginated via the Link header. |
| GET | `/current_location` | `list_current_location` | List all current location, paginated via the Link header. |
| PATCH | `/current_location` | `update_current_location` | Update a current location by ID. |
| GET | `/customer_tags` | `list_customer_tags` | List all customer tags, paginated via the Link header. |
| POST | `/customer_tags` | `create_customer_tag` | Create a customer tag. |
| GET | `/customer_tags/{id}` | `show_customer_tag` | Show a customer tag by ID. |
| PATCH | `/customer_tags/{id}` | `update_customer_tag` | Update a customer tag by ID. |
| PATCH | `/customer_tags/{id}/archive` | `update_customer_tags_archive` | Update a customer tags archive by ID. |
| PATCH | `/customer_tags/{id}/restore` | `update_customer_tags_restore` | Update a customer tags restore by ID. |
| PATCH | `/customer_tags/{id}/trash` | `update_customer_tags_trash` | Update a customer tags trash by ID. |
| GET | `/customers` | `list_customers` | List all customers, paginated via the Link header. |
| POST | `/customers` | `create_customer` | Create a customer. |
| GET | `/customers/1043910089/merge` | `list_customers_merge` | List all customers merge, paginated via the Link header. |
| GET | `/customers/check_duplicate` | `check_customer_duplicate` | Check duplicate |
| POST | `/customers/export` | `create_customers_export` | Create a customers export. |
| GET | `/customers/export/{id}/download` | `list_customers_export_download` | List all customers export download, paginated via the Link header. |
| GET | `/customers/lookup` | `lookup_customer` | Lookup |
| GET | `/customers/{customer_id}/drivers` | `list_customers_drivers` | List all customers drivers, paginated via the Link header. |
| POST | `/customers/{customer_id}/drivers` | `create_driver` | Create a driver. |
| DELETE | `/customers/{customer_id}/drivers/{id}` | `delete_driver` | Delete a driver by ID. |
| GET | `/customers/{customer_id}/drivers/{id}` | `show_driver` | Show a driver by ID. |
| PATCH | `/customers/{customer_id}/drivers/{id}` | `update_driver` | Update a driver by ID. |
| GET | `/customers/{customer_id}/statements` | `list_customers_statements` | List all customers statements, paginated via the Link header. |
| GET | `/customers/{customer_id}/vehicles` | `list_customers_vehicles` | List all customers vehicles, paginated via the Link header. |
| GET | `/customers/{customer_id}/vehicles/{vehicle_id}/history` | `list_customers_vehicles_history` | List all customers vehicles history, paginated via the Link header. |
| GET | `/customers/{customer_id}/work_orders` | `list_customers_work_orders` | List all customers work orders, paginated via the Link header. |
| GET | `/customers/{id}` | `show_customer` | Show a customer by ID. |
| PATCH | `/customers/{id}` | `update_customer` | Update a customer by ID. |
| PATCH | `/customers/{id}/archive` | `archive_customer` | Archive |
| POST | `/customers/{id}/merges` | `merge_customer` | Create |
| PATCH | `/customers/{id}/restore` | `restore_customer` | Restore |
| PATCH | `/customers/{id}/trash` | `trash_customer` | Trash |
| GET | `/expenses` | `list_expenses` | List all expenses, paginated via the Link header. |
| POST | `/expenses` | `create_expense` | Create a expense. |
| GET | `/expenses/data_transfer` | `list_expenses_data_transfer` | List all expenses data transfer, paginated via the Link header. |
| POST | `/expenses/export` | `create_expenses_export` | Create a expenses export. |
| GET | `/expenses/export/{id}/download` | `list_expenses_export_download` | List all expenses export download, paginated via the Link header. |
| POST | `/expenses/imports/commit` | `create_expenses_imports_commit` | Create a expenses imports commit. |
| GET | `/expenses/imports/template` | `list_expenses_imports_template` | List all expenses imports template, paginated via the Link header. |
| POST | `/expenses/imports/validate` | `create_expenses_imports_validate` | Create a expenses imports validate. |
| DELETE | `/expenses/{id}` | `delete_expense` | Delete a expense by ID. |
| GET | `/expenses/{id}` | `show_expense` | Show a expense by ID. |
| PATCH | `/expenses/{id}` | `update_expense` | Update a expense by ID. |
| GET | `/fleets` | `list_fleets` | List all fleets, paginated via the Link header. |
| POST | `/inspection_reports` | `create_inspection_report` | Create a inspection report. |
| DELETE | `/inspection_reports/{id}` | `delete_inspection_report` | Delete a inspection report by ID. |
| GET | `/inspection_reports/{id}` | `show_inspection_report` | Show a inspection report by ID. |
| PATCH | `/inspection_reports/{id}/complete` | `complete_inspection_report` | Complete |
| GET | `/inspection_reports/{id}/group` | `show_inspection_report_group` | Show a inspection report group by ID. |
| POST | `/inspection_reports/{id}/mark_all` | `mark_all_inspection_report` | Mark all |
| PATCH | `/inspection_reports/{id}/publish` | `publish_inspection_report` | Publish |
| PATCH | `/inspection_reports/{id}/reassign` | `reassign_inspection_report` | Reassign |
| PATCH | `/inspection_reports/{id}/reopen` | `reopen_inspection_report` | Reopen |
| PATCH | `/inspection_reports/{id}/reset` | `reset_inspection_report` | Reset |
| POST | `/inspection_reports/{id}/retry_recording` | `retry_inspection_report_recording` | Retry recording |
| PATCH | `/inspection_reports/{id}/unpublish` | `unpublish_inspection_report` | Unpublish |
| GET | `/inspections` | `list_inspections` | List all inspections, paginated via the Link header. |
| POST | `/inspections` | `create_inspection` | Create a inspection. |
| GET | `/inspections/{id}` | `show_inspection` | Show a inspection by ID. |
| PATCH | `/inspections/{id}` | `update_inspection` | Update a inspection by ID. |
| PATCH | `/inspections/{id}/archive` | `archive_inspection` | Archive |
| PATCH | `/inspections/{id}/remove_default` | `remove_default_inspection` | Remove default |
| PATCH | `/inspections/{id}/restore` | `restore_inspection` | Restore |
| PATCH | `/inspections/{id}/set_default` | `set_default_inspection` | Set default |
| PATCH | `/inspections/{id}/trash` | `trash_inspection` | Trash |
| GET | `/inventory_levels` | `list_inventory_levels` | List all inventory levels, paginated via the Link header. |
| POST | `/inventory_levels` | `create_inventory_level` | Create a inventory level. |
| GET | `/inventory_levels/barcode_lookup` | `list_inventory_levels_barcode_lookup` | List all inventory levels barcode lookup, paginated via the Link header. |
| DELETE | `/inventory_levels/{id}` | `delete_inventory_level` | Delete a inventory level by ID. |
| GET | `/inventory_levels/{id}` | `show_inventory_level` | Show a inventory level by ID. |
| PATCH | `/inventory_levels/{id}` | `update_inventory_level` | Update a inventory level by ID. |
| PATCH | `/inventory_levels/{id}/stock` | `update_inventory_levels_stock` | Update a inventory levels stock by ID. |
| GET | `/labor_matrices` | `list_labor_matrices` | List all labor matrices, paginated via the Link header. |
| POST | `/labor_matrices` | `create_labor_matrice` | Create a labor matrice. |
| DELETE | `/labor_matrices/{id}` | `delete_labor_matrice` | Delete a labor matrice by ID. |
| PATCH | `/labor_matrices/{id}` | `update_labor_matrice` | Update a labor matrice by ID. |
| GET | `/labor_rates` | `list_labor_rates` | List all labor rates, paginated via the Link header. |
| POST | `/labor_rates` | `create_labor_rate` | Create a labor rate. |
| PATCH | `/labor_rates/{id}/archive` | `archive_labor_rate` | Archive |
| PATCH | `/labor_rates/{id}/restore` | `restore_labor_rate` | Restore |
| PATCH | `/labor_rates/{id}/trash` | `trash_labor_rate` | Trash |
| GET | `/labor_templates` | `list_labor_templates` | List all labor templates, paginated via the Link header. |
| DELETE | `/labor_templates/{id}` | `delete_labor_template` | Delete a labor template by ID. |
| PATCH | `/labor_templates/{id}` | `update_labor_template` | Update a labor template by ID. |
| GET | `/lead_sources` | `list_lead_sources` | List all lead sources, paginated via the Link header. |
| POST | `/lead_sources` | `create_lead_source` | Create a lead source. |
| POST | `/lead_sources/seed_defaults` | `create_lead_sources_seed_default` | Create a lead sources seed default. |
| DELETE | `/lead_sources/{id}` | `delete_lead_source` | Delete a lead source by ID. |
| PATCH | `/lead_sources/{id}` | `update_lead_source` | Update a lead source by ID. |
| GET | `/locations/{id}` | `show_location` | Show a location by ID. |
| PATCH | `/locations/{id}` | `update_location` | Update a location by ID. |
| GET | `/locations/{id}/business_profile` | `list_locations_business_profile` | List all locations business profile, paginated via the Link header. |
| PATCH | `/locations/{id}/business_profile` | `update_locations_business_profile` | Update a locations business profile by ID. |
| GET | `/locations/{id}/operations` | `list_locations_operations` | List all locations operations, paginated via the Link header. |
| PATCH | `/locations/{id}/operations` | `update_locations_operations` | Update a locations operations by ID. |
| GET | `/locations/{location_id}/close_requirements` | `list_locations_close_requirements` | List all locations close requirements, paginated via the Link header. |
| PATCH | `/locations/{location_id}/close_requirements` | `update_locations_close_requirements` | Update a locations close requirements by ID. |
| GET | `/locations/{location_id}/courtesy_cars` | `list_locations_courtesy_cars` | List all locations courtesy cars, paginated via the Link header. |
| GET | `/locations/{location_id}/documents` | `list_locations_documents` | List all locations documents, paginated via the Link header. |
| PATCH | `/locations/{location_id}/documents` | `update_locations_documents` | Update a locations documents by ID. |
| GET | `/locations/{location_id}/lead_source_requirements` | `list_locations_lead_source_requirements` | List all locations lead source requirements, paginated via the Link header. |
| PATCH | `/locations/{location_id}/lead_source_requirements` | `update_locations_lead_source_requirements` | Update a locations lead source requirements by ID. |
| GET | `/locations/{location_id}/reminders` | `list_locations_reminders` | List all locations reminders, paginated via the Link header. |
| PATCH | `/locations/{location_id}/reminders` | `update_locations_reminders` | Update a locations reminders by ID. |
| GET | `/locations/{location_id}/schedule_config` | `list_locations_schedule_config` | List all locations schedule config, paginated via the Link header. |
| PATCH | `/locations/{location_id}/schedule_config` | `update_locations_schedule_config` | Update a locations schedule config by ID. |
| GET | `/me/api_tokens` | `list_me_api_tokens` | List all me api tokens, paginated via the Link header. |
| POST | `/me/api_tokens` | `create_me_api_token` | Create a me api token. |
| DELETE | `/me/api_tokens/{id}` | `delete_me_api_token` | Delete a me api token by ID. |
| GET | `/me/notifications` | `list_me_notifications` | List all me notifications, paginated via the Link header. |
| PATCH | `/me/notifications` | `update_me_notifications` | Update a me notifications by ID. |
| GET | `/me/preferences` | `list_me_preferences` | List all me preferences, paginated via the Link header. |
| PATCH | `/me/preferences` | `update_me_preferences` | Update a me preferences by ID. |
| GET | `/me/profile` | `list_me_profile` | List all me profile, paginated via the Link header. |
| PATCH | `/me/profile` | `update_me_profile` | Update a me profile by ID. |
| GET | `/notifications` | `list_notifications` | List all notifications, paginated via the Link header. |
| POST | `/notifications/bulk_mark_read` | `create_notifications_bulk_mark_read` | Create a notifications bulk mark read. |
| GET | `/notifications/{id}` | `show_notification` | Show a notification by ID. |
| PATCH | `/notifications/{id}` | `update_notification` | Update a notification by ID. |
| GET | `/orders/purchase_orders` | `list_orders_purchase_orders` | List all orders purchase orders, paginated via the Link header. |
| POST | `/orders/purchase_orders` | `create_orders_purchase_order` | Create a orders purchase order. |
| DELETE | `/orders/purchase_orders/{id}` | `delete_orders_purchase_order` | Delete a orders purchase order by ID. |
| GET | `/orders/purchase_orders/{id}` | `show_orders_purchase_order` | Show a orders purchase order by ID. |
| PATCH | `/orders/purchase_orders/{id}` | `update_orders_purchase_order` | Update a orders purchase order by ID. |
| POST | `/orders/purchase_orders/{purchase_order_id}/cancellations` | `create_orders_purchase_orders_cancellation` | Create a orders purchase orders cancellation. |
| GET | `/orders/return_orders` | `list_orders_return_orders` | List all orders return orders, paginated via the Link header. |
| DELETE | `/orders/return_orders/{id}` | `delete_orders_return_order` | Delete a orders return order by ID. |
| GET | `/orders/return_orders/{id}` | `show_orders_return_order` | Show a orders return order by ID. |
| PATCH | `/orders/return_orders/{id}` | `update_orders_return_order` | Update a orders return order by ID. |
| POST | `/orders/return_orders/{return_order_id}/refund_completions` | `create_orders_return_orders_refund_completion` | Create a orders return orders refund completion. |
| GET | `/orders/sublet_orders` | `list_orders_sublet_orders` | List all orders sublet orders, paginated via the Link header. |
| PATCH | `/orders/sublet_orders/mark_payment_complete` | `update_orders_sublet_orders_mark_payment_complete` | Update a orders sublet orders mark payment complete by ID. |
| GET | `/orders/sublet_orders/{id}` | `show_orders_sublet_order` | Show a orders sublet order by ID. |
| GET | `/packages` | `list_packages` | List all packages, paginated via the Link header. |
| POST | `/packages` | `create_package` | Create a package. |
| PATCH | `/packages/{id}` | `update_package` | Update a package by ID. |
| PATCH | `/packages/{id}/archive` | `archive_package` | Archive |
| POST | `/packages/{id}/duplicate` | `create_packages_duplicate` | Create a packages duplicate. |
| PATCH | `/packages/{id}/restore` | `restore_package` | Restore |
| PATCH | `/packages/{id}/trash` | `trash_package` | Trash |
| GET | `/parts_matrices` | `list_parts_matrices` | List all parts matrices, paginated via the Link header. |
| POST | `/parts_matrices` | `create_parts_matrice` | Create a parts matrice. |
| DELETE | `/parts_matrices/{id}` | `delete_parts_matrice` | Delete a parts matrice by ID. |
| PATCH | `/parts_matrices/{id}` | `update_parts_matrice` | Update a parts matrice by ID. |
| GET | `/payments` | `list_payments` | List all payments, paginated via the Link header. |
| GET | `/payments/pending` | `list_payments_pending` | List all payments pending, paginated via the Link header. |
| GET | `/payments/{id}` | `show_payment` | Show a payment by ID. |
| POST | `/payments/{id}/cancellation` | `create_payments_cancellation` | Create a payments cancellation. |
| POST | `/payments/{id}/confirmation` | `create_payments_confirmation` | Create a payments confirmation. |
| POST | `/payments/{id}/failure` | `create_payments_failure` | Create a payments failure. |
| GET | `/recent_searches` | `list_recent_searches` | List all recent searches, paginated via the Link header. |
| POST | `/recent_searches` | `create_recent_searche` | Create a recent searche. |
| DELETE | `/recent_searches/clear` | `delete_recent_searches_clear` | Delete a recent searches clear by ID. |
| DELETE | `/recent_searches/{id}` | `delete_recent_searche` | Delete a recent searche by ID. |
| GET | `/reports/accounting` | `list_reports_accounting` | List all reports accounting, paginated via the Link header. |
| GET | `/reports/ar_aging` | `list_reports_ar_aging` | List all reports ar aging, paginated via the Link header. |
| GET | `/reports/declined_work` | `list_reports_declined_work` | List all reports declined work, paginated via the Link header. |
| GET | `/reports/end_of_day` | `list_reports_end_of_day` | List all reports end of day, paginated via the Link header. |
| GET | `/reports/financial` | `list_reports_financial` | List all reports financial, paginated via the Link header. |
| GET | `/reports/open_work` | `list_reports_open_work` | List all reports open work, paginated via the Link header. |
| GET | `/reports/parts_purchases` | `list_reports_parts_purchases` | List all reports parts purchases, paginated via the Link header. |
| GET | `/reports/parts_usage` | `list_reports_parts_usage` | List all reports parts usage, paginated via the Link header. |
| GET | `/reports/performance` | `list_reports_performance` | List all reports performance, paginated via the Link header. |
| GET | `/reports/profit_and_loss` | `list_reports_profit_and_loss` | List all reports profit and loss, paginated via the Link header. |
| GET | `/reports/sales_summary` | `list_reports_sales_summary` | List all reports sales summary, paginated via the Link header. |
| GET | `/reports/service_categories` | `list_reports_service_categories` | List all reports service categories, paginated via the Link header. |
| GET | `/reports/store_credit` | `list_reports_store_credit` | List all reports store credit, paginated via the Link header. |
| GET | `/reports/technician_productivity` | `list_reports_technician_productivity` | List all reports technician productivity, paginated via the Link header. |
| GET | `/reports/work_order_profitability` | `list_reports_work_order_profitability` | List all reports work order profitability, paginated via the Link header. |
| GET | `/search` | `list_search` | List all search, paginated via the Link header. |
| GET | `/service_categories` | `list_service_categories` | List all service categories, paginated via the Link header. |
| POST | `/service_categories` | `create_service_category` | Create a service category. |
| POST | `/service_categories/seed_defaults` | `seed_defaults_service_categories` | Seed defaults |
| PATCH | `/service_categories/{id}` | `update_service_category` | Update a service category by ID. |
| PATCH | `/service_categories/{id}/archive` | `archive_service_category` | Archive |
| PATCH | `/service_categories/{id}/restore` | `restore_service_category` | Restore |
| PATCH | `/service_categories/{id}/trash` | `trash_service_category` | Trash |
| GET | `/shop_discounts` | `list_shop_discounts` | List all shop discounts, paginated via the Link header. |
| POST | `/shop_discounts` | `create_shop_discount` | Create a shop discount. |
| DELETE | `/shop_discounts/{id}` | `delete_shop_discount` | Delete a shop discount by ID. |
| PATCH | `/shop_discounts/{id}` | `update_shop_discount` | Update a shop discount by ID. |
| POST | `/shop_discounts/{id}/duplicate` | `create_shop_discounts_duplicate` | Create a shop discounts duplicate. |
| GET | `/shop_fees` | `list_shop_fees` | List all shop fees, paginated via the Link header. |
| POST | `/shop_fees` | `create_shop_fee` | Create a shop fee. |
| DELETE | `/shop_fees/{id}` | `delete_shop_fee` | Delete a shop fee by ID. |
| PATCH | `/shop_fees/{id}` | `update_shop_fee` | Update a shop fee by ID. |
| POST | `/shop_fees/{id}/duplicate` | `create_shop_fees_duplicate` | Create a shop fees duplicate. |
| GET | `/statements/{id}` | `show_statement` | Show a statement by ID. |
| GET | `/statements/{statement_id}/payments` | `list_statements_payments` | List all statements payments, paginated via the Link header. |
| POST | `/store_credits/{store_credit_id}/voids` | `create_store_credits_void` | Create a store credits void. |
| GET | `/sub_statuses` | `list_sub_statuses` | List all sub statuses, paginated via the Link header. |
| POST | `/sub_statuses` | `create_sub_statuse` | Create a sub statuse. |
| PATCH | `/sub_statuses/{id}` | `update_sub_statuse` | Update a sub statuse by ID. |
| GET | `/sublet_packages` | `list_sublet_packages` | List all sublet packages, paginated via the Link header. |
| POST | `/sublet_packages` | `create_sublet_package` | Create a sublet package. |
| DELETE | `/sublet_packages/{id}` | `delete_sublet_package` | Delete a sublet package by ID. |
| PATCH | `/sublet_packages/{id}` | `update_sublet_package` | Update a sublet package by ID. |
| PATCH | `/sublet_packages/{id}/deactivate` | `update_sublet_packages_deactivate` | Update a sublet packages deactivate by ID. |
| GET | `/tire_storage_slots` | `list_tire_storage_slots` | List all tire storage slots, paginated via the Link header. |
| POST | `/tire_storage_slots` | `create_tire_storage_slot` | Create a tire storage slot. |
| GET | `/tire_storage_slots/{id}` | `show_tire_storage_slot` | Show a tire storage slot by ID. |
| POST | `/tire_storage_slots/{tire_storage_slot_id}/check_outs` | `create_tire_storage_slots_check_out` | Create a tire storage slots check out. |
| GET | `/tires` | `list_tires` | List all tires, paginated via the Link header. |
| POST | `/tires` | `create_tire` | Create a tire. |
| DELETE | `/tires/{id}` | `delete_tire` | Delete a tire by ID. |
| GET | `/tires/{id}` | `show_tire` | Show a tire by ID. |
| PATCH | `/tires/{id}` | `update_tire` | Update a tire by ID. |
| GET | `/users` | `list_users` | List all users, paginated via the Link header. |
| POST | `/users` | `create_user` | Create a user. |
| GET | `/users/permission_groups` | `list_permission_groups` | List all permission groups, paginated via the Link header. |
| POST | `/users/permission_groups` | `create_permission_group` | Create a permission group. |
| PATCH | `/users/permission_groups/{id}` | `update_permission_group` | Update a permission group by ID. |
| DELETE | `/users/{id}` | `delete_user` | Delete a user by ID. |
| GET | `/users/{id}` | `show_user` | Show a user by ID. |
| PATCH | `/users/{id}` | `update_user` | Update a user by ID. |
| POST | `/users/{id}/disable` | `create_users_disable` | Create a users disable. |
| POST | `/users/{id}/enable` | `create_users_enable` | Create a users enable. |
| GET | `/users/{id}/qr_code` | `list_users_qr_code` | List all users qr code, paginated via the Link header. |
| POST | `/users/{id}/reset_pin` | `create_users_reset_pin` | Create a users reset pin. |
| POST | `/users/{id}/send_confirmation` | `create_users_send_confirmation` | Create a users send confirmation. |
| POST | `/users/{id}/send_password_reset` | `create_users_send_password_reset` | Create a users send password reset. |
| POST | `/users/{id}/unlock` | `create_users_unlock` | Create a users unlock. |
| GET | `/vehicles` | `list_vehicles` | List all vehicles, paginated via the Link header. |
| POST | `/vehicles` | `create_vehicle` | Create a vehicle. |
| GET | `/vehicles/autocomplete` | `list_vehicles_autocomplete` | List all vehicles autocomplete, paginated via the Link header. |
| GET | `/vehicles/check_duplicate` | `check_vehicle_duplicate` | Check duplicate |
| GET | `/vehicles/customer_vehicles` | `list_vehicles_customer_vehicles` | List all vehicles customer vehicles, paginated via the Link header. |
| GET | `/vehicles/lookup` | `lookup_vehicle` | Lookup |
| GET | `/vehicles/prefill` | `prefill_vehicle` | Prefill |
| GET | `/vehicles/vin_decode` | `decode_vin` | VIN decode |
| GET | `/vehicles/{id}` | `show_vehicle` | Show a vehicle by ID. |
| PATCH | `/vehicles/{id}` | `update_vehicle` | Update a vehicle by ID. |
| PATCH | `/vehicles/{id}/archive` | `archive_vehicle` | Archive |
| POST | `/vehicles/{id}/merges` | `merge_vehicle` | Create |
| PATCH | `/vehicles/{id}/restore` | `restore_vehicle` | Restore |
| POST | `/vehicles/{id}/transfers` | `transfer_vehicle` | Create |
| PATCH | `/vehicles/{id}/trash` | `trash_vehicle` | Trash |
| GET | `/vehicles/{vehicle_id}/work_orders` | `list_vehicles_work_orders` | List all vehicles work orders, paginated via the Link header. |
| GET | `/vendors` | `list_vendors` | List all vendors, paginated via the Link header. |
| POST | `/vendors` | `create_vendor` | Create a vendor. |
| GET | `/vendors/{id}` | `show_vendor` | Show a vendor by ID. |
| PATCH | `/vendors/{id}` | `update_vendor` | Update a vendor by ID. |
| PATCH | `/vendors/{id}/archive` | `archive_vendor` | Archive |
| PATCH | `/vendors/{id}/restore` | `restore_vendor` | Restore |
| PATCH | `/vendors/{id}/trash` | `trash_vendor` | Trash |
| GET | `/vendors/{vendor_id}/purchase_orders` | `list_vendors_purchase_orders` | List all vendors purchase orders, paginated via the Link header. |
| GET | `/work_order_tags` | `list_work_order_tags` | List all work order tags, paginated via the Link header. |
| POST | `/work_order_tags` | `create_work_order_tag` | Create a work order tag. |
| PATCH | `/work_order_tags/{id}` | `update_work_order_tag` | Update a work order tag by ID. |
| PATCH | `/work_order_tags/{id}/archive` | `update_work_order_tags_archive` | Update a work order tags archive by ID. |
| GET | `/work_orders` | `list_work_orders` | List all work orders, paginated via the Link header. |
| POST | `/work_orders` | `create_work_order` | Create a work order. |
| DELETE | `/work_orders/{id}` | `delete_work_order` | Delete a work order by ID. |
| GET | `/work_orders/{id}` | `show_work_order` | Show a work order by ID. |
| PATCH | `/work_orders/{id}` | `update_work_order` | Update a work order by ID. |
| POST | `/work_orders/{id}/account_postings` | `create_work_orders_account_posting` | Create a work orders account posting. |
| PATCH | `/work_orders/{id}/close` | `close_work_order` | Close |
| PATCH | `/work_orders/{id}/close_as_declined` | `update_work_orders_close_as_declined` | Update a work orders close as declined by ID. |
| PATCH | `/work_orders/{id}/close_as_voided` | `update_work_orders_close_as_voided` | Update a work orders close as voided by ID. |
| DELETE | `/work_orders/{id}/courtesy_car_assignment` | `delete_work_orders_courtesy_car_assignment` | Delete a work orders courtesy car assignment by ID. |
| PATCH | `/work_orders/{id}/courtesy_car_assignment` | `update_work_orders_courtesy_car_assignment` | Update a work orders courtesy car assignment by ID. |
| POST | `/work_orders/{id}/courtesy_car_assignment` | `create_work_orders_courtesy_car_assignment` | Create a work orders courtesy car assignment. |
| GET | `/work_orders/{id}/declined_services` | `show_work_order_declined_services` | Show a work order declined services by ID. |
| POST | `/work_orders/{id}/estimate_deliveries` | `create_work_orders_estimate_deliverie` | Create a work orders estimate deliverie. |
| POST | `/work_orders/{id}/invoice_summaries` | `create_work_orders_invoice_summarie` | Create a work orders invoice summarie. |
| POST | `/work_orders/{id}/reminders` | `create_work_orders_reminder` | Create a work orders reminder. |
| PATCH | `/work_orders/{id}/reopen` | `reopen_work_order` | Reopen |
| POST | `/work_orders/{id}/service_declines` | `create_work_orders_service_decline` | Create a work orders service decline. |
| GET | `/work_orders/{id}/service_history` | `show_work_order_service_history` | Show a work order service history by ID. |
| POST | `/work_orders/{work_order_id}/authorization_decisions` | `create_work_orders_authorization_decision` | Create a work orders authorization decision. |
| POST | `/work_orders/{work_order_id}/authorizations` | `create_work_orders_authorization` | Create a work orders authorization. |
| GET | `/work_orders/{work_order_id}/concerns` | `list_work_orders_concerns` | List all work orders concerns, paginated via the Link header. |
| GET | `/work_orders/{work_order_id}/estimate` | `show_work_order_estimate` | Show a work order estimate by ID. |
| POST | `/work_orders/{work_order_id}/fee_exclusions` | `create_work_orders_fee_exclusion` | Create a work orders fee exclusion. |
| DELETE | `/work_orders/{work_order_id}/fee_exclusions/{id}` | `delete_work_orders_fee_exclusion` | Delete a work orders fee exclusion by ID. |
| GET | `/work_orders/{work_order_id}/inspection` | `show_work_order_inspection` | Show a work order inspection by ID. |
| GET | `/work_orders/{work_order_id}/parts` | `show_work_order_parts` | Show a work order parts by ID. |
| GET | `/work_orders/{work_order_id}/payments` | `show_work_order_payments` | Show a work order payments by ID. |
| POST | `/work_orders/{work_order_id}/payments` | `create_work_order_payment` | Create a work order payment. |
| DELETE | `/work_orders/{work_order_id}/payments/reverse_ar` | `reverse_work_order_payment_ar` | Reverse ar |
| POST | `/work_orders/{work_order_id}/payments/send_to_ar` | `send_work_order_payment_to_ar` | Send to ar |
| POST | `/work_orders/{work_order_id}/refunds` | `create_work_orders_refund` | Create a work orders refund. |
| POST | `/work_orders/{work_order_id}/services` | `create_work_orders_service` | Create a work orders service. |
| PATCH | `/work_orders/{work_order_id}/services/reorder` | `update_work_orders_services_reorder` | Update a work orders services reorder by ID. |
| DELETE | `/work_orders/{work_order_id}/services/{id}` | `delete_work_orders_service` | Delete a work orders service by ID. |
| PATCH | `/work_orders/{work_order_id}/services/{id}` | `update_work_orders_service` | Update a work orders service by ID. |
| GET | `/work_orders/{work_order_id}/services/{id}/adjust_time` | `list_work_orders_services_adjust_time` | List all work orders services adjust time, paginated via the Link header. |
| PATCH | `/work_orders/{work_order_id}/services/{id}/adjust_time` | `update_work_orders_services_adjust_time` | Update a work orders services adjust time by ID. |
| DELETE | `/work_orders/{work_order_id}/services/{id}/authorization` | `delete_work_orders_services_authorization` | Delete a work orders services authorization by ID. |
| DELETE | `/work_orders/{work_order_id}/services/{id}/completion` | `delete_work_orders_services_completion` | Delete a work orders services completion by ID. |
| POST | `/work_orders/{work_order_id}/services/{id}/completion` | `create_work_orders_services_completion` | Create a work orders services completion. |
| POST | `/work_orders/{work_order_id}/services/{id}/copies` | `create_work_orders_services_copie` | Create a work orders services copie. |
| POST | `/work_orders/{work_order_id}/services/{id}/packages` | `create_work_orders_services_package` | Create a work orders services package. |
| PATCH | `/work_orders/{work_order_id}/services/{id}/pause` | `update_work_orders_services_pause` | Update a work orders services pause by ID. |
| PATCH | `/work_orders/{work_order_id}/services/{id}/publish` | `update_work_orders_services_publish` | Update a work orders services publish by ID. |
| PATCH | `/work_orders/{work_order_id}/services/{id}/revive` | `update_work_orders_services_revive` | Update a work orders services revive by ID. |
| POST | `/work_orders/{work_order_id}/services/{id}/time_entries` | `create_work_orders_services_time_entrie` | Create a work orders services time entrie. |
| PATCH | `/work_orders/{work_order_id}/services/{id}/toggle_labor_completion` | `update_work_orders_services_toggle_labor_completion` | Update a work orders services toggle labor completion by ID. |
| PATCH | `/work_orders/{work_order_id}/services/{id}/update_category` | `update_work_orders_services_update_category` | Update a work orders services update category by ID. |
| POST | `/work_orders/{work_order_id}/services/{service_id}/line_items` | `create_work_orders_services_line_item` | Create a work orders services line item. |
| DELETE | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}` | `delete_work_orders_services_line_item` | Delete a work orders services line item by ID. |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}` | `update_work_orders_services_line_item` | Update a work orders services line item by ID. |
| POST | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/copies` | `create_work_orders_services_line_items_copie` | Create a work orders services line items copie. |
| POST | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/inventory_additions` | `create_work_orders_services_line_items_inventory_addition` | Create a work orders services line items inventory addition. |
| POST | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/price_refreshes` | `create_work_orders_services_line_items_price_refreshe` | Create a work orders services line items price refreshe. |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/pull` | `pull_work_order_service_line_item` | Pull |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_pull` | `undo_pull_work_order_service_line_item` | Undo pull |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_return` | `undo_return_work_order_service_line_item` | Undo return |
| GET | `/work_orders/{work_order_id}/summary/activity` | `list_work_orders_summary_activity` | List all work orders summary activity, paginated via the Link header. |
| GET | `/work_orders/{work_order_id}/summary/appointments` | `list_work_orders_summary_appointments` | List all work orders summary appointments, paginated via the Link header. |
| GET | `/work_orders/{work_order_id}/summary/authorization_logs` | `list_work_orders_summary_authorization_logs` | List all work orders summary authorization logs, paginated via the Link header. |
| GET | `/work_orders/{work_order_id}/summary/vehicle_history` | `list_work_orders_summary_vehicle_history` | List all work orders summary vehicle history, paginated via the Link header. |
| POST | `/work_orders/{work_order_id}/voids` | `create_work_orders_void` | Create a work orders void. |
| GET | `/work_orders/{work_order_id}/wip` | `show_work_order_wip` | Show a work order wip by ID. |

For common error and pagination shapes, see [errors](errors.md) and [pagination](pagination.md).
