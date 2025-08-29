package user_po

import "fmt"

// UserPo 表
type UserPo struct {
	Name   string
	Status UserStatus // 状态,需要换成一个枚举,因为字符串的情况下没有约定只能哪些字符串,导致开发人员可以随意填写,但是具体值只有: enable,disable
}

func main() {
	po := UserPo{Name: "dfd", Status: Enable}
	fmt.Println(po)
}
