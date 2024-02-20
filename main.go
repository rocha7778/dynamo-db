package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rocha7778/dynamo-db/handlers"
)

func createNoteHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handlers := handlers.NewNoteHandler()

	switch request.HTTPMethod {
	case "GET":
		return handlers.GetNote(request), nil
	case "POST":
		return handlers.CreateNote(request), nil
	case "DELETE":
		return handlers.DeleteNote(request), nil
	case "PUT":
		return handlers.UpdateNote(request), nil
	default:
		return handlers.UnhandledMethod(), nil
	}
}

func main() {
	lambda.Start(createNoteHandler)
}
