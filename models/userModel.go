package models

import (
	"golang-rest-crud/entities"
	"time"
)

type RegisterUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type RegisterUserInputResponse struct {
	Email   string `json:"email"`
	Address string `json:"address"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	UserId   int       `json:"userId" gorm:"column:id"`
	Email    string    `json:"email" gorm:"column:email"`
	Address  string    `json:"address" gorm:"column:address"`
	Password string    `json:"password" gorm:"column:pwd"`
	UpdateAt time.Time `json:"updateAt" gorm:"column:update_at"`
}

type GetUserID struct {
	ID string `uri:"id" binding:"required"`
}

func UserFormatter(user entities.User) RegisterUserInputResponse {
	formatter := RegisterUserInputResponse{
		Email:   user.Email,
		Address: user.Address,
	}

	return formatter
}

func UsersFormatter(user []entities.User) []RegisterUserInputResponse {
	UsersFormatter := []RegisterUserInputResponse{}

	for _, user := range user {
		RegisterUserInputResponse := UserFormatter(user)
		UsersFormatter = append(UsersFormatter, RegisterUserInputResponse)
	}

	return UsersFormatter
}
