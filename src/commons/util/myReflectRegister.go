package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 全局注册表：名字 -> 类型(Type)
var typeRegistry = make(map[string]reflect.Type)

//func init() {
//	// 把将来可能用到的结构体注册进来；key 就是字符串名字
//	//registerByName("RefletDemo", reflect.TypeOf(RefletDemo{}))
//	//registerByStruct(&RefletDemo{})   // 或 RefletDemo{}
//}

func RegisterByName(name string, typ reflect.Type) {
	typeRegistry[name] = typ
}

// RegisterByStruct 注册一个结构体（传入指针或值都行）
func RegisterByStruct(v interface{}) {
	t := reflect.TypeOf(v)
	// 去掉指针层，拿到真正的结构体类型
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("registerType: need a struct or *struct")
	}
	// 用结构体自身的名字作为 key
	typeRegistry[t.Name()] = t
}

// 根据名字返回该类型的零值指针 (interface{})
func newByName(name string) (interface{}, error) {
	t, ok := typeRegistry[name]
	if !ok {
		return nil, fmt.Errorf("unknown type: %s", name)
	}
	return reflect.New(t).Interface(), nil
}

// NewStructFromJSONAndName 业务逻辑：传入 JSON 和类型名字符串，返回填充好的结构体对象
func NewStructFromJSONAndName(typeName string, jsonData []byte) (interface{}, error) {
	// 1. 从注册表拿到类型的 reflect.Type
	t, ok := typeRegistry[typeName]
	if !ok {
		return nil, fmt.Errorf("unknown type: %s", typeName)
	}

	// 2. 创建该类型的零值指针（*Struct）
	ptr := reflect.New(t)

	// 3. 把 JSON 填进去
	if err := json.Unmarshal(jsonData, ptr.Interface()); err != nil {
		return nil, err
	}

	// 4. 返回真正的结构体对象（去掉指针层）
	return ptr.Elem().Interface(), nil
}
