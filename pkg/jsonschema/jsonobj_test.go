package jsonschema

import "testing"

func Test(t *testing.T) {
	type jsonObj = map[string]interface{}
	a := make(jsonObj)
	b := make(jsonObj)
	a["bObj"] = b
	obj := a["bObj"].(jsonObj)
	print(obj)

	type jsonObj2 map[string]interface{}
	a2 := make(jsonObj2)
	b2 := make(jsonObj2)
	a2["bObj"] = b2
	obj2 := a["bObj"].(jsonObj2)
	print(obj2)
}
