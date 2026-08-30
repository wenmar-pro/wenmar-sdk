package wenmar

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSlogHooks_LogsOperations(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))
	hooks := NewSlogHooks(logger)
	cfg := DefaultConfig()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()
	cfg.BaseURL = ts.URL

	c, _ := NewClient(cfg, NewStaticTokenProvider("test-token"), WithHooks(hooks))
	_, _ = c.ListCustomers(context.Background(), nil)

	output := buf.String()
	if !strings.Contains(output, "operation=ListCustomers") {
		t.Errorf("expected log to contain operation=ListCustomers, got: %s", output)
	}
}
