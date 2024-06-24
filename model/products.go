package model

import (
	"context"
	db2 "flover-market/db"
	"fmt"
)

type Product struct {
	ProductId   int
	ProductName string
	Description string
	Height      int
	Diameter    string
	ImgSrc      string
	CategoryId  int
}

func GetProducts() ([]Product, error) {
	db := db2.GetDB()

	rows, err := db.Query(context.Background(), "SELECT product_id, product_name, description, height, diameter, img_src, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ProductId, &product.ProductName, &product.Description, &product.Height, &product.Diameter, &product.ImgSrc, &product.CategoryId); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil

}

func GetProduct(id int) (Product, error) {
	db := db2.GetDB()
	var product Product
	row := db.QueryRow(context.Background(), "SELECT product_id, product_name, description, height, diameter, img_src, category_id FROM products WHERE product_id = $1", id)
	err := row.Scan(&product.ProductId, &product.ProductName, &product.Description, &product.Height, &product.Diameter, &product.ImgSrc, &product.CategoryId)
	if err != nil {
		return Product{}, fmt.Errorf("product with id %d not found", id)
	}
	return product, nil
}

func AddProduct(product Product) (int, error) {
	db := db2.GetDB()
	var id int
	fmt.Println(product)
	err := db.QueryRow(context.Background(), `INSERT INTO products (product_name, category_id, img_src, description, height, diameter) VALUES ($1, $2, $3, $4, $5, $6) RETURNING product_id`, product.ProductName, product.CategoryId, product.ImgSrc, product.Description, product.Height, product.Diameter).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func UpdateProduct(product Product) error {
	db := db2.GetDB()
	_, err := db.Exec(context.Background(), "UPDATE products SET product_name = $1, category_id = $2, img_src = $3, description = $4, height = $5, diameter = $6 WHERE product_id = $7", product.ProductName, product.CategoryId, product.ImgSrc, product.Description, product.Height, product.Diameter, product.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int) error {
	db := db2.GetDB()

	_, err := db.Exec(context.Background(), "DELETE FROM products WHERE product_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
