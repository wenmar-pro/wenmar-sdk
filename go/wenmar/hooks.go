package wenmar

import "context"

// OperationInfo describes an SDK-level operation (e.g. ListCustomers).
type OperationInfo struct {
	Operation string
}

// OperationResult is the outcome of an operation.
type OperationResult struct {
	Operation string
	Err       error
}

// RequestInfo describes a single HTTP request.
type RequestInfo struct {
	Method string
	URL    string
}

// RequestResult is the outcome of an HTTP request.
type RequestResult struct {
	Method     string
	URL        string
	StatusCode int
	Err        error
}

// Hooks receives callbacks for SDK operations and HTTP requests. Implement
// this interface for logging, tracing, and metrics. Use NoopHooks (the
// default) for zero overhead.
//
// OnOperationStart returns a context so hooks (e.g. OTel) can attach span
// context that subsequent callbacks receive. NoopHooks returns the ctx
// unchanged.
type Hooks interface {
	OnOperationStart(ctx context.Context, info OperationInfo) context.Context
	OnOperationEnd(ctx context.Context, info OperationInfo, result OperationResult)
	OnRequestStart(ctx context.Context, info RequestInfo)
	OnRequestEnd(ctx context.Context, info RequestInfo, result RequestResult)
	OnRetry(ctx context.Context, info RequestInfo, attempt int, err error)
	OnPaginate(ctx context.Context, url string, page int)
}

// NoopHooks implements Hooks with no-op methods. Zero cost.
type NoopHooks struct{}

func (NoopHooks) OnOperationStart(ctx context.Context, _ OperationInfo) context.Context { return ctx }
func (NoopHooks) OnOperationEnd(context.Context, OperationInfo, OperationResult)      {}
func (NoopHooks) OnRequestStart(context.Context, RequestInfo)                          {}
func (NoopHooks) OnRequestEnd(context.Context, RequestInfo, RequestResult)             {}
func (NoopHooks) OnRetry(context.Context, RequestInfo, int, error)                     {}
func (NoopHooks) OnPaginate(context.Context, string, int)                              {}

// ClientOption configures a Client.
type ClientOption func(*Client)

// WithHooks sets the hooks for the client.
func WithHooks(hooks Hooks) ClientOption {
	return func(c *Client) {
		c.hooks = hooks
	}
}
