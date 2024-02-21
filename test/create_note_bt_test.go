package notes_test

import (
	"testing"

	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/notes_impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateNoteRepository struct {
	mock.Mock
}

func (m *MockCreateNoteRepository) PutItem(note *modelos.UserNote) error {
	args := m.Called(note)
	return args.Error(0)
}

func setup() (*notes_impl.CreateNoteService, *MockCreateNoteRepository) {
	mockRepo := new(MockCreateNoteRepository)
	service := &notes_impl.CreateNoteService{
		Repo: mockRepo, // Inyectar el mock aquí
	}
	return service, mockRepo
}

func TestCreateNoteSuccesstb(t *testing.T) {
	service, mockRepo := setup()
	mockRepo.On("PutItem", mock.Anything).Return(nil)
	body := `{"id":"1", "text":"test note"}`
	response := service.CreateNote(body)
	assert.Equal(t, 200, response.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestCreateNoteFailureUnmarshal(t *testing.T) {
	service, _ := setup()
	body := `{"id":1, "text":}`
	response := service.CreateNote(body)
	assert.Equal(t, 400, response.StatusCode)
}

func TestCreateNoteFailureMissingFields(t *testing.T) {
	service, _ := setup()
	body := `{"id":"", "text":""}`
	response := service.CreateNote(body)

	assert.Equal(t, 400, response.StatusCode)
}
