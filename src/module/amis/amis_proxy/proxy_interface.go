package amis_proxy

type IAmisProxy interface {
	BeforeAdd()
	AfterAdd()
	BeforeUpdate()
	AfterUpdate()

	BeforeDelete()
	AfterDelete()
	Table(tableFieldPre string, resTable map[string]interface{})
}
