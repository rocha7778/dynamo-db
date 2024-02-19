package notes_impl

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rocha7778/dynamo-db/db"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/variables"
)

// DefaultNoteService implements the NoteService interface
type DefaultNoteService struct{}
type CreateNoteRepository struct{}

func (s DefaultNoteService) CreateNote(body string, createNoteService db.CreateNoteRepository) (events.APIGatewayProxyResponse, error) {
	var note modelos.UserNote
	err := json.Unmarshal([]byte(body), &note)

	if err != nil {
		errMsg := fmt.Sprintf("THE BODY %s, Error unmarshaling request body: %s", body, err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errMsg}, nil
	}

	if note.ID == "" || note.Text == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "ID and Text fields are required"}, nil
	}

	// Marshal the note struct into JSON
	noteJSON, err := json.Marshal(note)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error marshaling JSON"}, nil
	}

	// Put the item into DynamoDB
	err = createNoteService.PutItem(&note)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error saving note to DynamoDB"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(noteJSON)}, nil
}

func (createNoteService CreateNoteRepository) PutItem(note *modelos.UserNote) error {
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
