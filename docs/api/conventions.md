# Wenmar Pro API Conventions

This document defines the conventions for the public API. Every endpoint
follows these rules without exception.

## Response format

Responses are **bare objects or arrays** — no envelope wrapper.

- Single resource: `GET /customers/1` returns `{ "id": 1, "full_name": "Jane Doe", ... }`
- Collection: `GET /customers` returns `[ { "id": 1, ... }, { "id": 2, ... } ]`

There is no `{ "data": ... }` wrapper. The response body is the resource
directly. This is a deliberate decision — the envelope adds a layer of
indirection without value for a single-tenant-per-token API.

## Pagination

Pagination is delivered via the RFC 5988 `Link` response header:

```
Link: <https://app.wenmarpro.com/customers?page=2>; rel="next"
```

- `rel="next"` — next page
- `rel="prev"` — previous page

No body-side cursor, offset, or page field is exposed to clients. The
header is the pagination contract. Page numbers are 1-based; page 1 is
the default. When there is no next page, the `Link` header is absent.

## Errors

All errors use one envelope:

```json
{ "error": { "code": "not_found", "message": "Customer not found", "details": {} } }
```

| `code` | HTTP status | When |
|---|---|---|
| `unauthorized` | 401 | Missing/invalid/expired bearer token |
| `forbidden` | 403 | Authorization failure |
| `not_found` | 404 | Record not found |
| `validation_failed` | 422 | Request body failed validation — `details` keyed by field |

The `details` object carries per-field validation errors, e.g.
`{ "full_name": ["can't be blank"] }`.

## Auth

The API supports **static bearer tokens**. Tokens are digest-stored
(SHA-256); the plaintext is shown once at generation time. Authentication
is exclusively via the `Authorization: Bearer <token>` header.

OAuth2 + PKCE is documented in the spec's `securitySchemes` as a stubbed
flow. It will be implemented in a later phase.

## Versioning

No `/v1/`, `/v2/` in the URL. No `/api/` prefix either — endpoints live at
root paths (`/customers`, `/vehicles/{id}`, `/work_orders`).

- Changes are additive-only: new fields, new endpoints, new resources.
- Never rename or repurpose an existing field's meaning.
- A field/response-shape bug is fixed as a defect, not a version bump.
- A genuine breaking redesign is a new namespace/product, not a version.

## Spec generation

The OpenAPI spec is generated from request tests in the wenmar-pro Rails
app using `rspec-openapi` (pinned to OpenAPI 3.0 for committee
compatibility). The `spec-check.yml` workflow in wenmar-pro:

1. Regenerates the spec from request tests (`OPENAPI=1 bin/rails test`).
2. Validates responses against the spec using `committee`.
3. Checks for drift between the generated and committed spec.
4. On merge to main, redacts internal endpoints and publishes to this repo.

## Endpoint paths

All public endpoints use bare paths — no `/api/` prefix:

- `GET /customers` — list customers
- `POST /customers` — create customer
- `GET /customers/{id}` — show customer
- `PATCH /customers/{id}` — update customer

The `.json` format suffix is supported (e.g. `/customers.json`) but not
required. The spec documents paths without the suffix.
