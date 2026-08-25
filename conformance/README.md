# Conformance Suite

Shared behavioral tests that run against both the Go and Ruby SDKs,
ensuring identical behavior for retry, pagination, error mapping, and
auth.

## Running

```bash
./conformance/run.sh
```

This starts the mock server, runs the Go conformance tests, then the
Ruby conformance tests, and cleans up.

## Test cases

Test cases are JSON files in `conformance/tests/`. Each file defines
one or more scenarios with:

- `operation` — the SDK method to call (e.g. `list_customers`)
- `method` / `path` — the HTTP request shape
- `mockResponses` — array of mock server responses (for retry scenarios)
- `expect` — assertions (noError, errorCode, errorStatus, responseBody, requestCount)

See `conformance/schema.json` for the full JSON schema.

## Adding a test case

1. Create a JSON file in `conformance/tests/`
2. Define the scenario following the schema
3. Add the operation to both runners if it doesn't exist:
   - `conformance/go/conformance_test.go` — `executeOperation` switch
   - `conformance/ruby/conformance_spec.rb` — `execute_operation` case
4. Run `./conformance/run.sh` to verify
