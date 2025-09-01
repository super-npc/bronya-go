package sys_env_obj_field

import (
	"database/sql"
)

// SysEnvObjField 变量属性
type SysEnvObjField struct {
	_ struct{} `module:"系统" group:"数据管理" menu:"系统变量" comment:"变量属性"`

	Id uint `json:"id" gorm:"primaryKey"` // id主键

	DataKey string `json:"dataKey"` // 键

	DataValue string `json:"dataValue"` // 值

	Description *string `json:"description"` // 描述

	EnvObjId uint `json:"envObjId"` // 对象

	CreateBy *uint `json:"createBy"` // 创建人

	UpdateBy *uint `json:"updateBy"` // 最后更新人

	CreateTime sql.NullTime `json:"createTime"` // 创建时间

	UpdateTime sql.NullTime `json:"updateTime"` // 最后更新时间

}

// EnvPropertyExt 变量属性 拓展
type EnvPropertyExt struct {
	Status EnvStatus `json:"status"` // 状态

}
