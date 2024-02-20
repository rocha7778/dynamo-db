package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/notes_impl"
)

var (
	createNoteServiceRepo = &notes_impl.CreateNoteRepository{}
	getNotesServiceRepo   = &notes_impl.GetNotesServiceRepository{}
	getNoteServiceRepo    = &notes_impl.GetNoteServiceRepository{}
	deleteServiceRepo     = &notes_impl.DeleteServiceRepository{}
	updateServiceRepo     = &notes_impl.UpdateNoteServiceRepository{}
)

var (
	noteService        = notes_impl.CreateNoteService{Repo: createNoteServiceRepo}
	getNoteService     = notes_impl.GetNotesCreateService{Repo: getNotesServiceRepo}
	deleteNoteService  = notes_impl.DeleteNoteService{Repo: deleteServiceRepo}
	updateNoteService  = notes_impl.UpdateNoteService{Repo: updateServiceRepo}
	getNoteByIdService = notes_impl.GetNoteServiceById{Repo: getNoteServiceRepo}
)

func CreateNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	body := request.Body
	return noteService.CreateNote(body)
}

func GetNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	noteID := request.PathParameters["id"]
	if request.Path == "/notes" {
		return getNotes()
	}
	return getNoteById(noteID)

}

func getNotes() events.APIGatewayProxyResponse {
	return getNoteService.GetNotes()
}

func getNoteById(noteID string) events.APIGatewayProxyResponse {
	return getNoteByIdService.GetNoteById(noteID)
}

func DeleteNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	noteID := request.PathParameters["id"]
	return deleteNoteService.DeleteNote(noteID)
}
func UpdateNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	noteID := request.PathParameters["id"]
	body := request.Body
	return updateNoteService.UpdateNote(noteID, body)
}

func UnhandledMethod() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Unsupported method"}
}
