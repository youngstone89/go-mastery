package rabbitmq

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type RabbitMQConsumer struct {
	RabbitMQClient
	ID    string
	queue string
}

func NewRabbitMQConsumer(client *RabbitMQClient, name string) *RabbitMQConsumer {
	c := &RabbitMQConsumer{
		RabbitMQClient: *client,
		ID:             uuid.New().String(),
		queue:          name,
	}
	return c
}

func (c *RabbitMQConsumer) Consume() {
	msgs, err := c.ch.Consume(
		c.queue, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		counter := 0
		for d := range msgs {
			counter++
			log.Printf("[Total:%d][Consumer:%v]Received a message: %s", counter, c, d.Body)
			time.Sleep(1 * time.Second)
		}
	}()

}
