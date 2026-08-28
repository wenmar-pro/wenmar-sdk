.PHONY: check test conformance enrich generate

generate: ## Regenerate the Go client from the enriched spec
	cd go && go generate ./...

enrich: ## Enrich the spec from the canonical openapi.yaml
	ruby scripts/enrich_spec.rb spec/openapi.yaml spec/openapi.enriched.yaml

test: ## Run Go and Ruby unit tests
	cd go && go test ./...
	cd ruby && bundle exec ruby -Ilib spec/client_spec.rb
	cd ruby && bundle exec ruby -Ilib spec/pagination_spec.rb
	cd ruby && bundle exec ruby -Ilib spec/error_spec.rb
	cd ruby && bundle exec ruby -Ilib spec/config_spec.rb

conformance: ## Run both conformance suites
	cd conformance/go && go test ./...
	cd conformance/ruby && bundle exec ruby conformance_spec.rb

check: enrich generate test conformance ## Full CI check — fails on drift
	@echo "Checking for spec drift..."
	@if ! git diff --exit-code spec/openapi.enriched.yaml; then \
		echo "::error::openapi.enriched.yaml has drifted. Run 'make enrich' and commit."; \
		exit 1; \
	fi
	@if ! git diff --exit-code go/pkg/generated/client.gen.go; then \
		echo "::error::generated Go client has drifted. Run 'go generate ./...' and commit."; \
		exit 1; \
	fi
	ruby scripts/check_fixture_coverage.rb
	@echo "All checks passed."
