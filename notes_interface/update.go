package notes_interface

import (
	"github.com/aws/aws-lambda-go/events"
)

type UodateNoteService interface {
	UpdateNote(noteID string, body string) (events.APIGatewayProxyResponse, error)
}
