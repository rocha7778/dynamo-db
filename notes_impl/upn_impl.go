package notes_impl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/validations"
)

type UpdateNoteService struct {
	Repo db.UpdateNoteRepository
}

func (NoteService *UpdateNoteService) UpdateNote(noteID string, body string) events.APIGatewayProxyResponse {
	// Extract note ID from request path parameters
	if noteID == "" || !validations.IsValidNoteID(noteID) {
		return handleError("Note ID is required in path parameters", http.StatusBadRequest)
	}
	// Parse request body into UserNote struct
	var updatedNote modelos.UserNote
	err := json.Unmarshal([]byte(body), &updatedNote)
	if err != nil {
		errMsg := fmt.Sprintf("Error unmarshaling request body: %s", err.Error())
		return handleError(errMsg, http.StatusBadRequest)
	}

	// Update the item in DynamoDB
	err = NoteService.Repo.UpdateItem(&updatedNote)
	if err != nil {
		return handleError("Error updating note in DynamoDB", http.StatusInternalServerError)
	}
	// Respond with success message and updated note
	responseBody, _ := json.Marshal(updatedNote)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody)}
}
