package notes_test

import (
	"errors"
	"testing"

	"github.com/rocha7778/dynamo-db/notes_impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock de DeleteServiceRepositoryInterface
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) DeleteItem(noteId string) error {
	args := m.Called(noteId)
	return args.Error(0)
}

// Caso de prueba positivo para DeleteNote
func TestDeleteNoteSuccess(t *testing.T) {
	mockRepo := new(MockRepo)
	noteID := "12346"
	mockRepo.On("DeleteItem", noteID).Return(nil)

	service := notes_impl.DeleteNoteService{
		Repo: mockRepo,
	}

	resp := service.DeleteNote(noteID)

	assert.Equal(t, 204, resp.StatusCode)
	assert.Equal(t, "Note deleted successfully", resp.Body)
	mockRepo.AssertExpectations(t)
}

// Caso de prueba negativo para DeleteNote - NoteID vacío
func TestDeleteNoteEmptyNoteID(t *testing.T) {
	service := notes_impl.DeleteNoteService{}

	resp := service.DeleteNote("")

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "Note ID is required in path parameters", resp.Body)
}

// Caso de prueba negativo para DeleteNote - Error al eliminar nota
func TestDeleteNoteDeleteItemError(t *testing.T) {
	mockRepo := new(MockRepo)
	noteID := "123456"
	mockRepo.On("DeleteItem", noteID).Return(errors.New("internal error"))

	service := notes_impl.DeleteNoteService{
		Repo: mockRepo,
	}

	resp := service.DeleteNote(noteID)

	assert.Equal(t, 500, resp.StatusCode)
	assert.Equal(t, "Error deleting note from DynamoDB", resp.Body)
	mockRepo.AssertExpectations(t)
}
