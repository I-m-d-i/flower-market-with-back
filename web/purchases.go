package web

import (
	"encoding/json"
	"flover-market/controller"
	"flover-market/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func setHandlersPurchases(router *gin.Engine) {
	router.POST("/api/purchases/add", handlerAddPurchase)
	//router.GET("/api/purchase", handlerGetPurchase)
	router.POST("/api/purchase/detail", handlerGetPurchaseDetail)
	//router.GET("/api/purchases/get/all", handlerGetAllPurchases)
	//router.GET("/api/purchases/get/user", handlerGetUserPurchases)
}

func handlerAddPurchase(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if err != nil {
		log.Println(err)
		c.AbortWithError(401, err)
		return
	}

	var purchase = struct {
		StoreId    int   `json:"store_id"`
		ProductsId []int `json:"products"`
	}{}
	// decode json request in purchase
	err = json.NewDecoder(c.Request.Body).Decode(&purchase)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, err)
		return
	}

	fmt.Println(purchase)

	addPurchase, err := controller.AddPurchase(token, purchase.StoreId, purchase.ProductsId)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, addPurchase)
}

func handlerGetPurchaseDetail(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if err != nil {
		log.Println(err)
		c.AbortWithError(401, err)
		return
	}
	purchaseIdstr := c.PostForm("purchase_id")
	fmt.Println("Номер заказа:", purchaseIdstr)
	purchaseId, err := strconv.Atoi(purchaseIdstr)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, err) // TODO not 500
		return
	}
	purchase, items, err := controller.GetDetailsPurchase(token, purchaseId)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, err)
		return
	}
	var DetailPurchase = struct {
		Purchase model.Purchase `json:"purchase"`
		Items    []struct {
			PurchaseId   int    `json:"purchase_id"`
			ProductId    int    `json:"product_id"`
			ProductName  string `json:"product_name"`
			ProductCount int    `json:"product_count"`
			ProductPrice int    `json:"product_price"`
		} `json:"items,omitempty"`
	}{
		Purchase: purchase,
	}

	DetailPurchase.Purchase = purchase
	for _, item := range items {
		product, err := controller.GetProduct(item.ProductId)
		if err != nil {
			log.Println(err)
			c.AbortWithError(500, err)
			return
		}
		DetailPurchase.Items = append(DetailPurchase.Items, struct {
			PurchaseId   int    `json:"purchase_id"`
			ProductId    int    `json:"product_id"`
			ProductName  string `json:"product_name"`
			ProductCount int    `json:"product_count"`
			ProductPrice int    `json:"product_price"`
		}{
			PurchaseId:   item.PurchaseId,
			ProductId:    item.ProductId,
			ProductName:  product.ProductName,
			ProductCount: item.ProductCount,
			ProductPrice: item.ProductPrice,
		})
	}
	c.JSON(200, DetailPurchase)
}
