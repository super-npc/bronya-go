package sys_thread_pool

type ProcessStatus string

const (
	CREATED ProcessStatus = "CREATED" // 已创建

	PENDING ProcessStatus = "PENDING" // 队列中

	RUNNING ProcessStatus = "RUNNING" // 运行中

	PAUSED ProcessStatus = "PAUSED" // 已暂停

	SUCCESS ProcessStatus = "SUCCESS" // 完成

	FAIL ProcessStatus = "FAIL" // 失败

	CANCELED ProcessStatus = "CANCELED" // 已取消

)
