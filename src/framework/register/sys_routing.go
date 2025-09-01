package register

import (
	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/module/amis/controller"
)

func InitRouting(e *echo.Echo) {
	routingAdmin(e)
}

func routingAdmin(e *echo.Echo) {
	admin := e.Group("/admin/base")
	{
		admin.POST("/amis/page", controller.Page)
		admin.POST("/amis/view", controller.View)
		admin.POST("/amis/create", controller.Create)
		admin.POST("/amis/update", controller.Update)
		admin.POST("/amis/delete-batch", controller.DeleteBatch)

		admin.GET("/site", controller.Site)
		admin.GET("/sys/top-right-header", controller.TopRightHeader)
		admin.GET("/sys/app-info", controller.AppInfo)

	}
}
