package notes_impl

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/variables"
)

type DefaultGetNotesCreateService struct{}

func (s DefaultGetNotesCreateService) GetNotes() (events.APIGatewayProxyResponse, error) {

	// Scan the DynamoDB table to retrieve all users
	result, err := db.DBClient().Scan(&dynamodb.ScanInput{
		TableName: aws.String(variables.TableName),
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
