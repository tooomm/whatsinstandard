package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func RenderApp(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for k, v := range request.Headers {
		fmt.Printf("  %s: %s\n", k, v)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       request.Body,
	}, nil
}

func main() {
	lambda.Start(RenderApp)
}
