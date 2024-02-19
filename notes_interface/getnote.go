package notes_interface

import "github.com/aws/aws-lambda-go/events"

type NoteGetService interface {
	GetNoteById(noteID string) (events.APIGatewayProxyResponse, error)
}
