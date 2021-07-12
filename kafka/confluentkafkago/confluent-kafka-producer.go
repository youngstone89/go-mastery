package main



import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func main() {

	for i := 0; i < 100; i++ {
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"client.id": "test",
			"acks": "all",
			"message.timeout.ms" : 5000,
		})

		fmt.Printf("Failed to create producer: %s\n", p.GetFatalError())

		fmt.Printf("Failed to create producer: %s\n", p.Logs())

		fmt.Printf("Failed to create producer: %s\n", p.Len())

		if err != nil {
			fmt.Printf("Failed to create producer: %s\n", err)
			os.Exit(1)
		}


		value := "{\"topic\":\"yeongseok_topic\",\"key\":{\"service\":\"DLGaaS\",\"id\":\"977f87ae-5923-4281-8e91-1f1968baea7b\"},\"value\":{\"specversion\":\"1.0\",\"service\":\"DLGaaS\",\"source\":\"NIIEventLogger\",\"type\":\"NIIEventLog\",\"id\":\"977f87ae-5923-4281-8e91-1f1968baea7b\",\"timestamp\":\"2021-01-19T19:06:07.113Z\",\"appid\":\"DEMO-OMNICHANNEL-APP-DEV\",\"datacontenttype\":\"application/json\",\"data\":{\"dataContentType\":\"application/x-nuance-dlg-nii-logs.v1+json\",\"traceid\":\"8bbedbb2353cdc1d\",\"requestid\":\"2c06142d-71b7-4077-bf4b-e3892811b3d8\",\"sessionid\":\"0c234bb6-e2c4-4198-a7f7-53020b02d734\",\"locale\":\"en-US\",\"clientData\":{},\"seqid\":\"1\",\"events\":[{\"name\":\"session-start\",\"value\":{\"runtime\":{\"api\":\"v1\"},\"project\":{\"name\":\"TestMixClient\",\"namespace\":\"alex.smith@company.com\",\"deployed\":\"2020-12-10T19:52:25.224Z\",\"contextTag\":\"A1880_C1\",\"id\":\"8655\"},\"selector\":{\"channel\":\"default\",\"language\":\"en-US\"},\"version\":{\"asr\":{\"en_US\":\"2\"},\"nlu\":{\"en_US\":\"2\"},\"dlg\":\"2\"},\"user\":{\"systemID\":null,\"userChannelID\":null,\"userAuxiliaryID\":null,\"userGlobalID\":null,\"location\":null},\"timeout\":900}}]}},\"partition\":0,\"offset\":312}"
		topic := "yeongseok_topic"
		delivery_chan := make(chan kafka.Event, 10000)
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value: []byte(value)},
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
