package rabbitmq_test

import (
	"go-mastery/pkg/rabbitmq"
	"testing"
)

func TestLoadBalance(t *testing.T) {

	go rabbitmq.RunNewWorker("elc-dead-letter", 1)
	go rabbitmq.RunNewWorker("elc-dead-letter", 2)
	go rabbitmq.RunNewTask("elc-dead-letter", 1)
	go rabbitmq.RunNewTask("elc-dead-letter", 2)
	go rabbitmq.RunNewTask("elc-dead-letter", 3)
	go rabbitmq.RunNewTask("elc-dead-letter", 4)
	go rabbitmq.RunNewTask("elc-dead-letter", 5)
	select {}

}

func TestLoadBalanceFromExchange(t *testing.T) {

	go rabbitmq.RunNewWorker("elc-dead-letter", 1)
	go rabbitmq.RunNewWorker("elc-dead-letter", 2)
	select {}

}
