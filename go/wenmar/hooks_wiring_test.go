package wenmar

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
)

// recordingHooksRecords is a recording Hooks implementation that captures
// request, retry, and paginate callbacks for assertions.
type recordingHooksRecords struct {
	mu        sync.Mutex
	reqStarts []RequestInfo
	reqEnds   []RequestResult
	retries   []retryRecord
	paginates []paginateRecord
}

type retryRecord struct {
	info    RequestInfo
	attempt int
	err     error
}

type paginateRecord struct {
	url  string
	page int
}

func (h *recordingHooksRecords) OnOperationStart(ctx context.Context, _ OperationInfo) context.Context {
	return ctx
}
func (h *recordingHooksRecords) OnOperationEnd(context.Context, OperationInfo, OperationResult) {}
func (h *recordingHooksRecords) OnRequestStart(_ context.Context, info RequestInfo) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.reqStarts = append(h.reqStarts, info)
}
func (h *recordingHooksRecords) OnRequestEnd(_ context.Context, _ RequestInfo, result RequestResult) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.reqEnds = append(h.reqEnds, result)
}
func (h *recordingHooksRecords) OnRetry(_ context.Context, info RequestInfo, attempt int, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.retries = append(h.retries, retryRecord{info: info, attempt: attempt, err: err})
}
func (h *recordingHooksRecords) OnPaginate(_ context.Context, url string, page int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.paginates = append(h.paginates, paginateRecord{url: url, page: page})
}

func (h *recordingHooksRecords) reqEndsCount() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return len(h.reqEnds)
}

func TestHooksWiring_RequestStartEnd(t *testing.T) {
	var reqMethod, reqPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqMethod = r.Method
		reqPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	rec := &recordingHooksRecords{}
	cfg := DefaultConfig()
	cfg.BaseURL = ts.URL
	c, err := NewClient(cfg, NewStaticTokenProvider("tok"), WithHooks(rec))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if _, err := c.ListCustomers(context.Background(), nil); err != nil {
		t.Fatalf("ListCustomers: %v", err)
	}

	rec.mu.Lock()
	defer rec.mu.Unlock()
	if len(rec.reqStarts) != 1 {
		t.Fatalf("expected 1 OnRequestStart, got %d", len(rec.reqStarts))
	}
	if len(rec.reqEnds) != 1 {
		t.Fatalf("expected 1 OnRequestEnd, got %d", len(rec.reqEnds))
	}
	if rec.reqStarts[0].Method != http.MethodGet || reqMethod != http.MethodGet {
		t.Errorf("expected GET method, got start=%q actual=%q", rec.reqStarts[0].Method, reqMethod)
	}
	if rec.reqEnds[0].StatusCode != http.StatusOK {
		t.Errorf("expected status 200 in OnRequestEnd, got %d", rec.reqEnds[0].StatusCode)
	}
	if rec.reqEnds[0].Err != nil {
		t.Errorf("expected nil Err in OnRequestEnd, got %v", rec.reqEnds[0].Err)
	}
	if want := ts.URL + reqPath; rec.reqEnds[0].URL != want {
		t.Errorf("expected URL %q, got %q", want, rec.reqEnds[0].URL)
	}
}

func TestHooksWiring_RetryOn500Then200(t *testing.T) {
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":{"code":"internal_error","message":"fail","details":{}}}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	rec := &recordingHooksRecords{}
	cfg := DefaultConfig()
	cfg.BaseURL = ts.URL
	cfg.MaxRetries = 3
	c, err := NewClient(cfg, NewStaticTokenProvider("tok"), WithHooks(rec))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if _, err := c.ListCustomers(context.Background(), nil); err != nil {
		t.Fatalf("expected success after retry, got: %v", err)
	}

	rec.mu.Lock()
	defer rec.mu.Unlock()
	if len(rec.retries) < 1 {
		t.Fatalf("expected at least 1 OnRetry, got %d", len(rec.retries))
	}
	if rec.retries[0].info.Method != http.MethodGet {
		t.Errorf("expected GET method in OnRetry, got %q", rec.retries[0].info.Method)
	}
	if rec.retries[0].attempt != 1 {
		t.Errorf("expected first retry attempt=1, got %d", rec.retries[0].attempt)
	}
	if rec.retries[0].err == nil {
		t.Error("expected non-nil err in OnRetry")
	}
	// OnRequestEnd should report the final successful status.
	if len(rec.reqEnds) != 1 {
		t.Fatalf("expected 1 OnRequestEnd (outermost), got %d", len(rec.reqEnds))
	}
	if rec.reqEnds[0].StatusCode != http.StatusOK {
		t.Errorf("expected final status 200 in OnRequestEnd, got %d", rec.reqEnds[0].StatusCode)
	}
}

func TestHooksWiring_Paginate(t *testing.T) {
	var serverURL string
	var calls int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("Content-Type", "application/json")
		if atomic.LoadInt32(&calls) == 1 {
			w.Header().Set("Link", fmt.Sprintf(`<%s/customers?page=2>; rel="next"`, serverURL))
			w.Write([]byte(`[{"id":1,"type":"Customer","first_name":"A","last_name":"B","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
			return
		}
		if atomic.LoadInt32(&calls) == 2 {
			w.Header().Set("Link", fmt.Sprintf(`<%s/customers?page=3>; rel="next"`, serverURL))
			w.Write([]byte(`[{"id":2,"type":"Customer","first_name":"C","last_name":"D","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
			return
		}
		w.Write([]byte(`[{"id":3,"type":"Customer","first_name":"E","last_name":"F","url":"x","app_url":"y","created_at":"t","updated_at":"t"}]`))
	}))
	defer ts.Close()
	serverURL = ts.URL

	rec := &recordingHooksRecords{}
	cfg := DefaultConfig()
	cfg.BaseURL = ts.URL
	c, err := NewClient(cfg, NewStaticTokenProvider("tok"), WithHooks(rec))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if _, err := c.GetAllCustomers(context.Background(), nil, nil); err != nil {
		t.Fatalf("GetAllCustomers: %v", err)
	}

	rec.mu.Lock()
	defer rec.mu.Unlock()
	if len(rec.paginates) != 2 {
		t.Fatalf("expected 2 OnPaginate calls (pages 2 and 3), got %d", len(rec.paginates))
	}
	if rec.paginates[0].page != 2 {
		t.Errorf("expected first paginate page=2, got %d", rec.paginates[0].page)
	}
	if rec.paginates[1].page != 3 {
		t.Errorf("expected second paginate page=3, got %d", rec.paginates[1].page)
	}
	if want := serverURL + "/customers?page=2"; rec.paginates[0].url != want {
		t.Errorf("expected paginate URL %q, got %q", want, rec.paginates[0].url)
	}
}

func TestHooksWiring_NoopNoPanic(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	}))
	defer ts.Close()

	cfg := DefaultConfig()
	cfg.BaseURL = ts.URL
	// No Hooks configured (defaults to NoopHooks) and nil opts.
	c, err := NewClient(cfg, NewStaticTokenProvider("tok"))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if _, err := c.GetAllCustomers(context.Background(), nil, nil); err != nil {
		t.Fatalf("GetAllCustomers: %v", err)
	}
}
