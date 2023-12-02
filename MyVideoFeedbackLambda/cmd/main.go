package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/WriteRightProject/WriteRightLambda/app"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Hello API Gateway event! %s\n", event.Body)

	app := app.App{}

	feedback := app.GetFeedbackOfYoutubeVideo(event)

	//Lets marshal this feedback and send as a response
	data, err := json.Marshal(feedback)
	if err != nil {
		log.Fatal("Could not marshal youtube video feedback, ", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	//We parsed the feedback, lets send this
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(data),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
