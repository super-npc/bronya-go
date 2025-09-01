package log

import (
	"time"

	"go.uber.org/zap"
)

// LogContext 日志上下文
func LogContext(requestID, traceID string) *zap.Logger {
	return Logger.With(
		zap.String("request_id", requestID),
		zap.String("trace_id", traceID),
	)
}

// LogWithContext 带上下文的日志
func LogWithContext(requestID, traceID string) *zap.Logger {
	return LogContext(requestID, traceID)
}

// LogErrorWithContext 记录错误日志
func LogErrorWithContext(requestID, traceID, errMsg string, err error) {
	LogContext(requestID, traceID).Error(errMsg,
		zap.Error(err),
		zap.String("error_type", "application_error"),
	)
}

// LogRequest 记录请求日志
func LogRequest(method, path, ip string, duration time.Duration, status int) {
	Info("HTTP请求",
		zap.String("method", method),
		zap.String("path", path),
		zap.String("ip", ip),
		zap.Duration("duration", duration),
		zap.Int("status", status),
	)
}

// LogDatabase 记录数据库操作日志
func LogDatabase(operation, table string, duration time.Duration, rows int64) {
	Debug("数据库操作",
		zap.String("operation", operation),
		zap.String("table", table),
		zap.Duration("duration", duration),
		zap.Int64("rows_affected", rows),
	)
}

// LogPerformance 记录性能日志
func LogPerformance(operation string, duration time.Duration, extraFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("operation", operation),
		zap.Duration("duration", duration),
		zap.String("type", "performance"),
	}
	fields = append(fields, extraFields...)
	Info("性能监控", fields...)
}

// LogBusiness 记录业务日志
func LogBusiness(operation, status string, extraFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("operation", operation),
		zap.String("status", status),
		zap.String("type", "business"),
	}
	fields = append(fields, extraFields...)
	Info("业务操作", fields...)
}

// LogSecurity 记录安全日志
func LogSecurity(event, ip, user string, extraFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("event", event),
		zap.String("ip", ip),
		zap.String("user", user),
		zap.String("type", "security"),
	}
	fields = append(fields, extraFields...)
	Warn("安全事件", fields...)
}
