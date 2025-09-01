package framework

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/framework/conf"
	"github.com/super-npc/bronya-go/src/framework/log"
	"github.com/super-npc/bronya-go/src/framework/middle_ware"
	"github.com/super-npc/bronya-go/src/framework/register"
	"github.com/super-npc/bronya-go/src/module/amis/controller"
	"github.com/super-npc/bronya-go/src/module/sys/user_po"
	"go.uber.org/zap"
)

func AppStart(e *echo.Echo) {
	// 初始化配置
	err := conf.InitSettings()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	err = log.InitLogger()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	log.Info("应用启动中...",
		zap.String("mode", conf.Conf.Mode),
		zap.String("version", conf.Conf.Version),
	)

	registerDatabaseBean() // 初始化数据库表对象

	// 初始化数据库
	register.InitDatabase()
	log.Info("数据库初始化完成")

	// 注入依赖，打破循环引用
	controller.SetDbProvider(register.GetDbProvider())

	register.InitRouting(e) // 注册路由

	// 注册中间件
	// 使用自定义日志中间件替换默认的
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())                             // 捕获 panic
	e.HTTPErrorHandler = middle_ware.CustomHttpErrorHandler // 全局异常处理

	// 配置资源
	e.Static("/", "public")

	log.Info("应用启动完成", zap.Int("port", conf.Conf.Port))
}

func registerDatabaseBean() {
	log.Info("注册数据库表对象")
	util.RegisterFramework(util.RegisterReq{Po: &user_po.UserPo{}, Proxy: &user_po.UserPoProxy{}})

	log.Info("注册数据库表对象.finish")
}
