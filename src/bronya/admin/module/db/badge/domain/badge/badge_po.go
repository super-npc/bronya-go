package badge

import (
	"time"
)

// Badge 徽章事件
type Badge struct {
	_ struct{} `module:"系统" group:"数据管理" menu:"菜单" comment:"徽章事件"`

	Id uint `json:"id" gorm:"primaryKey"` // id主键

	ReadStatus ReadStatus `json:"readStatus"` // 状态

	Bean string `json:"bean"` // bean

	PrimaryKey uint `json:"primaryKey"` // 主键

	Reason *string `json:"reason"` // 原因

	CreateTime time.Time `json:"createTime"` // 创建时间

}

// BadgeExt 徽章事件 拓展
type BadgeExt struct {
}
