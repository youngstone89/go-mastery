package rabbitmq_test

import (
	"go-mastery/pkg/rabbitmq"
	"testing"
	"time"
)

func TestPublishToQueue(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.QueueDeclare("elc-dead-letter")
	p.Publish()
	defer p.Close()

	c := rabbitmq.NewRabbitMQConsumer(client, "elc-dead-letter")
	c.Consume()
}

func TestPublishToExchange(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.ExchangeDeclare("elc-dead-letter-exchange", "fanout")
	p.QueueDeclare("elc-dead-letter")
	p.QueueBind()

	defer p.Close()

	for i := 0; i < 100; i++ {
		p.Publish()
		time.Sleep(1 * time.Second)
	}
}
