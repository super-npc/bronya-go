package register

import (
	"github.com/acmestack/gorm-plus/gplus"
	"github.com/super-npc/bronya-go/src/commons/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:RrCkEBbmlyktSXhuELo2Fa4SIA3ktKdA@tcp(139.199.207.29:33068)/bronya-demo-one?charset=utf8mb4&parseTime=True&loc=Local"

	mysqlDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "app_", // 表前缀
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	MyDb = mysqlDb

	// 初始化gplus
	gplus.Init(mysqlDb)
}
