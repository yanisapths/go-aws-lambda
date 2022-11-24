package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// This data structure is understand by Golang
var todo = []Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Wash Dishes", Completed: false},
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var newTodo Todo

	if req.RequestContext.HTTP.Method == "POST" {
		json.Unmarshal([]byte(req.Body), &newTodo)
	}

	todo, err := json.Marshal(todo)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Something went Wrong",
		}, nil
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(todo),
	}, nil
}
