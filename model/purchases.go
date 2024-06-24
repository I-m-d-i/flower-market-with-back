package model

import (
	"context"
	database "flover-market/db"
)

type Purchase struct {
	PurchaseId   int `json:"purchase_id"`
	CustomerId   int `json:"customer_id"`
	StoreId      int `json:"store_id"`
	PurchaseDate int `json:"purchase_date"`
}

type PurchaseItem struct {
	PurchaseId   int `json:"purchase_id"`
	ProductId    int `json:"product_id"`
	ProductCount int `json:"product_count"`
	ProductPrice int `json:"product_price"`
}

func GetPurchases() ([]Purchase, error) {
	db := database.GetDB()

	rows, err := db.Query(context.Background(), "SELECT purchase_id, customer_id, store_id, purchase_date FROM purchases")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var purchases []Purchase
	for rows.Next() {
		var purchase Purchase
		if err := rows.Scan(&purchase.PurchaseId, &purchase.CustomerId, &purchase.StoreId, &purchase.PurchaseDate); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}
	return purchases, nil
}

func GetPurchase(purchaseId int) (Purchase, error) {
	db := database.GetDB()

	row := db.QueryRow(context.Background(), "SELECT purchase_id, customer_id, store_id, purchase_date FROM purchases WHERE purchase_id = $1", purchaseId)
	var purchase Purchase
	if err := row.Scan(&purchase.PurchaseId, &purchase.CustomerId, &purchase.StoreId, &purchase.PurchaseDate); err != nil {
		return Purchase{}, err
	}
	return purchase, nil
}

func GetItemsPurchase(purchaseId int) ([]PurchaseItem, error) {
	db := database.GetDB()

	rows, err := db.Query(context.Background(), "SELECT purchase_id, product_id, product_count, product_price FROM purchase_items WHERE purchase_id = $1", purchaseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PurchaseItem
	for rows.Next() {
		var item PurchaseItem
		if err := rows.Scan(&item.PurchaseId, &item.ProductId, &item.ProductCount, &item.ProductPrice); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func AddPurchase(purchase Purchase, purchaseItems []PurchaseItem) (Purchase, error) {
	db := database.GetDB()

	err := db.QueryRow(context.Background(), "INSERT INTO purchases (customer_id, store_id, purchase_date) VALUES ($1, $2, $3) RETURNING purchase_id", purchase.CustomerId, purchase.StoreId, purchase.PurchaseDate).Scan(&purchase.PurchaseId)
	if err != nil {
		return Purchase{}, err
	}

	for _, item := range purchaseItems {
		_, err = db.Exec(context.Background(), "INSERT INTO purchase_items (purchase_id, product_id, product_count, product_price) VALUES ($1, $2, $3, $4)", purchase.PurchaseId, item.ProductId, item.ProductCount, item.ProductPrice)
		if err != nil {
			return Purchase{}, err
		}
	}

	return purchase, nil
}
