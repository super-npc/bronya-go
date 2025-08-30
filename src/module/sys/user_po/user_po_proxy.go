package user_po

import (
	"fmt"

	"github.com/super-npc/bronya-go/src/commons/util"
)

type UserPoProxy struct {
}

func (this *UserPoProxy) BeforeAdd() {
	fmt.Println("调用代理了...")
}

func (this *UserPoProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	ext := UserPoExt{ExtName: "测试拓展字段"}
	toMap := util.StructToMap(ext)
	for k, v := range toMap {
		resTable[tableFieldPre+k] = v
	}
}
