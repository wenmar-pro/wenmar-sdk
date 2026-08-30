package wenmar

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestOTelHooks_CreatesSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	defer tp.Shutdown(context.Background())

	hooks := NewOTelHook(tp.Tracer("wenmar"))
	cfg := DefaultConfig()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()
	cfg.BaseURL = ts.URL

	ctx, span := tp.Tracer("test").Start(context.Background(), "parent")

	c, _ := NewClient(cfg, NewStaticTokenProvider("test-token"), WithHooks(hooks))
	_, _ = c.ListCustomers(ctx, nil)
	span.End()

	spans := exporter.GetSpans()
	if len(spans) < 2 {
		t.Fatalf("expected at least 2 spans (parent + child), got %d", len(spans))
	}
	// Find the operation span
	var opSpan *tracetest.SpanStub
	for _, s := range spans {
		if s.Name == "ListCustomers" {
			opSpan = &s
			break
		}
	}
	if opSpan == nil {
		t.Fatal("expected a span named 'ListCustomers'")
	}
}
