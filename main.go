package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"userregistry/controllers"
	mongo2 "userregistry/repo/mongo"
	"userregistry/service"
	validator2 "userregistry/utils/validator"
)

const (
	MongoDBUri = "mongodb://localhost:27017"
	Database   = "testdb"
	Collection = "users"
)

func main() {
	e := echo.New()
	userRepo, err := mongo2.NewUsersRepo(MongoDBUri, Database, Collection)
	if err != nil {
		log.Fatalln("failed to connect to mongodb", err)
	}
	defer userRepo.Close()

	userService := service.NewUserService(userRepo)
	userController := controllers.NewUsersController(userService)

	e.POST("/user", userController.Create)
	e.PATCH("/user/:username", userController.Update)
	e.GET("/user/:username", userController.Get)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = validator2.NewCustomValidator(validator.New())

	log.Println("starting the server")
	e.Logger.Fatal(e.Start(":8080"))
}
