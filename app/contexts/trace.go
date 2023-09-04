package contexts

import (
	"context"
	"errors"
)

const traceIDKeyName = "traceId"

type traceIDKey string

func GetTraceID(ctx context.Context) (string, error) {
	err := errors.New("failed to get traceID from context")
	var key traceIDKey = traceIDKeyName
	value := ctx.Value(key)
	if value == nil {
		return "", err
	}
	traceID, ok := value.(string)
	if !ok {
		return "", err
	}
	return traceID, nil
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	var key traceIDKey = traceIDKeyName
	return context.WithValue(ctx, key, traceID)
}
