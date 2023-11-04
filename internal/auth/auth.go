package auth

import "errors"

type Credentials struct {
	// username string
	// password string

	// todo temporary solution
	Name string
	ID   int
}

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

type AbstractUser interface {
	GetID() string
	GetRole() string
}

type AuthProvider interface {
	Authorize(cred Credentials) (authToken string, err error)
	CheckAuth(authToken string) (err error)
	GetAuthUser() AbstractUser
}

func InitAuth() AuthProvider {
	return &JWTAuthProvider{}
}
