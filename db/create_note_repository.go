package db

import (
	"github.com/rocha7778/dynamo-db/modelos"
)

type CreateNoteRepository interface {
	PutItem(note *modelos.UserNote) error
}
