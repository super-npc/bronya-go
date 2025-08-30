package user_po

import "fmt"

type UserPoProxy struct {
}

func (this *UserPoProxy) BeforeAdd() {
	fmt.Println("调用代理了...")
}
