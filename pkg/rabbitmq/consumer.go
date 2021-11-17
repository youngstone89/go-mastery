package rabbitmq

import (
	"log"

	"github.com/google/uuid"
)

type RabbitMQConsumer struct {
	RabbitMQClient
	ID string
}

func NewRabbitMQConsumer(client *RabbitMQClient) *RabbitMQConsumer {
	c := &RabbitMQConsumer{
		RabbitMQClient: *client,
		ID:             uuid.New().String(),
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
		counter := 0
		for d := range msgs {
			counter++
			log.Printf("[Total:%d][Consumer:%v]Received a message: %s", counter, c, d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
