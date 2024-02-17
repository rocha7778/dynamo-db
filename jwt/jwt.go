package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/rocha7778/dynamo-db/modelos"
)

var Email string
var IdUsuario string

func ProcessToken(tk string, Jwtsing string) (*modelos.Claim, bool, string, error) {

	miClave := []byte(Jwtsing)

	var claims modelos.Claim

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Token format is incorrect")

	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		// 	TODO: validate the token
	}

	if tkn.Valid {
		return &claims, false, string(""), errors.New("Invalid token")
	}

	return &claims, false, string(""), nil

}
