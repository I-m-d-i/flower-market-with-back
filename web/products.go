package web

import (
	"flover-market/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func setHandlersProduct(router *gin.Engine) {
	router.GET("/api/products/get/all", handlerProductsAll)
}

func handlerProductsAll(c *gin.Context) {
	products, err := controller.GetAllProducts()
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, products)
}
