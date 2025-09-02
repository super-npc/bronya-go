package sys_thread_pool

type SysThreadPoolProxy struct {
}

func (this *SysThreadPoolProxy) BeforeAdd(reqBody map[string]interface{}) {

}

func (this *SysThreadPoolProxy) AfterAdd(poBean interface{}) {

}

func (this *SysThreadPoolProxy) BeforeUpdate(reqBody map[string]interface{}) {

}
func (this *SysThreadPoolProxy) AfterUpdate(poBean interface{}) {

}

func (this *SysThreadPoolProxy) BeforeDelete(ids []uint) {

}

func (this *SysThreadPoolProxy) AfterDelete(ids []uint) {

}

func (this *SysThreadPoolProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	// ext := DemoExt{}
	// toMap := util.StructToMap(ext)
	// for k, v := range toMap {
	// 		resTable[tableFieldPre+k] = v
	// }
}
