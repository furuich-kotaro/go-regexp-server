package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

type RequestBody struct {
	Text  string `json:"Text"`
	Regex string `json:"Regex"`
}

type ResponseBody struct {
	Result [][]string `json:"Result"`
}

func ConvertInputDataToStruct(inputs string) (*RequestBody, error) {
	var req RequestBody
	err := json.Unmarshal([]byte(inputs), &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func verify(h map[string]string) error {
	if h["authentication"] != os.Getenv("AUTHENTICATION_VALUE") {
		log.Printf("[ERROR] failed to Authentication: %s", h["Authentication"])
		return errors.New("invalid Authentication value")
	}

	return nil
}

func Handler(req events.APIGatewayProxyRequest) (Response, error) {
	e, err := json.Marshal(req)
	if err != nil {
		log.Printf("[ERROR] failed to Marshal: %v", err)
	}
	log.Println(string(e))

	if err := verify(req.Headers); err != nil {
		log.Printf("[ERROR] failed to verify: %v", err)
		return Response{StatusCode: 403}, nil
	}
	jsonBody, err := ConvertInputDataToStruct(req.Body)
	if err != nil {
		log.Printf("[ERROR] failed to convert to struct: %v", err)
		return Response{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	s := jsonBody.Text
	r := regexp.MustCompile(jsonBody.Regex)
	matchAllStrings := r.FindAllStringSubmatch(s, -1)

	var resBody ResponseBody
	resBody.Result = matchAllStrings
	jsonBytes, _ := json.Marshal(resBody)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonBytes),
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
