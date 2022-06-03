package jsonparsing

import (
	"encoding/json"
	"fmt"
)

func doMarshalTest1() {
	data := make(map[string]interface{})

	data["name"] = "kim"
	data["age"] = 31
	valuedata := make(map[string]interface{})
	valuedata["key"] = "abcdefg-zxvasdfa"
	data["value"] = valuedata

	doc, _ := json.Marshal(data)

	fmt.Println(string(doc))

	doc2, _ := json.MarshalIndent(data, "", " ")

	fmt.Println(string(doc2))

	doc3 := `
	{
	"name": "kim",
	"age": 31
	}
			`
	var data3 map[string]interface{}
	json.Unmarshal([]byte(doc3), &data3)
	fmt.Println(data3)

	doc4, _ := json.MarshalIndent(data3, "", " ")
	fmt.Println(string(doc4))
}

func DoUnMarshalJsonString() {

	stringKey := `{"key": "{\"service\":\"TTSaaS\",\"id\":\"32c1e4d6-2bc0-4ad2-a89e-fde66363fed1\"}"}`
	var stringKeyObj map[string]interface{}
	err := json.Unmarshal([]byte(stringKey), &stringKeyObj)
	if err != nil {
		fmt.Println(err)
	}
	if val, ok := stringKeyObj["key"]; ok {

		if isString(val) {
			fmt.Println("reconstructing key field")
		}

	}
	objectKey := ` {"key": {
		"service": "DLGaaS",
		"id": "ca6e4281-47c1-4cc8-a3e2-48d0bcda29de"
	}}`
	var objectKeyObj map[string]interface{}
	err = json.Unmarshal([]byte(objectKey), &objectKeyObj)
	if err != nil {
		fmt.Println(err)
	}
	if val, ok := objectKeyObj["key"]; ok {
		if isString(val) {
			fmt.Println("reconstructing key field")
		} else {
			fmt.Println("keeping original keyf field")
		}
	}
}

func isString(val interface{}) bool {
	switch v := val.(type) {
	case string:
		fmt.Println("from onpremise", v)
		return true
	default:
		return false

	}
}
