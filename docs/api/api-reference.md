# Wenmar Pro API Reference

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

Base URL: `https://app.wenmarpro.com`

For detailed per-resource docs, see [sections/](sections/).

## Endpoints

| Method | Path | Operation | Description |
|---|---|---|---|
| GET | `/account` | `list_account` | List all account, paginated via the Link header. |
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
| POST | `/customers/{id}/merge` | `merge_customer` | Perform merge |
| GET | `/locations/{id}` | `show_location` | Show a location by ID. |
| GET | `/service_categories` | `list_service_categories` | List all service categories, paginated via the Link header. |
| POST | `/service_categories` | `create_service_category` | Create a service category. |
| POST | `/service_categories/seed_defaults` | `seed_defaults_service_categories` | Seed defaults |
| DELETE | `/service_categories/{id}` | `delete_service_category` | Delete a service category by ID. |
| PATCH | `/service_categories/{id}` | `update_service_category` | Update a service category by ID. |
| PATCH | `/service_categories/{id}/deactivate` | `deactivate_service_category` | Deactivate |
| PATCH | `/service_categories/{id}/move_down` | `move_down_service_category` | Move down |
| PATCH | `/service_categories/{id}/move_up` | `move_up_service_category` | Move up |
| PATCH | `/service_categories/{id}/reactivate` | `reactivate_service_category` | Reactivate |
| GET | `/settings/tags` | `list_tags` | List all tags, paginated via the Link header. |
| PATCH | `/settings/tags` | `update_tags` | Update a tags by ID. |
| GET | `/statements/{id}` | `show_statement` | Show a statement by ID. |
| GET | `/team/members` | `list_team` | List all team, paginated via the Link header. |
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
| POST | `/vehicles/{id}/merge` | `merge_vehicle` | Perform merge |
| PATCH | `/vehicles/{id}/transfer` | `transfer_vehicle` | Perform transfer |
| GET | `/vehicles/{vehicle_id}/work_orders` | `list_vehicles_work_orders` | List all vehicles work orders, paginated via the Link header. |
| GET | `/vendors` | `list_vendors` | List all vendors, paginated via the Link header. |
| GET | `/vendors/{id}` | `show_vendor` | Show a vendor by ID. |
| GET | `/work_orders` | `list_work_orders` | List all work orders, paginated via the Link header. |
| POST | `/work_orders` | `create_work_order` | Create a work order. |
| DELETE | `/work_orders/{id}` | `delete_work_order` | Delete a work order by ID. |
| GET | `/work_orders/{id}` | `show_work_order` | Show a work order by ID. |
| PATCH | `/work_orders/{id}` | `update_work_order` | Update a work order by ID. |
| GET | `/work_orders/{work_order_id}/estimate` | `show_work_order_estimate` | Show a work order estimate by ID. |
| GET | `/work_orders/{work_order_id}/inspection` | `show_work_order_inspection` | Show a work order inspection by ID. |
| GET | `/work_orders/{work_order_id}/parts` | `show_work_order_parts` | Show a work order parts by ID. |
| GET | `/work_orders/{work_order_id}/payments` | `show_work_order_payments` | Show a work order payments by ID. |
| POST | `/work_orders/{work_order_id}/payments` | `create_work_order_payment` | Create a work order payment. |
| GET | `/work_orders/{work_order_id}/summary` | `get_work_orders_summary` | Show |
| GET | `/work_orders/{work_order_id}/wip` | `show_work_order_wip` | Show a work order wip by ID. |

For common error and pagination shapes, see [errors](errors.md) and [pagination](pagination.md).
