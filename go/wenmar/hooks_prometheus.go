package wenmar

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusHook exposes operation/request/retry counters.
type PrometheusHook struct {
	opsTotal   *prometheus.CounterVec
	httpTotal  *prometheus.CounterVec
	retryTotal *prometheus.CounterVec
}

func NewPrometheusHooks(reg prometheus.Registerer, _ prometheus.Gatherer) *PrometheusHook {
	h := &PrometheusHook{
		opsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "wenmar_operations_total",
			Help: "Total Wenmar SDK operations",
		}, []string{"operation", "status"}),
		httpTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "wenmar_http_requests_total",
			Help: "Total HTTP requests made by the Wenmar SDK",
		}, []string{"method", "status"}),
		retryTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "wenmar_retry_total",
			Help: "Total retries by the Wenmar SDK",
		}, []string{"method"}),
	}
	reg.MustRegister(h.opsTotal, h.httpTotal, h.retryTotal)
	return h
}

func (h *PrometheusHook) OnOperationStart(ctx context.Context, _ OperationInfo) context.Context { return ctx }
func (h *PrometheusHook) OnOperationEnd(_ context.Context, info OperationInfo, result OperationResult) {
	status := "success"
	if result.Err != nil {
		status = "error"
	}
	h.opsTotal.WithLabelValues(info.Operation, status).Inc()
}
func (h *PrometheusHook) OnRequestStart(_ context.Context, _ RequestInfo)    {}
func (h *PrometheusHook) OnRequestEnd(_ context.Context, info RequestInfo, result RequestResult) {
	h.httpTotal.WithLabelValues(result.Method, statusLabel(result.StatusCode)).Inc()
}
func (h *PrometheusHook) OnRetry(_ context.Context, info RequestInfo, _ int, _ error) {
	h.retryTotal.WithLabelValues(info.Method).Inc()
}
func (h *PrometheusHook) OnPaginate(_ context.Context, _ string, _ int) {}

func statusLabel(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "2xx"
	case code >= 300 && code < 400:
		return "3xx"
	case code >= 400 && code < 500:
		return "4xx"
	case code >= 500:
		return "5xx"
	default:
		return "unknown"
	}
}
