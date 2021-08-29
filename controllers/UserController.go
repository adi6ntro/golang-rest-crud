package controllers

import (
	"golang-rest-crud/models"
	"golang-rest-crud/services"
	"golang-rest-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userController {
	return &userController{userService}
}

func (h *userController) AddUser(c *gin.Context) {
	var input models.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Tidak Bisa Menambah Akun", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, err := h.userService.AddUser(input)
	if err != nil {
		response := utils.APIResponse("Tidak Bisa Menambah Akun", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := models.UserFormatter(users, "token")

	response := utils.APIResponse("Akun baru berhasil ditambahkan", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}
