package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenStr, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token error: token is invalid")
	}

	return token.Claims.(jwt.MapClaims), nil
}
