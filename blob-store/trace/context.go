package trace

import (
	"context"
	"github.com/google/uuid"
)

type ctxKey int8

const (
	LogFieldName           = "trace_id"
	traceIDFieldKey ctxKey = iota
)

func NewContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDFieldKey, traceID)
}

func FromContext(ctx context.Context) string {
	traceID, ok := ctx.Value(traceIDFieldKey).(string)
	if !ok {
		return uuid.NewString()
	}

	return traceID
}
