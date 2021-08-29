package controllers

import (
	"golang-rest-crud/models"
	"golang-rest-crud/services"
	"golang-rest-crud/utils"
	"net/http"
	"strconv"

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
		errorMessage := gin.H{"errors": err.Error()}
		response := utils.APIResponse("Tidak Bisa Menambah Akun", http.StatusBadRequest, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := models.UserFormatter(users)

	response := utils.APIResponse("Akun baru berhasil ditambahkan", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var input models.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := utils.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// token, err := h.authService.GenerateToken(loggedinUser.ID)
	// if err != nil {
	// 	response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	formatter := models.UserFormatter(loggedinUser)

	response := utils.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userController) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		response := utils.APIResponse("Error to get Users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("List of Users", http.StatusOK, "success", models.UsersFormatter(users))
	c.JSON(http.StatusOK, response)
}

func (h *userController) GetUser(c *gin.Context) {
	id := c.Query("id")
	userid, _ := strconv.Atoi(id)

	user, err := h.userService.GetUserByID(userid)
	if err != nil {
		response := utils.APIResponse("Failed to get user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("", http.StatusOK, "success", models.UserFormatter(user))
	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateUser(c *gin.Context) {
	var input models.UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.UpdateUser(input)
	if err != nil {
		response := utils.APIResponse("Failed to update user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update user", http.StatusOK, "success", models.UserFormatter(user))
	c.JSON(http.StatusOK, response)
}

func (h *userController) DeleteUser(c *gin.Context) {
	id := c.Query("id")
	userid, _ := strconv.Atoi(id)

	_, err := h.userService.DeleteUser(userid)
	if err != nil {
		response := utils.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errorMessage := gin.H{"note": "success deleted"}
	response := utils.APIResponse("", http.StatusOK, "success", errorMessage)
	c.JSON(http.StatusOK, response)
}
