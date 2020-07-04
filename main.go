package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-restful-api/database"
	"golang-restful-api/employee"
	"golang-restful-api/product"
	"net/http"
)

var databaseConnection *sql.DB

func main() {
	databaseConnection = database.InitDB()
	defer databaseConnection.Close()
	r := chi.NewRouter()

	// Productos
	var (
		productRepository = product.NewRepository(databaseConnection)
		productService    product.Service
	)
	productService = product.NewService(productRepository)

	// Empleados
	var (
		employeeRepository = employee.NewRepository(databaseConnection)
		employeeService    employee.Service
	)
	employeeService = employee.NewService(employeeRepository)

	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))

	http.ListenAndServe(":3000", r)
}
