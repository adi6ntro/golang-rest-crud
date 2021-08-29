package services

import (
	"golang-rest-crud/entities"
	"golang-rest-crud/models"
	"golang-rest-crud/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	AddUser(input models.RegisterUserInput) (entities.User, error)
}

type userService struct {
	repository repository.UserRepo
}

func NewService(repository repository.UserRepo) *userService {
	return &userService{repository}
}

func (s *userService) AddUser(input models.RegisterUserInput) (entities.User, error) {
	user := entities.User{}
	user.Email = input.Email
	user.Address = input.Address

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Pwd = string(passwordHash)
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
