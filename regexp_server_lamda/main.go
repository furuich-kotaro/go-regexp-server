package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(r events.APIGatewayProxyRequest) (Response, error) {
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "slack-interaction-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
