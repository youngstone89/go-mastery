package webhook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	webhookData := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&webhookData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("got webhook payload: ")
	for k, v := range webhookData {
		fmt.Printf("%s : %v\n", k, v)
	}

	switch name := webhookData["username"]; name{
	case "test":
		fmt.Println("Invoked")
	default:
		fmt.Println("Not invoked")
	}
}

func DoWebhookTest()  {
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}