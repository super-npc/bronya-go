package amis_proxy

type IAmisProxy interface {
	BeforeAdd()
	Table(tableFieldPre string, resTable map[string]interface{})
}
