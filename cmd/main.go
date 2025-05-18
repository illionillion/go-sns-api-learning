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
	// swaggerのリンクのログ出力
	domain := "http://" + os.Getenv("API_DOMAIN")
	port := os.Getenv("PORT")
	fmt.Printf("Swagger UI: %s:%s/swagger/index.html\n", domain, port)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":" + port))
}
