# Wenmar Pro API Reference

Base URL: `https://app.wenmarpro.com`

All requests require a bearer token (see [authentication](authentication.md)).
Paths use no prefix — endpoints live at root paths (`/customers`, `/vehicles/{id}`).
Responses are bare objects/arrays — no `{ "data": ... }` wrapper.

## Endpoints

| Method | Path | Operation | Description |
|---|---|---|---|
| GET | `/account` | `list_account` | Show the current account |
| GET | `/customers` | `list_customers` | List customers (paginated) |
| POST | `/customers` | `create_customer` | Create a customer |
| GET | `/customers/{id}` | `show_customer` | Show a customer |
| PATCH | `/customers/{id}` | `update_customer` | Update a customer |
| GET | `/locations/{id}` | `show_location` | Show a location |
| GET | `/vehicles` | `list_vehicles` | List vehicles (paginated) |
| POST | `/vehicles` | `create_vehicle` | Create a vehicle |
| GET | `/vehicles/{id}` | `show_vehicle` | Show a vehicle |
| PATCH | `/vehicles/{id}` | `update_vehicle` | Update a vehicle |
| DELETE | `/vehicles/{id}` | `delete_vehicle` | Delete a vehicle |
| GET | `/vehicles/vin_decode` | `decode_vin` | Decode a VIN |
| GET | `/vehicles/check_duplicate` | `check_duplicate` | Check for duplicate vehicles |
| GET | `/work_orders` | `list_work_orders` | List work orders (paginated) |
| POST | `/work_orders` | `create_work_order` | Create a work order |
| GET | `/work_orders/{id}` | `show_work_order` | Show a work order |
| PATCH | `/work_orders/{id}` | `update_work_order` | Update a work order |
| DELETE | `/work_orders/{id}` | `delete_work_order` | Delete a work order |

For detailed per-resource docs, see [sections/](sections/).

## Common response shapes

### Error

```json
{ "error": { "code": "not_found", "message": "Customer not found", "details": {} } }
```

### Pagination (Link header)

```
Link: <https://app.wenmarpro.com/customers?page=2>; rel="next"
```

See [pagination](pagination.md) for details.
