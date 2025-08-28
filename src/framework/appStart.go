package framework

import (
	"bronya-go-demo/src/appFramework"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/framework/conf"
	"github.com/super-npc/bronya-go/src/framework/register"
	"github.com/super-npc/bronya-go/src/module/sys/domain"
)

func AppStart(e *echo.Echo) {
	registerDatabaseBean() // 初始化数据库表对象

	register.InitRouting(e) // 注册路由

	e.HTTPErrorHandler = appFramework.CustomHttpErrorHandler // 全局异常处理

	err := conf.InitSettings()
	if err != nil {
		panic(err)
	}
	// 初始化中间件
	register.InitDatabase()
}

func registerDatabaseBean() {
	util.RegisterByStruct(&domain.User{})
}
