package main

import (
	"encoding/json"
	"main/config"
	"main/model"
	"main/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	dynoClient, err := config.LocalDyno()

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to initialize DynamoDB",
		}, err
	}

	repos := repository.NewTodoRepository(dynoClient, "todo")

	res, err := repos.FindAll()

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed fo fetch data in DynamoDB",
		}, err
	}

	webResp := model.TodosResponse{
		Code:  200,
		Todos: res,
	}

	j, _ := json.MarshalIndent(webResp, "", " ")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(j),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
