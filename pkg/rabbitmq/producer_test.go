package rabbitmq_test

import (
	"go-mastery/pkg/rabbitmq"
	"testing"
)

func TestPublishToQueue(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.QueueDeclare()
	p.Publish()
	defer p.Close()

	c := rabbitmq.NewRabbitMQConsumer(client)
	c.Consume()
}
