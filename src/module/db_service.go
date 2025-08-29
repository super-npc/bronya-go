package module

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/super-npc/bronya-go/src/framework/register"
	"gorm.io/gorm"
)

func Create[T any](obj *T) error {
	ctx := context.Background()
	if err := gorm.G[T](register.MyDb).Create(ctx, obj); err != nil {
		log.Errorf("create failed: %v", err)
		return err
	}
	return nil
}
