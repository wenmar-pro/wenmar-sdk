# Wenmar SDK

Go and Ruby SDKs for the Wenmar Pro automotive shop management API, plus
the OpenAPI spec and conformance suite.

## Architecture

The public API surface spans four repos. The Rails app is the source of
truth; the API docs repo holds the spec; this repo and the CLI are derived.

- `wenmar-pro` (`~/Projects/wenmar-pro`) — Rails app, **source of truth** for the API
- `wenmar-api` (`~/Projects/wenmar-api`) — OpenAPI spec + human-readable docs (CC BY-SA)
- `wenmar-sdk` (this repo) — Go/Ruby SDKs + conformance suite (pulls spec from wenmar-api)
- `wenmar-cli` (`~/Projects/wenmar-cli`) — Go CLI (Cobra) consuming the Go SDK

```
wenmar-pro (Rails, source of truth)
   │  exports OpenAPI spec
   ▼
wenmar-api (OpenAPI spec + docs, the contract)
   │  spec consumed by
   ▼
wenmar-sdk (Go/Ruby SDKs + conformance)
   │  Go SDK consumed by
   ▼
wenmar-cli (Go CLI, Cobra)
```

## Key decisions

- **No API versioning** — the API is additive-only. No `/v1/` prefix. New
  fields may appear; existing fields keep their meaning. Do not introduce
  versioned URL prefixes.
- **No Smithy** — OpenAPI is the spec format. Revisit Smithy only when
  ~50+ operations, 3+ SDK languages, or conformance schema generation
  from the spec is needed.
- **Spec is read-only here** — `spec/openapi.yaml` is synced from
  wenmar-pro. Do not hand-edit it; sync it out from the Rails app.

## Quick reference

- Regenerate Go client: `cd go && go generate ./...`
- Enrich spec: `ruby scripts/enrich_spec.rb spec/openapi.yaml spec/openapi.enriched.yaml`
- Fixture coverage: `ruby scripts/check_fixture_coverage.rb`
- Go tests: `cd go && go test ./...`
- Ruby tests: `cd ruby && bundle install && ruby -Ilib spec/client_spec.rb`
- Conformance: `cd conformance/go && go test ./...` and `cd conformance/ruby && bundle exec ruby conformance_spec.rb`

## Full architecture

For the full architecture, codegen pipeline, and the checklist for
touching `Api::` controllers, load the `sdk-cli` skill from
`wenmar-pro/.opencode/skills/sdk-cli/SKILL.md`.
