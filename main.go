package main

import "go-mastery/pkg/form"

func main() {

	//confluentkafkago.DoProduce()
	//regex.DoRegexTest()
	//generics.DoGenericsTest()
	//generics.DoStackTest()
	//generics.DoGenericsTest()

	//webhook.DoWebhookTest()

	go form.DoReceiveTest()
	form.DoMultipartFormTest()
}

