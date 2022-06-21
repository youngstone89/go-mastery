package rabbitmq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	exchangeName string
	conn         *amqp.Connection
	ch           *amqp.Channel
	q            amqp.Queue
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
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type RabbitMQProducer struct {
	RabbitMQClient
}

func NewRabbitMQProducer(client *RabbitMQClient) *RabbitMQProducer {
	p := &RabbitMQProducer{
		RabbitMQClient: *client,
	}
	return p
}

func (r *RabbitMQProducer) Publish(file string) {

	data, err := ioutil.ReadFile("earthquake.json")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	var jsonArr []interface{}
	err = json.Unmarshal(data, &jsonArr)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	for _, obj := range jsonArr {
		data, _ := json.Marshal(obj)
		err := r.ch.Publish(
			r.exchangeName, // exchange
			r.q.Name,       // routing key
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(data),
			})
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}

}
