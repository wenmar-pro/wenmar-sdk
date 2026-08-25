# Wenmar SDK

Public SDK and OpenAPI spec for the Wenmar Pro API.

## Structure

- `spec/openapi.yaml` — synced from the private wenmar-pro repo on every merge to main. Read-only here.
- `ruby/` — (Phase 2) generated Ruby gem
- `go/` — (Phase 2) generated Go module
- `docs/` — (Phase 2) generated docs site

## Spec sync

The `spec/openapi.yaml` file is auto-published by the `spec-check.yml`
workflow in the private `wenmar-pro` repo. Do not edit it directly —
changes must come from the Rails app's request tests.
