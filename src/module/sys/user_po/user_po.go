package user_po

import (
	"reflect"

	"github.com/super-npc/bronya-go/src/framework/logger"
	"go.uber.org/zap"
)

// UserPo 表
type UserPo struct {
	//_      struct{} `meta:"table=user;desc=用户表"` // 不会被导出，但反射能拿到
	Id     uint       `json:"id" gorm:"primaryKey"`
	Name   string     `json:"name"`
	Status UserStatus `json:"status"`
}

type UserPoExt struct {
	ExtName string `json:"extName"`
}

func temp() {
	t := reflect.TypeOf(UserPo{})
	for i := 0; i < t.NumField(); i++ {
		if tag := t.Field(i).Tag.Get("meta"); tag != "" {
			logger.Info("Struct meta", zap.String("tag", tag))
		}
	}
}
