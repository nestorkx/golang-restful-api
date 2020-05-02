package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetProductById(productId int) (*Product, error) {
	const sql = `SELECT
	products.id,
	products.product_code,
	products.product_name,
	COALESCE ( products.description, '' ) AS description,
	products.standard_cost,
	products.list_price,
	products.category 
	FROM
		products 
	WHERE
		products.id=?`
	row := repo.db.QueryRow(sql, productId)
	product := &Product{}

	err := row.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
	if err != nil {
		panic(err)
	}
	return product, err
}
