package amis_proxy

type IAmisProxy interface {
	BeforeAdd(reqBody map[string]interface{})
	AfterAdd(PoBean interface{})
	BeforeUpdate()
	AfterUpdate()

	BeforeDelete()
	AfterDelete()
	Table(tableFieldPre string, resTable map[string]interface{})
}
