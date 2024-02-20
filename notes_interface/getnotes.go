package notes_interface

import "github.com/aws/aws-lambda-go/events"

type NoteGetsService interface {
	GetNotes() events.APIGatewayProxyResponse
}
