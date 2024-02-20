package notes_impl

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/variables"
)

type DefaultGetNotesCreateService struct {
	Repo db.GetNotesRepository
}

type GetNotesServiceRepository struct{}

func (s DefaultGetNotesCreateService) GetNotes() (events.APIGatewayProxyResponse, error) {

	// Get the item from DynamoDB
	result, err := s.Repo.Scam()

	// Check for errors
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error scanning DynamoDB table"}, errors.New("error scanning DynamoDB table")
	}

	// Check if any users were found
	if len(result.Items) == 0 {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: "Users not found"}, errors.New("user not found")
	}

	// Unmarshal the items into a slice of user structs
	var users []modelos.UserNote
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error unmarshaling users from DynamoDB"}, errors.New("error unmarshaling users from DynamoDB")
	}

	// Marshal the users slice into JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error marshaling users to JSON"}, errors.New("error marshaling users to JSON")
	}

	// Return the users as a response
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(usersJSON)}, nil

}

func (s *GetNotesServiceRepository) Scam() (*dynamodb.ScanOutput, error) {

	result, err := db.DBClient().Scan(&dynamodb.ScanInput{
		TableName: aws.String(variables.TableName),
	})
	return result, err
}
