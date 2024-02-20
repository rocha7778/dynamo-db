package notes_impl

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/validations"
	"github.com/rocha7778/dynamo-db/variables"
)

// DefaultNoteService implements the NoteService interface
type CreateNoteService struct {
	Repo db.CreateNoteRepository // Asume que db.CreateNoteRepository es una interfaz
}
type CreateNoteRepository struct{}

func (service *CreateNoteService) CreateNote(body string) events.APIGatewayProxyResponse {
	var note modelos.UserNote

	if err := json.Unmarshal([]byte(body), &note); err != nil {
		errMsg := "Error al procesar el cuerpo de la solicitud"
		return handleError(errMsg, 400)
	}

	if note.ID == "" || note.Text == "" {
		return handleError("ID and Text fields are required", http.StatusBadRequest)

	}

	if !validations.IsValidNoteID(note.ID) {
		return handleError("ID needs to take a valid value", http.StatusBadRequest)
	}

	if err := service.Repo.PutItem(&note); err != nil {
		return handleError("Error saving note to DynamoDB", http.StatusInternalServerError)
	}

	noteJSON, _ := json.Marshal(note)
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(noteJSON)}
}

func (*CreateNoteRepository) PutItem(note *modelos.UserNote) error {
	// Put the item into DynamoDB
	_, err := db.DBClient().PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(variables.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id":   {S: aws.String(note.ID)},
			"text": {S: aws.String(note.Text)},
		},
	})

	return err
}
