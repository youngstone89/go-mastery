package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/confluentinc/confluent-kafka-go/kafka"
//	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/client_golang/prometheus/promhttp"
//	"net/http"
//	"sync"
//	"time"
//)
//
//
//
//
//func main1() {
//	http.Handle("/metrics", promhttp.Handler())
//	go func() {
//		http.ListenAndServe(":8078", nil)
//	}()
//
//	consumerWaiterGroup2()
//
//	//var wg sync.WaitGroup
//	//for i := 0; i < 3; i++ {
//	//	wg.Add(1)
//	//	consumerWaiterGroup(&wg)
//	//}
//	//wg.Wait()
//}
//
//type Counter struct {
//	i int
//	mu sync.Mutex
//	elcConnectionErrorCounterVec *prometheus.CounterVec
//}
//var counter *Counter
//
//func init()  {
//	counter = NewCounter(0)
//	counter.elcConnectionErrorCounterVec = prometheus.NewCounterVec(
//		prometheus.CounterOpts{
//			Name: "elcConnectionErrorCounterVec",
//			Help: "The number of processors configured",
//		},[]string{"topic"})
//	prometheus.Register(counter.elcConnectionErrorCounterVec)
//}
//
//func NewCounter(i int) *Counter {
//	return &Counter{i: i}
//}
//
//func (c *Counter) increment()  {
//	c.mu.Lock()
//	c.i++
//	c.mu.Unlock()
//}
//
//func consumerWaiterGroup2() {
//
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//
//	c, err := kafka.NewConsumer(&kafka.ConfigMap{
//		"bootstrap.servers": "localhost:9092",
//		"group.id":          "myGroup",
//		"auto.offset.reset": "earliest",
//		"go.application.rebalance.enable": true,
//	})
//
//	if err != nil {
//		panic(err)
//	}
//	c.SubscribeTopics([]string{"yeongseok_topic"}, nil)
//
//	//ev := c.Events()
//	//go func() {
//	//	for  {
//	//		select {
//	//		case <- ev:
//	//			fmt.Println(ev)
//	//		}
//	//	}
//	//}()
//
//	//time.Sleep(time.Duration(time.Second*60))
//	messages := make(chan map[string]interface{})
//	go func(counter *Counter) {
//		defer wg.Done()
//		for  {
//			select {
//			case msg:= <-messages:
//				counter.increment()
//				msgJson, _ := json.MarshalIndent(msg,"","")
//				counter.elcConnectionErrorCounterVec.WithLabelValues("yeongseok_topic").Inc()
//				fmt.Printf("\n!!!!!!!! got error message from go consumer gorounte: %v",string(msgJson))
//			}
//
//		}
//
//	}(counter)
//
//
//
//	go func(messages chan<- map[string]interface{}) {
//		defer wg.Done()
//		for {
//			msg, err := c.ReadMessage(-1)
//			//ev <- msg
//			if err == nil {
//				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
//				jsonObj := make(map[string]interface{})
//
//			} else {
//				// The client will automatically try to recover from all errors.
//				fmt.Printf("\nSending Consumer error: %v (%v)\n", err.(kafka.Error).Code(), err.(kafka.Error).String())
//				var msgMap = make(map[string]interface{})
//				msgMap["connectionError"] = err.(kafka.Error).String()
//				messages <- msgMap
//			}
//			time.Sleep(time.Duration(2*time.Second))
//		}
//		c.Close()
//	}(messages)
//	wg.Wait()
//	fmt.Printf("\nCounter: %v\n", counter.i)
//
//}
//func consumerWaiterGroup(wgp *sync.WaitGroup) {
//	defer wgp.Done()
//	//var wg sync.WaitGroup
//	//wg.Add(1)
//
//	c, err := kafka.NewConsumer(&kafka.ConfigMap{
//		"bootstrap.servers": "localhost",
//		"group.id":          "myGroup",
//		"auto.offset.reset": "earliest",
//	})
//
//	if err != nil {
//		panic(err)
//	}
//	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
//	//ev := c.Events()
//	//for  {
//	//	select {
//	//	case <- ev:
//	//		fmt.Println(ev)
//	//	}
//	//}
//	//time.Sleep(time.Duration(time.Second*60))
//
//	go func() {
//		//defer wg.Done()
//		for {
//			msg, err := c.ReadMessage(-1)
//			if err == nil {
//				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
//			} else {
//				// The client will automatically try to recover from all errors.
//				fmt.Printf("Consumer error: %v (%v)\n", err.(kafka.Error).Code(), err.(kafka.Error).String())
//			}
//		}
//		//c.Close()
//	}()
//	//wg.Wait()
//}
