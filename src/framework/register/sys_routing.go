package register

import (
	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/module/amis/controller"
)

func InitRouting(e *echo.Echo) {
	routingAdmin(e)
}

func routingAdmin(e *echo.Echo) {
	admin := e.Group("/admin")
	{
		admin.GET("/", controller.Page)
	}
}
