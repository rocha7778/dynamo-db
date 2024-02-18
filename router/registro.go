package router

import (
	"encoding/json"
	"errors"

	"github.com/rocha7778/dynamo-db/modelos"
)

func ruserRegister(body string) error {
	var u modelos.User

	err := json.Unmarshal([]byte(body), &u)

	if err != nil {
		return err
	}

	if len(u.Email) == 0 {
		return errors.New("email is required")
	}

	if len(u.Password) < 6 {
		return errors.New("the password must be at least 6 characters")
	}

	_, userExist, _ := user.FindByEmail(u.Email)

	if userExist {
		return errors.New("user already exists")
	}

	_, status, err := user.save(u)

	if err {
		return errors.New("user already exists")

	}

	return nil

}
