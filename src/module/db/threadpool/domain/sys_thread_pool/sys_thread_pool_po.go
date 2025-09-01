package sys_thread_pool

// SysThreadPool 线程池
type SysThreadPool struct {
	_ struct{} `module:"系统" group:"配置" menu:"系统配置"`

	Id uint `json:"id" gorm:"primaryKey"` // id主键

	PrefixName string `json:"prefixName"` // 线程前缀

	CorePoolSize uint8 `json:"corePoolSize"` // 核心线程

	MaximumPoolSize uint8 `json:"maximumPoolSize"` // 最大线程

	RejectedStrategy RejectedExecutionHandlerEnum `json:"rejectedStrategy"` // 拒绝策略

}

// SysThreadPoolExt
type SysThreadPoolExt struct {
	Status ActivationStatus `json:"status"` // 状态

	Process ProcessStatus `json:"process"` // 流程

	Running *string `json:"running"` // 运行信息

}
