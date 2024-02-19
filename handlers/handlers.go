package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/notes"
)

func CreateNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := request.Body
	return notes.CreateNote(body)
}

func GetNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	noteID := request.PathParameters["id"]

	if request.Path == "/notes" {
		return notes.GetNotes()
	} else {
		return notes.GetNoteById(noteID)
	}
}

func DeleteNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	noteID := request.PathParameters["id"]
	return notes.DeleteNote(noteID)
}
func UpdateNote(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	noteID := request.PathParameters["id"]
	body := request.Body
	return notes.UpdateNote(noteID, body)
}

func UnhandledMethod() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Unsupported method"}, nil
}
