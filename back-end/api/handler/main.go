package main

import (
	"context"
	"net/http"

	awsdynamodb "github.com/VinciusSouzaDosReis/vi-todo/libs/aws/dynamodb"
	"github.com/VinciusSouzaDosReis/vi-todo/libs/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var helloRepository *repository.RepositoryHello

func main() {
	dynamodbConnection := awsdynamodb.NewDynamoDBConnection()

	helloRepository = repository.NewRepositoryHello(
		dynamodbConnection,
	)

	lambda.Start(Handler)
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res, err := helloRepository.GetHello(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       res.Name,
	}, nil
}
