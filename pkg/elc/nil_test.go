package elc_test

import (
	"fmt"
	"testing"
)

type Consumer struct {
	url string
}

func (c *Consumer) Do() {
	fmt.Printf("Do %v \n", c.url)
}

func TestNilInstancePointerMethod(t *testing.T) {

	c := &Consumer{}
	// c = nil
	c.url = "test url"
	c.Do()

	// c = nil
	c.url = ""

	c.Do()

}

func TestNilMethod(t *testing.T) {

	c := &Consumer{url: "http://localhost"}
	// t.Logf("%v", c.url)
	c.Do()

	c = nil

	c.Do()

}
