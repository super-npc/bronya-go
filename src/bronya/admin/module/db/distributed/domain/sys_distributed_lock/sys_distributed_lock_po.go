package sys_distributed_lock

import (
	"time"
)

// SysDistributedLock 分布式锁
type SysDistributedLock struct {
	_ struct{} `module:"系统" group:"配置" menu:"系统配置" comment:"分布式锁"`

	Id uint `json:"id" gorm:"primaryKey"` // id主键

	LockKey string `json:"lockKey"` // 名称

	CreateTime time.Time `json:"createTime"` // 创建时间

}
