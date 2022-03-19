package router

import (
	"eea/controller"
	"eea/middleware"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logrusLogger(), gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile("../web/dist", true)))
	api := r.Group("/api")
	{
		api.POST("/login", controller.Login)

		api.Use(middleware.Auth)
		api.GET("/user_info", controller.UserInfo)
		api.POST("/transferin", controller.TransferInReq)
		api.POST("/withdraw", controller.WithDrawReq)
		api.GET("/balance", controller.Balance)

		adm := api.Group("/admin")
		adm.Use(middleware.Admin)
		adm.POST("/balance/:uid", controller.SetBalance)
		adm.GET("/balances", controller.GetBalanceList)
		adm.POST("/invert_strategy/:uid", controller.SetInvestStrategy)
		adm.POST("/transfer/:id/confim", controller.TransferConfim)
		adm.GET("/transfers", controller.GetTransferList)
		adm.POST("/user_info", controller.UpdateUserInfo)
		adm.DELETE("/user_info/:id", controller.DeletetUser)
		adm.GET("/users", controller.GetUserList)

	}
	r.NoRoute(func(c *gin.Context) {
		c.File("../web/dist/index.html")
	})
	return r
}

func logrusLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		log.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
