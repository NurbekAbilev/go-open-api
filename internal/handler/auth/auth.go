package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthEmployee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CustomAuthClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateSignedJWT(id int) (signedString string, err error) {
	secretKey := os.Getenv("JWT_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomAuthClaims{
		Name: "Nurbek",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "Server localhost",
		},
	})

	signedString, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ValidateSignedJWT(token string, claim *CustomAuthClaims) (*jwt.Token, error) {
	secretKey := os.Getenv("JWT_KEY")

	tkn, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return tkn, nil
}
