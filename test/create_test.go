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
	args := m.Called(note)
	return args.Error(0)
}

func setupService() (*notes_impl.CreateNoteService, *MockCreateNoteService) {
	mockRepo := new(MockCreateNoteService)
	service := &notes_impl.CreateNoteService{
		Repo: mockRepo, // Inyectar el mock aquí
	}
	return service, mockRepo
}

func TestCreateNoteSuccess(t *testing.T) {
	// Arrange
	service, mockRepository := setupService()

	note := &modelos.UserNote{
		ID:   "123",
		Text: "test-text",
	}

	body, _ := json.Marshal(note)
	requestBody := string(body)

	mockRepository.On("PutItem", note).Return(nil)

	// Act
	response := service.CreateNote(requestBody)

	// Assert
	assert.Equal(t, events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, response)
}

func TestCreateNoteInvalidRequestBody(t *testing.T) {
	// Creating CreateNoteService instance
	service, _ := setupService()

	note := &modelos.UserNote{
		ID:   "",
		Text: "",
	}
	body, _ := json.Marshal(note)
	requestBody := string(body)

	// Calling the CreateNote function with invalid request body
	response := service.CreateNote(requestBody)

	// Asserting the response
	assert.Equal(t, 400, response.StatusCode)
	assert.Contains(t, response.Body, "ID and Text fields are required")
}

func TestCreateNoteEmptyBody(t *testing.T) {
	// Creating CreateNoteService instance
	service, _ := setupService()

	// Calling the CreateNote function with invalid request body
	response := service.CreateNote("")

	// Asserting the response
	assert.Equal(t, 400, response.StatusCode)
}
