package model

import (
	"context"
	"database/sql"
	"errors"

	database "flover-market/db"
)

/*CREATE TABLE IF NOT EXISTS customers
(
    customer_id    INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    customer_name  text                              NOT NULL,
    customer_city TEXT,
    customer_phone TEXT                              NOT NULL
);*/

type Customer struct {
	CustomerId       int64
	CustomerName     string
	CustomerPassword string
	CustomerCity     string
	CustomerPhone    string
}

func GetCustomers() ([]Customer, error) {
	db := database.GetDB()

	rows, err := db.Query(context.Background(), "SELECT customer_id, customer_name, customer_city, customer_phone FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.CustomerId, &customer.CustomerName, &customer.CustomerCity, &customer.CustomerPhone); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomer(id int64) (Customer, error) {
	db := database.GetDB()

	var customer Customer
	err := db.QueryRow(context.Background(), "SELECT customer_name, customer_city, customer_phone FROM customers WHERE customer_id = $1", id).Scan(&customer.CustomerName, &customer.CustomerCity, &customer.CustomerPhone)
	if err != nil {
		return customer, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return customer, errors.New("customer not found")
	}
	return customer, nil
}

func GetCustomerByPhone(phone string) (Customer, error) {
	db := database.GetDB()
	rows, err := db.Query(context.Background(), "SELECT customer_id, customer_name, customer_password, customer_city, customer_phone FROM customers WHERE customer_phone = $1", phone)
	if err != nil {
		return Customer{}, err
	}
	defer rows.Close()
	var customer Customer
	for rows.Next() {
		if err := rows.Scan(&customer.CustomerId, &customer.CustomerName, &customer.CustomerPassword, &customer.CustomerCity, &customer.CustomerPhone); err != nil {
			return Customer{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return Customer{}, err
	}
	if customer.CustomerId == 0 {
		return Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

func AddCustomer(customer Customer) (int64, error) {
	db := database.GetDB()
	var lastId int64
	err := db.QueryRow(context.Background(), "INSERT INTO customers (customer_name, customer_city, customer_phone, customer_password) VALUES ($1, $2, $3, $4)", customer.CustomerName, customer.CustomerCity, customer.CustomerPhone, customer.CustomerPassword).Scan(&lastId)
	if err != nil {
		return 0, err
	}
	return lastId, nil

}

func UpdateCustomer(customer Customer) error {
	db := database.GetDB()
	_, err := db.Exec(context.Background(), "UPDATE customers SET customer_name = $1, customer_city = $2, customer_phone = $3 WHERE customer_id = $4", customer.CustomerName, customer.CustomerCity, customer.CustomerPhone, customer.CustomerId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCustomer(id int) {
	// TODO
}
