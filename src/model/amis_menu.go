package model

import "reflect"

type AmisMenu struct {
	ModulePath   string // 框架模块基础路径,例如: github.com/super-npc/bronya-go
	ModuleMenu   ModuleMenu
	One2ManyTags map[string]One2ManyTag
	IsFramework  bool // 是否为框架所属的bean
}

type One2ManyTag struct {
	Field_            reflect.StructField
	Bind1vNBean       string
	Bind1vNValueField string
	Bind1vNLabelField string
}

type ModuleMenu struct {
	Field_  reflect.StructField
	Module  string
	Group   string
	Menu    string
	Comment string
}
