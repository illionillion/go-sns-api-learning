package main

import (
	"fmt"
	"os"

	"github.com/illionillion/go-sns-api-learning/controller"
	"github.com/illionillion/go-sns-api-learning/db"
	"github.com/illionillion/go-sns-api-learning/repository"
	"github.com/illionillion/go-sns-api-learning/router"
	"github.com/illionillion/go-sns-api-learning/usecase"
)

func main() {
	fmt.Println("Hello, World!")
	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
