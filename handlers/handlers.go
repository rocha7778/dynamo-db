package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/notes"
	"github.com/rocha7778/dynamo-db/notes_impl"
)

var repo = &notes_impl.GetNotesServiceRepository{}
var service = notes_impl.DefaultGetNotesCreateService{
	Repo: repo,
}

var repoById = &notes_impl.GetNoteServiceRepository{}
var serviceById = notes_impl.DefaultNoteGetService{
	Repo: repoById,
}

func CreateNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := request.Body
	service := notes_impl.DefaultNoteService{}
	createNoteService := &notes_impl.CreateNoteRepository{}
	return service.CreateNote(body, createNoteService)
}

func GetNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	noteID := request.PathParameters["id"]
	if request.Path == "/notes" {
		return service.GetNotes()
	} else {
		return serviceById.GetNoteById(noteID)
	}
}

func DeleteNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	noteID := request.PathParameters["id"]
	var repo = &notes_impl.DeleteServiceRepository{}
	var service = notes_impl.DefaultNoteDeleteService{
		Repo: repo,
	}
	return service.DeleteNote(noteID)
}
func UpdateNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	noteID := request.PathParameters["id"]
	body := request.Body
	return notes.UpdateNote(noteID, body)
}

func UnhandledMethod() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Unsupported method"}, nil
}
