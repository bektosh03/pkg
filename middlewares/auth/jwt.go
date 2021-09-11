package auth

import (
	"github.com/golang-jwt/jwt"
)

func (a JwtAuth) ExtractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return a.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
