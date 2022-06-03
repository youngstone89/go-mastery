package types

import "testing"

var _ Processor = new(ProcessorFunc)

type Processor interface {
	Process()
}

type ProcessorFunc func()

func (f ProcessorFunc) Process() {
	f()
}

func hi() {
	print("hi")
}
func hoi() {
	print("hoi")
}
func himsg(msg string) {
	print(msg)
}

func RunProcess(p Processor) {
	p.Process()
}

func Test_process(t *testing.T) {
	p := ProcessorFunc(hi)
	t.Logf("%T, %v", p, p)
	// p.Process()
	RunProcess(p)

	p2 := ProcessorFunc(hoi)
	t.Logf("%T, %v", p2, p2)
	RunProcess(p2)

}
