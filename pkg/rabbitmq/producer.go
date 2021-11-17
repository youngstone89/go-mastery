package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
	RabbitMQClient
}

func NewRabbitMQProducer(client *RabbitMQClient) *RabbitMQProducer {
	p := &RabbitMQProducer{
		RabbitMQClient: *client,
	}
	return p
}

func (r *RabbitMQProducer) Publish() {

	body := `
	{
		"envelope": {
			"letter": {
				"records": [
					{
						"key": "some key",
						"partition":"0",
						"offset": "0"
					}
				]
			},
			"error": {
				"retry_count": 1,
				"description": "some error message",
				"reason": "some cause"
			}
		}
	}
	`

	err := r.ch.Publish(
		"",       // exchange
		r.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}
