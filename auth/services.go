package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Services interface {
	// GenerateToken(userID int) (string, error)
	GenerateToken(userID int) (string, error)
}

type jwtServices struct {
}

func NewService() *jwtServices {
	return &jwtServices{}
}

var SecretKey []byte

func (s *jwtServices) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	key := os.Getenv("SECRET_KEY")
	SecretKey = []byte(key)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token.SignedString(SecretKey)

	signedToken, err := token.SignedString(SecretKey)

	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
