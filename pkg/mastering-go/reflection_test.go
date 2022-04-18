package masteringgo_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflection(t *testing.T) {

	val1 := 1
	val2 := "1"
	val3 := true
	val4 := 6.6
	val5 := new(interface{})
	val6 := new(struct{})
	val7 := new(map[string]interface{})
	val8 := make(map[string]interface{}, 0)
	fmt.Println(reflect.TypeOf(val1), ":", reflect.ValueOf(val1))
	fmt.Println(reflect.TypeOf(val2), ":", reflect.ValueOf(val2))
	fmt.Println(reflect.TypeOf(val3), ":", reflect.ValueOf(val3))
	fmt.Println(reflect.TypeOf(val4), ":", reflect.ValueOf(val4))
	fmt.Println(reflect.TypeOf(val5), ":", reflect.ValueOf(val5))
	fmt.Println(reflect.TypeOf(val6), ":", reflect.ValueOf(val6))
	fmt.Println(reflect.TypeOf(val7), ":", reflect.ValueOf(val7))
	fmt.Println(reflect.TypeOf(*val7), ":", reflect.ValueOf(*val7))
	fmt.Println(reflect.TypeOf(val8), ":", reflect.ValueOf(val8))

	type Secret struct {
		Username string
		Password string
	}

	type Record struct {
		Field1 string
		Field2 float64
		Field3 Secret
	}
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}
	r := reflect.ValueOf(A)                  // value object
	fmt.Println("String value:", r.String()) // String value: <masteringgo_test.Record Value>
	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)                            //i Type: masteringgo_test.Record
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType) //The 3 fields of masteringgo_test.Record are
	for i := 0; i < r.NumField(); i++ {
		// Field1  with type: string       and value _String value_
		// Field2  with type: float64      and value _-12.123_
		// Field3  with type: masteringgo_test.Secret      and value _{Mihalis Tsoukalos}_
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())

		// Check whether there are other structures embedded in Record
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		fmt.Printf("\tand kind _%v_\n", k.String())

		// Need to convert it to string in order to compare it
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type()) //masteringgo_test.Secret

		}
		// Same as before but using the internal value
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type()) //masteringgo_test.Secret --> full unique name of the structure as defined by Go
		}
	}

}

func TestChangeStructureValuesUsingReflection(t *testing.T) {
	type T struct {
		F1 int
		F2 string
		F3 float64
	}
	A := T{1, "F2", 3.0}
	fmt.Println("A:", A) //A: {1 F2 3}
	r := reflect.ValueOf(&A).Elem()
	fmt.Println("String value:", r.String()) //String value: <masteringgo_test.T Value>
	typeOfA := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tOfA := typeOfA.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tOfA, f.Type(), f.Interface())
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r.Field(i).SetString("Changed!")
		} else if k == reflect.Float64 {
			r.Field(i).SetFloat(-6.6)
		}
		fmt.Println("A:", A)
	}
}
