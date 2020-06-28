package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
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

func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
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
	ORDER BY
		products.id
	LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}

	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = `SELECT
		COUNT(*) AS total
	FROM
		products`
	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, nil
}

func (repo *repository) InsertProduct(params *getAddProductRequest) (int64, error) {
	const sql = `INSERT INTO products (product_code, product_name, category, description, list_price, standard_cost) VALUES(?,?,?,?,?,?)`
	result, err := repo.db.Exec(sql, params.ProductCode, params.ProductName, params.Category, params.Description, params.ListPrice, params.StandardCost)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) UpdateProduct(params *updateProductRequest) (int64, error) {
	const sql = `UPDATE products 
	SET products.product_code = ?,
		products.product_name = ?,
		products.category = ?,
		products.description = ?,
		products.list_price = ?,
		products.standard_cost = ? 
	WHERE
		products.id = ?`
	_, err := repo.db.Exec(sql, params.ProductCode, params.ProductName, params.Category, params.Description, params.ListPrice, params.StandardCost, params.ID)
	if err != nil {
		panic(err)
	}
	return params.ID, nil
}

func (repo *repository) DeleteProduct(params *deleteProductRequest) (int64, error) {
	const sql = `DELETE FROM products WHERE products.id = ?`
	res, err := repo.db.Exec(sql, params.ProductID)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, nil
}
