.PHONY: check test conformance enrich generate docs scalar check-conformance-parity check-ruby-signature-parity

generate: enrich ## Run the full codegen pipeline from the enriched spec
	cd go && go generate ./...
	ruby scripts/generate_manifest.rb spec/openapi.enriched.yaml spec/operations.json
	ruby scripts/generate_go_wrapper.rb
	ruby scripts/generate_ruby_resources.rb
	ruby scripts/generate_conformance_dispatch.rb

enrich: ## Enrich the spec from the canonical openapi.yaml
	ruby scripts/enrich_spec.rb spec/openapi.yaml spec/openapi.enriched.yaml

docs: enrich ## Generate markdown API docs from the enriched spec + fixtures
	ruby scripts/generate_docs.rb spec/openapi.enriched.yaml docs/api spec/fixtures

scalar: ## Copy enriched spec into docs/scalar for local preview
	cp spec/openapi.enriched.yaml docs/scalar/openapi.enriched.yaml
	@echo "Open docs/scalar/index.html in a browser"

test: ## Run Go and Ruby unit tests
	cd go && go test ./...
	cd ruby && bundle exec ruby -Ilib spec/client_spec.rb
	cd ruby && bundle exec ruby -Ilib spec/pagination_spec.rb
	cd ruby && bundle exec ruby -Ilib spec/error_spec.rb
	cd ruby && bundle exec ruby -Ilib spec/config_spec.rb

conformance: ## Run both conformance suites
	cd conformance/go && go test ./...
	cd conformance/ruby && bundle exec ruby conformance_spec.rb

check-conformance-parity: ## Verify manifest <-> dispatch <-> test parity
	ruby scripts/check_conformance_parity.rb

check-ruby-signature-parity: ## Verify Ruby dispatch kwargs are accepted by the client
	ruby scripts/check_ruby_signature_parity.rb

check: enrich generate docs test conformance check-conformance-parity check-ruby-signature-parity ## Full CI check — fails on drift
	@echo "Checking for generated-file drift..."
	@if ! git diff --exit-code spec/openapi.enriched.yaml; then \
		echo "::error::openapi.enriched.yaml has drifted. Run 'make enrich' and commit."; \
		exit 1; \
	fi
	@if ! git diff --exit-code spec/operations.json; then \
		echo "::error::spec/operations.json has drifted. Run 'make generate' and commit."; \
		exit 1; \
	fi
	@if ! git diff --exit-code go/pkg/generated/client.gen.go; then \
		echo "::error::generated Go client has drifted. Run 'go generate ./...' and commit."; \
		exit 1; \
	fi
	@if ! git diff --exit-code go/wenmar/operations.gen.go go/wenmar/models.gen.go; then \
		echo "::error::generated Go wrapper has drifted. Run 'make generate' and commit."; \
		exit 1; \
	fi
	@if ! git diff --exit-code ruby/lib/wenmar/resources.rb; then \
		echo "::error::generated Ruby resources have drifted. Run 'make generate' and commit."; \
		exit 1; \
	fi
	@if ! git diff --exit-code conformance/go/dispatch.gen.go conformance/ruby/dispatch.gen.rb; then \
		echo "::error::generated conformance dispatch has drifted. Run 'make generate' and commit."; \
		exit 1; \
	fi
	@if ! git diff --exit-code docs/api/sections/ docs/api/api-reference.md docs/api/llm-compact.md; then \
		echo "::error::generated API docs have drifted. Run 'make docs' and commit."; \
		exit 1; \
	fi
	ruby scripts/check_fixture_coverage.rb
	@echo "All checks passed."
