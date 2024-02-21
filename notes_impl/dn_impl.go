package notes_impl

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/validations"
)

type DeleteNoteService struct {
	Repo db.DeleteServiceRepository
}

// DeleteNote deletes a note
func (NoteService *DeleteNoteService) DeleteNote(noteID string) events.APIGatewayProxyResponse {
	if noteID == "" || !validations.IsValidNoteID(noteID) {
		return handleError("Note ID is required in path parameters", http.StatusBadRequest)
	}
	err := NoteService.Repo.DeleteItem(noteID)

	if err != nil {
		return handleError("Error deleting note from DynamoDB", http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{StatusCode: 204, Body: "Note deleted successfully"}
}
