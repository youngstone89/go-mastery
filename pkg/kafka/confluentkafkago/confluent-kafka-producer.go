package confluentkafkago



import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func DoProduce() {

	//ProduceEventMessages("clientdata_test_topic_tts","TTSaaS")
	//ProduceEventMessages("clientdata_test_topic_nlu","NLUaaS")
	//ProduceEventMessages("clientdata_test_topic_asr","ASRaaS")
	//ProduceEventMessages("clientdata_test_topic","ASRaaS")
	//ProduceEventMessages("clientdata_test_topic","TTSaaS")
	ProduceEventMessages("yeogseok_topic","DLGaaS")

}

func ProduceEventMessages(topic string, service string) {
	for i := 0; i < 10; i++ {
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers":  "localhost:9092",
			"client.id":          "test.client1",
			"acks":               "all",
			"message.timeout.ms": 5000,
		})

		fmt.Printf("Failed to create producer: %s\n", p.GetFatalError())

		fmt.Printf("Failed to create producer: %s\n", p.Logs())

		fmt.Printf("Failed to create producer: %s\n", p.Len())

		if err != nil {
			fmt.Printf("Failed to create producer: %s\n", err)
			os.Exit(1)
		}

		value := fmt.Sprintf("{\"key\":{\"id\":\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\",\"service\":\"DLGaaS\"},\"offset\":%d,\"partition\":0,\"topic\":\"%s\",\"value\":{\"appid\":\"Rakuten-VA-Assistant-Dev\",\"data\":{\"clientData\":{},\"dataContentType\":\"application/x-nuance-dlg-interaction-summary.v1+json\",\"events\":[{\"name\":\"transition\",\"value\":{\"from\":{\"component\":\"LoyaltyPoint_PaymentMethod\",\"name\":\"LoyaltyPoint_PaymentMethod_start\",\"type\":\"decision\",\"uuid\":\"d51b3358-13ed-49d9-af88-1342fe70dd70\"},\"to\":{\"component\":\"LoyaltyPoint_PaymentMethod\",\"name\":\"1_LoyaltyPoint_PaymentMethod_Message\",\"type\":\"playprompt\",\"uuid\":\"0c791767-2f9b-4628-bbed-6a676614cd55\"}}}],\"locale\":\"ja-JP\",\"request\":{\"clientData\":{\"x-nuance-event-log-source\":\"on-premise\"}},\"requestid\":\"67e5abd6-84ae-4b11-ad4c-ce88297d0ebb\",\"seqid\":\"71\",\"sessionid\":\"ec1c2329-32eb-4cc3-ae6f-4e78f7ea6ece\",\"traceid\":\"null\"},\"datacontenttype\":\"application/json\",\"id\":\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\",\"partitionKey\":\"{\\\"service\\\": \\\"DLGaaS\\\", \\\"id\\\": \\\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\\\"}\",\"service\":\"DLGaaS\",\"source\":\"NIIEventLogger\",\"specversion\":\"1.0\",\"timestamp\":\"2021-07-21T00:03:41.958Z\",\"type\":\"NIIEventLog\"}}",i,topic)
		//value := fmt.Sprintf("{\"key\":{\"id\":\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\",\"service\":\"DLGaaS\"},\"offset\":%d,\"partition\":0,\"topic\":\"%s\",\"value\":{\"appid\":\"Rakuten-VA-Assistant-Dev\",\"data\":{\"clientData\":{},\"dataContentType\":\"application/x-nuance-asr-recognitioninitmessage\",\"events\":[{\"name\":\"transition\",\"value\":{\"from\":{\"component\":\"LoyaltyPoint_PaymentMethod\",\"name\":\"LoyaltyPoint_PaymentMethod_start\",\"type\":\"decision\",\"uuid\":\"d51b3358-13ed-49d9-af88-1342fe70dd70\"},\"to\":{\"component\":\"LoyaltyPoint_PaymentMethod\",\"name\":\"1_LoyaltyPoint_PaymentMethod_Message\",\"type\":\"playprompt\",\"uuid\":\"0c791767-2f9b-4628-bbed-6a676614cd55\"}}}],\"locale\":\"ja-JP\",\"request\":{\"clientData\":{\"x-nuance-event-log-source\":\"on-premise\"}},\"requestid\":\"67e5abd6-84ae-4b11-ad4c-ce88297d0ebb\",\"seqid\":\"71\",\"sessionid\":\"ec1c2329-32eb-4cc3-ae6f-4e78f7ea6ece\",\"traceid\":\"null\"},\"datacontenttype\":\"application/json\",\"id\":\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\",\"partitionKey\":\"{\\\"service\\\": \\\"DLGaaS\\\", \\\"id\\\": \\\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\\\"}\",\"service\":\"DLGaaS\",\"source\":\"NIIEventLogger\",\"specversion\":\"1.0\",\"timestamp\":\"2021-07-21T00:03:41.958Z\",\"type\":\"NIIEventLog\"}}",i,topic)
		//value := fmt.Sprintf("{\"key\":{\"id\":\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\",\"service\":\"DLGaaS\"},\"offset\":%d,\"partition\":0,\"topic\":\"%s\",\"value\":{\"appid\":\"Rakuten-VA-Assistant-Dev\",\"data\":{\"clientData\":{},\"dataContentType\":\"iamnobody\",\"events\":[{\"name\":\"transition\",\"value\":{\"from\":{\"component\":\"LoyaltyPoint_PaymentMethod\",\"name\":\"LoyaltyPoint_PaymentMethod_start\",\"type\":\"decision\",\"uuid\":\"d51b3358-13ed-49d9-af88-1342fe70dd70\"},\"to\":{\"component\":\"LoyaltyPoint_PaymentMethod\",\"name\":\"1_LoyaltyPoint_PaymentMethod_Message\",\"type\":\"playprompt\",\"uuid\":\"0c791767-2f9b-4628-bbed-6a676614cd55\"}}}],\"locale\":\"ja-JP\",\"request\":{\"clientData\":{\"x-nuance-event-log-source\":\"on-premise\"}},\"requestid\":\"67e5abd6-84ae-4b11-ad4c-ce88297d0ebb\",\"seqid\":\"71\",\"sessionid\":\"ec1c2329-32eb-4cc3-ae6f-4e78f7ea6ece\",\"traceid\":\"null\"},\"datacontenttype\":\"application/json\",\"id\":\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\",\"partitionKey\":\"{\\\"service\\\": \\\"DLGaaS\\\", \\\"id\\\": \\\"ffff7f7a-14fb-462e-86bf-5dd3a3c947f3\\\"}\",\"service\":\"DLGaaS\",\"source\":\"NIIEventLogger\",\"specversion\":\"1.0\",\"timestamp\":\"2021-07-21T00:03:41.958Z\",\"type\":\"NIIEventLog\"}}",i,topic)

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
