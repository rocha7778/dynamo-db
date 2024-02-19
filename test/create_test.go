package notes_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/notes_impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCreateNoteRepository struct {
	mock.Mock
}

func (m *mockCreateNoteRepository) PutItem(note *modelos.UserNote) error {
	return nil
}

func TestCreateNoteSuccess(t *testing.T) {
	// Arrange
	mockRepository := mockCreateNoteRepository{}
	mockRepository.On("PutItem").Return(nil)

	service := notes_impl.DefaultNoteService{}

	note := &modelos.UserNote{
		ID:   "test-id",
		Text: "test-text",
	}
	body, _ := json.Marshal(note)
	requestBody := string(body)

	// Act
	response, err := service.CreateNote(requestBody, &mockRepository)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, response)
}

func TestCreateNoteInvalidRequestBody(t *testing.T) {
	// Creating DefaultNoteService instance
	service := notes_impl.DefaultNoteService{}

	note := &modelos.UserNote{
		ID:   "",
		Text: "",
	}
	body, _ := json.Marshal(note)
	requestBody := string(body)

	// Calling the CreateNote function with invalid request body
	response, err := service.CreateNote(requestBody, nil)

	// Asserting the response
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.StatusCode)
	assert.Contains(t, response.Body, "ID and Text fields are required")
}

func TestCreateNoteEmptyBody(t *testing.T) {
	// Creating DefaultNoteService instance
	service := notes_impl.DefaultNoteService{}

	// Calling the CreateNote function with invalid request body
	response, err := service.CreateNote("", nil)

	// Asserting the response
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.StatusCode)
}
