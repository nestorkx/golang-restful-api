package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"golang-restful-api/database"
	"net/http"
)

var databaseConnection *sql.DB

type Product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	databaseConnection = database.InitDB()
	defer databaseConnection.Close()
	r := chi.NewRouter()

	// Traer todos los productos
	r.Get("/products", AllProucts)

	// Insertar productos
	r.Post("/products", CreateProduct)

	// Actualizar productos
	r.Put("/products/{id}", UpdateProduct)

	// Eliminar productos
	r.Delete("/products/{id}", DeleteProduct)

	http.ListenAndServe(":3000", r)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := databaseConnection.Prepare("DELETE FROM products WHERE id=?")
	catch(err)

	_, er := query.Exec(id)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Successfully Deleted"})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&product)

	query, err := databaseConnection.Prepare("UPDATE products SET product_code=?, description=? WHERE id=?")
	catch(err)

	_, er := query.Exec(product.ProductCode, product.Description, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Update Successfully"})
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)

	query, err := databaseConnection.Prepare("INSERT products SET product_code=?, description=?")
	catch(err)

	_, er := query.Exec(product.ProductCode, product.Description)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func AllProucts(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT
	products.id,
	products.product_code,
	COALESCE(products.description, '') AS description
	FROM
		products
	ORDER BY
		products.id;`
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.Description)
		catch(err)
		products = append(products, product)
	}
	respondwithJSON(w, http.StatusOK, products)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
