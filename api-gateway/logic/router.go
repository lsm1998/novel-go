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

func routerInit(route *gin.Engine) {
	imS := route.Group("/im")
	{
		imS.GET("/send", sendMessage)
	}

	userS := route.Group("/user")
	{
		userS.GET("/login", login)
	}
}
