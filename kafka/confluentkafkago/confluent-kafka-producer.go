package main



import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func main() {

	//ProduceEventMessages("clientdata_test_topic_tts","TTSaaS")
	//ProduceEventMessages("clientdata_test_topic_nlu","NLUaaS")
	//ProduceEventMessages("clientdata_test_topic_asr","ASRaaS")
	//ProduceEventMessages("clientdata_test_topic","ASRaaS")
	//ProduceEventMessages("clientdata_test_topic","TTSaaS")
	ProduceEventMessages("error_topic_5","TTSaaS")

}

func ProduceEventMessages(topic string, service string) {
	for i := 0; i < 10000; i++ {
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

		value := fmt.Sprintf("{\n  \"appid\": \"%s\",\n  \"data\": \n{\"clientData\": {}, \n    \"dataContentType\": \"application/x-nuance-nluaas-interpretation.v1+json\",\n    \"locale\": \"jp-JP\",\n    \"processingTime\": {\n      \"durationMs\": 45,\n      \"startTime\": \"2021-07-15T01:29:33.241Z\"\n    },\n    \"request\": {\n      \"clientData\": {},\n      \"input\": {\n        \"inputUnion\": \"text\",\n        \"text\": \"こんにちは\"\n      },\n      \"model\": {\n        \"requestTimeoutMs\": 0,\n        \"type\": \"SEMANTIC_MODEL\",\n        \"uri\": \"urn:nuance-mix:tag:model/Mobile_CS_Active_Model/mix.nlu?=language=jpn-JPN\"\n      },\n      \"parameters\": {\n        \"interpretationInputLoggingMode\": \"PLAINTEXT\",\n        \"interpretationResultType\": \"SINGLE_INTENT\",\n        \"maxInterpretations\": 0,\n        \"postProcessingScriptParameters\": {}\n      },\n      \"resources\": [],\n      \"userId\": \"\"\n    },\n    \"response\": {\n      \"result\": null,\n      \"status\": {\n        \"code\": 500,\n        \"details\": \"com.nuance.nlu.runtime.api.interpret.processor.NtpeException: com.nuance.entrd.ntpe.tokenize.NtpeException: RequestId: 5f39c494-5e41-4846-9a7b-6e1af320573f, Severity: ERROR, Summary: Error while sending request to NTpE server. ntpe-ndp-jpn-jpn-rak-4-2-0\\ncom.nuance.entrd.ntpe.tokenize.NtpeException: RequestId: 5f39c494-5e41-4846-9a7b-6e1af320573f, Severity: ERROR, Summary: Error while sending request to NTpE server. ntpe-ndp-jpn-jpn-rak-4-2-0\\nRequestId: 5f39c494-5e41-4846-9a7b-6e1af320573f, Severity: ERROR, Summary: Error while sending request to NTpE server. ntpe-ndp-jpn-jpn-rak-4-2-0\\n\",\n        \"message\": \"Internal Server Error\"\n      }\n    }\n  },\n  \"datacontenttype\": \"application/json\",\n  \"id\": \"d07b4fce-a336-4ad2-92e6-924d359284cc\",\n  \"key\": \"{\\\"service\\\":\\\"NLUaaS\\\"}\",\n  \"offset\": %d,\n  \"partition\": 0,\n  \"partitionKey\": \"{\\\"service\\\":\\\"NLUaaS\\\",\\\"id\\\":\\\"d07b4fce-a336-4ad2-92e6-924d359284cc\\\"}\",\n  \"service\": \"NLUaaS\",\n  \"source\": \"nuance.nlu.v1.Runtime/Interpret\",\n  \"specversion\": \"1.0\",\n  \"timestamp\": \"2021-07-15T01:29:33.287Z\",\n  \"topic\": \"%s\",\n  \"type\": \"Interpret\",\n  \"value\": {\n    \"appid\": \"%s\",\n    \"data\": {\n      \"dataContentType\": \"application/x-nuance-nluaas-interpretation.v1+json\",\n      \"locale\": \"jp-JP\",\n      \"processingTime\": {\n        \"durationMs\": 45,\n        \"startTime\": \"2021-07-15T01:29:33.241Z\"\n      },\n      \"request\": {\n        \"clientData\": {},\n        \"input\": {\n          \"inputUnion\": \"text\",\n          \"text\": \"こんにちは\"\n        },\n        \"model\": {\n          \"requestTimeoutMs\": 0,\n          \"type\": \"SEMANTIC_MODEL\",\n          \"uri\": \"urn:nuance-mix:tag:model/Mobile_CS_Active_Model/mix.nlu?=language=jpn-JPN\"\n        },\n        \"parameters\": {\n          \"interpretationInputLoggingMode\": \"PLAINTEXT\",\n          \"interpretationResultType\": \"SINGLE_INTENT\",\n          \"maxInterpretations\": 0,\n          \"postProcessingScriptParameters\": {}\n        },\n        \"resources\": [],\n        \"userId\": \"\"\n      },\n      \"response\": {\n        \"result\": null,\n        \"status\": {\n          \"code\": 500,\n          \"details\": \"com.nuance.nlu.runtime.api.interpret.processor.NtpeException: com.nuance.entrd.ntpe.tokenize.NtpeException: RequestId: 5f39c494-5e41-4846-9a7b-6e1af320573f, Severity: ERROR, Summary: Error while sending request to NTpE server. ntpe-ndp-jpn-jpn-rak-4-2-0\\ncom.nuance.entrd.ntpe.tokenize.NtpeException: RequestId: 5f39c494-5e41-4846-9a7b-6e1af320573f, Severity: ERROR, Summary: Error while sending request to NTpE server. ntpe-ndp-jpn-jpn-rak-4-2-0\\nRequestId: 5f39c494-5e41-4846-9a7b-6e1af320573f, Severity: ERROR, Summary: Error while sending request to NTpE server. ntpe-ndp-jpn-jpn-rak-4-2-0\\n\",\n          \"message\": \"Internal Server Error\"\n        }\n      }\n    },\n    \"datacontenttype\": \"application/json\",\n    \"id\": \"d07b4fce-a336-4ad2-92e6-924d359284cc\",\n    \"partitionKey\": \"{\\\"service\\\":\\\"%s\\\",\\\"id\\\":\\\"d07b4fce-a336-4ad2-92e6-924d359284cc\\\"}\",\n    \"service\": \"%s\",\n    \"source\": \"nuance.nlu.v1.Runtime/Interpret\",\n    \"specversion\": \"1.0\",\n    \"timestamp\": \"2021-07-15T01:29:33.287Z\",\n    \"type\": \"Interpret\"\n  }\n}",topic,i,topic,topic, service,service)
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
