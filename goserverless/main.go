package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.APIGatewayProxyResponse) (Response, error) {
	resp := Response{ StatusCode: 200, Body: "hello" }

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
