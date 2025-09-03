package build_sql

import (
	"strings"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/module/amis/controller/req"
)

// GetOne2ManySql
// SELECT distinct hobby.*, (SELECT name FROM `student` WHERE  id = hobby.student_id ) as student_id_desc FROM `hobby` AS `hobby` WHERE  hobby.student_id = 2
func GetOne2ManySql(poBeanStr string, req *req.PageReq) string {
	if strings.EqualFold(req.One2ManyReq.Entity, "") {
		// 没有1:n 不需要拼接
		return ""
	}
	if !util.ExistRegisterBean(req.One2ManyReq.Entity) {
		panic(req.One2ManyReq.Entity + " 未注册")
	}
	if !util.ExistRegisterBean(poBeanStr) {
		panic(poBeanStr + " 未注册")
	}
	poTable := getPoBeanTable(poBeanStr)
	refSql := getRefSql(poBeanStr, req)
	refFieldSnakeCase := getRefFieldSnakeCase(req)
	refValStr := getRefValStr(req)
	// 生成 SELECT distinct hobby.*, FROM `hobby` AS `hobby` WHERE  hobby.student_id = 2
	return "SELECT distinct " + poTable + ".*," + refSql + " FROM `" + poTable + "` AS `" + poTable + "` WHERE  " + poTable + "." + refFieldSnakeCase + " = " + refValStr
}

// @BindMany2One(entity = Student.class, valueField = Student.Fields.id, labelField = Student.Fields.name)
// 生成 (SELECT name FROM `student` WHERE  id = hobby.student_id ) as student_id_desc
func getRefSql(poBeanStr string, req *req.PageReq) string {
	refBeanStr := req.One2ManyReq.Entity
	//refBean := util.NewStructFromName(refBeanStr)
	var labelField = "name" // 从 tag注入,对应labelField
	var refBeanTable = strutil.SnakeCase(refBeanStr)
	refFieldSnakeCase := getRefFieldSnakeCase(req)
	poTable := getPoBeanTable(poBeanStr)

	return "(SELECT " + labelField + " FROM `" + refBeanTable + "` WHERE  id = " + poTable + "." + refFieldSnakeCase + ") as " + refFieldSnakeCase + "_desc"
}

func getPoBeanTable(poBeanStr string) string {
	return strutil.SnakeCase(poBeanStr)
}

func getRefFieldSnakeCase(req *req.PageReq) string {
	refField := req.One2ManyReq.EntityField
	return strutil.SnakeCase(refField) // student_id
}

func getRefValStr(req *req.PageReq) string {
	refVal := req.One2ManyReq.EntityFieldVal
	return convertor.ToString(refVal)
}
