package notes_test

import (
	"errors"
	"testing"

	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/notes_impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para el repositorio
type MockUpdateNoteRepo struct {
	mock.Mock
}

func (m *MockUpdateNoteRepo) UpdateItem(note *modelos.UserNote) error {
	args := m.Called(note)
	return args.Error(0)
}

// Test positivo
func TestUpdateNoteSuccess(t *testing.T) {
	mockRepo := new(MockUpdateNoteRepo)
	noteService := notes_impl.NoteService{Repo: mockRepo}
	expectedNote := modelos.UserNote{ID: "123", Text: "Nota actualizada"}

	mockRepo.On("UpdateItem", &expectedNote).Return(nil)

	response := noteService.UpdateNote("123", `{ "id": "123", "text":"Nota actualizada"}`)

	assert.Equal(t, 200, response.StatusCode)
	mockRepo.AssertExpectations(t)
}

// Test negativo
func TestUpdateNoteFailureIDMissing(t *testing.T) {
	mockRepo := new(MockUpdateNoteRepo)
	noteService := notes_impl.NoteService{Repo: mockRepo}

	response := noteService.UpdateNote("", `{ "id": "123", "text":"Nota actualizada"}`)
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, "Note ID is required in path parameters", response.Body)
	mockRepo.AssertExpectations(t)

}

func TestUpdateNoteFailureUnMarsharl(t *testing.T) {
	mockRepo := new(MockUpdateNoteRepo)
	noteService := notes_impl.NoteService{Repo: mockRepo}

	response := noteService.UpdateNote("123", `{ "id": "123", "text":"Nota actualizada"`)
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, "Error unmarshaling request body: unexpected end of JSON input", response.Body)
	mockRepo.AssertExpectations(t)

}

func TestUpdateNoteFailureDynamoDBOperation(t *testing.T) {
	mockRepo := new(MockUpdateNoteRepo)
	noteService := notes_impl.NoteService{Repo: mockRepo}
	expectedNote := modelos.UserNote{ID: "123", Text: "Nota actualizada"}
	mockRepo.On("UpdateItem", &expectedNote).Return(errors.New("error updating note in DynamoDB"))

	response := noteService.UpdateNote("123", `{ "id": "123", "text":"Nota actualizada"}`)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, "Error updating note in DynamoDB", response.Body)
	mockRepo.AssertExpectations(t)

}
