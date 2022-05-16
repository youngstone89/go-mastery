package confluentkafkago_test

import (
	"context"
	"encoding/json"
	"go-mastery/pkg/kafka/confluentkafkago"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestProduce(t *testing.T) {
	value := `{
		"bootstrap.servers":  "localhost:9092",
		"client.id":          "client1",
		"transactional.id": "1"
	}`
	var config confluentkafkago.JsonObj
	json.Unmarshal([]byte(value), &config)

	p := confluentkafkago.NewConfluentKafkaProducer(config)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	err := p.Producer.InitTransactions(ctx)
	if err != nil {
		cancelFunc()
		t.Fatalf(err.Error())
	}

	topic := "dead_letter_queues"
	delivery_chan := make(chan kafka.Event, 1)

	value = `{
		"count": "5",
		"cause": "download failure",
		"event: "original event"
	}
	`
	p.Producer.BeginTransaction()
	err = p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("example key"),
		Value:          []byte(value),
	},

		delivery_chan)
	if err != nil {
		log.Fatal(err)
	}

	err = p.Producer.CommitTransaction(ctx)
	if err != nil {
		cancelFunc()
		t.Fatalf(err.Error())
	}
	e := <-delivery_chan
	m := e.(*kafka.Message)

	t.Logf("message:%v", m)

	close(delivery_chan)

}

func TestProduceMultiple(t *testing.T) {

	value := `{
		"bootstrap.servers":  "localhost:9092",
		"client.id":          "client1"
	}`
	var config confluentkafkago.JsonObj
	json.Unmarshal([]byte(value), &config)

	p := confluentkafkago.NewConfluentKafkaProducer(config)

	value = `{
		"bootstrap.servers":  "localhost:9092",
		"client.id":          "client1"
	}`
	var config2 confluentkafkago.JsonObj
	json.Unmarshal([]byte(value), &config2)

	// p2 := p
	p2 := confluentkafkago.NewConfluentKafkaProducer(config2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(p *confluentkafkago.ConfluentKafkaProducer) {
		defer wg.Done()
		// err := p.InitTransactions()
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		value := `{
			"producer": "p1",
			"count": "5",
			"cause": "download failure",
			"event: "original event"
		}
		`
		err := p.ProduceMessage("key1", value)
		if err != nil {
			log.Fatal(err)
			return
		}
	}(p)

	go func(p *confluentkafkago.ConfluentKafkaProducer) {
		defer wg.Done()
		// err := p2.InitTransactions()
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		value := `{
			"producer": "p2",
			"count": "5",
			"cause": "download failure",
			"event: "original event"
		}`
		err := p2.ProduceMessage("key2", value)
		if err != nil {
			log.Fatal(err)
			return
		}
	}(p)
	wg.Wait()

}

func TestProduceMockData(t *testing.T) {

	confluentkafkago.ProduceEventMessages("mock2")

}
