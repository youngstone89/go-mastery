package openandclose_test

import (
	"fmt"
	"testing"
)

type NoopProcessor struct {
	Validator
}

func (np NoopProcessor) Validate() {
	np.Validator.Validate()
	fmt.Printf("%T validating \n", np)
}

type Validator struct {
}

func (v Validator) Validate() {
	fmt.Printf("%T validating \n", v)
}

func Test(t *testing.T) {
	np := NoopProcessor{}
	np.Validate()

}
