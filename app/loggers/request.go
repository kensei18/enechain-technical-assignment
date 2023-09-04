package loggers

import (
	"context"
	"io"
	"log/slog"

	"github.com/kensei18/enechain-technical-assignment/app/contexts"
)

type RequestLogger struct {
	logger func(context.Context) *slog.Logger
}

func NewDefaultLogger(writer io.Writer, level slog.Level) *RequestLogger {
	logger := slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: level}))

	loggerFunc := func(ctx context.Context) *slog.Logger {
		traceID, _ := contexts.GetTraceID(ctx)
		return logger.With("traceId", traceID)
	}

	return &RequestLogger{logger: loggerFunc}
}

func (l *RequestLogger) Debug(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.logger(ctx).LogAttrs(ctx, slog.LevelDebug, msg, attrs...)
}

func (l *RequestLogger) Info(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.logger(ctx).LogAttrs(ctx, slog.LevelInfo, msg, attrs...)
}

func (l *RequestLogger) Warn(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.logger(ctx).LogAttrs(ctx, slog.LevelWarn, msg, attrs...)
}

func (l *RequestLogger) Error(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.logger(ctx).LogAttrs(ctx, slog.LevelError, msg, attrs...)
}
