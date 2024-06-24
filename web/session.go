package web

import (
	"flover-market/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func setHandlersSession(router *gin.Engine) {
	router.GET("/api/session/check", handlerCheckSession)
}

func handlerCheckSession(c *gin.Context) {
	token, err := c.Cookie("session_token")
	session, err := controller.GetSession(token)
	if err != nil {
		log.Println(err)
		c.AbortWithError(401, err)
		return
	}
	c.JSON(200, session)
}
