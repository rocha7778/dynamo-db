package modelos

import (
	"time"
)

type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Birthday time.Time `json:"birthday"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
