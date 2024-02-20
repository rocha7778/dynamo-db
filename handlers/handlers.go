package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/notes_impl"
)

// Considerando la posibilidad de utilizar instancias singleton o inyección de dependencias para reutilizar instancias de servicios
var (
	createNoteService   = &notes_impl.CreateNoteRepository{}
	getNotesServiceRepo = &notes_impl.GetNotesServiceRepository{}
	getNoteServiceRepo  = &notes_impl.GetNoteServiceRepository{}
	deleteServiceRepo   = &notes_impl.DeleteServiceRepository{}
	updateServiceRepo   = &notes_impl.UpdateNoteServiceRepository{}
)

var (
	noteService       = notes_impl.DefaultNoteService{}
	getNoteService    = notes_impl.DefaultGetNotesCreateService{Repo: getNotesServiceRepo}
	deleteNoteService = notes_impl.DefaultNoteDeleteService{Repo: deleteServiceRepo}
	updateNoteService = notes_impl.NoteService{Repo: updateServiceRepo}
)

func CreateNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	body := request.Body
	return noteService.CreateNote(body, createNoteService)
}

func GetNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	noteID := request.PathParameters["id"]
	if request.Path == "/notes" {
		return getNotes()
	} else {
		return getNoteById(noteID)
	}
}

func getNotes() events.APIGatewayProxyResponse {
	return getNoteService.GetNotes()
}

func getNoteById(noteID string) events.APIGatewayProxyResponse {
	service := notes_impl.DefaultNoteGetService{Repo: getNoteServiceRepo}
	return service.GetNoteById(noteID)
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
