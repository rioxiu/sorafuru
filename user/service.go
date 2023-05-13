package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	//mapping struct input ke struct user
	// menyimpan struct user melalui repository

	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Fullname = input.Fullname
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error")
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}
	return newUser, nil

}
