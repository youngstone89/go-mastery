package main

import (
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {


	retry := 0
	maxRetry := -1
	run := true

	for run == true {
		err := runKafkaConsumer()
		if err != nil {
			run = false
		}
		if retry > maxRetry && maxRetry > 0 {
			run = false
		}
		time.Sleep(time.Duration(5 * time.Second))
	}
}

func runKafkaConsumer() error {
	topic := "error_topic_5"
	//Create initial consumer
	c, err := NewConsumer()

	if err != nil {
		return err
	}

	//Subscribe
	err = c.Subscribe(topic, nil)

	if err != nil {
		return err
	}

	adminClient, err := kafka.NewAdminClientFromConsumer(c)

	if err != nil {
		return err
	}
	topicPointer := &topic
	md, err := adminClient.GetMetadata(topicPointer, false, -1)

	if err != nil {
		return err
	}

	var partitions []float64
	for _, v := range md.Topics[topic].Partitions {
		partitions = append(partitions, float64(v.ID))
	}

	if len(partitions) <= 0 {
		return errors.New("Partition Count is zero or less.")
	}

	if err != nil {
		return err
	}

	errChan := make(chan error)
	sigChan := make(chan os.Signal)
	eofChan := make(chan kafka.PartitionEOF)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	var wg sync.WaitGroup
	err = c.Close()
	if err != nil {
		errChan <- err
		return err
	}
	partitionsCount := len(partitions)
	for p := range partitions {
		c, err := NewConsumer()
		if err != nil {
			errChan <- err
			return err
		}
		wg.Add(1)
		go process(p, c, topic, topicPointer, errChan, &wg, sigChan,eofChan)
	}

	run := true
	for run {
		select {
		case cerr := <-errChan:
			run = false
			log.Printf("Error: %s", cerr)
		case eof := <-eofChan:
			partitionsCount--
			log.Printf("[Event: %s][Partition Count Now : %v]",eof,partitionsCount)
			if partitionsCount == 0 {
				run = false
			}
		}
	}
	wg.Wait()
	log.Printf("================== Process Ended =========================")
	return nil
}

func NewConsumer() (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    "localhost:9092",
		"group.id":             "08",
		"auto.offset.reset":    "earliest",
		"enable.auto.commit":   "false",
		"enable.partition.eof": "true",
		"auto.commit.interval.ms": "5000",
		"go.application.rebalance.enable": true,
		//"debug": "cgrp,topic,fetch",
		//"session.timeout.ms": "60000",
		//"heartbeat.interval.ms": "15000",
		//"max.poll.interval.ms": "300000",
	})
	return c, err
}

func process(p int, c *kafka.Consumer, topic string, topicPointer *string, errChan chan error, wg *sync.WaitGroup, sigChan chan os.Signal, eofChan chan kafka.PartitionEOF) error {
	defer wg.Done()
	log.Printf("Partition ID:%v \n", p)
	low, high, err := c.QueryWatermarkOffsets(topic, int32(p), -1)
	if err != nil {
		sendErrorToChannel(errChan, err)
		return err
	}
	log.Printf("[Consumer: %s][Partition: %d][Offset Range low:%v, high:%v]  \n", c.String(),p, low, high)
	var topicParitions []kafka.TopicPartition
	tp := kafka.TopicPartition{
		Topic:     topicPointer,
		Partition: int32(p),
	}
	topicParitions = append(topicParitions, tp)
	aerr := c.Assign(topicParitions)
	if aerr != nil {
		sendErrorToChannel(errChan, err)
		return aerr
	}
	committedOffsets, err := c.Committed(topicParitions, -1)

	if err != nil {
		sendErrorToChannel(errChan, err)
		return err
	}

	log.Printf("[Consumer: %s][Partition: %d][Committed Offsets :%v] \n", c.String(),p, committedOffsets[0].Offset.String())
	currentOffset := committedOffsets[0].Offset

	if int64(currentOffset) > high || int64(currentOffset) < low {
		currentOffset = 0
	}
	log.Printf("[Consumer: %s][Partition: %d][Committed Offsets :%v] \n", c.String(),p, currentOffset)

	run := true
	offset, err := strconv.Atoi(currentOffset.String())
	if err != nil {
		sendErrorToChannel(errChan, err)
		return err
	}
	seekTp := kafka.TopicPartition{
		topicPointer, int32(p), kafka.Offset(offset), nil, nil,
	}
	serr := c.Seek(seekTp, -1)
	if serr != nil {
		sendErrorToChannel(errChan, err)
		return serr
	}
	for run {
		select {
		case sig := <- sigChan:
			log.Printf("Caught signal %v: terminating\n", sig)
			err  := c.Close()
			if err != nil {
				errChan <- err
				return err
			}
			run = false
		default:
			ev := c.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				//record := cast(e)
				//raw, _ := json.Marshal(record)
				var partitionsToCommit []kafka.TopicPartition
				tp := kafka.TopicPartition{
					Topic:     topicPointer,
					Partition: int32(p),
				}
				offset++
				err := tp.Offset.Set(offset)
				if err != nil {
					sendErrorToChannel(errChan, err)
					return err
				}
				partitionsToCommit = append(partitionsToCommit, tp)

				committedOffsets, err := c.CommitOffsets(partitionsToCommit)
				if err != nil {
					sendErrorToChannel(errChan, err)
					return err
				}

				intOffset, err := strconv.Atoi(committedOffsets[0].Offset.String())
				if err != nil {
					errChan <- err
					return err
				}
				_, high, err := c.QueryWatermarkOffsets(topic, int32(p), -1)
				intHigh:= int(high)
				//intOffset, err := strconv.Atoi(committedOffsets[0].Offset.String())
				left := intHigh - intOffset

				//log.Printf("[Consumer: %s][Partition: %d][Committed Offsets :%v][Left Offsets :%d][End Offsets :%d] \n", c.String(),p, committedOffsets[0].Offset.String(),left,intHigh)
				committedOffsetsToCheck, cerr := c.Committed(partitionsToCommit, -1)


				if cerr != nil {
					sendErrorToChannel(errChan, err)
					return cerr
				}
				log.Printf("[Consumer: %s][Partition: %d][Committed Offsets: %v][Left Offsets :%d][End Offsets :%d] \n", c.String(),p, committedOffsetsToCheck[0].Offset.String(),left,intHigh)

			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				log.Printf("Error: %v \n: %v \n", e.Code(), e)
				err = e
			case kafka.PartitionEOF:
				log.Printf("PartitionEOF [Consumer: %s][Partition: %d] \n", c.String(),p)
				c.Unassign()
				c.Unsubscribe()
				eofChan <- e
				wg.Done()
			case kafka.OffsetsCommitted:
				log.Printf("OffsetsCommitted [Consumer: %s][Partition: %d][OffsetCommitted Event: %v] \n", c.String(),p, e)

			case kafka.AssignedPartitions:
				log.Printf("AssignedPartitions [Consumer: %s][Partition: %d][AssignedPartitions Event: %v] \n", c.String(),p, e)

			default:
				//log.Printf("Default Event [Consumer: %s][Partition: %d][Event: %s] \n", c.String(),p,e.String())
				//log.Printf("Default Event [Consumer: %s][Partition: %d][Event: %s] \n", c.String(),p,e)

			}
			//time.Sleep(time.Duration(5 * time.Second))
		}
		}

	return nil
}

func sendErrorToChannel(errChan chan error, err error) {
	errChan <- err
}
