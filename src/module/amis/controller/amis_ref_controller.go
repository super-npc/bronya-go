package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
)

func BindMany2OneView(c echo.Context) error {
	bindMany2OneClassField := c.Request().Header.Get("bind-many2-one-class-field") // 绑定的 many2one 字段
	fmt.Printf(bindMany2OneClassField)
	return resp.Success(c, nil)
}
