package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)
type Services struct{
	GenerateToken(userID int) (string, error)
}

type jwtServices struct{}

key := os.Getenv("SECRET_KEY")
var SecretKey []byte(key)

func (s* jwtServices) GenerateToken (userID int) (string, error){
	claim := jwt.MapClaims{}
	claim("user_id") = userID

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	
}