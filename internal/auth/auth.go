package auth

import "errors"

type Credentials struct {
	Login string
	ID    string
}

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

type AbstractUser interface {
	GetID() int
	GetRole() string
}

type AuthProvider interface {
	GenerateToken(cred Credentials) (authToken string, err error)
	CheckAuth(authToken string) (err error)
	GetAuthUser() AbstractUser
}

func InitAuth() AuthProvider {
	return &JWTAuthProvider{}
}
