package consumer_test

import (
	"fmt"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestSetOffset(t *testing.T) {

	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     "localhost",
		"group.id":              "myGroup",
		"auto.offset.reset":     "earliest",
		"broker.address.family": "v4",
		"session.timeout.ms":    6000,
	})
	if err != nil {
		t.Fail()
	}
	adminClient, err := kafka.NewAdminClientFromConsumer(kc)

	if err != nil {
		t.Fail()
	}
	topic := new(string)
	*topic = "NMDPTRIAL_kim_seok_nuance_com_20210427T044351526748"
	md, err := adminClient.GetMetadata(topic, false, 1000)
	if err != nil {
		t.Error(err)
	}
	// close admin client
	adminClient.Close()

	// get parition array
	var partitions []int32

	for _, p := range md.Topics[*topic].Partitions {
		partitions = append(partitions, p.ID)
	}
	//declare TopicPartiton slice to commit
	var topicPartitions []kafka.TopicPartition

	for _, p := range partitions {
		tp := &kafka.TopicPartition{
			Topic:     topic,
			Partition: int32(p),
		}
		err := tp.Offset.Set(10)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		topicPartitions = append(topicPartitions, *tp)
	}

	committedTopicPartitions, err := kc.CommitOffsets(topicPartitions)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	for _, tp := range committedTopicPartitions {
		fmt.Printf("partition %d set to offset %s ", tp.Partition, tp.Offset.String())
	}

	// close consumer
	err = kc.Close()
	if err != nil {
		t.Error(err)
	}

}

func TestShowOffset(t *testing.T) {

	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     "localhost",
		"group.id":              "myGroup",
		"auto.offset.reset":     "earliest",
		"broker.address.family": "v4",
		"session.timeout.ms":    6000,
	})
	if err != nil {
		t.Fail()
	}
	adminClient, err := kafka.NewAdminClientFromConsumer(kc)

	if err != nil {
		t.Fail()
	}
	topic := new(string)
	*topic = "NMDPTRIAL_kim_seok_nuance_com_20210427T044351526748"
	md, err := adminClient.GetMetadata(topic, false, 1000)
	if err != nil {
		t.Error(err)
	}
	// close admin client
	adminClient.Close()

	// get parition array
	var partitions []int32

	for _, p := range md.Topics[*topic].Partitions {
		partitions = append(partitions, p.ID)
	}

	for _, p := range partitions {
		low, high, err := kc.QueryWatermarkOffsets(*topic, p, -1)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		tp := &kafka.TopicPartition{
			Topic:     topic,
			Partition: int32(p),
		}
		var partitions []kafka.TopicPartition
		partitions = append(partitions, *tp)
		offsets, err := kc.Committed(partitions, -1)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		current := offsets[0].Offset
		t.Logf("low: %d, high: %d, current: %d", low, high, current)
	}

	// close consumer
	err = kc.Close()
	if err != nil {
		t.Error(err)
	}

}

func TestResetOffset(t *testing.T) {

	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     "localhost",
		"group.id":              "myGroup",
		"auto.offset.reset":     "earliest",
		"broker.address.family": "v4",
		"session.timeout.ms":    6000,
	})
	if err != nil {
		t.Fail()
	}
	adminClient, err := kafka.NewAdminClientFromConsumer(kc)

	if err != nil {
		t.Fail()
	}
	topic := new(string)
	*topic = "NMDPTRIAL_kim_seok_nuance_com_20210427T044351526748"
	md, err := adminClient.GetMetadata(topic, false, 1000)
	if err != nil {
		t.Error(err)
	}
	// close admin client
	adminClient.Close()

	// get parition array
	var partitions []int32

	for _, p := range md.Topics[*topic].Partitions {
		partitions = append(partitions, p.ID)
	}

	var topicPartitions []kafka.TopicPartition

	for _, p := range partitions {
		low, _, err := kc.QueryWatermarkOffsets(*topic, p, -1)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		tp := &kafka.TopicPartition{
			Topic:     topic,
			Partition: int32(p),
		}
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		err = tp.Offset.Set(low)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		topicPartitions = append(topicPartitions, *tp)
	}
	committedTopicPartitions, err := kc.CommitOffsets(topicPartitions)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	for _, tp := range committedTopicPartitions {
		fmt.Printf("partition %d reset to offset %s ", tp.Partition, tp.Offset.String())
	}

	// close consumer
	err = kc.Close()
	if err != nil {
		t.Error(err)
	}

}
