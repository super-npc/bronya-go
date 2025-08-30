package util

import (
	"encoding/json"
	"reflect"
)

// 全局注册表：名字 -> 类型(Type)
var typeRegistry = make(map[string]RegisterResp)

//func init() {
//	// 把将来可能用到的结构体注册进来；key 就是字符串名字
//	//registerByName("RefletDemo", reflect.TypeOf(RefletDemo{}))
//	//registerByStruct(&RefletDemo{})   // 或 RefletDemo{}
//}

type RegisterReq struct {
	Po    interface{}
	Proxy interface{}
}

type RegisterResp struct {
	Po    reflect.Type
	Proxy reflect.Type
}

type RegisterObj struct {
	Po    interface{}
	Proxy interface{}
}

func Register(registerAmis RegisterReq) {
	if registerAmis.Po == nil {
		panic("未传递po类")
	}
	poType := RegisterByStruct(registerAmis.Po)
	resp := RegisterResp{Po: poType}
	// 代理类可有可无
	if registerAmis.Proxy == nil {
		proxyType := RegisterByStruct(registerAmis.Proxy)
		resp.Proxy = proxyType
	}
	typeRegistry[poType.Name()] = resp
}

// RegisterByStruct 注册一个结构体（传入指针或值都行）
func RegisterByStruct(v interface{}) reflect.Type {
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

func NewStructFromName(typeName string) (RegisterObj, error) {
	// 1. 从注册表拿到类型的 reflect.Type
	t, ok := typeRegistry[typeName]
	if !ok {
		//return nil, fmt.Errorf("unknown type: %s", typeName)
		panic("无法注册:" + typeName)
	}

	// 2. 创建该类型的零值指针（*Struct）
	obj := RegisterObj{Po: reflect.New(t.Po).Interface()}
	if t.Proxy != nil {
		obj.Proxy = reflect.New(t.Proxy).Interface()
	}

	// 4. 返回指针而不是值，以便GORM可以使用反射
	return obj, nil
}

// NewStructFromJSONAndName 业务逻辑：传入 JSON 和类型名字符串，返回填充好的结构体对象
func NewStructFromJSONAndName(typeName string, jsonData []byte) RegisterObj {
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
	obj := RegisterObj{Po: ptr.Interface()}
	if t.Proxy != nil {
		obj.Proxy = reflect.New(t.Proxy).Interface()
	}

	// 4. 返回指针而不是值，以便GORM可以使用反射
	return obj
}
