package wenmar

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// OTelHook creates OpenTelemetry spans for each SDK operation.
type OTelHook struct {
	tracer trace.Tracer
}

func NewOTelHook(tracer trace.Tracer) *OTelHook {
	if tracer == nil {
		tracer = otel.Tracer("wenmar-sdk")
	}
	return &OTelHook{tracer: tracer}
}

func (h *OTelHook) OnOperationStart(ctx context.Context, info OperationInfo) context.Context {
	ctx, _ = h.tracer.Start(ctx, info.Operation)
	return ctx
}

func (h *OTelHook) OnOperationEnd(ctx context.Context, info OperationInfo, result OperationResult) {
	span := trace.SpanFromContext(ctx)
	if result.Err != nil {
		span.SetStatus(codes.Error, result.Err.Error())
		span.RecordError(result.Err)
	}
	span.SetAttributes(attribute.String("operation", result.Operation))
	span.End()
}

func (h *OTelHook) OnRequestStart(_ context.Context, _ RequestInfo)    {}
func (h *OTelHook) OnRequestEnd(_ context.Context, _ RequestInfo, _ RequestResult) {}
func (h *OTelHook) OnRetry(_ context.Context, _ RequestInfo, _ int, _ error)        {}
func (h *OTelHook) OnPaginate(_ context.Context, _ string, _ int)                    {}
