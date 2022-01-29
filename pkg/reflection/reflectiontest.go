package reflection

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}

var x Foo

func DoSomething(f Foo) {
	fmt.Println(f.A, f.B)
}

func DoReflectionTest() {

	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name()) // returns int
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name()) // returns Foo
	fmt.Println(ft.Kind()) // struct

	//slice, pointer don't have names
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())        // returns an empty string
	fmt.Println(xpt.Kind())        // kind
	fmt.Println(xpt.Elem().Name()) // pointer's name
	fmt.Println(xpt.Elem().Kind()) // pointer's kind

	var is []interface{}
	sft := reflect.TypeOf(is)
	fmt.Println(sft.Name())
	fmt.Println(sft.Kind())

	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(),
			curField.Tag.Get("myTag"))
	}

	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)        // sv is of type reflect.Value
	s2 := sv.Interface().([]string) // s2 is of type []string
	fmt.Println(s2)

	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	fmt.Println(ivv)
	ivv.SetInt(20)
	fmt.Println(i)
	fmt.Println(iv.Elem().Interface().(int))
}
