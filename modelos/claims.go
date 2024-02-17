package modelos

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

type Claim struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.RegisteredClaims
}
