package handlers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/notes"
	// luisamaria.ariastorres@emeal.nttdata.com
)

func CreateNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	return notes.CreateNote(ctx, request, tableName, dynamoDBClient)
}

func GetNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	if request.Path == "/notes" {
		return notes.GetNotes(ctx, request, tableName, dynamoDBClient)
	} else {
		return notes.GetNote(ctx, request, tableName, dynamoDBClient)
	}
}

func DeleteNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	return notes.DeleteNote(ctx, request, tableName, dynamoDBClient)
}
func UpdateNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	return notes.UpdateNote(ctx, request, tableName, dynamoDBClient)
}

func UnhandledMethod(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Unsupported method"}, nil
}
