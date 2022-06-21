package generics_test

import (
	"fmt"
	"go-mastery/pkg/generics"
	"testing"
)

// Initialize a map for the integer values
var ints = map[string]int64{
	"first":  34,
	"second": 12,
}

// Initialize a map for the float values
var floats = map[string]float64{
	"first":  35.98,
	"second": 26.99,
}

func TestSum(t *testing.T) {

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		generics.SumInts(ints),
		generics.SumFloats(floats))

}

func TestSumIntsOrFloats(t *testing.T) {

	fmt.Printf("Generic Sums: %v and %v\n",
		generics.SumIntsOrFloats[string, int64](ints),
		generics.SumIntsOrFloats[string, float64](floats))

}

func TestSumNumbers(t *testing.T) {

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		generics.SumNumbers(ints),
		generics.SumNumbers(floats))

}
