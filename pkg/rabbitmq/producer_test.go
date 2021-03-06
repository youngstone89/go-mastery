package rabbitmq_test

import (
	"go-mastery/pkg/rabbitmq"
	"testing"
)

func TestPublishToQueue(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://user:bitnami@:::5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.QueueDeclare("earthquake")
	p.Publish()
	defer p.Close()

	// c := rabbitmq.NewRabbitMQConsumer(client, "earthquake")
	// c.Consume()
}

func TestPublishToExchange(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://guest:guest@:::5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.ExchangeDeclare("local.exchange", "direct")
	p.QueueDeclare("earthquake")
	p.QueueBind()

	defer p.Close()

	for i := 0; i < 100; i++ {
		p.Publish()
	}
}
