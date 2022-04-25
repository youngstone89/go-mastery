package pipelinesrpgood

type Reader struct {
}

func (r Reader) Read() {

}

type Processor struct {
}

func (p Processor) Process(data []byte) error {

	return nil
}

type Writer struct {
}

func (w Writer) Write(data []byte) error {
	return nil
}
