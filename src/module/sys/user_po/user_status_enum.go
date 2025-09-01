package user_po

// UserStatus 1. 定义枚举类型
type UserStatus string

// 2. 枚举值常量
const (
	Enable  UserStatus = "enable"  //
	Disable UserStatus = "disable" //
)
