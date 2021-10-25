package confluentkafkago

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func DoProduce() {

	//ProduceEventMessages("clientdata_test_topic_tts","TTSaaS")
	//ProduceEventMessages("clientdata_test_topic_nlu","NLUaaS")
	//ProduceEventMessages("clientdata_test_topic_asr","ASRaaS")
	//ProduceEventMessages("clientdata_test_topic","ASRaaS")
	//ProduceEventMessages("clientdata_test_topic","TTSaaS")
	ProduceEventMessages("partition_test", "DLGaaS")

}

func ProduceEventMessages(topic string, service string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"client.id":          "test.client1",
		"acks":               "all",
		"message.timeout.ms": 5000,
	})

	for i := 0; i < 10000000; i++ {

		fmt.Printf("Failed to create producer: %s\n", p.GetFatalError())

		fmt.Printf("Failed to create producer: %s\n", p.Logs())

		fmt.Printf("Failed to create producer: %s\n", p.Len())

		if err != nil {
			fmt.Printf("Failed to create producer: %s\n", err)
			os.Exit(1)
		}
		value := ``
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
