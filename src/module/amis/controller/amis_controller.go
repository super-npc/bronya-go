package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/module/amis/controller/req"
)

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
	var body map[string]interface{}
	if c.Bind(&body) != nil {
		return errors.New("非法参数")
	}
	amisHeader := getAmisHeader(c.Request().Header)
	byStruct := changeMapByStruct(amisHeader, body)
	marshal, _ := json.Marshal(byStruct)
	var json1 = string(marshal)
	fmt.Println(json1)
	bean, err := util.NewStructFromJSONAndName(amisHeader.Bean, marshal)
	if err != nil {
		return errors.New("未注册bean")
	}
	fmt.Println(bean)
	return c.String(http.StatusOK, "Hello, World!")
}

func changeMapByStruct(header req.AmisHeader, body map[string]interface{}) map[string]interface{} {
	var bodyNew map[string]interface{} = make(map[string]interface{})
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
