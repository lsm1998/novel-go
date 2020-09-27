package logic

import (
	"api-gateway/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartHttpServer(router *gin.Engine) {
	// favicon.ico
	router.GET("favicon.ico", func(context *gin.Context) {

	})
	routerInit(router)
	_ = router.Run(fmt.Sprintf(":%d", config.Config.Gateway.Port))
}

func routerInit(router *gin.Engine) {
	// 来访记录
	router.Use(AccessRecord())

	// 限流
	router.Use(LimitHandler())

	// 基础服务
	baseS := router.Group("/base")
	{
		baseS.GET("/info", info)
		baseS.GET("/show/:id", show)
		baseS.GET("/bannerList", bannerList)
		baseS.POST("/login", login)
	}

	// 授权
	router.Use(Authorize())

	imS := router.Group("/im")
	{
		imS.GET("/send", sendMessage)
	}

	userS := router.Group("/user")
	{
		userS.GET("/info/:id", userInfo)
	}
}
