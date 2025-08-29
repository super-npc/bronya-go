package user_po

import (
	"github.com/labstack/gommon/log"
	"github.com/super-npc/bronya-go/src/framework/register"
	"gorm.io/gorm"
)

func Insert(user UserPo) *gorm.DB {
	res := register.MyDb.Create(user)
	log.Info("插入条数", res.RowsAffected)
	return res
}
