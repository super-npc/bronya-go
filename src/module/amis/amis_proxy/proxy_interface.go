package amis_proxy

type IAmisProxy interface {
	BeforeAdd(reqBody map[string]interface{})
	AfterAdd(PoBean interface{})

	BeforeUpdate(reqBody map[string]interface{})
	AfterUpdate(PoBean interface{})

	BeforeDelete(ids []uint)
	AfterDelete(ids []uint)
	Table(tableFieldPre string, resTable map[string]interface{})
}
