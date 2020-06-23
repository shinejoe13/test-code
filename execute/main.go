package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func ExecuteOperation(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	test := "Lambda executed after TTL expired"
	log.Println("test", test)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{"Access-Control-Allow-Origin": "*",
			"Content-Type": "application/json",
		},
		Body: test,
	}, nil

}

func main() {
	lambda.Start(ExecuteOperation)
}
