package sys_data_audit

type SysDataAuditProxy struct {
}

func (this *SysDataAuditProxy) BeforeAdd(reqBody map[string]interface{}) {

}

func (this *SysDataAuditProxy) AfterAdd(poBean interface{}) {

}

func (this *SysDataAuditProxy) BeforeUpdate(reqBody map[string]interface{}) {

}
func (this *SysDataAuditProxy) AfterUpdate(poBean interface{}) {

}

func (this *SysDataAuditProxy) BeforeDelete(ids []uint) {

}

func (this *SysDataAuditProxy) AfterDelete(ids []uint) {

}

func (this *SysDataAuditProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	// ext := DemoExt{}
	// toMap := util.StructToMap(ext)
	// for k, v := range toMap {
	// 		resTable[tableFieldPre+k] = v
	// }
}
