package sys_env_obj_field

type EnvStatus string

const (
	SYNC EnvStatus = "SYNC" // 同步

	NON_EXIST EnvStatus = "NON_EXIST" // 不存在

	DIFF EnvStatus = "DIFF" // 差异

)
