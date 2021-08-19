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

	doc, _ :=json.Marshal(data)

	fmt.Println(string(doc))

	doc2, _ :=json.MarshalIndent(data,""," ")

	fmt.Println(string(doc2))

	doc3 := `
	{
	"name": "kim",
	"age": 31
	}
			`
	var data3 map[string]interface{}
	json.Unmarshal([]byte(doc3),&data3)
	fmt.Println(data3)

	doc4, _ := json.MarshalIndent(data3,""," ")
	fmt.Println(string(doc4))
}