package notes_test

import (
	"testing"

	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/notes_impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateNoteService struct {
	mock.Mock
}

func (m *MockCreateNoteService) PutItem(note *modelos.UserNote) error {
	args := m.Called(note)
	return args.Error(0)
}

func TestCreateNoteSuccesstb(t *testing.T) {
	service := notes_impl.DefaultNoteService{}
	mockRepo := new(MockCreateNoteService)
	mockRepo.On("PutItem", mock.Anything).Return(nil)

	body := `{"id":"1", "text":"test note"}`
	response := service.CreateNote(body, mockRepo)
	assert.Equal(t, 200, response.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestCreateNoteFailureUnmarshal(t *testing.T) {
	service := notes_impl.DefaultNoteService{}
	mockRepo := new(MockCreateNoteService)

	body := `{"id":1, "text":}`
	response := service.CreateNote(body, mockRepo)

	assert.Equal(t, 400, response.StatusCode)
}

func TestCreateNoteFailureMissingFields(t *testing.T) {
	service := notes_impl.DefaultNoteService{}
	mockRepo := new(MockCreateNoteService)

	body := `{"id":"", "text":""}`
	response := service.CreateNote(body, mockRepo)

	assert.Equal(t, 400, response.StatusCode)
}
