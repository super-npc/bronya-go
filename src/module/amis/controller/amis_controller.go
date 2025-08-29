package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/db"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/module/amis/controller/req"
)

// 使用全局变量存储DB提供者
var dbProvider db.DBProvider

// SetDbProvider 设置数据库提供者（由外部注入）
func SetDbProvider(provider db.DBProvider) {
	dbProvider = provider
}

func Page(c echo.Context) error {
	pageReq := new(req.PageReq)
	if c.Bind(pageReq) != nil {
		return errors.New("非法参数")
	}
	amisHeader := getAmisHeader(c.Request().Header)
	_, err := util.NewStructFromJSONAndName(amisHeader.Bean, []byte{})

	if err != nil {
		return errors.New("")
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func View(c echo.Context) error {
	viewReq := new(req.ViewReq)
	if c.Bind(viewReq) != nil {
		return errors.New("非法参数")
	}

	_, err := util.NewStructFromJSONAndName("", []byte{})
	if err != nil {
		return errors.New("")
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func Create(c echo.Context) error {
	bean := findBean(c)
	// 使用接口获取数据库连接
	res := dbProvider.GetDb().Create(bean)
	if res.RowsAffected == 0 {
		return errors.New("新增失败")
	}
	return c.String(http.StatusOK, "Hello, World! ")
}

func Update(c echo.Context) error {
	bean := findBean(c)
	// 使用接口获取数据库连接
	res := dbProvider.GetDb().Updates(bean)
	if res.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return c.String(http.StatusOK, "Hello, World! ")
}

func DeleteBatch(c echo.Context) error {
	deleteBatchReq := new(req.DeleteBatchReq)
	if c.Bind(deleteBatchReq) != nil {
		return errors.New("非法参数")
	}
	amisHeader := getAmisHeader(c.Request().Header)
	bean, err := util.NewStructFromName(amisHeader.Bean)
	if err != nil {
		return errors.New("未注册bean")
	}

	// 使用接口获取数据库连接
	ids := strings.Split(deleteBatchReq.Ids, ",")
	res := dbProvider.GetDb().Delete(bean, ids)
	if res.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return c.String(http.StatusOK, "Hello, World! ")
}

func findBean(c echo.Context) interface{} {
	var body map[string]interface{}
	if c.Bind(&body) != nil {
		return nil
	}
	amisHeader := getAmisHeader(c.Request().Header)
	byStruct := changeMapByStruct(amisHeader, body)
	marshal, _ := json.Marshal(byStruct)
	bean, err := util.NewStructFromJSONAndName(amisHeader.Bean, marshal)
	if err != nil {
		return nil
	}
	return bean
}

func changeMapByStruct(header req.AmisHeader, body map[string]interface{}) map[string]interface{} {
	var bodyNew = make(map[string]interface{})
	var beanPre = header.Bean + "__"
	for s, i := range body {
		key := strings.TrimPrefix(s, beanPre)
		bodyNew[key] = i
	}
	return bodyNew
}

func getAmisHeader(header http.Header) req.AmisHeader {
	bean := header.Get("bean")
	site := header.Get("site")
	return req.AmisHeader{Bean: bean, Site: site}
}
