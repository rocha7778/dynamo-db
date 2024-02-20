package notes_impl

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/validations"
	"github.com/rocha7778/dynamo-db/variables"
)

type DefaultNoteDeleteService struct {
	Repo db.DeleteServiceRepositoryInterface
}
type DeleteServiceRepository struct{}

// DeleteNote deletes a note
func (NoteService *DefaultNoteDeleteService) DeleteNote(noteID string) events.APIGatewayProxyResponse {
	if noteID == "" || !validations.IsValidNoteID(noteID) {
		return handleError("Note ID is required in path parameters", http.StatusBadRequest)
	}
	err := NoteService.Repo.DeleteItem(noteID)

	if err != nil {
		return handleError("Error deleting note from DynamoDB", http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{StatusCode: 204, Body: "Note deleted successfully"}
}

func (*DeleteServiceRepository) DeleteItem(noteId string) error {
	_, err := db.DBClient().DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(variables.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: &noteId},
		},
	})
	return err
}
