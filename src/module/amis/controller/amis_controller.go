package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/labstack/echo/v4"
	"github.com/super-npc/bronya-go/src/commons"
	"github.com/super-npc/bronya-go/src/commons/constant"
	"github.com/super-npc/bronya-go/src/commons/db"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/framework/log"
	"github.com/super-npc/bronya-go/src/module/amis/amis_proxy"
	"github.com/super-npc/bronya-go/src/module/amis/controller/req"
	"github.com/super-npc/bronya-go/src/module/amis/controller/resp"
	"github.com/super-npc/bronya-go/src/module/amis/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 使用全局变量存储DB提供者
var dbProvider db.DBProvider

// SetDbProvider 设置数据库提供者（由外部注入）
func SetDbProvider(provider db.DBProvider) {
	dbProvider = provider
	log.Info("数据库提供者已设置", zap.String("provider_type", fmt.Sprintf("%T", provider)))
}

func Page(c echo.Context) error {
	start := time.Now()
	pageReq := new(req.PageReq)
	err := c.Bind(pageReq)
	if err != nil {
		log.Warn("Page接口参数绑定失败", zap.String("error", "invalid_params"))
		panic(err)
	}

	amisHeader := getAmisHeader(c.Request().Header)
	registerObj := util.NewStructFromName(amisHeader.Bean)
	poBean := registerObj.Po

	page := 1      // 第几页，从 1 开始
	pageSize := 10 // 每页条数
	offset := (page - 1) * pageSize

	log.Debug("Page接口查询开始",
		zap.String("bean", amisHeader.Bean),
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.Int("offset", offset),
	)

	// 根据 bean 的类型创建对应切片
	sliceType := reflect.SliceOf(reflect.TypeOf(poBean))
	slicePtr := reflect.New(sliceType)

	// 使用 GORM 查询数据
	queryStart := time.Now()
	dbProvider.GetDb().Limit(pageSize).Offset(offset).Find(slicePtr.Interface())
	//sql := dbProvider.GetDb().Limit(pageSize).Offset(offset).Find(slicePtr.Interface()).Statement.SQL.String()
	//dbProvider.GetDb().Raw("select * from ...").Limit(pageSize).Offset(offset).Find(slicePtr.Interface())
	service.One2Many(amisHeader.Bean, pageReq) // 准备1:n sql拼装

	queryDuration := time.Since(queryStart)

	// 获取实际的切片值
	sliceValue := slicePtr.Elem().Interface()

	// 将结果转换为 []map[string]amis_proxy{}
	var list []map[string]interface{}
	if sliceValue != nil {
		jsonData, err := json.Marshal(sliceValue)
		if err != nil {
			log.Error("Page接口数据序列化失败",
				zap.Error(err),
				zap.String("operation", "marshal_data"),
			)
			return errors.New("序列化数据失败")
		}
		err = json.Unmarshal(jsonData, &list)
		if err != nil {
			log.Error("Page接口数据反序列化失败",
				zap.Error(err),
				zap.String("operation", "unmarshal_data"),
			)
			return errors.New("反序列化数据失败")
		}
	}

	// 获取总数
	countStart := time.Now()
	var total int64
	dbProvider.GetDb().Model(poBean).Count(&total)
	countDuration := time.Since(countStart)

	// 处理动态代理数据 todo 后续优化点,不需要顺序处理,go协程处理多任务
	processStart := time.Now()
	var pageFinalRes = make([]map[string]interface{}, 0)
	for _, m := range list {
		pageFinalRes = append(pageFinalRes, table(amisHeader, registerObj, m))
	}
	processDuration := time.Since(processStart)

	log.Info("Page接口查询完成",
		zap.String("bean", amisHeader.Bean),
		zap.Int64("total", total),
		zap.Int("rows", len(list)),
		zap.Duration("query_duration", queryDuration),
		zap.Duration("count_duration", countDuration),
		zap.Duration("process_duration", processDuration),
		zap.Duration("total_duration", time.Since(start)),
	)

	res := resp.PageResp{Total: total, Rows: pageFinalRes}
	return resp.Success(c, res)
}

func View(c echo.Context) error {
	start := time.Now()
	viewReq := new(req.ViewReq)
	if c.Bind(viewReq) != nil {
		log.Warn("View接口参数绑定失败", zap.String("error", "invalid_params"))
		return errors.New("非法参数")
	}

	amisHeader := getAmisHeader(c.Request().Header)
	registerObj := util.NewStructFromName(amisHeader.Bean)
	poBean := registerObj.Po

	log.Debug("View接口查询开始",
		zap.String("bean", amisHeader.Bean),
		zap.Uint("id", viewReq.Id),
	)

	queryStart := time.Now()
	res := dbProvider.GetDb().First(poBean, viewReq.Id)
	queryDuration := time.Since(queryStart)

	if res.RowsAffected == 0 {
		log.Warn("View接口记录不存在",
			zap.Uint("id", viewReq.Id),
			zap.String("bean", amisHeader.Bean),
		)
		return errors.New("不能存在记录")
	}

	beanTableMap := table(amisHeader, registerObj, poBean)
	log.Info("View接口查询完成",
		zap.String("bean", amisHeader.Bean),
		zap.Uint("id", viewReq.Id),
		zap.Duration("duration", time.Since(start)),
		zap.Duration("query_duration", queryDuration),
	)
	return resp.Success(c, beanTableMap)
}

func Create(c echo.Context) error {
	start := time.Now()
	var reqBody map[string]interface{}
	if c.Bind(&reqBody) != nil {
		log.Error("Create接口参数绑定失败", zap.String("error", "bind_failed"))
		panic("无法绑定请求参数")
	}

	_, registerObj := findBean(c, reqBody)
	if registerObj.Proxy != nil {
		proxy := registerObj.Proxy.(amis_proxy.IAmisProxy)
		proxy.BeforeAdd(reqBody)
		// 修改后重新序列化
		_, registerObj = findBean(c, reqBody)
	}

	poBean := registerObj.Po

	log.Debug("Create接口开始",
		zap.String("bean_type", fmt.Sprintf("%T", poBean)),
		zap.Any("request_body", reqBody),
	)

	// 通过反射调用BeforeAdd方法
	var res *gorm.DB
	createStart := time.Now()
	if registerObj.Proxy != nil {
		proxy := registerObj.Proxy.(amis_proxy.IAmisProxy)
		res = dbProvider.GetDb().Create(poBean)
		proxy.AfterAdd(poBean)
	} else {
		res = dbProvider.GetDb().Create(poBean)
	}
	createDuration := time.Since(createStart)

	if res.RowsAffected == 0 {
		log.Warn("Create接口新增失败", zap.String("error", "create_failed"))
		return errors.New("新增失败")
	}

	idValue := reflect.ValueOf(poBean).Elem().FieldByName("ID")
	var newID uint
	if idValue.IsValid() && idValue.CanUint() {
		newID = uint(idValue.Uint())
		log.Info("Create接口完成",
			zap.Uint("id", newID),
			zap.Int64("rows_affected", res.RowsAffected),
			zap.Duration("duration", time.Since(start)),
			zap.Duration("create_duration", createDuration),
		)
	} else {
		log.Info("Create接口完成",
			zap.Int64("rows_affected", res.RowsAffected),
			zap.Duration("duration", time.Since(start)),
			zap.Duration("create_duration", createDuration),
		)
	}
	return resp.Success(c, true)
}

func Update(c echo.Context) error {
	start := time.Now()
	var reqBody map[string]interface{}

	if c.Bind(&reqBody) != nil {
		log.Error("Update接口参数绑定失败", zap.String("error", "bind_failed"))
		panic("无法绑定请求参数")
	}

	_, registerObj := findBean(c, reqBody)
	if registerObj.Proxy != nil {
		proxy := registerObj.Proxy.(amis_proxy.IAmisProxy)
		proxy.BeforeUpdate(reqBody)
		// 修改后重新序列化
		_, registerObj = findBean(c, reqBody)
	}

	poBean := registerObj.Po

	log.Debug("Update接口开始",
		zap.String("bean_type", fmt.Sprintf("%T", poBean)),
		zap.Any("request_body", reqBody),
	)

	// 使用接口获取数据库连接
	updateStart := time.Now()
	var res *gorm.DB
	if registerObj.Proxy != nil {
		proxy := registerObj.Proxy.(amis_proxy.IAmisProxy)
		res = dbProvider.GetDb().Updates(poBean)
		proxy.AfterUpdate(poBean)
	} else {
		res = dbProvider.GetDb().Updates(poBean)
	}
	updateDuration := time.Since(updateStart)

	if res.RowsAffected == 0 {
		log.Warn("Update接口更新失败", zap.String("error", "update_failed"))
		return errors.New("更新失败")
	}

	log.Info("Update接口完成",
		zap.Int64("rows_affected", res.RowsAffected),
		zap.Duration("duration", time.Since(start)),
		zap.Duration("update_duration", updateDuration),
	)
	return resp.Success(c, true)
}

func DeleteBatch(c echo.Context) error {
	start := time.Now()
	deleteBatchReq := new(req.DeleteBatchReq)
	if c.Bind(deleteBatchReq) != nil {
		log.Warn("DeleteBatch接口参数绑定失败", zap.String("error", "invalid_params"))
		return errors.New("非法参数")
	}
	amisHeader := getAmisHeader(c.Request().Header)
	registerObj := util.NewStructFromName(amisHeader.Bean)
	poBean := registerObj.Po

	// 使用接口获取数据库连接
	//deleteBatchReq.Ids
	idsStr := convertor.ToString(deleteBatchReq.Ids)
	idsSplit := strings.Split(idsStr, ",")
	ids, _ := commons.StringsToUints(idsSplit)

	log.Debug("DeleteBatch接口开始",
		zap.String("bean", amisHeader.Bean),
		zap.Uints("ids", ids),
		zap.String("raw_ids", idsStr),
	)

	deleteStart := time.Now()
	var tx *gorm.DB
	if registerObj.Proxy != nil {
		proxy := registerObj.Proxy.(amis_proxy.IAmisProxy)
		proxy.BeforeDelete(ids)
		tx = dbProvider.GetDb().Delete(poBean, ids)
		proxy.AfterDelete(ids)
	} else {
		tx = dbProvider.GetDb().Delete(poBean, ids)
	}
	deleteDuration := time.Since(deleteStart)

	if tx.RowsAffected == 0 {
		log.Warn("DeleteBatch接口删除失败",
			zap.String("bean", amisHeader.Bean),
			zap.Uints("ids", ids),
			zap.String("error", "delete_failed"),
		)
		return errors.New("删除失败")
	}

	log.Info("DeleteBatch接口完成",
		zap.String("bean", amisHeader.Bean),
		zap.Uints("ids", ids),
		zap.Int64("deleted_count", tx.RowsAffected),
		zap.Duration("duration", time.Since(start)),
		zap.Duration("delete_duration", deleteDuration),
	)
	return resp.Success(c, true)
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

func findBean(c echo.Context, reqBody map[string]interface{}) (map[string]interface{}, util.RegisterResp) {
	amisHeader := getAmisHeader(c.Request().Header)
	byStruct := changeMapByStruct(amisHeader, reqBody)
	marshal, _ := json.Marshal(byStruct)
	bean := util.NewStructFromJSONAndName(amisHeader.Bean, marshal)

	log.Debug("findBean处理完成",
		zap.String("bean", amisHeader.Bean),
		zap.String("site", amisHeader.Site),
		zap.Int("request_body_size", len(reqBody)),
	)

	return reqBody, bean
}

func changeMapByStruct(header req.AmisHeader, body map[string]interface{}) map[string]interface{} {
	var bodyNew = make(map[string]interface{})
	var beanPre = header.Bean + constant.AmisSplitSymbol

	log.Debug("changeMapByStruct处理",
		zap.String("bean", header.Bean),
		zap.String("prefix", beanPre),
		zap.Int("original_fields", len(body)),
	)

	for s, i := range body {
		key := strings.TrimPrefix(s, beanPre)
		if strings.EqualFold(key, "createTime") || strings.EqualFold(key, "updateTime") {
			continue
		}
		bodyNew[key] = i
	}

	log.Debug("changeMapByStruct完成",
		zap.String("bean", header.Bean),
		zap.Int("processed_fields", len(bodyNew)),
	)

	return bodyNew
}

func getAmisHeader(header http.Header) req.AmisHeader {
	bean := header.Get("bean")
	site := header.Get("site")
	log.Debug("获取AmisHeader",
		zap.String("bean", bean),
		zap.String("site", site),
	)
	return req.AmisHeader{Bean: bean, Site: site}
}
