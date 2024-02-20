package notes_impl

import (
	"encoding/json"
	"net/http"

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

func (NoteService DefaultGetNotesCreateService) GetNotes() events.APIGatewayProxyResponse {

	// Get the item from DynamoDB
	result, err := NoteService.Repo.Scam()

	// Check for errors
	if err != nil {
		return handleError("Error scanning DynamoDB table", http.StatusInternalServerError)
	}

	// Check if any users were found
	if len(result.Items) == 0 {
		return handleError("Notes not found", http.StatusOK)
	}

	// Unmarshal the items into a slice of user structs
	var users []modelos.UserNote
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return handleError("Error unmarshaling users from DynamoDB", http.StatusInternalServerError)
	}

	// Marshal the users slice into JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		return handleError("Error marshaling notes to JSON", http.StatusInternalServerError)
	}

	// Return the users as a response
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(usersJSON)}

}

func (*GetNotesServiceRepository) Scam() (*dynamodb.ScanOutput, error) {

	result, err := db.DBClient().Scan(&dynamodb.ScanInput{
		TableName: aws.String(variables.TableName),
	})
	return result, err
}
