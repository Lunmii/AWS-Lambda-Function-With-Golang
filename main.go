package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name       string `json:"What's your name?"`
	Age        int    `json:"How Old are you?"`
	Occupation string `json:"What job do you do?"`
	Experience string `json:"What Experience do you have?"`
}
type MyResponse struct {
	Message string `json:"Answer"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%v is %v years old!, I work as a/an %v, currently posses above%v", event.Name, event.Age, event.Occupation, event.Experience)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
