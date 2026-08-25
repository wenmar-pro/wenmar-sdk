# Authentication

Every request to the Wenmar Pro API must carry a bearer token in the
`Authorization` header:

```
Authorization: Bearer YOUR_API_TOKEN
```

All endpoints are private and require a valid token. Missing or invalid tokens
return a `401` response:

```json
{ "error": { "code": "unauthorized", "message": "Invalid or missing API token", "details": {} } }
```

## Getting a token

An API token is issued per account. Contact your Wenmar Pro administrator to
provision one for your shop.

## Per-SDK usage

**Go**

```go
client, err := wenmar.NewClient("https://app.wenmarpro.com", "YOUR_API_TOKEN")
```

**Ruby**

```ruby
client = Wenmar::Client.new(token: "YOUR_API_TOKEN")
```

## Security notes

- Never commit tokens to source control.
- Store tokens in environment variables or a secret manager.
- The token grants access to a single account's data — treat it like a
  password.
