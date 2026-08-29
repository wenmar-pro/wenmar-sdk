The Wenmar Pro API
=================

Connecting to Wenmar Pro? Here's how:

- **AI agents** — Install the Wenmar Skills or the [Wenmar CLI](https://github.com/Wenmar-Pro/wenmar-cli).
- **Integrating with your app or service** — The [Wenmar SDK](https://github.com/Wenmar-Pro/wenmar-sdk) provides full-featured API clients and an [OpenAPI spec](../../spec/openapi.yaml) for code generation.
- **Going fully custom** — The OpenAPI spec and this REST API reference cover authentication, pagination, and every endpoint in detail.


Making a request
----------------

All URLs start with **`https://app.wenmarpro.com/`**. URLs are HTTPS only. There is no `/api/v1` API prefix — endpoints live at root paths (`/customers`, `/vehicles/{id}`, `/work_orders`).

To make a request for all the customers on your account, append the `customers` index path to the base URL to form something like `https://app.wenmarpro.com/customers`. In cURL, it looks like this:

```shell
curl -H "Authorization: Bearer $ACCESS_TOKEN" -A 'MyApp (yourname@example.com)' https://app.wenmarpro.com/customers
```

To create something, it's the same idea, but you also have to include the `Content-Type` header and the JSON data:

```shell
curl -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H 'Content-Type: application/json' \
  -A 'User-Agent: MyApp (yourname@example.com)' \
  -d '{ "customer": { "first_name": "Jane", "last_name": "Doe" } }' \
  https://app.wenmarpro.com/customers
```

Throughout the Wenmar API docs, we include "Copy as cURL" examples. To try the examples in your shell, copy your access token into your clipboard and run:

```shell
export ACCESS_TOKEN=PASTE_ACCESS_TOKEN_HERE
```

Then you should be able to copy/paste any example from the docs. After pasting a cURL example, you can pipe it to a JSON pretty printer to make it more readable. Try [jsonpp](https://jmhodges.github.io/jsonpp/) or `json_pp` on OSX:

```shell
curl -s -H "Authorization: Bearer $ACCESS_TOKEN" https://app.wenmarpro.com/customers | json_pp
```


Authentication
--------------

The Wenmar Pro API authenticates with **bearer tokens**. Every request must carry a token in the `Authorization` header:

```
Authorization: Bearer YOUR_API_TOKEN
```

Tokens are issued per account and digest-stored (SHA-256); the plaintext is shown once at generation time. All endpoints are private and require a valid token.

> **OAuth 2.0 is planned but not yet available.** The OpenAPI spec documents an OAuth2 + PKCE flow as a stubbed `securityScheme`, but it is not implemented. Today, only static bearer tokens work.

Read the [authentication guide](authentication.md) to get started.


Identifying your application
----------------------------

You must include a `User-Agent` header with **both**:

* The name of your application
* A link to your application or your email address

We use this information to get in touch if you're doing something wrong (so we can warn you) or something awesome (so we can congratulate you). Here are examples of acceptable `User-Agent` headers:

* `User-Agent: WenmarSync (https://example.com/contact)`
* `User-Agent: Fabian's Shop Integration (fabian@example.com)`

If you don't include a `User-Agent` header, you'll get a `400 Bad Request` response.


JSON only
---------

We use JSON for all API data. The style is **no root element and no envelope** — responses are bare objects or arrays, and object keys are `snake_case`. This means you must send the `Content-Type: application/json; charset=utf-8` header when POSTing or PATCHing data into Wenmar Pro.

- Single resource: `GET /customers/1` returns `{ "id": 1, "full_name": "Jane Doe", ... }`
- Collection: `GET /customers` returns `[ { "id": 1, ... }, { "id": 2, ... } ]`

There is no `{ "data": ... }` wrapper. The response body is the resource directly. The `.json` format suffix is supported (e.g. `/customers.json`) but not required.

You'll receive a `415 Unsupported Media Type` response code if you don't include the `Content-Type` header.


Pagination
----------

Most collection APIs paginate their results. We use the [RFC 5988](https://www.rfc-editor.org/rfc/rfc5988) convention of using the `Link` header to provide the URL for the `next` page. Follow this convention to retrieve the next page of data—please don't build the pagination URLs yourself!

Here's an example response header from requesting the second page of [customers](sections/customers.md):

```
Link: <https://app.wenmarpro.com/customers?page=3>; rel="next"
```

If the `Link` header is absent (or carries no `rel="next"`), that's the last page. Page numbers are 1-based; page 1 is the default. The `next` URL includes any query params and must be followed verbatim.

See [pagination](pagination.md) for SDK-specific helpers.


Using HTTP caching
------------------

You must use HTTP freshness headers to speed up your application and lighten the load on our servers. Most API responses will include an `ETag` or `Last-Modified` header. When you first request a resource, store these values. On subsequent requests, submit them back to us as `If-None-Match` and `If-Modified-Since`, respectively. If the resource hasn't changed since your last request, you'll get a `304 Not Modified` response with no body, saving you the time and bandwidth of sending something you already have.


Handling errors
---------------

Every error response uses one consistent envelope:

```json
{ "error": { "code": "not_found", "message": "Customer not found", "details": {} } }
```

| HTTP status | `code` | When |
|---|---|---|
| 401 | `unauthorized` | Missing or invalid bearer token |
| 403 | `forbidden` | Authorization failure |
| 404 | `not_found` | Resource not found |
| 422 | `validation_failed` | Request body failed validation — `details` keyed by field |
| 429 | `rate_limited` | Too many requests — retry after delay |
| 500+ | `internal_error` | Server-side failure |

API clients must expect and gracefully handle transient server errors and rate limits. We recommend baking graceful 5xx and 429 retries into your integration from the beginning so errors are handled automatically.

### Rate limiting (429 Too Many Requests)

We return a [429 Too Many Requests](http://tools.ietf.org/html/draft-nottingham-http-new-status-02) response when you've exceeded a rate limit. Consult the `Retry-After` response header to determine how long to wait (in seconds) before retrying the request. The SDKs retry 429 responses with exponential backoff and respect `Retry-After`.

### 5xx server errors

If Wenmar Pro is having trouble, you'll get a response with a 5xx status code indicating a server error. 500, 502, 503, and 504 may be retried with [exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff). The SDKs retry 5xx responses up to 3 times.

### 404 Not Found

API requests may 404 due to deleted content, an inactive account, or insufficient permissions. Detect these conditions to give your users a clear explanation. Do not automatically retry these requests.

See [errors](errors.md) for SDK-specific error handling.


Key concepts
------------

Understanding Wenmar Pro's domain model helps you navigate the API effectively.

### No URL prefixes

All public endpoints use bare paths — no `/api/` prefix, no `/v1/` version segment:

```
GET /customers
POST /work_orders
GET /vehicles/{id}
```

Changes are additive-only: new fields, new endpoints, new resources. We never rename or repurpose an existing field's meaning.

### Account and locations

An **account** is the top-level tenant — the shop. The token determines the account, so `/account` takes no path parameter. An account has one or more **locations** (bays, sites).

### Customers, vehicles, work orders

The core hierarchy:

```
Account
  └── Location (optional scope)
        └── Customer
              ├── Vehicle (one or more)
              │     └── Work Order
              └── Work Order (direct, e.g. walk-in)
```

- A **customer** owns vehicles and work orders. Nested objects like `emails` and `phones` are inline arrays.
- A **vehicle** carries VIN, odometer, and plate data. Use `/vehicles/vin_decode` to decode a VIN before creating, and `/vehicles/check_duplicate` to detect existing records.
- A **work order** is the central operational record — services, payments, messages, and activity logs are reached via `_url` fields on the work order (`services_url`, `payments_url`, `messages_url`, `activity_url`).

### Linked resources via `_url` fields

Resources include `_url` fields pointing to related collections rather than embedding them. A customer exposes `vehicles_url` and `work_orders_url`; a work order exposes `services_url`, `payments_url`, `messages_url`, and `activity_url`. Follow these URLs rather than constructing paths yourself.

### Bare objects, no envelope

Responses are the resource directly — no `{ "data": ... }` wrapper. A `type` field is present for convenience but is not required to route the response. This is a deliberate decision for a single-tenant-per-token API.

### Money fields

All monetary values are expressed in cents as integers with a `_cents` suffix (`total_cents`, `outstanding_balance_cents`). The currency code (`CAD`, `USD`) is carried alongside in a `currency` field on objects that include totals.

### Timestamps

All timestamps are ISO 8601 with timezone offset (e.g. `2026-08-27T12:00:00.000-04:00`).


API endpoints
-------------

- [Authentication](authentication.md#authentication)
<!-- START API ENDPOINTS -->
- [Account](sections/account.md#account)
- [Customers](sections/customers.md#customers)
- [Locations](sections/locations.md#locations)
- [Vehicles](sections/vehicles.md#vehicles)
- [Work Orders](sections/work_orders.md#work-orders)

<!-- END API ENDPOINTS -->

See the [full API reference](api-reference.md) for a single-table view of every endpoint.


How the spec flows
------------------

```
wenmar-pro (Rails app, source of truth)
   │  generates the spec from request tests, redacts + publishes on merge to main
   ▼
wenmar-sdk (this repo)
   │  spec/openapi.yaml is the canonical public spec; docs/ describe it for humans and agents
   │  enrichment + Go/Ruby codegen + conformance
   ▼
wenmar-cli
   │  consumes the Go SDK
```

The spec is generated from request tests in the [wenmar-pro](https://github.com/Wenmar-Pro/wenmar-pro) Rails app and pushed here by its `spec-check.yml` workflow on every merge to `main`.


Getting Help
------------

If you have a question about the API, please [open an issue](https://github.com/Wenmar-Pro/wenmar-sdk/issues).


License
-------

These API docs are licensed under [Creative Commons (CC BY-SA 4.0)](LICENSE). The SDK code in this repo is licensed separately under the MIT license (see the root [LICENSE](../../LICENSE)).