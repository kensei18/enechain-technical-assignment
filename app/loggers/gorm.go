package loggers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type gormLogger struct {
	*RequestLogger
}

var slowThreshold = 200 * time.Millisecond

func NewGormLogger(logger *RequestLogger) logger.Interface {
	return &gormLogger{logger}
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.RequestLogger.Info(ctx, fmt.Sprintf(msg, data...))
}

func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.RequestLogger.Warn(ctx, fmt.Sprintf(msg, data...))
}

func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.RequestLogger.Error(ctx, fmt.Sprintf(msg, data...))
}

func (l *gormLogger) Trace(
	ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error,
) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	switch {
	case err != nil && !errors.Is(err, logger.ErrRecordNotFound):
		msg := "%s %s\n[%.3fms] [rows:%v] %s"
		if rows == -1 {
			l.Error(ctx, fmt.Sprintf(msg, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.Error(ctx, fmt.Sprintf(msg, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	case elapsed > slowThreshold:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		msg := "%s %s\n[%.3fms] [rows:%v] %s"
		if rows == -1 {
			l.Warn(
				ctx, fmt.Sprintf(msg, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql),
			)
		} else {
			l.Warn(
				ctx, fmt.Sprintf(msg, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql),
			)
		}
	default:
		msg := "%s\n[%.3fms] [rows:%v] %s"
		if rows == -1 {
			l.Info(ctx, fmt.Sprintf(msg, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.Info(ctx, fmt.Sprintf(msg, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}

	}
}
