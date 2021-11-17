package rabbitmq_test

import (
	"go-mastery/pkg/rabbitmq"
	"sync"
	"testing"
)

func TestPublishToQueue(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.QueueDeclare("elc-dead-letter")
	p.Publish()
	defer p.Close()

	c := rabbitmq.NewRabbitMQConsumer(client)
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
	}
}

func TestConsume(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.ExchangeDeclare("elc-dead-letter-exchange", "fanout")
	p.QueueDeclare("elc-dead-letter")
	p.QueueBind()

	defer p.Close()

	c := rabbitmq.NewRabbitMQConsumer(client)
	c.Consume()

}

func TestConsumeMultiple(t *testing.T) {
	client := rabbitmq.NewRabbitMQClient("amqp://elc:elc@localhost:5672/")
	p := rabbitmq.NewRabbitMQProducer(client)
	p.ExchangeDeclare("elc-dead-letter-exchange", "fanout")
	p.QueueDeclare("elc-dead-letter")
	p.QueueBind()

	defer p.Close()

	c := rabbitmq.NewRabbitMQConsumer(client)
	go c.Consume()

	c1 := rabbitmq.NewRabbitMQConsumer(client)
	go c1.Consume()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
