package liskovsubstitution


import "fmt"

// Adder is implemented by objects that can add two integers together.
type Adder interface {
	Add(int, int) int
}

type Age struct {

}

func (a *Age) Add(i int, i2 int) int {
	return i + i2
}

func PrintSum(a, b int, adder Adder) string {
	return fmt.Sprintf("%d + %d = %d", a, b, adder.Add(a, b))
}


// Double adds two double values.
type Double struct{}

// Add returns the sum a+b.
func (Double) Add(a, b float64) float64 { return a + b }
