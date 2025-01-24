package main

import (
	"example/libs/queue"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var ENDPOINT_URL = "http://localhost:4566"

type Product struct {
	Action string `json:"action"`
	Data   any    `json:"data"`
}

var product_queue = queue.NewQueue[Product]("http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/teste", aws.Config{
	Region:       "us-west-2",
	BaseEndpoint: &ENDPOINT_URL,
})

func main() {
	product_queue.Send(Product{
		Action: "create",
		Data:   "product",
	})

	for {
		data, err := product_queue.Receive()

		if err != nil {
			fmt.Println(err)
		}

		if len(data) > 0 {
			fmt.Println(data)
		}
	}
}
