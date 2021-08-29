package models

import "golang-rest-crud/entities"

type RegisterUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type RegisterUserInputResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func UserFormatter(user entities.User, token string) RegisterUserInputResponse {
	formatter := RegisterUserInputResponse{
		Email:    user.Email,
		Address:  user.Address,
		Password: user.Pwd,
	}

	return formatter
}
