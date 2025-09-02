package sys_thread_pool

type ActivationStatus string

const (
	ENABLE ActivationStatus = "ENABLE" // 生效

	DISABLE ActivationStatus = "DISABLE" // 失效

)
