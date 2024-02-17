package notes

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rocha7778/dynamo-db/modelos"
)

func GetNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {

	noteID := request.PathParameters["id"]

	// Check if the note ID is empty
	if noteID == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Note ID is required"}, nil
	}

	// Get the item from DynamoDB
	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(noteID)},
		},
	})

	// Check for errors
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error getting note from DynamoDB"}, nil
	}

	// Check if the item exists
	if result.Item == nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: "Note not found"}, nil
	}

	// Unmarshal the item into a note struct
	var note modelos.UserNote
	err = dynamodbattribute.UnmarshalMap(result.Item, &note)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error unmarshaling note from DynamoDB"}, nil
	}

	// Marshal the note struct into JSON
	noteJSON, err := json.Marshal(note)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error marshaling note to JSON"}, nil
	}

	// Return the note as a response
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(noteJSON)}, nil
}
