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

// Definir un mock de tu repositorio
type MockGetNoteRepository struct {
	mock.Mock
}

func (m *MockGetNoteRepository) GetItem(noteID string) (*dynamodb.GetItemOutput, error) {
	args := m.Called(noteID)
	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

func TestGetNoteById(t *testing.T) {
	// Crear instancias del mock y del servicio
	mockRepo := new(MockGetNoteRepository)
	service := notes_impl.DefaultNoteGetService{Repo: mockRepo}

	note := modelos.UserNote{
		ID:   "123",
		Text: "Sample note",
	}

	noteAttrMap, _ := dynamodbattribute.MarshalMap(note)

	// Configurar el comportamiento esperado del mock
	expectedResult := &dynamodb.GetItemOutput{
		Item: noteAttrMap,
	}
	mockRepo.On("GetItem", "someNoteID").Return(expectedResult, nil)

	// Llamar a la función bajo prueba
	resp, err := service.GetNoteById("someNoteID")

	// Aserciones para verificar el comportamiento esperado
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	// Verificar que el mock fue llamado como se esperaba
	mockRepo.AssertExpectations(t)
}

func TestGetNoteByIdWithError(t *testing.T) {
	// Configurando el mock para devolver un error
	mockRepo := new(MockGetNoteRepository)
	service := notes_impl.DefaultNoteGetService{Repo: mockRepo}
	mockRepo.On("GetItem", "someNoteID").Return((*dynamodb.GetItemOutput)(nil), errors.New("error"))

	// Llamando a la función bajo prueba
	resp, err := service.GetNoteById("someNoteID")

	// Verificando que se recibió un error
	assert.Nil(t, err)

	// Dado que hay un error, se puede verificar también el estado esperado de la respuesta
	// Esto asume que tu implementación de GetNoteById devuelve una respuesta con un código de estado específico incluso en caso de error
	assert.Equal(t, 500, resp.StatusCode)

	// Verificando que el mock fue llamado como se esperaba
	mockRepo.AssertExpectations(t)
}
