# Wenmar Pro API Reference

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

Base URL: `https://app.wenmarpro.com`

For detailed per-resource docs, see [sections/](sections/).

## Endpoints

| Method | Path | Operation | Description |
|---|---|---|---|
| GET | `/account` | `list_account` | List all account, paginated via the Link header. |
| GET | `/conversations` | `list_conversations` | List all conversations, paginated via the Link header. |
| POST | `/conversations/bulk_mark_read` | `create_conversation` | Create a conversation. |
| POST | `/conversations/{conversation_id}/customer_links` | `post_conversations_customer_link` | Create |
| POST | `/conversations/{conversation_id}/ignores` | `post_conversations_ignore` | Create |
| GET | `/conversations/{conversation_id}/messages` | `get_conversations_message` | Index |
| POST | `/conversations/{conversation_id}/messages` | `post_conversations_message` | Create |
| POST | `/conversations/{conversation_id}/messages/{id}/resends` | `post_conversations_message_resends` | Create |
| GET | `/conversations/{id}` | `show_conversation` | Show a conversation by ID. |
| PATCH | `/conversations/{id}` | `update_conversation` | Update a conversation by ID. |
| GET | `/counter_sales` | `list_counter_sales` | List all counter sales, paginated via the Link header. |
| POST | `/counter_sales` | `create_counter_sale` | Create a counter sale. |
| GET | `/counter_sales/{id}` | `show_counter_sale` | Show a counter sale by ID. |
| PATCH | `/counter_sales/{id}` | `update_counter_sale` | Update a counter sale by ID. |
| GET | `/courtesy_cars` | `list_courtesy_cars` | List all courtesy cars, paginated via the Link header. |
| GET | `/customer_tags` | `list_customer_tags` | List all customer tags, paginated via the Link header. |
| POST | `/customer_tags` | `create_customer_tag` | Create a customer tag. |
| DELETE | `/customer_tags/{id}` | `delete_customer_tag` | Delete a customer tag by ID. |
| PATCH | `/customer_tags/{id}` | `update_customer_tag` | Update a customer tag by ID. |
| GET | `/customers` | `list_customers` | List all customers, paginated via the Link header. |
| POST | `/customers` | `create_customer` | Create a customer. |
| GET | `/customers/check_duplicate` | `check_customer_duplicate` | Check duplicate |
| GET | `/customers/lookup` | `lookup_customer` | Lookup |
| GET | `/customers/{customer_id}/drivers` | `list_customers_drivers` | List all customers drivers, paginated via the Link header. |
| POST | `/customers/{customer_id}/drivers` | `create_driver` | Create a driver. |
| DELETE | `/customers/{customer_id}/drivers/{id}` | `delete_driver` | Delete a driver by ID. |
| GET | `/customers/{customer_id}/drivers/{id}` | `show_driver` | Show a driver by ID. |
| PATCH | `/customers/{customer_id}/drivers/{id}` | `update_driver` | Update a driver by ID. |
| GET | `/customers/{customer_id}/statements` | `list_customers_statements` | List all customers statements, paginated via the Link header. |
| GET | `/customers/{customer_id}/vehicles` | `list_customers_vehicles` | List all customers vehicles, paginated via the Link header. |
| GET | `/customers/{customer_id}/vehicles/{vehicle_id}/history` | `get_customers_vehicle_history` | Show |
| GET | `/customers/{customer_id}/work_orders` | `list_customers_work_orders` | List all customers work orders, paginated via the Link header. |
| DELETE | `/customers/{id}` | `delete_customer` | Delete a customer by ID. |
| GET | `/customers/{id}` | `show_customer` | Show a customer by ID. |
| PATCH | `/customers/{id}` | `update_customer` | Update a customer by ID. |
| POST | `/customers/{id}/merges` | `merge_customer` | Create |
| GET | `/expenses` | `list_expenses` | List all expenses, paginated via the Link header. |
| POST | `/expenses` | `create_expense` | Create a expense. |
| DELETE | `/expenses/{id}` | `delete_expense` | Delete a expense by ID. |
| PATCH | `/expenses/{id}` | `update_expense` | Update a expense by ID. |
| GET | `/fleets` | `list_fleets` | List all fleets, paginated via the Link header. |
| GET | `/locations/{id}` | `show_location` | Show a location by ID. |
| GET | `/notifications` | `list_notifications` | List all notifications, paginated via the Link header. |
| POST | `/notifications/bulk_mark_read` | `create_notification` | Create a notification. |
| GET | `/notifications/{id}` | `show_notification` | Show a notification by ID. |
| PATCH | `/notifications/{id}` | `update_notification` | Update a notification by ID. |
| GET | `/packages` | `list_packages` | List all packages, paginated via the Link header. |
| POST | `/packages` | `create_package` | Create a package. |
| PATCH | `/packages/{id}` | `update_package` | Update a package by ID. |
| POST | `/packages/{id}/duplicate` | `post_packages_duplicate` | Duplicate |
| GET | `/service_categories` | `list_service_categories` | List all service categories, paginated via the Link header. |
| POST | `/service_categories` | `create_service_category` | Create a service category. |
| POST | `/service_categories/seed_defaults` | `seed_defaults_service_categories` | Seed defaults |
| DELETE | `/service_categories/{id}` | `delete_service_category` | Delete a service category by ID. |
| PATCH | `/service_categories/{id}` | `update_service_category` | Update a service category by ID. |
| GET | `/settings/tags` | `list_tags` | List all tags, paginated via the Link header. |
| PATCH | `/settings/tags` | `update_tags` | Update a tags by ID. |
| GET | `/statements/{id}` | `show_statement` | Show a statement by ID. |
| GET | `/sublet_packages` | `list_sublet_packages` | List all sublet packages, paginated via the Link header. |
| POST | `/sublet_packages` | `create_sublet_package` | Create a sublet package. |
| DELETE | `/sublet_packages/{id}` | `delete_sublet_package` | Delete a sublet package by ID. |
| PATCH | `/sublet_packages/{id}` | `update_sublet_package` | Update a sublet package by ID. |
| GET | `/team/members` | `list_team` | List all team, paginated via the Link header. |
| GET | `/tire_storage_slots` | `list_tire_storage_slots` | List all tire storage slots, paginated via the Link header. |
| POST | `/tire_storage_slots` | `create_tire_storage_slot` | Create a tire storage slot. |
| GET | `/tire_storage_slots/{id}` | `show_tire_storage_slot` | Show a tire storage slot by ID. |
| POST | `/tire_storage_slots/{tire_storage_slot_id}/check_outs` | `post_tire_storage_slots_check_out` | Create |
| GET | `/tires` | `list_tires` | List all tires, paginated via the Link header. |
| POST | `/tires` | `create_tire` | Create a tire. |
| DELETE | `/tires/{id}` | `delete_tire` | Delete a tire by ID. |
| GET | `/tires/{id}` | `show_tire` | Show a tire by ID. |
| PATCH | `/tires/{id}` | `update_tire` | Update a tire by ID. |
| GET | `/vehicle_tags` | `list_vehicle_tags` | List all vehicle tags, paginated via the Link header. |
| POST | `/vehicle_tags` | `create_vehicle_tag` | Create a vehicle tag. |
| DELETE | `/vehicle_tags/{id}` | `delete_vehicle_tag` | Delete a vehicle tag by ID. |
| PATCH | `/vehicle_tags/{id}` | `update_vehicle_tag` | Update a vehicle tag by ID. |
| GET | `/vehicles` | `list_vehicles` | List all vehicles, paginated via the Link header. |
| POST | `/vehicles` | `create_vehicle` | Create a vehicle. |
| GET | `/vehicles/check_duplicate` | `check_vehicle_duplicate` | Check duplicate |
| GET | `/vehicles/lookup` | `lookup_vehicle` | Lookup |
| GET | `/vehicles/prefill` | `prefill_vehicle` | Prefill |
| GET | `/vehicles/vin_decode` | `decode_vin` | VIN decode |
| DELETE | `/vehicles/{id}` | `delete_vehicle` | Delete a vehicle by ID. |
| GET | `/vehicles/{id}` | `show_vehicle` | Show a vehicle by ID. |
| PATCH | `/vehicles/{id}` | `update_vehicle` | Update a vehicle by ID. |
| POST | `/vehicles/{id}/merges` | `merge_vehicle` | Create |
| POST | `/vehicles/{id}/transfers` | `transfer_vehicle` | Create |
| GET | `/vehicles/{vehicle_id}/work_orders` | `list_vehicles_work_orders` | List all vehicles work orders, paginated via the Link header. |
| GET | `/vendors` | `list_vendors` | List all vendors, paginated via the Link header. |
| POST | `/vendors` | `create_vendor` | Create a vendor. |
| DELETE | `/vendors/{id}` | `delete_vendor` | Delete a vendor by ID. |
| GET | `/vendors/{id}` | `show_vendor` | Show a vendor by ID. |
| PATCH | `/vendors/{id}` | `update_vendor` | Update a vendor by ID. |
| GET | `/vendors/{vendor_id}/purchase_orders` | `get_vendors_purchase_order` | Index |
| GET | `/work_orders` | `list_work_orders` | List all work orders, paginated via the Link header. |
| POST | `/work_orders` | `create_work_order` | Create a work order. |
| DELETE | `/work_orders/{id}` | `delete_work_order` | Delete a work order by ID. |
| GET | `/work_orders/{id}` | `show_work_order` | Show a work order by ID. |
| PATCH | `/work_orders/{id}` | `update_work_order` | Update a work order by ID. |
| PATCH | `/work_orders/{id}/close` | `close_work_order` | Close |
| PATCH | `/work_orders/{id}/close_as_paid` | `close_work_order_as_paid` | Close as paid |
| PATCH | `/work_orders/{id}/close_zero` | `close_work_order_zero` | Close zero |
| POST | `/work_orders/{id}/courtesy_car_assignment` | `post_work_orders_courtesy_car_assignment` | Create |
| PATCH | `/work_orders/{id}/decline_all` | `decline_all_work_order_services` | Decline all |
| GET | `/work_orders/{id}/declined_services` | `show_work_order_declined_services` | Show a work order declined services by ID. |
| DELETE | `/work_orders/{id}/hard_delete` | `delete_work_orders_hard_delete` | Delete a work orders hard delete by ID. |
| PATCH | `/work_orders/{id}/post_to_account` | `patch_work_orders_post_to_account` | Post to account |
| PATCH | `/work_orders/{id}/reopen` | `reopen_work_order` | Reopen |
| PATCH | `/work_orders/{id}/return_to_board` | `return_work_order_to_board` | Return to board |
| PATCH | `/work_orders/{id}/save_for_later` | `save_work_order_for_later` | Save for later |
| PATCH | `/work_orders/{id}/send_estimate` | `patch_work_orders_send_estimate` | Send estimate |
| PATCH | `/work_orders/{id}/send_invoice_summary` | `patch_work_orders_send_invoice_summary` | Send invoice summary |
| PATCH | `/work_orders/{id}/send_reminder` | `patch_work_orders_send_reminder` | Send reminder |
| GET | `/work_orders/{id}/service_history` | `show_work_order_service_history` | Show a work order service history by ID. |
| PATCH | `/work_orders/{id}/toggle_waiting_for_customer` | `patch_work_orders_toggle_waiting_for_customer` | Toggle waiting for customer |
| POST | `/work_orders/{work_order_id}/authorization` | `create_work_order_authorization` | Create a work order authorization. |
| POST | `/work_orders/{work_order_id}/authorization/update_decisions` | `update_work_order_authorization_decisions` | Update a work order authorization decisions by ID. |
| GET | `/work_orders/{work_order_id}/concerns` | `get_work_orders_concern` | Index |
| GET | `/work_orders/{work_order_id}/estimate` | `show_work_order_estimate` | Show a work order estimate by ID. |
| GET | `/work_orders/{work_order_id}/inspection` | `show_work_order_inspection` | Show a work order inspection by ID. |
| POST | `/work_orders/{work_order_id}/labels` | `post_work_orders_label` | Create |
| DELETE | `/work_orders/{work_order_id}/labels/{id}` | `delete_work_orders_label` | Delete a work orders label by ID. |
| GET | `/work_orders/{work_order_id}/parts` | `show_work_order_parts` | Show a work order parts by ID. |
| GET | `/work_orders/{work_order_id}/payments` | `show_work_order_payments` | Show a work order payments by ID. |
| POST | `/work_orders/{work_order_id}/payments` | `create_work_order_payment` | Create a work order payment. |
| DELETE | `/work_orders/{work_order_id}/payments/reverse_ar` | `reverse_work_order_payment_ar` | Reverse ar |
| POST | `/work_orders/{work_order_id}/payments/send_to_ar` | `send_work_order_payment_to_ar` | Send to ar |
| POST | `/work_orders/{work_order_id}/services` | `post_work_orders_service` | Create |
| DELETE | `/work_orders/{work_order_id}/services/{id}` | `delete_work_orders_service` | Delete a work orders service by ID. |
| PATCH | `/work_orders/{work_order_id}/services/{id}` | `patch_work_orders_service` | Update |
| PATCH | `/work_orders/{work_order_id}/services/{id}/acknowledge_parts` | `patch_work_orders_service_acknowledge_parts` | Acknowledge parts |
| POST | `/work_orders/{work_order_id}/services/{id}/add_line_item` | `post_work_orders_service_add_line_item` | Add line item |
| POST | `/work_orders/{work_order_id}/services/{id}/add_package` | `post_work_orders_service_add_package` | Add package |
| GET | `/work_orders/{work_order_id}/services/{id}/adjust_time` | `get_work_orders_service_adjust_time` | Adjust time |
| PATCH | `/work_orders/{work_order_id}/services/{id}/adjust_time` | `patch_work_orders_service_adjust_time` | Adjust time |
| PATCH | `/work_orders/{work_order_id}/services/{id}/apply_discount` | `patch_work_orders_service_apply_discount` | Apply discount |
| PATCH | `/work_orders/{work_order_id}/services/{id}/complete_service` | `patch_work_orders_service_complete_service` | Complete service |
| POST | `/work_orders/{work_order_id}/services/{id}/duplicate` | `post_work_orders_service_duplicate` | Duplicate |
| PATCH | `/work_orders/{work_order_id}/services/{id}/pause` | `patch_work_orders_service_pause` | Pause |
| PATCH | `/work_orders/{work_order_id}/services/{id}/publish` | `patch_work_orders_service_publish` | Publish |
| PATCH | `/work_orders/{work_order_id}/services/{id}/reset_completion` | `patch_work_orders_service_reset_completion` | Reset completion |
| PATCH | `/work_orders/{work_order_id}/services/{id}/revive` | `patch_work_orders_service_revive` | Revive |
| PATCH | `/work_orders/{work_order_id}/services/{id}/start` | `patch_work_orders_service_start` | Start |
| PATCH | `/work_orders/{work_order_id}/services/{id}/toggle_labor_completion` | `patch_work_orders_service_toggle_labor_completion` | Toggle labor completion |
| PATCH | `/work_orders/{work_order_id}/services/{id}/toggle_labor_tax` | `patch_work_orders_service_toggle_labor_tax` | Toggle labor tax |
| POST | `/work_orders/{work_order_id}/services/{id}/unauthorize` | `post_work_orders_service_unauthorize` | Unauthorize |
| PATCH | `/work_orders/{work_order_id}/services/{id}/update_category` | `patch_work_orders_service_update_category` | Update category |
| PATCH | `/work_orders/{work_order_id}/services/{id}/update_pricing_mode` | `patch_work_orders_service_update_pricing_mode` | Update pricing mode |
| PATCH | `/work_orders/{work_order_id}/services/{id}/update_service_technician` | `patch_work_orders_service_update_service_technician` | Update service technician |
| POST | `/work_orders/{work_order_id}/services/{service_id}/line_items` | `post_work_orders_service_line_items` | Create |
| DELETE | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}` | `delete_work_orders_service_line_items` | Delete a work orders service line items by ID. |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}` | `patch_work_orders_service_line_items` | Update |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/add_to_inventory` | `add_work_order_service_line_item_to_inventory` | Add to inventory |
| POST | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/duplicate` | `duplicate_work_order_service_line_item` | Duplicate |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/pull` | `pull_work_order_service_line_item` | Pull |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/refresh_price` | `refresh_work_order_service_line_item_price` | Refresh price |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_pull` | `undo_pull_work_order_service_line_item` | Undo pull |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/undo_return` | `undo_return_work_order_service_line_item` | Undo return |
| PATCH | `/work_orders/{work_order_id}/services/{service_id}/line_items/{id}/update_part_status` | `update_work_order_service_line_item_part_status` | Update a work order service line item part status by ID. |
| GET | `/work_orders/{work_order_id}/wip` | `show_work_order_wip` | Show a work order wip by ID. |

For common error and pagination shapes, see [errors](errors.md) and [pagination](pagination.md).
