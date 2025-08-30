package user_po

import (
	"fmt"
	"reflect"
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
			fmt.Println("Struct meta:", tag)
		}
	}
}
