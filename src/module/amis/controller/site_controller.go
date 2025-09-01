package controller

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/util"
	tool "github.com/super-npc/bronya-go/src/commons/util/guava"
	"github.com/super-npc/bronya-go/src/module/amis/controller/dto"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
)

var SiteModuleMap = map[string]*tool.Table[string, string, *[]dto.SiteDto]{}

func GetSiteModuleMap() {
	amisMenus := util.GetAmisMenus()

	// module, menu
	var tbl = tool.NewTable[string, string, *[]dto.SiteDto]()

	// 整理菜单, module
	for beanName, amisMenu := range amisMenus {
		path := amisMenu.ModulePath // github.com/super-npc/bronya-go
		field := amisMenu.Field_    //github.com/super-npc/bronya-go/src/module/sys/user_po
		pkgPath := field.PkgPath
		beanPath := strings.TrimPrefix(pkgPath, path) // 路径相减得到bean路径  /src/module/sys/user_po
		//计算json资源路径 "get:/src/module/sys/user_po/SysThreadPool.json",
		var jsonPath = "get:" + beanPath + "/" + beanName + ".json"
		menu := amisMenu.Menu

		siteDto := dto.SiteDto{}
		siteDto.Label = menu.Comment
		siteDto.SchemaApi = jsonPath

		var leafs, ok = tbl.Get(menu.Group, menu.Menu)
		if !ok {
			// 需要初始化
			si := make([]dto.SiteDto, 0)
			leafs = &si
			tbl.Put(menu.Group, menu.Menu, leafs)
		}
		*leafs = append(*leafs, siteDto)
		tbl.Put(menu.Group, menu.Menu, leafs)

		// 绑定模块
		module := menu.Module
		SiteModuleMap[module] = tbl
	}
}

func Site(c echo.Context) error {
	GetSiteModuleMap()
	menuTable := SiteModuleMap["系统"]
	groups := menuTable.Rows()
	//var sort = 1;
	for _, group := range groups {
		groupId := util.ToPinyin(group)

		groupLeaf := resp.Leaf{}
		groupLeaf.Id = groupId
		groupLeaf.ParentId = "0"
		groupLeaf.Label = group
		// 组
		menuSiteDto := menuTable.Row(group)
		for menuName, siteDtoList := range menuSiteDto {
			// 第一层菜单
			menuId := util.ToPinyin(group)
			menuLeaf := resp.Leaf{}
			menuLeaf.Id = groupId
			menuLeaf.ParentId = "0"
			menuLeaf.Label = group

			for _, siteDto := range *siteDtoList {
				// 叶子
				// 菜单叶子集合
				leaf := resp.Leaf{}
				leaf.ParentId = menuId
				leaf.Id = ""
				leaf.Label = menuName
				leaf.Url = ""
				leaf.SchemaApi = siteDto.SchemaApi
				leaf.Icon = "/icon/香蕉水果.svg"
			}
		}
	}

	siteResp := resp.SiteResp{}

	return resp.Success(c, siteResp)
}
