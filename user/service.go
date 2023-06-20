package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	CheckEmail(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

type service struct {
	//mapping struct input ke struct user
	// menyimpan struct user melalui repository
	repository Repository
}

type UserResult struct {
	User *User
	Err  error
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Fullname
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

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckEmail(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) *UserResult {
	// mencari user berdasarkan ID yang tersimpan
	// update attribute avatar file name
	// simpan perubahan avatar file name

	user, err := s.repository.FindById(ID)
	if err != nil {
		return &UserResult{
			User: &user,
			Err:  err,
		}
	}

	user.Avatar_filename = fileLocation
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return &UserResult{
			User: &updatedUser,
			Err:  err,
		}
	}

	return nil

	// user.Avatar_filename = fileLocation
	// updatedUser, err := s.repository.Update(user)
	// if err != nil {
	// 	return updatedUser, err
	// }
	// return updatedUser, nil
}
