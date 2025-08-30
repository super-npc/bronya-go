package user_po

import (
	"github.com/labstack/gommon/log"
	"github.com/super-npc/bronya-go/src/commons/util"
)

type UserPoProxy struct {
}

func (this *UserPoProxy) BeforeAdd(reqBody map[string]interface{}) {
	log.Infof("新增前: %s", reqBody)
}

func (this *UserPoProxy) AfterAdd(PoBean interface{}) {
	log.Infof("新增后: %s", PoBean)
}

func (this *UserPoProxy) BeforeUpdate() {

}
func (this *UserPoProxy) AfterUpdate() {

}

func (this *UserPoProxy) BeforeDelete() {

}

func (this *UserPoProxy) AfterDelete() {

}

func (this *UserPoProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	ext := UserPoExt{ExtName: "测试拓展字段"}
	toMap := util.StructToMap(ext)
	for k, v := range toMap {
		resTable[tableFieldPre+k] = v
	}
}
