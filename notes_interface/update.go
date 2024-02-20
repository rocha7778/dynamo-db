package notes_interface

import (
	"github.com/aws/aws-lambda-go/events"
)

// NoteService defines the interface for managing notes
type UodateNoteService interface {
	UpdateNote(noteID string, body string) (events.APIGatewayProxyResponse, error)
}
