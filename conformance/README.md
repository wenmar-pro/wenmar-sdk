# Conformance Suite

Shared behavioral tests that run against both the Go and Ruby SDKs,
ensuring identical behavior for retry, pagination, error mapping, and
auth.

## Running

```bash
make conformance
```

This runs the Go conformance tests (`conformance/go`) and the Ruby
conformance tests (`conformance/ruby`) against the shared scenarios in
`conformance/tests/`.

## Test cases

Test cases are JSON files in `conformance/tests/`. Each file defines
one or more scenarios with:

- `operation` — the SDK method to call (e.g. `list_customers`)
- `method` / `path` — the HTTP request shape
- `mockResponses` — array of mock server responses (for retry scenarios)
- `expect` — assertions applied by the runners. Supported keys:
  - `noError` — expect the call to succeed
  - `errorCode` / `errorStatus` — the API error code / status when a call fails
  - `responseBody` — a path/`equals` assertion against the response body
  - `requestCount` — number of HTTP requests the SDK must make
  - `requestHeaders` — exact request-header values to assert
  - `requestHeadersPresent` — request headers that must be present (value may be environment-specific)
  - `repeat` — call the operation N times to exercise stateful behavior (e.g. conditional-GET caching)

## Adding a test case

1. Create a JSON file in `conformance/tests/`
2. Define the scenario following the schema
3. Regenerate the dispatch tables so the new operation is callable by the
   runners:

   ```
   ruby scripts/generate_conformance_dispatch.rb
   ```

   This reads `spec/operations.json` and emits both
   `conformance/go/dispatch.gen.go` and `conformance/ruby/dispatch.gen.rb`
   (both marked `DO NOT EDIT`). If the operation is already in the
   manifest, no manual edit to the runners is needed.
4. Run `make conformance` to verify
