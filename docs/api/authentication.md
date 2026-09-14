# Authentication

Every request to the Wenmar Pro API must carry a bearer token in the
`Authorization` header:

```
Authorization: Bearer YOUR_API_TOKEN
```

All endpoints are private and require a valid token. Missing or invalid tokens
return a `401` response:

```json
{ "error": { "code": "unauthorized", "message": "Invalid or missing API token", "field_errors": {} } }
```

## Token types

The API accepts two kinds of bearer token:

- **API tokens (static PATs)** — long-lived tokens minted per shop account and
  scoped to a location. Create one from the web UI under *Account → API tokens*
  or via the JSON API. This is the simplest integration path.
- **OAuth access tokens** — issued by the browser `authorization_code` + PKCE
  flow against the public `wenmar-cli` OAuth application. Access tokens expire
  after 2 hours and can be refreshed with the `refresh_token` grant.

## Location scoping

Tokens are **location-pinned**: a token belongs to exactly one location. Sending
the `X-Wenmar-Location` header scopes the request to a location the token's user
may access:

```
X-Wenmar-Location: 42
```

- For a location-scoped API token the header is **required**; omitting it returns
  `400 bad_request`, and sending a location that doesn't match the token returns
  `403 forbidden`.
- For OAuth tokens (which are not pinned to a location) the header is optional;
  without it the request uses the user's home location.

The Go SDK's `ForLocation` and the Ruby SDK's `for_location` build a
location-scoped sub-client that injects this header on every request.

## Per-SDK usage

**Go**

```go
cfg := wenmar.DefaultConfig()
cfg.BaseURL = "https://app.wenmarpro.com"
client, err := wenmar.NewClient(cfg, wenmar.NewStaticTokenProvider("YOUR_API_TOKEN"))
scoped := client.ForLocation("42")
```

**Ruby**

```ruby
client = Wenmar::Client.new(token: "YOUR_API_TOKEN")
scoped = client.for_location("42")
```

## OAuth refresh

Both SDKs can exchange a stored refresh token for a new access token:

**Go**

```go
newTok, err := auth.RefreshToken(ctx, baseURL+"/oauth/token", "wenmar-cli", refreshToken)
```

`auth.NewAuthManagerWithOAuth(store, provider, baseURL, "wenmar-cli")` wires the
refresh into an `AuthManager` whose `CredentialStoreProvider` auto-refreshes
tokens that are expired or within five minutes of expiry.

**Ruby**

```ruby
token = Wenmar::OAuth.refresh(base_url: "https://app.wenmarpro.com", refresh_token: rt)
store = Wenmar::CredentialStore.new          # JSON file, 0600
store.save_token(token)
manager = Wenmar::AuthManager.new(store: store, oauth: { base_url: "https://app.wenmarpro.com" })
manager.refresh
```

`Wenmar::CredentialStoreProvider.new(store: store, manager: manager)` resolves a
usable access token, refreshing automatically when needed.

## Security notes

- Never commit tokens to source control.
- Store tokens in environment variables, a secret manager, or the SDK's
  credential store (`~/.config/wenmar/credentials.json`, written `0600`).
- The token grants access to a single account/location's data — treat it like a
  password.
