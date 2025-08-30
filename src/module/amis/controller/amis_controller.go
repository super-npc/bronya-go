package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
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
	amisHeader := getAmisHeader(c.Request().Header)
	registerObj := util.NewStructFromName(amisHeader.Bean)
	poBean := registerObj.Po

	page := 1      // 第几页，从 1 开始
	pageSize := 10 // 每页条数
	offset := (page - 1) * pageSize

	// 根据 bean 的类型创建对应切片
	sliceType := reflect.SliceOf(reflect.TypeOf(poBean))
	slicePtr := reflect.New(sliceType)

	// 使用 GORM 查询数据
	dbProvider.GetDb().Limit(pageSize).Offset(offset).Find(slicePtr.Interface())

	// 获取实际的切片值
	sliceValue := slicePtr.Elem().Interface()

	// 将结果转换为 []map[string]interface{}
	var list []map[string]interface{}
	if sliceValue != nil {
		jsonData, err := json.Marshal(sliceValue)
		if err != nil {
			return errors.New("序列化数据失败")
		}
		err = json.Unmarshal(jsonData, &list)
		if err != nil {
			return errors.New("反序列化数据失败")
		}
	}

	// 获取总数
	var total int64
	dbProvider.GetDb().Model(poBean).Count(&total)

	res := resp.PageRes{Total: total, Rows: list}
	return resp.Success(c, res)
}

func View(c echo.Context) error {
	viewReq := new(req.ViewReq)
	if c.Bind(viewReq) != nil {
		return errors.New("非法参数")
	}

	amisHeader := getAmisHeader(c.Request().Header)
	registerObj := util.NewStructFromName(amisHeader.Bean)
	poBean := registerObj.Po
	res := dbProvider.GetDb().First(poBean, viewReq.Id)

	if res.RowsAffected == 0 {
		return errors.New("不能存在记录")
	}
	return resp.Success(c, poBean)
}

func Create(c echo.Context) error {
	registerObj := findBean(c)
	poBean := registerObj.Po
	// 使用接口获取数据库连接
	res := dbProvider.GetDb().Create(poBean)
	if res.RowsAffected == 0 {
		return errors.New("新增失败")
	}
	return resp.Success(c, fmt.Sprintf("新增 %d", res.RowsAffected))
}

func Update(c echo.Context) error {
	registerObj := findBean(c)
	poBean := registerObj.Po
	// 使用接口获取数据库连接
	res := dbProvider.GetDb().Updates(poBean)
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
	registerObj := util.NewStructFromName(amisHeader.Bean)
	poBean := registerObj.Po

	// 使用接口获取数据库连接
	ids := strings.Split(deleteBatchReq.Ids, ",")
	res := dbProvider.GetDb().Delete(poBean, ids)
	if res.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return resp.Success(c, poBean)
}

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
