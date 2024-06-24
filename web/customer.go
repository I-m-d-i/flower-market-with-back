package web

import (
	"encoding/json"
	"flover-market/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func setHandlersCustomer(router *gin.Engine) {
	router.POST("/api/customer/login", handlerCustomerLogin)
	router.POST("/api/customer/register", handlerCustomerRegister)
	router.POST("/api/customer/logout", handlerCustomerLogout)
	router.GET("/api/customer/profile", handlerCustomerProfile)
	router.POST("/api/customer/update", handlerCustomerUpdate)
}

func handlerCustomerUpdate(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}

	var customer controller.Customer

	err = json.NewDecoder(c.Request.Body).Decode(&customer)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	ses, err := controller.GetSession(token)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	customer.CustomerId = ses.CustomerId
	err = controller.UpdateCustomer(customer)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
}

func handlerCustomerProfile(c *gin.Context) {
	token, err := c.Cookie("session_token")
	session, err := controller.GetSession(token)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}

	customer, err := controller.GetCustomer(session.CustomerId)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	var customerView = struct {
		CustomerId    int64  `json:"customer_id"`
		CustomerName  string `json:"customer_name"`
		CustomerPhone string `json:"customer_phone"`
		CustomerCity  string `json:"customer_city"`
	}{
		customer.CustomerId,
		customer.CustomerName,
		customer.CustomerPhone,
		customer.CustomerCity,
	}

	c.JSON(200, customerView)
}

func handlerCustomerLogin(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	token, customer, err := controller.Login(phone, password)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}

	c.SetCookie("session_token", token, 3600, "/", "", false, false)
	c.JSON(200, customer)
}

func handlerCustomerLogout(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}

	err = controller.Logout(token)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}

	c.SetCookie("session_token", "", -1, "/", "", false, false)
	c.JSON(200, "ok")
}

func handlerCustomerRegister(c *gin.Context) {

	name := c.PostForm("name")
	//email := c.PostForm("email")
	city := c.PostForm("city")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	_, err := controller.Register(controller.Customer{
		CustomerName:     name,
		CustomerCity:     city,
		CustomerPhone:    phone,
		CustomerPassword: password,
	})

	if err != nil {
		log.Println(err)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	// TODO создать сессию
	// TODO вертуть токен сессии

	c.JSON(200, "ok")
}
