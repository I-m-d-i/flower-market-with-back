package model

import (
	"context"
	database "flover-market/db"
	"fmt"
)

type Category struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func GetCategories() ([]Category, error) {
	db := database.GetDB()

	rows, err := db.Query(context.Background(), "SELECT category_id, category_name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.CategoryId, &category.CategoryName); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategory(id int) (Category, error) {
	db := database.GetDB()
	var category Category
	row := db.QueryRow(context.Background(), "SELECT category_id, category_name FROM categories WHERE category_id = $1", id)
	err := row.Scan(&category.CategoryId, &category.CategoryName)
	if err != nil {
		return Category{}, fmt.Errorf("category with id %d not found", id)
	}
	return category, nil
}
