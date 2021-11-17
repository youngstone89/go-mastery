package rabbitmq

import "log"

type RabbitMQConsumer struct {
	RabbitMQClient
}

func NewRabbitMQConsumer(client *RabbitMQClient) *RabbitMQConsumer {
	c := &RabbitMQConsumer{
		RabbitMQClient: *client,
	}
	return c
}

func (c *RabbitMQConsumer) Consume() {
	msgs, err := c.ch.Consume(
		c.q.Name, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
