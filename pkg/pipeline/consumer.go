package pipeline

type Consumer struct {
	count int
}

func NewConsumer(count int) *Consumer {
	c := &Consumer{count: count}

	return c
}

func (c *Consumer) Read(stream chan interface{}, errc chan error) error {

	for i := 0; i < c.count; i++ {
		// fmt.Printf("read index : %v \n", i)
		stream <- i

	}
	return nil
}
