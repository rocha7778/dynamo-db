package notes

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/variables"
)

func DeleteNote(noteID string) (events.APIGatewayProxyResponse, error) {

	if noteID == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Note ID is required in path parameters"}, nil
	}

	// Delete the item from DynamoDB
	_, err := db.DBClient().DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(variables.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(noteID)},
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error deleting note from DynamoDB"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 204, Body: "Note deleted successfully"}, nil
}
