package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-restful-api/database"
)

func main() {
	databaseConnection := database.InitDB()

	// Lógica

	defer databaseConnection.Close()
	fmt.Println(databaseConnection)
}
