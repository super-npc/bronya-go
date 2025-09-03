package build_sql

import (
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/super-npc/bronya-go/src/module/amis/controller/req"
)

// GetOne2ManySql
// SELECT distinct hobby.*, (SELECT name FROM `student` WHERE  id = hobby.student_id ) as student_id_desc FROM `hobby` AS `hobby` WHERE  hobby.student_id = 2
func GetOne2ManySql(poBeanStr string, req *req.PageReq) string {
	if strings.EqualFold(req.One2ManyReq.Entity, "") {
		// 没有1:n 不需要拼接
		return ""
	}
	//if !util.ExistRegisterBean(req.One2ManyReq.Entity) { // todo 测试通过要放开
	//	panic(req.One2ManyReq.Entity + " 未注册")
	//}
	//if !util.ExistRegisterBean(poBeanStr) { // todo 测试通过要放开
	//	panic(poBeanStr + " 未注册")
	//}
	poTable := getPoBeanTable(poBeanStr)
	refSql := getRefSql(poBeanStr, req)
	refFieldSnakeCase := getRefFieldSnakeCase(req)
	refValStr := GetOne2ManyRefValStr(req)
	// 生成 SELECT distinct hobby.*, FROM `hobby` AS `hobby` WHERE  hobby.student_id = 2

	poCol := "distinct " + poTable + ".*"
	sql, _, err := squirrel.Select(poCol, refSql).From(poTable + " AS " + poTable).Where(squirrel.Eq{poTable + "." + refFieldSnakeCase: refValStr}).ToSql()
	if err != nil {
		panic(err)
	}
	//log.Info("参数", zap.String("args",args))
	return sql
	//return "SELECT distinct " + poTable + ".*," + refSql + " FROM `" + poTable + "` AS `" + poTable + "` WHERE  " + poTable + "." + refFieldSnakeCase + " = " + refValStr
}

// @BindMany2One(entity = Student.class, valueField = Student.Fields.id, labelField = Student.Fields.name)
// 生成 (SELECT name FROM `student` WHERE  id = hobby.student_id ) as student_id_desc
func getRefSql(poBeanStr string, req *req.PageReq) string {
	refBeanStr := req.One2ManyReq.Entity
	//refBean := util.NewStructFromName(refBeanStr)
	var labelField = "name" // todo 从 tag注入,对应labelField
	var refBeanTable = strutil.SnakeCase(refBeanStr)
	refFieldSnakeCase := getRefFieldSnakeCase(req)
	idVal := GetOne2ManyRefIdVal(poBeanStr, req)
	sql, _, err := squirrel.Select(labelField).
		From(refBeanTable).
		Where(squirrel.Eq{"id": idVal}).
		ToSql()
	if err != nil {
		panic(err)
	}
	return "(" + sql + ") as " + refFieldSnakeCase + "_desc"
	//return "(SELECT " + labelField + " FROM `" + refBeanTable + "` WHERE  id = " + poTable + "." + refFieldSnakeCase + ") as " + refFieldSnakeCase + "_desc"
}

func GetOne2ManyRefIdVal(poBeanStr string, req *req.PageReq) string {
	refFieldSnakeCase := getRefFieldSnakeCase(req)
	poTable := getPoBeanTable(poBeanStr)
	return poTable + "." + refFieldSnakeCase
}

func getPoBeanTable(poBeanStr string) string {
	return strutil.SnakeCase(poBeanStr)
}

func getRefFieldSnakeCase(req *req.PageReq) string {
	refField := req.One2ManyReq.EntityField // todo 需要检查该字段是否存在,防止注入风险
	return strutil.SnakeCase(refField)      // student_id
}

func GetOne2ManyRefValStr(req *req.PageReq) string {
	refVal := req.One2ManyReq.EntityFieldVal
	return convertor.ToString(refVal)
}
