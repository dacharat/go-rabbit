package routers

import (
	"github.com/dacharat/go-rabbit/cmd/server/apps/rabbit"
	"github.com/dacharat/go-rabbit/cmd/server/apps/welcome"
	"github.com/dacharat/go-rabbit/config"
	"github.com/gin-gonic/gin"
)

func GenerateRouter() *gin.Engine {
	gin.SetMode(config.GinMode)

	router := gin.New()

	router.GET("/", welcome.WelcomeHandler)
	router.POST("/rabbit", rabbit.PublishHandler)

	return router
}
