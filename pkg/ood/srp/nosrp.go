package srp

import (
	"fmt"
	"io"
)

// CalculatorNotSrp calculates the test coverage for a directory
// and it's sub-directories
type CalculatorNotSrp struct {
	// coverage data populated by `Calculate()` method
	data map[string]float64
}

// Calculate will calculate the coverage
func (c *CalculatorNotSrp) Calculate(path string) error {
	// run `go test -cover ./[path]/...` and store the results
	return nil
}

// Output will print the coverage data to the supplied writer
func (c CalculatorNotSrp) Output(writer io.Writer) {
	for path, result := range c.data {
		fmt.Fprintf(writer, "%s -> %.1f\n", path, result)
	}
}

// OutputCSV will print the coverage data to the supplied writer
func (c CalculatorNotSrp) OutputCSV(writer io.Writer) {
	for path, result := range c.data {
		fmt.Fprintf(writer, "%s,%.1f\n", path, result)
	}
}
