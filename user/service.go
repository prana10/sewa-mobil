package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	GetAllUser() ([]User, error)
	GetUserByID(ID int) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
}

type serviceUser struct {
	repository Repository
}

func NewService(repository Repository) *serviceUser {
	return &serviceUser{repository}
}

func (s *serviceUser) RegisterUser(input RegisterUserInput) (User, error) {
	userStruct := User{}
	userStruct.Name = input.Name
	userStruct.Email = input.Email
	userStruct.Role = "user"

	// hashing
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return userStruct, err
	}
	userStruct.PasswordHash = string(hashPassword)

	newUser, err := s.repository.Save(userStruct)
	if err != nil {
		return userStruct, err
	}

	return newUser, nil
}

func (s *serviceUser) LoginUser(input LoginUserInput) (User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *serviceUser) GetAllUser() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *serviceUser) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *serviceUser) IsEmailAvailable(input CheckEmailInput) (bool, error) {
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
