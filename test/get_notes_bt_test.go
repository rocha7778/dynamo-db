package notes_test

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rocha7778/dynamo-db/modelos"
	"github.com/rocha7778/dynamo-db/notes_impl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock de GetNotesRepository
type MockGetNotesRepo struct {
	mock.Mock
}

func (m *MockGetNotesRepo) Scam() (*dynamodb.ScanOutput, error) {
	args := m.Called()
	return args.Get(0).(*dynamodb.ScanOutput), args.Error(1)
}

var mockNotes = []modelos.UserNote{
	{
		ID:   "1",
		Text: "Nota de prueba 1",
	},
	{
		ID:   "2",
		Text: "Nota de prueba 2",
	},
}

// Test positivo para GetNotes
func TestGetNotesSuccess(t *testing.T) {
	mockRepo := new(MockGetNotesRepo)
	service := notes_impl.GetNotesCreateService{Repo: mockRepo}
	items := []map[string]*dynamodb.AttributeValue{}

	for _, note := range mockNotes {
		item, _ := dynamodbattribute.MarshalMap(note)
		items = append(items, item)
	}

	mockScanOutput := &dynamodb.ScanOutput{Items: items} // Aquí se debería simular un resultado válido
	mockRepo.On("Scam").Return(mockScanOutput, nil)
	resp := service.GetNotes()
	assert.Equal(t, 200, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}

// Test negativo para GetNotes - Error al escanear DynamoDB
func TestGetNotesScanError(t *testing.T) {
	mockRepo := new(MockGetNotesRepo)
	service := notes_impl.GetNotesCreateService{Repo: mockRepo}
	mockRepo.On("Scam").Return((*dynamodb.ScanOutput)(nil), errors.New("Error scanning DynamoDB"))
	resp := service.GetNotes()
	assert.Equal(t, 500, resp.StatusCode)
	assert.Contains(t, resp.Body, "Error scanning DynamoDB")
	mockRepo.AssertExpectations(t)
}

// Test negativo para GetNotes - No se encontraron usuarios
func TestGetNotesNoUsersFound(t *testing.T) {
	mockRepo := new(MockGetNotesRepo)
	service := notes_impl.GetNotesCreateService{Repo: mockRepo}
	mockScanOutput := &dynamodb.ScanOutput{Items: []map[string]*dynamodb.AttributeValue{}}
	mockRepo.On("Scam").Return(mockScanOutput, nil)
	resp := service.GetNotes()
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, resp.Body, "Notes not found")
	mockRepo.AssertExpectations(t)
}
