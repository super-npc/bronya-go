package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/framework/log"
	"go.uber.org/zap"
)

// LoggerMiddleware 自定义日志中间件
func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// 继续处理请求
			err := next(c)

			// 计算请求耗时
			latency := time.Since(start)

			// 获取状态码
			status := c.Response().Status
			if err != nil {
				if httpErr, ok := err.(*echo.HTTPError); ok {
					status = httpErr.Code
				} else {
					status = 500
				}
			}

			// 记录日志
			log.Info("HTTP请求",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.String("query", c.Request().URL.RawQuery),
				zap.String("ip", c.RealIP()),
				zap.Int("status", status),
				zap.Duration("latency", latency),
				zap.String("user_agent", c.Request().UserAgent()),
				zap.Int64("response_size", c.Response().Size),
			)

			return err
		}
	}
}

// RequestLogger 请求日志中间件（简化版）
func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			latency := time.Since(start)
			status := c.Response().Status

			log.Info("HTTP",
				zap.String("method", c.Request().Method),
				zap.String("uri", c.Request().RequestURI),
				zap.Int("status", status),
				zap.String("latency", latency.String()),
				zap.String("ip", c.RealIP()),
			)

			return err
		}
	}
}
