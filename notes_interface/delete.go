package notes_interface

import (
	"github.com/aws/aws-lambda-go/events"
)

// NoteService defines the interface for managing notes
type DeleteNoteService interface {
	DeleteNote(noteID string) events.APIGatewayProxyResponse
}
