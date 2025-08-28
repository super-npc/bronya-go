package framework

import (
	"github.com/super-npc/bronya-go/src/commons/util"
	"github.com/super-npc/bronya-go/src/module/sys/domain"
)

func Start() {
	util.RegisterByStruct(&domain.User{})
}
