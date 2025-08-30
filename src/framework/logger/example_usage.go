package logger

// ExampleUsage 日志使用示例
//
// 基本用法:
// import "github.com/super-npc/bronya-go/src/framework/logger"
//
// logger.Info("用户登录成功", zap.String("username", "john"), zap.Int("user_id", 123))
// logger.Error("数据库连接失败", zap.Error(err))
// logger.Debug("调试信息", zap.Any("data", someStruct))
//
// 性能监控:
// start := time.Now()
// // ... 业务逻辑
// logger.LogPerformance("用户查询", time.Since(start))
//
// 业务日志:
// logger.LogBusiness("订单创建", "success", zap.String("order_id", "ORD123"))
//
// 数据库日志:
// logger.LogDatabase("INSERT", "users", time.Since(start), rowsAffected)
//
// 安全日志:
// logger.LogSecurity("登录失败", "192.168.1.1", "unknown_user")
//
// 环境变量设置:
// export APP_MODE=production
// export APP_LOG_LEVEL=debug
// export APP_MYSQL_HOST=localhost
//
// 启动应用:
// APP_MODE=production ./your-app
// 或
// APP_MODE=develop ./your-app
