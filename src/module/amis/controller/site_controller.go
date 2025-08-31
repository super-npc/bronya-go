package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
)

func Site(c echo.Context) error {
	amisMenus := util.GetAmisMenus()

	for beanName, amisMenuObj := range amisMenus {
		fmt.Println(beanName, amisMenuObj)
		field := amisMenuObj.Field_
		amisMenu := amisMenuObj.Menu
		fmt.Println(amisMenu)

		fmt.Println(field.PkgPath) // github.com/super-npc/bronya-go/src/module/sys/user_po
	}

	siteResp := resp.SiteResp{}

	return resp.Success(c, siteResp)
}
