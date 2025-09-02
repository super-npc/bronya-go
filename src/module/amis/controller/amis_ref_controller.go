package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
)

func BindMany2OneView(c echo.Context) error {
	fmt.Printf("")
	return resp.Success(c, nil)
}
