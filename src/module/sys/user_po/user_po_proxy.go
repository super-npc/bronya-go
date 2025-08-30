package user_po

import "fmt"

type UserPoProxy struct {
}

func (this *UserPoProxy) BeforeAdd() {
	fmt.Println("调用代理了...")
}

func (this *UserPoProxy) Table(record interface{}) {
	fmt.Println("调用table")
}
