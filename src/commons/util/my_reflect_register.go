package util

import (
	"encoding/json"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/super-npc/bronya-go/src/commons/constant"
	"github.com/super-npc/bronya-go/src/framework/log"
	"github.com/super-npc/bronya-go/src/model"
	"go.uber.org/zap"
)

// 全局注册表：名字 -> 类型(Type)
var typeRegistry = make(map[string]RegisterRefType)
var amisMenus = make(map[string]model.AmisMenu)

type RegisterReq struct {
	Po    interface{}
	PoExt interface{}
	Proxy interface{}
}

type RegisterRefType struct {
	Po    reflect.Type
	PoExt reflect.Type
	Proxy reflect.Type
}

type RegisterResp struct {
	Po    interface{}
	PoExt interface{}
	Proxy interface{}
}

func GetAmisMenus() map[string]model.AmisMenu {
	return amisMenus
}

func RegisterProject(registerAmis RegisterReq) {
	register(false, registerAmis)
}

func RegisterFramework(registerAmis RegisterReq) {
	register(true, registerAmis)
}

func register(isFramework bool, registerAmis RegisterReq) {
	if registerAmis.Po == nil {
		panic("未传递po类")
	}
	poType := registerByStruct(registerAmis.Po)
	refType := typeRegistry[poType.Name()]
	if refType.Po != nil {
		panic("重复注册" + refType.Po.Name())
	}

	tags := getPoFieldTags(isFramework, poType)
	resp := RegisterRefType{Po: poType}
	// 代理类可有可无
	if registerAmis.Proxy != nil {
		resp.Proxy = registerByStruct(registerAmis.Proxy)
		// 记录所有的tag
	}
	// 拓展类,可有可无
	if registerAmis.PoExt != nil {
		resp.PoExt = registerByStruct(registerAmis.PoExt)
	}
	typeRegistry[poType.Name()] = resp
	amisMenus[poType.Name()] = tags
}

func getPoFieldTags(isFramework bool, poType reflect.Type) model.AmisMenu {
	var field_ reflect.StructField
	for i := range poType.NumField() {
		field := poType.Field(i)
		if !strings.EqualFold(field.Name, "_") {
			continue
		}
		// 所有构造字段,用于提取非普通字段的tag
		field_ = field
	}
	var tag = field_.Tag
	module := tag.Get("module")
	group := tag.Get("group")
	menu := tag.Get("menu")
	comment := tag.Get("comment")
	groupMenu := model.Menu{Module: module, Group: group, Menu: menu, Comment: comment}
	return model.AmisMenu{ModulePath: getModulePath(isFramework), Field_: field_, Menu: groupMenu}
}

func getModulePath(isFramework bool) string {
	if isFramework {
		return constant.FrameworkModule
	}
	info, _ := debug.ReadBuildInfo()
	return info.Path
}

// RegisterByStruct 注册一个结构体（传入指针或值都行）
func registerByStruct(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	// 去掉指针层，拿到真正的结构体类型
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("registerType: need a struct or *struct")
	}
	// 用结构体自身的名字作为 key
	//typeRegistry[t.Name()] = t
	log.Info("注册结构体", zap.String("struct", t.Name()))
	return t
}

func NewStructFromName(typeName string) RegisterResp {
	// 1. 从注册表拿到类型的 reflect.Type
	t, ok := typeRegistry[typeName]
	if !ok {
		//return nil, fmt.Errorf("unknown type: %s", typeName)
		panic("无法注册:" + typeName)
	}

	// 2. 创建该类型的零值指针（*Struct）
	obj := RegisterResp{Po: reflect.New(t.Po).Interface()}
	if t.Proxy != nil {
		obj.Proxy = reflect.New(t.Proxy).Interface()
	}
	if t.PoExt != nil {
		obj.PoExt = reflect.New(t.PoExt).Interface()
	}

	// 4. 返回指针而不是值，以便GORM可以使用反射
	return obj
}

// NewStructFromJSONAndName 业务逻辑：传入 JSON 和类型名字符串，返回填充好的结构体对象
func NewStructFromJSONAndName(typeName string, jsonData []byte) RegisterResp {
	// 1. 从注册表拿到类型的 reflect.Type
	t, ok := typeRegistry[typeName]
	if !ok {
		panic("无法注册:" + typeName)
	}

	// 2. 创建该类型的零值指针（*Struct）
	ptr := reflect.New(t.Po)

	// 3. 把 JSON 填进去
	if err := json.Unmarshal(jsonData, ptr.Interface()); err != nil {
		panic("json无法转成po实体:" + typeName)
	}

	// 2. 创建该类型的零值指针（*Struct）
	obj := RegisterResp{Po: ptr.Interface()}
	if t.Proxy != nil {
		obj.Proxy = reflect.New(t.Proxy).Interface()
	}
	if t.PoExt != nil {
		obj.PoExt = reflect.New(t.PoExt).Interface()
	}

	// 4. 返回指针而不是值，以便GORM可以使用反射
	return obj
}
