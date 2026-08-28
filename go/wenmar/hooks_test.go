package wenmar

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestHooks_OnOperationStartEnd(t *testing.T) {
	recorder := &recordingHooks{}
	cfg := DefaultConfig()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()
	cfg.BaseURL = ts.URL

	c, _ := NewClient(cfg, NewStaticTokenProvider("test-token"), WithHooks(recorder))
	_, _ = c.ListCustomers(context.Background())

	if len(recorder.starts) != 1 {
		t.Errorf("expected 1 OnOperationStart, got %d", len(recorder.starts))
	}
	if len(recorder.ends) != 1 {
		t.Errorf("expected 1 OnOperationEnd, got %d", len(recorder.ends))
	}
	if recorder.ends[0].Operation != "ListCustomers" {
		t.Errorf("expected operation 'ListCustomers', got %q", recorder.ends[0].Operation)
	}
}

type recordingHooks struct {
	mu     sync.Mutex
	starts []OperationInfo
	ends   []OperationResult
}

func (h *recordingHooks) OnOperationStart(ctx context.Context, info OperationInfo) context.Context {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.starts = append(h.starts, info)
	return ctx
}
func (h *recordingHooks) OnOperationEnd(_ context.Context, info OperationInfo, result OperationResult) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.ends = append(h.ends, result)
}
func (h *recordingHooks) OnRequestStart(context.Context, RequestInfo)    {}
func (h *recordingHooks) OnRequestEnd(context.Context, RequestInfo, RequestResult) {}
func (h *recordingHooks) OnRetry(context.Context, RequestInfo, int, error)        {}
func (h *recordingHooks) OnPaginate(context.Context, string, int)                  {}
