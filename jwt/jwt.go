package jwt

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	jwt "github.com/golang-jwt/jwt/v5"
	modelos "github.com/rocha7778/dynamo-db/modelos"
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

func ValidateToken(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, *modelos.Claim) {

	path := ctx.Value(modelos.Key("path")).(string)

	if path == "signup" || path == "login" {
		return true, 200, "OK", &modelos.Claim{}
	}

	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "Unauthorized", &modelos.Claim{}
	}

	claim, tokenOk, msg, err := ProcessToken(token, ctx.Value(modelos.Key("JWT_SIGN")).(string))

	if !tokenOk {
		if err != nil {
			return false, 401, "Unauthorized", &modelos.Claim{}
		} else {
			return false, 401, msg, &modelos.Claim{}
		}
	}

	return true, 200, "OK", claim
}

func GenerateJWT(ctx context.Context, t modelos.User) (string, error) {

	jwtSign := "rocha7778"
	miClave := []byte(jwtSign)

	payload := jwt.MapClaims{
		"id":        t.ID,
		"name":      t.Name,
		"last_name": t.LastName,
		"birthday":  t.Birthday,
		"email":     t.Email,
		"password":  t.Password,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
