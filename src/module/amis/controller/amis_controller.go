package controller

import (
	"encoding/json"
	"errors"
	"fmt"
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
	fmt.Println(res)
	return c.String(http.StatusOK, "Hello, World! ")
}

func Update(c echo.Context) error {
	bean := findBean(c)
	// 使用接口获取数据库连接
	res := dbProvider.GetDb().Updates(bean) // 得到该bean是通过反射得到的,导致执行后reflect: reflect.Value.SetUint
	fmt.Println(res)
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
