package logger

import (
	"github.com/super-npc/bronya-go/src/framework/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

// InitLogger 初始化日志系统
func InitLogger() error {
	var config zap.Config

	if conf.Conf.Mode == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	// 设置日志级别
	level, err := zap.ParseAtomicLevel(conf.Conf.LogConfig.Level)
	if err != nil {
		return err
	}
	config.Level = level

	// 设置编码器
	if conf.Conf.Mode == "production" {
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.StacktraceKey = "stacktrace"
	} else {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// 设置输出路径
	if conf.Conf.LogConfig.Filename != "" {
		// 使用 lumberjack 进行日志切割
		logWriter := &lumberjack.Logger{
			Filename:   conf.Conf.LogConfig.Filename,
			MaxSize:    conf.Conf.LogConfig.MaxSize,
			MaxBackups: conf.Conf.LogConfig.MaxBackups,
			MaxAge:     conf.Conf.LogConfig.MaxAge,
			Compress:   true,
		}

		// 设置输出到文件和控制台
		config.OutputPaths = []string{conf.Conf.LogConfig.Filename}
		config.ErrorOutputPaths = []string{conf.Conf.LogConfig.Filename}

		// 使用日志写入器
		_ = logWriter

		// 创建日志器
		Logger, err = config.Build(
			zap.AddCallerSkip(1),
			zap.WrapCore(func(core zapcore.Core) zapcore.Core {
				return zapcore.NewTee(
					core,
					zapcore.NewCore(
						zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
							MessageKey:  "msg",
							LevelKey:    "level",
							EncodeLevel: zapcore.CapitalColorLevelEncoder,
							TimeKey:     "ts",
							EncodeTime:  zapcore.ISO8601TimeEncoder,
						}),
						zapcore.AddSync(os.Stdout),
						level,
					),
				)
			}),
		)
	} else {
		// 开发环境输出到控制台
		config.OutputPaths = []string{"stdout"}
		config.ErrorOutputPaths = []string{"stderr"}
		Logger, err = config.Build(zap.AddCallerSkip(1))
	}

	if err != nil {
		return err
	}

	return nil
}

// Sync 确保所有日志都被写入
func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}

// Debug 输出调试日志
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Info 输出信息日志
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Warn 输出警告日志
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Error 输出错误日志
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Fatal 输出致命错误日志
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

// WithFields 创建带字段的日志器
func WithFields(fields ...zap.Field) *zap.Logger {
	return Logger.With(fields...)
}

// GetLevel 获取当前日志级别
func GetLevel() zapcore.Level {
	return Logger.Level()
}
