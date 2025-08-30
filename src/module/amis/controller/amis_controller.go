package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons/constant"
	"github.com/super-npc/bronya-go/src/commons/db"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/module/amis/amis_proxy"
	"github.com/super-npc/bronya-go/src/module/amis/controller/req"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
	"gorm.io/gorm"
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

	// 将结果转换为 []map[string]amis_proxy{}
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

	// 处理动态代理数据 todo 后续优化点,不需要顺序处理,go协程处理多任务
	for _, m := range list {
		table(amisHeader, registerObj, m)
	}

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
	beanTableMap := table(amisHeader, registerObj, poBean)
	return resp.Success(c, beanTableMap)
}

func Create(c echo.Context) error {
	reqBody, registerObj := findBean(c)
	poBean := registerObj.Po

	// 通过反射调用BeforeAdd方法
	var res *gorm.DB
	if registerObj.Proxy != nil {
		proxy := registerObj.Proxy.(amis_proxy.IAmisProxy)
		proxy.BeforeAdd(reqBody)
		res = dbProvider.GetDb().Create(poBean)
		proxy.AfterAdd(poBean)
	} else {
		res = dbProvider.GetDb().Create(poBean)
	}

	if res.RowsAffected == 0 {
		return errors.New("新增失败")
	}
	return resp.Success(c, fmt.Sprintf("新增 %d", res.RowsAffected))
}

func Update(c echo.Context) error {
	reqBody, registerObj := findBean(c)
	poBean := registerObj.Po
	// 使用接口获取数据库连接
	var res *gorm.DB

	if registerObj.Proxy != nil {
		proxy := registerObj.Proxy.(amis_proxy.IAmisProxy)
		proxy.BeforeAdd(reqBody)
		res = dbProvider.GetDb().Updates(poBean)
		proxy.AfterUpdate(poBean)
	} else {
		res = dbProvider.GetDb().Updates(poBean)
	}
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
	tx := dbProvider.GetDb().Delete(poBean, ids)
	if tx.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return resp.Success(c, poBean)
}

func table(header req.AmisHeader, registerObj util.RegisterResp, record interface{}) map[string]interface{} {
	var tableFieldPre = header.Bean + constant.AmisSplitSymbol
	var resTable = make(map[string]interface{})
	for k, v := range util.StructToMap(record) {
		if strings.EqualFold(k, constant.PrimaryKey) {
			resTable[k] = v
		}
		resTable[tableFieldPre+k] = v
	}
	// 对数据加上前缀
	if registerObj.Proxy != nil {
		a := registerObj.Proxy.(amis_proxy.IAmisProxy)
		a.Table(tableFieldPre, resTable)
	}
	return resTable
}

func findBean(c echo.Context) (map[string]interface{}, util.RegisterResp) {
	var reqBody map[string]interface{}
	if c.Bind(&reqBody) != nil {
		panic("无法绑定请求参数")
	}
	amisHeader := getAmisHeader(c.Request().Header)
	byStruct := changeMapByStruct(amisHeader, reqBody)
	marshal, _ := json.Marshal(byStruct)
	bean := util.NewStructFromJSONAndName(amisHeader.Bean, marshal)

	return reqBody, bean
}

func changeMapByStruct(header req.AmisHeader, body map[string]interface{}) map[string]interface{} {
	var bodyNew = make(map[string]interface{})
	var beanPre = header.Bean + constant.AmisSplitSymbol
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
