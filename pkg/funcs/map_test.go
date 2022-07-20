package funcs_test

import (
	"fmt"
	"testing"
)

func TestMapPreSize(t *testing.T) {
	presizedMap := make(map[string]int, 1)
	presizedMap["key1"] = 1
	presizedMap["key2"] = 2
	fmt.Println(len(presizedMap))

	emptyMap := map[string]int{}
	fmt.Println(len(emptyMap))
	if emptyMap == nil {
		fmt.Println("nil")
	}

}
