# Team

<!-- AUTO-GENERATED from spec/openapi.enriched.yaml. Do not edit.
     Run: make docs -->

## List team permission groups

```
GET /team/permission_groups
```

List all team permission groups, paginated via the Link header.

**Response 200** — array

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `role` | string \| null | Yes |
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

**Response 403** — [Error](#error-schema) error envelope

```bash
curl -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" https://app.wenmarpro.com/team/permission_groups.json
```

## Create team permission group

```
POST /team/permission_groups
```

Create a team permission group.

**Response 201**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `role` | string \| null | Yes |
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X POST -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/team/permission_groups.json
```

## Update team permission group

```
PATCH /team/permission_groups/{id}
```

Update a team permission group by ID.

| Param | Type | Required |
|---|---|---|
| `id` | integer | Yes |

**Response 200**

| Field | Type | Required |
|---|---|---|
| `id` | integer | Yes |
| `name` | string | Yes |
| `description` | string \| null | Yes |
| `role` | string \| null | Yes |
| `can_perform_work` | boolean | Yes |
| `can_dispatch_work` | boolean | Yes |
| `can_message_customers` | boolean | Yes |
| `can_manage_technicians` | boolean | Yes |
| `can_override_inspections` | boolean | Yes |
| `can_perform_inspections` | boolean | Yes |
| `can_view_all_active_work_orders` | boolean | Yes |
| `can_close_reopen_work_orders` | boolean | Yes |
| `can_hard_delete_work_orders` | boolean | Yes |
| `can_edit_permissions` | boolean | Yes |
| `can_view_job_board` | boolean | Yes |
| `can_view_metrics` | boolean | Yes |
| `can_view_activity_feed` | boolean | Yes |
| `created_at` | string | Yes |
| `updated_at` | string | Yes |
| `url` | string | Yes |
| `app_url` | string | Yes |

```bash
curl -X PATCH -H "User-Agent: wenmar-cli/0.2" -H "Authorization: Bearer $WENMAR_TOKEN" -H "Content-Type: application/json" \
     -d '{"...":"..."}' https://app.wenmarpro.com/team/permission_groups/<id>.json
```

---

### Error schema {#error-schema}

| Field | Type | Required |
|---|---|---|
| `code` | string | Yes |
| `message` | string | Yes |
| `field_errors` | object | Yes |

`field_errors` — object:

---

### CreateTeamPermissionGroupRequest schema {#createteampermissiongrouprequest-schema}

| Field | Type | Required |
|---|---|---|
| `permission_group` | object | Yes |

`permission_group` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |
| `can_perform_work` | boolean | Yes |

---

### UpdateTeamPermissionGroupRequest schema {#updateteampermissiongrouprequest-schema}

| Field | Type | Required |
|---|---|---|
| `permission_group` | object | Yes |

`permission_group` — object:
| Field | Type | Required |
|---|---|---|
| `name` | string | Yes |

