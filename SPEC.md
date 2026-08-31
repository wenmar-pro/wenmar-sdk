# Wenmar SDK Behavioral Contract

This document is the cross-SDK behavioral contract for the Wenmar Pro API
clients (Go and Ruby). It describes the guarantees both SDKs provide so that
behavior is consistent regardless of language. The canonical OpenAPI spec
lives in [`spec/openapi.yaml`](spec/openapi.yaml) in this repo.

## Retry Policy

| Status | Retried? | Methods | Notes |
|--------|----------|---------|-------|
| 429 Too Many Requests | Yes | All | Idempotent throttle; honors `Retry-After` |
| 500/502/503/504 | Yes | GET only | Exponential backoff + jitter; max 3 |
| 507 Insufficient Storage | No | — | Account/plan limit; not transient |
| 404 Not Found | No | — | Deleted, inaccessible, or forbidden |
| 304 Not Modified | No | — | Not an error; returns cached body |

Mutations (POST/PATCH/DELETE) are **not** retried on 5xx because the server
may have processed the request before the response was lost, and retrying
would duplicate the side effect. 429 is safe to retry because the throttle
response means the request was **not** processed.

Retry backoff is exponential with jitter: `delay = base * 2^attempt + jitter`,
base 500ms, max 3 retries. A `Retry-After` header (seconds or HTTP-date)
overrides the computed delay.

## Pagination

- Pagination uses RFC 5988 `Link` headers only (no `{ "data": ... }` envelope).
- The `next` link is followed for subsequent pages.
- **Same-origin validation:** the SDK refuses to follow a `next` URL that is
  not same-origin as the configured base URL, preventing token exfiltration
  via a malicious `Link` header. This returns an `invalid_pagination` error.
- `X-Total-Count` and `X-Per-Page` headers are surfaced as pagination metadata
  when present.
- `GetAll` auto-pagination supports `max_items` / `max_pages` caps. The default
  safety cap is 1000 items. When a cap is hit, the result is marked truncated.

## Error Codes

The API returns errors in the `{ "error": { code, message, field_errors } }`
envelope. The SDK maps each code to an error type with the following
properties:

| Code | HTTP | Retryable | Notes |
|------|------|-----------|-------|
| `unauthorized` | 401 | No | Invalid or missing API token |
| `forbidden` | 403 | No | Authenticated but not permitted |
| `not_found` | 404 | No | Resource missing or inaccessible |
| `rate_limited` | 429 | Yes | Throttled; honor `Retry-After` |
| `validation_failed` | 422 | No | Field-level errors in `field_errors` |
| `internal_error` | 500 | Yes (GET) | Server error |
| `limit_exceeded` | 507 | No | Account/plan limit reached |

`field_errors` may contain field-level validation errors keyed by field name
(e.g. `{ "first_name": ["can't be blank"] }`). The SDK exposes these via a
`FieldErrors` / `field_errors` accessor.

## Authentication

- Authentication is **bearer token only**: `Authorization: Bearer <token>`.
- Tokens are **location-pinned**: a token is bound to a single location. A
  caller may scope a request to a specific location by sending the
  `X-Wenmar-Location` header. The `for_location` / `ForLocation` sub-client
  injects this header on every request. The server rejects a location the
  token is not permitted to access with `403 forbidden`.
- `X-Request-Id` is captured from error responses and surfaced on the error
  object for support correlation.

## Conditional GET

- The SDK sends `If-None-Match` (ETag) and `If-Modified-Since` headers on
  repeat GETs.
- A `304 Not Modified` response returns the previously cached body.

## Headers

| Header | Value |
|--------|-------|
| `User-Agent` | `wenmar-sdk-{lang}/{version}` (e.g. `wenmar-sdk-go/0.2.0`) |
| `Accept` | `application/json` |
| `Authorization` | `Bearer <token>` |

## HTTPS Enforcement

- Non-https base URLs are rejected unless the host is `localhost` or
  `127.0.0.1` (for development).
- On cross-origin redirects, the `Authorization` header is stripped to prevent
  credential leakage to third parties.

## Observability

The Go SDK exposes a `Hooks` interface for logging, tracing, and metrics:

- `SlogHooks` — structured debug logging via `log/slog`.
- `OTelHook` — OpenTelemetry spans per SDK operation.
- `PrometheusHook` — operation/request/retry counters.

Both operation-level and request-level hooks propagate a `context.Context`:

```go
type Hooks interface {
    OnOperationStart(ctx context.Context, op OperationInfo) context.Context
    OnOperationEnd(ctx context.Context, op OperationInfo, err error)
    OnRequestStart(ctx context.Context, req RequestInfo) context.Context
    OnRequestEnd(ctx context.Context, req RequestInfo, resp *ResponseInfo, err error)
    OnRetry(ctx context.Context, retry RetryInfo)
    OnPaginate(ctx context.Context, paginate PaginateInfo)
}
```

`OnOperationStart` and `OnRequestStart` return a (possibly derived) context
that is threaded through to the matching `*End` callback. This enables
request-level OTel child spans and hook chaining (e.g. OTel + Prometheus
composed).
