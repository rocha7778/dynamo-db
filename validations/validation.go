package validations

import "regexp"

func IsValidNoteID(noteID string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(noteID)
}
