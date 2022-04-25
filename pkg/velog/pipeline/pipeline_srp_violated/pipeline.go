package pipelinesrpviolated

type Pipeline struct {
}

func (p Pipeline) Read() {

}
func (p Pipeline) Process(data []byte) error {

	return nil
}
func (p Pipeline) Write(data []byte) error {
	return nil
}
