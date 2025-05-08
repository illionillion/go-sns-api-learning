package main

import (
	"fmt"

	"github.com/illionillion/go-sns-api-learning/db"
	"github.com/illionillion/go-sns-api-learning/models"
)

// データベースをマイグレートするための関数
func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&models.User{})
}
