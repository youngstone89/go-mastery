package maps_test

import (
	"fmt"
	"testing"
)

type jsonObj = map[string]any

func TestModifyMap(t *testing.T) {
	obj := make(jsonObj)
	modify(obj)
	fmt.Printf("%p \n", obj)
	t.Log(obj)

	newObj := obj
	fmt.Printf("%p \n", newObj)
	t.Log(obj)

	modifyAgain(newObj)
	t.Log(newObj)
	t.Log(obj)

}

func modify(obj jsonObj) error {
	obj["modify"] = true
	fmt.Printf("%p \n", obj)
	return nil
}
func modifyAgain(obj jsonObj) error {
	obj["modify"] = false
	fmt.Printf("%p \n", obj)
	return nil
}
