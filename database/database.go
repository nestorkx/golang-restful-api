package database

import "database/sql"

func InitDB() *sql.DB {
	connectionStr := "root:admin@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionStr)

	if err != nil {
		panic(err.Error()) // Error Handling Manejo de Errores
	}
	return databaseConnection
}
