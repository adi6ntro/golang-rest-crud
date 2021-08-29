package services

import (
	"errors"
	"golang-rest-crud/entities"
	"golang-rest-crud/models"
	"golang-rest-crud/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	AddUser(input models.RegisterUserInput) (entities.User, error)
	Login(input models.LoginInput) (entities.User, error)
	GetUserByID(ID int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(input models.UpdateUserInput) (entities.User, error)
	DeleteUser(ID int) (entities.User, error)
}

type userService struct {
	repository repository.UserRepo
}

func NewService(repository repository.UserRepo) *userService {
	return &userService{repository}
}

func (s *userService) AddUser(input models.RegisterUserInput) (entities.User, error) {
	user := entities.User{}
	found, err := s.isEmailAvailable(input.Email)
	if err != nil {
		return user, err
	}

	if !found {
		return user, errors.New("Email sudah terdaftar")
	}

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

func (s *userService) Login(input models.LoginInput) (entities.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Akun tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) isEmailAvailable(email string) (bool, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *userService) GetUserByID(ID int) (entities.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Akun dengan ID tersebut tidak ditemukan")
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]entities.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *userService) UpdateUser(input models.UpdateUserInput) (entities.User, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	user.Email = input.Email
	if input.Address != "" {
		user.Address = input.Address
	}
	if input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return user, err
		}
		user.Pwd = string(passwordHash)
	}

	user.UpdateAt = time.Now()

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *userService) DeleteUser(ID int) (entities.User, error) {
	user, err := s.repository.Delete(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Akun tidak ditemukan")
	}

	return user, nil
}
