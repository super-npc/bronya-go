package user_po

// UserPo 表
type UserPo struct {
	Id     uint `gorm:"primaryKey"`
	Name   string
	Status UserStatus
}
