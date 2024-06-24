package controller

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"flover-market/model"
	"fmt"
	"log"
	"regexp"
)

type Customer struct {
	CustomerId       int64  `json:"customer_id"`
	CustomerName     string `json:"customer_name"`
	CustomerPassword string `json:"customer_password"`
	CustomerCity     string `json:"customer_city"`
	CustomerPhone    string `json:"customer_phone"`
}

func GetCustomer(id int64) (Customer, error) {
	customerModel, err := model.GetCustomer(id)
	if err != nil {
		log.Println(err)
		return Customer{}, fmt.Errorf("не удалось получить пользователя")
	}

	customer := convertCustomerFromModel(customerModel)

	return customer, nil
}

func UpdateCustomer(customer Customer) error {
	customerModel := convertCustomerToModel(customer)
	return model.UpdateCustomer(customerModel)
}

func Register(customer Customer) (int64, error) {
	customerModel := convertCustomerToModel(customer)
	if !validatePhone(customer.CustomerPhone) {
		return 0, fmt.Errorf("некорректный номер телефона")
	}
	if !validatePassword(customer.CustomerPassword) {
		return 0, fmt.Errorf("некорректный пароль")
	}
	if customer.CustomerName == "" {
		return 0, fmt.Errorf("имя не может быть пустым")
	}
	if customer.CustomerCity == "" {
		return 0, fmt.Errorf("город не может быть пустым")
	}
	cust, err := model.GetCustomerByPhone(customer.CustomerPhone)
	if err != nil && !errors.Is(err, sql.ErrNoRows) && errors.Is(err, errors.New("customer not found")) {
		return 0, fmt.Errorf("регистраци не удалась: %w", err)
	}
	if cust.CustomerId != 0 {
		return 0, fmt.Errorf("такой пользователь уже существует")
	}

	customerModel.CustomerPassword = hash(customer.CustomerPassword)
	return model.AddCustomer(customerModel)
}

// Login создает новую пользовательскую сессию и возвращает токен сессии
func Login(phone string, password string) (string, Customer, error) {
	var customer Customer
	if !validatePhone(phone) {
		return "", customer, fmt.Errorf("некорректный номер телефона")
	}

	customerModel, err := model.GetCustomerByPhone(phone)
	if err != nil {
		return "", customer, fmt.Errorf("неверный логин или пароль")
	}

	if customerModel.CustomerPassword != hash(password) {
		return "", customer, fmt.Errorf("неверный логин или пароль")
	}

	sessionToken, err := model.CreateSession(customerModel.CustomerId)
	if err != nil {
		return "", customer, fmt.Errorf("не удалось создать сессию")
	}

	customer = convertCustomerFromModel(customerModel)

	return sessionToken, customer, nil
}

// Logout удаляет пользовательскую сессию
func Logout(sessionToken string) error {
	return model.DeleteSession(sessionToken)
}

func validatePhone(phone string) bool {
	compile, err := regexp.Compile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`)
	if err != nil {
		return false
	}
	return compile.MatchString(phone)
}

// TODO Add validation
func validatePassword(password string) bool {
	return len(password) >= 1
}

func hash(s string) string { return fmt.Sprintf("%x", md5.Sum([]byte(s))) }

func convertCustomerFromModel(customer model.Customer) Customer {
	return Customer{
		CustomerId:    customer.CustomerId,
		CustomerName:  customer.CustomerName,
		CustomerCity:  customer.CustomerCity,
		CustomerPhone: customer.CustomerPhone,
	}
}

func convertCustomerToModel(customer Customer) model.Customer {
	return model.Customer{
		CustomerId:    customer.CustomerId,
		CustomerName:  customer.CustomerName,
		CustomerCity:  customer.CustomerCity,
		CustomerPhone: customer.CustomerPhone,
	}
}
