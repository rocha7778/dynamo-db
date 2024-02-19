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

func (m mockCreateNoteRepository) PutItem(note *modelos.UserNote) error {
	// Mock implementation, you can customize as needed for your test cases
	return nil
}

func TestCreateNote_Success(t *testing.T) {
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
	response, err := service.CreateNote(requestBody, mockRepository)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, response)
}
