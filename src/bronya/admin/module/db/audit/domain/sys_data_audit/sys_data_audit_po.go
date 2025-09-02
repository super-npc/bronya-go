package sys_data_audit

import (
	"time"
)

// SysDataAudit 数据审计
type SysDataAudit struct {
	_ struct{} `module:"系统" group:"数据管理" menu:"审计管理" comment:"数据审计"`

	Id uint `json:"id" gorm:"primaryKey"` // id主键

	TableBean string `json:"tableBean"` // 表

	RecordPrimaryKey string `json:"recordPrimaryKey"` // 主键

	OldData string `json:"oldData"` // 修改前

	NewData string `json:"newData"` // 修改后

	UpdateBy uint `json:"updateBy"` // 更新人

	UpdateTime time.Time `json:"updateTime"` // 更新时间

}

// SysDataAuditExt 数据审计 拓展
type SysDataAuditExt struct {
	Diff *string `json:"diff"` // 差异

}
