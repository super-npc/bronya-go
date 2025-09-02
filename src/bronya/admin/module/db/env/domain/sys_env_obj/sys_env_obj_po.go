package sys_env_obj

import (
	"time"
)

// SysEnvObj 变量对象
type SysEnvObj struct {
	_ struct{} `module:"系统" group:"数据管理" menu:"系统变量" comment:"变量对象"`

	Id uint `json:"id" gorm:"primaryKey"` // id主键

	ObjName string `json:"objName"` // 对象名

	Description *string `json:"description"` // 描述

	CreateTime time.Time `json:"createTime"` // 创建时间

}

// SysEnvObjExt 变量对象 拓展
type SysEnvObjExt struct {
	Env *string `json:"env"` // 环境变量

}
