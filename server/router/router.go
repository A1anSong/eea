package router

import (
	"eea/controller"
	"net/http"

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
		//api.Use(Auth)
		api.POST("/rsadecrypt", controller.RSADecrypt)
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "ok"})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.File("../web/dist/index.html")
	})

	return r
}
