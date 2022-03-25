package models

import d "test_revel/app/db"

const (
	InsertProductSQL  = `INSERT INTO products (name, price) VALUES (?, ?)`
	GetProductSQL     = `SELECT id, name, price FROM products`
	GetProductByIdSQL = GetProductSQL + ` WHERE id=?`
	UpdateProductSQL  = `UPDATE products SET name=?, price=? WHERE id=?`
	DeleteProductSQL  = `DELETE FROM products WHERE id=?`
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func GetProductById(id string) (Product, int) {
	db := d.Connect()
	defer db.Close()

	var product Product
	rows, err := db.Query(GetUserByIdSQL, id)
	if err != nil {
		return product, 0
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return product, 0
		}
	}

	return product, 1
}
