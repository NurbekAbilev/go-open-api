package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nurbekabilev/go-open-api/internal/repo"
)

type AuthEmployee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CustomAuthClaims struct {
	Login string
	ID    string
	jwt.RegisteredClaims
}

type JWTAuthProvider struct {
	currentUser AbstractUser
}

func (a *JWTAuthProvider) GenerateToken(cred Credentials) (authToken string, err error) {
	token, err := generateSignedJWT(cred)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *JWTAuthProvider) CheckAuth(authToken string) (err error) {
	empl := repo.Employee{
		ID:        "1",
		FirstName: "Nurbek",
	}
	a.currentUser = empl

	claim := CustomAuthClaims{}

	tkn, err := validateSignedJWT(authToken, &claim)
	if err != nil {
		return err
	}

	if !tkn.Valid {
		return ErrInvalidCredentials
	}

	return nil
}

func (a *JWTAuthProvider) GetAuthUser() AbstractUser {
	return a.currentUser
}

func generateSignedJWT(cred Credentials) (signedString string, err error) {
	secretKey := os.Getenv("JWT_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomAuthClaims{
		ID:    cred.ID,
		Login: cred.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Go open api project",
		},
	})

	signedString, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func validateSignedJWT(token string, claim *CustomAuthClaims) (*jwt.Token, error) {
	secretKey := os.Getenv("JWT_KEY")

	tkn, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return tkn, nil
}
