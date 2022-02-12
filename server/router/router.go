package router

import (
	"eea/controller"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("../web/dist", true)))

	api := r.Group("/api")
	{
		api.POST("/login", controller.Login)
		api.POST("/rsadecrypt", controller.RSADecrypt)
	}

	r.NoRoute(func(c *gin.Context) {
		c.File("../web/dist/index.html")
	})

	return r
}
