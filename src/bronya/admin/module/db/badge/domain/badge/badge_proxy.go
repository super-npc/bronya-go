package badge

type BadgeProxy struct {
}

func (this *BadgeProxy) BeforeAdd(reqBody map[string]interface{}) {

}

func (this *BadgeProxy) AfterAdd(poBean interface{}) {

}

func (this *BadgeProxy) BeforeUpdate(reqBody map[string]interface{}) {

}
func (this *BadgeProxy) AfterUpdate(poBean interface{}) {

}

func (this *BadgeProxy) BeforeDelete(ids []uint) {

}

func (this *BadgeProxy) AfterDelete(ids []uint) {

}

func (this *BadgeProxy) Table(tableFieldPre string, resTable map[string]interface{}) {
	// 处理拓展类
	// ext := DemoExt{}
	// toMap := util.StructToMap(ext)
	// for k, v := range toMap {
	// 		resTable[tableFieldPre+k] = v
	// }
}
