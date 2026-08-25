#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "=== Starting mock server ==="
go run conformance/mock/server.go -port 18080 &
MOCK_PID=$!
trap 'kill $MOCK_PID 2>/dev/null || true' EXIT
sleep 2

echo ""
echo "=== Go conformance ==="
(cd conformance/go && go test -v -timeout 30s)

echo ""
echo "=== Ruby conformance ==="
(cd conformance/ruby && bundle install --quiet && bundle exec ruby conformance_spec.rb)

echo ""
echo "=== All conformance tests passed ==="
