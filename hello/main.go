package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Hello handler")
	log.Println(request)
	var message = fmt.Sprintf("Hello %s, from AWS Lambda using Golang!", request.QueryStringParameters["name"])

	return events.APIGatewayProxyResponse{Body: message, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
