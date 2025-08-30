package amis_proxy

type IAmisProxy interface {
	BeforeAdd()
	Table(record interface{})
}
