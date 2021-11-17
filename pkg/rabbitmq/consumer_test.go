package rabbitmq_test

import (
	"go-mastery/pkg/rabbitmq"
	"testing"
	"time"
)

func TestConsume(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.ExchangeDeclare("elc-dead-letter-exchange", "fanout")
	p.QueueDeclare("elc-dead-letter")
	p.QueueBind()

	defer p.Close()

	c := rabbitmq.NewRabbitMQConsumer(client, "elc-dead-letter")
	c.Consume()

}

func TestMultipleConsumersFromExchange(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	client.QueueDeclare("elc-dead-letter")
	c := rabbitmq.NewRabbitMQConsumer(client, "elc-dead-letter")
	c.Consume()

	client1 := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	client1.QueueDeclare("elc-dead-letter")
	c1 := rabbitmq.NewRabbitMQConsumer(client1, "elc-dead-letter")
	c1.Consume()

	defer c.Close()
	time.Sleep(100 * time.Second)

}
