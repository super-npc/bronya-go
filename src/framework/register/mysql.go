package register

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/super-npc/bronya-go/src/commons/db"
	"github.com/super-npc/bronya-go/src/framework/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MyDb *gorm.DB

// 实现 DBProvider 接口
type dbProviderImpl struct{}

func (d *dbProviderImpl) GetDb() *gorm.DB {
	return MyDb
}

// GetDbProvider 返回数据库提供者实例
func GetDbProvider() db.DBProvider {
	return &dbProviderImpl{}
}

func InitDatabase() {
	config := conf.Settings.MySQLConfig
	host := config.Host
	port := config.Port
	user := config.User
	password := config.Password
	dbName := config.DbName
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	mysqlDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "app_", // 表前缀
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出
			logger.Config{
				SlowThreshold: time.Second, // 慢查询阈值
				LogLevel:      logger.Info, // 日志级别：Silent、Error、Warn、Info
				Colorful:      true,        // 彩色输出
			},
		),
	})
	if err != nil {
		panic(err)
	}
	MyDb = mysqlDb

	// 初始化gplus
	gplus.Init(mysqlDb)
}
