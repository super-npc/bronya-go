package service

import (
	"context"
	"fmt"

	"github.com/super-npc/bronya-go/src/framework/register"
	"github.com/super-npc/bronya-go/src/module/sys/domain"
	"gorm.io/gorm"
)

func CreateUserPo() {
	user := domain.UserPo{Name: "Jinzhu"}
	ctx := context.Background()
	err := gorm.G[domain.UserPo](register.MyDb).Create(ctx, &user) // pass pointer of data to Create
	if err != nil {
		panic(err)
	}
}

func CreateUserPo2() {
	user := domain.UserPo{Name: "Jinzhu"}
	res := register.MyDb.Create(user)
	fmt.Printf("插入条数: %d", res.RowsAffected)
}
