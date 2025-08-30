package framework

import (
	"bronya-go-demo/src/app_framework"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/framework/register"
	"github.com/super-npc/bronya-go/src/module/amis/controller"
	"github.com/super-npc/bronya-go/src/module/sys/user_po"
)

func AppStart(e *echo.Echo) {
	registerDatabaseBean() // 初始化数据库表对象

	// 初始化数据库
	register.InitDatabase()

	// 注入依赖，打破循环引用
	controller.SetDbProvider(register.GetDbProvider())

	register.InitRouting(e) // 注册路由

	e.Static("/", "public")

	e.HTTPErrorHandler = app_framework.CustomHttpErrorHandler // 全局异常处理

	// todo 配置文件读取
	//err := conf.InitSettings()
	//if err != nil {
	//	panic(err)
	//}
}

func registerDatabaseBean() {
	util.Register(util.RegisterReq{Po: &user_po.UserPo{}, Proxy: &user_po.UserPoProxy{}})
}
