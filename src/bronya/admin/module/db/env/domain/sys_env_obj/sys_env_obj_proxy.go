package sys_env_obj

type SysEnvObjProxy struct {
}

func (this *SysEnvObjProxy) BeforeAdd(reqBody map[string]interface{}) {

}

func (this *SysEnvObjProxy) AfterAdd(poBean interface{}) {

}

func (this *SysEnvObjProxy) BeforeUpdate(reqBody map[string]interface{}) {

}
func (this *SysEnvObjProxy) AfterUpdate(poBean interface{}) {

}

func (this *SysEnvObjProxy) BeforeDelete(ids []uint) {

}

func (this *SysEnvObjProxy) AfterDelete(ids []uint) {

}

func (this *SysEnvObjProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	// ext := DemoExt{}
	// toMap := util.StructToMap(ext)
	// for k, v := range toMap {
	// 		resTable[tableFieldPre+k] = v
	// }
}
