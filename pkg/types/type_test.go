package types_test

import (
	"fmt"
	"testing"
)

func TestZeroValue(t *testing.T) {
	var a uint
	fmt.Println(a)
	var b uint16
	fmt.Println(b)
	var c uint32
	fmt.Println(c)
	var d uint64
	fmt.Println(d)
	var e uintptr
	fmt.Println(e)

	var f int
	fmt.Println(f)

	var h int16
	fmt.Println(h)

	var i int32
	fmt.Println(i)

	var j int64
	fmt.Println(j)

	var k string
	fmt.Println(k)

	if k == "" {
		fmt.Println("zero value is \"\"")
	}
	var l float32
	fmt.Printf("%f \n", l)

	var m float64
	fmt.Printf("%f \n", m)

	fmt.Println(struct{}{})
	var s struct{}
	fmt.Println(s)

	var y interface{}
	fmt.Println(y)
	fmt.Printf("%v \n", y)

	p := &y
	fmt.Println(p)
	fmt.Printf("%v \n", p)

	var ch chan struct{}
	ch = make(chan struct{}, 1)
	fmt.Println(ch)
	ch <- struct{}{}
	fmt.Println(<-ch)
	var x byte
	fmt.Println(x)
	var arr [1]string
	arr[0] = "hey"
	fmt.Println(arr)
	fmt.Println(len(arr))
	var arr2 []int
	fmt.Println(len(arr2))

	_ = "hey"
	_ = 7

}
