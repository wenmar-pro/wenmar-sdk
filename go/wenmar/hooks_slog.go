package wenmar

import (
	"context"
	"log/slog"
)

// SlogHooks logs operations and requests at debug level.
type SlogHooks struct {
	logger *slog.Logger
}

func NewSlogHooks(logger *slog.Logger) *SlogHooks {
	if logger == nil {
		logger = slog.Default()
	}
	return &SlogHooks{logger: logger}
}

func (h *SlogHooks) OnOperationStart(ctx context.Context, info OperationInfo) context.Context {
	h.logger.Debug("operation start", "operation", info.Operation)
	return ctx
}

func (h *SlogHooks) OnOperationEnd(_ context.Context, info OperationInfo, result OperationResult) {
	if result.Err != nil {
		h.logger.Debug("operation end", "operation", result.Operation, "error", result.Err)
	} else {
		h.logger.Debug("operation end", "operation", result.Operation)
	}
}

func (h *SlogHooks) OnRequestStart(_ context.Context, info RequestInfo) {
	h.logger.Debug("request start", "method", info.Method, "url", info.URL)
}

func (h *SlogHooks) OnRequestEnd(_ context.Context, info RequestInfo, result RequestResult) {
	h.logger.Debug("request end", "method", result.Method, "url", result.URL, "status", result.StatusCode)
	if result.Err != nil {
		h.logger.Debug("request error", "error", result.Err)
	}
}

func (h *SlogHooks) OnRetry(_ context.Context, info RequestInfo, attempt int, err error) {
	h.logger.Debug("retry", "method", info.Method, "url", info.URL, "attempt", attempt, "error", err)
}

func (h *SlogHooks) OnPaginate(_ context.Context, url string, page int) {
	h.logger.Debug("paginate", "url", url, "page", page)
}
