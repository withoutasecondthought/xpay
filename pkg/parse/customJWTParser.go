package parse

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"xpay/pkg/service"
)

func CustomJWTParser(token string) (int, error) {
	ltok, err := jwt.ParseWithClaims(token, &service.CustomJWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := ltok.Claims.(*service.CustomJWTClaims); ok && ltok.Valid {
		return claims.Id, nil
	}

	return 0, errors.New("invalid access token")
}
