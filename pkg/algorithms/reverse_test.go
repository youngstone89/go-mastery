package algorithms_test

import (
	"fmt"
	"go-mastery/pkg/algorithms"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []string{"foo", "bar", "baz", "go", "stop"}
	reversedS := algorithms.Reverse(s)
	fmt.Println(reversedS)

}
