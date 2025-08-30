package util

import (
	"encoding/json"
	"reflect"
)

// 全局注册表：名字 -> 类型(Type)
var typeRegistry = make(map[string]RegisterRefType)

//func init() {
//	// 把将来可能用到的结构体注册进来；key 就是字符串名字
//	//registerByName("RefletDemo", reflect.TypeOf(RefletDemo{}))
//	//registerByStruct(&RefletDemo{})   // 或 RefletDemo{}
//}

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

func Register(registerAmis RegisterReq) {
	if registerAmis.Po == nil {
		panic("未传递po类")
	}
	poType := registerByStruct(registerAmis.Po)
	resp := RegisterRefType{Po: poType}
	// 代理类可有可无
	if registerAmis.Proxy != nil {
		resp.Proxy = registerByStruct(registerAmis.Proxy)
	}
	// 拓展类,可有可无
	if registerAmis.PoExt != nil {
		resp.PoExt = registerByStruct(registerAmis.PoExt)
	}
	typeRegistry[poType.Name()] = resp
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
