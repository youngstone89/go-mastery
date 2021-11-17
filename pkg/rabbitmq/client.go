package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewRabbitMQClient(host string) *RabbitMQClient {
	conn, err := amqp.Dial(host)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	rc := &RabbitMQClient{
		conn: conn,
		ch:   ch,
	}
	return rc
}
func (r *RabbitMQClient) Close() {
	r.conn.Close()
	r.ch.Close()
}
func (r *RabbitMQClient) QueueDeclare() {
	q, err := r.ch.QueueDeclare(
		"elc-dead-letter", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare a queue")
	r.q = q
}

func (r *RabbitMQClient) ExchangeDeclare() {
	err := r.ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
