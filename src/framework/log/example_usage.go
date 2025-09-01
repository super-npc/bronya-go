package log

//
//import (
//	"errors"
//	"time"
//
//	"github.com/super-npc/bronya-go/src/framework/log"
//	"go.uber.org/zap"
//)
//
////ExampleUsage 日志使用示例
////
////环境变量设置:
////export APP_MODE=production
////export APP_LOG_LEVEL=debug
////export APP_MYSQL_HOST=localhost
////
////启动应用:
////APP_MODE=production ./your-app
////或
////APP_MODE=develop ./your-app
//
////基本用法:
//
//func demo() {
//	log.Info("用户登录成功", zap.String("username", "john"), zap.Int("user_id", 123))
//	err := errors.New("异常1")
//	log.Error("数据库连接失败", zap.Error(err))
//	log.Debug("调试信息", zap.Any("data", "val"))
//
//	//性能监控:
//	start := time.Now()
//	// ... 业务逻辑
//	log.LogPerformance("用户查询", time.Since(start))
//
//	//业务日志:
//	log.LogBusiness("订单创建", "success", zap.String("order_id", "ORD123"))
//
//	//数据库日志:
//	log.LogDatabase("INSERT", "users", time.Since(start), 1)
//
//	//安全日志:
//	log.LogSecurity("登录失败", "192.168.1.1", "unknown_user")
//}
