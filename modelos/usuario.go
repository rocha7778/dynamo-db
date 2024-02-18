package modelos

import (
	"errors"
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

func (u *User) ValidEmail() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	return nil
}
