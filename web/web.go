package web

import (
	"flover-market/model"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"path"
)

func StartGin(port string) {
	router := gin.Default()

	// Serve static files
	router.Use(static.Serve("/", static.LocalFile(path.Join("frontend", "dist"), true)))
	//router.Use(static.Serve("/img", static.LocalFile(".\\img", true)))
	router.GET("/images/:file", handlerGetImg)

	router.LoadHTMLGlob(path.Join(".", "frontend", "dist", "index.html"))
	setHandlersCustomer(router)
	setHandlersSession(router)
	setHandlersProduct(router)
	setHandlersPurchases(router)
	setHandlersAdmin(router)

	router.Run(port)
}

func handlerGetImg(c *gin.Context) {
	imgName := c.Param("file")
	// Get img
	img, err := model.GetImg(imgName)
	if err != nil {
		c.Abort()
		return
	}
	c.Data(200, "image/png", img.ImageData)
}

// проверяет токен сессии в куки на валидность
func handlerSession(c *gin.Context) {
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		c.Abort()
		return
	}
	_, err = model.GetSession(sessionToken)
	if err != nil {
		c.Abort()
		return
	}
	c.Next()
}
