package wenmar

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestNewPrometheusHooks_DuplicateRegistrationDoesNotPanic(t *testing.T) {
	reg := prometheus.NewRegistry()
	if _, err := NewPrometheusHooks(reg, nil); err != nil {
		t.Fatalf("first registration failed: %v", err)
	}
	if _, err := NewPrometheusHooks(reg, nil); err != nil {
		t.Fatalf("duplicate registration failed: %v", err)
	}
}

func TestPrometheusHooks_CountsOperations(t *testing.T) {
	reg := prometheus.NewRegistry()
	hooks, err := NewPrometheusHooks(reg, reg)
	if err != nil {
		t.Fatalf("NewPrometheusHooks: %v", err)
	}
	cfg := DefaultConfig()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()
	cfg.BaseURL = ts.URL

	c, _ := NewClient(cfg, NewStaticTokenProvider("test-token"), WithHooks(hooks))
	_, _ = c.ListCustomers(context.Background(), nil)

	mfs, err := reg.Gather()
	if err != nil {
		t.Fatalf("gather failed: %v", err)
	}
	var found bool
	for _, mf := range mfs {
		if mf.GetName() == "wenmar_operations_total" {
			found = true
			for _, m := range mf.GetMetric() {
				if m.GetCounter().GetValue() < 1 {
					t.Error("expected counter >= 1")
				}
			}
		}
	}
	if !found {
		t.Error("wenmar_operations_total metric not found")
	}
}
