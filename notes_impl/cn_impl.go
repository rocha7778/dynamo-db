package notes_impl

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/validations"
)

// DefaultNoteService implements the NoteService interface
type CreateNoteService struct {
	Repo db.CreateNoteRepository // Asume que db.CreateNoteRepository es una interfaz
}

func (service *CreateNoteService) CreateNote(body string) events.APIGatewayProxyResponse {
	var note modelos.UserNote

	if err := json.Unmarshal([]byte(body), &note); err != nil {
		errMsg := "Error al procesar el cuerpo de la solicitud"
		return handleError(errMsg, 400)
	}

	if note.ID == "" || note.Text == "" {
		return handleError("ID and Text fields are required", http.StatusBadRequest)

	}

	if !validations.IsValidNoteID(note.ID) {
		return handleError("ID needs to take a valid value", http.StatusBadRequest)
	}

	if err := service.Repo.PutItem(&note); err != nil {
		return handleError("Error saving note to DynamoDB", http.StatusInternalServerError)
	}

	noteJSON, _ := json.Marshal(note)
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(noteJSON)}
}
