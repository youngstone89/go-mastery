package enumerators_test

import (
	"fmt"
	"testing"
)

//go:generate stringer -type=Pill

type Pill int

const (
	Placebo Pill = iota

	Aspirin

	Ibuprofen

	Paracetamol

	Acetaminophen = Paracetamol
)

func Test(t *testing.T) {
	fmt.Println(Placebo, Aspirin, Ibuprofen, Paracetamol, Acetaminophen)
}
