package list_test

import (
	"fmt"
	"go-mastery/pkg/datastructure/list"
	"testing"
)

func TestAddToHead(t *testing.T) {
	var linkedList list.LinkedList
	linkedList = list.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	fmt.Println(linkedList.HeadNode.Property)
}

func TestAddToEnd(t *testing.T) {
	var linkedList list.LinkedList
	linkedList = list.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(5)
	linkedList.AddToHead(3)
}
