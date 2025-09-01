package sys_env_obj_field

type SysEnvObjFieldProxy struct {
}

func (this *SysEnvObjFieldProxy) BeforeAdd(reqBody map[string]interface{}) {

}

func (this *SysEnvObjFieldProxy) AfterAdd(poBean interface{}) {

}

func (this *SysEnvObjFieldProxy) BeforeUpdate(reqBody map[string]interface{}) {

}
func (this *SysEnvObjFieldProxy) AfterUpdate(poBean interface{}) {

}

func (this *SysEnvObjFieldProxy) BeforeDelete(ids []uint) {

}

func (this *SysEnvObjFieldProxy) AfterDelete(ids []uint) {

}

func (this *SysEnvObjFieldProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	// ext := DemoExt{}
	// toMap := util.StructToMap(ext)
	// for k, v := range toMap {
	// 		resTable[tableFieldPre+k] = v
	// }
}
