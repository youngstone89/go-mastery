package kafka

import "fmt"

type ConfluentKafkaClient struct {
}

func New() *ConfluentKafkaClient {
	return &ConfluentKafkaClient{}
}

func (ckc *ConfluentKafkaClient) Init() {
	fmt.Println("initializing confluent kafka client...")
}
