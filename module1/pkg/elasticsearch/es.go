package es

import (
	"fmt"
	"log"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

type ESClient struct {
}

func New() *ESClient {
	return &ESClient{}
}

func (esc *ESClient) Init() {
	fmt.Println("initializing ElasticsearchClient...")
	es8, err := elasticsearch8.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es8.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}
