package controller

import (
	"maps"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/framework/log"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
	"go.uber.org/zap"
)

func ChangeSite(c echo.Context) error {
	param := c.QueryParam("site")
	log.Info("切换主菜单", zap.String("site", param))
	return resp.Success(c, true)
}

func AppInfo(c echo.Context) error {
	infoResp := resp.AppInfoResp{}
	infoResp.AppName = "系统1"
	infoResp.GitBaseVersion = "11111"
	infoResp.GitProjectVersion = "22222"

	return resp.Success(c, infoResp)
}

func TopRightHeader(c echo.Context) error {
	siteHeader := c.Request().Header.Get("site")
	if siteHeader == "" {

	}
	moduleMap := SiteModuleMap
	modules := maps.Keys(moduleMap)

	res := resp.TopRightHeaderResp{}
	var btnList = make([]resp.SubSystemButtons, 0)
	for module := range modules {
		btnList = append(btnList, resp.SubSystemButtons{Label: module, Level: "light"})
	}
	//btnList = append(btnList, resp.SubSystemButtons{Label: "杂货铺1", Level: "light"})
	//btnList = append(btnList, resp.SubSystemButtons{Label: "系统1", Level: "primary"})
	//btnList = append(btnList, resp.SubSystemButtons{Label: "杂货铺2", Level: "light"})
	//btnList = append(btnList, resp.SubSystemButtons{Label: "杂货铺3", Level: "light"})
	res.SubSystemButtons = btnList

	roles := make([]resp.Role, 0)
	roles = append(roles, resp.Role{Id: "1", RoleName: "超级管理员"})
	roles = append(roles, resp.Role{Id: "2", RoleName: "普通用户"})

	user := resp.LoginUser{Id: "1", UserAvatar: "/admin/sys-file/image/null", UserName: "admin1", Roles: roles}
	res.LoginUser = user

	return resp.Success(c, res)
}
