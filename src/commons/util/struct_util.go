package util

import "encoding/json"

func StructToMap(in interface{}) map[string]interface{} {
	var m map[string]interface{}
	data, _ := json.Marshal(in)
	err := json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	return m
}
