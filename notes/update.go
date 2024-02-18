package notes

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/variables"
)

func UpdateNote(noteID string, body string) (events.APIGatewayProxyResponse, error) {
	// Extract note ID from request path parameters
	if noteID == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Note ID is required in path parameters"}, nil
	}

	// Parse request body into UserNote struct
	var updatedNote modelos.UserNote
	err := json.Unmarshal([]byte(body), &updatedNote)
	if err != nil {
		errMsg := fmt.Sprintf("Error unmarshaling request body: %s", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errMsg}, nil
	}

	// Update the item in DynamoDB
	_, err = db.DBClient().UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(variables.TableName),
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
