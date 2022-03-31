package consumer

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaConsumer struct {
	consumer *kafka.Consumer
}
