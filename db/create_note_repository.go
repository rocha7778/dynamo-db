package db

import (
	"github.com/rocha7778/dynamo-db/modelos"
)

type CreateNoteRepository interface {
	PutItem(note *modelos.UserNote) error
}

type DeleteServiceRepositoryInterface interface {
	DeleteItem(NoteId string) error
}

type GetNoteRepository interface {
	GetItem(NoteId string) (*modelos.UserNote, error)
}

type UpdateNoteRepository interface {
	UpdateItem(note *modelos.UserNote) error
}
