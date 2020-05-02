package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-restful-api/database"
	"golang-restful-api/product"
	"net/http"
)

var databaseConnection *sql.DB

func main() {
	databaseConnection = database.InitDB()
	defer databaseConnection.Close()
	r := chi.NewRouter()

	// Productos
	var productRepository = product.NewRepository(databaseConnection)
	var productService product.Service
	productService = product.NewService(productRepository)

	r.Mount("/products", product.MakeHttpHandler(productService))

	http.ListenAndServe(":3000", r)
}
