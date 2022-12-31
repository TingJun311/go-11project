package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name"`
	Age int `json:"How old are you"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse {
		Message: fmt.Sprint("%s is %d years old!", event.Name, event.Age),
	}, nil
}