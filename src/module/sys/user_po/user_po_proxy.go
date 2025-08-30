package user_po

import (
	"github.com/labstack/gommon/log"
	"github.com/super-npc/bronya-go/src/commons/util"
)

type UserPoProxy struct {
}

func (this *UserPoProxy) BeforeAdd(reqBody map[string]interface{}) {
	log.Infof("新增参数: %s", reqBody)

}

func (this *UserPoProxy) AfterAdd(PoBean interface{}) {
	log.Infof("新增后: %s", PoBean)
}

func (this *UserPoProxy) BeforeUpdate(reqBody map[string]interface{}) {
	log.Infof("修改参数: %s", reqBody)
}
func (this *UserPoProxy) AfterUpdate(PoBean interface{}) {
	log.Infof("修改后: %s", PoBean)
}

func (this *UserPoProxy) BeforeDelete(ids []uint) {
	log.Infof("删除前: %d", ids)
}

func (this *UserPoProxy) AfterDelete(ids []uint) {
	log.Infof("删除后: %d", ids)
}

func (this *UserPoProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	ext := UserPoExt{ExtName: "测试拓展字段"}
	toMap := util.StructToMap(ext)
	for k, v := range toMap {
		resTable[tableFieldPre+k] = v
	}
}
