package register

import (
	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/module/amis/controller"
)

func InitRouting(e *echo.Echo) {
	routingAdmin(e)
}

func routingAdmin(e *echo.Echo) {
	admin := e.Group("/admin/base/amis")
	{
		admin.POST("/page", controller.Page)
		admin.POST("/view", controller.View)
		admin.POST("/create", controller.Create)
		admin.POST("/update", controller.Update)
	}
}
