package jsonparsing

import (
	"encoding/json"
	"fmt"
)

type jsonObj map[string]interface{}

func DoMarshal() {

	var jsonArr = make([]jsonObj, 0)
	var sample = make(jsonObj)
	var nest = make(jsonObj)
	var nest2 = make(jsonObj)
	nest2["iam2"] = "here2"
	nest["iam"] = "here"
	nest["2ndchild"] = nest2
	sample["hi"] = "there"
	sample["child"] = nest
	d, e := json.Marshal(sample)
	if e != nil {
		fmt.Errorf("err %v", e)
	}
	fmt.Println(string(d))
	jsonArr = append(jsonArr, sample)

	r, e := json.Marshal(jsonArr)
	if e != nil {
		fmt.Errorf("err %v", e)
	}
	fmt.Println(string(r))

}
