package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/handlers"
)

const tableName = "UserNote"

var dynamoDBClient *dynamodb.DynamoDB

func init() {
	// Create a new DynamoDB client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamoDBClient = dynamodb.New(sess)
}

func createNoteHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return handlers.GetNote(ctx, request, tableName, dynamoDBClient)
	case "POST":
		return handlers.CreateNote(ctx, request, tableName, dynamoDBClient)
	case "DELETE":
		return handlers.DeleteNote(ctx, request, tableName, dynamoDBClient)
	case "PUT":
		return handlers.UpdateNote(ctx, request, tableName, dynamoDBClient)
	default:
		return handlers.UnhandledMethod(ctx, request, tableName, dynamoDBClient)
	}
}

func main() {
	lambda.Start(createNoteHandler)
}
