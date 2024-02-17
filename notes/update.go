package notes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/modelos"
)

func UpdateNote(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	// Extract note ID from request path parameters
	noteID := request.PathParameters["id"]
	if noteID == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Note ID is required in path parameters"}, nil
	}

	// Parse request body into UserNote struct
	var updatedNote modelos.UserNote
	err := json.Unmarshal([]byte(request.Body), &updatedNote)
	if err != nil {
		errMsg := fmt.Sprintf("Error unmarshaling request body: %s", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errMsg}, nil
	}

	// Update the item in DynamoDB
	_, err = dynamoDBClient.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(noteID)},
		},
		UpdateExpression: aws.String("SET #text = :text"),
		ExpressionAttributeNames: map[string]*string{
			"#text": aws.String("text"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":text": {S: aws.String(updatedNote.Text)},
		},
		ReturnValues: aws.String("ALL_NEW"),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error updating note in DynamoDB"}, nil
	}

	// Respond with success message and updated note
	responseBody, _ := json.Marshal(updatedNote)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}, nil
}
