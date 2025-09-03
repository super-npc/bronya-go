package controller

import (
	"fmt"
	"path/filepath"
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
		if strings.EqualFold(path, "") {
			panic("app module path is empty")
		}
		moduleMenu := amisMenu.ModuleMenu
		field := moduleMenu.Field_ //github.com/super-npc/bronya-go/src/module/sys/user_po
		pkgPath := field.PkgPath
		beanPath := strings.TrimPrefix(pkgPath, path) // 路径相减得到bean路径  /src/module/sys/user_po
		//计算json资源路径 "get:/src/module/sys/user_po/SysThreadPool.json",
		//snakeCase := strutil.SnakeCase(beanName)
		basePath := filepath.Dir(beanPath)
		newPath := strings.Replace(basePath, "/src", "/json", 1)
		var jsonPath = "get:" + newPath + "/" + beanName + ".json"

		siteDto := dto.SiteDto{}
		siteDto.Label = moduleMenu.Comment
		siteDto.SchemaApi = jsonPath

		var leafs, ok = tbl.Get(moduleMenu.Group, moduleMenu.Menu)
		if !ok {
			// 需要初始化
			si := make([]dto.SiteDto, 0)
			leafs = &si
			tbl.Put(moduleMenu.Group, moduleMenu.Menu, leafs)
		}
		*leafs = append(*leafs, siteDto)
		tbl.Put(moduleMenu.Group, moduleMenu.Menu, leafs)

		// 绑定模块
		module := moduleMenu.Module
		SiteModuleMap[module] = tbl
	}
}

func Site(c echo.Context) error {
	GetSiteModuleMap()
	menuTable := SiteModuleMap["系统"]
	groups := menuTable.Rows()
	groupLeafs := make([]resp.Module, 0)
	for _, group := range groups {
		groupId := util.ToPinyin(group)

		groupLeaf := resp.Module{}
		groupLeaf.Id = groupId
		groupLeaf.ParentId = "0"
		groupLeaf.Label = group
		// 组
		menuSiteDto := menuTable.Row(group)

		menuLeafs := make([]resp.Menu, 0)
		for menuName, siteDtoList := range menuSiteDto {
			// 第一层菜单
			menuId := util.ToPinyin(menuName)
			menuLeaf := resp.Menu{}
			menuLeaf.Id = groupId
			menuLeaf.ParentId = "0"
			menuLeaf.Label = group

			leafs := make([]resp.Leaf, 0)
			for _, siteDto := range *siteDtoList {
				// 菜单叶子集合
				leafId := util.ToPinyin(siteDto.Label)
				leaf := resp.Leaf{}
				leaf.ParentId = menuId
				leaf.Id = leafId
				leaf.Label = siteDto.Label
				leaf.Url = fmt.Sprintf("/%s/%s/%s", groupId, menuId, leafId)
				leaf.SchemaApi = siteDto.SchemaApi
				//leaf.Icon = "/icon/香蕉水果.svg"
				leafs = append(leafs, leaf)
			}
			menuLeaf.Children = leafs
			menuLeafs = append(menuLeafs, menuLeaf)
		}
		groupLeaf.Menu = menuLeafs
		groupLeafs = append(groupLeafs, groupLeaf)
	}

	siteResp := resp.SiteResp{}
	siteResp.Pages = groupLeafs

	return resp.Success(c, siteResp)
}
