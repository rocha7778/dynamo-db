package notes_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/notes_impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCreateNoteRepositoryFaild struct {
	mock.Mock
}

func (m *mockCreateNoteRepositoryFaild) PutItem(note *modelos.UserNote) error {
	arg := m.Called(note)
	return arg.Error(0)
}

func TestCreateNoteErrorSavingToDynamoDB(t *testing.T) {
	mockRepository := &mockCreateNoteRepositoryFaild{}
	service := notes_impl.DefaultNoteService{}
	note := modelos.UserNote{
		ID:   "123",
		Text: "Sample note",
	}
	noteJSON, _ := json.Marshal(note)
	mockRepository.On("PutItem", &note).Return(errors.New("error saving to DynamoDB"))
	response, err := service.CreateNote(string(noteJSON), mockRepository)
	assert.NotNil(t, err)
	assert.Equal(t, 500, response.StatusCode)
	assert.Contains(t, response.Body, "Error saving note to DynamoDB")
	mockRepository.AssertCalled(t, "PutItem", &note)
}
