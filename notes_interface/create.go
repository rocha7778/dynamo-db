package notes_interface

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/db"
)

type NoteService interface {
	CreateNote(body string, createNoteRepository db.CreateNoteRepository) (events.APIGatewayProxyResponse, error)
}
