package logic

import (
	"api-gateway/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartHttpServer(router *gin.Engine) {
	routerInit(router)
	_ = router.Run(fmt.Sprintf(":%d", config.Config.Gateway.Port))
}

func routerInit(router *gin.Engine) {
	// 授权
	router.Use(Authorize())
	imS := router.Group("/im")
	{
		imS.GET("/send", sendMessage)
	}

	userS := router.Group("/user")
	{
		userS.GET("/login", login)
	}
}
