package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rocha7778/dynamo-db/handlers"
)

func createNoteHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	switch request.HTTPMethod {
	case "GET":
		return handlers.GetNote(request)
	case "POST":
		return handlers.CreateNote(request)
	case "DELETE":
		return handlers.DeleteNote(request)
	case "PUT":
		return handlers.UpdateNote(request)
	default:
		return handlers.UnhandledMethod()
	}
}

func main() {
	lambda.Start(createNoteHandler)
}
