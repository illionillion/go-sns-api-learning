package main

import (
	"fmt"
	"os"

	"github.com/illionillion/go-sns-api-learning/db"
	"github.com/illionillion/go-sns-api-learning/router"
)

func main() {
	fmt.Println("Hello, World!")
	db.NewDB()

	e := router.NewRouter()
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
