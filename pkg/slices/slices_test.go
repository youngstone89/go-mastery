package slices_test

import (
	"go-mastery/pkg/slices"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		elements []int
	}{
		{
			desc:     "",
			elements: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			slices.DoWithPointerIntegerSlice(tC.elements)
			slices.DoWithPointerSlice(tC.elements...)
		})
	}
}
