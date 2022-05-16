package confluentkafkago

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func DoProduce() {

	//ProduceEventMessages("clientdata_test_topic_tts","TTSaaS")
	//ProduceEventMessages("clientdata_test_topic_nlu","NLUaaS")
	//ProduceEventMessages("clientdata_test_topic_asr","ASRaaS")
	//ProduceEventMessages("clientdata_test_topic","ASRaaS")
	//ProduceEventMessages("clientdata_test_topic","TTSaaS")
	// ProduceEventMessages("partition_test", "DLGaaS")

}

type ConfluentKafkaProducer struct {
	Producer       *kafka.Producer
	ctx            context.Context
	Topic          *string
	DeliveryReport chan kafka.Event
}
type JsonObj map[string]interface{}

func NewConfluentKafkaProducer(config JsonObj) *ConfluentKafkaProducer {
	value, err := json.Marshal(config)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}
	var configMap *kafka.ConfigMap
	json.Unmarshal(value, &configMap)

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil
	}
	topic := "dead_letter_queues"
	producer := &ConfluentKafkaProducer{
		Producer:       p,
		ctx:            context.Background(),
		Topic:          &topic,
		DeliveryReport: make(chan kafka.Event, 1),
	}
	return producer
}

func (c *ConfluentKafkaProducer) InitTransactions() error {
	ctx, _ := context.WithTimeout(c.ctx, time.Duration(5*time.Second))
	// defer cancelFun()
	err := c.Producer.InitTransactions(ctx)
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}
	return nil
}

func (c *ConfluentKafkaProducer) ProduceMessage(key string, value string) error {
	// c.Producer.BeginTransaction()
	err := c.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: c.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(value),
	},

		c.DeliveryReport)
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}

	// ctx, _ := context.WithTimeout(c.ctx, time.Duration(10*time.Second))
	// defer cancelFun()
	// Retry := true
	// for Retry {
	// 	err = c.Producer.CommitTransaction(ctx)
	// 	if err == nil {
	// 		break
	// 	} else if err.(kafka.Error).TxnRequiresAbort() {
	// 		c.Producer.AbortTransaction(ctx)
	// 		fmt.Errorf("abort")
	// 		break
	// 	} else if err.(kafka.Error).IsRetriable() {
	// 		fmt.Errorf("retry")
	// 		continue
	// 	} else { // treat all other errors as fatal errors
	// 		fmt.Errorf(err.Error())
	// 		break
	// 	}
	// }

	e := <-c.DeliveryReport
	m := e.(*kafka.Message)

	fmt.Printf("message:%v", m)
	return nil

}

func ProduceEventMessages(topic string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group1",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		fmt.Printf("Failed to create producer: %s\n", p.GetFatalError())
		os.Exit(1)
	}
	for i := 0; i < 100; i++ {

		value := `{
			"hi": "there",
		}`
		delivery_chan := make(chan kafka.Event, 10000)
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(value)},
			delivery_chan,
		)
		if err != nil {
			log.Fatal(err)
		}

		e := <-delivery_chan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

		close(delivery_chan)

	}
}
