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
	"github.com/rocha7778/dynamo-db/validations"
	"github.com/rocha7778/dynamo-db/variables"
)

type GetNoteServiceById struct {
	Repo db.GetNoteRepository
}

type GetNoteServiceRepository struct{}

func (NoteService *GetNoteServiceById) GetNoteById(noteID string) events.APIGatewayProxyResponse {

	// Check if the note ID is empty

	if noteID == "" || !validations.IsValidNoteID(noteID) {
		return handleError("Note Id is mandatory and needs to take a valid value", http.StatusNotFound)
	}

	// Get the item from DynamoDB
	result, err := NoteService.Repo.GetItem(noteID)

	// Check for errors
	if err != nil {
		return handleError("Error getting note from DynamoDB", http.StatusInternalServerError)
	}

	// Check if the item exists
	if result.Item == nil {
		return handleError("Note not found", http.StatusNotFound)
	}

	// Unmarshal the item into a note struct
	var note modelos.UserNote
	err = dynamodbattribute.UnmarshalMap(result.Item, &note)
	if err != nil {
		return handleError("Error unmarshaling note from DynamoDB", http.StatusInternalServerError)
	}

	// Marshal the note struct into JSON
	noteJSON, err := json.Marshal(note)
	if err != nil {
		return handleError("Error marshaling note to JSON", http.StatusInternalServerError)
	}

	// Return the note as a response
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(noteJSON)}

}

func (*GetNoteServiceRepository) GetItem(noteID string) (*dynamodb.GetItemOutput, error) {
	result, err := db.DBClient().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(variables.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(noteID)},
		},
	})

	return result, err
}
