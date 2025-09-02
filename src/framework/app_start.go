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
	"github.com/super-npc/bronya-go/src/module/db/audit/domain/sys_data_audit"
	"github.com/super-npc/bronya-go/src/module/db/badge/domain/badge"
	"github.com/super-npc/bronya-go/src/module/db/distributed/domain/sys_distributed_lock"
	"github.com/super-npc/bronya-go/src/module/db/env/domain/sys_env_obj"
	"github.com/super-npc/bronya-go/src/module/db/env/domain/sys_env_obj_field"
	"github.com/super-npc/bronya-go/src/module/db/threadpool/domain/sys_thread_pool"
	"go.uber.org/zap"
)

func Start(e *echo.Echo) {
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

	log.Info("应用启动完成", zap.Int("port", conf.Conf.Port))
}

func registerDatabaseBean() {
	log.Info("注册数据库表对象")
	util.RegisterFramework(util.RegisterReq{Po: &sys_data_audit.SysDataAudit{}, Proxy: &sys_data_audit.SysDataAuditProxy{}})
	util.RegisterFramework(util.RegisterReq{Po: &badge.Badge{}, Proxy: &badge.BadgeProxy{}})
	util.RegisterFramework(util.RegisterReq{Po: &sys_distributed_lock.SysDistributedLock{}})
	util.RegisterFramework(util.RegisterReq{Po: &sys_env_obj.SysEnvObj{}, Proxy: &sys_env_obj.SysEnvObjProxy{}})
	util.RegisterFramework(util.RegisterReq{Po: &sys_env_obj_field.SysEnvObjField{}, Proxy: &sys_env_obj_field.SysEnvObjFieldProxy{}})
	util.RegisterFramework(util.RegisterReq{Po: &sys_thread_pool.SysThreadPool{}, Proxy: &sys_thread_pool.SysThreadPoolProxy{}})

	log.Info("注册数据库表对象.finish")
}
