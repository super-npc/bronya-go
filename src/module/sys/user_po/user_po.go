package user_po

// UserPo 表
type UserPo struct {
	_      struct{}   `module:"系统" group:"数据管理" menu:"系统变量"`
	Id     uint       `json:"id" gorm:"primaryKey"`
	Name   string     `json:"name"`
	Status UserStatus `json:"status"`
}

type UserPoExt struct {
	ExtName string `json:"extName"`
}
