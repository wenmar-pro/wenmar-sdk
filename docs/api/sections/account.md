# Account

## Show account

```
GET /account
```

Returns the account associated with the current API token. No path
parameter — the token determines the account.

**Response 200**:

```json
{ "id": 1, "name": "Main Shop", "currency": "CAD" }
```
