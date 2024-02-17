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

func GetNotes(ctx context.Context, request events.APIGatewayProxyRequest, tableName string, dynamoDBClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {

	// Scan the DynamoDB table to retrieve all users
	result, err := dynamoDBClient.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	// Check for errors
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error scanning DynamoDB table"}, nil
	}

	// Check if any users were found
	if len(result.Items) == 0 {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: "No users found"}, nil
	}

	// Unmarshal the items into a slice of user structs
	var users []modelos.UserNote
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error unmarshaling users from DynamoDB"}, nil
	}

	// Marshal the users slice into JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error marshaling users to JSON"}, nil
	}

	// Return the users as a response
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(usersJSON)}, nil
}
