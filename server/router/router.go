package router

import (
	"eea/controller"
	"eea/middleware"
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

		api.Use(middleware.Auth)
		//api.GET("/user_info", controller.UserInfo)
		//api.POST("/transferin", controller.TransferInReq)
		//api.POST("/withdraw", controller.WithDrawReq)
		//api.GET("/balance", controller.Balance)

		adm := api.Group("/admin")
		adm.Use(middleware.Admin)
		//adm.POST("/balance/:uid", controller.SetBalance)
		//adm.POST("/invert_strategy/:uid", controller.SetInvertStrategy)
		//adm.POST("/withdraw/:id/confim", controller.WithDrawConfim)
		//adm.POST("/transferin/:id/confim", controller.TransferInConfim)
		//adm.POST("/user_info", controller.UpdateUserInfo)
	}
	r.NoRoute(func(c *gin.Context) {
		c.File("../web/dist/index.html")
	})
	return r
}
