package make_test

import (
	"go-mastery/pkg/effective-go/data/make"
	"testing"
)

func TestDoTwoDimensionalSlice(t *testing.T) {
	make.DoTwoDimensionalSlice()
	make.DoTwoDimensionalSliceOneAllocation()
}
