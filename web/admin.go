package web

import (
	"bytes"
	"flover-market/controller"
	"flover-market/env"
	"flover-market/model"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func setHandlersAdmin(router *gin.Engine) {
	apiGroup := router.Group("/api")
	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middlewareAdmin)
	//adminGroup.GET("/api/admin/customers", handlerGetAllCustomers)
	adminGroup.GET("/purchases", handlerGetAllPurchases)
	adminGroup.GET("/products", handlerGetAllProducts)
	adminGroup.POST("/product/add", handlerAddProduct)
	adminGroup.POST("/image/upload", handlerUploadImg)
	//adminGroup.POST("/api/admin/product/update", handlerUpdateProduct)
	//adminGroup.POST("/api/admin/product/delete", handlerDeleteProduct)
}
func middlewareAdmin(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if token == "" || err != nil {
		c.AbortWithError(401, err)
		return
	}
	ses, err := controller.GetSession(token)
	if err != nil {
		c.Abort()
		return
	}
	cust, err := controller.GetCustomer(ses.CustomerId)
	if err != nil {
		c.Abort()
		return
	}
	if cust.CustomerPhone != env.AdminPhone {
		c.AbortWithError(403, err)
		return
	}
	c.Next()
}
func handlerGetAllPurchases(c *gin.Context) {
	purchases, err := controller.GetPurchasesAdmin()
	if err != nil {
		c.Abort()
		return
	}
	c.JSON(200, purchases)
}

func handlerGetAllProducts(c *gin.Context) {
	products, err := controller.GetAllProducts()
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, products)
}

func handlerAddProduct(c *gin.Context) {
	var newProduct = controller.Product{}

	err := c.BindJSON(&newProduct)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, err)
		return
	}

	err = controller.AddProduct(newProduct)
	if err != nil {
		log.Println(err)
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, newProduct)
}

func handlerUploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	// Проверка имени файла на расширение
	if file.Filename[len(file.Filename)-4:] != ".png" && file.Filename[len(file.Filename)-4:] != ".jpg" &&
		file.Filename[len(file.Filename)-5:] != ".jpeg" {
		log.Println("Invalid file type")
		c.Abort()
		return
	}

	err = c.SaveUploadedFile(file, "/img/"+file.Filename)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	ff, err := os.OpenFile("/img/"+file.Filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	defer ff.Close()

	buf := bytes.NewBuffer([]byte{})

	_, err = buf.ReadFrom(ff)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}

	img := model.Img{FileName: file.Filename, ImageData: buf.Bytes()}

	err = model.AddImage(img)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	c.JSON(200, "/images/"+img.FileName)

}
