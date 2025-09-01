package controller

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/commons/util/guava"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
)

func Site(c echo.Context) error {
	amisMenus := util.GetAmisMenus()

	// module, menu
	var tbl = tool.NewTable[string, string, string]()
	// 整理菜单, module
	for beanName, amisMenu := range amisMenus {
		path := amisMenu.ModulePath // github.com/super-npc/bronya-go
		field := amisMenu.Field_    //github.com/super-npc/bronya-go/src/module/sys/user_po
		pkgPath := field.PkgPath
		beanPath := strings.TrimPrefix(pkgPath, path) // 路径相减得到bean路径  /src/module/sys/user_po
		//计算json资源路径 "get:/src/module/sys/user_po/SysThreadPool.json",
		var jsonPath = "get:" + beanPath + "/" + beanName + ".json"
		fmt.Println(jsonPath)

		menu := amisMenu.Menu

		tbl.Put(menu.Group, menu.Menu, "1")
	}

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
