package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Services interface {
	// GenerateToken(userID int) (string, error)
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtServices) ValidateToken(encodeToken string) (*jwt.Token, error) {
	key := os.Getenv("SECRET_KEY")
	SecretKey = []byte(key)
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
