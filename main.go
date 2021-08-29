package main

import (
	"fmt"
	"golang-rest-crud/controllers"
	"golang-rest-crud/db"
	"golang-rest-crud/repository"
	"golang-rest-crud/services"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	db := db.GetDbCon()

	userRepository := repository.NewUserRepo(db)
	userService := services.NewService(userRepository)
	userController := controllers.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("v1/golang-rest-crud/")

	api.POST("/users", userController.AddUser)

	router.Run()
}
