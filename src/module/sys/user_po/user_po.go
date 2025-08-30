package user_po

import (
	"fmt"
	"reflect"
)

// UserPo 表
type UserPo struct {
	//_      struct{} `meta:"table=user;desc=用户表"` // 不会被导出，但反射能拿到
	Id     uint `gorm:"primaryKey"`
	Name   string
	Status UserStatus
}

type UserPoExt struct {
}

func temp() {
	t := reflect.TypeOf(UserPo{})
	for i := 0; i < t.NumField(); i++ {
		if tag := t.Field(i).Tag.Get("meta"); tag != "" {
			fmt.Println("Struct meta:", tag)
		}
	}
}
