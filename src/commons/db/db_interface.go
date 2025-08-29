package db

import "gorm.io/gorm"

// DBProvider 定义数据库操作接口
type DBProvider interface {
	GetDb() *gorm.DB
}
