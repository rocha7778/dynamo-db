package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/notes_impl"
)

type NoteHandler struct {
	CreateNoteService  notes_impl.CreateNoteService
	GetNotesService    notes_impl.GetNotesCreateService
	DeleteNoteService  notes_impl.DeleteNoteService
	UpdateNoteService  notes_impl.UpdateNoteService
	GetNoteByIdService notes_impl.GetNoteServiceById
}

func NewNoteHandler() *NoteHandler {
	// Inicialización de repositorios
	createNoteRepo := &notes_impl.CreateNoteRepository{}
	getNotesRepo := &notes_impl.GetNotesServiceRepository{}
	getNoteRepo := &notes_impl.GetNoteServiceRepository{}
	deleteRepo := &notes_impl.DeleteServiceRepository{}
	updateRepo := &notes_impl.UpdateNoteServiceRepository{}

	// Inicialización de servicios con inyección de dependencias
	return &NoteHandler{
		CreateNoteService:  notes_impl.CreateNoteService{Repo: createNoteRepo},
		GetNotesService:    notes_impl.GetNotesCreateService{Repo: getNotesRepo},
		DeleteNoteService:  notes_impl.DeleteNoteService{Repo: deleteRepo},
		UpdateNoteService:  notes_impl.UpdateNoteService{Repo: updateRepo},
		GetNoteByIdService: notes_impl.GetNoteServiceById{Repo: getNoteRepo},
	}
}

func (h *NoteHandler) CreateNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	body := request.Body
	return h.CreateNoteService.CreateNote(body)
}

func (h *NoteHandler) GetNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	noteID := request.PathParameters["id"]
	if request.Path == "/notes" {
		return getNotes(h)
	}
	return getNoteById(h, noteID)

}

func getNotes(h *NoteHandler) events.APIGatewayProxyResponse {
	return h.GetNotesService.GetNotes()
}

func getNoteById(h *NoteHandler, noteID string) events.APIGatewayProxyResponse {
	return h.GetNoteByIdService.GetNoteById(noteID)
}

func (h *NoteHandler) DeleteNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	noteID := request.PathParameters["id"]
	return h.DeleteNoteService.DeleteNote(noteID)
}
func (h *NoteHandler) UpdateNote(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	noteID := request.PathParameters["id"]
	body := request.Body
	return h.UpdateNoteService.UpdateNote(noteID, body)
}

func (h *NoteHandler) UnhandledMethod() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Unsupported method"}
}
