#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "Enriching spec..."
ruby scripts/enrich_spec.rb spec/openapi.yaml spec/openapi.enriched.yaml

echo "Building Scalar docs..."
PORT="${SCALAR_PORT:-18080}"
npx --yes @scalar/cli document serve spec/openapi.enriched.yaml --port "$PORT" \
  > /tmp/scalar-build.log 2>&1 &
SERVE_PID=$!

cleanup() {
  kill "$SERVE_PID" 2>/dev/null || true
}
trap cleanup EXIT

# Wait for the server to come up (bounded loop).
for i in $(seq 1 30); do
  if curl -fsSL "http://localhost:$PORT/" -o docs/index.html 2>/dev/null; then
    echo "Docs built to docs/index.html"
    exit 0
  fi
  sleep 1
done

echo "Scalar server did not become ready on port $PORT" >&2
cat /tmp/scalar-build.log >&2
exit 1
