package service

import (
	"fmt"
	"strings"

	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/module/amis/controller/req"
)

// One2Many
// SELECT distinct hobby.*, (SELECT name FROM `student` WHERE  id = hobby.student_id ) as student_id_desc FROM `hobby` AS `hobby` WHERE  hobby.student_id = 2
func One2Many(poBeanStr string, req *req.PageReq) {
	fmt.Println(req)
	if strings.EqualFold(req.One2ManyReq.Entity, "") {
		// 没有1:n 不需要拼接
		return
	}
	refBeanStr := req.One2ManyReq.Entity
	refField := req.One2ManyReq.EntityField
	refVal := req.One2ManyReq.EntityFieldVal

	poBean := util.NewStructFromName(poBeanStr)
	refBean := util.NewStructFromName(refBeanStr)
}
