package register

import (
	"fmt"

	"github.com/go-redis/redis"
	setting "github.com/super-npc/bronya-go/src/framework/conf"
)

// 声明一个全局的client变量
var (
	client   *redis.Client
	redisNil = redis.Nil
)

// InitLogger Init 初始化连接
func InitLogger(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err = client.Ping().Result()
	return
}

func Close() {
	_ = client.Close()
}
