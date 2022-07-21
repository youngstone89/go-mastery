package defer_test

import (
	"fmt"
	"testing"
)

func TestPrintSutff(t *testing.T) {
	v := printStuff()
	fmt.Println(v)
}

func printStuff() (value string) {

	defer fmt.Println("exiting")

	defer func() {

		value = "we returned this"

	}()

	fmt.Println("I am printing stuff")

	return ""

}
