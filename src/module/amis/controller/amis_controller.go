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
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
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
	//amisHeader := getAmisHeader(c.Request().Header)
	//bean, err := util.NewStructFromName(amisHeader.Bean)
	////query.Eq(&model.Username, users[0].Username).Or().Eq(&model.Username, users[5].Username)
	//if err != nil {
	//	return errors.New("未注册bean")
	//}

	//page := 1      // 第几页，从 1 开始
	//pageSize := 10 // 每页条数
	//offset := (page - 1) * pageSize

	//var aa = []bean
	//dbProvider.GetDb().Limit(pageSize).Offset(offset).Find(&aa)
	//fmt.Println("page", aa)

	//page := gplus.NewPage[bean](1, 10)
	//query, _ := gplus.NewQuery[bean]()
	//resultPage, _ := gplus.SelectPage(page, query)
	//var list []map[string]interface{}
	//for _, record := range resultPage.Records {
	//	var result map[string]interface{}
	//	marshal, err := json.Marshal(record)
	//	err1 := json.Unmarshal(marshal, &result)
	//	if err1 != nil {
	//		panic(err)
	//	}
	//	list = append(list, result)
	//}
	//res := resp.PageRes{Total: resultPage.Total, Rows: list}
	//return resp.Success(c, res)
	return c.String(http.StatusOK, "Hello, World! ")
}

func View(c echo.Context) error {
	viewReq := new(req.ViewReq)
	if c.Bind(viewReq) != nil {
		return errors.New("非法参数")
	}

	amisHeader := getAmisHeader(c.Request().Header)
	bean, err := util.NewStructFromName(amisHeader.Bean)

	if err != nil {
		return errors.New("未注册bean")
	}
	res := dbProvider.GetDb().First(bean, viewReq.Id)

	if res.RowsAffected == 0 {
		return errors.New("不能存在记录")
	}
	return resp.Success(c, bean)
}

func Create(c echo.Context) error {
	bean := findBean(c)
	// 使用接口获取数据库连接
	res := dbProvider.GetDb().Create(bean)
	if res.RowsAffected == 0 {
		return errors.New("新增失败")
	}
	return resp.Success(c, fmt.Sprintf("新增 %d", res.RowsAffected))
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
	return resp.Success(c, bean)
}

// todo 返回类型变了
func findBean(c echo.Context) util.RegisterObj {
	var body map[string]interface{}
	if c.Bind(&body) != nil {
		panic("无法绑定请求参数")
	}
	amisHeader := getAmisHeader(c.Request().Header)
	byStruct := changeMapByStruct(amisHeader, body)
	marshal, _ := json.Marshal(byStruct)
	bean := util.NewStructFromJSONAndName(amisHeader.Bean, marshal)

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
