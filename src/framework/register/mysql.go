package register

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MyDb *gorm.DB

func Conn() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:RrCkEBbmlyktSXhuELo2Fa4SIA3ktKdA@tcp(139.199.207.29:33068)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	MyDb = db
}
