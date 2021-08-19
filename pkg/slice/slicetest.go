package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var jsonArr []interface{}
	obj := make(map[string]interface{})

	obj["key1"] = "val1"
	obj["key2"] = "val2"
	obj["key2"] = "val3"
	jsonArr = append(jsonArr, obj)

	jsonDoc, _ := json.MarshalIndent(jsonArr,""," ")
	fmt.Println(string(jsonDoc))
}


