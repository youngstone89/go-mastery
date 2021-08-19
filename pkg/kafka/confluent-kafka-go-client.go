package main

func main() {
	//consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
	//	"bootstrap.servers":    "host1:9092,host2:9092",
	//	"group.id":             "foo",
	//	"auto.offset.reset":    "smallest"})

	//
	//
	//a := "NMDPTRIAL_peter_freshman_nuance_com_20191015T195942962874"
	//topic := &a
	//topicPartition := kafka.TopicPartition{
	//	Topic: topic,
	//	Partition: 0,
	//	//Offset: 0,
	//}
	//
	//kc, err := kafka.NewConsumer(&kafka.ConfigMap{
	//	"bootstrap.servers": "localhost:9092",
	//	"group.id":          "01",
	//	"auto.offset.reset": "earliest",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//topicPartitions := []kafka.TopicPartition{topicPartition}
	//err3 := kc.Assign(topicPartitions)
	//if err3 != nil {
	//	panic(err3)
	//}
	//assignedTopicPartitions, err4 := kc.Assignment()
	//
	//if err4 != nil {
	//	panic(err4)
	//}
	//fmt.Printf("assignedTopicPartitions :: %s", assignedTopicPartitions)
	//err5 := kc.Subscribe("NMDPTRIAL_peter_freshman_nuance_com_20191015T195942962874", nil)
	//if err5 != nil {
	//	panic(err5)
	//}
	//topics, err6 := kc.Subscription()
	//if err6 != nil {
	//	panic(err6)
	//}
	//fmt.Printf("topics:%s \n", topics)
	//
	//defer kc.Close()
	//offsetCount := 0
	//for {
	//	msg, err := kc.ReadMessage(-1)
	//	if err == nil {
	//		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	//		offsetCount++
	//	} else {
	//		// The client will automatically try to recover from all errors.
	//		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	//	}
	//}
	//fmt.Printf("OffsetCount: %v \n", offsetCount)
	//for e := range kc.Events() {
	//	switch ev := e.(type) {
	//	case	*kafka.Message:
	//		if ev.TopicPartition.Error != nil {
	//			fmt.Printf("Delivery Failed: %v\n", ev.TopicPartition)
	//		}else {
	//			fmt.Printf("Delivered Message : %v\n", ev.TopicPartition)
	//		}
	//	default:
	//		fmt.Printf("What's going on here : %v\n", ev.String())
	//	}
	//}
	//
	//
	//
	//
	//run := true
	//for run == true {
	//	ev := kc.Poll(0)
	//	switch e := ev.(type) {
	//	case *kafka.Message:
	//		var keyJsonObject map[string]interface{}
	//		var valueJsonObject map[string]interface{}
	//		json.Unmarshal(e.Key,&keyJsonObject)
	//		json.Unmarshal(e.Value,&valueJsonObject)
	//		fmt.Printf("Message Key: %s, \n Value: %s \n",keyJsonObject, valueJsonObject)
	//		fmt.Printf("Topic: %s \n",*e.TopicPartition.Topic )
	//		fmt.Printf("Partition: %s \n",e.TopicPartition.Partition )
	//		fmt.Printf("Offset: %s \n",e.TopicPartition.Offset )
	//		fmt.Printf("Metadata: %s \n",*e.TopicPartition.Metadata )
	//
	//
	//	case kafka.Error:
	//		fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
	//		run = false
	//	default:
	//		fmt.Printf("Ignored %v\n", e)
	//	}
	//}
	//
	//
	//msg, err := kc.ReadMessage(time.Duration(5000))
	//fmt.Printf("message :: %s", msg)
	//
	//
	//
	//err2 := kc.Seek(topicPartition,5000)
	//if err2 != nil {
	//	panic(err2)
	//}
	//
	//fmt.Printf("error: %s", err2)

}
