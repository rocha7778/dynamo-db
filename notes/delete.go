package notes

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	// Extract note ID from request path parameters
	noteID := request.PathParameters["id"]

	if noteID == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Note ID is required in path parameters"}, nil
	}

	// Delete the item from DynamoDB
	_, err := dynamoDBClient.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(noteID)},
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error deleting note from DynamoDB"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 204, Body: "Note deleted successfully"}, nil
}
