package user_po

import (
	"github.com/super-npc/bronya-go/src/commons/util"
)

type UserPoProxy struct {
}

func (this *UserPoProxy) BeforeAdd(reqBody map[string]interface{}) {

}

func (this *UserPoProxy) AfterAdd(poBean interface{}) {

}

func (this *UserPoProxy) BeforeUpdate(reqBody map[string]interface{}) {

}
func (this *UserPoProxy) AfterUpdate(poBean interface{}) {

}

func (this *UserPoProxy) BeforeDelete(ids []uint) {

}

func (this *UserPoProxy) AfterDelete(ids []uint) {

}

func (this *UserPoProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	ext := UserPoExt{}
	toMap := util.StructToMap(ext)
	for k, v := range toMap {
		resTable[tableFieldPre+k] = v
	}
}
