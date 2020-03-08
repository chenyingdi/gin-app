package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GeneToken(args map[string]interface{}, secretKey string) (string, error) {
	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24*3)).Unix()
	claims["iat"] = time.Now().Unix()

	if args != nil {
		for k, v := range args {
			claims[k] = v
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = claims

	return token.SignedString([]byte(secretKey))

}
