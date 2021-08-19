package jsonparsing

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID            string  `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID    string  `json:"customer_id"`
	Items         []Item  `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Records struct {

	Records []KeyValuePair `json:"records"`
	
}

func DoTest()  {

	str := `
{
    "id":"12345",
    "date_ordered":"2020-05-01T13:01:02Z",
    "customer_id":"3",
    "items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]
}
`
	var order Order
	err := json.Unmarshal([]byte(str),&order)
	if err != nil {
		panic(err)
	}
	input := `
{
"records": 
[
{
    "key":"12345",
    "value":"2020-05-01T13:01:02Z"
},
{
    "key":"12345",
    "value":"2020-05-01T13:01:02Z"
}
]
}
`
	//var kvPair KeyValuePair
	//var kvPairArr []KeyValuePair
	var records Records
	err = json.Unmarshal([]byte(input),&records)
	if err != nil {
		panic(err)
	}

}
