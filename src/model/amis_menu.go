package model

import "reflect"

type AmisMenu struct {
	ModulePath string // 框架模块基础路径,例如: github.com/super-npc/bronya-go
	Field_     reflect.StructField
	Menu       Menu
}

type Menu struct {
	Module  string
	Group   string
	Menu    string
	Comment string
}
