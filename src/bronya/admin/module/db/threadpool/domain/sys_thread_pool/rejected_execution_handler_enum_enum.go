package sys_thread_pool

type RejectedExecutionHandlerEnum string

const (
	AbortPolicy RejectedExecutionHandlerEnum = "AbortPolicy" // 中止策略

	DiscardPolicy RejectedExecutionHandlerEnum = "DiscardPolicy" // 丢弃策略

	DiscardOldestPolicy RejectedExecutionHandlerEnum = "DiscardOldestPolicy" // 丢弃最旧任务

	CallerRunsPolicy RejectedExecutionHandlerEnum = "CallerRunsPolicy" // 调用者运行策略

)
